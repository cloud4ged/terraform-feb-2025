# Day 3

## Lab - Creating a custom module in golang

First let's create a directory name hello
```
cd ~
mkdir hello
cd hello
go mod init tektutor.org/hello
```
Let's create a hello.go with the below content
<pre>
package hello

import "fmt"

func SayHello( name string ) string {
	message := fmt.Sprintf("Hello, %v !", name ) 
	return message
}
</pre>  

Expected output
![image](https://github.com/user-attachments/assets/7f56365c-e5af-45bb-b719-4c9247b472ff)
![image](https://github.com/user-attachments/assets/75e3fc7d-d153-43d4-a486-19049fa740c9)

Let's create a main.go file with the below content

## Lab - Running two or more functions in parallel using concurrency feature in go
```
cd ~/terraform-feb-2025
git pull
cd Day3/concurrency
go run ./main.go
```

Expected output
![image](https://github.com/user-attachments/assets/13516d7c-2cc1-4900-9183-1bf7c782eb3b)

To exit, you may press any key.

## Info - Golang modules
<pre>
- is a way we can write reusable code
- go modules can have multiple go files
- go modules can export reusable functions to the external world by ensure the function name first letter begins with Upper case
- any internal implementation details in go module can be abstracted by naming those functions with _ or by starting the function name with a lowercase character
</pre>

## Lab - Creating a custom go module

In this lab exercise, you will be creating two go lang modules namely hello and tektutor.

Let's first begin with the hello module. We need to place the module code in a separate directory
```
cd ~
mkdir hello
cd hello
go mod init rps.com/hello
go mod tidy
```
The above go mod command, will create a go.mod file capturing the module name i.e rps.com/hello.  In addition to that, it will also capture the go version that is required to use the custom module.  The go mod init is used while creating a new go module.

The go mod tidy command will download the dependencies mentioned in the go.mod file.

Within the ~/hello folder, create a file named hello.go with the below content
<pre>
package hello

import "fmt"

func SayHello( name string ) string {
	message := fmt.Sprintf("Hello, %v !", name ) 
	return message
}	
</pre>

Now, let's step up of hello folder and create another folder name tektutor at the same level as hello module.
```
cd ~
mkdir tektutor
cd tektutor
go mod init tektutor.org/tektutor
go mod tidy
```

Now, let's create a file named main.go with the below content
<pre>
package main 

import (
	"fmt"

	"tektutor.org/hello"
)

func main() {
	msg := hello.SayHello("Golang")
	fmt.Println(msg)
}	
</pre>

Let's try to run the above main.go
```
go run ./main.go
```
We are supposed get an error saying unable to download tektutor.org/hello module as it is not published in public terraform registry(website). Hence, we are going to point to our local hello module directory as shown below
```
cd ~/tektutor
go mod edit --replace tektutor.org/hello=../hello
go mod tidy
```

If everything went well, we should be able to run the main.go
```
cd ~/tektutor
go run ./main.go
```

Expected output
