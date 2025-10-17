# Nicole Terraform Configuration

This directory contains Terraform configuration for provisioning AWS infrastructure for the Nicole project.

## Resources Created

- **S3 Bucket**: Public bucket for storing ranking JSON data
  - Public read access (GetObject) for all users
  - CORS configuration for browser access
  - Versioning disabled
- **OIDC Provider**: References existing GitHub Actions OIDC provider
- **IAM Role**: Role for GitHub Actions with OIDC authentication
- **IAM Policy**: Allows PutObject and ListBucket on the S3 bucket

## Prerequisites

1. AWS CLI installed and configured
2. Terraform installed (>= 1.0)
3. AWS credentials with sufficient permissions to create S3 buckets and IAM resources

## Usage

### Initialize Terraform

```bash
cd terraform
terraform init
```

### Plan Changes

```bash
terraform plan
```

### Apply Configuration

```bash
terraform apply
```

### View Outputs

```bash
terraform output
```

## Configuration Variables

You can customize variables in `variables.tf` or create a `terraform.tfvars` file:

```hcl
aws_region                = "ap-northeast-1"
bucket_name               = "shibu1x-public"
environment               = "production"
github_actions_role_name  = "nicole-github-actions-role"
github_repository         = "shibu1x/nicole"
```

## GitHub Actions Setup

After applying the Terraform configuration:

1. **Get the IAM Role ARN**:
   ```bash
   terraform output github_actions_role_arn
   ```

2. **Add Role ARN to GitHub Secrets**:
   - Go to your repository Settings → Secrets and variables → Actions
   - Add a new repository secret:
     - Name: `AWS_ROLE_ARN`
     - Value: The ARN from step 1 (e.g., `arn:aws:iam::123456789012:role/nicole-github-actions-role`)

3. **Update GitHub Actions Workflow** to use OIDC authentication:
   ```yaml
   permissions:
     id-token: write
     contents: read

   steps:
     - name: Configure AWS credentials
       uses: aws-actions/configure-aws-credentials@v4
       with:
         role-to-assume: ${{ secrets.AWS_ROLE_ARN }}
         aws-region: ap-northeast-1
   ```

4. **No long-lived access keys required** - OIDC authentication uses temporary credentials

## Security Notes

- Uses OIDC (OpenID Connect) for secure, keyless authentication
- No long-lived access keys required
- IAM role can only be assumed by GitHub Actions from the specified repository (default: `shibu1x/nicole`)
- The S3 bucket allows public read access for all objects
- Only the GitHub Actions role can write to the bucket
- Temporary credentials are automatically rotated

## Destroy Resources

To remove all created resources:

```bash
terraform destroy
```

**Warning**: This will delete the S3 bucket and all its contents.
