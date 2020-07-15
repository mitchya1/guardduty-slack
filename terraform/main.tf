provider "aws" {
  region  = "us-east-2"
  version = "2.67"
}

data "aws_caller_identity" "current" {}
data "aws_region" "current" {}
