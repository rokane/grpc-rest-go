# grpc-rest-template-go

This repository shows an example of how to expose a gRPC API alongside a compatible RESTful API using the grpc-gateway.

## Generating Protobuf Definitions

To generate the golang protobuf definitions contained in `pkg/pingv1/ping_api.pb.go`

```bash
protoc -I $PATH_TO_GOOGLEAPIS -I api/proto/ api/proto/ping_api.proto --go_out=plugins=grpc:pkg/pingv1
```

To generate reverse proxy allowing translation from gRPC to REST contained in `pkg/pingv1/ping_api.pb.gw.go`

```bash
protoc -I $PATH_TO_GOOGLEAPIS -I api/proto/ api/proto/ping_api.proto --grpc-gateway_out=logtostderr=true:pkg/pingv1
```

To generate swagger definitions matching the RESTful API contained in `api/rest/ping_api.swagger.json`

```bash
protoc -I $PATH_TO_GOOGLEAPIS -I api/proto/ api/proto/ping_api.proto --swagger_out=logtostderr=true:api/rest
```
