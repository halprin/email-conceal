variable "environment" {
  type        = string
  default     = "dev"
  description = "The dev or prod environment"
}

variable "forward_email_prefix" {
  type        = string
  default     = "forwarder"
  description = "The email address that sends the e-mail onto the true receipient, but just the part to the left of the domain name and @"
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
