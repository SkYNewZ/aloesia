# Enable Storage Component API
resource "google_project_service" "storage" {
  service            = "storage-component.googleapis.com"
  disable_on_destroy = false
  project            = module.generic-project.project_id
}

resource "google_storage_bucket" "source_code" {
  project       = module.generic-project.project_id
  name          = format("%s-functions-source-code", module.generic-project.project_id)
  location      = var.region
  force_destroy = true

  depends_on = [
    google_project_service.storage
  ]
}

data "archive_file" "send_email" {
  type        = "zip"
  output_path = format("%s/../bin/send_email.zip", path.module)

  source {
    content  = file(format("%s/../functions/send_mail/main.py", path.module))
    filename = "main.py"
  }
}

# Copy source code to this bucket
resource "google_storage_bucket_object" "send_email" {
  name   = format("send_email_%s.zip", data.archive_file.send_email.output_md5)
  bucket = google_storage_bucket.source_code.name
  source = data.archive_file.send_email.output_path
}
