variable "deployment_bucket" {
  type        = string
  description = "The deployment bucket name"
}

variable "stage" {
  type        = string
  description = "Stage for deployment"
}

variable "acm_certificate" {
  type        = string
  description = "ARN of public certificate for API"
}

variable "domain_name" {
  type        = string
  description = "Root domain for API"
}

variable "route_domain_zone_id" {
  type        = string
  description = "Route53 Zone ID of domain"
}