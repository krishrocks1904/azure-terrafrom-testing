###### IMPORTANT LEGAL NOTICE: PLEASE BE AWARE THAT THIS CODE AND/OR SCRIPT (SOFTWARE) IS THE EXCLUSIVE  AND PROTECTED.  #
###### IT IS A CRIMINAL OFFENCE TO COPY OR RE-PRODUCE IN WHOLE OR PART ANY ELEMENT OF THIS SOFTWARE.                                                  #
######                                                                                                                                                #
###### Company Confidential                                                                                                                           #
######                                                                                                                                                #
###### Original Author: Rakesh Suryawanshi                                                                                                            #
###### Creation Date: 09/06/2023                                                                                                                      #
###### Description: Terraform Infrastrucutre Creation                                                                                                 #
###### Pre-requisites: GitHub Actions, Azure Service Principal                                                                                        #
###################################################################################################################################################

output "id" {
  value       = azurerm_storage_account.stg.id
  description = "id - The ID of the Storage Account."
}

output "name" {
  value       = azurerm_storage_account.stg.name
  description = "Storage Account name."
}

output "primary_location" {
  value       = azurerm_storage_account.stg.primary_location
  description = "primary_location - The primary location of the storage account."
}

output "primary_blob_endpoint" {
  value       = azurerm_storage_account.stg.primary_blob_endpoint
  description = "primary_blob_endpoint - The endpoint URL for blob storage in the primary location."
}

output "primary_queue_endpoint" {
  value       = azurerm_storage_account.stg.primary_queue_endpoint
  description = "The endpoint URL for queue storage in the primary location."
}

output "primary_table_endpoint" {
  value       = azurerm_storage_account.stg.primary_table_endpoint
  description = "primary_table_endpoint - The endpoint URL for table storage in the primary location."
}

output "primary_file_endpoint" {
  value       = azurerm_storage_account.stg.primary_file_endpoint
  description = "primary_file_endpoint - The endpoint URL for file storage in the primary location."
}

output "primary_access_key" {
  value       = azurerm_storage_account.stg.primary_access_key
  description = "The primary access key for the storage account."
  sensitive = true
}

output "primary_connection_string" {
  value       = azurerm_storage_account.stg.primary_connection_string
  description = "The connection string associated with the primary location."
  sensitive   = true
}
