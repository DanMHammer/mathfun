FROM golang:alpine as builder
RUN apk --no-cache add ca-certificates git
WORKDIR /build/src

# Fetch dependencies
COPY src/go.mod ./
RUN go mod download

# # Test
COPY src/. ./
RUN CGO_ENABLED=0 GOOS=linux GARCH=amd64 go test ./...

# Build
RUN CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build

# Create final image
FROM alpine
WORKDIR /root
COPY --from=builder /build/src/ .
EXPOSE 3000
ENTRYPOINT ["./mathfun"]