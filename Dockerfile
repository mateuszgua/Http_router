FROM golang:1.19-alpine3.16 AS build

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /http_router

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /http_router /http_router

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/http_router"]