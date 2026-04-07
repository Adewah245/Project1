package main

import (
	"errors"
	"strconv"
)

func convert(s string, base string) (int64, error) {
	switch base {
	case "1":
		return strconv.ParseInt(s, 16, 64)
	case "2":
		return strconv.ParseInt(s, 2, 64)
	case "3":
		return strconv.ParseInt(s, 8, 64)
	default:
		return 0, errors.New("error parsing the base")
	}
}
