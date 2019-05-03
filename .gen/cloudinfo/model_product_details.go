/*
 * Product Info.
 *
 * The product info application uses the cloud provider APIs to asynchronously fetch and parse instance type attributes and prices, while storing the results in an in memory cache and making it available as structured data through a REST API.
 *
 * API version: 0.5.0
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudinfo

// ProductDetails extended view of the virtual machine details
type ProductDetails struct {
	Attributes map[string]string `json:"attributes,omitempty"`
	// Burst this is derived for now
	Burst bool `json:"burst,omitempty"`
	Category string `json:"category,omitempty"`
	CpusPerVm float64 `json:"cpusPerVm,omitempty"`
	// CurrentGen signals whether the instance type generation is the current one. Only applies for amazon
	CurrentGen bool `json:"currentGen,omitempty"`
	GpusPerVm float64 `json:"gpusPerVm,omitempty"`
	MemPerVm float64 `json:"memPerVm,omitempty"`
	NtwPerf string `json:"ntwPerf,omitempty"`
	NtwPerfCategory string `json:"ntwPerfCategory,omitempty"`
	OnDemandPrice float64 `json:"onDemandPrice,omitempty"`
	SpotPrice []ZonePrice `json:"spotPrice,omitempty"`
	Type string `json:"type,omitempty"`
	Zones []string `json:"zones,omitempty"`
}
