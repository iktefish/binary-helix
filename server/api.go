package server

import (
	"fmt"
	// "reflect"
	"strconv"
	"strings"

	"github.com/iktefish/binary-helix/analyser"
	"github.com/iktefish/binary-helix/types"
	"github.com/iktefish/binary-helix/utils"
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

func (a *API) CallReverseComplement(argFromCaller string, resultFromFunction *string) error {
	// var to_be_returned string = analyser.Complement(argFromCaller[0])
	// var to_be_returned string = "Alive"
	// *resultFromFunction = to_be_returned
	*resultFromFunction = analyser.ReverseComplement(argFromCaller)

	return nil
}

func (a *API) CallExactMatch(argFromCaller []string, resultFromFunction *string) error {
	out := analyser.ExactMatch(argFromCaller[3], argFromCaller[1])
	send := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(out)), " "), "[]")
	*resultFromFunction = send

	return nil
}

func (a *API) CallKMer(argFromCaller []string, resultFromFunction *string) error {
	strVar := argFromCaller[3]
	intVar, err := strconv.Atoi(strVar)
	utils.HandleError(err)
	// fmt.Println(intVar, err, reflect.TypeOf(intVar))

	out := analyser.ConstructIA(argFromCaller[1], intVar)

	var new []string
	for _, v := range out.I {
		sVal := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(v.Val)), " "), "[]")
		new = append(new, v.Key, ":", sVal, "., ")
	}
	send := strings.Join(new, ".")

	// send := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(out)), " "), "[]")
	*resultFromFunction = send

	return nil
}

func (a *API) CallLongestCommonPrefix(argFromCaller []string, resultFromFunction *string) error {
	out := analyser.LongestCommonPrefix(argFromCaller[1], argFromCaller[3])
	// send := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(out)), " "), "[]")
	*resultFromFunction = out

	return nil
}

func (a *API) CallBoyerMoore(argFromCaller []string, resultFromFunction *string) error {
	pBM := types.ConstructBM(argFromCaller[3])
	out := analyser.BoyerMoore(argFromCaller[3], pBM, argFromCaller[1])
	send := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(out)), " "), "[]")
	*resultFromFunction = send

	return nil
}
