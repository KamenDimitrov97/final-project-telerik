# Hello World Go Project with DevOps Infrastructure

## Overview

This is a simple Hello World Go project that serves as a starting point for building a DevOps infrastructure. The goal is to demonstrate a basic Go application and provide guidelines for setting up CI/CD, containerization, and deployment using kubernetes with flux, Google cloud as a public cloud, helmcharts as deployments and .

## Project Structure

- `.deployment` dir contains Helm charts and configurations for deploying the Go application.
    - `go-app-chart` dir contains deployment helm charts.
    - `production-chart` dir contains production helm charts.
     It includes subdirectories such as `go-app-chart` and `production-chart`, which host Helm charts tailored for specific deployment environments (e.g., development and production).
- `.github` dir houses workflows defined using GitHub Actions.
    - `workflows` subdir, you'll find configurations like `ci.yaml`, `development.yaml`, and `production.yaml`, which automate processes such as continuous integration for testing and deployment. These workflows are designed to streamline development and maintain consistency across different stages of the project.
- `config` dir contains configurations for the go-app.
```
├── .deployment
│   ├── go-app-chart
│   ├── production-chart
├── .github
│   └── workflows
├── config
```

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) installed on your local machine.
- [Docker](https://www.docker.com/get-started) for containerization.

### Running Locally

```bash
# Clone the repository
git clone https://github.com/yourusername/hello-world-go.git

# Check available make scripts
make help

# Run the application
make run

# Test it out
curl http://localhost:3333
```


### Building containers

```bash
# Check available make scripts
make help

# Run the application
make run-container

# Test it out
curl http://localhost:3333
```


## DevOps Infrastructure

### Continuous Integration (CI)

This project uses GitHub Actions for CI. The workflow is defined in the `.github/workflows/ci.yml`.
Workflow is run on
Steps For feature branches:
- Pre-Build:
    Linting commits using wagoid/commitlint-github-action@v5.
    (Optional) Git repository leak detection using gitleaks/gitleaks-action@v2 (commented out due to issue ["#1316"](https://github.com/gitleaks/gitleaks/issues/1316)).

- Code Security:
    Running the Gosec security scanner using securego/gosec@master.

- Dependency Security:
    Running govulncheck to check for vulnerabilities in project dependencies using golang/govulncheck-action@v1.

- Linting:
    Running golangci-lint for linting Go code using golangci/golangci-lint-action@v3.

- Test and Build:
    Running tests using go test.
    Checking editorconfig compliance using editorconfig-checker/action-editorconfig-checker@main.

Steps For development environment:
- Build and push Docker image:
    Builds and pushes a docker image tagged with `master-runner_id`. Development kubernetes environment will pull the latest changes and create new pods.

- Run Trivy vulnerability scanner:
    Checks the image for vunlerabilities.

Steps For production environment:
- Build and push Docker image:
    Builds and pushes a docker image tagged with `x.x.x` on a new tag being published. Production kubernetes environment will pull the latest changes and create new pods.

- Run Trivy vulnerability scanner:
    Checks the image for vunlerabilities.

The project is containerized using Docker. The Dockerfile defines the image.

### Additional check with GitGuardian on repo.

<img alt="gitleaks badge" src="https://img.shields.io/badge/protected%20by-gitleaks-blue">

### Kubernetes Infrastructure
Automated Deployments with Flux.

This project is set up for automated deployments to a Kubernetes cluster using Flux. The Kubernetes manifests for development and production can be found in the ["Infrastructure repo"](https://github.com/KamenDimitrov97/final-project-infrastructure) and helmcharts can be found in the .deployment directory here.

Flux continuously monitors the Dockerhub repository for changes and automatically deploys updated helmcharts to the specified namespaces by commiting the immutable tags to the infrastructure repo. In this case, deployments are automated for the telerik namespace.

## Contributing

We welcome contributions! Feel free to open issues, suggest improvements, or submit pull requests.
However please follow this ["guide"](./contributions/CONTRIBUTING.md)!
