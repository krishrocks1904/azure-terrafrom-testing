package test
//https://brendanthompson.com/posts/2021/09/getting-started-with-terratest-on-azure
import (
	"testing"
	//"os"
	// "path/filepath"
	 "fmt"
	
    "strings"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	//"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

const (
    fixtures = "../"
)

var (
	globalBackendConf = make(map[string]interface{})
	globalEnvVars     = make(map[string]string)
	uniquePostfix     = strings.ToLower(random.UniqueId())
	expected_name              = "stourcloudschool9"	
)

// func setTerraformVariables() (map[string]string, error) {

// 	// Getting enVars from environment variables
// 	ARM_CLIENT_ID := os.Getenv("AZURE_CLIENT_ID")
// 	ARM_CLIENT_SECRET := os.Getenv("AZURE_CLIENT_SECRET")
// 	ARM_TENANT_ID := os.Getenv("AZURE_TENANT_ID")
// 	ARM_SUBSCRIPTION_ID := os.Getenv("AZURE_SUBSCRIPTION_ID")

// 	// Creating globalEnVars for terraform call through Terratest
// 	if ARM_CLIENT_ID != "" {
// 		globalEnvVars["ARM_CLIENT_ID"] = ARM_CLIENT_ID
// 		globalEnvVars["ARM_CLIENT_SECRET"] = ARM_CLIENT_SECRET
// 		globalEnvVars["ARM_SUBSCRIPTION_ID"] = ARM_SUBSCRIPTION_ID
// 		globalEnvVars["ARM_TENANT_ID"] = ARM_TENANT_ID
// 	}

// 	return globalEnvVars, nil
// }
type TestCondition int
const (
    TestConditionEquals   TestCondition = 0
    TestConditionNotEmpty TestCondition = 1
	TestConditionContains TestCondition = 2
)
func terraformOptions() *terraform.Options {
    return &terraform.Options{
        TerraformDir: fixtures,
		VarFiles: []string{"test.tfvars"},
		// // Variables to pass to our Terraform code using -var options
		// Vars: map[string]interface{}{
		// 	"resource_group_name": resource_group_name, //os.Getenv("TF_VAR_resource_group_name"),
		// 	"location":            location,//os.Getenv("TF_VAR_location"),
		// 	"name":name,
		// 	"settings":settings,
		// },

		// // globalvariables for user account
		// EnvVars: globalEnvVars,
		// // Backend values to set when initialziing Terraform
		// BackendConfig: globalBackendConf,
		// // Disable colors in Terraform commands so its easier to parse stdout/stderr
		// NoColor: true,
		// // Reconfigure is required if module deployment and go test pipelines are running in one stage
		// Reconfigure: true,
        //NoColor:      true,

    }
}

func Test_automation(t *testing.T) {
	t.Parallel()
	
	//  // Set the Azure subscription ID and resource group where the Storage Account will be created
	//  uniqueID := random.UniqueId()
	//  subscriptionID := "YOUR_AZURE_SUBSCRIPTION_ID"
	//  resourceGroupName := fmt.Sprintf("test-rg-%s", uniqueID)

	terraform.InitAndApplyAndIdempotent(t, terraformOptions())
	// Defer destroy the infrastructure using the Terraform destroy command
	defer terraform.Destroy(t, terraformOptions())


	t.Run("Output Validation", OutputValidation)

	// // Validate that the resource group was created
	// name := terraform.Output(t, terraformOptions, "name")
	// primary_location := terraform.Output(t, terraformOptions, "primary_location")
	// id := terraform.Output(t, terraformOptions, "id")
	// require.Equal(t, name, name)
	// assert.Equal(t, name, name)
	// assert.NotEmpty(t, primary_location)
	// assert.Contains(t, id, name)
}

func OutputValidation(t *testing.T) {
    testCases := []struct {
        Name      string
        Got       string
        Want      string
        Condition TestCondition
    }{
        {"Storage account name", terraform.Output(t, terraformOptions(), "name"),expected_name, TestConditionEquals},
        {"primary location", terraform.Output(t, terraformOptions(), "primary_location"), "", TestConditionNotEmpty},
		{"resource Id container Storage account name", terraform.Output(t, terraformOptions(), "id"), "", TestConditionNotEmpty},
		{"primary_blob_endpoint", terraform.Output(t, terraformOptions(), "primary_blob_endpoint"),fmt.Sprintf("https://%s.blob.core.windows.net/", expected_name), TestConditionEquals},
    }

    for _, tc := range testCases {
        t.Run(tc.Name, func(t *testing.T) {
            switch tc.Condition {
            case TestConditionEquals:
                assert.Equal(t, tc.Got, tc.Want)
            case TestConditionNotEmpty:
                assert.NotEmpty(t, tc.Got)
			case TestConditionContains:
                assert.Contains(t, tc.Got,tc.Want)
            }

        })
    }
}
