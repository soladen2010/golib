package slices

import (
	"fmt"
	"testing"
)

type User struct {
	id   int
	name string
}

func TestForEach(t *testing.T) {
	arr := []User{{id: 1, name: "tom"}, {id: 2, name: "jack"}}
	ForEach(arr, func(e *User) {
		fmt.Println(e.id, e.name)
	})
}

func TestForEachParal(t *testing.T) {
	arr := []User{{id: 1, name: "tom"}, {id: 2, name: "jack"}}
	ForEachParal(arr, 2, func(e *User) {
		fmt.Println(e.id, e.name)
	})
}

func TestMap(t *testing.T) {
	arr := []User{{id: 1, name: "tom"}, {id: 2, name: "jack"}}
	intArr := Map(arr, func(e User) int {
		return e.id
	})
	fmt.Println(intArr)
}

func TestMapParal(t *testing.T) {
	arr := []User{{id: 1, name: "tom"}, {id: 2, name: "jack"}, {id: 3, name: "marry"}}
	intArr := MapParal(arr, 2, func(e User) int {
		return e.id
	})
	fmt.Println(intArr)
}

func TestReduce(t *testing.T) {
	arr := []User{{id: 1, name: "tom"}, {id: 2, name: "jack"}, {id: 3, name: "marry"}}
	sum := Reduce(arr, 0, func(s int, e User) int {
		return s + e.id
	})
	fmt.Println(sum)
}

func TestFilter(t *testing.T) {
	arr := []User{{id: 1, name: "tom"}, {id: 2, name: "jack"}, {id: 3, name: "marry"}}
	filteredArr := Filter(arr, func(e User) bool {
		return e.id > 1
	})
	fmt.Println(filteredArr)
}

func TestReverse(t *testing.T) {
	arr := []User{{id: 1, name: "tom"}, {id: 2, name: "jack"}, {id: 3, name: "marry"}}
	reversedArr := Reverse(arr)
	fmt.Println(reversedArr)
}

func TestGroupCount(t *testing.T) {
	arr := []User{{id: 1, name: "tom"}, {id: 2, name: "jack"}, {id: 3, name: "marry"}, {id: 2, name: "jack"}}
	countMap := GroupCount(arr)
	fmt.Println(countMap)
}

func TestFindLast(t *testing.T) {
	arr := []User{{id: 1, name: "tom"}, {id: 2, name: "jack"}, {id: 3, name: "marry"}, {id: 2, name: "jackNew"}}
	last := FindLast(arr, func(e User) bool {
		return e.id == 2
	})
	fmt.Println(last)
}

func TestFindAll(t *testing.T) {
	arr := []User{{id: 1, name: "tom"}, {id: 2, name: "jack"}, {id: 3, name: "marry"}, {id: 2, name: "jackNew"}}
	all := FindAll(arr, func(e User) bool {
		return e.id == 2
	})
	fmt.Println(all)
}

// create a test function for FindFirst
func TestFindFirst(t *testing.T) {
	arr := []User{{id: 1, name: "tom"}, {id: 2, name: "jack"}, {id: 3, name: "marry"}, {id: 2, name: "jackNew"}}
	first := FindFirst(arr, func(e User) bool {
		return e.id == 2
	})
	fmt.Println(first)
}
