variable "lambda_schedule" {
  type    = string
  default = "cron(*/5 * * * ? *)"
}
variable "region" {
  type = string
}
variable "cw_metric_name" {
  type        = string
  description = "CloudWatch metric name"
}
variable "cw_metric_namespace" {
  type        = string
  description = "CloudWatch metric namespace"
}
variable "target_url" {
  type        = string
  description = "URL which is checked by the lambda function"
}
variable "name" {
  type = string
}
variable "secret_name" {
  type    = string
  default = ""
}

variable "subnet_ids" {
  type = list(string)
}
variable "sg_ids" {
  type = list(string)
}
