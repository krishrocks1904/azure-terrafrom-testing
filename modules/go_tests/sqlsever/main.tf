
provider "azurerm" {
  features {}
}

resource "azurerm_mssql_server" "sql" {
  name                         = var.sqlserver_name
  resource_group_name          = var.sqlserver_resource_group_name
  location                     = var.sqlserver_location
  version                      = var.sqlserver_version
  administrator_login          = var.sqlserver_administrator_login
  administrator_login_password = var.sqlserver_administrator_password
  minimum_tls_version          = "1.2"
  tags = var.tags
}
