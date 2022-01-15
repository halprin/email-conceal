provider "aws" {
  region = var.region
}

terraform {
  backend "s3" {
    bucket         = "terraforms-state"
    key            = "email-conceal-us-east-1/prod.tfstate"
    region         = "us-east-1"
    dynamodb_table = "terraform_lock"
  }
}

module "template" {
  source = "../../template/"

  environment = "prod"

  email_lifetime = 4
  domain         = var.domain
}
