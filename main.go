package main

import (
	"fmt"
	"golang-workshop/exampleDocumentation/dog"
)

type lab struct {
	name string
	age  int
}

func main() {
	lab1 := lab{
		name: "Maissie",
		age:  dog.Years(2),
	}
	fmt.Println(lab1)
}
