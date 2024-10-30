# database/main.tf

resource "aws_db_instance" "main" {
  allocated_storage      = 20 # Storage in GB
  engine                 = "postgres"
  engine_version         = "16.3"
  instance_class         = "db.t3.micro"
  db_name                = var.db_name
  username               = var.db_user
  password               = var.db_password
  vpc_security_group_ids = [var.rds_sg_id]
  db_subnet_group_name   = var.db_subnet_name
  publicly_accessible    = false
  skip_final_snapshot    = true

  tags = {
    Name = "gitmarks-rds-instance"
  }
}

output "db_vars" {
  description = "Application secrets"
  value = {
    DB_PORT      = var.db_port
    DATABASE_URL = "postgresql://${var.db_user}:${var.db_password}@${aws_db_instance.main.endpoint}/${var.db_name}"
    DB_HOST      = aws_db_instance.main.endpoint
    DB_NAME      = var.db_name
    DB_USER      = var.db_user
    DB_PASSWORD  = var.db_password
  }
}