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

// AttributeResponse holds attribute values
type AttributeResponse struct {
	AttributeName string `json:"attributeName,omitempty"`
	AttributeValues []float64 `json:"attributeValues,omitempty"`
}
