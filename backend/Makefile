.PHONY: generated

generated:
	@echo "Generating code..."

generated: generated/api.gen.go

generated/api.gen.go:
	@echo "Generating API server..."
	mkdir generated
	oapi-codegen --package=generated --generate server,types,spec ../api/api.yaml > generated/api.gen.go 
