generate-api:
	oapi-codegen -config docs/oapi-codegen.types.yaml docs/openapi.yaml

gomod:
	cd apps/api && go mod tidy && cd ../../