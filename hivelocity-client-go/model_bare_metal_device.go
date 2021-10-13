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
	VlanId int32 `json:"vlanId,omitempty"`
	PublicSshKeyId int32 `json:"publicSshKeyId,omitempty"`
	LocationName string `json:"locationName,omitempty"`
	ProductId int32 `json:"productId,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Script string `json:"script,omitempty"`
	DeviceId int32 `json:"deviceId,omitempty"`
	PrimaryIp string `json:"primaryIp,omitempty"`
	OrderId int32 `json:"orderId,omitempty"`
	PowerStatus string `json:"powerStatus,omitempty"`
	OsName string `json:"osName,omitempty"`
	ServiceId int32 `json:"serviceId,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	ProductName string `json:"productName,omitempty"`
	Period string `json:"period,omitempty"`
}
