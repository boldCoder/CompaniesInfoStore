FROM golang:alpine 

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# RUN mkdir /app
WORKDIR /app  

COPY go.mod go.sum ./ 

ENV GO111MODULE=on

RUN go mod download

COPY . ./

RUN go build -o main ./cmd/ 

EXPOSE 9000

CMD ["/main"]

