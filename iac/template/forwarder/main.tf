data "null_data_source" "names" {
  inputs = {
    queue_name = "email-conceal-forwarder-${var.environment}"
    bucket_name = "email-conceal-${var.environment}"
  }
}
