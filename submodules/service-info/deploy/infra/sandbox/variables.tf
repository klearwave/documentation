variable "public_ip" {
  type        = string
  description = <<EOT
  Public IP that will host the API server hostname.  This should already be associated with the USG and
  hosted on the USG public interface.
EOT
}

variable "public_domain" {
  type        = string
  description = "Public domain where the 'hostname' will be created."
}

variable "hostname" {
  type        = string
  description = "Hostname, within 'public_domain' that is associated with the 'public_ip'."
}

#
# namecheap authentication information
#
variable "namecheap_user_name" {
  type        = string
  description = "User name associated with the account hosting the DNS records for 'public_hostname'."
  sensitive   = true
}

variable "namecheap_api_user_name" {
  type        = string
  default     = null
  description = "API User name associated with the account hosting the DNS records for 'public_hostname'.  Defaults to 'namecheap_user_name' if not provided."
  sensitive   = true
  nullable    = true
}

variable "namecheap_api_key" {
  type        = string
  description = "API key used for updating DNS records associated with the account hosting the DNS records for 'public_hostname'."
  sensitive   = true
}

#
# ubiquiti usg variables
#
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

variable "usg_site" {
  type        = string
  default     = "Default"
  description = "Site description to use for management of USG."
}

variable "usg_port" {
  type        = number
  description = "Front end port used to forward to the backend service.  This is the port you connect to publically."
  sensitive   = true
}

#
# backend service
#
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
