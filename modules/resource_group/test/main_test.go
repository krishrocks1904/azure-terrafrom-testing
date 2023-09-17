package test

import (
	"testing"
	//"os"
	// "path/filepath"
	// "fmt"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerraformAzureInfrastructure(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		// Set the path to your Terraform code that will be tested
		TerraformDir: "../",
		
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"resource_group_name": "rg-demo-001", //os.Getenv("TF_VAR_resource_group_name"),
			"location":            "eastus",//os.Getenv("TF_VAR_location"),
		},
	}

	// Defer destroy the infrastructure using the Terraform destroy command
	defer terraform.Destroy(t, terraformOptions)

	// Run terraform init and apply
	terraform.InitAndApply(t, terraformOptions)

	// Validate that the resource group was created
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	require.Equal(t, "rg-demo-001", resourceGroupName)
	assert.Equal(t, "rg-demo-001", resourceGroupName)

	// You can add more assertions here based on the resources you're creating
}

