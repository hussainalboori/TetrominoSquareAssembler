package src

import "strings"

func CleanData(s string) string{
	return strings.Trim(s, "/n/t")
}