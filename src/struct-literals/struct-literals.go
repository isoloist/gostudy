package main

import "fmt"

type Vertex struct {
	x, y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{x: 1}
	v3 = Vertex{}
	p = &Vertex{4, 5}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}