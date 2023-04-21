# Pet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Category** | Pointer to [**Category**](Category.md) |  | [optional] 
**Name** | **string** |  | 
**PhotoUrls** | **[]string** |  | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Status** | Pointer to **string** | pet status in the store | [optional] 
**ScalarInt32** | Pointer to **int32** |  | [optional] [default to 19]
**ScalarInt64** | Pointer to **int64** |  | [optional] [default to 19]
**ScalarFloat32** | Pointer to **float32** |  | [optional] [default to 19.19]
**ScalarFloat64** | Pointer to **float64** |  | [optional] [default to 19.19]
**ArrayInt32** | Pointer to **[]int32** |  | [optional] 
**ArrayInt64** | Pointer to **[]int64** |  | [optional] 
**ArrayFloat32** | Pointer to **[]float32** |  | [optional] 
**ArrayFloat64** | Pointer to **[]float64** |  | [optional] 
**ScalarString** | Pointer to **string** |  | [optional] [default to "19"]
**ScalarBoolean** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewPet

`func NewPet(name string, photoUrls []string, ) *Pet`

NewPet instantiates a new Pet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPetWithDefaults

`func NewPetWithDefaults() *Pet`

NewPetWithDefaults instantiates a new Pet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Pet) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pet) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pet) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Pet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategory

`func (o *Pet) GetCategory() Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Pet) GetCategoryOk() (*Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Pet) SetCategory(v Category)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Pet) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetName

`func (o *Pet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pet) SetName(v string)`

SetName sets Name field to given value.


### GetPhotoUrls

`func (o *Pet) GetPhotoUrls() []string`

GetPhotoUrls returns the PhotoUrls field if non-nil, zero value otherwise.

### GetPhotoUrlsOk

`func (o *Pet) GetPhotoUrlsOk() (*[]string, bool)`

GetPhotoUrlsOk returns a tuple with the PhotoUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoUrls

`func (o *Pet) SetPhotoUrls(v []string)`

SetPhotoUrls sets PhotoUrls field to given value.


### GetTags

`func (o *Pet) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Pet) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Pet) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Pet) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetStatus

`func (o *Pet) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Pet) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Pet) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Pet) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScalarInt32

`func (o *Pet) GetScalarInt32() int32`

GetScalarInt32 returns the ScalarInt32 field if non-nil, zero value otherwise.

### GetScalarInt32Ok

`func (o *Pet) GetScalarInt32Ok() (*int32, bool)`

GetScalarInt32Ok returns a tuple with the ScalarInt32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalarInt32

`func (o *Pet) SetScalarInt32(v int32)`

SetScalarInt32 sets ScalarInt32 field to given value.

### HasScalarInt32

`func (o *Pet) HasScalarInt32() bool`

HasScalarInt32 returns a boolean if a field has been set.

### GetScalarInt64

`func (o *Pet) GetScalarInt64() int64`

GetScalarInt64 returns the ScalarInt64 field if non-nil, zero value otherwise.

### GetScalarInt64Ok

`func (o *Pet) GetScalarInt64Ok() (*int64, bool)`

GetScalarInt64Ok returns a tuple with the ScalarInt64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalarInt64

`func (o *Pet) SetScalarInt64(v int64)`

SetScalarInt64 sets ScalarInt64 field to given value.

### HasScalarInt64

`func (o *Pet) HasScalarInt64() bool`

HasScalarInt64 returns a boolean if a field has been set.

### GetScalarFloat32

`func (o *Pet) GetScalarFloat32() float32`

GetScalarFloat32 returns the ScalarFloat32 field if non-nil, zero value otherwise.

### GetScalarFloat32Ok

`func (o *Pet) GetScalarFloat32Ok() (*float32, bool)`

GetScalarFloat32Ok returns a tuple with the ScalarFloat32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalarFloat32

`func (o *Pet) SetScalarFloat32(v float32)`

SetScalarFloat32 sets ScalarFloat32 field to given value.

### HasScalarFloat32

`func (o *Pet) HasScalarFloat32() bool`

HasScalarFloat32 returns a boolean if a field has been set.

### GetScalarFloat64

`func (o *Pet) GetScalarFloat64() float64`

GetScalarFloat64 returns the ScalarFloat64 field if non-nil, zero value otherwise.

### GetScalarFloat64Ok

`func (o *Pet) GetScalarFloat64Ok() (*float64, bool)`

GetScalarFloat64Ok returns a tuple with the ScalarFloat64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalarFloat64

`func (o *Pet) SetScalarFloat64(v float64)`

SetScalarFloat64 sets ScalarFloat64 field to given value.

### HasScalarFloat64

`func (o *Pet) HasScalarFloat64() bool`

HasScalarFloat64 returns a boolean if a field has been set.

### GetArrayInt32

`func (o *Pet) GetArrayInt32() []int32`

GetArrayInt32 returns the ArrayInt32 field if non-nil, zero value otherwise.

### GetArrayInt32Ok

`func (o *Pet) GetArrayInt32Ok() (*[]int32, bool)`

GetArrayInt32Ok returns a tuple with the ArrayInt32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayInt32

`func (o *Pet) SetArrayInt32(v []int32)`

SetArrayInt32 sets ArrayInt32 field to given value.

### HasArrayInt32

`func (o *Pet) HasArrayInt32() bool`

HasArrayInt32 returns a boolean if a field has been set.

### GetArrayInt64

`func (o *Pet) GetArrayInt64() []int64`

GetArrayInt64 returns the ArrayInt64 field if non-nil, zero value otherwise.

### GetArrayInt64Ok

`func (o *Pet) GetArrayInt64Ok() (*[]int64, bool)`

GetArrayInt64Ok returns a tuple with the ArrayInt64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayInt64

`func (o *Pet) SetArrayInt64(v []int64)`

SetArrayInt64 sets ArrayInt64 field to given value.

### HasArrayInt64

`func (o *Pet) HasArrayInt64() bool`

HasArrayInt64 returns a boolean if a field has been set.

### GetArrayFloat32

`func (o *Pet) GetArrayFloat32() []float32`

GetArrayFloat32 returns the ArrayFloat32 field if non-nil, zero value otherwise.

### GetArrayFloat32Ok

`func (o *Pet) GetArrayFloat32Ok() (*[]float32, bool)`

GetArrayFloat32Ok returns a tuple with the ArrayFloat32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayFloat32

`func (o *Pet) SetArrayFloat32(v []float32)`

SetArrayFloat32 sets ArrayFloat32 field to given value.

### HasArrayFloat32

`func (o *Pet) HasArrayFloat32() bool`

HasArrayFloat32 returns a boolean if a field has been set.

### GetArrayFloat64

`func (o *Pet) GetArrayFloat64() []float64`

GetArrayFloat64 returns the ArrayFloat64 field if non-nil, zero value otherwise.

### GetArrayFloat64Ok

`func (o *Pet) GetArrayFloat64Ok() (*[]float64, bool)`

GetArrayFloat64Ok returns a tuple with the ArrayFloat64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayFloat64

`func (o *Pet) SetArrayFloat64(v []float64)`

SetArrayFloat64 sets ArrayFloat64 field to given value.

### HasArrayFloat64

`func (o *Pet) HasArrayFloat64() bool`

HasArrayFloat64 returns a boolean if a field has been set.

### GetScalarString

`func (o *Pet) GetScalarString() string`

GetScalarString returns the ScalarString field if non-nil, zero value otherwise.

### GetScalarStringOk

`func (o *Pet) GetScalarStringOk() (*string, bool)`

GetScalarStringOk returns a tuple with the ScalarString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalarString

`func (o *Pet) SetScalarString(v string)`

SetScalarString sets ScalarString field to given value.

### HasScalarString

`func (o *Pet) HasScalarString() bool`

HasScalarString returns a boolean if a field has been set.

### GetScalarBoolean

`func (o *Pet) GetScalarBoolean() bool`

GetScalarBoolean returns the ScalarBoolean field if non-nil, zero value otherwise.

### GetScalarBooleanOk

`func (o *Pet) GetScalarBooleanOk() (*bool, bool)`

GetScalarBooleanOk returns a tuple with the ScalarBoolean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalarBoolean

`func (o *Pet) SetScalarBoolean(v bool)`

SetScalarBoolean sets ScalarBoolean field to given value.

### HasScalarBoolean

`func (o *Pet) HasScalarBoolean() bool`

HasScalarBoolean returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


