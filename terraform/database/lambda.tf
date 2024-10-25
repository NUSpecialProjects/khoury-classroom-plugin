# lambda.tf

resource "aws_lambda_function" "drop_db_function" {
  filename         = "./database/drop_db.zip"
  function_name    = "drop_db"
  role             = aws_iam_role.lambda_execution_role.arn
  handler          = "drop_db.lambda_handler"
  runtime          = "python3.9"
  source_code_hash = filebase64sha256("./database/drop_db.zip")

  environment {
    variables = {
      DB_HOST     = aws_db_instance.main.endpoint
      DB_PORT     = var.db_port
      DB_USERNAME = var.db_user
      DB_PASSWORD = var.db_password
      TARGET_DB   = var.db_name
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