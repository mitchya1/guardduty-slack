resource "aws_lambda_function" "lambda" {
  function_name                  = format("%s-%s", var.environment, var.app_name)
  s3_bucket                      = var.lambda_bundle_s3_bucket
  s3_key                         = format("%s/latest.zip", var.environment)
  role                           = aws_iam_role.role.arn
  handler                        = "main" // binary will be named api
  memory_size                    = 128
  runtime                        = "go1.x"
  reserved_concurrent_executions = -1 // Disable concurrency limits
  timeout                        = 3  // 5 second timeout
  environment {
    variables = {
      SLACK_CHANNEL_ID = var.slack_channel_id
      SLACK_TOKEN      = var.slack_bot_token
    }
  }

  tags = {
    env      = var.environment
    app_name = var.app_name
  }
}