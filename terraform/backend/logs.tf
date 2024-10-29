# logs.tf

# Store logs
resource "aws_cloudwatch_log_group" "gitmarks_log_group" {
  name              = "/ecs/gitmarks-app"
  retention_in_days = var.log_retention

  tags = {
    Name = "gitmarks-log-group"
  }
}
resource "aws_cloudwatch_log_stream" "gitmarks_log_stream" {
  name           = "gitmarks-log-stream"
  log_group_name = aws_cloudwatch_log_group.gitmarks_log_group.name

  depends_on = [aws_cloudwatch_log_group.gitmarks_log_group]
}