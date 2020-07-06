variable "lambda_schedule" {
  default = "cron(*/5 * * * ? *)"
}
variable "region" {}
variable "cw_metric_name" {
  description = "CloudWatch metric name"
}
variable "cw_metric_namespace" {
  description = "CloudWatch metric namespace"
}
variable "target_url" {
  description = "URL which is checked by the lambda function"
}
variable "name" {}
variable "secret_name" {
  default = ""
}

variable "subnet_ids" {}
variable "sg_ids" {}
