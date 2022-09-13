.PHONY: install
install:
	go install .

.PHONY: gen
gen:
	go generate ./...

.PHONY: test
test:
	go test -v -mod=vendor ./...
