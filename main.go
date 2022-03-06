package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/duysmile/go-webrtc-server/localpool"
)

type Config struct {
	Port string
}

func NewConfig() *Config {
	port := flag.String("port", ":8000", "http port")
	return &Config{
		Port: *port,
	}
}

func main() {
	mainConfig := NewConfig()

	poolWs := localpool.NewPool()

	go poolWs.Start()

	http.HandleFunc("/ping", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		localpool.ServeWs(poolWs, w, r)
	})

	err := http.ListenAndServe(mainConfig.Port, nil)
	if err != nil {
		log.Fatal("Error to start server", err)
	}
}
