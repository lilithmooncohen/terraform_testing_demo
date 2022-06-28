variable "name" {
  description = "Name to give created resources"
  type        = string
  default     = "example-complete"
}

variable "region" {
  description = "The region where the WebACL will be created"
  type        = string
  default     = "us-east-1"
}


variable "random_suffix" {
  description = "Whether to use a random suffix for the name of created resources"
  type        = bool
  default     = true
}

variable "assume_role_arn" {
  description = "The ARN of the role to assume when making API calls"
  type        = string
  default     = ""
}
