locals {
  project = "email-conceal"
}

module "forwarder" {
  source = "./forwarder/"

  environment = var.environment
  configuration_database_name = aws_dynamodb_table.configuration.name
  domain                      = var.domain
  application_key_arn         = aws_kms_key.application_key.arn
}

module "manager" {
  source = "./manager/"

  environment = var.environment
  configuration_database_name = aws_dynamodb_table.configuration.name
  domain                      = var.domain
  application_key_arn         = aws_kms_key.application_key.arn
}

data "aws_caller_identity" "current" {}
