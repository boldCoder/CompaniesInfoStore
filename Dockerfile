FROM golang:latest 

WORKDIR /app  

COPY go.mod go.sum ./ 

RUN go mod download

COPY . .  
RUN CD_ENABLED=0 GOOS-linux go build -a -installsuffix cgo -o main .

EXPOSE 9000

RUN go run cmd/main.go 

CMD ["./main"]

