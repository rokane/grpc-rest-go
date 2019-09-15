package main

import (
	"log"

	"github.com/rokane/grpc-rest-template-go/pkg/pingclient/cmd"
)

func main() {
	pcCmd := cmd.NewInitialisedPingClientCmd()
	if err := pcCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
