COVERAGE_FILE=coverage/out

test:
	go test -coverprofile=$(COVERAGE_FILE)

html-coverage:
	go tool cover -html=$(COVERAGE_FILE)