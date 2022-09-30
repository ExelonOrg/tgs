variable "identifier" {
  type = object({
      primary = string
      secondary = string
      type = string
  })
}

variable "tags" {
  type = object({})
}

variable "service_endpoint" {
  type = string
}

variable "backends" {
  type = list
  default = []
}

variable "revision" {
  type = string
  default = "1"
}

variable "apim_instance" {
  type = object({
      name = string
      resource_group_name = string
  })
}

variable "operations" {
  type = list
  default = ["GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH", "HEAD", "TRACE"]
}

variable "add_product" {
  type = bool
  default = false
}

variable "path_override" {
  type = bool
  default = false
}

variable "service_url_override" {
  type = bool
  default = false
}

variable "service_url" {
  type = string
  default = ""
}

variable "policy_override" {
  type = bool
  default = false
}

variable "path" {
  type = string
  default = ""
}

variable "soap_pass_through" {
  type = bool
  default = false
}

variable "subscription_required" {
  type = bool
  default = false
}

variable "wsdl_content_value" {
  type = string
  default = ""
}

variable "wsdl_service_name" {
  type = string
  default = ""
}

variable "wsdl_endpoint_name" {
  type = string
  default = ""
}

variable "api_tags" {
  type = list
  default = []
}

variable "policy_xml_content" {
  type = string
}

variable "named_values" {
  type = list
  default = []
}

variable "authorization_server_name" {
  type = list
  default = []
}