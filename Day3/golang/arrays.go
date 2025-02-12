package main

import "fmt"

func sayHello( yourName string ) string {
   return "Hello, " + yourName + " !"
}

func main() {

	fmt.Println( sayHello("Golang") )


	//We have declared an array of integers of size 5
	//we can maximum store 5 integers in it
	//Go lang array are fixed in size
	var arr[5] int

	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	fmt.Println ( arr )

	count := len(arr)

	// "Hello" + "World" - valid
	// 10 + "Hello" - is not allowed
	fmt.Println ( "Length of array : ", count )
	fmt.Println ( "Hello" + "World" )

	fmt.Println("Array elements are ...")
	for i:=0; i<count; i++ {
		fmt.Printf("%d\t", arr[i] )
	}
	fmt.Println()

	arr[2] = 25
	fmt.Println ("Modified array")
	fmt.Println ( arr )
}
