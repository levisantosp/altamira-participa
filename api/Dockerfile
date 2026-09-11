FROM golang:1.27.0-alpine AS build

WORKDIR /app

RUN apk add --no-cache make

COPY go.work go.work.sum /app/
COPY api/go.mod api/go.sum /app/api/

RUN go mod download

COPY . .

RUN make build

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/bin/api ./bin/api

ENTRYPOINT [ "./bin/api" ]
