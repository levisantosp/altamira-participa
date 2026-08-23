dev:
	go run ./api

fmt:
	find . -path './pg' -prune -o -type f -name '*.go' -exec go tool gofumpt -w {} +
	find . -path './pg' -prune -o -type f -name '*.go' -exec go tool golines -m 80 -w {} +
	find . -path './pg' -prune -o -type f -name '*.go' -exec go tool goimports -w {} +

create-schema:
ifndef name
	$(error name is required. Usage: make create-schema name=SchemaName)
endif
	go tool ent new $(name)

generate:
	go generate ./api/ent/...

generate-clean:
	rm -rf api/ent/generated && go generate ./api/ent/...

build:
	rm -rf bin && go build -o bin/api ./api

start:
	bin/api

push:
	go run ./api/cmd/push
