resource "google_artifact_registry_repository" "artifact-registry-repo" {
  provider      = google-beta
  project       = var.project.id
  location      = var.region
  repository_id = var.artifact_registry.repository
  description   = "Artifact Registry Repository"
  format        = "DOCKER"

  depends_on = [
    google_project_service.gcp_services
  ]
}
