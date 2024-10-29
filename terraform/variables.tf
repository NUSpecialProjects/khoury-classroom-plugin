# variables.tf

variable "domain_name" {
  description = "The domain name for the React app"
  type        = string
  default     = "gitmarks.org"
}

variable "app_port" {
  description = "Application port"
  type        = number
  default     = 8080
}

variable "db_port" {
  description = "Database port"
  type        = number
  default     = 5432
}

variable "aws_account_id" {
  description = "AWS account ID"
  type        = string
  default     = "058264409130"
}

variable "ecs_cluster_name" {
  description = "ECS cluster name"
  type        = string
  default     = "gitmarks-cluster"
}

variable "ecs_service_name" {
  description = "ECS service name"
  type        = string
  default     = "gitmarks-service"
}

variable "ecr_repo_name" {
  description = "ECR repository name"
  type        = string
  default     = "khoury-classroom/backend"
}