package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0, 10)
	fmt.Printf("len:%d cap:%d\n", len(s), cap(s))
	s = append(s, 1)
	s = append(s, 2, 3)
	a := []int{4, 5}
	s = append(s, a...)
	fmt.Printf("%#v\n", s)
}
