/*
 * Hivelocity API for Partners
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ARecordUpdate struct {
	Ttl int32 `json:"ttl,omitempty"`
	Name string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}
