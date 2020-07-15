resource "aws_guardduty_detector" "guardduty" {
  enable = true
  tags = {
    env      = var.environment
    app_name = var.app_name
  }
}
