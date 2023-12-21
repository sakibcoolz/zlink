include compose/.env
include variables.mk

docker-image:
	make -C docker docker-image

run:
	SERVICEHOST=$(localhost) \
	VERSION=$(version) \
	SERVICEPORT=$(port) \
	$(GORUN) cmd/main.go

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(build) ./cmd/

docker-start:
	make -C compose start

docker-stop:
	make -C compose stop


.PHONY: run