package main

import (
	"fmt"

	"calhoun.io/go-gettable/demo"
)

func main() {
	a := demo.Author{
		Name: "jon calhoun",
	}
	fmt.Println(a)
}
