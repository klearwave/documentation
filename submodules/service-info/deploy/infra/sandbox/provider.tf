terraform {
  required_providers {
    namecheap = {
      source  = "namecheap/namecheap"
      version = ">= 2.2.0"
    }

    unifi = {
      source  = "paultyng/unifi"
      version = "0.41.0"
    }

    http = {
      source  = "hashicorp/http"
      version = "3.4.5"
    }
  }
}

locals {
  namecheap_api_user_name = (var.namecheap_api_user_name == null || var.namecheap_api_user_name == "") ? var.namecheap_user_name : var.namecheap_api_user_name
}

provider "namecheap" {
  user_name = var.namecheap_user_name
  api_user  = local.namecheap_api_user_name
  api_key   = var.namecheap_api_key
}

provider "unifi" {
  username = var.usg_username
  password = var.usg_password
  api_url  = var.usg_api_url

  allow_insecure = true
}
