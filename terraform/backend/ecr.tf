# ecr.tf

resource "aws_ecr_repository" "gitmarks_repo" {
  name                 = var.ecr_repo_name
  image_tag_mutability = "MUTABLE"
}


resource "aws_ecr_lifecycle_policy" "gitmarks_repo_lifecycle" {
  repository = aws_ecr_repository.gitmarks_repo.name

  policy = jsonencode({
    rules = [
      {
        rulePriority = 1
        description  = "Keep only latest image"
        selection = {
          countType   = "imageCountMoreThan"
          countNumber = 1
          tagStatus   = "untagged"
        }
        action = {
          type = "expire"
        }
      }
    ]
  })
}