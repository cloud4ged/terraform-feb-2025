# Day 2

## Lab - Starting Ansible automation platform
```
minikube status
minikube start
kubectl get pods -n ansible-awx
kubectl get svc -n ansible-awx
minikube service awx-demo-service --url -n ansible-awx
```

Expected output
![image](https://github.com/user-attachments/assets/8d5f02cf-75cd-43a6-a84d-f6ce776b9846)

AWX Login Credentials
```
username - admin
```

To retrieve password
```
kubectl get secret -n ansible-awx | grep -i password
kubectl get secret awx-demo-admin-password -o jsonpath="{.data.password}" -n ansible-awx | base64 --decode; echo
```

To access the awx dashboard from other machines, you need to do port-forwarding
```
kubectl port-forward service/awx-demo-service -n ansible-awx --address 0.0.0.0 10445:80 &> /dev/null &
```

You may now access the dashboard from other machines as
```
http://10.0.1.72:31225
```

Expected output
![image](https://github.com/user-attachments/assets/3081631a-89f0-46cc-b84d-412c0ea41dd0)
![image](https://github.com/user-attachments/assets/6c8a2104-c97f-4f74-80cd-d9adb6443dc0)


## Lab - Creating a project in Ansible Tower
Navigate to Resource --> Projects
![image](https://github.com/user-attachments/assets/ee41a93c-96cc-44f5-b799-bc2913d25f2c)

Click "Add" button
![image](https://github.com/user-attachments/assets/0169dd9c-89c1-42d5-a006-9a99674b3f6a)
Select "Git" under Source Control Type
![image](https://github.com/user-attachments/assets/9d9b571c-dfa4-4a04-84e0-69aeddd4c094)
![image](https://github.com/user-attachments/assets/96644d63-8b22-4902-a888-f0fa4b412a77)
![image](https://github.com/user-attachments/assets/5dfeb1aa-7fa4-492d-93ae-3f289a7abe85)

Click "Save" button
![image](https://github.com/user-attachments/assets/fcfe69a8-2fff-41c9-b353-16f811affb69)
![image](https://github.com/user-attachments/assets/47964725-b73d-466f-b865-3cc2d39f657a)
![image](https://github.com/user-attachments/assets/cd1c6d54-417c-4dfc-92e6-066cb1a8e0ab)

## Lab - Let's create a credential in Ansible Tower to store the private key we generated

Navigate to Resources --> Credentials
![image](https://github.com/user-attachments/assets/59c2b03d-6c02-4fc0-b9df-a944ea598b6f)

Click "Add" button
![image](https://github.com/user-attachments/assets/6d9007c9-06d9-4d06-9e04-441b4ccf7420)
![image](https://github.com/user-attachments/assets/09414934-4e19-4fdb-9d26-c433c951e630)
![image](https://github.com/user-attachments/assets/dd1cae37-154a-4aaf-8727-62f91b06f80f)
![image](https://github.com/user-attachments/assets/349c5f61-8da1-4c4d-b97e-922aa5c977d2)

Click "Save" button
![image](https://github.com/user-attachments/assets/d2d2503c-6d65-4f2a-9d4e-0e694e2b10a3)

## Lab - Let's create an inventory in Ansible Tower

Navigate to Resources --> Inventories
![image](https://github.com/user-attachments/assets/fa42b60e-920b-40f0-abbd-725cb4c7f9de)

Click "Add" button
![image](https://github.com/user-attachments/assets/13b13825-5d00-4b4d-80da-f121baa66319)

Select "Add Inventory"
![image](https://github.com/user-attachments/assets/fc05d87c-b27d-4ed1-8118-a2a76b9d365b)

Type some name for the inventory, may be "Docker Inventory"
![image](https://github.com/user-attachments/assets/afc26d16-af2e-4bfe-a743-6da6fb9ef59a)

Click "Save" button
![image](https://github.com/user-attachments/assets/5bf49418-6dd8-432c-bdd1-e96f4e13da2c)

Click on the "Hosts" Tab
![image](https://github.com/user-attachments/assets/62fe33dc-18d5-4d1e-b57f-49096f1a8e80)

Click on "Add" button
![image](https://github.com/user-attachments/assets/36c5152a-1015-4b14-bf3a-8804fbbfe103)

Let's add Ubuntu1 connection details as shown below
![image](https://github.com/user-attachments/assets/679d60ee-f05a-4e25-b96c-5cf5cee9ca68)
Click on "Save" button
![image](https://github.com/user-attachments/assets/13ca7028-d350-43d8-8c97-cdd005ef7833)

Let's add Ubuntu2 connection details as shown below
Click on Hosts tab
![image](https://github.com/user-attachments/assets/fda47161-2e7e-45e4-800d-91c06bdd367f)
Click on "Save" button
![image](https://github.com/user-attachments/assets/ddc3d6d3-adcc-4d06-92f7-92a78e60be1a)

Let's add Rocky1 connection details as shown below
Click on Hosts tab
![image](https://github.com/user-attachments/assets/c897a76e-2478-4edf-a7d2-18b8d81cd610)
Click on "Save" button

Let's add Rocky2 connection details as shown below
Click on Hosts tab
![image](https://github.com/user-attachments/assets/df055741-51f5-48de-ae26-a4e501af8edd)

Click on "Save" button
![image](https://github.com/user-attachments/assets/5d0d302e-f486-4cb9-a598-ca2164018063)


Now your hosts tab would look as shown below
![image](https://github.com/user-attachments/assets/bb7b00ea-4891-4a8b-9301-6c65ec970be6)

## Lab - Let's create a Template(Job Template) in Ansible Tower to run an ansible playbook

Navigate to Resources --> Templates
![image](https://github.com/user-attachments/assets/b003f499-6efd-4876-8459-d4eb6a6cf416)

Click on "Add" button and select "Add job template"
![image](https://github.com/user-attachments/assets/6546cc33-7b75-47e2-bc51-1f24951e7466)
![image](https://github.com/user-attachments/assets/d628f135-af10-44bb-b8d1-4ab330e47049)
![image](https://github.com/user-attachments/assets/9fa432a0-81ad-450f-88d1-aa7aa98ae4f4)
![image](https://github.com/user-attachments/assets/0fb1a926-11de-4d66-a688-b80835718cd6)
![image](https://github.com/user-attachments/assets/cab8cfdb-4224-473a-84bd-99ecac0429b8)
![image](https://github.com/user-attachments/assets/8ba518bf-3ef8-4243-aefc-be9b9fc12926)
![image](https://github.com/user-attachments/assets/9af670f4-c188-40fc-b5ee-96a0980e3992)
![image](https://github.com/user-attachments/assets/a66bba06-f5fd-427f-b149-4069c5f60c4a)
![image](https://github.com/user-attachments/assets/f05c646e-cb02-4364-80bd-d79b69490a93)

Click on "Save" button
![image](https://github.com/user-attachments/assets/83e1584d-307e-49d2-be3e-c2672076b6c8)

To run the playbook, click the "Launch" button
![image](https://github.com/user-attachments/assets/a8a73519-4b3a-49cc-bdb4-f244493f52e6)
![image](https://github.com/user-attachments/assets/ec647343-d59b-4825-9255-8ca807223323)
Click on "Launch" button
![image](https://github.com/user-attachments/assets/96c62946-0777-4306-ba62-f7e667ec99d3)
![image](https://github.com/user-attachments/assets/7b8d7fc0-dd68-4877-8624-69224836a673)
![image](https://github.com/user-attachments/assets/ed5ff31d-6eae-4063-b053-98a1c7571992)
![image](https://github.com/user-attachments/assets/af34b233-0f4d-4c09-9c91-0fb0b61e2942)

## Info - Golang overview
<pre>
- a programming developed and maintained by Google
- it is an open source progamming language
- popular tools like Docker, Kubernetes and Openshift were all developed using Go lang
- its syntax resembles C language, originally developed in C/C++ but later rebuild using go lang
- it is a compiled language
- For example
  - python is an interpretted language, where each line of code is interpretted while running the script
  - golang source is compiled like C++ to its machine executable format, hence all the compiler errors are reported at one shot unlike python or any other interpretted programming language
  - unlike C/C++, the memory management is handled by golang itself, it supports pointers like C/C++
  - when golang source code is compiled, it is observed the compilation is faster than any interpretted programming languages like python, ruby, etc.,
  - while golang compilation is slighly slower than C/C++
  - golang pretty much can be used to develop system tools just like C/C++, while it also offer modern features like python
</pre>

## Lab - Checking the go lang version
```
go version
```

Expected output
![image](https://github.com/user-attachments/assets/435162a2-fceb-4f9f-b177-de5bf497ead3)

## Lab - Writing a hello program in go lang

Create a file named hello.go with the below content
<pre>
package main

import "fmt"

func main() {
  fmt.Println( "Hello Golang !" )
}  
</pre>

To run your go program
```
go run ./hello.go
```

If you wish to build the hello.go to an exe that can be directly executed
```
go build ./hello.go
ls
./hello
```

Expected output
![image](https://github.com/user-attachments/assets/34f85c03-c50a-4922-b5a0-6b743ae15c5b)

## Lab - Go variables

Let's create file named variables.go with the below content
<pre>
package main

import "fmt"

func main() {
   var name = "James Gosling"
   fmt.Println(name)

   var firstNumber, secondNumber int = 10, 20
   fmt.Println( "Value of first number is ", firstNumber )
   fmt.Println( "Value of second number is ", secondNumber )

   var thirdNumber = 30
   fmt.Println ( "Value of third number is ", thirdNumber )
}  
</pre>

Let's run the program as
```
go run ./variables.go
```

Expected output
![image](https://github.com/user-attachments/assets/1df67cf4-b862-4e12-9ed8-40822f97e86d)
![image](https://github.com/user-attachments/assets/5ed32b65-ed8f-4f14-ba29-adf59cf6343c)
![image](https://github.com/user-attachments/assets/e4282f4d-acb6-412b-8399-33bc31a8c588)
![image](https://github.com/user-attachments/assets/fe4d4125-ff6a-4c25-a214-9303bc501b5d)

## Info - Go basic types
<pre>
- bool
- string
- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64 uintptr
- byte ( alias of uint8 )
- rune ( alias for int32 - represents a unicode )
- float32 float64
- complex64 complex128
</pre>

## Lab - Print formatting

Create a file named print-formatting.go with the below content
<pre>
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

}  
</pre>

Let's run the program
```
go run ./print-formatting.go
```

Expected output
![image](https://github.com/user-attachments/assets/aa493b93-6cc6-4e01-9c5d-b443f279b546)
![image](https://github.com/user-attachments/assets/dd7a11e9-16d8-42db-83ac-f5275bc0b766)
![image](https://github.com/user-attachments/assets/fa71f931-e031-48d7-99c6-8eeb5987c5ab)
![image](https://github.com/user-attachments/assets/d8330549-38ee-4513-a656-7a2fb55c8ec5)

## Lab - Creating a text file

Create a file names files.go with the below content
<pre>
package main

import (
   "fmt"
   "os"
   "log"
)

func main() {

	myfile, err := os.Create("./myfile.txt")

	if err != nil {
	   log.Fatal(err)
	}

	str := "Some content"

	bytesWritten, err := myfile.WriteString( str + "\n")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Wrote %d bytes into the file\n", bytesWritten)
	myfile.Sync()
}  
</pre>

Run the program
```
go run ./files.go
```

Expected output
![image](https://github.com/user-attachments/assets/76109000-823b-428d-b783-108db9da24f6)

The updated files.go looks as shown below
<pre>
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
</pre>

To run the program
```
go run ./files.go
```

Expected output
![image](https://github.com/user-attachments/assets/460fb276-61ce-4d16-82de-2ec39b38491e)
