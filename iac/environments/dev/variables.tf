variable "region" {
  type    = string
  default = "us-east-1"
}

variable "domain" {
  type        = string
  description = "The domain name that should be utilized"
}

variable "docker_image" {
  type        = string
  description = "The Docker image URI to deploy"
}
