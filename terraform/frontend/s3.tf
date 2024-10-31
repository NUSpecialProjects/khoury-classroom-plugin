# frontend/s3.tf

# Bucket to host React App
resource "aws_s3_bucket" "frontend" {
  bucket        = "gitmarks-frontend-${var.aws_account_id}"
  force_destroy = true
}
resource "aws_s3_bucket_ownership_controls" "frontend_ownership" {
  bucket = aws_s3_bucket.frontend.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

# Configure bucket for static website hosting
resource "aws_s3_bucket_website_configuration" "frontend_config" {
  bucket = aws_s3_bucket.frontend.id

  index_document {
    suffix = var.index_document
  }
  error_document {
    key = var.error_document
  }
}

# Allow public access
resource "aws_s3_bucket_public_access_block" "frontend_public_access" {
  bucket = aws_s3_bucket.frontend.id

  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
}
resource "aws_s3_bucket_acl" "frontend_acl" {
  depends_on = [
    aws_s3_bucket_ownership_controls.frontend_ownership,
    aws_s3_bucket_public_access_block.frontend_public_access,
  ]

  bucket = aws_s3_bucket.frontend.id
  acl    = "public-read"
}
