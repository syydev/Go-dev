# 1. Build Container
FROM golang:1.13.5-alpine AS build

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# First add modules list to better utilize caching
COPY go.mod .
COPY go.sum .
# Download dependencies
RUN go mod download

COPY . .
# Build the Go app
RUN go build -o main .


# 2. Runtime Container
FROM alpine

ENV TZ=Asia/Seoul

COPY --from=build /app/main .
COPY --from=build /app/config-container.yaml ./config.yaml

EXPOSE 1323

CMD ["./main"]