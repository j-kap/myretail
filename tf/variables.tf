variable "gcp_project_id" {
  type = string
}

variable "gcp_region" {
  type    = string
  default = "us-central1"
}

variable "gcp_firestore_region" {
  type    = string
  default = "us-central"
}

variable "myretail_service_name" {
  type        = string
  default     = "myretail"
  description = "Name of cloud run service to create"
}

variable "myretail_image_name" {
  type        = string
  default     = "myretail"
  description = "Docker image name to deploy to cloud run"
}

variable "myretail_image_tag" {
  type        = string
  default     = "latest"
  description = "Docker image tag to deploy to cloud run"
}

variable "products_url" {
  type        = string
  description = "URL of remote products api"
}
