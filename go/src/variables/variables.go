package main

import "fmt"

func main() {
	// Without assigning value, Go automatically assings 0.
	var age int
	fmt.Println("My age is", age)
	age = 29
	fmt.Println("My age is", age)
	age = 54
	fmt.Println("My age is", age)

	// Starting a new one, this time assigning a value.
	var age2 int = 21
	fmt.Println("My age2 is", age2)

	// Type inference: when you don't need to explicitly declare the variable type,
	// as Go will understand it's type based on the assigned value.
	var age3 = 66
	fmt.Println("My age3 is", age3)

	// Declaring multiple variables in a single statement.
	var age4, age5 int = 100, 50
	fmt.Println("My ages 4 and 5 are", age4, "and", age5)

	// We can also declare multiple variables using type inference as well :)
	var age6, age7 = 51, 99
	fmt.Println("My ages 6 and 7 are", age6, "and", age7)

	// And, as we can guess, multiple variable declarations can be done without
	// value assignment.
	var age8, age9 int
	fmt.Println("My ages 8 and 9 are both", age8, "and", age9)
}
