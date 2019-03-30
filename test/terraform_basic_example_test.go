package test

import (
	"testing"
	// "fmt"
	// "time"
    // "math/rand"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var terraformDirectory = "../examples"
var region             = "us-east-1"
// var account            = ""

// func randSeq(n int) string {
// 	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
//     b := make([]rune, n)
//     for i := range b {
//         b[i] = letters[rand.Intn(len(letters))]
//     }
//     return string(b)
// }

// func Test_SetUp(t *testing.T) {
// 	rand.Seed(time.Now().UnixNano()) //IF YOU USE randSeq

// 	terraformOptions := &terraform.Options{
// 		TerraformDir: terraformDirectory,
// 		Vars: map[string]interface{}{
// 			"aws_region": region,
// 			"example": randSeq(10), //EXAMPLE FOR THIS MODULE! NOT NECESSARY
// 		},
// 	}
// 	defer terraform.Destroy(t, terraformOptions)
// 	terraform.Init(t, terraformOptions)
// 	terraform.Apply(t, terraformOptions)
// 	account = terraform.Output(t, terraformOptions, "account_id")
// 	fmt.Println("Your account id: " + account)
// }

// An example of how to test the simple Terraform module in examples/terraform-basic-example using Terratest.
func TestTerraformBasicExample(t *testing.T) {
	t.Parallel()

	expectedText := "test"

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: terraformDirectory,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"aws_region": region,
			"example": expectedText,
		},

		// Variables to pass to our Terraform code using -var-file options
		VarFiles: []string{"../varfile.tfvars"},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	actualTextExample := terraform.Output(t, terraformOptions, "example")
	actualTextExample2 := terraform.Output(t, terraformOptions, "example2")

	// Verify we're getting back the outputs we expect
	assert.Equal(t, expectedText, actualTextExample)
	assert.Equal(t, expectedText, actualTextExample2)
}

