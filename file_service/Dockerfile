FROM alpine AS base
RUN apk add --no-cache curl wget

FROM golang:1.13 AS go-builder
WORKDIR /go/app
COPY . /go/app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/app/file-service /go/app/file_service/file-service.go

FROM base
COPY --from=go-builder /go/app/file-service /file-service
COPY --from=go-builder /go/app/file_service/file-local-settings.json /file-local-settings.json
COPY --from=go-builder /go/app/file_service/file-settings.json /file-settings.json
CMD ["/file-service"]
