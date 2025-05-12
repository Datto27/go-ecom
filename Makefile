build:
	@echo "Building application"
	@go build -o bin/goecom main.go
# go build -o go-ecom -mod=vendor main.go

run: build
	@echo "Running application"
	@./bin/goecom

run-dev:
	@echo "Running dev server"
	@air

test:
	@go test -v ./..

