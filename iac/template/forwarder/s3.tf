resource "aws_s3_bucket" "email_storage" {
  bucket = local.bucket_name
  acl    = "private"

  policy = data.aws_iam_policy_document.ses_write_to_s3.json

  lifecycle_rule {
    id      = "DeleteAfterTime"
    enabled = true

    expiration {
      days = var.email_lifetime
    }
  }

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm     = "aws:kms"
        kms_master_key_id = var.application_key_arn
      }
    }
  }

  tags = {
    project     = local.project
    environment = var.environment
  }
}

resource "aws_s3_bucket_notification" "add_email_notification" {
  bucket = aws_s3_bucket.email_storage.id

  queue {
    id        = "NotifyQueue"
    queue_arn = aws_sqs_queue.email_storage_add_event_queue.arn
    events    = ["s3:ObjectCreated:*"]
  }
}

data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "ses_write_to_s3" {
  statement {
    sid    = "AllowSESPuts"
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["ses.amazonaws.com"]
    }
    actions   = ["s3:PutObject"]
    resources = ["arn:aws:s3:::${local.bucket_name}/*"]
    condition {
      test     = "StringEquals"
      variable = "aws:Referer"
      values   = [data.aws_caller_identity.current.account_id]
    }
  }
}
