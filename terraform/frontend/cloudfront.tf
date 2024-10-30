# frontend/cloudfront.tf

resource "aws_cloudfront_distribution" "frontend" {
  # The S3 bucket with our React App
  origin {
    domain_name = "${aws_s3_bucket.frontend.bucket}.s3.amazonaws.com"
    origin_id   = aws_s3_bucket.frontend.id

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.oai.cloudfront_access_identity_path
    }
  }

  # CloudFront Configuration
  enabled             = true
  is_ipv6_enabled     = true
  default_root_object = var.index_document

  # SSL Configuration for HTTPS
  viewer_certificate {
    acm_certificate_arn      = aws_acm_certificate.frontend_cert.arn
    ssl_support_method       = "sni-only"
    minimum_protocol_version = "TLSv1.2_2019"
  }

  # Enforcing HTTPS and Caching Behavior
  default_cache_behavior {
    target_origin_id       = aws_s3_bucket.frontend.id
    viewer_protocol_policy = "redirect-to-https"

    allowed_methods = ["GET", "HEAD", "OPTIONS", "PUT", "POST", "PATCH", "DELETE"]
    cached_methods  = ["GET", "HEAD", "OPTIONS"]

    forwarded_values {
      query_string = false

      cookies {
        forward           = "whitelist"
        whitelisted_names = ["jwt_token"]
      }

      headers = ["Origin", "Access-Control-Request-Method", "Access-Control-Request-Headers"]
    }

    # Caching TTLs (Time To Live)
    min_ttl     = 0
    default_ttl = 3600
    max_ttl     = 86400
    compress    = true
  }

  # No restrictions
  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  # Routes to homepage on error
  custom_error_response {
    error_code         = 404
    response_code      = 200
    response_page_path = "/${var.error_document}"
  }

  # Custom Domain Setup
  aliases = [var.domain_name]

  # Pricing options
  price_class = "PriceClass_100"

  depends_on = [
    aws_acm_certificate.frontend_cert,
    aws_cloudfront_origin_access_identity.oai,
    aws_s3_bucket.frontend
  ]
}
