package main

import (
	"fmt"
	"practice/lib"
)

func main() {
	i1 := lib.GetInstance()
	i2 := lib.GetInstance()

	fmt.Println(i1.AddOne(), i2.AddOne())
}
