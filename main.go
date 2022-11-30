package main

import (
	"fmt"
	"strings"

	"github.com/soladen2010/golib/slices"
)

func main() {
	// fmt.Println(strs.SplitFirst("sd ff d", " "))
	// fmt.Println(strs.SplitLast("sd fe d", "fe"))

	ss := []string{"aaa", "bbbaa", "ccc"}
	fmt.Println(slices.FindAll(ss, func(s string) bool {
		if strings.Contains(s, "a") {
			return true
		} else {
			return false
		}
	}))

	fmt.Println(slices.FindFirst(ss, func(s string) bool {
		if strings.Contains(s, "d") {
			return true
		} else {
			return false
		}
	}))

	fmt.Println(slices.FindLast(ss, func(s string) bool {
		if strings.Contains(s, "a") {
			return true
		} else {
			return false
		}
	}))
}
