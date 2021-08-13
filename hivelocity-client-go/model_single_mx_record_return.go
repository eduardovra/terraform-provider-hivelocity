/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SingleMxRecordReturn struct {
	Ttl int32 `json:"ttl"`
	Id int32 `json:"id"`
	Exchange string `json:"exchange"`
	Preference int32 `json:"preference"`
	DomainId int32 `json:"domainId"`
	Type_ string `json:"type"`
	Name string `json:"name"`
}
