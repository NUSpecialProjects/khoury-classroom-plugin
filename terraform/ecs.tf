# ecs.tf

resource "aws_ecs_cluster" "main" {
    name = "cb-cluster"
}

# Define the environment variables
locals {
    app_secrets = {
        APP_PRIVATE_KEY        = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAoHQT+RKgjEdkq7SBiQ0LkKRaWyeh8AD8zHBovlVa6yV/ik2C\nieDwIWGtlx2OU3gqf13ZZDEGvNwi1PDt3m+ny3Xsy901vyvlGSyqfmAn0oTj1rJu\nnOq9MfcXWfMjbPXyaONRj0/1aHNN2Bvl5Ye/dgCOVwYZOgsu8Ka33aH5dRWfnPf7\ng4XQoO1oBoe4EufnPJ2uciTTZDuvg1rnav6e+uMEKopNpKFvtg4GAQl4x0Qq6xkQ\n6VfA8opnit7mOaWcDdXc3thgN7CHTPCD4XqvcmFsBJNclGEfjCbdavrIIp4JKkpw\nDmk1Ou6Lz78qnbTxQ9CYeBJDfvarLPDJyKGsDQIDAQABAoIBAQCgRlMNIYYtmcL9\noTkjZVyABywakeQ4kUP0EvUN6sT+zl4wEGysvXwgXCnCIUviJM6Om3hjlHVegaZp\nfqCc6Ht7yTfYDAd8BqS6GNvVkMc2infsJiBHrlN+bYtt1mk0lhimnSsDNKO2yjag\nAH4MYSTnAncshnL8f99Lk71mLj24rWOa7SjmRKga6pgNacHZ8cCBR0TKmMJNlEFL\noHPZveTmg+uu7obaGJ7Lo08hQJsocB7nrqUtrnpcPGDz+yXa+nRwYCTza4+4AKvN\nkX0680js2adth8ELpL5Gy+9ukp8LMf7gArqSd+ZAGhSgSIaklLfixGHcWBgry2gx\nOepXnmMhAoGBANNFzqSYNH+MlxbW8n1Wj/cp396R1SUKbXP6oFPMUS0mMtOU7RFS\nqSoNIgT5bhJk5B8/D+f4XEZNWWMBnq5EgL39hmstULkl26htKPKHmDvX5X3hkXRH\nyAKTeYWEblE+PpDqKFE9O3IpPovyJ0Oe19U+100i/Gry6bgRDzHG3DAHAoGBAMJs\nDnOH789IQjiQprhj+3ue9Xl2JEU959PbsCk2jmp2NMuuYf3fyK9iZhCXbBysmaSH\nOUGbl86SFOdwtTcGwYKzyNaCUdzNYTqLUoWiJkLMn3PZbggiJ2MQ2e42UNiCAsQb\ndjPh5/66dyMlJuBH3kfh0yP5ShK4heEiMULYSRZLAoGBAJ/TTVICuqRLDPmAPg1H\ncL1/9hV/qQjObKKyVJtQE5DeNtEM9pKGP+bJ7JRqxTQxEsn4gOXxYozkctyNGyem\nNuaDZi6qJ0kJNLSjb7iZjzamSrwB6nFW5B3exq2U04euWNJz8XATrGbegKyJ0d47\nyfdOBL4b22xkux49+Yqkb2n9AoGAb45Y7Gl/bExl0tcNEpgr4E7hQwRK44AV2TYg\n6kTniqawvH4es/EH0bqAHd0Ep59RuVntvHtuq5Secf31vNEfj8Ng5dR47FzcAR+Y\nBh14HrQSegK0Y+5U8z7kDQ8VbGWM+MFZHYPt/fc4DO5wVBhoro4g/G851WwTRY68\n/UHlDekCgYBGEMBiOUu3H3ZrHpv7w3qVEAEJtWUgRg3HC+oyIFf29uSKCzU/Jvbe\n11T7rdu05HOt0+V/mcYfH/ElkQ6Aj7vErkrDONsoCC9s0fw4MNlkNVwXmxSAxnFZ\nc8bA31O3uNVkJ5GPjoJscLyE1U1LBtNaxlJQDYggSmKn9aLeY7QZDQ==\n-----END RSA PRIVATE KEY-----"
        APP_ID                 = "1012223"
        APP_INSTALLATION_ID    = "55454507"
        APP_WEBHOOK_SECRET     = "abc123"
        CLIENT_REDIRECT_URL    = "http://khoury-classroom-frontend-prod-east-2.s3-website.us-east-2.amazonaws.com/oauth/callback"
        CLIENT_ID              = "Ov23lip0iKQiglFSl90d"
        CLIENT_SECRET          = "2c11232818c8e5ab02114a03323856cde07001fe"
        CLIENT_URL             = "https://github.com/login/oauth/authorize"
        CLIENT_TOKEN_URL       = "https://github.com/login/oauth/access_token"
        CLIENT_SCOPES          = "repo,read:org,classroom"
        CLIENT_JWT_SECRET      = "H96GlVdJaaz9+rvUxHuTfI4owA8XyiH1eTsaup1LkTg="
        DATABASE_URL           = "postgresql://${var.db_user}:${var.db_password}@terraform-20241016161021390000000005.che2wqk4qebu.us-east-2.rds.amazonaws.com:5432/${var.db_name}"
        DB_HOST                = "terraform-20241016161021390000000005.che2wqk4qebu.us-east-2.rds.amazonaws.com"
        DB_PORT                = "5432"
        DB_NAME                = "khouryclassroomdb"
        DB_USER                = var.db_user
        DB_PASSWORD            = var.db_password
  }
}

# Create the ECS task definition
resource "aws_ecs_task_definition" "app" {
    family                   = "cb-app-task"
    execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
    requires_compatibilities = ["FARGATE"]
    network_mode             = "awsvpc"
    cpu                      = var.fargate_cpu
    memory                   = var.fargate_memory
    container_definitions = jsonencode([{
        name      = "cb-app"
        image     = var.app_image
        cpu       = var.fargate_cpu
        memory    = var.fargate_memory
        logConfiguration = {
            logDriver = "awslogs"
            options = {
                "awslogs-group"         = "/ecs/cb-app"
                "awslogs-region"        = var.aws_region
                "awslogs-stream-prefix" = "ecs"
            }
        }
        portMappings = [{
            containerPort = var.app_port
            hostPort      = var.app_port
        }]
        environment = [
            for key, value in local.app_secrets : {
                name  = key
                value = value
            }
        ]
    }])

    depends_on = [ aws_db_instance.main ]
}

resource "aws_ecs_service" "main" {
    name            = "cb-service"
    cluster         = aws_ecs_cluster.main.id
    task_definition = aws_ecs_task_definition.app.arn
    desired_count   = var.app_count
    launch_type     = "FARGATE"

    network_configuration {
        security_groups  = [aws_security_group.ecs_tasks.id]
        subnets          = aws_subnet.private.*.id
        assign_public_ip = true
    }

    load_balancer {
        target_group_arn = aws_alb_target_group.app.id
        container_name   = "cb-app"
        container_port   = var.app_port
    }

    depends_on = [aws_alb_listener.front_end, aws_iam_role_policy_attachment.ecs-task-execution-role-policy-attachment]
}