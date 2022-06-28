package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// An example of how to test the simple Terraform module in examples/basic using Terratest.
func TestTerraformCompleteExample(t *testing.T) {
	t.Parallel()

	expectedPrefix := "test-example-complete"
	expectedSuffix := StringWithCharset(4, "abcdefghijklmnopqrstuvwxyz0123456789")

	expectedName := expectedPrefix + "-" + expectedSuffix

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		// The path to where our Terraform code is located
		TerraformDir: "../examples/complete",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name":          expectedName,
			"random_suffix": false,
		},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,

		// Perfrom upgrade of providers on Terraform init
		Upgrade: true,
	})

	// Clean up resources with "terraform destroy". Using "defer" runs the command at the end of the test, whether the test succeeds or fails.
	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply".
	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of outputs
	// actualOutputName := terraform.Output(t, terraformOptions, "name")
	// actualOutputDescription := terraform.Output(t, terraformOptions, "description")

	// Setup expected values
	// expectedDescription := expectedName + " created by Terraform"

	// Check the output against expected values.
	// Verify we're getting back the outputs we expect
	// assert.Equal(t, expectedName, actualOutputName)
	// assert.Equal(t, expectedDescription, actualOutputDescription)
}
