module "pubsub_send_email" {
  source     = "terraform-google-modules/pubsub/google"
  version    = "~> 1.0"
  project_id = module.generic-project.project_id
  topic      = "send_email"
}
