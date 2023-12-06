terraform {
    cloud {
        organization = "d4n13l-4lf4"
        workspaces {
          name = "gomid-aws-lambda"
        }
    }

    required_providers {
      aws = {
        source = "hashicorp/aws"
        version = "~> 5.0"
      }
    }
}

provider "aws" {
        region = "us-east-1"
}