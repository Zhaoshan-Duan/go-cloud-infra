# Go Cloud infrastructure with AWS CDK

This project provides a modular cloud infrastructure solution using Go and AWS CDK. It provisions serverless resources such as Lambda functions, APIs, and databases, enabling scalable and maintainable cloud deployments.

## Features
- Infrastructure-as-code with AWS CDK
- Go-based Lambda functions
- Modular API and database components
- Automated deployment and resource management

## Structure
- `go_app.go` – Main application logic
- `lambda/` – Lambda function code and modules
- `cdk.out/` – Generated CDK assets and templates

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests
 
