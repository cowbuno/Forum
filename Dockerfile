FROM golang:1.20

WORKDIR /web

COPY . .
RUN go build -o web ./cmd/web
EXPOSE 8080
ENTRYPOINT ["./web"]
