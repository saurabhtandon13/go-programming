package main

import "fmt"


func main() {

	// this is one way to declare the variables

	// variable declaration of type int

	var age int = 35
	var newage = 4
    intdeclare := 34    // this way also we can declate the int value
	
	// variable type declaration of type string

	var name string
	name = "Saurabh"

	var newname = "Dviti"

	stringdeclare := "Shruti"   // new way to declare the string
	

	// unknown variable types example

	var z int = 00     // value of the z is not known and that why it is no defined
	x := 0.0


	// printing the int variable type
	
	fmt.Println("Age is=", age)
	fmt.Println("NewAge is=", newage)
	fmt.Println("Name age declae is ", intdeclare)

	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")

	// printing the string

	fmt.Println("Name is= ", name)
	fmt.Println("Name New is= ", newname)
	fmt.Println("Name string declare is ", stringdeclare)	


	// printig the value of unknown which is not used 

	_ = z
	_ = x
}
