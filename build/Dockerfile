FROM golang:1.18.3-alpine3.16 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY config/config-docker.yaml .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /app/earthquake-collector cmd/main.go

FROM scratch
COPY --from=builder /app/config/config-docker.yaml /app/config/config-docker.yaml
COPY --from=builder /app/earthquake-collector /app/earthquake-collector
EXPOSE 8085
ENTRYPOINT ["/app/earthquake-collector"]
