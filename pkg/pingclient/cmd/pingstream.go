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

// NewPingStreamCmd returns a new cobra command which is used to handle the
// 'pingstream' command.
func NewPingStreamCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pingstream",
		Short: "PingStream invokes the PingStream method on the PingServer",
		Args:  validSingleIntArg,
		Run: func(cmd *cobra.Command, args []string) {
			count, err := strconv.Atoi(args[0])
			if err != nil {
				os.Exit(1)
			}
			invokePingStreamMethod(count)
		},
	}
	return cmd
}

// invokePingStreamMethod invokes the PingStream method on the PingServer and
// logs the resulting response once closing the stream.
func invokePingStreamMethod(count int) {
	conn := getClientConnection()
	defer conn.Close()
	client := pv1.NewPingAPIClient(conn)

	stream, err := client.PingStream(context.Background())
	if err != nil {
		log.Println("unable to establish stream for communication")
		os.Exit(1)
	}

	// Iterate and send a request on each iteration
	for i := 1; i <= count; i++ {
		req := pv1.PingStreamRequest{
			Id:      int32(i),
			Message: "Ping ...",
		}
		log.Println("Sending PingStream Request:", req.GetMessage(), req.GetId())
		if err := stream.Send(&req); err != nil {
			if err == io.EOF {
				break
			}
			log.Println("Received error from stream:", err.Error())
		}
	}

	// Close connection stream and log the response from server
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Println("Received error from stream:", err.Error())
	}
	log.Println("Received PingStream Response:", resp.GetCount())
}
