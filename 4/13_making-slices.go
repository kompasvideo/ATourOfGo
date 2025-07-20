package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice2(a)

	b := make([]int, 0, 5)
	printSlice2(b)

	c := b[:2]
	printSlice2(c)

	d := c[2:5]
	printSlice2(d)
}
func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
