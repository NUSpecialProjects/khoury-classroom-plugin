# logs.tf

# Set up CloudWatch group and log stream and retain logs for 30 days
resource "aws_cloudwatch_log_group" "gitmarks_log_group" {
  name              = "/ecs/gitmarks-app"
  retention_in_days = 7

  tags = {
    Name = "gitmarks-log-group"
  }
}

resource "aws_cloudwatch_log_stream" "gitmarks_log_stream" {
  name           = "gitmarks-log-stream"
  log_group_name = aws_cloudwatch_log_group.gitmarks_log_group.name

  depends_on = [ aws_cloudwatch_log_group.gitmarks_log_group ]
}

# S3 Bucket for Access Logs
resource "aws_s3_bucket" "logs_prod" {
  bucket = "logs-prod-${var.aws_account_id}"
  force_destroy = true
}
resource "aws_s3_bucket_policy" "logs_prod_policy" {
  bucket = aws_s3_bucket.logs_prod.id

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::033677994240:root"
      },
      "Action": "s3:PutObject",
      "Resource": "arn:aws:s3:::logs-prod-${var.aws_account_id}/alb/alb-prod/AWSLogs/${var.aws_account_id}/*"
    }
  ]
}
POLICY

  depends_on = [aws_s3_bucket.logs_prod]
}