# ------------------------------------------------
# IAM Roles for ECS
# ------------------------------------------------

resource "aws_iam_role" "ecs_task_execution_role" {
  name = "role-name"
 
  assume_role_policy = <<EOF
{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Action": "sts:AssumeRole",
     "Principal": {
       "Service": "ecs-tasks.amazonaws.com"
     },
     "Effect": "Allow",
     "Sid": ""
   }
 ]
}
EOF
}
resource "aws_iam_role" "ecs_task_role" {
  name = "role-name-task"
 
  assume_role_policy = <<EOF
{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Action": "sts:AssumeRole",
     "Principal": {
       "Service": "ecs-tasks.amazonaws.com"
     },
     "Effect": "Allow",
     "Sid": ""
   }
 ]
}
EOF
}

# Attach the Amazon ECS task execution role policy to the role
resource "aws_iam_role_policy_attachment" "ecs-task-execution-role-policy-attachment" {
  role       = aws_iam_role.ecs_task_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}
resource "aws_iam_role_policy_attachment" "task_s3" {
  role       = "${aws_iam_role.ecs_task_role.name}"
  policy_arn = "arn:aws:iam::aws:policy/AmazonS3FullAccess"
}

# ------------------------------------------------
# ECS Auto Scale Role
# ------------------------------------------------

data "aws_iam_policy_document" "ecs_auto_scale_role" {
  version = "2012-10-17"
  statement {
    effect = "Allow"
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["application-autoscaling.amazonaws.com"]
    }
  }
}

# Attach the Amazon ECS auto scale role policy to the role
resource "aws_iam_role" "ecs_auto_scale_role" {
  name               = var.ecs_auto_scale_role_name
  assume_role_policy = data.aws_iam_policy_document.ecs_auto_scale_role.json
}
resource "aws_iam_role_policy_attachment" "ecs_auto_scale_role" {
  role       = aws_iam_role.ecs_auto_scale_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceAutoscaleRole"
}


# ------------------------------------------------
# IAM Role for GitHub Actions Deployment
# ------------------------------------------------

# GitHub Actions Role for Deployment
resource "aws_iam_role" "github_actions_deploy_role" {
  name = "github-actions-deploy-role"

  # Trust policy for GitHub OIDC
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Federated": "arn:aws:iam::058264409130:oidc-provider/token.actions.githubusercontent.com"
      },
      "Action": "sts:AssumeRoleWithWebIdentity",
      "Condition": {
          "StringEquals": {
              "token.actions.githubusercontent.com:aud": "sts.amazonaws.com"
          },
          "StringLike": {
              "token.actions.githubusercontent.com:sub": "repo:NUSpecialProjects/khoury-classroom-plugin:*"
          }
      }
    }
  ]
}
EOF
}


# Policies to allow GitHub Actions to push Docker images to ECR and update ECS service
resource "aws_iam_policy" "github_actions_ecr_ecs_policy" {
  name = "github-actions-ecr-ecs-policy"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "ecr:GetAuthorizationToken",
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": [
                "ecr:BatchCheckLayerAvailability",
                "ecr:PutImage",
                "ecr:InitiateLayerUpload",
                "ecr:UploadLayerPart",
                "ecr:CompleteLayerUpload",
                "ecr:DescribeRepositories"
            ],
            "Resource": "arn:aws:ecr:${var.aws_region}:${var.aws_account_id}:repository/khoury-classroom/backend"
        },
        {
            "Effect": "Allow",
            "Action": [
                "ecs:UpdateService",
                "ecs:RegisterTaskDefinition",
                "ecs:DescribeServices",
                "ecs:DescribeTaskDefinition",
                "ecs:DescribeTasks"
            ],
            "Resource": [
              "arn:aws:ecs:${var.aws_region}:${var.aws_account_id}:cluster/cb-cluster",
              "arn:aws:ecs:${var.aws_region}:${var.aws_account_id}:service/cb-cluster/*"
            ]
        }
    ]
}
EOF
}

# Attach the GitHub Actions ECR/ECS policy to the role
resource "aws_iam_role_policy_attachment" "github_actions_deploy_policy_attachment" {
  role       = aws_iam_role.github_actions_deploy_role.name
  policy_arn = aws_iam_policy.github_actions_ecr_ecs_policy.arn
}