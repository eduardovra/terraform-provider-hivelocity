/*
 * Hivelocity API for Partners
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type IpmiInfo struct {
	IpmiVersion string `json:"ipmiVersion,omitempty"`
	SdrRepositoryDevice string `json:"sdrRepositoryDevice,omitempty"`
	IpmbEventGenerator string `json:"ipmbEventGenerator,omitempty"`
	FruInventoryDevice string `json:"fruInventoryDevice,omitempty"`
	Bridge string `json:"bridge,omitempty"`
	ProductId string `json:"productId,omitempty"`
	DeviceAvailable string `json:"deviceAvailable,omitempty"`
	DeviceId string `json:"deviceId,omitempty"`
	DeviceRevision string `json:"deviceRevision,omitempty"`
	ChassisDevice string `json:"chassisDevice,omitempty"`
	SensorDevice string `json:"sensorDevice,omitempty"`
	DeviceSDRs string `json:"deviceSDRs,omitempty"`
	ManufacturerId string `json:"manufacturerId,omitempty"`
	IpmbEventReceiver string `json:"ipmbEventReceiver,omitempty"`
	SelDevice string `json:"selDevice,omitempty"`
	AuxFirmwareRevInfo string `json:"auxFirmwareRevInfo,omitempty"`
	FirmwareRevision string `json:"firmwareRevision,omitempty"`
}
