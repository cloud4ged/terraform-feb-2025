package main

import (
   "fmt"
   "os"
   "io/ioutil"
   "log"
)

func createFile( filename string ) {
	myfile, err := os.Create( filename )

	if err != nil {
	   //log.Fatal(err)
	   panic(err)
	}

	str := "Some content"

	bytesWritten, err := myfile.WriteString( str + "\n")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Wrote %d bytes into the file\n", bytesWritten)
	myfile.Sync()
}

func readFile( filename string ) {
	content, err := ioutil.ReadFile(filename)
        if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File content: %s", content
}

//This is the entry point function that will be automatically by go
func main() {
	file := "./myfile.txt"

	createFile( file )
	readFile(file) 

}
