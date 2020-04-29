resource "aws_s3_bucket" "email_storage" {
  bucket = data.null_data_source.names.outputs.bucket_name
  acl = "private"

  policy = data.aws_iam_policy_document.ses_write_to_s3.json

  lifecycle_rule {
    id = "DeleteAfterTime"
    enabled = true

    expiration {
      days = var.email_lifetime
    }
  }

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "aws:kms"
        kms_master_key_id = aws_kms_key.application_key.arn
      }
    }
  }

  tags = {
    environment = var.environment
  }
}

resource "aws_s3_bucket_notification" "add_email_notification" {
  bucket = aws_s3_bucket.email_storage.id

  queue {
    id = "NotifyQueue"
    queue_arn = aws_sqs_queue.email_storage_add_event_queue.arn
    events = ["s3:ObjectCreated:*"]
  }

  depends_on = [aws_kms_key.application_key]  //depend on the KMS key, with its policy allowing the S3 bucket to send encrypted messages to the SQS queue
}

data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "ses_write_to_s3" {
  statement {
    sid = "AllowSESPuts"
    effect = "Allow"
    principals {
      type = "Service"
      identifiers = ["ses.amazonaws.com"]
    }
    actions = ["s3:PutObject"]
    resources = ["arn:aws:s3:::${data.null_data_source.names.outputs.bucket_name}/*"]
    condition {
      test = "StringEquals"
      variable = "aws:Referer"
      values = [data.aws_caller_identity.current.account_id]
    }
  }
}