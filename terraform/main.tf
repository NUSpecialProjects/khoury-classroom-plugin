# main.tf

module "networking" {
  source   = "./networking"
  app_port = var.app_port
  db_port  = var.db_port
}

module "database" {
  source         = "./database"
  aws_account_id = var.aws_account_id
  db_port        = var.db_port

  db_subnet_name     = module.networking.db_subnet_name
  private_subnet_ids = module.networking.private_subnet_ids

  lambda_sg_id = module.networking.lambda_sg_id
  rds_sg_id    = module.networking.rds_sg_id

  ecs_cluster_name = var.ecs_cluster_name
  ecs_service_name = var.ecs_service_name
}

module "frontend" {
  source         = "./frontend"
  domain_name    = var.domain_name
  aws_account_id = var.aws_account_id

  providers = {
    aws      = aws
    aws.cert = aws.us-east-1 # CloudFront ceritificate requires us-east-1
  }
}

module "backend" {
  source             = "./backend"
  public_subnet_ids  = module.networking.public_subnet_ids
  private_subnet_ids = module.networking.private_subnet_ids
  vpc_id             = module.networking.vpc_id

  lb_sg_id        = module.networking.lb_sg_id
  ecs_tasks_sg_id = module.networking.ecs_tasks_sg_id

  db_vars        = module.database.db_vars
  app_port       = var.app_port
  domain_name    = var.domain_name
  aws_account_id = var.aws_account_id

  ecs_cluster_name = var.ecs_cluster_name
  ecs_service_name = var.ecs_service_name
  ecr_repo_name    = var.ecr_repo_name
}