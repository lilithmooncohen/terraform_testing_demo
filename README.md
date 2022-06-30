# Terraform Testing Demo

This repository is a starting point point for a Terraform testing
workshop demo.

## Prep

Clone the repo

```bash
git clone git@github.com:lilithmooncohen/terraform_testing_demo.git
cd terraform_testing_demo
```

Add the needed tools via asdf:

```bash
asdf plugin-add terraform
asdf plugin-add golang
asdf
```

Initialize the `test` go module:

```bash
# Where <MODULE_NAME> is the name of your module, typically in the format github.com/<REPO_NAMESPACE>/<REPO_NAME>.
cd test
go mod init "<MODULE_NAME"
go get github.com/gruntwork-io/terratest/modules/terraform
```

## Working with Terratest

I have already created a [complete_test.go](test/complete_test.go) file that closely mirrors the test file found
in the [Terratest Quick Start](https://terratest.gruntwork.io/docs/getting-started/quick-start/)
and a [helpers.go](test/helpers.go) file for housing helper functions to get us up and running quickly.

Now we have everything we need in place to get running.
We aren't yet validating any outputs,
but Terratest can now validate that it can at least run a `terraform apply` succesfully.
Let's run a test to see if a `terraform apply` is working.

```bash
go test -timeout 30m
```

Output:

```bash
An argument named "nam" is not expected here. Did you mean "name"?}
--- FAIL: TestTerraformCompleteExample (4.80s)
    apply.go:15:
                Error Trace:    apply.go:15
                                                        complete_test.go:42
                Error:          Received unexpected error:
                                FatalError{Underlying: error while running command: exit status 1;
                                Error: Missing required argument

                                  on main.tf line 1, in module "example":
                                   1: module "example" {

                                The argument "name" is required, but no definition was found.

                                Error: Unsupported argument

                                  on main.tf line 4, in module "example":
                                   4:   nam           = var.name

                                An argument named "nam" is not expected here. Did you mean "name"?}
                Test:           TestTerraformCompleteExample
    destroy.go:11:
                Error Trace:    destroy.go:11
                                                        panic.go:482
                                                        testing.go:864
                                                        apply.go:15
                                                        complete_test.go:42
                Error:          Received unexpected error:
                                FatalError{Underlying: error while running command: exit status 1;
                                Error: Missing required argument

                                  on main.tf line 1, in module "example":
                                   1: module "example" {

                                The argument "name" is required, but no definition was found.

                                Error: Unsupported argument

                                  on main.tf line 4, in module "example":
                                   4:   nam           = var.name

                                An argument named "nam" is not expected here. Did you mean "name"?}
                Test:           TestTerraformCompleteExample
FAIL
exit status 1
FAIL    github.com/lilithmooncohen/terraform_testing_demo       5.284s
```

As you can see it looks like we have a problem in line 4 of [examples/complete/main.tf](examples/complete/main.tf).
Let's fix that.

```diff
-  nam           = var.name
+  name          = var.name
```

Now that it's fixed, let's run another test:

```bash
go test -timeout 30m
```

Output:

```bash
PASS
ok      github.com/lilithmooncohen/terraform_testing_demo       6.380s
```

It passed! Now lets validate some outputs!

```diff
--- a/test/complete_test.go
+++ b/test/complete_test.go
@@ -4,6 +4,7 @@ import (
        "testing"

        "github.com/gruntwork-io/terratest/modules/terraform"
+       "github.com/stretchr/testify/assert"
 )

 // An example of how to test the simple Terraform module in examples/basic using Terratest.
@@ -42,14 +43,14 @@ func TestTerraformCompleteExample(t *testing.T) {
        terraform.InitAndApply(t, terraformOptions)

        // Run `terraform output` to get the values of outputs
-       // actualOutputName := terraform.Output(t, terraformOptions, "name")
-       // actualOutputDescription := terraform.Output(t, terraformOptions, "description")
+       actualOutputName := terraform.Output(t, terraformOptions, "name")
+       actualOutputDescription := terraform.Output(t, terraformOptions, "description")

-       // Setup expected values
-       // expectedDescription := expectedName + " created by Terraform"
+       // Setup additional expected values
+       expectedDescription := expectedName + " created by Terraform"

        // Check the output against expected values.
        // Verify we're getting back the outputs we expect
-       // assert.Equal(t, expectedName, actualOutputName)
-       // assert.Equal(t, expectedDescription, actualOutputDescription)
+       assert.Equal(t, expectedName, actualOutputName)
+       assert.Equal(t, expectedDescription, actualOutputDescription)
```

Now let's run a test to see if our output assertions are working:
```bash
go test -timeout 30m
```

Output:

```bash
--- FAIL: TestTerraformCompleteExample (7.15s)
    complete_test.go:55:
                Error Trace:    complete_test.go:55
                Error:          Not equal:
                                expected: "test-example-complete-gbe3 created by Terraform"
                                actual  : "test-example-complete-gbe3 created by TF"

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1 +1 @@
                                -test-example-complete-gbe3 created by Terraform
                                +test-example-complete-gbe3 created by TF
                Test:           TestTerraformCompleteExample
FAIL
exit status 1
FAIL    github.com/lilithmooncohen/terraform_testing_demo       7.518s
```

Looks like our description in line 5 of [examples/complete/main.tf](examples/complete/main.tf) doesn't match what we expect in our tests.
Let's fix that.

```diff
-  description   = "${var.name} created by TF"
+  description   = "${var.name} created by Terraform"
```

Now that it's fixed, let's run another test:

```bash
go test -timeout 30m
```

Output:

```bash
PASS
ok      github.com/lilithmooncohen/terraform_testing_demo       13.540s
```
