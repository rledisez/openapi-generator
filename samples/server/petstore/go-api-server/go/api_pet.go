/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstoreserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// PetApiController binds http requests to an api service and writes the service results to the http response
type PetApiController struct {
	service PetApiServicer
	errorHandler ErrorHandler
}

// PetApiOption for how the controller is set up.
type PetApiOption func(*PetApiController)

// WithPetApiErrorHandler inject ErrorHandler into controller
func WithPetApiErrorHandler(h ErrorHandler) PetApiOption {
	return func(c *PetApiController) {
		c.errorHandler = h
	}
}

// NewPetApiController creates a default api controller
func NewPetApiController(s PetApiServicer, opts ...PetApiOption) Router {
	controller := &PetApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the PetApiController
func (c *PetApiController) Routes() Routes {
	return Routes{
		"AddPet": Route{
			strings.ToUpper("Post"),
			"/v2/pet",
			c.AddPet,
		},
		"DeletePet": Route{
			strings.ToUpper("Delete"),
			"/v2/pet/{petId}/{scalarInt32}/{scalarInt64}/{scalarFloat32}/{scalarFloat64}",
			c.DeletePet,
		},
		"FindPetsByStatus": Route{
			strings.ToUpper("Get"),
			"/v2/pet/findByStatus",
			c.FindPetsByStatus,
		},
		"FindPetsByTags": Route{
			strings.ToUpper("Get"),
			"/v2/pet/findByTags",
			c.FindPetsByTags,
		},
		"GetPetById": Route{
			strings.ToUpper("Get"),
			"/v2/pet/{petId}/{scalarInt32}/{scalarInt64}/{scalarFloat32}/{scalarFloat64}",
			c.GetPetById,
		},
		"UpdatePet": Route{
			strings.ToUpper("Put"),
			"/v2/pet",
			c.UpdatePet,
		},
		"UpdatePetWithForm": Route{
			strings.ToUpper("Post"),
			"/v2/pet/{petId}/{scalarInt32}/{scalarInt64}/{scalarFloat32}/{scalarFloat64}",
			c.UpdatePetWithForm,
		},
		"UploadFile": Route{
			strings.ToUpper("Post"),
			"/v2/pet/{petId}/uploadImage",
			c.UploadFile,
		},
	}
}

// AddPet - Add a new pet to the store
func (c *PetApiController) AddPet(w http.ResponseWriter, r *http.Request) {
	petParam := Pet{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&petParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPetRequired(petParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertPetConstraints(petParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddPet(r.Context(), petParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeletePet - Deletes a pet
func (c *PetApiController) DeletePet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	petIdParam, err := parseInt64Parameter(params["petId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	scalarInt32Param, err := parseInt32Parameter(params["scalarInt32"], true, WithMinimum(int64(17)), WithMaximum(int64(42)))
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	scalarInt64Param, err := parseInt64Parameter(params["scalarInt64"], true, WithMinimum(int64(17)), WithMaximum(int64(42)))
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	scalarFloat32Param, err := parseFloat32Parameter(params["scalarFloat32"], true, WithMinimum(float64(17.17)), WithMaximum(float64(42.42)))
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	scalarFloat64Param, err := parseFloat64Parameter(params["scalarFloat64"], true, WithMinimum(float64(17.17)), WithMaximum(float64(42.42)))
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	apiKeyParam := r.Header.Get("api_key")
	result, err := c.service.DeletePet(r.Context(), petIdParam, scalarInt32Param, scalarInt64Param, scalarFloat32Param, scalarFloat64Param, apiKeyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// FindPetsByStatus - Finds Pets by status
func (c *PetApiController) FindPetsByStatus(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	statusParam := strings.Split(query.Get("status"), ",")
	scalarInt32Param, err := parseInt32Parameter(query.Get("scalarInt32"), false, WithDefault(int64(19)), WithMinimum(int64(17)), WithMaximum(int64(42)))
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	scalarInt64Param, err := parseInt64Parameter(query.Get("scalarInt64"), false, WithDefault(int64(19)), WithMinimum(int64(17)), WithMaximum(int64(42)))
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	scalarFloat32Param, err := parseFloat32Parameter(query.Get("scalarFloat32"), false, WithDefault(float64(19.19)), WithMinimum(float64(17.17)), WithMaximum(float64(42.42)))
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	scalarFloat64Param, err := parseFloat64Parameter(query.Get("scalarFloat64"), false, WithDefault(float64(19.19)), WithMinimum(float64(17.17)), WithMaximum(float64(42.42)))
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	arrayInt32Param, err := parseInt32ArrayParameter(query.Get("arrayInt32"), ",", false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	arrayInt64Param, err := parseInt64ArrayParameter(query.Get("arrayInt64"), ",", false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	arrayFloat32Param, err := parseFloat32ArrayParameter(query.Get("arrayFloat32"), ",", false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
	return
	}
	arrayFloat64Param, err := parseFloat64ArrayParameter(query.Get("arrayFloat64"), ",", false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	scalarStringParam := "19"
	if query.Has("scalarString") {
		scalarStringParam = query.Get("scalarString")
	}
	scalarBooleanParam, err := parseBoolParameter(query.Get("scalarBoolean"), false, WithDefault(true))
	if err != nil {
		w.WriteHeader(500)
		return
	}
	result, err := c.service.FindPetsByStatus(r.Context(), statusParam, scalarInt32Param, scalarInt64Param, scalarFloat32Param, scalarFloat64Param, arrayInt32Param, arrayInt64Param, arrayFloat32Param, arrayFloat64Param, scalarStringParam, scalarBooleanParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// FindPetsByTags - Finds Pets by tags
// Deprecated
func (c *PetApiController) FindPetsByTags(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	tagsParam := strings.Split(query.Get("tags"), ",")
	result, err := c.service.FindPetsByTags(r.Context(), tagsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetPetById - Find pet by ID
func (c *PetApiController) GetPetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	petIdParam, err := parseInt64Parameter(params["petId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	scalarInt32Param, err := parseInt32Parameter(params["scalarInt32"], true, WithMinimum(int64(17)), WithMaximum(int64(42)))
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	scalarInt64Param, err := parseInt64Parameter(params["scalarInt64"], true, WithMinimum(int64(17)), WithMaximum(int64(42)))
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	scalarFloat32Param, err := parseFloat32Parameter(params["scalarFloat32"], true, WithMinimum(float64(17.17)), WithMaximum(float64(42.42)))
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	scalarFloat64Param, err := parseFloat64Parameter(params["scalarFloat64"], true, WithMinimum(float64(17.17)), WithMaximum(float64(42.42)))
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetPetById(r.Context(), petIdParam, scalarInt32Param, scalarInt64Param, scalarFloat32Param, scalarFloat64Param)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdatePet - Update an existing pet
func (c *PetApiController) UpdatePet(w http.ResponseWriter, r *http.Request) {
	petParam := Pet{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&petParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPetRequired(petParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertPetConstraints(petParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdatePet(r.Context(), petParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdatePetWithForm - Updates a pet in the store with form data
func (c *PetApiController) UpdatePetWithForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	params := mux.Vars(r)
	petIdParam, err := parseInt64Parameter(params["petId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	scalarInt32Param, err := parseInt32Parameter(params["scalarInt32"], true, WithMinimum(int64(17)), WithMaximum(int64(42)))
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	scalarInt64Param, err := parseInt64Parameter(params["scalarInt64"], true, WithMinimum(int64(17)), WithMaximum(int64(42)))
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	scalarFloat32Param, err := parseFloat32Parameter(params["scalarFloat32"], true, WithMinimum(float64(17.17)), WithMaximum(float64(42.42)))
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	scalarFloat64Param, err := parseFloat64Parameter(params["scalarFloat64"], true, WithMinimum(float64(17.17)), WithMaximum(float64(42.42)))
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
				nameParam := r.FormValue("name")
				statusParam := r.FormValue("status")
	result, err := c.service.UpdatePetWithForm(r.Context(), petIdParam, scalarInt32Param, scalarInt64Param, scalarFloat32Param, scalarFloat64Param, nameParam, statusParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UploadFile - uploads an image
func (c *PetApiController) UploadFile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	params := mux.Vars(r)
	petIdParam, err := parseInt64Parameter(params["petId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
				additionalMetadataParam := r.FormValue("additionalMetadata")
	
	fileParam, err := ReadFormFileToTempFile(r, "file")
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
			result, err := c.service.UploadFile(r.Context(), petIdParam, additionalMetadataParam, fileParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
