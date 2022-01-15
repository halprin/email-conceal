resource "aws_lambda_function" "forwarder_lambda" {
  function_name = "email-conceal-forwarder-${var.environment}"

  filename         = data.archive_file.lambda_zip_archive.output_path
  source_code_hash = data.archive_file.lambda_zip_archive.output_base64sha256
  handler          = "forwarder"
  timeout          = 10
  memory_size      = 128
  runtime          = "go1.x"

  environment {
    variables = {
      DOMAIN                 = var.domain
      TABLE_NAME             = var.configuration_database_name
      ENVIRONMENT            = var.environment
      FORWARDER_EMAIL_PREFIX = var.forward_email_prefix
    }
  }

  role = aws_iam_role.execution_role.arn

  tags = {
    project     = local.project
    environment = var.environment
  }
}

data "archive_file" "lambda_zip_archive" {
  type        = "zip"
  source_file = "${path.module}/../../../src/forwarder"
  output_path = "${path.module}/forwarder_lambda.zip"
}

resource "aws_lambda_event_source_mapping" "attach_sqs" {
  event_source_arn = aws_sqs_queue.email_storage_add_event_queue.arn
  function_name    = aws_lambda_function.forwarder_lambda.arn

  batch_size                         = 10
  maximum_batching_window_in_seconds = 0

  function_response_types = ["ReportBatchItemFailures"]
}
