package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

const (
	fixtures   = "../"
	apiVersion = "2021-11-01"
)

var (
	globalBackendConf   = make(map[string]interface{})
	globalEnvVars       = make(map[string]string)
	uniquePostfix       = strings.ToLower(random.UniqueId())
	expected_name       = "sqlserverdemo-9901"
	expected_location   = "eastus"
	resource_group_name = "rg-tf-deployment"
)
var subscriptionId string

type TestCondition int

const (
	TestConditionEquals   TestCondition = 0
	TestConditionNotEmpty TestCondition = 1
	TestConditionContains TestCondition = 2
)

func Test_automation(t *testing.T) {
	t.Parallel()
	setTerraformVariables()

	terraform.InitAndApply(t, terraformOptions())
	// Defer destroy the infrastructure using the Terraform destroy command
	defer terraform.Destroy(t, terraformOptions())

	id := terraform.Output(t, terraformOptions(), "id")

	reponseData, err := getResourceFromRESTAPI(id, subscriptionId)
	if err != nil {
		log.Fatalf("failed to obtain a terraform var output as json: %v", err)
	}

	fmt.Printf("Resource ID: %s\n", reponseData.ResourceId)

	//Aserts
	testCases := []struct {
		Name      string
		Got       string
		Want      string
		Condition TestCondition
	}{
		{"resource name matching", reponseData.Name, expected_name, TestConditionEquals},
		{"resource location matching", reponseData.Location, expected_location, TestConditionEquals},
		{"FQDN contains resource name", reponseData.Properties.FullyQualifiedDomainName, expected_name, TestConditionContains},
		{"FQDN Matching", reponseData.Properties.FullyQualifiedDomainName, fmt.Sprintf("%s.database.windows.net", expected_name), TestConditionEquals},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			switch tc.Condition {
			case TestConditionEquals:
				assert.Equal(t, tc.Got, tc.Want)
			case TestConditionNotEmpty:
				assert.NotEmpty(t, tc.Got)
			case TestConditionContains:
				assert.Contains(t, tc.Got, tc.Want)
			}
		})
	}
	//t.Run("Output Validation", OutputValidation)
}

func setTerraformVariables() (map[string]string, error) {

	// Getting enVars from environment variables
	ARM_CLIENT_ID := os.Getenv("AZURE_CLIENT_ID")
	ARM_CLIENT_SECRET := os.Getenv("AZURE_CLIENT_SECRET")
	ARM_TENANT_ID := os.Getenv("AZURE_TENANT_ID")
	ARM_SUBSCRIPTION_ID := os.Getenv("AZURE_SUBSCRIPTION_ID")

	// Creating globalEnVars for terraform call through Terratest
	if ARM_CLIENT_ID != "" {
		globalEnvVars["ARM_CLIENT_ID"] = ARM_CLIENT_ID
		globalEnvVars["ARM_CLIENT_SECRET"] = ARM_CLIENT_SECRET
		globalEnvVars["ARM_SUBSCRIPTION_ID"] = ARM_SUBSCRIPTION_ID

		globalEnvVars["ARM_TENANT_ID"] = ARM_TENANT_ID
	}

	// Set the variable value so that it can be used later
	subscriptionId = globalEnvVars["ARM_SUBSCRIPTION_ID"]
	return globalEnvVars, nil
}

func terraformOptions() *terraform.Options {
	return &terraform.Options{
		TerraformDir: fixtures,
		VarFiles:     []string{"test.tfvars"},

		// // Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"sqlserver_name":                   expected_name,
			"sqlserver_resource_group_name":    resource_group_name,
			"sqlserver_location":               expected_location,
			"sqlserver_version":                "12.0",
			"sqlserver_administrator_login":    "sqladmin",
			"sqlserver_administrator_password": "thisIsDog11",
		},

		// // globalvariables for user account
		// EnvVars: globalEnvVars,
		// // Backend values to set when initialziing Terraform
		// BackendConfig: globalBackendConf,
		// // Disable colors in Terraform commands so its easier to parse stdout/stderr
		// NoColor: true,
		// // Reconfigure is required if module deployment and go test pipelines are running in one stage
		// Reconfigure: true,
	}
}

func getResourceFromRESTAPI(resourceId, azure_subscription_id string) (ReponseBase, error) {

	ctx := context.Background()
	cred, err := azidentity.NewDefaultAzureCredential(nil)

	if err != nil {
		log.Fatalf("Authentication failure: %+v", err)
	}

	// Azure SDK Azure Resource Management clients accept the credential as a parameter
	client, err := armresources.NewClient(azure_subscription_id, cred, nil)

	resource, err := client.GetByID(ctx, resourceId, apiVersion, nil)
	if err != nil {
		log.Fatalf("failed to obtain a response: %v", err)
	}

	// Use below line of code to know the return type dataType
	dataType := reflect.TypeOf(resource)

	fmt.Printf("Data Type: %v\n", dataType)

	// Use the retrieved resource
	fmt.Printf("Resource ID: %s\n", *resource.ID)
	fmt.Printf("Resource Name: %s\n", *resource.Name)

	b, _ := json.Marshal(resource)
	// Convert bytes to string.
	sOutput := string(b)
	fmt.Println(sOutput)

	// Get bytes.
	bytes := []byte(sOutput)

	// Unmarshal JSON to Result struct.
	var result ReponseBase
	json.Unmarshal(bytes, &result)
	fmt.Printf("Result ResourceId:: %s\n", result.ResourceId)

	return result, err
}

func OutputValidation(t *testing.T) {
	// testCases := []struct {
	// 	Name      string
	// 	Got       string
	// 	Want      string
	// 	Condition TestCondition
	// }{

	// 	{"resource name", reponseData.ResourceId, expected_name, TestConditionEquals},
	// 	{"resource location", reponseData.Location, expected_location, TestConditionEquals},
	// 	{"FullyQualifiedDomainName", reponseData.Properties.FullyQualifiedDomainName, expected_name, TestConditionContains},
	// 	{"FullyQualifiedDomainName Matching", reponseData.Properties.FullyQualifiedDomainName, fmt.Sprintf("%s.database.windows.net", expected_name), TestConditionEquals},
	// }

	// for _, tc := range testCases {
	// 	t.Run(tc.Name, func(t *testing.T) {
	// 		switch tc.Condition {
	// 		case TestConditionEquals:
	// 			assert.Equal(t, tc.Got, tc.Want)
	// 		case TestConditionNotEmpty:
	// 			assert.NotEmpty(t, tc.Got)
	// 		case TestConditionContains:
	// 			assert.Contains(t, tc.Got, tc.Want)
	// 		}

	// 	})
	// }
}
