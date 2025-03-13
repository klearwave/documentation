#
# build image
#
FROM golang:1.23-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o service ./pkg/cmd
RUN export GOBIN=/src && go install github.com/pressly/goose/v3/cmd/goose@latest

#
# runtime image
#
FROM alpine AS runtime
RUN apk add --no-cache ca-certificates
COPY --from=build /src/migrations /migrations
COPY --from=build /src/service /bin/service
COPY --from=build /src/goose /bin/goose
ENTRYPOINT ["service"]
