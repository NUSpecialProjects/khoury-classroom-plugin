# ecs.tf

resource "aws_ecs_cluster" "main" {
  name = "gitmarks-cluster"
}

# Create the ECS task definition
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
      for key, value in var.app_secrets : {
        name  = key
        value = value
      }
    ]
  }])
}

resource "aws_ecs_service" "main" {
  name            = "gitmarks-service"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.app.arn
  desired_count   = var.app_count
  launch_type     = "FARGATE"

  network_configuration {
    security_groups  = [var.ecs_tasks_sg_id]
    subnets          = var.private_subnet_ids
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.app.id
    container_name   = "gitmarks-app"
    container_port   = var.app_port
  }

  depends_on = [
    aws_lb_listener.front_end_http,
    aws_lb_listener.front_end_https,
    aws_iam_role_policy_attachment.ecs-task-execution-role-policy-attachment
  ]
}