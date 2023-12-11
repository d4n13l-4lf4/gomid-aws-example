terraform {
  cloud {
    # omit settings to be taken from deployment
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }

    local = {
      source  = "hashicorp/local"
      version = "2.4.0"
    }
  }

  required_version = "~> 1.6.4"
}

locals {
  service_name = "gomid-aws-example"
}

provider "aws" {
  region = "us-west-2"
}