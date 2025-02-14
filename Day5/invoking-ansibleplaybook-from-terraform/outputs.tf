output "container_id" {
  description = "Docker container ID"
  value       =  docker_container.tektutor_container1.id
}

output "image_id" {
  description = "Docker image ID"
  value       =  data.docker_image.tektutor.id
}
