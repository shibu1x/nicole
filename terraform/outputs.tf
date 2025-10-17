output "bucket_name" {
  description = "Name of the S3 bucket"
  value       = aws_s3_bucket.nicole_rankings.id
}

output "bucket_arn" {
  description = "ARN of the S3 bucket"
  value       = aws_s3_bucket.nicole_rankings.arn
}

output "bucket_domain_name" {
  description = "Domain name of the S3 bucket"
  value       = aws_s3_bucket.nicole_rankings.bucket_domain_name
}

output "bucket_regional_domain_name" {
  description = "Regional domain name of the S3 bucket"
  value       = aws_s3_bucket.nicole_rankings.bucket_regional_domain_name
}

output "github_actions_role_name" {
  description = "IAM role name for GitHub Actions"
  value       = aws_iam_role.github_actions.name
}

output "github_actions_role_arn" {
  description = "ARN of the GitHub Actions IAM role"
  value       = aws_iam_role.github_actions.arn
}

output "oidc_provider_arn" {
  description = "ARN of the GitHub Actions OIDC provider"
  value       = data.aws_iam_openid_connect_provider.github_actions.arn
}

output "s3_upload_policy_arn" {
  description = "ARN of the S3 upload policy"
  value       = aws_iam_policy.github_actions_s3_upload.arn
}
