
resource "azurerm_storage_account" "st" {
  name                     = var.resource_name
  resource_group_name      = var.resource_group_name
  location                 = var.location
  account_tier             = var.account_tier //"Standard"
  account_replication_type = var.account_replication_type //"GRS"

  tags = {
    environment = "staging"
  }
}

