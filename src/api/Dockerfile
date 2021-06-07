FROM golang:1.15.12 as builder

WORKDIR /go/src/
COPY . .

# Build the command inside the container.
RUN go get -u github.com/codegangsta/gin \
  && CGO_ENABLED=0 GOOS=linux go build -v -o api

# Use a Docker multi-stage build to create a lean production image.
FROM alpine

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/api /api

# set ENV
ENV GIN_MODE="release"

# Run the web service on container startup.
CMD ["/api"]