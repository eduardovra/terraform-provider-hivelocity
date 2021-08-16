/*
 * Hivelocity API for Partners
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Stock struct {
	ProductMonthlyPrice float32 `json:"product_monthly_price,omitempty"`
	Edge bool `json:"edge,omitempty"`
	ProductAnnuallyPrice float32 `json:"product_annually_price,omitempty"`
	ProductDisplayPrice float32 `json:"product_display_price,omitempty"`
	Core bool `json:"core,omitempty"`
	MonthlyLocationPremium float32 `json:"monthly_location_premium,omitempty"`
	DataCenter string `json:"data_center,omitempty"`
	ProductCpuCores string `json:"product_cpu_cores,omitempty"`
	ProductTriennialPrice float32 `json:"product_triennial_price,omitempty"`
	Stock string `json:"stock,omitempty"`
	ProductGpu string `json:"product_gpu,omitempty"`
	ProductDisabledBillingPeriods []string `json:"product_disabled_billing_periods,omitempty"`
	BiennialLocationPremium float32 `json:"biennial_location_premium,omitempty"`
	ProductName string `json:"product_name,omitempty"`
	ProductDrive string `json:"product_drive,omitempty"`
	ProductCpu string `json:"product_cpu,omitempty"`
	ProductOnSale bool `json:"product_on_sale,omitempty"`
	TriennialLocationPremium float32 `json:"triennial_location_premium,omitempty"`
	ProductHourlyPrice float32 `json:"product_hourly_price,omitempty"`
	AnnuallyLocationPremium float32 `json:"annually_location_premium,omitempty"`
	ProductQuarterlyPrice float32 `json:"product_quarterly_price,omitempty"`
	ProductBiennialPrice float32 `json:"product_biennial_price,omitempty"`
	ProductOriginalPrice float32 `json:"product_original_price,omitempty"`
	QuarterlyLocationPremium float32 `json:"quarterly_location_premium,omitempty"`
	ProductBandwidth string `json:"product_bandwidth,omitempty"`
	ProductMemory string `json:"product_memory,omitempty"`
	ProductId int32 `json:"product_id,omitempty"`
	SemiAnnuallyLocationPremium float32 `json:"semi_annually_location_premium,omitempty"`
	HourlyLocationPremium float32 `json:"hourly_location_premium,omitempty"`
	ProductSemiAnnuallyPrice float32 `json:"product_semi_annually_price,omitempty"`
}
