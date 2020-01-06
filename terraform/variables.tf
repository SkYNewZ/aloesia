variable "project_id" {
  type        = string
  description = "Project it to set to"
}

variable "region" {
  type        = string
  description = "Where to deploy resources"
  default     = "europe-west1"
}

variable "send_email_environment_variables" {
  type = map
  description = "Map of key/values HCL pairs to set into mail Cloud Function"
}