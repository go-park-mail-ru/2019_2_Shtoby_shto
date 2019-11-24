FROM golang:1.13

EXPOSE 40000

RUN go get github.com/go-delve/delve/cmd/dlv

COPY . /go/app/src
WORKDIR /go/app/src
# COPY . /go/app

# COPY ./trello-local-settings.json /trello-local-settings.json
# COPY ./trello-settings.json /trello-settings.json

# COPY ./trello-local-settings.json /go/app/src/trello-local-settings.json
# COPY ./trello-settings.json /go/app/src/trello-settings.json

COPY ./trello-local-settings.json .
COPY ./trello-settings.json .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /trello-api ./src
CMD ["/trello-api"]

# CMD ["dlv", "debug", "./src", "--listen=:40000", "--headless=true", "--api-version=2"]
# CMD ["dlv", "exec", "/trello-api", "--listen=:40000", "--headless=true", "--api-version=2"]
# CMD ["dlv", "debug", "/go/app/src/trello-api.go", "--listen=:40000", "--headless=true", "--api-version=2"]
# CMD ["dlv", "exec", "/trello-api"]
# CMD ["ls", "/app/src"]
# CMD ["ls", "app/src"]

# FROM alpine AS base
# RUN apk add --no-cache curl wget

# FROM golang:1.13 AS go-builder
# WORKDIR /go/app
# COPY . /go/app
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/app/trello-api /go/app/src/trello-api.go

# FROM base
# COPY --from=go-builder /go/app/trello-api /trello-api
# COPY --from=go-builder /go/app/trello-local-settings.json /trello-local-settings.json
# COPY --from=go-builder /go/app/trello-settings.json /trello-settings.json
# COPY --from=go-builder /go/app/keys /keys
# CMD ["/trello-api"]
