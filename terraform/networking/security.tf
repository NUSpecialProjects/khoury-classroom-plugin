# security.tf

resource "aws_security_group" "lb" {
  name        = "cb-load-balancer-security-group"
  description = "controls access to the ALB"
  vpc_id      = aws_vpc.main.id

  ingress {
    protocol    = "tcp"
    from_port   = 443
    to_port     = 443
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    protocol    = "-1"
    from_port   = 0
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "ecs_tasks" {
  name        = "cb-ecs-tasks-security-group"
  description = "allow inbound access directly to ECS tasks"
  vpc_id      = aws_vpc.main.id

  # Allow inbound traffic from the load balancer
  ingress {
    protocol        = "tcp"
    from_port       = var.app_port
    to_port         = var.app_port
    security_groups = [aws_security_group.lb.id]
  }

  # Default outbound rule to allow all other traffic
  egress {
    protocol    = "-1"
    from_port   = 0
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "rds_sg" {
  name        = "cb-rds-security-group"
  description = "Allow inbound access from ECS tasks"
  vpc_id      = aws_vpc.main.id

  # Allow inbound traffic from the ECS tasks and Lambda functions
  ingress {
    description     = "Allow inbound traffic from ECS tasks to RDS"
    protocol        = "tcp"
    from_port       = var.db_port
    to_port         = var.db_port
    security_groups = [aws_security_group.ecs_tasks.id, aws_security_group.lambda_sg.id]
  }

  # Allow all outbound traffic
  egress {
    protocol    = "-1"
    from_port   = 0
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "lambda_sg" {
  name        = "drop_db_sg"
  description = "Security group for Lambda to access RDS"
  vpc_id      = aws_vpc.main.id

  egress {
    from_port   = var.db_port
    to_port     = var.db_port
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = var.app_port
    to_port     = var.app_port
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  # Default outbound rule to allow all other traffic
  egress {
    protocol    = "-1"
    from_port   = 0
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }
}