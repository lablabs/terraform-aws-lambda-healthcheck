locals {
  name = "${var.name}-healthcheck"
}

resource "aws_iam_role" "this" {
  name = local.name

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "this" {
  name = local.name
  path = "/"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:*",
      "Effect": "Allow"
     },
     {
      "Action": [
        "cloudwatch:PutMetricData",
        "ec2:CreateNetworkInterface",
        "ec2:DescribeNetworkInterfaces",
        "ec2:DeleteNetworkInterface",
        "kms:GenerateDataKey",
        "kms:Decrypt"
      ],
      "Resource": "*",
      "Effect": "Allow"
     },
     {
      "Action": [
        "secretsmanager:GetSecretValue"
      ],
      "Resource": "arn:aws:secretsmanager:${var.region}:*:secret:${var.secret_name}*",
      "Effect": "Allow"
     }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "this" {
  role       = aws_iam_role.this.name
  policy_arn = aws_iam_policy.this.arn
}

resource "aws_cloudwatch_log_group" "this" {
  name              = "/aws/lambda/${local.name}"
  retention_in_days = 14
}

resource "aws_lambda_function" "this" {
  filename      = "${path.module}/lambda.zip"
  function_name = local.name
  role          = aws_iam_role.this.arn
  handler       = "main"

  timeout = "60"

  source_code_hash = filebase64sha256("${path.module}/lambda.zip")

  runtime = "go1.x"

  vpc_config {
    subnet_ids         = var.subnet_ids
    security_group_ids = var.sg_ids
  }

  environment {
    variables = {
      REGION              = var.region
      CW_METRIC_NAME      = var.cw_metric_name
      CW_METRIC_NAMESPACE = var.cw_metric_namespace
      TARGET_URL          = var.target_url
      SECRET_NAME         = var.secret_name
    }
  }
}

resource "aws_cloudwatch_event_rule" "trigger" {
  name        = local.name
  description = "Trigger creation of RDS snapshot on schedule"

  schedule_expression = var.lambda_schedule
}

resource "aws_cloudwatch_event_target" "this" {
  rule = aws_cloudwatch_event_rule.trigger.name
  arn  = aws_lambda_function.this.arn
}

resource "aws_lambda_permission" "this" {
  statement_id  = "${aws_lambda_function.this.function_name}-AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.this.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.trigger.arn
}

resource "aws_cloudwatch_metric_alarm" "default" {
  provider                  = aws.alarm
  alarm_name                = "${var.target_url} not healthy"
  comparison_operator       = "GreaterThanThreshold"
  evaluation_periods        = "2"
  metric_name               = var.cw_metric_name
  namespace                 = var.cw_metric_namespace
  period                    = "300"
  statistic                 = "Maximum"
  threshold                 = "200"
  alarm_description         = "Monitors health of a specified endpoint"
  insufficient_data_actions = []
  treat_missing_data        = "breaching"
}
