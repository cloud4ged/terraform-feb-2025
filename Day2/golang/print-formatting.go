package main

import "fmt"

func main() {

  str1 := "a :"
  str2 := "ab :"
  str3 := "abc :"
  str4 := "abcd :"

  fmt.Printf("%10v  %5v\n", str1, 10 )
  fmt.Printf("%10v  %5v\n", str2, 203 )
  fmt.Printf("%10v  %5v\n", str3, 1999 )
  fmt.Printf("%10v  %5v\n", str4, 19999 )


  //%15.5f means total width of the float is 15 
  //out of which 5 decimal places precision must be considered
  fmt.Printf("%10v  %015.5f\n", str1, 10.12 )
  fmt.Printf("%10v  %015.5f\n", str2, 203.123 )
  fmt.Printf("%10v  %015.5f\n", str3, 1999.123 ) //total width is 15 and 5 decimal precisions
  fmt.Printf("%10v  %.2f\n", str4, 19999.1234 ) //default width with 2 decimal precisions

  //print numbers in different number system format
  fmt.Printf("Decimal 5 Printed in binary format: %05b\n", 5 )
  fmt.Printf("Decimal 10 Printed in decimal format: %05d\n", 10 )
  fmt.Printf("Decimal 10 Printed in octal format: %05o\n", 10 )
  fmt.Printf("Decimal 1019 Printed in hexadecimal format: %05x\n", 1019 )


}
