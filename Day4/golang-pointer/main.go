package main

import "fmt"

func someFunction(x *int) {
	*x = 200
	fmt.Println("Value of x inside someFunction is ", *x )
}

func main() {
	i := 100
	fmt.Println("Value of i in main function before calling someFunction is ", i)
	someFunction( &i )
	fmt.Println("Value of i in main function after calling someFunction is ", i)
}
