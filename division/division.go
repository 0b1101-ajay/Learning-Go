package main

import "fmt"

func main() {
	
	var i int = 97
	var j int = 5
	fmt.Println(i / j)

	var c float64 = 97
	var d float64 = 5
	fmt.Println(c / d)

	var e int = 97
	fmt.Println(e /5)

	var f float64 = 97
	fmt.Println(f / 5)


	// Throws error:invalid operation: a / b (mismatched types float64 and int)
	// var a float64 = 97
	// var b int = 5
	// fmt.Println(a / b)
}
