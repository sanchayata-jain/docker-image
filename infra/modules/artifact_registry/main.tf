resource "google_artifact_registry_repository" "artifact_registry" {
  provider = google-beta
  project = var.project
  location  = var.region

  repository_id = "example"
  format = "DOCKER"
}