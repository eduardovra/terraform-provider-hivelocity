/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ContactUpdate struct {
	Email string `json:"email,omitempty"`
	Active int32 `json:"active,omitempty"`
	Phone string `json:"phone,omitempty"`
	Description string `json:"description,omitempty"`
	FullName string `json:"fullName,omitempty"`
}
