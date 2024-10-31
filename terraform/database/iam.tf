# database/iam.tf

data "aws_region" "current" {}

# Allow basic Lambda execution permissions
resource "aws_iam_role" "lambda_execution_role" {
  name = "gitmarks-drop-db-execution"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Action = "sts:AssumeRole",
      Effect = "Allow",
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })
}

# Allow Lambda interaction with backend components
resource "aws_iam_policy" "lambda_policy" {
  name = "gitmarks-drop-db-policy"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      { # Allow logging
        Effect = "Allow",
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        Resource = "arn:aws:logs:${data.aws_region.current.name}:${var.aws_account_id}:*"
      },
      { # Allow VPC Access
        Effect = "Allow",
        Action = [
          "ec2:CreateNetworkInterface",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DeleteNetworkInterface"
        ],
        Resource = "*"
      },
      { # Allow restarting ECS service
        Effect = "Allow",
        Action = [
          "ecs:UpdateService"
        ],
        Resource = "arn:aws:ecs:${data.aws_region.current.name}:${var.aws_account_id}:service/*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_policy_attachment" {
  role       = aws_iam_role.lambda_execution_role.name
  policy_arn = aws_iam_policy.lambda_policy.arn
}