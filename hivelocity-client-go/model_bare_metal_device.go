/*
 * Hivelocity API for Partners
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BareMetalDevice struct {
	Script string `json:"script,omitempty"`
	PublicSshKeyId int32 `json:"publicSshKeyId,omitempty"`
	OrderId int32 `json:"orderId,omitempty"`
	ProductName string `json:"productName,omitempty"`
	DeviceId int32 `json:"deviceId,omitempty"`
	LocationName string `json:"locationName,omitempty"`
	PowerStatus string `json:"powerStatus,omitempty"`
	Period string `json:"period,omitempty"`
	VlanId int32 `json:"vlanId,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	OsName string `json:"osName,omitempty"`
	ServiceId int32 `json:"serviceId,omitempty"`
	ProductId int32 `json:"productId,omitempty"`
	Tags []string `json:"tags,omitempty"`
	PrimaryIp string `json:"primaryIp,omitempty"`
}
