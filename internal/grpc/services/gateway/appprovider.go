// Copyright 2018-2021 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

package gateway

import (
	"context"
	"crypto/tls"
	"net/url"
	"strings"

	providerpb "github.com/cs3org/go-cs3apis/cs3/app/provider/v1beta1"
	registry "github.com/cs3org/go-cs3apis/cs3/app/registry/v1beta1"
	gateway "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	ocmprovider "github.com/cs3org/go-cs3apis/cs3/ocm/provider/v1beta1"
	rpc "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	storageprovider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	typespb "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
	"github.com/cs3org/reva/pkg/appctx"
	"github.com/cs3org/reva/pkg/errtypes"
	"github.com/cs3org/reva/pkg/rgrpc/status"
	"github.com/cs3org/reva/pkg/rgrpc/todo/pool"
	"github.com/cs3org/reva/pkg/token"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func (s *svc) OpenInApp(ctx context.Context, req *gateway.OpenInAppRequest) (*providerpb.OpenInAppResponse, error) {
	p, st := s.getPath(ctx, req.Ref)
	if st.Code != rpc.Code_CODE_OK {
		if st.Code == rpc.Code_CODE_NOT_FOUND {
			return &providerpb.OpenInAppResponse{
				Status: status.NewNotFound(ctx, "gateway: resource not found:"+req.Ref.String()),
			}, nil
		}
		return &providerpb.OpenInAppResponse{
			Status: st,
		}, nil
	}

	if s.isSharedFolder(ctx, p) {
		return &providerpb.OpenInAppResponse{
			Status: status.NewInvalid(ctx, "gateway: can't open shares folder"),
		}, nil
	}

	resName, resChild := p, ""
	if s.isShareChild(ctx, p) {
		resName, resChild = s.splitShare(ctx, p)
	}

	statRes, err := s.stat(ctx, &storageprovider.StatRequest{
		Ref: &storageprovider.Reference{Path: resName},
	})
	if err != nil {
		return &providerpb.OpenInAppResponse{
			Status: status.NewInternal(ctx, err, "gateway: error calling Stat on the resource path for the app provider: "+req.Ref.GetPath()),
		}, nil
	}
	if statRes.Status.Code != rpc.Code_CODE_OK {
		err := status.NewErrorFromCode(statRes.Status.GetCode(), "gateway")
		return &providerpb.OpenInAppResponse{
			Status: status.NewInternal(ctx, err, "Stat failed on the resource path for the app provider: "+req.Ref.GetPath()),
		}, nil
	}

	fileInfo := statRes.Info

	// The file is a share
	if fileInfo.Type == storageprovider.ResourceType_RESOURCE_TYPE_REFERENCE {
		uri, err := url.Parse(fileInfo.Target)
		if err != nil {
			return &providerpb.OpenInAppResponse{
				Status: status.NewInternal(ctx, err, "gateway: error parsing target uri: "+fileInfo.Target),
			}, nil
		}
		if uri.Scheme == "webdav" {
			insecure, skipVerify := getGRPCConfig(req.Opaque)
			return s.openFederatedShares(ctx, fileInfo.Target, req.ViewMode, req.App, insecure, skipVerify, resChild)
		}

		res, err := s.Stat(ctx, &storageprovider.StatRequest{
			Ref: req.Ref,
		})
		if err != nil {
			return &providerpb.OpenInAppResponse{
				Status: status.NewInternal(ctx, err, "gateway: error calling Stat on the resource path for the app provider: "+req.Ref.GetPath()),
			}, nil
		}
		if res.Status.Code != rpc.Code_CODE_OK {
			err := status.NewErrorFromCode(res.Status.GetCode(), "gateway")
			return &providerpb.OpenInAppResponse{
				Status: status.NewInternal(ctx, err, "Stat failed on the resource path for the app provider: "+req.Ref.GetPath()),
			}, nil
		}
		fileInfo = res.Info
	}
	return s.openLocalResources(ctx, fileInfo, req.ViewMode, req.App)
}

func (s *svc) OpenFileInAppProvider(ctx context.Context, req *gateway.OpenFileInAppProviderRequest) (*providerpb.OpenFileInAppProviderResponse, error) {
	// TODO to be removed in a future PR
	return nil, errtypes.NotSupported("Deprecated")
}

func (s *svc) openFederatedShares(ctx context.Context, targetURL string, vm gateway.OpenInAppRequest_ViewMode, app string,
	insecure, skipVerify bool, nameQueries ...string) (*providerpb.OpenInAppResponse, error) {
	log := appctx.GetLogger(ctx)
	targetURL, err := appendNameQuery(targetURL, nameQueries...)
	if err != nil {
		return nil, err
	}
	ep, err := s.extractEndpointInfo(ctx, targetURL)
	if err != nil {
		return nil, err
	}

	ref := &storageprovider.Reference{Path: ep.filePath}
	appProviderReq := &gateway.OpenInAppRequest{
		Ref:      ref,
		ViewMode: vm,
		App:      app,
	}

	meshProvider, err := s.GetInfoByDomain(ctx, &ocmprovider.GetInfoByDomainRequest{
		Domain: ep.endpoint,
	})
	if err != nil {
		return nil, errors.Wrap(err, "gateway: error calling GetInfoByDomain")
	}
	var gatewayEP string
	for _, s := range meshProvider.ProviderInfo.Services {
		if strings.ToLower(s.Endpoint.Type.Name) == "gateway" {
			gatewayEP = s.Endpoint.Path
		}
	}
	log.Debug().Msgf("Forwarding OpenInApp request to: %s", gatewayEP)

	conn, err := getConn(gatewayEP, insecure, skipVerify)
	if err != nil {
		err = errors.Wrap(err, "gateway: error connecting to remote reva")
		return &providerpb.OpenInAppResponse{
			Status: status.NewInternal(ctx, err, "error error connecting to remote reva"),
		}, nil
	}

	gatewayClient := gateway.NewGatewayAPIClient(conn)
	remoteCtx := token.ContextSetToken(context.Background(), ep.token)
	remoteCtx = metadata.AppendToOutgoingContext(remoteCtx, token.TokenHeader, ep.token)

	res, err := gatewayClient.OpenInApp(remoteCtx, appProviderReq)
	if err != nil {
		log.Err(err).Msg("error reaching remote reva")
		return nil, errors.Wrap(err, "gateway: error calling OpenInApp")
	}
	return res, nil
}

func (s *svc) openLocalResources(ctx context.Context, ri *storageprovider.ResourceInfo,
	vm gateway.OpenInAppRequest_ViewMode, app string) (*providerpb.OpenInAppResponse, error) {

	accessToken, ok := token.ContextGetToken(ctx)
	if !ok || accessToken == "" {
		return &providerpb.OpenInAppResponse{
			Status: status.NewUnauthenticated(ctx, errtypes.InvalidCredentials("Access token is invalid or empty"), ""),
		}, nil
	}

	provider, err := s.findAppProvider(ctx, ri, app)
	if err != nil {
		err = errors.Wrap(err, "gateway: error calling findAppProvider")
		var st *rpc.Status
		if _, ok := err.(errtypes.IsNotFound); ok {
			st = status.NewNotFound(ctx, "app provider not found")
		} else {
			st = status.NewInternal(ctx, err, "error searching for app provider")
		}
		return &providerpb.OpenInAppResponse{
			Status: st,
		}, nil
	}

	appProviderClient, err := pool.GetAppProviderClient(provider.Address)
	if err != nil {
		err = errors.Wrap(err, "gateway: error calling GetAppProviderClient")
		return &providerpb.OpenInAppResponse{
			Status: status.NewInternal(ctx, err, "error getting appprovider client"),
		}, nil
	}

	appProviderReq := &providerpb.OpenInAppRequest{
		ResourceInfo: ri,
		ViewMode:     providerpb.OpenInAppRequest_ViewMode(vm),
		AccessToken:  accessToken,
	}

	res, err := appProviderClient.OpenInApp(ctx, appProviderReq)
	if err != nil {
		return nil, errors.Wrap(err, "gateway: error calling OpenInApp")
	}

	return res, nil
}

func (s *svc) findAppProvider(ctx context.Context, ri *storageprovider.ResourceInfo, app string) (*registry.ProviderInfo, error) {
	c, err := pool.GetAppRegistryClient(s.c.AppRegistryEndpoint)
	if err != nil {
		err = errors.Wrap(err, "gateway: error getting appregistry client")
		return nil, err
	}

	if app == "" {
		// We need to get the default provider in case app is not set
		// If the default isn't set as well, we'll return the first provider which matches the mimetype
		res, err := c.GetDefaultAppProviderForMimeType(ctx, &registry.GetDefaultAppProviderForMimeTypeRequest{
			MimeType: ri.MimeType,
		})
		if err == nil && res.Status.Code == rpc.Code_CODE_OK && res.Provider != nil {
			return res.Provider, nil
		}
	}

	res, err := c.GetAppProviders(ctx, &registry.GetAppProvidersRequest{
		ResourceInfo: ri,
	})
	if err != nil {
		err = errors.Wrap(err, "gateway: error calling GetAppProviders")
		return nil, err
	}
	if res.Status.Code != rpc.Code_CODE_OK {
		if res.Status.Code == rpc.Code_CODE_NOT_FOUND {
			return nil, errtypes.NotFound("gateway: app provider not found for resource: " + ri.String())
		}
		return nil, errtypes.InternalError("gateway: error finding app providers")
	}

	if app != "" {
		for _, p := range res.Providers {
			if p.Name == app {
				return p, nil
			}
		}
		return nil, errtypes.NotFound("gateway: app provider not found: " + app)
	}

	// As a fallback, return the first provider in the list
	return res.Providers[0], nil
}

func getGRPCConfig(opaque *typespb.Opaque) (bool, bool) {
	if opaque == nil {
		return false, false
	}
	_, insecure := opaque.Map["insecure"]
	_, skipVerify := opaque.Map["skip-verify"]
	return insecure, skipVerify
}

func getConn(host string, insecure, skipverify bool) (*grpc.ClientConn, error) {
	if insecure {
		return grpc.Dial(host, grpc.WithInsecure())
	}

	// TODO(labkode): if in the future we want client-side certificate validation,
	// we need to load the client cert here
	tlsconf := &tls.Config{InsecureSkipVerify: skipverify}
	creds := credentials.NewTLS(tlsconf)
	return grpc.Dial(host, grpc.WithTransportCredentials(creds))
}
