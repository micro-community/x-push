package main

import (
	"net/http"

	"github.com/micro-community/x-push/config"
	"github.com/micro-community/x-push/server"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	// New Service
	service := web.NewService()

	// Initialize service
	service.Init(web.Name(config.ServiceName),
		web.Version(config.Version),
	)
	// static files
	service.Handle("/", http.FileServer(http.Dir("html")))

	// Handle websocket connection
	service.HandleFunc("/ws", server.HandleConn)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
