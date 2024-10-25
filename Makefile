build:
# with @ sign this command won't print out when run
	@go build -o bin/goecom cmd/main.go

run: build
	@./bin/goecom

test:
	@got test -v ./..

