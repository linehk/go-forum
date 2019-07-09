package main

import (
	"log"
	"net/http"
	"time"

	"github.com/linehk/go-forum/config"
	"github.com/linehk/go-forum/controller"
)

var sc = config.Cfg.Server

func main() {
	server := &http.Server{
		Addr:           sc.Addr,
		Handler:        controller.Setup(),
		ReadTimeout:    time.Duration(sc.ReadTimeout * int(time.Second)),
		WriteTimeout:   time.Duration(sc.WriteTimeout * int(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}
