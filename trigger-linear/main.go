package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{TaskQueue: "MakeCookiesQueue"}, "MakeCookies")

	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
