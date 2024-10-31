# frontend/iam.tf

# ------------------------------------------------
# CloudFront Origin Access Identity
# ------------------------------------------------

resource "aws_cloudfront_origin_access_identity" "oai" {
  comment = "OAI for S3 access via CloudFront"
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