locals {
  subdomain_suffix = var.environment != "prod" ? "-dev" : ""
}

data "aws_route53_zone" "domain" {
  name = var.domain
}

data "aws_acm_certificate" "domain_certificate" {
  domain      = var.domain
  statuses    = ["ISSUED"]
  most_recent = true
}

resource "aws_api_gateway_domain_name" "domain_for_api" {
  domain_name     = "api${local.subdomain_suffix}.${var.domain}"
  certificate_arn = data.aws_acm_certificate.domain_certificate.arn
  security_policy = "TLS_1_2"

  endpoint_configuration {
    types = ["EDGE"]
  }

  tags = {
    project     = local.project
    environment = var.environment
  }
}

resource "aws_route53_record" "api_route" {
  name    = aws_api_gateway_domain_name.domain_for_api.domain_name
  type    = "A"
  zone_id = data.aws_route53_zone.domain.id

  alias {
    name                   = aws_api_gateway_domain_name.domain_for_api.cloudfront_domain_name
    zone_id                = aws_api_gateway_domain_name.domain_for_api.cloudfront_zone_id
    evaluate_target_health = true
  }
}
