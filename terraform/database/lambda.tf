# database/lambda.tf

resource "aws_lambda_function" "drop_db_function" {
  filename         = "./database/func.zip"
  function_name    = "drop_db"
  role             = aws_iam_role.lambda_execution_role.arn
  handler          = "func/drop_db.lambda_handler"
  runtime          = "python3.9"
  source_code_hash = filebase64sha256("./database/func.zip")
  architectures    = ["arm64"]
  layers           = ["arn:aws:lambda:us-east-2:898466741470:layer:psycopg2-py39:1"]
  timeout          = 30

  environment {
    variables = {
      DB_HOST     = split(":", aws_db_instance.main.endpoint)[0] # get rid of the port number
      DB_PORT     = var.db_port
      DB_USERNAME = var.db_user
      DB_PASSWORD = var.db_password
      TARGET_DB   = var.db_name
      ECS_CLUSTER = var.ecs_cluster_name
      ECS_SERVICE = var.ecs_service_name
    }
  }

  vpc_config {
    subnet_ids         = var.private_subnet_ids
    security_group_ids = [var.lambda_sg_id]
  }

  depends_on = [
    aws_iam_role_policy_attachment.lambda_policy_attachment,
  ]
}