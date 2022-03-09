module github.com/percona/pmm

go 1.16

replace github.com/go-openapi/spec => github.com/Percona-Lab/spec v0.19.8-percona

require (
	github.com/go-openapi/errors v0.19.6
	github.com/go-openapi/jsonreference v0.19.4 // indirect
	github.com/go-openapi/runtime v0.19.20
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.9
	github.com/go-openapi/validate v0.19.10
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.2.0
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.1.0
	github.com/prometheus/common v0.10.0 // indirect
	github.com/prometheus/procfs v0.1.3 // indirect
	github.com/stretchr/testify v1.7.0
	go.mongodb.org/mongo-driver v1.7.1
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007
	google.golang.org/genproto v0.0.0-20210207032614-bba0dbe2a9ea
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/reform.v1 v1.5.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
