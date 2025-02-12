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
