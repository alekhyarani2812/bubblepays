vendor:
	@echo "Run go mod vendor"
	go mod tidy
	go mod vendor
build:
	@echo "Running build target"
	go get ./...
	go mod tidy
	go build
junit:
	go test -v ./...
#	go test -v 2>&1 | go-junit-report > report.xml
	go test -json ./... | go-junit-report > report.xml
