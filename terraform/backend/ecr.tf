# ecr.tf

resource "aws_ecr_repository" "gitmarks_repo" {
  name                 = "khoury-classroom/backend"
  image_tag_mutability = "MUTABLE"
}

resource "aws_ecr_lifecycle_policy" "gitmarks_repo_lifecycle" {
  repository = aws_ecr_repository.gitmarks_repo.name

  policy = jsonencode({
    rules = [
      {
        rulePriority = 1
        description   = "Keep only latest image"
        selection     = {
          countType        = "imageCountMoreThan"
          countNumber      = 1
          tagStatus        = "tagged"
          tagPrefixList    = ["latest"]
        }
        action = {
          type = "expire"
        }
      },
      {
        rulePriority = 2,
        description  = "Expire images older than 7 days",
        selection    = {
          tagStatus = "untagged",
          countType = "sinceImagePushed",
          countUnit = "days",
          countNumber = 7
        },
        action = {
          type = "expire"
        }
      }
    ]
  })
}

output "ecr_repo_url" {
  value = aws_ecr_repository.gitmarks_repo.repository_url
}