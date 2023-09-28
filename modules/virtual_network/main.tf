provider "azurerm" {
  features {}
}
data "azurerm_client_config" "current" {}

resource "azurerm_virtual_network" "example" {
  name                = var.resource_name
  location            = var.location
  resource_group_name = var.resource_group_name
  address_space       = ["10.0.0.0/16"]
  dns_servers         = ["10.0.0.4", "10.0.0.5"]

  subnet {
    name           = "subnet1"
    address_prefix = "10.0.1.0/24"
  }

  subnet {
    name           = "subnet2"
    address_prefix = "10.0.2.0/24"
    //security_group = azurerm_network_security_group.example.id
  }

  tags = {
    environment = "Production"
  }
}