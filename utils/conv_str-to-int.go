package utils

import (
	"flag"
	"strconv"
)

func Conv_StrToInt(s string) (int, error) {
	flag.Parse()
	i, err := strconv.Atoi(s)

	return i, err
}
