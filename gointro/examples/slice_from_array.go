package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:4]
	fmt.Printf("%#v\n", s)
}
