resource "aws_iam_role" "ecs_task_role" {
  name = "email-conceal-forwarder-${var.environment}"

  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_policy" "permissions_for_forwarder" {
  name   = "email-conceal-forwarder-${var.environment}"
  policy = data.aws_iam_policy_document.permissions.json
}

data "aws_iam_policy_document" "permissions" {
  statement {
    sid       = "GetObjectFromEmailStore"
    effect    = "Allow"
    actions   = ["s3:GetObject"]
    resources = ["${aws_s3_bucket.email_storage.arn}/*"]
  }

  statement {
    sid    = "DecryptSqSAndS3BucketStuff"
    effect = "Allow"
    actions = [
      "kms:Decrypt",
      "kms:DescribeKey",
    ]
    resources = [var.application_key_arn]
  }

  statement {
    sid       = "SendEmail"
    effect    = "Allow"
    actions   = ["ses:SendRawEmail"]
    resources = ["arn:aws:ses:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:identity/${var.domain}"]
  }

  statement {
    sid    = "WorkWithEventQueue"
    effect = "Allow"
    actions = [
      "sqs:GetQueueUrl",
      "sqs:ReceiveMessage",
      "sqs:DeleteMessage",
    ]
    resources = [aws_sqs_queue.email_storage_add_event_queue.arn]
  }

  statement {
    sid    = "ReadConfigurationFromDynamo"
    effect = "Allow"
    actions = [
      "dynamodb:Query",
    ]
    resources = ["arn:aws:dynamodb:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:table/${var.configuration_database_name}"]
  }
}

resource "aws_iam_role_policy_attachment" "attach_permission_to_role" {
  role       = aws_iam_role.ecs_task_role.name
  policy_arn = aws_iam_policy.permissions_for_forwarder.arn
}

data "aws_iam_role" "execution_role" {
  name = "ecsTaskExecutionRole"
}
