dev:
	cd api && $(MAKE) dev

fmt:
	cd api && $(MAKE) fmt

lint:
	cd api && $(MAKE) lint $(filter-out $@,$(MAKECMDGOALS))

lintf:
	cd api && $(MAKE) lintf

create-schema:
	cd api && $(MAKE) create-schema

generate:
	cd api && $(MAKE) generate

generate-clean:
	cd api && $(MAKE) generate-clean

build:
	cd api && $(MAKE) build

start:
	cd api && $(MAKE) start

push:
	cd api && $(MAKE) push

test:
	cd api && $(MAKE) test