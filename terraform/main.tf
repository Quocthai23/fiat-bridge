terraform {
  required_version = ">= 1.5.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
  }
}

# Configure default AWS Provider
provider "aws" {
  region = var.aws_region
}

# Configure default GCP Provider
provider "google" {
  project = var.gcp_project_id
  region  = var.gcp_region
}

variable "aws_region" {
  type        = string
  default     = "us-east-1"
  description = "AWS region for resources"
}

variable "gcp_project_id" {
  type        = string
  default     = "my-gcp-project-id"
  description = "GCP Project ID"
}

variable "gcp_region" {
  type        = string
  default     = "us-central1"
  description = "GCP region for resources"
}
