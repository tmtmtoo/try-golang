.PHONY: fmt lint test clean list

define run-go-command
	@echo "$(1) Go modules..."
	@find . -name "go.mod" -exec dirname {} \; | xargs -I {} sh -c 'echo "$(1) {}" && cd {} && $(2)'
endef

list:
	@echo "Available make targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-10s - %s\n", $$1, $$2}'

fmt: ## Format all Go modules
	$(call run-go-command,Formatting,go fmt ./...)

lint: ## Lint all Go modules  
	$(call run-go-command,Linting,go vet ./...)

test: ## Test all Go modules
	$(call run-go-command,Testing,go test -v ./...)

clean: ## Clean all Go modules
	$(call run-go-command,Cleaning,go clean ./...)

add-migration: ## Create a new database migration. Usage: make add-migration name=<migration_name>
	@migrate create -ext sql -dir migrations -seq $(name)

migrate: ## Apply all up database migrations
	@migrate -path migrations -database "$(DATABASE_URL)" up
