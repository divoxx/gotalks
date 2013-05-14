package main

func main() {
	a3 := [3]int{1, 2, 3}
	a5 := [5]int{}

	// [3]int and [5]int are incompatible types
	a3 = a5
}
