package main

import "fmt"

func add( x, y int ) int {
	return x+y
}

func subtract( x, y int ) int {
	return x-y
}

func multiply( x, y int ) int {
	return x*y
}

func divide( x, y int ) int {
	return x/y
}

func main() {
	a := 10
	b := 20

	fmt.Printf( "The sum of %d + %d is %d\n", a, b, add(a,b) )
	fmt.Printf( "The diff between %d and %d is %d\n", a, b, subtract(a,b) )
	fmt.Printf( "The product of %d and %d is %d\n", a, b, multiply(a,b) )
	fmt.Printf( "%d / %d is %d\n", a, b, divide(a,b) )
}
