FROM golang:1.13.7 as builder

WORKDIR /go/src/github.com/shinjiezumi/vue-go-samples/src/api
COPY . .

# Build the command inside the container.
RUN go get -u github.com/codegangsta/gin \
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure \
  && CGO_ENABLED=0 GOOS=linux go build -v -o api

# Use a Docker multi-stage build to create a lean production image.
FROM alpine

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/github.com/shinjiezumi/vue-go-samples/src/api/api /api

# set ENV
ENV GIN_MODE="release"

# Run the web service on container startup.
CMD ["/api"]