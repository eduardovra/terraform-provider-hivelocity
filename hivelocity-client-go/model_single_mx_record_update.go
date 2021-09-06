/*
 * Hivelocity API for Partners
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SingleMxRecordUpdate struct {
	Preference int32 `json:"preference,omitempty"`
	Name string `json:"name,omitempty"`
	Ttl int32 `json:"ttl,omitempty"`
	Exchange string `json:"exchange"`
}
