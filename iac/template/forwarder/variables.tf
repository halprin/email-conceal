variable "environment" {
  type        = string
  default     = "dev"
  description = "The dev or prod environment"
}

variable "concealed_email_prefix" {
  type        = string
  description = "The email address that is receiving the e-mail, but just the part to the left of the domain name and @"
}

variable "forward_email_prefix" {
  type        = string
  default     = "forwarder"
  description = "The email address that sends the e-mail onto the true receipient, but just the part to the left of the domain name and @"
}

variable "receiving_email" {
  type        = string
  description = "The email address that receives the e-mails"
}

variable "domain" {
  type        = string
  description = "The domain name that should be utilized"
}

variable "email_lifetime" {
  type        = number
  default     = 4
  description = "The time an e-mail sticks around until it is deleted, whether or not it was delivered.  Measured in days."
}

variable "docker_image" {
  type        = string
  description = "The Docker image URI to deploy"
}
