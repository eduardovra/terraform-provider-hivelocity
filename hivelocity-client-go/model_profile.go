/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Profile struct {
	Company interface{} `json:"company,omitempty"`
	Zip interface{} `json:"zip,omitempty"`
	City interface{} `json:"city,omitempty"`
	FullName interface{} `json:"full_name,omitempty"`
	Id int32 `json:"id,omitempty"`
	Country interface{} `json:"country,omitempty"`
	First string `json:"first,omitempty"`
	Phone string `json:"phone,omitempty"`
	Login string `json:"login,omitempty"`
	State interface{} `json:"state,omitempty"`
	Last string `json:"last,omitempty"`
	Created interface{} `json:"created,omitempty"`
	MetaData interface{} `json:"meta_data,omitempty"`
	Fax interface{} `json:"fax,omitempty"`
	Address interface{} `json:"address,omitempty"`
	Email string `json:"email,omitempty"`
	IsClient bool `json:"is_client,omitempty"`
}
