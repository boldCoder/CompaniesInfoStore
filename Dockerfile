FROM golang:latest 

WORKDIR /app  

COPY go.mod ./ 
COPY go.sum ./ 

RUN go mod tidy  

COPY . .  

EXPOSE 9000

RUN go run cmd/main.go 

