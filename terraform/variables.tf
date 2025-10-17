variable "aws_region" {
  description = "AWS region for resources"
  type        = string
  default     = "ap-northeast-1"
}

variable "bucket_name" {
  description = "Name of the S3 bucket for ranking data"
  type        = string
  default     = "shibu1x-public"
}

variable "environment" {
  description = "Environment name"
  type        = string
  default     = "production"
}

variable "github_actions_role_name" {
  description = "IAM role name for GitHub Actions"
  type        = string
  default     = "nicole-github-actions-role"
}

variable "github_repository" {
  description = "GitHub repository in the format 'owner/repo'"
  type        = string
  default     = "shibu1x/nicole"
}
