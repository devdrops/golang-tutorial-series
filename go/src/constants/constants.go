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
	// untyped constant has a default type associated with them. So, every
	// constant that doesn't explicitly declares it's type is a untyped contanst.
	// BUT, the constant can assume the type based on the value only if the code
	// requires it.
	// Weird, eh? As the example below:
	const nomezito = "Jose de Los Vegetales"
	fmt.Printf("El tipo %T e valor %v\n", nomezito, nomezito)
	// When we write `const nomezito`, this constant is untyped; however, as we're
	// assigning `"Jose de Los Vegetales"`, which is another constant but it's a
	// string, our constant `nomezito` interprets the very same type to it's own
	// type. Just because the value assignment, as a code operation, demands it.
	// Honestly this sounds like a lie to me, it's like the assignment itself is
	// the only responsible for the constant type, period.
	// For me, this is weird AF.
	// To make things clear, we can also create typed constants:
	const typedSomething string = "Something"
	fmt.Printf("type %T and value %v\n", typedSomething, typedSomething)
	// As a proof that typed constants can only assign values from the same type,
	// the line below raises a `cannot convert 100 (type untyped number) to type
	// string` error.
	//const eitaPreula string = 100
	// And something kinda strange happens if we try the below:
	var firstName = "Foo"
	type myString string
	var anotherName myString = "Bar"
	// Until here, everything is 200 OK. But the line below has a surprise:
	anotherName = firstName
	// It raises the error `cannot use firstName (type string) as type myString in
	// assignment`, which is telling us that type string is != of type myString,
	// even if they both accepts strings! The type myString above is an alias for
	// string.
}

