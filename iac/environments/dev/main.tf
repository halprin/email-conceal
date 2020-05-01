provider "aws" {
  region = var.region
}

terraform {
  backend "s3" {
    bucket         = "terraforms-state"
    key            = "email-conceal-us-east-1/dev.tfstate"
    region         = "us-east-1"
    dynamodb_table = "terraform_lock"
  }
}

module "forwarder" {
  source = "../../template/forwarder/"

  environment = "dev"

  email_lifetime         = 4
  concealed_email_prefix = var.concealed_email_prefix
  receiving_email        = var.receiving_email
  domain                 = var.domain
  docker_image           = var.docker_image
}
