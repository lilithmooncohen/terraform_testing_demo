module "example" {
  source = "../../"

  name          = var.name
  description   = "${var.name} created by Terraform"
  random_suffix = var.random_suffix
}
