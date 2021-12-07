terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.2.1"
    }
  }
}

provider "google" {
  credentials = file(var.credentials_file)
  project = var.project.name
  region  = var.region
  zone    = var.zone
}

provider "google-beta" {
  credentials = file(var.credentials_file)
  project = var.project.name
  region  = var.region
  zone    = var.zone
}

resource "google_project_service" "gcp_services" {
  for_each                   = toset(var.gcp_service_list)
  project                    = var.project.id
  service                    = each.key
  disable_dependent_services = true
}

