# syntax=docker/dockerfile:1

FROM golang:1.16

ENV PORT 8080
ENV DB_NAME getir.db

WORKDIR /app

COPY ./ ./

RUN go mod download
RUN go build -o ./bin/getir ./cmd/getir/main.go

EXPOSE 8080

CMD [ "./bin/getir" ]