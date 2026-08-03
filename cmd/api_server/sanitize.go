package main

import "strings"

func normalize(s string) string {
	lowered := strings.ToLower(s)
	final := strings.ReplaceAll(lowered, " ", "_")
	return final

}
