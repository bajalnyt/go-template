FROM golang:1.21-alpine3.18 as builder

WORKDIR /app

# COPY go.mod .
# RUN go mod download

COPY . .
RUN ls -lR

RUN CGO_ENABLED=0 GOOS=linux go build -o /myapp ./...

# Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=builder /myapp /myapp

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/myapp"]
