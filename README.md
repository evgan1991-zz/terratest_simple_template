# Terraform Basic Example

## Description
[ORIGINAL MODULE](https://github.com/gruntwork-io/terratest/tree/master/examples/terraform-basic-example)
This folder contains a very simple Terraform module to demonstrate how you can use Terratest to write automated tests
for your Terraform code. This module takes in an input variable called `example`, renders it using a `template_file`
data source, and outputs the result in an output variable called `example`.

## Usage

```hcl
data "template_file" "example" {
  template = "${var.example}"
}

data "template_file" "example2" {
  template = "${var.example2}"
}

resource "local_file" "example" {
  content  = "${data.template_file.example.rendered} + ${data.template_file.example2.rendered}"
  filename = "example.txt"
}

resource "local_file" "example2" {
  content  = "${data.template_file.example2.rendered}"
  filename = "example2.txt"
}
```

## Running this module manually

1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`.
1. Run `terraform init`.
1. Run `terraform apply`.
1. When you're done, run `terraform destroy`.



## Running automated tests against this module

1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`.
1. Install [Golang](https://golang.org/) and make sure this code is checked out into your `GOPATH`.
1. `cd test`
1. `dep ensure`
1. `go test -v -run TestTerraformBasicExample`




## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|:----:|:-----:|:-----:|
| example | Example variable | string | example | no |
| example2 | Example variable2 | string | example2 | no |



## Outputs

| Name | Description |
|------|-------------|
| example | - |
| example2 | - |


## Terraform versions

Terraform version 0.11.11 or newer is required for this module to work.

## Contributing

None

## License

Apache2.0 Licensed.

## Authors

Yauheni Anashkin
