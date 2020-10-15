resource "aws_cloudwatch_event_rule" "guardduty" {
  name        = "guardduty-lambda"
  description = "Send GuardDuty events to Lambda"
  //role_arn = aws_lambda_permission.allow_cloudwatch.arn
  is_enabled = true
  tags = {
    env      = var.environment
    app_name = var.app_name
  }
  // Severity levels https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings.html
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
      4,
      4.0,
      4.1,
      4.2,
      4.3,
      4.4,
      4.5,
      4.6,
      4.7,
      4.8,
      4.9,
      5,
      5.0,
      5.1,
      5.2,
      5.3,
      5.4,
      5.5,
      5.6,
      5.7,
      5.8,
      5.9,
      6,
      6.0,
      6.1,
      6.2,
      6.3,
      6.4,
      6.5,
      6.6,
      6.7,
      6.8,
      6.9,
      7,
      7.0,
      7.1,
      7.2,
      7.3,
      7.4,
      7.5,
      7.6,
      7.7,
      7.8,
      7.9,
      8,
      8.0,
      8.1,
      8.2,
      8.3,
      8.4,
      8.5,
      8.6,
      8.7,
      8.8,
      8.9
    ]
  }
}
PATTERN
}


resource "aws_cloudwatch_event_target" "lambda_target" {
    rule = aws_cloudwatch_event_rule.guardduty.name
    arn = aws_lambda_function.lambda.arn
}

data "aws_cloudwatch_log_group" "lambda_log_group" {
  name = format("/aws/lambda/%s-%s", var.environment, var.app_name)
  retention_in_days = 3

  tags = {
    env = var.environment
    app_name = var.app_name
  }
}