
variable "sqlserver_name" {
    type = string
    description = "(optional) sqlserver name"
}

variable "sqlserver_resource_group_name" {
    type = string
    description = "(required) sqlserver resource group name"
}

variable "sqlserver_location" {
    type = string
    description = "(required) sqlserver location name"
}

variable "sqlserver_version" {
    type = string
    description = "(required) sqlserver version"
}

variable "sqlserver_administrator_login" {
    type = string
    description = "(required) sqlserver admin login name"
}

variable "sqlserver_administrator_password" {
    type = string
    description = "(required) sqlserver admin login password"
}

variable "tags" {
    type = map(string)
    description = "(required) tags"
    default = {
    environment = "production"
  }
}