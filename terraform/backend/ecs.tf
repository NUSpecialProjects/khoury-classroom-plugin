# ecs.tf

resource "aws_ecs_cluster" "main" {
  name = var.ecs_cluster_name
}

# Fetch environment variables from AWS Secrets Manager
data "aws_secretsmanager_secret" "env_variables" {
  name = "prod/gitmarks/app_secrets"
}
data "aws_secretsmanager_secret_version" "env_variables_version" {
  secret_id = data.aws_secretsmanager_secret.env_variables.id
}
locals {
  app_secrets = jsondecode(data.aws_secretsmanager_secret_version.env_variables_version.secret_string)
}

# Create the ECS task definition
data "aws_region" "current" {}
resource "aws_ecs_task_definition" "app" {
  family                   = "gitmarks-app-task"
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = var.fargate_cpu
  memory                   = var.fargate_memory
  container_definitions = jsonencode([{
    name   = "gitmarks-app"
    image  = aws_ecr_repository.gitmarks_repo.repository_url
    cpu    = var.fargate_cpu
    memory = var.fargate_memory

    logConfiguration = {
      logDriver = "awslogs"
      options = {
        "awslogs-group"         = "/ecs/gitmarks-app"
        "awslogs-region"        = data.aws_region.current.name
        "awslogs-stream-prefix" = "ecs"
      }
    }

    portMappings = [{
      containerPort = var.app_port
      hostPort      = var.app_port
    }]

    environment = [
      for key, value in var.db_vars : {
        name  = key
        value = value
      }
    ]
    secrets = [
      for key, value in local.app_secrets : {
        name      = key
        valueFrom = "${data.aws_secretsmanager_secret.env_variables.arn}:${key}::"
      }
    ]
  }])
}

# Create the ECS service
resource "aws_ecs_service" "main" {
  name            = var.ecs_service_name
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.app.arn
  desired_count   = var.app_count
  launch_type     = "FARGATE"

  network_configuration {
    security_groups = [var.ecs_tasks_sg_id]
    subnets         = var.private_subnet_ids
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.app.id
    container_name   = "gitmarks-app"
    container_port   = var.app_port
  }

  depends_on = [
    aws_lb_listener.front_end_http,
    aws_lb_listener.front_end_https,
    aws_iam_role_policy_attachment.ecs_task_execution_role_policy_attachment
  ]
}