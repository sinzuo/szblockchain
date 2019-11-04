package main

import (
	"strconv"
)

func HexToString(value int64) []byte {
	return strconv.AppendInt(nil, value, 10)
}
