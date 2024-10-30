# s3.tf

resource "aws_s3_bucket" "frontend" {
  bucket        = "khoury-classroom-frontend-${var.aws_account_id}"
  force_destroy = true
}

resource "aws_s3_bucket_website_configuration" "frontend_config" {
  bucket = aws_s3_bucket.frontend.id

  index_document {
    suffix = var.index_document
  }
  error_document {
    key = var.error_document
  }
}

resource "aws_s3_bucket_policy" "frontend_policy" {
  bucket = aws_s3_bucket.frontend.id

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = ["s3:GetObject"],
        Effect = "Allow",
        Principal = {
          AWS = "${aws_cloudfront_origin_access_identity.oai.iam_arn}"
        },
        Resource = "${aws_s3_bucket.frontend.arn}/*"
      }
    ]
  })
}

resource "aws_s3_bucket_ownership_controls" "frontend_ownership" {
  bucket = aws_s3_bucket.frontend.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

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
