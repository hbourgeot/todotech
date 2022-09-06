FROM golang:alpine AS build

ENV GOPROXY=https://proxy.golang.org

WORKDIR /go/src/todotech
RUN apk update && apk add --no-cache git

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/todotech cmd/web/*

EXPOSE 4000

FROM alpine
RUN apk --no-cache add ca-certificates

WORKDIR /
COPY --from=build /go/bin/todotech /go/bin/todotech
COPY --from=build /go/src/todotech /go/src/todotech

ENTRYPOINT ["go/bin/todotech"]