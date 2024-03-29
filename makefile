.PHONY: test
test:
# -timeout 10m
	go test -v ./...

.PHONY: gen
gen: gen-go

.PHONY: gen-go
gen-go:
	go mod tidy
	go generate ./...

