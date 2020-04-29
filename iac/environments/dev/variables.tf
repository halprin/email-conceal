variable "region" {
  type = string
  default = "us-east-1"
}

variable "email_address" {
  type = string
  description = "The email address, but just the part to the left of the domain name and @"
}

variable "domain" {
  type = string
  description = "The domain name that should be utilized"
}
