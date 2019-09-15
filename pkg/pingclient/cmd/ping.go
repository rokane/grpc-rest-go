package cmd

import (
	"context"
	"log"

	pv1 "github.com/rokane/grpc-rest-template-go/pkg/pingv1"
	"github.com/spf13/cobra"
)

// NewPingCmd returns a new cobra command which is used to handle the 'ping'
// command.
func NewPingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ping",
		Short: "Ping invokes the Ping method on the PingServer",
		Run: func(cmd *cobra.Command, args []string) {
			invokePingMethod()
		},
	}
	return cmd
}

// invokePingMethod invokes the Ping method on the PingServer and will
// log the result.
func invokePingMethod() {
	conn := getClientConnection()
	defer conn.Close()
	client := pv1.NewPingAPIClient(conn)

	log.Println("Sending Ping Request:")
	resp, err := client.Ping(context.Background(), &pv1.PingRequest{})
	if err != nil {
		log.Println("error received from PingServer", err.Error())
	}
	log.Println("Received Ping Response:", resp.GetMessage())
}
