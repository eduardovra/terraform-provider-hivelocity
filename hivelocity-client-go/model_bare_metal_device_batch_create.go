/*
 * Hivelocity API for Partners
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BareMetalDeviceBatchCreate struct {
	Devices []BareMetalDeviceCreate `json:"devices,omitempty"`
	OrderGroupId int32 `json:"orderGroupId,omitempty"`
}
