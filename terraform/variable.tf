# variables.tf

variable "aws_access_key" {
    description = "The IAM public access key"
}

variable "aws_secret_key" {
    description = "IAM secret access key"
}

variable "aws_region" {
    description = "The AWS region things are created in"
}

variable "ec2_task_execution_role_name" {
    description = "ECS task execution role name"
    default = "khoury-classroom-backend-task-execution-role"
}

variable "ecs_auto_scale_role_name" {
    description = "ECS auto scale role name"
    default = "khoury-classroom-backend-autoscale-role"
}

variable "az_count" {
    description = "Number of AZs to cover in a given region"
    default = "1"
}

variable "app_image" {
    description = "Docker image to run in the ECS cluster"
    default = "058264409130.dkr.ecr.us-east-2.amazonaws.com/khoury-classroom/backend:latest"
}

variable "app_port" {
    description = "Port exposed by the docker image to redirect traffic to"
    default = 3000
}

variable "app_count" {
    description = "Number of docker containers to run in the ECS cluster"
    default = 1
}

variable "health_check_path" {
    description = "Path to health check endpoint"
    default = "/"
}

variable "fargate_cpu" {
    description = "Fargate instance CPU units to provision (1 vCPU = 1024 CPU units)"
    default = "256"
}

variable "fargate_memory" {
    description = "Fargate instance memory to provision (in MiB)"
    default = "512"
}

variable "scaling_policy_interval" {
    description = "Number of consecutive periods required to trigger the scaling policy"
    default = "2"
}