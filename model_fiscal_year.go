/*
 * Beanie ERP API
 *
 * An API specification for interacting with the Beanie ERP system
 *
 * API version: 0.2
 * Contact: dev@bean.ie
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package beanie

type FiscalYear struct {
	Year int32 `json:"year"`
	Closed bool `json:"closed"`
}