variable "domain_name" {
  description = "The domain name for the React app"
  type        = string
  default     = "gitmarks.org"
}

variable "app_port" {
  description = "Application port"
  type        = number
  default     = 8080
}

variable "db_port" {
  description = "Database port"
  type        = number
  default     = 5432
}

variable "aws_account_id" {
  description = "AWS account ID"
  type        = string
  default     = "058264409130"
}