resource "docker_image" "nginx18" {
  name = "nginx:1.18"
}

resource "docker_image" "nginx19" {
  name = "nginx:1.19"
}

resource "docker_image" "nginx20" {
  name = "nginx:1.20"
}
