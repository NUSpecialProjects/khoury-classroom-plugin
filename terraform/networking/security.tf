# security.tf

resource "aws_security_group" "lb" {
  name        = "gitmarks-load-balancer-security-group"
  description = "controls access to the ALB"
  vpc_id      = aws_vpc.main.id

  ingress {
    protocol    = "tcp"
    from_port   = 443
    to_port     = 443
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    protocol    = "tcp"
    from_port   = 80
    to_port     = 80
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    protocol    = "tcp"
    from_port   = var.app_port
    to_port     = var.app_port
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "ecs_tasks" {
  name        = "gitmarks-ecs-tasks-security-group"
  description = "allow inbound access directly to ECS tasks"
  vpc_id      = aws_vpc.main.id

  # Allow inbound traffic from the load balancer
  ingress {
    protocol        = "tcp"
    from_port       = var.app_port
    to_port         = var.app_port
    security_groups = [aws_security_group.lb.id]
  }

  # Allow ECS tasks to communicate with the RDS instance on the database port
  egress {
    protocol    = "tcp"
    from_port   = var.db_port
    to_port     = var.db_port
    cidr_blocks = [aws_subnet.private[0].cidr_block]
  }
}

resource "aws_security_group" "rds_sg" {
  name        = "gitmarks-rds-security-group"
  description = "Allow inbound access from ECS tasks"
  vpc_id      = aws_vpc.main.id

  # Allow inbound traffic from the ECS tasks on the database port
  ingress {
    description     = "Allow inbound traffic from ECS tasks to RDS"
    protocol        = "tcp"
    from_port       = var.db_port
    to_port         = var.db_port
    security_groups = [aws_security_group.ecs_tasks.id]
  }

  # Allow outbound traffic to the ECS tasks
  egress {
    protocol    = "tcp"
    from_port   = var.app_port
    to_port     = var.app_port
    cidr_blocks = ["0.0.0.0/0"]
  }
}