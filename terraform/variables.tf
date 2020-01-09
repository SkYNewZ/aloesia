variable "project_id" {
  type        = string
  description = "Project it to set to"
}

variable "billing_account" {
  type    = string
  default = "019C2D-785389-9E0921"
}

variable "region" {
  type        = string
  description = "Where to deploy resources"
  default     = "europe-west1"
}

variable "send_email_environment_variables" {
  type        = map
  description = "Map of key/values HCL pairs to set into mail Cloud Function"
}
