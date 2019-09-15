package cmd

import (
	"context"
	"io"
	"log"
	"os"
	"strconv"

	pv1 "github.com/rokane/grpc-rest-template-go/pkg/pingv1"
	"github.com/spf13/cobra"
)

// NewPingPongStreamCmd returns a new cobra command which is used to handle the
// 'pingpongstream' command.
func NewPingPongStreamCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pingpongstream",
		Short: "PingPongStream invokes the PingPongStream method on the PingServer",
		Args:  validSingleIntArg,
		Run: func(cmd *cobra.Command, args []string) {
			count, err := strconv.Atoi(args[0])
			if err != nil {
				os.Exit(1)
			}
			invokePingPongStream(count)
		},
	}
	return cmd
}

// invokePingPongStream invokes the PingPongStream method on the PingServer and
// logs each response received on the stream.
func invokePingPongStream(count int) {
	conn := getClientConnection()
	defer conn.Close()
	client := pv1.NewPingAPIClient(conn)

	stream, err := client.PingPongStream(context.Background())
	if err != nil {
		log.Println("unable to establish stream for communication")
		os.Exit(1)
	}

	// Handles receiving of responses
	wait := make(chan struct{})
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(wait)
				return
			}
			if err != nil {
				log.Println("Received error from stream:", err.Error())
			}
			log.Println("Received PingPongStream Response:",
				resp.GetMessage(), resp.GetId())
		}
	}()

	// Send messages along the stream
	for i := 1; i <= count; i++ {
		req := pv1.PingPongRequest{
			Id:      int32(i),
			Message: "Ping ...",
		}
		log.Println("Sending PingPongStream Request:",
			req.GetMessage(), req.GetId())
		if err := stream.Send(&req); err != nil {
			log.Println("Received error from stream:", err.Error())
		}
	}

	// Close connection to the stream
	if err := stream.CloseSend(); err != nil {
		log.Println("Received error closing stream:", err.Error())
	}

	// Wait until all responses have been received
	<-wait
}
