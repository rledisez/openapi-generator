/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstoreserver



import "errors" // FIXME: why not in #imports
import "encoding/json" // FIXME: why not in #imports

var _ = errors.New("") // to enforce the use of errors
var _ = json.NewDecoder(nil) // to enforce the use of encoding/json


// Pet - A pet for sale in the pet store
type Pet struct {

	Id int64 `json:"id,omitempty"`

	Category *Category `json:"category,omitempty"`

	Name string `json:"name"`

	PhotoUrls *[]string `json:"photoUrls"`

	Tags *[]Tag `json:"tags,omitempty"`

	// pet status in the store
	Status string `json:"status,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *Pet) UnmarshalJSON(data []byte) error {

	type Alias Pet // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertPetRequired checks if the required fields are not zero-ed
func AssertPetRequired(obj Pet) error {
	elements := map[string]interface{}{
		"name": obj.Name,
		"photoUrls": obj.PhotoUrls,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if obj.Category != nil {
		if err := AssertCategoryRequired(*obj.Category); err != nil {
			return err
		}
	}
	if obj.Tags != nil {
		for _, el := range *obj.Tags {
			if err := AssertTagRequired(el); err != nil {
				return err
			}
		}
	}
	return nil
}

// AssertPetConstraints checks if the values respects the defined constraints
func AssertPetConstraints(obj Pet) error {
	return nil
}
