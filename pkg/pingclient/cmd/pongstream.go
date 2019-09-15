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

// NewPongStreamCmd returns a new cobra command which is used to handle the
// 'pongstream' command.
func NewPongStreamCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pongstream",
		Short: "PongStream invokes the PongStream method on the PingServer",
		Args:  validSingleIntArg,
		Run: func(cmd *cobra.Command, args []string) {
			count, err := strconv.Atoi(args[0])
			if err != nil {
				os.Exit(1)
			}
			invokePongStreamMethod(count)
		},
	}
	return cmd
}

// invokePongStreamMethod invokes the PongStream method on the PingServer and
// log each of the response received on the stream.
func invokePongStreamMethod(count int) {
	conn := getClientConnection()
	defer conn.Close()
	client := pv1.NewPingAPIClient(conn)

	req := pv1.PongStreamRequest{
		Count:   int32(count),
		Message: "Ping ...",
	}
	log.Println("Sending PongStream Request:", req.GetCount())
	stream, err := client.PongStream(context.Background(), &req)
	if err != nil {
		log.Println("unable to establish stream for communication")
		os.Exit(1)
	}

	// Iterate until the stream closes, logging each response
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Received error from stream:", err.Error())
		}
		log.Println("Received PongStream Response:",
			resp.GetMessage(), resp.GetId())
	}
}
