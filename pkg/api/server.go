package api

import (
	"context"
	"io"
	"log"

	pv1 "github.com/rokane/grpc-rest-template-go/pkg/pingv1"
)

// PingServer implements the PingAPIServer interface contained in
// pkg/pingv1/ping_api.pb.go
type PingServer struct {
	pv1.UnimplementedPingAPIServer
}

// NewPingServer generates and initialised a new PingServer.
func NewPingServer() *PingServer {
	return &PingServer{}
}

// Ping receives an empty request from a client, and response with a Pong msg.
func (ps *PingServer) Ping(ctx context.Context, req *pv1.PingRequest) (*pv1.PingResponse, error) {
	log.Println("Received Ping Request:")
	resp := pv1.PingResponse{Message: "Pong"}
	log.Println("Sending Ping Response:", resp.GetMessage())
	return &resp, nil
}

// PingStream processes a stream of PingRequests and responds with the count
// of requests it processed.
func (ps *PingServer) PingStream(stream pv1.PingAPI_PingStreamServer) error {
	count := 0
	for {
		// Receive request details from stream
		req, err := stream.Recv()

		// If end of stream, return result
		if err == io.EOF {
			log.Println("Sending PingStream Response:", count)
			return stream.SendAndClose(&pv1.PingStreamResponse{
				Count: int32(count),
			})
		}

		// Check for stream error
		if err != nil {
			return err
		}

		// Process request
		log.Println("Received PingStream Request:", req.GetMessage(), req.GetId())
		count++
	}
}

// PongStream receives a request stating information on how long to stream
// a response for, and proceeds to send a streaming response.
func (ps *PingServer) PongStream(req *pv1.PongStreamRequest, stream pv1.PingAPI_PongStreamServer) error {
	log.Println("Received PongStream Request:", req.GetCount())
	for i := 1; i <= int(req.Count); i++ {
		resp := pv1.PongStreamResponse{
			Id:      int32(i),
			Message: "Pong ...",
		}
		// Send response on stream
		log.Println("Sending PongStream Response:", resp.GetMessage(), resp.GetId())
		if err := stream.Send(&resp); err != nil {
			return err
		}
	}
	return nil
}

// PingPongStream receives and sends streams bidirectionally to a client.
func (ps *PingServer) PingPongStream(stream pv1.PingAPI_PingPongStreamServer) error {
	count := 0
	for {
		// Receive request from stream
		req, err := stream.Recv()

		// If end of stream, terminate
		if err == io.EOF {
			return nil
		}

		// Check for error on stream
		if err != nil {
			return err
		}

		// Process response
		log.Println("Received PingPongStream Request:",
			req.GetMessage(), req.GetId())

		count++
		resp := pv1.PingPongResponse{
			Id:      int32(count),
			Message: "Pong ...",
		}

		log.Println("Sending PingPongStream Response:",
			resp.GetMessage(), resp.GetId())

		// Send response on stream
		if err := stream.Send(&resp); err != nil {
			return err
		}
	}
}
