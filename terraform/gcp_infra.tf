# GCP Infrastructure Configuration

# 1. Google Artifact Registry Repository for Docker images
resource "google_artifact_registry_repository" "fiat_bridge_repo" {
  location      = var.gcp_region
  repository_id = "fiat-bridge-repo"
  description   = "Docker registry for Fiat Bridge microservices"
  format        = "DOCKER"
}

# 2. VPC Network
resource "google_compute_network" "vpc_network" {
  name                    = "fiat-bridge-vpc"
  auto_create_subnetworks = false
}

# 3. Subnet for GKE
resource "google_compute_subnetwork" "gke_subnet" {
  name          = "fiat-bridge-gke-subnet"
  ip_cidr_range = "10.10.0.0/16"
  region        = var.gcp_region
  network       = google_compute_network.vpc_network.id

  secondary_ip_range {
    range_name    = "pods"
    ip_cidr_range = "10.20.0.0/16"
  }

  secondary_ip_range {
    range_name    = "services"
    ip_cidr_range = "10.30.0.0/20"
  }
}

# 4. GKE Cluster
resource "google_container_cluster" "gke" {
  name     = "fiat-bridge-gke"
  location = "${var.gcp_region}-a"

  # We create a network-policy-enabled cluster with a custom subnet
  network    = google_compute_network.vpc_network.name
  subnetwork = google_compute_subnetwork.gke_subnet.name

  ip_allocation_policy {
    cluster_secondary_range_name  = "pods"
    services_secondary_range_name = "services"
  }

  # We recommend using a separate node pool resource, so we delete the default node pool on creation
  remove_default_node_pool = true
  initial_node_count       = 1
}

# 5. GKE Custom Node Pool
resource "google_container_node_pool" "primary_nodes" {
  name       = "fiat-bridge-node-pool"
  location   = google_container_cluster.gke.location
  cluster    = google_container_cluster.gke.name
  node_count = 2

  node_config {
    preemptible  = false
    machine_type = "e2-medium"

    # Google Service Account permissions
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]

    labels = {
      env = "production"
    }

    tags = ["gke-node", "fiat-bridge-gke"]
  }
}
