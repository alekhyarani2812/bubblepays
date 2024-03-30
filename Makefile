Vendor:
	@echo "Run go mod vendor"
	go mod vendor
Build:
	@echo "Running build target"
	make vendor
RunUnitTests:
	go test ./...
