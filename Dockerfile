FROM golang:1.19-alpine3.16 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

COPY . ./

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /http_router

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /http_router /http_router

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/http_router" ]