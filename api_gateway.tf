locals {
  base_domain = "api-${var.stage}.${var.domain_name}"
  base_path   = "hello"
}

resource "aws_api_gateway_rest_api" "hello_api" {
  name        = "hello_api"
  description = "Greeting API"

  disable_execute_api_endpoint = true
}

resource "aws_api_gateway_domain_name" "hello_api_domain_name" {
  domain_name              = local.base_domain
  regional_certificate_arn = var.acm_certificate

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

resource "aws_route53_record" "hello_api_regional_record" {
  name    = local.base_domain
  type    = "A"
  zone_id = var.route_domain_zone_id

  alias {
    evaluate_target_health = true
    name                   = aws_api_gateway_domain_name.hello_api_domain_name.regional_domain_name
    zone_id                = aws_api_gateway_domain_name.hello_api_domain_name.regional_zone_id
  }
}

resource "aws_api_gateway_deployment" "hello_deployment" {
  rest_api_id = aws_api_gateway_rest_api.hello_api.id
  stage_name  = aws_api_gateway_stage.hello_stage.stage_name
  depends_on  = [aws_api_gateway_integration.lambda_hello]

  triggers = {
    redeployment = sha1(jsonencode([
      aws_api_gateway_rest_api.hello_api.body
    ]))
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "hello_stage" {
  deployment_id = aws_api_gateway_deployment.hello_deployment.id
  rest_api_id   = aws_api_gateway_rest_api.hello_api.id
  stage_name    = var.stage
}

resource "aws_api_gateway_base_path_mapping" "hello_world_path_mapping" {
  api_id      = aws_api_gateway_rest_api.hello_api.id
  stage_name  = aws_api_gateway_stage.hello_stage.stage_name
  domain_name = local.base_domain
  base_path   = local.base_path
}

output "base_url" {
  value = aws_api_gateway_deployment.api_gtw_deploy_hello.invoke_url
}