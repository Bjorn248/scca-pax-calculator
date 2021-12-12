terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3"
    }
  }

  backend "s3" {
    bucket = "bjornaws-terraform"
    key    = "paxcalculator"
    region = "us-east-1"
  }
}

provider "aws" {
  region                  = "us-east-2"
  shared_credentials_file = "~/.aws/credentials"
}

module "s3-static-site" {
  source = "github.com/Bjorn248/s3-static-site"

  root-domain            = "paxcalculator.com"
  target-domain          = "www.paxcalculator.com"
  cloudfront-price-class = "PriceClass_100"

  global-tags = {
    project = "paxcalculator"
  }
}

module "s3-website-deployment-iam" {
  source = "github.com/Bjorn248/s3-website-deployment-iam"

  s3-bucket-name = module.s3-static-site.s3_bucket_name

  global-tags = {
    project = "paxcalculator"
  }
}
