# Step 1 Build the binary for the microservice
FROM golang:alpine AS builder

# Install 
RUN apk update && apk add --no-cache git gcc musl-dev

RUN mkdir /build 
WORKDIR /build 

ADD go.mod /build
ADD go.sum /build

RUN go mod download

ADD . /build/

# Build the binary.
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main /build/cmd/userservice

# Step 2 package the binary in a minimal container
FROM alpine

COPY --from=builder /build/main /app/
WORKDIR /app

# Run the binary.
CMD ["./main"]