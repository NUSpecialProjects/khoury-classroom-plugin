# main.tf

module "networking" {
  source   = "./networking"
  app_port = var.app_port
  db_port  = var.db_port
}

module "database" {
  source         = "./database"
  db_port        = var.db_port
  rds_sg_id      = module.networking.rds_sg_id
  db_subnet_name = module.networking.db_subnet_name
}

module "frontend" {
  source      = "./frontend"
  domain_name = var.domain_name

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

  app_secrets = module.database.app_secrets
  app_port    = var.app_port
  domain_name = var.domain_name
}
