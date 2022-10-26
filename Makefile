run:
	@go run main.go --env ./app/config/.env.dev
test:
	@go test -v ./... -coverprofile cover.out
godoc:
	@godoc -http=0.0.0.0:6060 -v -timestamps=true -links=true -play=true
lint:
	@golangci-lint run -v --config ./.golangci.yml