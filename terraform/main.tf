terraform {
  required_version = ">= 1.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

# S3 Bucket
resource "aws_s3_bucket" "nicole_rankings" {
  bucket = var.bucket_name

  tags = {
    Name        = "Nicole Rankings Data"
    Environment = var.environment
    ManagedBy   = "Terraform"
  }
}

# S3 Bucket Public Access Block Configuration
resource "aws_s3_bucket_public_access_block" "nicole_rankings" {
  bucket = aws_s3_bucket.nicole_rankings.id

  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
}

# S3 Bucket Policy for Public Read Access
resource "aws_s3_bucket_policy" "nicole_rankings_public_read" {
  bucket = aws_s3_bucket.nicole_rankings.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid       = "PublicReadGetObject"
        Effect    = "Allow"
        Principal = "*"
        Action    = "s3:GetObject"
        Resource  = "${aws_s3_bucket.nicole_rankings.arn}/*"
      }
    ]
  })

  depends_on = [aws_s3_bucket_public_access_block.nicole_rankings]
}

# Reference existing OIDC Provider for GitHub Actions
data "aws_iam_openid_connect_provider" "github_actions" {
  url = "https://token.actions.githubusercontent.com"
}

# IAM Role for GitHub Actions
resource "aws_iam_role" "github_actions" {
  name               = var.github_actions_role_name
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Federated = data.aws_iam_openid_connect_provider.github_actions.arn
        }
        Action = "sts:AssumeRoleWithWebIdentity"
        Condition = {
          StringEquals = {
            "token.actions.githubusercontent.com:aud" = "sts.amazonaws.com"
          }
          StringLike = {
            "token.actions.githubusercontent.com:sub" = "repo:${var.github_repository}:*"
          }
        }
      }
    ]
  })

  tags = {
    Name        = "GitHub Actions Role"
    Purpose     = "Upload ranking data to S3"
    ManagedBy   = "Terraform"
  }
}

# IAM Policy for S3 PutObject
resource "aws_iam_policy" "github_actions_s3_upload" {
  name        = "${var.bucket_name}-upload-policy"
  description = "Allow GitHub Actions to upload objects to S3 bucket"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "AllowS3Upload"
        Effect = "Allow"
        Action = [
          "s3:PutObject",
          "s3:PutObjectAcl"
        ]
        Resource = "${aws_s3_bucket.nicole_rankings.arn}/*"
      },
      {
        Sid    = "AllowS3ListBucket"
        Effect = "Allow"
        Action = [
          "s3:ListBucket"
        ]
        Resource = aws_s3_bucket.nicole_rankings.arn
      }
    ]
  })
}

# Attach Policy to IAM Role
resource "aws_iam_role_policy_attachment" "github_actions_s3_upload" {
  role       = aws_iam_role.github_actions.name
  policy_arn = aws_iam_policy.github_actions_s3_upload.arn
}

# CORS Configuration (if needed for browser access)
resource "aws_s3_bucket_cors_configuration" "nicole_rankings" {
  bucket = aws_s3_bucket.nicole_rankings.id

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["GET", "HEAD"]
    allowed_origins = ["*"]
    expose_headers  = ["ETag"]
    max_age_seconds = 3000
  }
}

# Versioning disabled
resource "aws_s3_bucket_versioning" "nicole_rankings" {
  bucket = aws_s3_bucket.nicole_rankings.id

  versioning_configuration {
    status = "Disabled"
  }
}
