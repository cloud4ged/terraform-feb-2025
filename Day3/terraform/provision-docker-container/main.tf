# Pull ubuntu:20.04 docker image from Docker Hub Remote Registry to Local Docker Registry
resource "docker_image" "ubuntu" {
  name = "ubuntu:latest"
}

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
