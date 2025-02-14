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

# Pull tektutor/ubuntu-ansible:1.0 docker image from Docker Hub Remote Registry to Local Docker Registry
data "docker_image" "tektutor" {
  name = var.image_name 
}

resource "docker_container" "tektutor_container1" {
  image		= data.docker_image.tektutor.name
  name		= var.container_name 
  must_run	= true

  ports {
    internal = "22"
    external = "2001"
  }

  ports {
    internal = "80"
    external = "8001"
  }
}

resource "local_file" "invoke_ansible_playbook" {
  content  = "172.17.0.2" 
  filename = "./ip.txt"

  provisioner "local-exec" {
    command = "ansible-playbook -i ./hosts install-nginx-playbook.yml"
  }

//  depends_on = [ time_sleep.wait_120_seconds ] 
}
