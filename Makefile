dev:
	go run main.go

fmt:
	go tool gofumpt -l -w .
	go tool golines -m 80 -w .
	go tool goimports -l -w .

create-schema:
ifndef name
	$(error name is required. Usage: make create-schema name=SchemaName)
endif
	go run entgo.io/ent/cmd/ent new $(name)

generate:
	rm -rf ent/generated && go generate ./ent/...

build:
	rm -rf bin && go build -o bin/api
