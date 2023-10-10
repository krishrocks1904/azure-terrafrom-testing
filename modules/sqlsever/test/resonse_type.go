package main

// Import key modules.

type Properties struct {
	AdministratorLogin            string `json:"administratorLogin"`
	FullyQualifiedDomainName      string `json:"fullyQualifiedDomainName"`
	MinimalTlsVersion             string `json:"minimalTlsVersion"`
	PublicNetworkAccess           string `json:"publicNetworkAccess"`
	Version                       string `json:"version"`
	RestrictOutboundNetworkAccess string `json:"restrictOutboundNetworkAccess"`
}

type ReponseBase struct {
	ResourceId string     `json:"id"`
	Kind       string     `json:"kind"`
	Location   string     `json:"location"`
	Name       string     `json:"name"`
	Properties Properties `json:"properties"`
	Type       string     `json:"type"` // "Microsoft.Sql/servers"
}
