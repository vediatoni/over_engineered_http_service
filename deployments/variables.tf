variable "credentials_file" {}

variable "region" {
  default = "us-central1"
}

variable "zone" {
  default = "us-central1-c"
}

variable "project" {
  type    = object({
    name = string
    id   = string
  })
  default = {
    name = "Over engineered http service"
    id   = "over-engineered-http-service"
  }
}

variable "artifact_registry" {
  type    = object({
    repository = string
    location   = string
  })
  default = {
    repository = "over-engineered-http-service"
    location   = "us-central1-docker.pkg.dev"
  }
}

variable "gcp_service_list" {
  description = "The list of all GCP API's necessary for this project"
  type        = list(string)
  default     = [
    "cloudresourcemanager.googleapis.com",
    "iam.googleapis.com",
    "cloudbuild.googleapis.com",
    "run.googleapis.com",
    "artifactregistry.googleapis.com",
  ]
}

variable "organization" {
  default = ""
}