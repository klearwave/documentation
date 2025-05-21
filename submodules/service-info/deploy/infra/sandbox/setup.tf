#
# public dns
#
resource "namecheap_domain_records" "service_info" {
  domain = var.public_domain

  record {
    hostname = var.hostname
    type     = "A"
    address  = var.public_ip
  }
}

#
# ubiquiti: get site name from proper name
#
locals {
  login_data = jsonencode(
    {
      username = var.usg_username,
      password = var.usg_password,
    }
  )
}

data "http" "login" {
  url      = "${var.usg_api_url}/api/auth/login"
  method   = "POST"
  insecure = true

  request_body = local.login_data

  request_headers = {
    "Content-Type" = "application/json"
  }

  lifecycle {
    postcondition {
      condition     = contains([200], self.status_code)
      error_message = "Could not login to Ubiquiti API.  Status code [${self.status_code}] invalid.  Expected 200."
    }

    postcondition {
      condition     = coalesce(lookup(self.response_headers, "Set-Cookie", null), null) != null
      error_message = "Could not login to Ubiquiti API.  Set-Cookie missing from response headers."
    }

    postcondition {
      condition     = coalesce(lookup(self.response_headers, "X-Csrf-Token", null), null) != null
      error_message = "Could not login to Ubiquiti API.  X-Csrf-Token missing from response headers."
    }
  }
}

locals {
  auth_token = regex("TOKEN=([^;]+)", data.http.login.response_headers["Set-Cookie"])[0]
}

data "http" "current_site_config" {
  url      = "${var.usg_api_url}/proxy/network/api/self/sites"
  method   = "GET"
  insecure = true

  request_headers = {
    "Content-Type" = "application/json",
    "Cookie"       = "TOKEN=${local.auth_token}"
  }

  lifecycle {
    postcondition {
      condition     = contains([200], self.status_code)
      error_message = "Could not GET site config.  Status code [${self.status_code}] invalid.  Expected 200."
    }
  }
}

locals {
  site_name = [
    for site in jsondecode(data.http.current_site_config.response_body).data : site.name if site.desc == var.usg_site
  ][0]
}

#
# ubiquiti: local port forward
#
resource "unifi_port_forward" "service_info" {
  site = local.site_name

  name = "klearwave_service_info"

  protocol = "tcp"
  dst_port = var.usg_port
  fwd_ip   = var.backend_service_address
  fwd_port = var.backend_service_port

  port_forward_interface = "wan"
}

#
# NOTE: firewall rule seems to be automatically created for port forward
#       but here may be an example of how to do it manually if the above 
#       fails for some reason.
#

# resource "unifi_firewall_group" "local_usg_ports" {
#   site = local.site_name
#
#   name    = "port_usg_klearwave_service_info"
#   type    = "port-group"
#   members = [tostring(var.usg_port), tostring(var.backend_service_port)]
# }
#
# resource "unifi_firewall_group" "local_usg_addresses" {
#   site = local.site_name
#
#   name    = "address_local_usg_klearwave_service_info"
#   type    = "address-group"
#   members = [var.public_ip, var.backend_service_address]
# }
#
# resource "unifi_firewall_rule" "service_info_wan_local" {
#   site = local.site_name
#
#   name       = "allow_klearwave_service_info_local_${var.usg_port}"
#   action     = "accept"
#   ruleset    = "WAN_LOCAL"
#   enabled    = true
#   rule_index = 2001
#   protocol   = "tcp"
#
#   dst_firewall_group_ids = [
#     unifi_firewall_group.local_usg_addresses.id,
#     unifi_firewall_group.local_usg_ports.id
#   ]
# }
#
# resource "unifi_firewall_rule" "service_info_wan_in" {
#   site = local.site_name
#
#   name       = "allow_klearwave_service_info_in_${var.usg_port}"
#   action     = "accept"
#   ruleset    = "WAN_IN"
#   enabled    = true
#   rule_index = 2002
#   protocol   = "tcp"
#
#   dst_firewall_group_ids = [
#     unifi_firewall_group.local_usg_addresses.id,
#     unifi_firewall_group.local_usg_ports.id
#   ]
# }
