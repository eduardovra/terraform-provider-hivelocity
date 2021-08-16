/*
 * Hivelocity API for Partners
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Webhook struct {
	WebhookId int32 `json:"webhookId"`
	Name string `json:"name"`
	Event string `json:"event"`
	Url string `json:"url"`
	Headers interface{} `json:"headers,omitempty"`
}
