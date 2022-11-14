FROM golang:1.19.3-alpine3.16 AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/swapi-app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod tidy
COPY . .

RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/swapi-app cmd/main/main.go

FROM alpine:3.16
RUN apk add ca-certificates

ENV DB_NAME dummy-db
ENV DB_USER dummy-user
ENV DB_PASSWORD dummy-password

COPY --from=build_base /tmp/swapi-app/out/swapi-app /app/swapi-app
EXPOSE 8080

CMD ["/app/swapi-app"]