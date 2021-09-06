/*
 * Hivelocity API for Partners
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SingleMxRecordReturn struct {
	Name string `json:"name"`
	Preference int32 `json:"preference"`
	Exchange string `json:"exchange"`
	Id int32 `json:"id"`
	Type_ string `json:"type"`
	Ttl int32 `json:"ttl"`
	DomainId int32 `json:"domainId"`
}
