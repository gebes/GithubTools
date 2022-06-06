FROM golang:latest AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d ./...

# Install the package
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/backend

FROM alpine:latest AS runtime
WORKDIR /app
COPY --from=builder /app/app /app/app
COPY --from=builder /app/.env /app/.env

CMD ["./app"]