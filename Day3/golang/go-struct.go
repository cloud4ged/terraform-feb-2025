package main

import "fmt"

type Employee struct {
  id int
  name string
  designation string
}

func printEmployee( employee Employee ) {
	fmt.Println()
	fmt.Println ( "Employee ID          : ", employee.id )
	fmt.Println ( "Name of the employee : ", employee.name )
	fmt.Println ( "Designation          : ", employee.designation)
	fmt.Println()
}

//Return Employee struct instance 
func createEmployee( id int, name string, designation string  ) Employee {
	var employee Employee

	employee.id = id
	employee.name = name
	employee.designation = designation 

	return employee
}

func main() {
	var sriram Employee
	var nitesh Employee

	sriram = createEmployee ( 1, "Sriram", "Software Enginer I" )
	nitesh = createEmployee ( 2, "Nitesh", "Software Enginer II" )

        printEmployee ( sriram )
        printEmployee ( nitesh )
/*
	sriram.name = "Sriram"
	sriram.id = 1
	sriram.designation = "Software Engineer 1"

	nitesh.name = "Nitesh"
	nitesh.id = 2
	nitesh.designation = "Sr.Software Engineer"
	fmt.Println()
	fmt.Println ( "Name of the employee : ", sriram.name )
	fmt.Println ( "Designation          : ", sriram.designation)
	fmt.Println()

	fmt.Println ( "Name of the employee : ", nitesh.name )
	fmt.Println ( "Designation          : ", nitesh.designation)
	fmt.Println()
*/
}
