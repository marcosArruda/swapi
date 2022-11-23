# syntax = docker/dockerfile:1-experimental
FROM golang:1.19.3-alpine3.16 AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/swapi-app

COPY go.* .
RUN go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build go build -o ./out/swapi-app ./cmd/main/main.go

# ===== UNCOMMENT THE FOLLOWING LINE TO ALSO RUN UNIT TESTS AFTER THE BUILD =====
#RUN --mount=type=cache,target=/root/.cache/go-build go test -v -coverpkg=./... -coverprofile=profile.cov ./...; go tool cover -func profile.cov

FROM alpine:3.16
RUN apk add ca-certificates

# RUN addgroup -S swappgroup && adduser -S swuser  -u 9999 -G swappgroup
USER guest

ENV DB_NAME dummy-db
ENV DB_USER dummy-user
ENV DB_PASSWORD dummy-password
ENV DB_HOSTPORT dummydb:3306

COPY --from=build_base --chown=guest /tmp/swapi-app/out/swapi-app /app/swapi-app
EXPOSE 8080

CMD ["/app/swapi-app"]