resource "aws_security_group" "ecs_tasks" {
    name        = "cb-ecs-tasks-security-group"
    description = "allow inbound access directly to ECS tasks"
    vpc_id      = aws_vpc.main.id

    # Allow inbound traffic from anywhere to the application port
    ingress {
        protocol    = "tcp"
        from_port   = var.app_port
        to_port     = var.app_port
        cidr_blocks = ["0.0.0.0/0"]
    }

    # Allow all outbound traffic (standard in ECS)
    egress {
        protocol    = "-1"
        from_port   = 0
        to_port     = 0
        cidr_blocks = ["0.0.0.0/0"]
    }
}
