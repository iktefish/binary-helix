package server

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func Server() {
	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("Error ocurred during API registration", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listening error", err)
	}

	log.Printf("Serving RPC on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}
