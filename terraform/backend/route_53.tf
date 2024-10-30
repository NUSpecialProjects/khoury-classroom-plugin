# backend/route_53.tf

# Fetch the Route 53 Zone
data "aws_route53_zone" "zone" {
  name         = var.domain_name
  private_zone = false
}

# Create DNS validation records for API cert
resource "aws_route53_record" "api_cert_validation" {
  for_each = {
    for dvo in aws_acm_certificate.api_cert.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      type   = dvo.resource_record_type
      record = dvo.resource_record_value
    }
  }

  zone_id = data.aws_route53_zone.zone.zone_id
  name    = each.value.name
  type    = each.value.type
  ttl     = 60
  records = [each.value.record]
}

# Map ALB IP to API DNS
resource "aws_route53_record" "alb_alias" {
  zone_id = data.aws_route53_zone.zone.zone_id
  name    = "api.${var.domain_name}"
  type    = "A"

  alias {
    name                   = aws_lb.main.dns_name
    zone_id                = aws_lb.main.zone_id
    evaluate_target_health = false
  }

  depends_on = [
    aws_acm_certificate_validation.api_cert_validation
  ]
}