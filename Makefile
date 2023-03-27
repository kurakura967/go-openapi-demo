.PHONY: run
run: ## Run server
	go1.19 run cmd/go-openapi-demo/main.go

.PHONY: codegen
codegen: ## Generate code from openapi
	rm -f api/generated/openapi/*.gen.go
	oapi-codegen -generate "chi-server" -old-config-style -package openapi openapi/api/openapi.yaml > api/generated/openapi/chi-server.gen.go
	oapi-codegen -generate "spec" -old-config-style -package openapi openapi/api/openapi.yaml > api/generated/openapi/spec.gen.go
	oapi-codegen -generate "types" -old-config-style -package openapi openapi/api/openapi.yaml > api/generated/openapi/types.gen.go

.PHONY: help
help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
