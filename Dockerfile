FROM golang:1.21-alpine3.18 as builder

WORKDIR /app

COPY . .

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /bin/myapp cmd/api

# minimize image
FROM alpine:latest

# copy only the executable
COPY --from=builder /bin/myapp /bin

# Run the web service on container startup.
CMD ["/bin/myapp"]
