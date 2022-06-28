# example complete example

This example illustrates how to use the example module.

<!-- BEGIN_TF_DOCS -->

<!-- prettier-ignore-start -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 3.63.0 |
| <a name="requirement_random"></a> [random](#requirement\_random) | 3.1.3 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_example"></a> [example](#module\_example) | ../../ | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_assume_role_arn"></a> [assume\_role\_arn](#input\_assume\_role\_arn) | The ARN of the role to assume when making API calls | `string` | `""` | no |
| <a name="input_name"></a> [name](#input\_name) | Name to give created resources | `string` | `"example-complete"` | no |
| <a name="input_random_suffix"></a> [random\_suffix](#input\_random\_suffix) | Whether to use a random suffix for the name of created resources | `bool` | `true` | no |
| <a name="input_region"></a> [region](#input\_region) | The region where the WebACL will be created | `string` | `"us-east-1"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_description"></a> [description](#output\_description) | Friendly description of the created WebACL |
| <a name="output_name"></a> [name](#output\_name) | Friendly name of the created WebACL |
<!-- prettier-ignore-end -->

<!-- END_TF_DOCS -->
