provider "aws" {
  region = "${var.region}"
}

data "aws_caller_identity" "current" {}

module "example" {
  source   = "../"
  example  = "${var.example}"
  example2 = "${var.example2}"
}




