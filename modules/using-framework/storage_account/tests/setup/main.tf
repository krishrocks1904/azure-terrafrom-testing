
variable resource_group_name {
  type = string
  description = "name of resource group of the storage account"
}

variable location {
  type = string
  description = "name of the storage account location"
}


resource "azurerm_resource_group" "example" {
  name     = var.resource_group_name
  location = var.location
}

output "resource_group_name" {
  value = azurerm_resource_group.example.name
}

output "resource_group_id" {
  value =  azurerm_resource_group.example.id
}