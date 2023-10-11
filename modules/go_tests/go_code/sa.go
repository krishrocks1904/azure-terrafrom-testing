package main

// Import key modules.

type NetworkAcls struct {
	Bypass              string   `json:"bypass"`
	DefaultAction       string   `json:"defaultAction"`
	VirtualNetworkRules []string `json:"virtualNetworkRules"`
	IpRules             []string `json:"ipRules"`
}
type Properties struct {
	AccessTier            string `json:"accessTier"`
	AllowBlobPublicAccess bool   `json:"allowBlobPublicAccess"`
	AllowSharedKeyAccess  bool   `json:"allowSharedKeyAccess"`
	MinimumTlsVersion     string `json:"minimumTlsVersion"`

	NetworkAcls NetworkAcls `json:"networkAcls"`
	//ProvisioningState string      `json:"provisioningState"` //  "Succeeded",

	//primaryEndpoints
}

type SKU struct {
	Name string `json:"name"`
	Tier string `json:"tier"`
}
type ReponseBase struct {
	ResourceId string     `json:"id"`
	Location   string     `json:"location"`
	Kind       string     `json:"kind"`
	Name       string     `json:"name"`
	Properties Properties `json:"properties"`

	SKU  SKU    `json:"sku"`
	Type string `json:"type"` // "Microsoft.Storage/storageAccounts"
}

// "primaryEndpoints": {
// 	"blob": "https://stocsdemostorage.blob.core.windows.net/",
// 	"dfs": "https://stocsdemostorage.dfs.core.windows.net/",
// 	"file": "https://stocsdemostorage.file.core.windows.net/",
// 	"queue": "https://stocsdemostorage.queue.core.windows.net/",
// 	"table": "https://stocsdemostorage.table.core.windows.net/",
// 	"web": "https://stocsdemostorage.z13.web.core.windows.net/"
// },
