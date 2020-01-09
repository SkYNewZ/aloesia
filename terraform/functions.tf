# Enable Cloud Function API
resource "google_project_service" "cloudfunctions" {
  service            = "cloudfunctions.googleapis.com"
  disable_on_destroy = false
  project            = module.generic-project.project_id
}

# Send emails function
resource "google_cloudfunctions_function" "send_email" {
  project               = module.generic-project.project_id
  name                  = "send_email"
  region                = var.region
  description           = "Send email from a topic message"
  runtime               = "python37"
  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.source_code.name
  source_archive_object = google_storage_bucket_object.send_email.name
  entry_point           = "send_mail"
  environment_variables = var.send_email_environment_variables

  event_trigger {
    event_type = "providers/cloud.pubsub/eventTypes/topic.publish"
    resource   = google_pubsub_topic.topic.id

    failure_policy {
      retry = false
    }
  }

  depends_on = [
    google_project_service.cloudfunctions
  ]
}
