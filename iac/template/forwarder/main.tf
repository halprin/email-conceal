locals {
  project     = "email-conceal"
  queue_name  = "email-conceal-forwarder-${var.environment}"
  bucket_name = "email-conceal-${var.environment}"
}