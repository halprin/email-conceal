resource "aws_dynamodb_table" "configuration" {
  name = "email-conceal-${var.environment}"

  billing_mode   = "PROVISIONED"
  read_capacity = 1
  write_capacity = 1

  hash_key = "primary"
  range_key      = "secondary"

  attribute {
    name = "primary"
    type = "S"
  }
  attribute {
    name = "secondary"
    type = "S"
  }

  tags = {
    environment = var.environment
  }
}