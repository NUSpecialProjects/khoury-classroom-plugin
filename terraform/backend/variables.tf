# variables.tf

# ------------------------------------------------
#  Local Variables
# ------------------------------------------------

variable "app_count" {
  description = "Number of docker containers to run in the ECS cluster"
  type        = number
  default     = 1
}

variable "fargate_cpu" {
  description = "Fargate instance CPU units to provision (1 vCPU = 1024 CPU units)"
  type        = number
  default     = 256
}

variable "fargate_memory" {
  description = "Fargate instance memory to provision (in MiB)"
  type        = number
  default     = 512
}

variable "scaling_policy_interval" {
  description = "Number of consecutive periods required to trigger the scaling policy"
  type        = number
  default     = "2"
}

variable "log_retention" {
  description = "Number of days to retain log events"
  type        = number
  default     = "7"
}

# ------------------------------------------------
#  Global Variables
# ------------------------------------------------

variable "aws_account_id" {
  description = "The AWS account ID"
  type        = string
}

variable "private_subnet_ids" {
  description = "List of private subnets"
  type        = list(any)
}

variable "public_subnet_ids" {
  description = "List of public subnets"
  type        = list(any)
}

variable "vpc_id" {
  description = "VPC ID"
  type        = string
}

variable "lb_sg_id" {
  description = "Security group ID for the load balancer"
  type        = string
}

variable "ecs_tasks_sg_id" {
  description = "Security group ID for ECS tasks"
  type        = string
}

variable "db_vars" {
  description = "Map of application secrets"
  type        = map(any)
}

variable "domain_name" {
  description = "The domain name for the React app"
  type        = string
}

variable "app_port" {
  description = "Application port"
  type        = number
}

variable "ecs_cluster_name" {
  description = "ECS cluster name"
  type        = string
}

variable "ecs_service_name" {
  description = "ECS service name"
  type        = string
}

variable "ecr_repo_name" {
  description = "ECR repository name"
  type        = string
}