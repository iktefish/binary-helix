package utils

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "FATAL: %s", err.Error())
		os.Exit(1)
	}
}
