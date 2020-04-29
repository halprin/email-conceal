resource "aws_kms_key" "application_key" {
  description = "Key to encrypt email-conceal application in ${var.environment} environment"
  enable_key_rotation = true
  tags = {
    environment = var.environment
  }
}
