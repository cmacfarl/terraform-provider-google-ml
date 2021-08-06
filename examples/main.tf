terraform {
    required_providers {
        googleml = {
            version = "0.1.0"
            source = "cmacfarl.org/mlapi/google-ml"
        }
    }
}

provider "googleml" {
    project = "<your-project-id>"
    credentials = "<your-service-account-credentials-file>"
}

data "ml_config" "cfg" {
    provider = googleml
}

output "tpu_service_account" {
    value = data.ml_config.cfg.tpu_service_account
}
