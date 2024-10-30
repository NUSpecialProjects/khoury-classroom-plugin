# backend/alb.tf

# Load Balancer Resource
resource "aws_lb" "main" {
  name                 = "gitmarks-load-balancer"
  internal             = false
  subnets              = var.public_subnet_ids
  security_groups      = [var.lb_sg_id]
  load_balancer_type   = "application"
  preserve_host_header = true
}

# Target Group Resource
resource "random_id" "id" {
  byte_length = 8
}
resource "aws_lb_target_group" "app" {
  name        = "gitmarks-tg-${random_id.id.hex}"
  port        = var.app_port
  protocol    = "HTTP"
  vpc_id      = var.vpc_id
  target_type = "ip"

  lifecycle {
    create_before_destroy = true
  }
}

# ALB Listener Resource
resource "aws_lb_listener" "front_end_https" {
  load_balancer_arn = aws_lb.main.arn
  port              = 443
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-2016-08"
  certificate_arn   = aws_acm_certificate.api_cert.arn
  default_action {
    target_group_arn = aws_lb_target_group.app.arn
    type             = "forward"
  }

  depends_on = [aws_lb.main, aws_lb_target_group.app]
}
resource "aws_lb_listener" "front_end_http" {
  load_balancer_arn = aws_lb.main.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type = "redirect"
    redirect {
      port        = "443"
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}
