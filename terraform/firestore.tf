# Needed for Cloud Firestore
resource "google_project_service" "file" {
  service            = "file.googleapis.com"
  disable_on_destroy = false
  project            = module.generic-project.project_id
}

resource "google_project_service" "firestore" {
  service            = "firestore.googleapis.com"
  disable_on_destroy = false
  project            = module.generic-project.project_id
}
