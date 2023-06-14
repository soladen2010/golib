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
