package main

import "fmt"

func main() {

	intArray := [5] int { 10, 20, 30, 40, 50 } 

	fmt.Println ( "Array elements are ..." )
	fmt.Println ( intArray )

	//slice uses array internally
	//the size of slice can change at runtime
	//2 - is the starting index 
	var slice []int = intArray[2:4]
	fmt.Println ( "Slice elements referenced from array are ..." )
	fmt.Println( slice )

	//it creates a new slice with appended 100 and that's how slice grows in size at runtime
	fmt.Println ( "Size of intArray before slice append is ", len(intArray) )
	slice = append( slice, 100, 200, 300 )
	fmt.Println( slice )

	fmt.Println ("Int array after modifying slice")
	fmt.Println ( intArray )
	fmt.Println ( "Size of intArray after appending is ", len(intArray) )

	//intArray[5] = 100 - This will report index out of bounds array as array size is fixed

	//internally slice creates an array of integer of size 3
	//anotherSlice is pointing the entire range of elements of the array
	anotherSlice := [] int { 1, 2, 3}
	fmt.Println ( "Another slice elements are ..." )
	fmt.Println ( anotherSlice )

	//slice will create a new array of size 6 and we need ensure the anotherSlice points 
	//to the new array
	anotherSlice = append( anotherSlice, 4, 5, 6 )
	fmt.Println ( "Another slice elements after appending are ..." )
	fmt.Println ( anotherSlice )
}
