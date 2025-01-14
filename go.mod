module github.com/cs3org/reva

require (
	bou.ke/monkey v1.0.2
	contrib.go.opencensus.io/exporter/jaeger v0.2.1
	contrib.go.opencensus.io/exporter/prometheus v0.3.0
	github.com/BurntSushi/toml v0.3.1
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible
	github.com/ReneKroon/ttlcache/v2 v2.7.0
	github.com/asaskevich/govalidator v0.0.0-20200428143746-21a406dcc535 // indirect
	github.com/aws/aws-sdk-go v1.40.0
	github.com/beevik/etree v1.1.0
	github.com/bluele/gcache v0.0.2
	github.com/c-bata/go-prompt v0.2.5
	github.com/cheggaaa/pb v1.0.29
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/cs3org/cato v0.0.0-20200828125504-e418fc54dd5e
	github.com/cs3org/go-cs3apis v0.0.0-20210726064751-4e567a2c8df1
	github.com/eventials/go-tus v0.0.0-20200718001131-45c7ec8f5d59
	github.com/gdexlab/go-render v1.0.1
	github.com/go-ldap/ldap/v3 v3.3.0
	github.com/go-openapi/errors v0.20.0 // indirect
	github.com/go-openapi/strfmt v0.19.5 // indirect
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang-jwt/jwt v3.2.1+incompatible
	github.com/golang/mock v1.4.4 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v1.8.5
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/huandu/xstrings v1.3.0 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/jedib0t/go-pretty v4.3.0+incompatible
	github.com/mattn/go-sqlite3 v1.14.8
	github.com/minio/minio-go/v7 v7.0.12
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.4.1
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.14.0
	github.com/pkg/errors v0.9.1
	github.com/pkg/xattr v0.4.3
	github.com/pquerna/cachecontrol v0.0.0-20180517163645-1555304b9b35 // indirect
	github.com/rs/cors v1.8.0
	github.com/rs/zerolog v1.23.0
	github.com/sciencemesh/meshdirectory-web v1.0.4
	github.com/sethvargo/go-password v0.2.0
	github.com/stretchr/testify v1.7.0
	github.com/studio-b12/gowebdav v0.0.0-20200303150724-9380631c29a1
	github.com/tidwall/pretty v1.1.0 // indirect
	github.com/tus/tusd v1.1.1-0.20200416115059-9deabf9d80c2
	go.mongodb.org/mongo-driver v1.5.1 // indirect
	go.opencensus.io v0.23.0
	golang.org/x/crypto v0.0.0-20201216223049-8b5274cf687f
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sys v0.0.0-20210423082822-04245dca01da
	golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98 // indirect
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	gotest.tools v2.2.0+incompatible
)

go 1.16

replace (
	github.com/eventials/go-tus => github.com/andrewmostello/go-tus v0.0.0-20200314041820-904a9904af9a
	github.com/oleiade/reflections => github.com/oleiade/reflections v1.0.1
	google.golang.org/grpc => google.golang.org/grpc v1.26.0 // temporary downgrade
)
