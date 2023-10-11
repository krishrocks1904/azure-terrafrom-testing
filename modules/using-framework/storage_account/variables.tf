variable resource_name {
  type = string
  description = "name of the storage account"
}

variable resource_group_name {
  type = string
  description = "name of resource group of the storage account"
}

variable location {
  type = string
  description = "name of the storage account location"
}

variable account_tier {
  type = string
  description = "storage account tier value"
  validation {
    condition     = contains(["Standard", "Premium"], var.account_tier)
    error_message = "storage account tier value must be 'Standard' or 'Premium'"
  }
}

variable account_replication_type {
  type = string
  description = "storage account replication type"
  validation {
    condition     = contains(["LRS" ,"GRS"], var.account_replication_type)
    error_message = "the only allowed value for storage account tier on this organization are 'LRS' , 'GRS'"
  }
}