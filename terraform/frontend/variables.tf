variable "aws_account_id" {
  description = "The AWS account ID"
  default     = "058264409130"
}

variable "index_document" {
  description = "The index document for the S3 bucket"
  type        = string
  default     = "index.html"
}

variable "error_document" {
  description = "The error document for the S3 bucket"
  type        = string
  default     = "index.html"
}

variable "domain_name" {
  description = "The domain name for the React app"
  type        = string
}