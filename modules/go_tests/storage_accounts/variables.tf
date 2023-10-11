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

variable "name" {
  type        = string
  description = "(Required) The name which should be used for this Service Plan. Changing this forces a new AppService to be created."
}

variable "resource_group_name" {
  type        = string
  description = "(Required) The name of the Resource Group where the AppService should exist. Changing this forces a new AppService to be created."
}

variable "location" {
  type        = string
  description = "(Required) The Azure Region where the Service Plan should exist. Changing this forces a new AppService to be created"
}

variable "tags" {
  type        = map(string)
  description = " (Optional) A mapping of tags which should be assigned to the AppService."
  default     = {}
}

variable "settings" {

}


variable "subnets" {

  default     = {}
  description = "colleciton of  the subnet."
}

variable "aad_groups" {

  default     = {}
  description = "collection of aad apps."
}


