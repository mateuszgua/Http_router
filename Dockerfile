FROM golang:1.19-alpine3.16 AS build

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

COPY . ./

RUN go build -o /http_router

FROM gcr.io/distroless/static-debian11

WORKDIR /http_router

COPY --from=build /http_router /http_router

EXPOSE 8080

USER nonroot:nonroot

CMD ["./http_router"]