# To run makefile use this command $ make commandname
build: tidy
	echo "build source"
	go build
tidy:
	echo "running go mod tidy"
	go mod tidy
run:
	echo "main is running..."
	go run main.go
# runtest:
# 	echo "run given test"
# 	#go test -timeout 30s -run ^TestPasswordVaidation$ github.com/bubblepays/learning -v 
# 	go test -v -run ^TestPasswordVaidation_invalidPwd$ github.com/bubblepays/learning

runtest:
	echo "running unit tests"
	go test -v ./...

junit:
	make test
#	go test -v 2>&1 | go-junit-report > report.xml
	go test -json ./... | go-junit-report > report.xml

lint:
	golint ./...


