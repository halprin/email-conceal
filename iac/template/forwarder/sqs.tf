resource "aws_sqs_queue" "email_storage_add_event_queue" {
  name = local.queue_name

  message_retention_seconds  = var.email_lifetime * 24 * 60 * 60
  visibility_timeout_seconds = 30

  kms_master_key_id                 = var.application_key_arn
  kms_data_key_reuse_period_seconds = 43200 //half a day

  policy = data.aws_iam_policy_document.s3_write_to_sqs.json

  tags = {
    environment = var.environment
  }
}

data "aws_region" "current" {}

data "aws_iam_policy_document" "s3_write_to_sqs" {
  statement {
    sid    = "AllowS3Sends"
    effect = "Allow"
    principals {
      type        = "*"
      identifiers = ["*"]
    }
    actions   = ["sqs:SendMessage"]
    resources = ["arn:aws:sqs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:${local.queue_name}"]
    condition {
      test     = "StringEquals"
      variable = "aws:SourceArn"
      values   = [aws_s3_bucket.email_storage.arn]
    }
  }
}
