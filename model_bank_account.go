/*
 * Beanie ERP API
 *
 * An API specification for interacting with the Beanie ERP system
 *
 * API version: 0.8
 * Contact: dev@bean.ie
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package beanie

type BankAccount struct {
	Name string `json:"name"`
	BankName string `json:"bank_name"`
	CurrencyCode string `json:"currency_code"`
	Swift string `json:"swift"`
	Iban string `json:"iban"`
	Address1 string `json:"address1,omitempty"`
	Address2 string `json:"address2,omitempty"`
	Address3 string `json:"address3,omitempty"`
	City string `json:"city,omitempty"`
	StateCounty string `json:"state_county,omitempty"`
	ZipPostcode string `json:"zip_postcode,omitempty"`
	CountryName string `json:"country_name"`
	ContactName string `json:"contact_name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
	Website string `json:"website,omitempty"`
	Id int64 `json:"id,omitempty"`
}
