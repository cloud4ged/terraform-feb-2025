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
