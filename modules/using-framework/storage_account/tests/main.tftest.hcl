variables {

 resource_name ="stdemoocs9901"
 resource_group_name = "rg-storage"
 location = "eastus"
 account_tier  = "Standard"
 account_replication_type ="LRS"

}
provider "azurerm" {
  features { }
}

run setup {
    command = apply
    
    variables {
        resource_group_name = "rg-storage"
        location            = "eastus"
    }

    module {
        // NOTE: The Path must be like this, starting from tests directory
        source              = "./tests/setup"
    }
}

run storage_account_variable_validation {

  command = plan
  
  variables {
    resource_name               = "stdemoocs9901"
    resource_group_name         = "rg-storage"
    location                    = "eastus"
    account_tier                = "Standard"
    account_replication_type    = "OKAY"
  }

  expect_failures  = [
    var.account_replication_type,
  ]
  
}


run storage_account_attribute_actual_vs_expected_test {

  command = plan
  
  variables {
    resource_group_name         = run.setup.resource_group_name
    location                    = var.location
    resource_name               = var.resource_name
    account_tier                = var.account_tier
    account_replication_type    = var.account_replication_type
  }

  assert {
     condition      = azurerm_storage_account.st.name == var.resource_name
     error_message  = "storage account name is not matching with given variable value"
  }

  assert {
     condition      = azurerm_storage_account.st.location == var.location
     error_message  = "storage account location is not matching with given variable value"
  }

  assert {
     condition      = azurerm_storage_account.st.account_tier == var.account_tier
     error_message  = "storage account account_tier is not matching with given variable value"
  }

  assert {
     condition      = azurerm_storage_account.st.account_replication_type == var.account_replication_type
     error_message  = "storage account location is not matching with given variable value"
  }  
}

// assert {
//     condition     = azurerm_storage_account.st == ""
//     error_message = "Invalid bucket name"
//   }