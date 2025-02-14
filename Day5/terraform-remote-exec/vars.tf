variable "container_name" {
  description = "Name for the Docker container"
  type        = string
  default     = "c1"
}

variable "image_name" {
  description = "Name for the Docker Image"
  type        = string
  default     = "tektutor/ubuntu-ansible:1.0"
}


