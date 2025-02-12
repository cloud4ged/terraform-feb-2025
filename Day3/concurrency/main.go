package main

import (
	"fmt"
	"time"
)

func printLoop( count int ) {
   for i:=0; i<count; i++ {
       fmt.Println(i)
       time.Sleep(time.Millisecond * 5)
   }
}

func printSomething() {
   for i := 0; i<10; i++ {
      fmt.Println ("something")
       time.Sleep(time.Millisecond * 5)
   }
}

func main() {
  go printLoop( 10 )
  go printSomething()
  var tmp string
  fmt.Scanln(&tmp)
}


