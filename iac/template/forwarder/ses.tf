resource "aws_ses_receipt_rule" "send_to_s3" {
  name          = "email-conceal-forward-${var.environment}"
  rule_set_name = "INBOUND_MAIL"
  enabled       = true

  recipients = [var.domain]

  scan_enabled = false
  tls_policy   = "Optional"

  s3_action {
    position    = 1
    bucket_name = aws_s3_bucket.email_storage.id
  }
}
