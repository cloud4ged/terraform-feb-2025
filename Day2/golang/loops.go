package main
import "fmt"

func main() {

	count := 5

	fmt.Println("----------------------------------------")
	for count > 0 {
		fmt.Println (count)
		count--
	}

	fmt.Println(count)

	fmt.Println("----------------------------------------")

	count = 0

	for count = 1; count < 10; count++ {
		fmt.Printf("%d\t",count)
	}
	fmt.Println()
	fmt.Println("----------------------------------------")
}
