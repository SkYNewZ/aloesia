resource "google_project_service" "iam" {
  service            = "iam.googleapis.com"
  disable_on_destroy = false
  project            = module.generic-project.project_id
}

# Service account to deploy code to AppEngine
resource "google_service_account" "deployer" {
  project      = module.generic-project.project_id
  account_id   = "deployer"
  display_name = "Deploy from CI"
  depends_on   = [google_project_service.iam]
}

resource "google_project_iam_member" "appengine_admin" {
  project = module.generic-project.project_id
  role    = "roles/appengine.appAdmin"
  member  = format("serviceAccount:%s", google_service_account.deployer.email)
}

resource "google_project_iam_member" "cloudbuild_builds_editor" {
  project = module.generic-project.project_id
  role    = "roles/cloudbuild.builds.editor"
  member  = format("serviceAccount:%s", google_service_account.deployer.email)
}

resource "google_project_iam_member" "storage_admin" {
  project = module.generic-project.project_id
  role    = "roles/storage.admin"
  member  = format("serviceAccount:%s", google_service_account.deployer.email)
}
