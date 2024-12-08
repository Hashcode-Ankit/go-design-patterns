package main

import (
	"fmt"
	"main/builder"
)

func main() {
	newCarBuilder := builder.NewBuilder(1, "Sports Car Wheels")
	fmt.Println(newCarBuilder.Stop())
}
