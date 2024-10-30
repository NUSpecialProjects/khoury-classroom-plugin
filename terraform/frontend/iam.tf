# iam.tf

# ------------------------------------------------
# CloudFront Origin Access Identity
# ------------------------------------------------

resource "aws_cloudfront_origin_access_identity" "oai" {
  comment = "OAI for S3 access via CloudFront"
}
