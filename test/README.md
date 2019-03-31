# Integration test using Terratest and Go

## Description
Sample test for tf-module. The "test/" directory may contain not one test, but many tests divided by different files. To use tests, you must install "Terraform", "Go" and download the following libraries with "go get": "[terratest](https://github.com/gruntwork-io/terratest/modules/terraform)", "[assert](https://github.com/stretchr/testify/assert)"
All modules must end with "_test.go" (example: iam_test.go alb_test.go MyAmazingModule_test.go)
All test-functions must start with "Test" (example: TestCreating(t *testing.T){} TestConnecting(t *testing.T){} TestRunning(t *testing.T){} )


## Usage

Run this for start testing
```hcl
go test
```
or
```hcl
go test ./path_to_go_file
```

## Important things

simple function
```hcl
package test

import (
    "testing"
    "github.com/gruntwork-io/terratest/modules/terraform"
)

func Test_TF(t *testing.T) {
    terraformOptions := &terraform.Options{
        TerraformDir: "../examples",
        Vars: map[string]interface{}{
            "aws_region": region,
        },
    }
    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)
}
```

Randomizer. Add random string to all names, to avoid name conflicts
```hcl
func randSeq(n int) string {
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

name := randSeq(10)
```

First test + getting account id.
```hcl
var terraformDirectory = "../examples"
var region             = "us-east-1"
var account            = ""

func Test_SetUp(t *testing.T) {
    rand.Seed(time.Now().UnixNano()) //IF YOU USE randSeq

    terraformOptions := &terraform.Options{
        TerraformDir: terraformDirectory,
        Vars: map[string]interface{}{
            "aws_region": region,
        },
    }
    defer terraform.Destroy(t, terraformOptions)
    terraform.Init(t, terraformOptions)
    terraform.Apply(t, terraformOptions)
    account = terraform.Output(t, terraformOptions, "account_id")
    fmt.Println("Your account id: " + account)
}
```
