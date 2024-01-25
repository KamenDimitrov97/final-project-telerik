# Use Fedora 39 as the base image
FROM fedora:39 as base

RUN dnf update -y

RUN dnf install -y golang

FROM base

WORKDIR /app

COPY . .

RUN go build -o main

EXPOSE 3333

# Command to run the Go application
CMD ["./main"]