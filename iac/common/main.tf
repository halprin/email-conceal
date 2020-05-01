provider "aws" {
  region = var.region
}

terraform {
  backend "s3" {
    bucket         = "terraforms-state"
    key            = "email-conceal-us-east-1/common.tfstate"
    region         = "us-east-1"
    dynamodb_table = "terraform_lock"
  }
}
