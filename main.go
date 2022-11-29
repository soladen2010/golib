package main

import (
	"fmt"

	"github.com/soladen2010/golib/strs"
)

func main() {
	fmt.Println(strs.SplitFirst("sd ff d", " "))
	fmt.Println(strs.SplitLast("sd ff d", " "))
}
