package main

import "fmt"

func main() {
   var name = "James Gosling"
   fmt.Println(name)

   var firstNumber, secondNumber int = 10, 20
   fmt.Println( "Value of first number is ", firstNumber )
   fmt.Println( "Value of second number is ", secondNumber )

   var thirdNumber = 30
   fmt.Println ( "Value of third number is ", thirdNumber )

   var isSuccess = true
   fmt.Println(isSuccess)

   startupName:= "TekTutor"
   fmt.Println(startupName)

   //Let's print the data types of the variables we used above
   fmt.Printf("%T\n", name )
   fmt.Printf("%T\n", firstNumber)
   fmt.Printf("%T\n", isSuccess)
   fmt.Printf("%T\n", startupName)
}
