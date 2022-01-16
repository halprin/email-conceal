variable "environment" {
  type        = string
  default     = "dev"
  description = "The dev or prod environment"
  nullable    = false
}

variable "forward_email_prefix" {
  type        = string
  default     = "forwarder"
  description = "The email address that sends the e-mail onto the true receipient, but just the part to the left of the domain name and @"
  nullable    = false
}

variable "domain" {
  type        = string
  description = "The domain name that should be utilized"
  nullable    = false
  validation {
    condition     = can(regex("^[^.]+\\.[^.]+$", var.domain))
    error_message = "The domain must be one or more non-period characters, followed by a period, followed by one or more non-period characters."
  }
}

variable "email_lifetime" {
  type        = number
  default     = 4
  description = "The time an e-mail sticks around until it is deleted, whether or not it was delivered.  Measured in days."
  nullable    = false
}
