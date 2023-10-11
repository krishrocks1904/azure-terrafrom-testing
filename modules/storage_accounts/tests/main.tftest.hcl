variables {
    resource_group_name = "rg-tf-deployment"
    location            = "eastus"
    settings            = {}
    name                = "stourcloudschool9"
}
# terraform {
#   required_providers {
#     azurerm = {
#       source  = "hashicorp/azurerm"
#       version = "=3.0.0"
#     }
#   }
# }
provider "azurerm" {
  features {}
}
run "unit_tests" {
    command = plan

    variables {
        resource_group_name = var.resource_group_name
        location            = var.location
        settings            = var.settings
        name                = var.name
    }

    assert {
        condition =  azurerm_storage_account.stg.account_kind == "StorageV2"
        error_message = "storage kind is not of type StorageV2"
    }
}

run "validate_storage_account_name" {
    command = plan

    variables {
        resource_group_name = var.resource_group_name
        location            = var.location
        settings            = var.settings
        name                = var.name
    }

    assert {
        condition =  azurerm_storage_account.stg.name == var.name
        error_message = "storage account name must match"
    }
}

run "create_storage_account_and_validate_name" {
    command = apply

    variables {
        resource_group_name = var.resource_group_name
        location            = var.location
        settings            = var.settings
        name                = var.name
    }

    assert {
        condition =  azurerm_storage_account.stg.name == var.name
        error_message = "storage account name must match"
    }
}