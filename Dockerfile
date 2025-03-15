FROM golang:1.24

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

EXPOSE 8080