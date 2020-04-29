resource "aws_ecr_repository" "application_repository" {
  name                 = "email-conceal-forwarder"
  image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}
