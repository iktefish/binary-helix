package server

import (
	"fmt"
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

	fmt.Println("Initial query from DB here!")

	chunk_1 := "Query this chunk_1"
	chunk_2 := "Query this chunk_2"
	chunk_3 := "Query this chunk_3"
	fmt.Println(chunk_1, chunk_2, chunk_3)

	fmt.Println("Query to add something to DB here!")
	fmt.Println("Delete something from DB here!")
	fmt.Println("Mod doc from DB here!")
}
