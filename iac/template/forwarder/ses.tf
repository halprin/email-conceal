resource "aws_ses_receipt_rule" "send_to_s3" {
  name          = "email-conceal-forward-${var.environment}"
  rule_set_name = "INBOUND_MAIL"
  enabled       = true

  recipients = ["${var.concealed_email_prefix}@${var.domain}"]

  scan_enabled = false
  tls_policy   = "Optional"

  s3_action {
    position    = 1
    bucket_name = aws_s3_bucket.email_storage.id
  }

  depends_on = [aws_kms_key.application_key] //depend on the KMS key, with its policy allowing the SES rule to send encrypted messages to the S3 bucket
}
