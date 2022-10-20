terraform {
  backend "gcs" {
    bucket  = "gh-pipeline-tf-state"
    prefix  = "terraform/state/glennadjrussell"
  }
}

