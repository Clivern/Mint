FROM golang:1.27 as build

RUN mkdir -p /app
RUN apt-get update

WORKDIR /app

COPY ./ ./

RUN go build -v -ldflags="-X 'main.version=v1.0.0' -X 'main.builtBy=docker'" mint.go

FROM ubuntu:20.04

RUN mkdir -p /app/configs
RUN mkdir -p /app/var/logs
RUN apt-get update
RUN apt-get install curl -y

WORKDIR /app

COPY --from=build /app/mint /app/mint
COPY --from=build /app/config.dist.yml /app/configs/config.dist.yml

EXPOSE 8000

VOLUME /app/configs
VOLUME /app/var

RUN ./mint version

CMD ["./mint", "server", "-c", "/app/configs/config.dist.yml"]
