.PHONY: gen deps

default: golang gen

golang:
	@echo "--> Go Version"
	@go version

deps:
	@echo "--> Installing go dependencies"
	@go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

gen: 
	@go generate ./...
	oapi-codegen --config=api/models.yaml api/Billingo-Billingo-3.0.14-resolved.yaml
	oapi-codegen --config=api/client.yaml api/Billingo-Billingo-3.0.14-resolved.yaml

authors:
	@echo "--> Updating the AUTHORS"
	@git log --format='%aN <%aE>' | sort -u > AUTHORS

changelog:
	git log $(shell git tag | tail -n1)..HEAD --no-merges --format=%B > changelog
