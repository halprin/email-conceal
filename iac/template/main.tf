module "forwarder" {
  source = "./forwarder/"

  configuration_database_name = aws_dynamodb_table.configuration.name
  docker_image = var.docker_image
  domain = var.domain
}
