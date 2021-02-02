resource "aws_guardduty_detector" "guardduty" {
  enable                       = true
  finding_publishing_frequency = "FIFTEEN_MINUTES"
  tags = {
    env      = var.environment
    app_name = var.app_name
  }
}
