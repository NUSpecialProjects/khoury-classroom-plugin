# Create an S3 bucket to store SQL migration scripts
resource "aws_s3_bucket" "sql_migrations_bucket" {
    bucket = "gitmarks-migrations-${var.aws_account_id}"

    tags = {
        Name        = "SQL Migrations Bucket"
        Environment = "gitmarks"
    }
}

# Add versioning to the bucket
resource "aws_s3_bucket_versioning" "sql_migrations_bucket_versioning" {
    bucket = aws_s3_bucket.sql_migrations_bucket.bucket

    versioning_configuration {
        status = "Enabled"
    }
}

# Configure bucket policy to allow access from Lambda
resource "aws_s3_bucket_policy" "sql_migrations_policy" {
    bucket = aws_s3_bucket.sql_migrations_bucket.id

    policy = jsonencode({
        Version = "2012-10-17"
        Statement = [
            {
                Effect = "Allow"
                Principal = {
                    Service = "lambda.amazonaws.com"
                }
                Action = "s3:GetObject"
                Resource = "arn:aws:s3:::${aws_s3_bucket.sql_migrations_bucket.id}/*"
            }
        ]
    })
}
