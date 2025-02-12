package main
import "fmt"

func compareString( str1, str2 string ) bool {
	var result bool

	if str1 == str2 {
		result = true
	} else {
		result = false
	}

	return result
}

func main() {
	str1 := "go"
	str2 := "go"

	fmt.Printf("%s and %s matching? : %v\n", str1, str2, compareString(str1,str2) )

	str1 = "Go"
	str2 = "go"
	fmt.Printf("%s and %s matching? : %v\n", str1, str2, compareString(str1,str2) )
}
