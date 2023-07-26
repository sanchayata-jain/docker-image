provider "google" {
  project     = local.project
  region      = local.region
}

# modules can be used to split up different resources, 
# in this case artifact registry (where we push our docker image to) and cloud run (deploys the image)
module "artifact_registry" {
  source = "./modules/artifact_registry"

  project = local.project
  region  = local.region
}

module "cloud_run" {
  source       = "./modules/cloud_run"
  service_name = "example-api"

  region           = local.region
  google_project   = local.project
}

#  a locals block can be used to store variables within the terraform
locals {
  project  = "docker-tutorial"
  region   = "europe-west2"
}