resource "aws_cloudwatch_event_rule" "guardduty" {
  name        = "guardduty-lambda"
  description = "Send GuardDuty events to Lambda"
  role_arn = aws_lambda_permission.allow_cloudwatch.arn
  is_enabled = true
  tags = {
    env      = var.environment
    app_name = var.app_name
  }

  event_pattern = <<PATTERN
{
  "source": [
    "aws.guardduty"
  ],
  "detail-type": [
    "GuardDuty Finding"
  ],
  "detail": {
    "severity": [
      1,
      8
    ]
  }
}
PATTERN
}


resource "aws_cloudwatch_event_target" "lambda_target" {
    rule = aws_cloudwatch_event_rule.guardduty.name
    arn = aws_lambda_function.lambda.arn
}