/*
 * Hivelocity API for Partners
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DeviceIpmiInfo struct {
	Info *IpmiInfo `json:"info,omitempty"`
	Sensors []IpmiSensor `json:"sensors,omitempty"`
}
