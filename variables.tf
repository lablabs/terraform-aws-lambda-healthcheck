variable "lambda_schedule" {
  default = "cron(*/5 * * * ? *)"
}
variable "region" {}
variable "cw_metric_name" {}
variable "cw_metric_namespace" {}
variable "target_url" {}
variable "name" {}