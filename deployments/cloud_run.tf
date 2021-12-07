resource "google_cloud_run_service" "app" {
  project  = var.project.id
  name     = "oehs-app"
  location = var.region

  template {
    spec {
      containers {
        image = "${var.container_registry_url}/oehs:${file("app_version.txt")}"
      }
    }
  }

  metadata {
    annotations = {
      "run.googleapis.com/ingress"     = "all"
      "run.googleapis.com/client-name" = "terraform"
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
  autogenerate_revision_name = true
  depends_on                 = [
    google_project_service.gcp_services,
    google_cloudbuild_trigger.push,
    google_artifact_registry_repository.artifact-registry-repo
  ]
}
