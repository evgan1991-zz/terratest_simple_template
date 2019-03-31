# Example module for running tests

## Description
The mediation module is used to run terratest tests. Terratest test does not test TF-files, but the module as a whole, so, an intermediary module was created that can call the main module, and also contains fields forbidden for the main module, such as a "region" in inputs or an "account id" in outputs
Optional for testing if the main module creates only local resources (for example, local files) and does not apply to aws

## Usage
Add provider aws with region and aws_caller_identity data. Call module with all parameters like (param = "${var.param}"). Use all variables from variables.tf for test all cases in your module after.

```hcl
provider "aws" {
  region = "${var.region}"
}

data "aws_caller_identity" "current" {}

module "example_name_for_module" {
  source   = "../"
  example  = "${var.example}"
  example2 = "${var.example2}"
}
```


## Inputs
Fully duplicate Inputs from the main module. Add and set a region (use any location)
```hcl
variable "region" {
  default = "us-east-1"
}
```

| Name | Description | Type | Default | Required |
|------|-------------|:----:|:-----:|:-----:|
| region | Region on which test resources will be created | string | us-east-1 | no |
| example | Example variable | string | example | no |
| example2 | Example variable2 | string | example2 | no |



## Outputs
Fully duplicate Outputs from the main module and add account_id, caller_arn and caller_user
```hcl
output "account_id" {
  value = "${data.aws_caller_identity.current.account_id}"
}

output "caller_arn" {
  value = "${data.aws_caller_identity.current.arn}"
}

output "caller_user" {
  value = "${data.aws_caller_identity.current.user_id}"
}
```

| Name | Description |
|------|-------------|
| example | - |
| example2 | - |
| account_id | Current AWS Account Id |
| caller_arn | Current AWS IAM-user ID  |
| caller_user | Current AWS IAM-user name |


## Terraform versions

Terraform version 0.11.11 or newer is required for this module to work.
