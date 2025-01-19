FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download \ 
    && go mod verify

COPY . .

RUN apt-get update -y \
    && apt-get install -q -y curl \
    && curl -fsSL https://deb.nodesource.com/setup_lts.x | bash \
    && apt-get update -y \
    && apt-get install nodejs -q -y

#RUN go build -v -o /usr/local/bin/app/ ./...

EXPOSE 8080

CMD ["sleep", "infinity"]
