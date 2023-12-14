# AWS Lambda example
[![Deployment](https://github.com/d4n13l-4lf4/gomid-aws-example/actions/workflows/terraform-apply.yaml/badge.svg)](https://github.com/d4n13l-4lf4/gomid-aws-example/actions/workflows/terraform-apply.yaml)
[![codecov](https://codecov.io/gh/d4n13l-4lf4/gomid-aws-example/graph/badge.svg?token=Ax9KxOOX58)](https://codecov.io/gh/d4n13l-4lf4/gomid-aws-example)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit)](https://github.com/pre-commit/pre-commit)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

An example usage of gomid. This project uses terraform to provision infrastructure on AWS.

### Usage
The next steps should be followed to deploy this application. 

Note: this guide assumes you have an AWS account and Terraform Cloud account.

### Step 1: Configure terraform cloud
Execute the following command to login into terraform cloud.
```bash
terraform login
```

### Step 2: Initialize terraform
Initialize terraform.
```bash
terraform init
```
### Step 3: Configure a new terraform workspace
Configure a new terraform workspace with the variables found at this repo [Terraform doc](terraform.md). Also, you wil have to setup AWS Credentials for terraform cloud remote execution.

### Step 4: Terraform plan
Next, overview the creation of your resources with planning.
```bash
terraform plan
```

### Step 5: Terraform apply
Finally, apply the configuration to start this deployment.
```bash
terraform apply
```
