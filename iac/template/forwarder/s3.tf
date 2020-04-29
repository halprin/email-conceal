resource "aws_s3_bucket" "email_storage" {
  bucket = "email-conceal-${var.environment}"
  acl = "private"

  policy = data.aws_iam_policy_document.ses_write_to_s3.json

  lifecycle_rule {
    id = "DeleteAfterTime"
    enabled = true

    expiration {
      days = var.email_lifetime
    }
  }

  tags = {
    environment = var.environment
  }
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
    resources = ["arn:aws:s3:::email-conceal-${var.environment}/*"]
    condition {
      test = "StringEquals"
      variable = "aws:Referer"
      values = [data.aws_caller_identity.current.account_id]
    }
  }
}