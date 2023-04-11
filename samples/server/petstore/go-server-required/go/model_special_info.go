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


// SpecialInfo - An order info for a pets from the pet store
type SpecialInfo struct {

	Promotion bool `json:"promotion,omitempty"`

	Type string `json:"type,omitempty"`
}

// UnmarshalJSON parse JSON while respecting the default values specified
func (o *SpecialInfo) UnmarshalJSON(data []byte) error {
    type Alias SpecialInfo // Avoid infinite recursion
    aux := Alias{
	}
    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }
    *o = SpecialInfo(aux)
    return nil
}

// AssertSpecialInfoRequired checks if the required fields are not zero-ed
func AssertSpecialInfoRequired(obj SpecialInfo) error {
	return nil
}

// AssertSpecialInfoConstraints checks if the values respects the defined constraints
func AssertSpecialInfoConstraints(obj SpecialInfo) error {
	return nil
}

// AssertRecurseSpecialInfoRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of SpecialInfo (e.g. [][]SpecialInfo), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseSpecialInfoRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aSpecialInfo, ok := obj.(SpecialInfo)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertSpecialInfoRequired(aSpecialInfo)
	})
}
