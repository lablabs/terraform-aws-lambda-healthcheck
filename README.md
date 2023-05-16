# AWS Lambda health check Terraform module

[![Labyrinth Labs logo](ll-logo.png)](https://www.lablabs.io)

We help companies build, run, deploy and scale software and infrastructure by embracing the right technologies and principles. Check out our website at https://lablabs.io/

---

![Terraform validation](https://github.com/lablabs/terraform-aws-lambda-healthcheck/workflows/Terraform%20validation/badge.svg?branch=master)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-success?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)

## Description

A terraform module to deploy a health check lambda function and to provide AWS CloudWatch Metric Alarm resource.

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 4.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_cloudwatch_event_rule.trigger](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_event_rule) | resource |
| [aws_cloudwatch_event_target.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_event_target) | resource |
| [aws_cloudwatch_log_group.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_log_group) | resource |
| [aws_cloudwatch_metric_alarm.default](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_metric_alarm) | resource |
| [aws_iam_policy.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_policy) | resource |
| [aws_iam_role.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role) | resource |
| [aws_iam_role_policy_attachment.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy_attachment) | resource |
| [aws_lambda_function.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_function) | resource |
| [aws_lambda_permission.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_permission) | resource |
| [aws_sns_topic.sns](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sns_topic) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_cw_metric_name"></a> [cw\_metric\_name](#input\_cw\_metric\_name) | CloudWatch metric name | `string` | n/a | yes |
| <a name="input_cw_metric_namespace"></a> [cw\_metric\_namespace](#input\_cw\_metric\_namespace) | CloudWatch metric namespace | `string` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | n/a | `string` | n/a | yes |
| <a name="input_region"></a> [region](#input\_region) | n/a | `string` | n/a | yes |
| <a name="input_sg_ids"></a> [sg\_ids](#input\_sg\_ids) | n/a | `list(string)` | n/a | yes |
| <a name="input_subnet_ids"></a> [subnet\_ids](#input\_subnet\_ids) | n/a | `list(string)` | n/a | yes |
| <a name="input_target_url"></a> [target\_url](#input\_target\_url) | URL which is checked by the lambda function | `string` | n/a | yes |
| <a name="input_lambda_schedule"></a> [lambda\_schedule](#input\_lambda\_schedule) | n/a | `string` | `"cron(*/5 * * * ? *)"` | no |
| <a name="input_secret_name"></a> [secret\_name](#input\_secret\_name) | n/a | `string` | `""` | no |

## Outputs

No outputs.
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## Contributing and reporting issues

Feel free to create an issue in this repository if you have questions, suggestions or feature requests.

### Validation, linters and pull-requests

We want to provide high quality code and modules. For this reason we are using
several [pre-commit hooks](.pre-commit-config.yaml) and
[GitHub Actions workflow](.github/workflows/main.yml). A pull-request to the
master branch will trigger these validations and lints automatically. Please
check your code before you will create pull-requests. See
[pre-commit documentation](https://pre-commit.com/) and
[GitHub Actions documentation](https://docs.github.com/en/actions) for further
details.


## License

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

See [LICENSE](LICENSE) for full details.

    Licensed to the Apache Software Foundation (ASF) under one
    or more contributor license agreements.  See the NOTICE file
    distributed with this work for additional information
    regarding copyright ownership.  The ASF licenses this file
    to you under the Apache License, Version 2.0 (the
    "License"); you may not use this file except in compliance
    with the License.  You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing,
    software distributed under the License is distributed on an
    "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
    KIND, either express or implied.  See the License for the
    specific language governing permissions and limitations
    under the License.
