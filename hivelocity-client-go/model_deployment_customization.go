/*
 * Hivelocity API for Partners
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DeploymentCustomization struct {
	// Operating System's Name or ID
	OperatingSystem string `json:"operatingSystem"`
	Hostnames []string `json:"hostnames"`
	AdditionalNotes []string `json:"additionalNotes,omitempty"`
	ProductId int32 `json:"productId"`
	PublicSshKeyId int32 `json:"publicSshKeyId,omitempty"`
	LocationCode string `json:"locationCode,omitempty"`
	Options []int32 `json:"options,omitempty"`
	Quantity int32 `json:"quantity,omitempty"`
	// must be one of ['monthly', 'quarterly', 'semi-annually', 'annually', 'biennial', 'triennial', 'hourly']
	BillingPeriod string `json:"billingPeriod,omitempty"`
}
