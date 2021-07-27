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

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	node "github.com/cs3org/reva/pkg/storage/utils/decomposedfs/node"

	os "os"

	providerv1beta1 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
)

// Tree is an autogenerated mock type for the Tree type
type Tree struct {
	mock.Mock
}

// CreateDir provides a mock function with given fields: ctx, _a1
func (_m *Tree) CreateDir(ctx context.Context, _a1 *node.Node) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *node.Node) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, _a1
func (_m *Tree) Delete(ctx context.Context, _a1 *node.Node) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *node.Node) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBlob provides a mock function with given fields: key
func (_m *Tree) DeleteBlob(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMD provides a mock function with given fields: ctx, _a1
func (_m *Tree) GetMD(ctx context.Context, _a1 *node.Node) (os.FileInfo, error) {
	ret := _m.Called(ctx, _a1)

	var r0 os.FileInfo
	if rf, ok := ret.Get(0).(func(context.Context, *node.Node) os.FileInfo); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(os.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *node.Node) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPathByID provides a mock function with given fields: ctx, id
func (_m *Tree) GetPathByID(ctx context.Context, id *providerv1beta1.Reference) (string, error) {
	ret := _m.Called(ctx, id)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *providerv1beta1.Reference) string); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *providerv1beta1.Reference) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListFolder provides a mock function with given fields: ctx, _a1
func (_m *Tree) ListFolder(ctx context.Context, _a1 *node.Node) ([]*node.Node, error) {
	ret := _m.Called(ctx, _a1)

	var r0 []*node.Node
	if rf, ok := ret.Get(0).(func(context.Context, *node.Node) []*node.Node); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*node.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *node.Node) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Move provides a mock function with given fields: ctx, oldNode, newNode
func (_m *Tree) Move(ctx context.Context, oldNode *node.Node, newNode *node.Node) error {
	ret := _m.Called(ctx, oldNode, newNode)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *node.Node, *node.Node) error); ok {
		r0 = rf(ctx, oldNode, newNode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Propagate provides a mock function with given fields: ctx, _a1
func (_m *Tree) Propagate(ctx context.Context, _a1 *node.Node) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *node.Node) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PurgeRecycleItemFunc provides a mock function with given fields: ctx, key
func (_m *Tree) PurgeRecycleItemFunc(ctx context.Context, key string) (*node.Node, func() error, error) {
	ret := _m.Called(ctx, key)

	var r0 *node.Node
	if rf, ok := ret.Get(0).(func(context.Context, string) *node.Node); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*node.Node)
		}
	}

	var r1 func() error
	if rf, ok := ret.Get(1).(func(context.Context, string) func() error); ok {
		r1 = rf(ctx, key)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(func() error)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, key)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ReadBlob provides a mock function with given fields: key
func (_m *Tree) ReadBlob(key string) (io.ReadCloser, error) {
	ret := _m.Called(key)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string) io.ReadCloser); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestoreRecycleItemFunc provides a mock function with given fields: ctx, key
func (_m *Tree) RestoreRecycleItemFunc(ctx context.Context, key string) (*node.Node, func() error, error) {
	ret := _m.Called(ctx, key)

	var r0 *node.Node
	if rf, ok := ret.Get(0).(func(context.Context, string) *node.Node); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*node.Node)
		}
	}

	var r1 func() error
	if rf, ok := ret.Get(1).(func(context.Context, string) func() error); ok {
		r1 = rf(ctx, key)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(func() error)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, key)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Setup provides a mock function with given fields: owner
func (_m *Tree) Setup(owner string) error {
	ret := _m.Called(owner)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(owner)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteBlob provides a mock function with given fields: key, reader
func (_m *Tree) WriteBlob(key string, size int64, reader io.Reader) error {
	ret := _m.Called(key, reader)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, io.Reader) error); ok {
		r0 = rf(key, reader)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
