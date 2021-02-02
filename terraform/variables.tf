variable "environment" {
  type    = string
  default = "any"
}

variable "slack_webhook" {
  type = string
}

variable "app_name" {
  type    = string
  default = "guardduty-alerts"
}

variable "lambda_bundle_s3_bucket" {
  type = string
}