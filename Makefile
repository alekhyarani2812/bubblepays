exevendor:
	@echo "Run go mod tidy"
	go mod tidy
	@echo "Run go mod vendor"
	go mod vendor
build:
	@echo "Running build target"
#	make exevendor
	go build
junit:
	go test -v ./...
#	go test -v 2>&1 | go-junit-report > report.xml
	go test -json ./... | go-junit-report > report.xml
