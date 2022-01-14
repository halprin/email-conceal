resource "aws_kms_key" "application_key" {
  description = "Key to encrypt email-conceal application in ${var.environment} environment"

  enable_key_rotation = true
  policy              = data.aws_iam_policy_document.s3_can_encrypt_for_sqs.json

  tags = {
    project     = local.project
    environment = var.environment
  }
}

resource "aws_kms_alias" "application_key_alias" {
  name          = "alias/email-conceal-${var.environment}"
  target_key_id = aws_kms_key.application_key.key_id
}

data "aws_iam_policy_document" "s3_can_encrypt_for_sqs" {
  statement {
    sid    = "DecryptDataForSqs"
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["s3.amazonaws.com"]
    }
    actions   = ["kms:GenerateDataKey", "kms:Decrypt"]
    resources = ["*"]
  }

  statement {
    sid    = "EncryptDataForSes"
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["ses.amazonaws.com"]
    }
    actions   = ["kms:GenerateDataKey", "kms:Encrypt"]
    resources = ["*"]
  }

  statement {
    sid    = "RootAccountCanStillAdministerKmsKey"
    effect = "Allow"
    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"]
    }
    actions   = ["kms:*"]
    resources = ["*"]
  }
}
