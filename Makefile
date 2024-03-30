vendor:
	@echo "Run go mod vendor"
	go mod vendor
build:
	@echo "Running build target"
	make vendor
