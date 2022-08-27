package main

import (
	"fmt"
	"strings"
)

func baseName(s string) string {

	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	fmt.Println(baseName("/a/b/c/defght.go"))
}
