package main

import (
	"log"
	"net/rpc"
)

/*
NOTE:
If you are catching the recieving data with a custom type (structure),
then you must define the type here too. Its best to share the type to
both server and client.
*/

func main() {
    client, err := rpc.DialHTTP("tcp", "172.17.0.2:4040")
    if err != nil {
        log.Fatal("Connection error: ", err)
    }

    s := "String from the client!"
    var response string

    client.Call("API.GetByName", s, &response)
}
