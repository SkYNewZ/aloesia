resource "google_pubsub_topic" "topic" {
  name    = "send_email"
  project = module.generic-project.project_id
}
