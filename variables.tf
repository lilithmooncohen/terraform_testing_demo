variable "name" {
  description = "Name to give created resources"
  type        = string
}

variable "description" {
  description = "Description to give created resources"
  type        = string
}

variable "random_suffix" {
  description = "Whether to use a random suffix for the name of created resources"
  type        = bool
  default     = true
}
