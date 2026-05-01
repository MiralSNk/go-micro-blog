package main

import (
	"fmt"
	"os"

	"github.com/MiralSNk/go-micro-blog/internal/gateway/config"
	"github.com/MiralSNk/go-micro-blog/internal/gateway/router"
	"github.com/MiralSNk/go-micro-blog/internal/gateway/server"
	"github.com/MiralSNk/go-micro-blog/internal/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log := logger.New(cfg.Name, cfg.LogLevel)
	r := router.New(log)

	srv := server.New(r, cfg, log)
	srv.Start()
}
