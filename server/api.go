package server

import "github.com/iktefish/binary-helix/analyser"

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

func (a *API) GetByName(argFromCaller []string, resultFromFunction *string) error {
	/* Some DB query */

	/* Do what you want here, get the output you want returned, store it in a var,
	   and pass the output as a pointer: */
	// var IWantThisReturned string
	// *resultFromFunction = IWantThisReturned

	// var IWantThisReturned string = "Pee Kaa Boo"
	// *resultFromFunction = IWantThisReturned
	*resultFromFunction = argFromCaller[1]

	return nil
}

func (a *API) ImAlive(argFromCaller string, resultFromFunction *string) error {
	var to_be_returned string = "Alive"
	*resultFromFunction = to_be_returned

	return nil
}

func (a *API) CallComplement(argFromCaller string, resultFromFunction *string) error {
	// var to_be_returned string = analyser.Complement(argFromCaller[0])
	// var to_be_returned string = "Alive"
	// *resultFromFunction = to_be_returned
	*resultFromFunction = analyser.Complement(argFromCaller)

	return nil
}
