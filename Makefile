MAIN_PKG = "./cmd/web"

unit-test:
	@go test -v ${MAIN_PKG} ./...

test-coverage:
	@go test -coverprofile=coverage.out ${MAIN_PKG} ./...
	@go tool cover -html=coverage.out

test: unit-test test-coverage

run:
	@go run ${MAIN_PKG}