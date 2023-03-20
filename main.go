package main

import (
	"fmt"

	"github.com/soladen2010/golib/slices"
)

func main() {
	// fmt.Println(strs.SplitFirst("sd ff d", " "))
	// fmt.Println(strs.SplitLast("sd fe d", "fe"))

	// ss := []string{"aaa", "bbbaa", "ccc"}
	// fmt.Println(slices.FindAll(ss, func(s string) bool {
	// 	if strings.Contains(s, "a") {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// }))

	// fmt.Println(slices.FindFirst(ss, func(s string) bool {
	// 	if strings.Contains(s, "d") {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// }))

	// fmt.Println(slices.FindLast(ss, func(s string) bool {
	// 	if strings.Contains(s, "a") {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// }))

	// slices.ForEach(ss, func(e string) {
	// 	fmt.Println(e)
	// })

	// slices.ForEachPt(ss, func(e *string) {
	// 	*e += "ZZZ"
	// 	fmt.Println(*e)
	// })

	numbers := []int{}
	for i := 0; i < 100; i++ {
		numbers = append(numbers, i)
	}
	slices.ForEachParal(numbers, 5, func(num int) {
		fmt.Println(num)
	})
	// slices.ForEachPtParal(numbers, 5, func(num *int) {
	// 	fmt.Println(*num)
	// })
	// slices.MapParal(numbers, 5, func(n int) int {
	// 	fmt.Println(100 * n)
	// 	return 100 * n
	// })
}
