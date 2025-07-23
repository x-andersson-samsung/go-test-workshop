coverage:
	go test -coverprofile=coverage.out -v .
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html