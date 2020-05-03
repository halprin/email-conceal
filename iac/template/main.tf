module "forwarder" {
  source = "./forwarder/"

  configuration_database_name = aws_dynamodb_table.configuration.name
  docker_image = var.docker_image
  domain = var.domain
  application_key_arn = aws_kms_key.application_key.arn
}

data "aws_caller_identity" "current" {}
