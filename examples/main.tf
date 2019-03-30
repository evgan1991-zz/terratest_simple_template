variable "access_key" {
  default = "EXAMPLEKEYFORTESTAWS"
}
variable "secret_key" {
  default = "ExampleSecretForTestPipelineNotRequired!"
}

provider "aws" {
  region = "${var.region}"
  access_key = "${var.access_key}"
  secret_key = "${var.secret_key}"
}

data "aws_caller_identity" "current" {}

module "example" {
  source   = "../"
  example  = "${var.example}"
  example2 = "${var.example2}"
}
