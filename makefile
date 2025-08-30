.PHONY: fmt lint test

define run-go-command
	@echo "$(1) Go modules..."
	@find . -name "go.mod" -exec dirname {} \; | xargs -I {} sh -c 'echo "$(1) {}" && cd {} && $(2)'
endef

fmt:
	$(call run-go-command,Formatting,go fmt ./...)

lint:
	$(call run-go-command,Linting,go vet ./...)

test:
	$(call run-go-command,Testing,go test -v ./...)
