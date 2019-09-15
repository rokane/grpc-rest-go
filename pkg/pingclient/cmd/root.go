package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	port string
)

// NewInitialisedPingClientCmd return an initialised the pingclient command
// ensuring flags are passed and connection to PingServer is established.
func NewInitialisedPingClientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pingclient",
		Short: "PingClient communicates with the PingServer",
		Long:  `Invoke methods on the PingServer over gRPC.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// Parse input flags
	cmd.PersistentFlags().StringVarP(&port, "port", "p",
		"8080", "Connection port for PingServer")

	// Add subcommands
	cmd.AddCommand(NewPingCmd())
	cmd.AddCommand(NewPingStreamCmd())
	cmd.AddCommand(NewPongStreamCmd())
	cmd.AddCommand(NewPingPongStreamCmd())

	return cmd
}

// getClientConnection attempts to connect to the grpc server and return a
// connection which will be used for the client to connect to.
func getClientConnection() *grpc.ClientConn {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", port), grpc.WithInsecure())
	if err != nil {
		os.Exit(1)
	}
	return conn
}

// validSingleIntArg returns an error if the args does not contain a single, int
// value. If it does no error is returned and it is considered valid.
func validSingleIntArg(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("requires a single int argument")
	}
	if _, err := strconv.Atoi(args[0]); err != nil {
		return fmt.Errorf("received non int arg: %v", args[0])
	}
	return nil
}
