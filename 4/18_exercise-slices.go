package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	return make([][]uint8, 0)
}

func main() {
	pic.Show(Pic)
}
