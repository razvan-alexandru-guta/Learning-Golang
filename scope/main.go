package main

import (
	"myapp/packageone"
)

//package level variable (small first letter, not exported)
var myVar = 2

func main() {
	blockVar := "just a main func var"
	packageone.PrintMe(myVar, blockVar)
}
