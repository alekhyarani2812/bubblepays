vendor:
	@echo "Run go mod vendor"
	go mod vendor -mod=mod
build:
	@echo "Running build target"
	make vendor
junit:
	go test -v ./...
#	go test -v 2>&1 | go-junit-report > report.xml
	go test -json ./... | go-junit-report > report.xml
