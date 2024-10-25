# outputs.tf

output "public_subnet_ids" {
  value = aws_subnet.public.*.id
}

output "private_subnet_ids" {
  value = aws_subnet.private.*.id
}

output "db_subnet_name" {
  value = aws_db_subnet_group.main.name
}

output "vpc_id" {
  value = aws_vpc.main.id
}

output "lb_sg_id" {
  value = aws_security_group.lb.id
}

output "rds_sg_id" {
  value = aws_security_group.rds_sg.id
}

output "ecs_tasks_sg_id" {
  value = aws_security_group.ecs_tasks.id
}

output "lambda_sg_id" {
  value = aws_security_group.lambda_sg.id
}