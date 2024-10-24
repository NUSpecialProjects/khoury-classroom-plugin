# main.tf

resource "aws_db_instance" "main" {
  allocated_storage      = 20 # Storage in GB
  engine                 = "postgres"
  engine_version         = "16.2"
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

output "app_secrets" {
  description = "Application secrets"
  value = {
    APP_PRIVATE_KEY     = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAoHQT+RKgjEdkq7SBiQ0LkKRaWyeh8AD8zHBovlVa6yV/ik2C\nieDwIWGtlx2OU3gqf13ZZDEGvNwi1PDt3m+ny3Xsy901vyvlGSyqfmAn0oTj1rJu\nnOq9MfcXWfMjbPXyaONRj0/1aHNN2Bvl5Ye/dgCOVwYZOgsu8Ka33aH5dRWfnPf7\ng4XQoO1oBoe4EufnPJ2uciTTZDuvg1rnav6e+uMEKopNpKFvtg4GAQl4x0Qq6xkQ\n6VfA8opnit7mOaWcDdXc3thgN7CHTPCD4XqvcmFsBJNclGEfjCbdavrIIp4JKkpw\nDmk1Ou6Lz78qnbTxQ9CYeBJDfvarLPDJyKGsDQIDAQABAoIBAQCgRlMNIYYtmcL9\noTkjZVyABywakeQ4kUP0EvUN6sT+zl4wEGysvXwgXCnCIUviJM6Om3hjlHVegaZp\nfqCc6Ht7yTfYDAd8BqS6GNvVkMc2infsJiBHrlN+bYtt1mk0lhimnSsDNKO2yjag\nAH4MYSTnAncshnL8f99Lk71mLj24rWOa7SjmRKga6pgNacHZ8cCBR0TKmMJNlEFL\noHPZveTmg+uu7obaGJ7Lo08hQJsocB7nrqUtrnpcPGDz+yXa+nRwYCTza4+4AKvN\nkX0680js2adth8ELpL5Gy+9ukp8LMf7gArqSd+ZAGhSgSIaklLfixGHcWBgry2gx\nOepXnmMhAoGBANNFzqSYNH+MlxbW8n1Wj/cp396R1SUKbXP6oFPMUS0mMtOU7RFS\nqSoNIgT5bhJk5B8/D+f4XEZNWWMBnq5EgL39hmstULkl26htKPKHmDvX5X3hkXRH\nyAKTeYWEblE+PpDqKFE9O3IpPovyJ0Oe19U+100i/Gry6bgRDzHG3DAHAoGBAMJs\nDnOH789IQjiQprhj+3ue9Xl2JEU959PbsCk2jmp2NMuuYf3fyK9iZhCXbBysmaSH\nOUGbl86SFOdwtTcGwYKzyNaCUdzNYTqLUoWiJkLMn3PZbggiJ2MQ2e42UNiCAsQb\ndjPh5/66dyMlJuBH3kfh0yP5ShK4heEiMULYSRZLAoGBAJ/TTVICuqRLDPmAPg1H\ncL1/9hV/qQjObKKyVJtQE5DeNtEM9pKGP+bJ7JRqxTQxEsn4gOXxYozkctyNGyem\nNuaDZi6qJ0kJNLSjb7iZjzamSrwB6nFW5B3exq2U04euWNJz8XATrGbegKyJ0d47\nyfdOBL4b22xkux49+Yqkb2n9AoGAb45Y7Gl/bExl0tcNEpgr4E7hQwRK44AV2TYg\n6kTniqawvH4es/EH0bqAHd0Ep59RuVntvHtuq5Secf31vNEfj8Ng5dR47FzcAR+Y\nBh14HrQSegK0Y+5U8z7kDQ8VbGWM+MFZHYPt/fc4DO5wVBhoro4g/G851WwTRY68\n/UHlDekCgYBGEMBiOUu3H3ZrHpv7w3qVEAEJtWUgRg3HC+oyIFf29uSKCzU/Jvbe\n11T7rdu05HOt0+V/mcYfH/ElkQ6Aj7vErkrDONsoCC9s0fw4MNlkNVwXmxSAxnFZ\nc8bA31O3uNVkJ5GPjoJscLyE1U1LBtNaxlJQDYggSmKn9aLeY7QZDQ==\n-----END RSA PRIVATE KEY-----"
    APP_ID              = "1012223"
    APP_INSTALLATION_ID = "55454507"
    APP_WEBHOOK_SECRET  = "abc123"
    CLIENT_REDIRECT_URL = "https://gitmarks.org/oauth/callback"
    CLIENT_ID           = "Ov23lip0iKQiglFSl90d"
    CLIENT_SECRET       = "2c11232818c8e5ab02114a03323856cde07001fe"
    CLIENT_URL          = "https://github.com/login/oauth/authorize"
    CLIENT_TOKEN_URL    = "https://github.com/login/oauth/access_token"
    CLIENT_SCOPES       = "repo,read:org,classroom"
    CLIENT_JWT_SECRET   = "H96GlVdJaaz9+rvUxHuTfI4owA8XyiH1eTsaup1LkTg="
    DB_PORT             = var.db_port
    DATABASE_URL        = "postgresql://${var.db_user}:${var.db_password}@${aws_db_instance.main.endpoint}/${var.db_name}"
    DB_HOST             = aws_db_instance.main.endpoint
    DB_NAME             = var.db_name
    DB_USER             = var.db_user
    DB_PASSWORD         = var.db_password
  }
}