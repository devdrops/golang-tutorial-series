package main

import (
	"fmt"
	"math"
)

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

	// We can also declare a pack of variables of different types in a single
	// statement. The rules for type inference work the same way :)
	var (
		age10 = 10
		age11 = "11"
		age12 int
	)
	fmt.Println("My ages 10, 11 and 12, which are different, are both", age10, ",", age11, "and", age12)

	// Variables can also be declared using a shorthand syntax, just like that ;)
	// Shorthand declaration requires an initial value for each variable, which,
	// as expected, uses type inference based on the assigned value.
	foo, bar := "Wut", 100
	fmt.Println("My name is", foo, "and my age is", bar)

	// We can only use the shorthand syntax when we have at least one newly
	// created variable. As below, we can reassign values because we are
	// also declaring new variables altogether.
	foo2, bar2 := 10, 20
	fmt.Println("Foo2 is", foo2, "while bar2 is", bar2)
	foo2, baz := 30, 40
	fmt.Println("Foo2 is", foo2, "while baz is", baz)
	foo2, baz = 50, 60
	fmt.Println("Foo2 is", foo2, "while baz is", baz)

	// If we try to reassign only variables that have already been created, well,
	// all hell's breaks loose.
	// This works OK, because we're creating two new variables :)
	foo3, bar3 := 1, 2
	fmt.Println("Foo3 is", foo3, "while bar3 is", bar3)
	// Remove the comments of the line below and OUCH! `no new variables on left
	// side of :=`
	//foo3, bar3 := 3, 4
	// Now, if we add a new variable on the left size, then Go will make it
	// happen.
	foo3, bar3, baz3 := 5, 6, 7
	fmt.Println("Foo3 is", foo3, "while bar3 is", bar3, "and Baz3 is", baz3)

	// We can also use the shorthand declaration at runtime, assigning values only
	// when the code is executed :)
	foo4, bar4 := 49.9, 49.8
	baz4 := math.Min(foo4, bar4)
	fmt.Println("Baz4 is", baz4)

	// As Go has a strong type relation, we cannot reassign a value which isn't of
	// the same type of it's original value.
	baz5 := 100
	//baz5 = "wuuuuuuuuuut"
	fmt.Println("Baz5 is", baz5)

	// Same goes for variables created without using shorthand syntax.
	var baz6 = 300
	//baz6 = "waaaaaaaaaat"
	fmt.Println("Baz6 is", baz6)
}
