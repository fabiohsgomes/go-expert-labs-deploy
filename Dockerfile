FROM golang:1.24-alpine AS build

WORKDIR /app
COPY . .
RUN go mod tidy \
&& GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -C cmd/previsao -o previsao

FROM scratch
WORKDIR /app
ENV AMBIENTE_PUBLICACAO=DEMO
ENV WEATHER_API_KEY=da26cd9b6c624664977234238250506
COPY --from=build /app/cmd/previsao/previsao previsao
COPY --from=build /etc/ssl/certs/ /etc/ssl/certs/
ENTRYPOINT ["./previsao"]