FROM golang:1.16-alpine AS backendBuild

WORKDIR /build

COPY . .

RUN go mod tidy

RUN go build -o ./bin/bot ./cmd/dom.go

FROM alpine:3 AS final

COPY --from=backendBuild /build/bin/bot .

ENV mode=prod

CMD ["./bot"]