FROM golang:alpine AS builder

WORKDIR /user/local/src

EXPOSE 80

RUN apk --no-cache add bash

COPY ["app/go.mod", "app/go.sum", "./"]

RUN go mod download

COPY app ./

RUN go build -o ./bin/app cmd/main.go

FROM alpine AS runner

COPY --from=builder /user/local/src/bin/app .

CMD ["/app"]