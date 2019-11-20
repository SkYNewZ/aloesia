module "pubsub_send_email" {
  source       = "terraform-google-modules/pubsub/google"
  version      = "~> 1.0"
  project_id   = var.project_id
  topic        = "send_email"
  topic_labels = var.labels
}
