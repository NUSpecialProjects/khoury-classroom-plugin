# RDS instance configuration
resource "aws_db_instance" "main" {
  allocated_storage    = 20 # Storage in GB
  engine               = "postgres"
  engine_version       = "16.2"
  instance_class       = "db.t3.micro"
  db_name              = var.db_name
  username             = var.db_user
  password             = var.db_password
  vpc_security_group_ids = [aws_security_group.rds_sg.id]
  db_subnet_group_name = aws_db_subnet_group.main.name
  publicly_accessible  = false 
  skip_final_snapshot  = true

  tags = {
    Name = "cb-rds-instance"
  }
}
