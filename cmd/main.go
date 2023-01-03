package main

import (
	"fmt"
	"ktd/config"
	"ktd/internal/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("starting...")

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
