package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

/* NOTE:
Criteria for Remote Procedure Calls: (refer to the manual for details)
- Functions need to be methods
- All functions have to be exported
- Functions need to have 2 args, both of those have to be exported types (built-in types are exported)
- Second arg of the functions must be a pointer
- Return type must be `error` type
*/

/* Using the API type to elevate all of out functions into methods. */
type API int

func (a *API) GetByName(argFromCaller string, resultFromFunction *string) error {
	/* Some DB query */

	/* Do what you want here, get the output you want returned, store it in a var,
	   and pass the output as a pointer: */
	// var IWantThisReturned string
	// *resultFromFunction = IWantThisReturned

	var IWantThisReturned string = "Pee Kaa Boo"
	*resultFromFunction = IWantThisReturned

	return nil
}

func (a *API) CreateDoc(argFromCaller string, resultFromFunction *string) error {
	/* Some DB create doc */

	/* Do what you want here, get the output you want returned, store it in a var,
	   and pass the output as a pointer: */
	// database = append(database, item)
	// *resultFromFunction = item

	return nil
}

func (a *API) InsertDoc(argFromCaller string, resultFromFunction *string) error {
	/* Some DB insert command */

	/* Do what you want here, get the output you want returned, store it in a var,
	   and pass the output as a pointer: */
	// database = append(database, item)
	// *resultFromFunction = item

	return nil
}

func (a *API) ModDoc(argFromCaller string, resultFromFunction *string) error {
	/* Some DB doc mod command */

	return nil
}

func(a *API) DeleteDoc(argFromCaller string, resultFromFunction *string) error {
	/* Some DB doc del command */

	return nil
}

func main() {
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
