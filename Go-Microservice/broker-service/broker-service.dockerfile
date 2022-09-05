#base go image

# comentei a linha abaixo para não baixar a imagem do docker hub novamente

# FROM golang:1.18-alpine AS builder

# RUN mkdir /app

# COPY . /app

# WORKDIR /app

# RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

# RUN chmod +x /app/brokerApp



# build a tiny docker image

FROM alpine:latest

RUN mkdir /app

# COPY --from=builder /app/brokerApp /app
# comentei a linha acima para não copiar o binário do docker hub novamente

COPY brokerApp /app

CMD [ "/app/brokerApp" ]