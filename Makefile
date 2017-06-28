install:
	go install -v

image:
	docker build -t cirocosta/nfsvol .

deps:
	glide install

fmt:
	gofmt -s -w main.go driver.go


.PHONY: install image deps fmt
