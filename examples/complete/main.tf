module "example" {
  source = "../../"

  nam           = var.name
  description   = "${var.name} created by TF"
  random_suffix = var.random_suffix
}
