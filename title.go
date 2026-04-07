package main

import "strings"

func titlewords(s, r string) string {
	switch r {
	case "cap":
		return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
	case "title":
		return strings.ToTitle(s)
	case "low":
		return strings.ToLower(s)
	case "head":
		return strings.Title(s)
	default:
		return s
	}
}
