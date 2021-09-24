/*
 * Hivelocity API for Partners
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DeploymentStart struct {
	Script string `json:"script,omitempty"`
	BillingInfo int32 `json:"billingInfo"`
}