terraform {
   required_providers {
       docker = {
          source = "tektutor/docker"
	  version = "1.0.0"
       }
   }
}

resource "docker_container" "nginx1" {
	image_name = "bitnami/nginx:latest"
	container_name = "nginx"
	host_name = "nginx"
}
