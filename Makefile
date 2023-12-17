include compose/.env
include variables.mk

docker-image:
	make -C docker docker-image

run:
	SERVICEHOST=$(localhost) \
	SERVICEPORT=$(port) \
	$(GORUN) cmd/main.go

.PHONY: run