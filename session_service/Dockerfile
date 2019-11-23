FROM alpine AS base
RUN apk add --no-cache curl wget

FROM golang:1.13 AS go-builder
WORKDIR /go/app
COPY . /go/app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/app/security-service /go/app/session_service/security-service.go

FROM base
COPY --from=go-builder /go/app/security-service /security-service
COPY --from=go-builder /go/app/session-local-settings.json /session-local-settings.json
COPY --from=go-builder /go/app/session-settings.json /session-settings.json
COPY --from=go-builder /go/app/keys /keys
CMD ["/security-service"]
