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

variable "configuration_database_name" {
  type        = string
  description = "Name to the configuration database"
}

variable "application_key_arn" {
  type        = string
  description = "The KMS key to encrypt things with"
}
