FROM fedora:39 as base

WORKDIR /app

RUN dnf update -y

RUN dnf install -y golang

COPY . .

RUN make build-bin

FROM ubuntu:23.10

RUN useradd -u 5000 server

WORKDIR /app

COPY --from=base app/build /app/

EXPOSE 3333

USER server:server

CMD ["./main"]
