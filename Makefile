build:
	echo "build source"
	go build
tidy:
	echo "running go mod tidy"
	go mod tidy
run:
	echo "main is running..."
	go run main.go
runtest:
	echo "run given test"
	go test -timeout 30s -run ^TestPasswordVaidation$ github.com/bubblepays/learning -v 
runAlltests:
	go test ./... -v
runmain:
	go run main.go
lint:
	golint ./...


