package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kurakura967/go-openapi-demo/api/generated/openapi"
)

func WriteInternalErrorResponse(w http.ResponseWriter, r *http.Request, err error) {

	res := &openapi.InternalServerError{
		Message: ToPtr("An expected error has occurred"),
		Type:    ToPtr("internal_server_error"),
	}
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println("failed to encode json.", err)
	}
}

func WriteBadRequestErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	var p, m string

	switch v := err.(type) {
	case *openapi.InvalidParamFormatError:
		p = v.ParamName
		m = v.Error()
	case *openapi.RequiredParamError:
		p = v.ParamName
		m = v.Error()
	}

	filedError := openapi.FieldError{
		FieldName: &p,
		Message:   &m,
	}

	res := openapi.BadRequestErrorResponse{
		Error: &openapi.BadRequestError{
			Field:   &filedError,
			Message: ToPtr("Request parameters are invalid"),
			Type:    ToPtr("invalid_parameter"),
		},
	}
	w.WriteHeader(http.StatusBadRequest)
	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Println("failed to encode json.", err)
	}
}

func ToPtr[T any](t T) *T {
	return &t
}
