package main

import (
	"fmt"
	"math"
)

func main() {
	// Constants
	//
	// Constants are fixed values, so, immutable.
	var foo string = "Eita preula"
	// When we declare a variable, as above, we are assigning to `foo` the
	// constant "Eita preula", even if we are not using the `const` 
	// declaration.
	const bar = "LOLWTFBBQ"
	// As constants are immutable, we cannot reassign values once a
	// constant is created. The line below will raise a
	// `cannot assign to bar` error.
	//bar = foo
	fmt.Println("Now bar is", bar, "while foo is", foo)
	// The reassignment is not possible due to how Go interprets a constant: it's
	// value can only be known at compile time. So, we also can't assign return
	// values from functions to a constant. Only constant values can be assigned
	// to constants ;)
	// The line bellow works fine:
	var eita = math.Sqrt(1)
	fmt.Println("The square of 1 is", eita)
	// While the line below will raise a `const initializer math.Sqrt(1) is not a
	// constant` error - the value 1 is a constant, but the value of math.Sqrt(1)
	// will only be known at runtime:
	//const preula = math.Sqrt(1)

	// String constants
	//
	// Anything enclosed between double quotes is a string constant in Go.
	// "LOLWTFBBQ", "Eita preula" and so on. But notice that these constants
	// doesn't have any type!
	fmt.Printf("The constant bar is of type %T\n", bar)
	fmt.Printf("The string constant is of type %T\n", "EITA")
	var name = "Sam"
	fmt.Printf("type %T and value %v\n", name, name)
	// HEY, I CAN SEE THE TYPE RIGHT THERE, WHAT'S GOING ON?
	// Ok, we can see the type, but this works a little bit different in Go. Every
	// untyped constant has a default type associated with them.
}
