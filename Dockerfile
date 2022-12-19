FROM golang:1.19-alpine3.16 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

COPY . ./

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /to-do-list

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /to-do-list /to-do-list

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/to-do-list" ]
