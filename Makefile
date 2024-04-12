.PHONY: help
help:
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

.PHONY: migrate
migrate: ## Migrate to DB set in config
	go run scripts/migrate/migrate.go

.PHONY: generate
generate: ## Run go generate on project level
	go generate ./...

.PHONY: tidy
tidy: ## Format code and tidy modfile
	go fmt ./...
	go mod tidy -v

.PHONY: audit
audit: ## Run quality control checks
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...


.PHONY: install-tools
install-tools: ## Install go tools required for this service
	go install go.uber.org/mock/mockgen@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest