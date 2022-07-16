resource "aws_cognito_user_pool" "authn_authz" {
  name = "${local.project}-user-pool-${var.environment}"

  username_attributes = "email"

  //password_policy is not specified to get the default settings

  mfa_configuration= "OFF"

  account_recovery_setting {
    recovery_mechanism {
      name     = "verified_email"
      priority = 1
    }
  }

  auto_verified_attributes = "email"
}
