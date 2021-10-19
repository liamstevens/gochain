provider "google" {
  project     = var.project_id
  region      = var.region
}

terraform {
  backend "gcs" {
    bucket  = var.state_bucket
    prefix  = "terraform/state"
  }
}