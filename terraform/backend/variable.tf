# variables.tf

variable "aws_account_id" {
  description = "The AWS account ID"
  type = string
}

variable "ec2_task_execution_role_name" {
  description = "ECS task execution role name"
  default     = "khoury-classroom-backend-task-execution-role"
}

variable "ecs_auto_scale_role_name" {
  description = "ECS auto scale role name"
  default     = "khoury-classroom-backend-autoscale-role"
}

variable "app_count" {
  description = "Number of docker containers to run in the ECS cluster"
  default     = 1
}

variable "health_check_path" {
  description = "Path to health check endpoint"
  default     = "/"
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
  default     = "2"
}

variable "backend_subdomain" {
  description = "Subdomain prefix for the backend service"
  type        = string
  default     = "api"
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

variable "app_secrets" {
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