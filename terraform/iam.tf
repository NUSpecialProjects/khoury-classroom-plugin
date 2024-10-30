data "aws_region" "current" {}

# ------------------------------------------------
#  GitHub Actions Deployment
# ------------------------------------------------

# GitHub Actions Role for Deployment
resource "aws_iam_role" "github_actions_deploy_role" {
  name = "github-actions-deploy-role"

  # Trust policy for GitHub OIDC
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        "Effect" : "Allow",
        "Principal" : {
          "Federated" : "arn:aws:iam::058264409130:oidc-provider/token.actions.githubusercontent.com"
        },
        "Action" : "sts:AssumeRoleWithWebIdentity",
        "Condition" : {
          "StringEquals" : {
            "token.actions.githubusercontent.com:aud" : "sts.amazonaws.com"
          },
          "StringLike" : {
            "token.actions.githubusercontent.com:sub" : "repo:NUSpecialProjects/khoury-classroom-plugin:*"
          }
        }
      }
    ]
  })
}

# Policies to allow GitHub Actions to push Docker images to ECR and update ECS service
resource "aws_iam_policy" "github_actions_deply_policy" {
  name = "github-actions-deploy-policy"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        "Effect" : "Allow",
        "Action" : [
          "iam:GetRole",
          "iam:GetPolicy",
          "iam:GetPolicyVersion",
          "iam:ListAttachedRolePolicies",
          "iam:ListPolicies",
          "iam:ListPolicyVersions",
          "iam:ListRoles",
          "iam:ListRolePolicies",
          "iam:PassRole",

          "acm:*",
          "application-autoscaling:*",
          "cloudwatch:*",
          "cloudfront:*",
          "dynamodb:*",
          "ec2:*",
          "ecr:*",
          "ecs:*",
          "elasticloadbalancing:*",
          "lambda:*",
          "logs:*",
          "rds:*",
          "route53:*",
          "s3:*",
          "secretsmanager:*",

          # "s3:DeleteObject",
          # "s3:GetBucketAcl",
          # "s3:GetBucketPolicy",
          # "s3:GetObject",
          # "s3:ListBucket",
          # "s3:PutObject",

          # "dynamodb:PutItem",
          # "dynamodb:GetItem",
          # "dynamodb:DeleteItem",
          # "dynamodb:DescribeContinuousBackups",
          # "dynamodb:DescribeTable",
          # "dynamodb:UpdateItem",
          # "dynamodb:Scan",
          # "dynamodb:Query",
          # "dynamodb:ListTables",
          
          # "secretsmanager:DescribeSecret",
          # "secretsmanager:GetResourcePolicy",
          # "secretsmanager:GetSecretValue",
          
          # "logs:DescribeLogGroups",
          # "logs:DescribeLogStreams",
          # "logs:GetLogEvents",
          
          # "route53:ListHostedZones",
          # "route53:ListTagsForResource",
          # "route53:GetHostedZone",
          
          # "cloudfront:GetCloudFrontOriginAccessIdentity",
          # "cloudfront:ListDistributions",
          
          # "ec2:DescribeAvailabilityZones",
          # "ec2:DescribeVpcs",
          # "ec2:DescribeSubnets",
          # "ec2:DescribeSecurityGroups",
        ]
        "Resource" : "*"
      },
      {
        "Effect" : "Allow",
        "Action" : [
          "ecr:BatchCheckLayerAvailability",
          "ecr:CompleteLayerUpload",
          "ecr:DescribeRepositories",
          "ecr:InitiateLayerUpload",
          "ecr:PutImage",
          "ecr:UploadLayerPart"
        ],
        "Resource" : "arn:aws:ecr:${data.aws_region.current.name}:${var.aws_account_id}:repository/${var.ecr_repo_name}"
      },
      {
        "Effect" : "Allow",
        "Action" : [
          "ecs:UpdateService",
          "ecs:RegisterTaskDefinition",
          "ecs:DescribeServices",
          "ecs:DescribeTaskDefinition",
          "ecs:DescribeTasks",
          "ecs:ListTasks",
          "ecs:DescribeClusters"
        ],
        "Resource" : [
          "arn:aws:ecs:${data.aws_region.current.name}:${var.aws_account_id}:cluster/${var.ecs_cluster_name}",
          "arn:aws:ecs:${data.aws_region.current.name}:${var.aws_account_id}:service/${var.ecs_cluster_name}/*"
        ]
      },
    ]
  })
}

# Attach the GitHub Actions ECR/ECS policy to the role
resource "aws_iam_role_policy_attachment" "github_actions_deploy_policy_attachment" {
  role       = aws_iam_role.github_actions_deploy_role.name
  policy_arn = aws_iam_policy.github_actions_deply_policy.arn
}