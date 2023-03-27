package api

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/kurakura967/go-openapi-demo/api/generated/openapi"
	embed "github.com/kurakura967/go-openapi-demo/openapi"
)

func NewRouter() *chi.Mux {

	swagger, err := openapi.GetSwagger()
	if err != nil {

	}
	swagger.Servers = nil

	r := chi.NewRouter()
	addHandlersForOpenAPI(r)
	middleware.OapiRequestValidator(swagger)

	ser := NewServer()

	openapi.HandlerWithOptions(ser, openapi.ChiServerOptions{
		BaseURL:    "/v1",
		BaseRouter: r,
		ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {

			switch err.(type) {
			case *openapi.InvalidParamFormatError:
				WriteBadRequestErrorResponse(w, r, err)
			default:
				WriteInternalErrorResponse(w, r, err)
			}
		},
	})

	return r
}

func addHandlersForOpenAPI(r *chi.Mux) {
	swagger, err := openapi.GetSwagger()
	if err != nil {
		log.Fatal("loading openapi spec is failed. ", err)
	}

	r.HandleFunc("/api-docs", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(swagger); err != nil {
			fmt.Fprintln(os.Stderr, "json encoding for openapi spec is failed. ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	fsys, _ := fs.Sub(embed.Content, "swagger-ui")
	r.Handle("/swagger-ui/*", http.StripPrefix("/swagger-ui", http.FileServer(http.FS(fsys))))
}
