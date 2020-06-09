package main

import (
	"fmt"
)

var a, b int

func main() {
	fmt.Println(Add(3, 5))
}

//Add は足す関数
func Add(a int, b int) int {
	return a + b
}
