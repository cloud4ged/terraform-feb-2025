package main

import (
	"fmt"
	"time"
)

func printLoop( count int ) {
   for i=0; i<count; i++ {
       fmt.Println(i)
   }
}

func main() {
  printLoop( 10 )
}


