install:
	go install -v

image:
	docker build -t cirocosta/nfsvol .

deps:
	glide install


.PHONY: install image deps
