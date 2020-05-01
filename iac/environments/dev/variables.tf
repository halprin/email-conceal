variable "region" {
  type    = string
  default = "us-east-1"
}

variable "concealed_email_prefix" {
  type        = string
  description = "The email address, but just the part to the left of the domain name and @"
}

variable "receiving_email" {
  type        = string
  description = "The email address that receives the e-mails"
}

variable "domain" {
  type        = string
  description = "The domain name that should be utilized"
}

variable "docker_image" {
  type        = string
  description = "The Docker image URI to deploy"
}
