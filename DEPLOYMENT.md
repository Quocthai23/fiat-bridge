# AWS and GCP Deployment Guide

This document provides detailed instructions for setting up the infrastructure using Terraform (IaC), configuring the Helm Chart, and automating CI/CD via GitHub Actions for the **Fiat Bridge** project.

---

## 1. Provisioning Infrastructure with Terraform

The `terraform/` directory contains configurations to automatically create a VPC, EKS (AWS) or GKE (GCP) cluster, along with a Docker Container Registry.

### Local Prerequisites:
- [Terraform CLI](https://developer.hashicorp.com/terraform/downloads) (>= 1.5.0)
- AWS CLI or Google Cloud SDK.

### Execution Steps:
1. Navigate to the `terraform` directory:
   ```bash
   cd terraform
   ```
2. Initialize the Terraform provider and backend:
   ```bash
   terraform init
   ```
3. Preview the execution plan (Dry Run):
   ```bash
   terraform plan -var="gcp_project_id=YOUR_PROJECT_ID"
   ```
4. Apply the infrastructure changes:
   ```bash
   terraform apply -var="gcp_project_id=YOUR_PROJECT_ID" -auto-approve
   ```

> [!NOTE]
> Upon successful application, you will receive the AWS ECR and GCP Artifact Registry URLs, as well as the Cluster IDs needed for CI/CD configuration.

---

## 2. Configuring GitHub Secrets for CI/CD

The GitHub Actions workflows require secure credentials to push images to the Container Registry and deploy them to the cluster.

Navigate to your GitHub repository: **Settings -> Secrets and variables -> Actions**, and create the following Repository Secrets:

### For Both Clouds (Database):
- `DB_PASSWORD`: The PostgreSQL database password that will be injected into the Helm Chart.

### For AWS Configuration:
- `AWS_ACCESS_KEY_ID`: The Access Key ID of the IAM User with ECR/EKS access permissions.
- `AWS_SECRET_ACCESS_KEY`: The Secret Access Key of that IAM User.

### For GCP Configuration:
- `GCP_PROJECT_ID`: Your GCP Project ID.
- `GCP_SA_KEY`: The JSON key content of the GCP Service Account with Artifact Registry and GKE Admin privileges.

---

## 3. Manual Configuration and Deployment via Helm

If you wish to deploy manually from your local machine to the Kubernetes Cluster (EKS/GKE):

1. Update your Kubeconfig to point to the cluster:
   - **AWS EKS**:
     ```bash
     aws eks update-kubeconfig --name fiat-bridge-eks --region us-east-1
     ```
   - **GCP GKE**:
     ```bash
     gcloud container clusters get-credentials fiat-bridge-gke --zone us-central1-a
     ```
2. Execute the Helm deployment (overriding parameters from `values.yaml`):
   ```bash
   helm upgrade --install fiat-bridge ./helm/fiat-bridge \
     --namespace default \
     --set image.repository=<YOUR_REGISTRY_URL>/fiat-bridge \
     --set image.tag=latest \
     --set secretEnv.dbPassword="my-strong-password"
   ```

---

## 4. Verifying the Deployment Status

After the CI/CD pipeline finishes or the manual Helm deployment completes:

1. Check the running Pods:
   ```bash
   kubectl get pods -l app.kubernetes.io/name=fiat-bridge
   ```
2. View the application logs:
   ```bash
   kubectl logs -f deployment/fiat-bridge
   ```
3. Check Service & Ingress information (to retrieve the external IP for access):
   ```bash
   kubectl get service fiat-bridge
   kubectl get ingress fiat-bridge
   ```
