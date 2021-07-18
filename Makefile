GOCMD=GO111MODULE=on go

lint:
	golangci-lint run

test:
	$(GOCMD) test -cover -race ./...

bench:
	$(GOCMD) test -bench=. -benchmem ./...

.PHONY: test bench lint