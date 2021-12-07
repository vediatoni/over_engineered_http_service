resource "google_cloudbuild_trigger" "push" {
  #   provider = google-beta

  name        = "push"
  description = "On push pipeline"
  filename    = "build/cloudbuild_on_push.yaml"
  project     = var.project.id


  github {
    owner = "vediatoni"
    name  = "over_engineered_http_service"
    push {
      branch = ".*"
    }
  }


  depends_on = [
    google_project_service.gcp_services
  ]
}

data "google_project" "project" {
  project_id = var.project.id
}

resource "google_project_iam_member" "cloud-svcacc-role-binding-svc" {
  project    = var.project.id
  role       = "roles/iam.serviceAccountUser"
  member     = "serviceAccount:${data.google_project.project.number}@cloudbuild.gserviceaccount.com"
  depends_on = [google_project_service.gcp_services, google_cloudbuild_trigger.push]
}

resource "google_project_iam_member" "cloud-svcacc-role-binding-run" {
  project    = var.project.id
  role       = "roles/run.admin"
  member     = "serviceAccount:${data.google_project.project.number}@cloudbuild.gserviceaccount.com"
  depends_on = [google_project_service.gcp_services, google_cloudbuild_trigger.push]
}

