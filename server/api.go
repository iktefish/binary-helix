package server

import (
	"fmt"
	"strings"

	"github.com/iktefish/binary-helix/analyser"
	"github.com/iktefish/binary-helix/types"
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

func (a *API) ImAlive(argFromCaller string, resultFromFunction *string) error {
	var to_be_returned string = "Alive"
	*resultFromFunction = to_be_returned

	return nil
}

func (a *API) CallBoyerMoore(argFromCaller []string, resultFromFunction *string) error {
	// func (a *API) CallBoyerMoore(argFromCaller []string, resultFromFunction *[]int) error {
	pBM := types.ConstructBM(argFromCaller[3])
	out := analyser.BoyerMoore(argFromCaller[3], pBM, argFromCaller[1])
	send := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(out)), " : "), "[]")
	*resultFromFunction = send
	// *resultFromFunction = out

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
	*resultFromFunction = ""

	return nil
}

func (a *API) CallLongestCommonPrefix(argFromCaller []string, resultFromFunction *string) error {
	*resultFromFunction = ""

	return nil
}

func (a *API) CallIdQual(argFromCaller []string, resultFromFunction *string) error {
	*resultFromFunction = ""

	return nil
}
