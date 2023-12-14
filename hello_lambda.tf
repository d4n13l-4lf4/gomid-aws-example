locals {
  gomid_aws_example_hello_s3_key = "${local.service_name}/${var.stage}/${local.service_name}-${var.stage}-hello.zip"
  hello_lambda                   = "${local.service_name}-${var.stage}-hello"
}

data "local_file" "lambda_hello_code" {
  filename = "${path.module}/build/hello.zip"
}

resource "aws_lambda_function" "gomid_aws_example_hello" {
  function_name = local.hello_lambda

  s3_bucket = aws_s3_object.lambda_hello_code.bucket
  s3_key    = aws_s3_object.lambda_hello_code.key

  handler       = "bootstrap"
  runtime       = "provided.al2023"
  architectures = ["arm64"]
  timeout       = 29

  publish      = true
  skip_destroy = true

  source_code_hash = data.local_file.lambda_hello_code.content_base64sha256


  role = aws_iam_role.gomid_aws_example_hello_role.arn
}

data "aws_iam_policy_document" "gomid_aws_example_hello_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "gomid_aws_example_hello_role" {
  name               = "${local.hello_lambda}-lambda-role"
  assume_role_policy = data.aws_iam_policy_document.gomid_aws_example_hello_role.json
}

resource "aws_api_gateway_resource" "hello" {
  rest_api_id = aws_api_gateway_rest_api.hello_api.id
  parent_id   = aws_api_gateway_rest_api.hello_api.root_resource_id
  path_part   = "greet"
}

resource "aws_api_gateway_method" "hello" {
  rest_api_id   = aws_api_gateway_rest_api.hello_api.id
  resource_id   = aws_api_gateway_resource.hello.id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "lambda_hello" {
  rest_api_id = aws_api_gateway_rest_api.hello_api.id
  resource_id = aws_api_gateway_resource.hello.id
  http_method = aws_api_gateway_method.hello.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = module.hello_lambda_alias_refresh.lambda_alias_invoke_arn
}

resource "aws_api_gateway_deployment" "api_gtw_deploy_hello" {
  depends_on  = [aws_api_gateway_integration.lambda_hello]
  rest_api_id = aws_api_gateway_rest_api.hello_api.id
  stage_name  = aws_api_gateway_stage.hello_stage.stage_name
}

resource "aws_lambda_permission" "api_gtw_hello" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.gomid_aws_example_hello.function_name
  qualifier     = module.hello_lambda_alias_refresh.lambda_alias_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.hello_api.execution_arn}/*/*"
}

resource "aws_cloudwatch_log_group" "gomid_aws_example_hello" {
  name              = "/aws/lambda/${aws_lambda_function.gomid_aws_example_hello.function_name}"
  retention_in_days = 7
}

resource "aws_iam_role_policy_attachment" "lambda_basic_role" {
  role       = aws_iam_role.gomid_aws_example_hello_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_s3_object" "lambda_hello_code" {
  bucket = var.deployment_bucket
  key    = local.gomid_aws_example_hello_s3_key
  source = "${path.module}/build/hello.zip"
  etag   = filemd5("${path.module}/build/hello.zip")
}

# Lambda alias
module "hello_lambda_alias_refresh" {
  source  = "terraform-aws-modules/lambda/aws//modules/alias"
  version = "6.5.0"

  depends_on = [aws_lambda_function.gomid_aws_example_hello]

  name          = var.stage
  refresh_alias = false
  function_name = aws_lambda_function.gomid_aws_example_hello.function_name

  function_version = aws_lambda_function.gomid_aws_example_hello.version
}

# CodeDeploy
module "hello_lambda_code_deploy" {
  source  = "terraform-aws-modules/lambda/aws//modules/deploy"
  version = "6.5.0"

  aws_cli_command = "aws --region ${data.aws_region.current.id}"

  depends_on = [module.hello_lambda_alias_refresh]

  alias_name             = module.hello_lambda_alias_refresh.lambda_alias_name
  function_name          = aws_lambda_function.gomid_aws_example_hello.function_name
  deployment_config_name = "CodeDeployDefault.LambdaLinear10PercentEvery1Minute"

  target_version = aws_lambda_function.gomid_aws_example_hello.version

  create_app = true
  app_name   = local.service_name

  create_deployment_group = true
  deployment_group_name   = local.hello_lambda

  create_deployment          = true
  run_deployment             = true
  wait_deployment_completion = false

}