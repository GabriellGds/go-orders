unit-tests:
	go test -v ./...
all-tests:
	go test -v -tags=integration ./...