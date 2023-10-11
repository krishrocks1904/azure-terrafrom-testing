
 the name of test file (golang file) must be contains keyword 'test'
Package Name must be same 'response_type.go' and 'main_test.go' file

In the test we are using Azure SDK package which will call to Azure REST API
 - The REST API will require to login using the service principal this requires some environment variables to be setup

```  
$env:AZURE_CLIENT_ID="Azure App registration client Id (AppId)"
$env:AZURE_SUBSCRIPTION_ID="Azure Subscription Id"
$env:AZURE_TENANT_ID="Azure tenant Id"
$env:AZURE_CLIENT_SECRET="Azure Client Secret"
```

Azure module must have features block to run the code from test, add below line if that is missing

```
provider "azurerm" {
  features {}
}

```

# Run Test
- Run following golang command to run the test cases.

```
go mod init testmod && go mod tidy
```

# Run test command

```
go test -timeout 30m -v
```