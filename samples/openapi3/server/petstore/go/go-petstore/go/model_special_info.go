// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 */

package petstoreserver




// SpecialInfo - An order info for a pets from the pet store
type SpecialInfo struct {

	Promotion bool `json:"promotion,omitempty"`

	RequireTest string `json:"requireTest"`

	Type string `json:"type,omitempty"`
}

// NewSpecialInfoWithDefaults instantiates a new SpecialInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecialInfoWithDefaults() SpecialInfo {
	this := SpecialInfo{}
	this.Type = "dog"
	return this
}

// AssertSpecialInfoRequired checks if the required fields are not zero-ed
func AssertSpecialInfoRequired(obj SpecialInfo) error {
	elements := map[string]interface{}{
		"requireTest": obj.RequireTest,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertSpecialInfoConstraints checks if the values respects the defined constraints
func AssertSpecialInfoConstraints(obj SpecialInfo) error {
	return nil
}
