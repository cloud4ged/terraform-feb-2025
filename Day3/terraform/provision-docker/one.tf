# Pull ubuntu:20.04 docker image from Docker Hub Remote Registry to Local Docker Registry
resource "docker_image" "ubuntu" {
  name = "ubuntu:20.04"
}

