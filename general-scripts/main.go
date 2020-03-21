package main

import "fmt"

var x int = 100
const y int = 10
var z float64 = 25

func main(){
	fmt.Println("\n")
	fmt.Println("Value of x defined is =", x)
	fmt.Println("Value of the y is= ", y)
	fmt.Println("The float value of z is=", z)

	fmt.Println("\n")

	var x int = x/2
	const y int = y/2
	var z float64= z/2

	fmt.Println("The Value of x has been changed to=", x)
	fmt.Println("New value of the y is= ", y)
	fmt.Println("value of the z float value is reduced to= ",z)
	fmt.Println("\n")
}