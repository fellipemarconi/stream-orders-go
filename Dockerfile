FROM golang:1.24.4-alpine AS builder

# Argumento para definir qual serviço buildar (api ou consumer)
ARG SERVICE=api

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

# Build do serviço especificado
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/${SERVICE}

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]