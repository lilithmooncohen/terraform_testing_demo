locals {
  name        = var.random_suffix ? "${var.name}-${random_string.suffix.result}" : var.name
  description = var.description
}
