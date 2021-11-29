package packageone

import "fmt"

// package level var (capital letter, exported)
var PackageVar = 3

func PrintMe(myVar int, blockVar string) {
	fmt.Println(myVar, blockVar, PackageVar)
}
