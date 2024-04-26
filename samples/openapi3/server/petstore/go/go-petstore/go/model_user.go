// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 */

package petstoreserver




// User - A User who is purchasing from the pet store
type User struct {

	Id int64 `json:"id,omitempty"`

	Username string `json:"username,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Email string `json:"email,omitempty"`

	Password string `json:"password,omitempty"`

	Phone *string `json:"phone,omitempty"`

	// User Status
	UserStatus int32 `json:"userStatus,omitempty"`

	// An array 1-deep.
	DeepSliceModel *[][][]Tag `json:"deepSliceModel"`

	// An array 1-deep.
	DeepSliceMap [][]AnObject `json:"deepSliceMap,omitempty"`
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() User {
	this := User{}
	return this
}

// AssertUserRequired checks if the required fields are not zero-ed
func AssertUserRequired(obj User) error {
	elements := map[string]interface{}{
		"deepSliceModel": obj.DeepSliceModel,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if obj.DeepSliceModel != nil {
		if err := AssertRecurseInterfaceRequired(*obj.DeepSliceModel, AssertTagRequired); err != nil {
			return err
		}
	}
	if err := AssertRecurseInterfaceRequired(obj.DeepSliceMap, AssertAnObjectRequired); err != nil {
		return err
	}
	return nil
}

// AssertUserConstraints checks if the values respects the defined constraints
func AssertUserConstraints(obj User) error {
    if obj.DeepSliceModel != nil {
     	if err := AssertRecurseInterfaceRequired(*obj.DeepSliceModel, AssertTagConstraints); err != nil {
     		return err
     	}
    }
	if err := AssertRecurseInterfaceRequired(obj.DeepSliceMap, AssertAnObjectConstraints); err != nil {
		return err
	}
	return nil
}
