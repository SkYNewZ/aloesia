variable "project_id" {
  type        = string
  description = "GCP project ID"
}

variable "region" {
  type    = string
  default = "europe-west2"
}

variable "labels" {
  type        = map
  description = "Labels map to apply to resources"
  default = {
    creator = "skynewz"
  }
}

variable "send_email_environment_variables" {
  type        = map
  description = "Map of environments variables to apply to send_email Cloud Function"
}
