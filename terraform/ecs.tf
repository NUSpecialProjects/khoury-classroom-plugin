# ecs.tf

resource "aws_ecs_cluster" "main" {
    name = "cb-cluster"
}

resource "aws_ecs_task_definition" "app" {
    family                   = "cb-app-task"
    execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
    requires_compatibilities = ["FARGATE"]
    network_mode             = "awsvpc"
    cpu                      = var.fargate_cpu
    memory                   = var.fargate_memory
    container_definitions     = templatefile("./templates/ecs/cb_app.json.tpl", {
        app_image      = var.app_image
        app_port       = var.app_port
        fargate_cpu    = var.fargate_cpu
        fargate_memory = var.fargate_memory
        aws_region     = var.aws_region
    })
}

resource "aws_ecs_service" "main" {
    name            = "cb-service"
    cluster         = aws_ecs_cluster.main.id
    task_definition = aws_ecs_task_definition.app.arn
    desired_count   = var.app_count
    launch_type     = "FARGATE"

    network_configuration {
        security_groups  = [aws_security_group.ecs_tasks.id]
        subnets          = aws_subnet.public.*.id
        assign_public_ip = true
    }

    depends_on = [aws_iam_role_policy_attachment.ecs-task-execution-role-policy-attachment]
}