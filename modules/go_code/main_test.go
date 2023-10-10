package main

// Import key modules.
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
)

var (
	globalBackendConf = make(map[string]interface{})
	globalEnvVars     = make(map[string]string)
	uniquePostfix     = strings.ToLower(random.UniqueId())
	prefix            = "vnet"
	separator         = "-"
)

// Define key global variables.
var (
	subscriptionId      = "2a04288a-8136-4880-b526-c6070e59f004"
	resource_group_name = "example-resources"
)

const (
	apiVersion              = "2019-06-01"
	resourceProvisionStatus = "Succeeded"
)

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

	return globalEnvVars, nil
}
func TestTerraform_azure_virtualNetwork(t *testing.T) {
	t.Parallel()

	setTerraformVariables()

	resource, err := getResourceFromRESTAPI("/subscriptions/2b973e4e-43f9-4abc-bbde-e7c7cb004949/resourceGroups/rg-adf-demo/providers/Microsoft.Storage/storageAccounts/stocsdemostorage")

	if err != nil {
		log.Fatalf("failed to obtain a terraform var output as json: %v", err)
	}

	fmt.Printf("Resource ID: %s\n", *resource.ID)

}

func getResourceFromRESTAPI(out string) (armresources.ClientGetByIDResponse, error) {

	//expected variable
	//expectedVnetName := strings.ToLower(fmt.Sprintf("%s%s%s", prefix, separator, uniquePostfix))

	log.Printf("json output: %s\n", out)

	ctx := context.Background()
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Authentication failure: %+v", err)
	}

	resourceId := fmt.Sprintf("%v", out) //result["resource_name"]
	//resourceId := fmt.Sprintf(resourceIdFormat, subscriptionId, resource_group_name, expectedResource_name)
	// Azure SDK Azure Resource Management clients accept the credential as a parameter
	client, err := armresources.NewClient(subscriptionId, cred, nil)

	resource, err := client.GetByID(ctx, resourceId, apiVersion, nil)
	if err != nil {
		log.Fatalf("failed to obtain a response: %v", err)
	}
	dataType := reflect.TypeOf(resource)

	fmt.Printf("Data Type: %v\n", dataType)

	// Use the retrieved resource
	fmt.Printf("Resource ID: %s\n", *resource.ID)
	fmt.Printf("Resource Name: %s\n", *resource.Name)

	b, _ := json.Marshal(resource)
	// Convert bytes to string.
	sOutput := string(b)
	//fmt.Println(sOutput)

	// Get bytes.
	bytes := []byte(sOutput)

	// Unmarshal JSON to Result struct.
	var result ReponseBase
	json.Unmarshal(bytes, &result)
	fmt.Printf("Result ResourceId:: %s\n", result.ResourceId)
	fmt.Printf("Result SKU Name:: %s\n", result.SKU.Name)

	return resource, err
}
