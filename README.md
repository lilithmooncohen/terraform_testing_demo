# Terraform Testing Demo

This repository is a starting point point for a Terraform testing
workshop demo.

# Prep

```bash
git clone git@github.com:lilithmooncohen/terraform_testing_demo.git
cd terraform_testing_demo
asdf install
```

```bash
# Where <MODULE_NAME> is the name of your module, typically in the format github.com/<REPO_NAMESPACE>/<REPO_NAME>.
cd test
go mod init "<MODULE_NAME"
go get github.com/gruntwork-io/terratest/modules/terraform
go test -v -timeout 30m
```
