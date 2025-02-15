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
go mod init tektutor.org/hello
go mod tidy
```
The above go mod command, will create a go.mod file capturing the module name i.e tektutor.org/hello.  In addition to that, it will also capture the go version that is required to use the custom module.  The go mod init is used while creating a new go module.

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
Expected output
![image](https://github.com/user-attachments/assets/a222ca36-4c34-4060-9918-616b8a6a55bc)


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
![image](https://github.com/user-attachments/assets/6cb4c9b1-7076-46ea-a2b8-1c6ab1b5e374)

## Lab - Golang concurrency and channel
```
cd ~/terraform-feb-2025
git pull
cd Day3/concurrency
go run ./channel.go
```

Expected output
![image](https://github.com/user-attachments/assets/6b10971d-d58b-4020-8766-358c2e403689)
![image](https://github.com/user-attachments/assets/f251a243-e445-40a4-962e-56ab8be8c369)

## Info - Terraform Overview
<pre>
- Terraform - a IOC, Infrastructure as a Code tool
- in otherwords, a provisioning tool
- helps in provisioning resources locally, private/public/hybrid cloud
- it is an alternate to AWS Cloudformation
- the advantage of using Terraform over AWS Cloudformation is, Terraform is cloud newtral(supports all public/private/hybrid cloud and locally too ), while Cloudformation supports only AWS
- Developed by a company called HashiCorp
- it comes in 2 flavours
  - open source
    - command-line
  - enterprise 
    - you get out of the box centralized state management
    - you also get world-wide support from Hashicorp 
    - supports web console interface
</pre>

## Info - Ansible vs Terraform
<pre>
- Terraform provision a new machine locally using VMWare, VirtualBox, Hyper-V, docker
- Terraform can also provision an ec2 instance in AWS, virtual machine in Azure, GCP, Digital Ocean, etc.,
- Terraform is a provisoner
- Terraform also has basic configuration management capabilities
  - Once a machine is provisioned successfully by terraform, terraform can invoke some shell scripts, power shells to further configure it
  - using shell scripts or powershell scripts is considered an imperative approach, not a declarative approach
- Ansible is a Configuration Management Tool, which helps us in
  - installing/uninstalling/upgrading softwares on an already provisioned machine
  - i.e The machine provisioned by Terraform can be configured further by Ansible
  - hence they solve two different problems, they are complementing tools not competing tools
  - ansible also has provisioning capabilities 	
</pre>


## Lab - Writing your first Terraform automation script

We need to first create a terraform project by creating a directory
```
cd ~
mkdir terraform-pull-image
cd terraform-pull-image
```

Inside the terraform-pull-image folder, create a file named main.tf
```
terraform {
  required_providers {
    docker = {
	source = "kreuzwerker/docker"
        version = "3.0.2"
    }
  }
}

# Creates an instance of the docker provider
provider "docker" {

}

# Pull ubuntu:20.04 docker image from Docker Hub Remote Registry to Local Docker Registry
resource "docker_image" "ubuntu" {
  name = "ubuntu:20.04"
}
```
Expected output
![image](https://github.com/user-attachments/assets/5d003c29-94a3-441b-a56e-b8ecc5280b9a)



Let's download the required provider plugins as shown below ( mandatory first step )
```
terraform init
ls -lha
tree .terraform
```
Expected output
![image](https://github.com/user-attachments/assets/ccddd455-9bc5-4817-b88c-24f89e372a57)
![image](https://github.com/user-attachments/assets/26d7f6bd-94e4-4a18-894a-17f78ad00a94)


Let's create plan to inspect what terraform will do if we run the script ( dry run - it won't really execute )
```
terraform plan
```
Expected output
![image](https://github.com/user-attachments/assets/168a29af-d0e3-4c38-b166-6ff26774d7c0)


Once you are happy with the terraform plan, you may apply it
```
terraform apply
```

Expected output
![image](https://github.com/user-attachments/assets/bdcfb0ac-27cb-42d4-9bd8-088e3d1c7e5c)
![image](https://github.com/user-attachments/assets/b7f150fb-5262-41b3-a8b1-ac5ac98ef480)
![image](https://github.com/user-attachments/assets/192fc775-5a4e-428f-ace5-e275fac31874)

Once you are done with this exercise, you may destroy all the resources we created via terraform
```
terraform destroy
```

Expected out
![image](https://github.com/user-attachments/assets/9d2361dd-cc6a-415c-9e87-580eb365c5fb)
![image](https://github.com/user-attachments/assets/5faf3c0e-f245-4681-bd7d-5aa499733b2e)
![image](https://github.com/user-attachments/assets/9d8a28b9-0214-4ff7-a30e-6ec72ad73a91)
![image](https://github.com/user-attachments/assets/a74414b7-bf6f-425d-a7bd-4eee93d2fd76)
![image](https://github.com/user-attachments/assets/c549f7eb-6c71-4b57-b4be-370bdde572aa)

You could update the main.tf as shown below
<pre>
terraform {
  required_providers {
    docker = {
	source = "kreuzwerker/docker"
        version = "3.0.2"
    }
  }
}

provider "docker" {

}

# Pull ubuntu:20.04 docker image from Docker Hub Remote Registry to Local Docker Registry
resource "docker_image" "ubuntu" {
  name = "ubuntu:20.04"
}

resource "docker_image" "nginx18" {
  name = "nginx:1.18"
}

resource "docker_image" "nginx19" {
  name = "nginx:1.19"
}

resource "docker_image" "nginx20" {
  name = "nginx:1.20"
}	
</pre>

Then you can try the below
```
terraform plan
terraform apply --auto-approve
docker images | grep nginx
```

Expected output
![image](https://github.com/user-attachments/assets/af9b3ddb-bfb1-4069-aee8-c9d7912edb14)
![image](https://github.com/user-attachments/assets/1793d9ea-1db4-4066-a34b-e43199c090d8)
![image](https://github.com/user-attachments/assets/ceb907d1-5097-4070-a250-ef65ce1fb104)
![image](https://github.com/user-attachments/assets/2f2824ee-f75d-45fb-8203-80b3bad22092)
![image](https://github.com/user-attachments/assets/fdd62ca1-a117-4e18-93b0-345794af49f2)
![image](https://github.com/user-attachments/assets/13a0b20f-db55-4713-986e-1fc2fd5989e4)
![image](https://github.com/user-attachments/assets/29723051-f9df-48d1-b532-2bfe34c33c57)

## Lab - Let's provision a docker container using terraform
```
cd ~
mkdir provision-docker-container
cd provision-docker-container
touch main.tf
touch providers.tf
```

Update your providers.tf file with the below content
<pre>
terraform {
  required_providers {
    docker = {
	source = "kreuzwerker/docker"
        version = "3.0.2"
    }
  }
}

provider "docker" {
}	
</pre>

Update your main.tf file with the below content
<pre>
# Pull ubuntu:20.04 docker image from Docker Hub Remote Registry to Local Docker Registry
resource "docker_image" "ubuntu" {
  name = "ubuntu:latest"
}
//HCL - Hashicorp Configuration Language
resource "docker_container" "ubuntu_container1" {
  image		= docker_image.ubuntu.name
  name		= "ubuntu1_cont"
  must_run	= true
  command	= [
	"tail",
	"-f",
	"/dev/null"
  ]
}	
</pre>

Now, let's download the provider plugins
```
cd ~/provision-docker-container
terraform init
ls -lha
tree .terraform
```
Expected ouput
![image](https://github.com/user-attachments/assets/2342e66b-0527-499c-9268-e72f6a082699)
![image](https://github.com/user-attachments/assets/79e33dd9-fbb1-4903-8a4c-7b292e07399d)

Let's check the plan
```
terraform plan
```
Expected ouput
![image](https://github.com/user-attachments/assets/1456df06-d4ce-40a6-a4bb-c1a5420b5b28)
![image](https://github.com/user-attachments/assets/4ae37286-fea8-4e35-af5e-ad900b13eb74)

Let's run the automation
```
terraform apply --auto-approve
```

Expected ouput
![image](https://github.com/user-attachments/assets/d7a5850a-70fa-4a95-82e1-8580b4a53c92)
![image](https://github.com/user-attachments/assets/b860ee7d-5c0d-4ca5-bb2f-c9d1d1c6017d)
