vendor:
	@echo "Run go mod vendor"
	go mod vendor
build:
	@echo "Running build target"
	make vendor
junit:
	go test -v ./...
#	go test -v 2>&1 | go-junit-report > report.xml
