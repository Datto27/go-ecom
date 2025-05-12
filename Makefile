CORE_DIR = ./microservices/core

.PHONY: build-core
build-core:
	@echo "Building application..."
	@cd $(CORE_DIR) && go build -o bin/core main.go
# go build -o go-ecom -mod=vendor main.go

.PHONY: run-core
run-core: build-core
	@echo "Running application"
	@cd $(CORE_DIR) && ./bin/core

.PHONY: run-dev-core
run-dev-core:
	@echo "Running dev server"
	@cd $(CORE_DIR) && air

.PHONY: test-core
test-core:
	@cd $(CORE_DIR) && go test ./...

.PHONY: swag-core
swag-core:
	@cd $(CORE_DIR) && swag init --output ./docs --parseDependency --parseInternal
