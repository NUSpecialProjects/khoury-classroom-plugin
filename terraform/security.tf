resource "aws_security_group" "ecs_tasks" {
    name        = "cb-ecs-tasks-security-group"
    description = "allow inbound access directly to ECS tasks"
    vpc_id      = aws_vpc.main.id

    # Allow inbound traffic from the ECS service
    ingress {
        protocol    = "tcp"
        from_port   = var.app_port
        to_port     = var.app_port
        cidr_blocks = ["0.0.0.0/0"]
    }
    
    # Allow ECS tasks to communicate with the RDS instance on the database port
    egress {
      protocol        = "tcp"
      from_port       = 5432
      to_port         = 5432
      cidr_blocks     = [aws_subnet.private[0].cidr_block]
    }

    # Default outbound rule to allow all other traffic
    egress {
      protocol    = "-1"
      from_port   = 0
      to_port     = 0
      cidr_blocks = ["0.0.0.0/0"]
    }
}

# Security Group for RDS
resource "aws_security_group" "rds_sg" {
  name        = "cb-rds-security-group"
  description = "Allow inbound access from ECS tasks"
  vpc_id      = aws_vpc.main.id

  # Allow inbound traffic from the ECS tasks on the database port
  ingress {
    description     = "Allow inbound traffic from ECS tasks to RDS on port 5432"
    protocol        = "tcp"
    from_port       = 5432
    to_port         = 5432
    security_groups = [aws_security_group.ecs_tasks.id]
  }

  # Allow all outbound traffic
  egress {
    protocol    = "-1"
    from_port   = 0
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }
}