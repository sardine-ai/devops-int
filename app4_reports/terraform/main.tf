resource "google_service_account" "ci" {
  account_id = "ci-deployer"
}

resource "google_project_iam_member" "ci" {
  project = "acme-prod"
  role    = "roles/editor"
  member  = "serviceAccount:${google_service_account.ci.email}"
}

resource "google_service_account_key" "ci" {
  service_account_id = google_service_account.ci.name
}

resource "google_storage_bucket_iam_member" "reports" {
  bucket = "acme-prod-reports"
  role   = "roles/storage.objectViewer"
  member = "allUsers"
}
