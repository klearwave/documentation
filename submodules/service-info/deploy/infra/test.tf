#
# variables
#
variable "public_ip" {
  type        = string
  description = <<EOT
  Public IP that will host the API server hostname.  This should already be associated with the USG and
  hosted on the USG public interface.
EOT
}

variable "hostname" {
  type        = string
  description = "Hostname, within 'public_domain' that is associated with the 'public_ip'."
}

variable "namecheap_user_name" {
  type        = string
  description = "User name associated with the account hosting the DNS records for 'public_hostname'."
  sensitive   = true
}

variable "namecheap_api_key" {
  type        = string
  description = "API key used for updating DNS records associated with the account hosting the DNS records for 'public_hostname'."
  sensitive   = true
}

variable "usg_username" {
  type        = string
  default     = "admin"
  description = "Username used for connecting to the USG for management."
  sensitive   = true
}

variable "usg_password" {
  type        = string
  description = "Password used for the 'usg_username' used for connecting to the USG for management."
  sensitive   = true
}

variable "usg_api_url" {
  type        = string
  description = "URL used for making API requests against the USG for management."
  sensitive   = true
}

variable "usg_port" {
  type        = number
  description = "Front end port used to forward to the backend service.  This is the port you connect to publically."
  sensitive   = true
}

variable "backend_service_address" {
  type        = string
  description = "Backend service IP to forward from the 'public_ip':'usg_port' from."
  sensitive   = true
}

variable "backend_service_port" {
  type        = number
  description = "Backend service port to forward from the 'public_ip':'usg_port' from."
  default     = 8888
}

locals {
  base_domain = "klearwave.io"
}

#
# test environments
# 
module "test_sandbox" {
  source = "./sandbox"

  public_ip     = var.public_ip
  public_domain = local.base_domain
  hostname      = var.hostname

  namecheap_user_name = var.namecheap_user_name
  namecheap_api_key   = var.namecheap_api_key

  usg_api_url             = var.usg_api_url
  usg_username            = var.usg_username
  usg_password            = var.usg_password
  usg_site                = "scott-home"
  usg_port                = var.usg_port
  backend_service_address = var.backend_service_address
  backend_service_port    = var.backend_service_port
}
