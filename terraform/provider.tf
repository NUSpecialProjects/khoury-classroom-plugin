# provider.tf

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.70.0"
    }
  }
}

provider "aws" {
  region  = "us-east-2"
}

provider "aws" {
  region  = "us-east-1"
  alias   = "us-east-1"
}