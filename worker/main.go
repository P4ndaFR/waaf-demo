package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	app "github.com/nd4pa/waaf-demo"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "MakeCookiesQueue", worker.Options{})

	w.RegisterWorkflow(app.MakeCookies)
	w.RegisterWorkflow(app.MakeCookiesParallelDough)

	w.RegisterActivity(app.MakeDough)
	w.RegisterActivity(app.ShapeCookies)
	w.RegisterActivity(app.BakeCookies)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
