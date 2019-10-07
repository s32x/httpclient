deps:
	-rm -rf ./vendor go.mod go.sum
	go mod init
	go mod vendor
	
test:
	go test ./...