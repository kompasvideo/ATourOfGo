package main

import "fmt"

type Vertex4 struct {
	lat, long float64
}

var m map[string]Vertex4

func main() {
	m = make(map[string]Vertex4)
	m["Bell Labs"] = Vertex4{
		49.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
