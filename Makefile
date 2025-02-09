.PHONY: help
help:
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run-app
run-app: ## Run main service
	go build -o ./bin/app ./cmd/app/
	./bin/app

.PHONY: run-app-reload
run-app-reload: ## Run main service via air
	air --build.cmd "go build -o ./bin/app ./cmd/app/" --build.bin "./bin/app"

.PHONY: check-config
check-config: ## Assert that all configurations are ok when initializing app
	go build -o ./bin/app ./cmd/app/
	./bin/app -t

.PHONY: test
test: ## Run all tests
	./scripts/testing/gocov.sh

.PHONY: migrate-db-new
migrate-db-new: ## Create a new sql db migration
ifndef file_name
	$(error file_name is not defined. Use make migrate-db-new file_name=<name>)
endif
	goose -dir db/migrations create $(file_name) sql

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
	go install github.com/air-verse/air@latest