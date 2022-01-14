resource "aws_api_gateway_rest_api" "api_gateway" {
  name = "email-conceal-manager-${var.environment}"
}

resource "aws_api_gateway_resource" "proxy_resource" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  parent_id = aws_api_gateway_rest_api.api_gateway.root_resource_id
  path_part = "{proxy+}"
}

locals {
  http_methods = [
    "GET",
    "POST",
    "PUT",
    "DELETE",
    "OPTIONS",
    "HEAD",
  ]
}

resource "aws_api_gateway_method" "api_method" {
  for_each = toset(local.http_methods)

  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.proxy_resource.id

  http_method = each.value

  authorization = "NONE"
}

resource "aws_api_gateway_integration" "api_integration" {
  for_each = aws_api_gateway_method.api_method

  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.proxy_resource.id
  http_method = each.value.http_method


  type                    = "AWS_PROXY"
  integration_http_method = "POST"
  uri                     = aws_lambda_function.api_lambda.invoke_arn
}
