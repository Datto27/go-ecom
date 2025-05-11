build:
# with @ sign this command won't print out when run
	@go build -o bin/goecom main.go
# go build -o go-ecom -mod=vendor main.go

run: build
	@./bin/goecom

test:
	@go test -v ./..

