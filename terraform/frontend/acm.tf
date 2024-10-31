# frontend/acm.tf

# Request SSL Certificate for CloudFront
resource "aws_acm_certificate" "frontend_cert" {
  provider          = aws.cert
  domain_name       = var.domain_name
  validation_method = "DNS"

  lifecycle {
    create_before_destroy = true
  }
}

# Validate the CloudFront certificate
resource "aws_acm_certificate_validation" "frontend_cert_validation" {
  provider                = aws.cert
  certificate_arn         = aws_acm_certificate.frontend_cert.arn
  validation_record_fqdns = [for record in aws_route53_record.frontend_cert_validation : record.fqdn]
}