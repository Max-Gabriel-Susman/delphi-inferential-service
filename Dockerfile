# syntax=docker/dockerfile:1

# Build the application from source
FROM --platform=linux/amd64 golang:1.19 AS build-stage

ENV GO111MODULE=on
ENV API_ADDRESS=0.0.0.0:8082 
ENV INFERENTIAL_DB_USER=usr 
ENV INFERENTIAL_DB_PASSWORD=identity 
ENV INFERENTIAL_DB_HOST=127.0.0.1 
ENV INFERENTIAL_DB_NAME=identity 
ENV INFERENTIAL_DB_PORT=3306 
ENV ENABLE_MIGRATE=true 

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# COPY *.go ./
COPY . . 

RUN go mod download
# RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o /delphi-inferential-service

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /delphi-inferential-service /delphi-inferential-service

EXPOSE 50051

USER nonroot:nonroot

ENTRYPOINT ["/delphi-inferential-service"]