# grpc-rest-go

This repository shows an example gRPC service contains set of methods to demonstrate  basic RPC as well as server, client and bidirection streaming RPC's.

The server exposes a gRPC API as well as a compatible REStful API using the grpc-gateway. Note, the RESTful API is not compatible with streaming RPC's and will only work on basic RPC calls.

A sample command line client is also provided to enable the sending of gRPC requests to the server.

## Pre Requisites

* go 1.12+
* protoc 3.0.0+

## Install

To build the executables `pingserver` and `pingclient` run the following commands

```bash
go build ./cmd/...
```

## Try it out

Running the below command will start a gRPC server listening on `localhost:8080` and a HTTP server listening on `http://localhost:8081`.

```bash
pingserver
```

With the above running, you can then connect to the server using the `pingclient` cli tool. The below commands invoke their corresponding methods on the server and log the responses.

### Invoke the Ping RPC

```bash
pingclient ping

> 2019/09/15 14:56:46 Sending Ping Request:
> 2019/09/15 14:56:46 Received Ping Response: Pong
```

### Invoke the PingStream RPC

```bash
pingclient pingstream 3

> 2019/09/15 14:58:15 Sending PingStream Request: Ping ... 1
> 2019/09/15 14:58:15 Sending PingStream Request: Ping ... 2
> 2019/09/15 14:58:15 Sending PingStream Request: Ping ... 3
> 2019/09/15 14:58:15 Received PingStream Response: 3
```

### Invoke the PongStream RPC

```bash
pingclient pongstream 3

> 2019/09/15 14:59:11 Sending PongStream Request: 3
> 2019/09/15 14:59:11 Received PongStream Response: Pong ... 1
> 2019/09/15 14:59:11 Received PongStream Response: Pong ... 2
> 2019/09/15 14:59:11 Received PongStream Response: Pong ... 3
```

### Invoke the PingPongStream RPC

```bash
pingclient pingpongstream 3

> 2019/09/15 14:59:54 Sending PingPongStream Request: Ping ... 1
> 2019/09/15 14:59:54 Sending PingPongStream Request: Ping ... 2
> 2019/09/15 14:59:54 Sending PingPongStream Request: Ping ... 3
> 2019/09/15 14:59:54 Received PingPongStream Response: Pong ... 1
> 2019/09/15 14:59:54 Received PingPongStream Response: Pong ... 2
> 2019/09/15 14:59:54 Received PingPongStream Response: Pong ... 3
```

### Sending a HTTP Request through the Reverse Proxy gRPC Gateway

```bash
curl http://localhost:8081/v1/ping
```

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
