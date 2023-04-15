
## https://cloud.google.com/run/docs/deploying?hl=ja#terraform
## https://cloud.google.com/docs/terraform/resource-management/managing-infrastructure-as-code?hl=ja
## image : create in cloudbuild.yaml


terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "4.63.1"
    }
  }
}

provider "google" {
  # Configuration options
  project = var.project
  region  = var.region 
}

resource "google_cloud_run_service" "default" {
  name     = "finmng-service-api"
  location = "asia-northeast1"

  template {
    spec {
      containers {
        image = var.image_id
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

data "google_iam_policy" "noauth" {
   binding {
     role = "roles/run.invoker"
     members = ["allUsers"]
   }
 }

 resource "google_cloud_run_service_iam_policy" "noauth" {
   location    = google_cloud_run_service.default.location
   project     = google_cloud_run_service.default.project
   service     = google_cloud_run_service.default.name

   policy_data = data.google_iam_policy.noauth.policy_data
}