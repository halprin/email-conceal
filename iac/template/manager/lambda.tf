resource "aws_lambda_function" "api_lambda" {
  function_name = "email-conceal-manager-${var.environment}"

  filename = data.archive_file.lambda_zip_archive.output_path
  source_code_hash = data.archive_file.lambda_zip_archive.output_base64sha256
  handler = "manager"
  timeout = 10
  memory_size = 128
  runtime = "go1.x"

  environment {
    variables = {
      DOMAIN                 = var.domain
      TABLE_NAME             = var.configuration_database_name
      ENVIRONMENT            = var.environment
      FORWARDER_EMAIL_PREFIX = var.forward_email_prefix
      GIN_MODE               = var.environment == "prod" ? "release" : "debug"
    }
  }

  role = aws_iam_role.execution_role.arn

  tags = {
    project     = local.project
    environment = var.environment
  }
}

data "archive_file" "lambda_zip_archive" {
  type             = "zip"
  source_file      = "${path.module}/../../../src/manager"
  output_path      = "${path.module}/manager_lambda.zip"
}
