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
	ProductId int32 `json:"productId"`
	Options []int32 `json:"options,omitempty"`
	Hostnames []string `json:"hostnames"`
	Quantity int32 `json:"quantity,omitempty"`
	LocationCode string `json:"locationCode,omitempty"`
	// must be one of ['monthly', 'quarterly', 'semi-annually', 'annually', 'biennial', 'triennial', 'hourly']
	BillingPeriod string `json:"billingPeriod,omitempty"`
	// Operating System's Name or ID
	OperatingSystem string `json:"operatingSystem"`
	PublicSshKeyId int32 `json:"publicSshKeyId,omitempty"`
	AdditionalNotes []string `json:"additionalNotes,omitempty"`
}
