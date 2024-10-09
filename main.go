package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/tirlochanarora16/go_microservice/application"
)

func main() {
	app := application.New(application.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)

	defer cancel()

	err := app.Start(ctx)

	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
