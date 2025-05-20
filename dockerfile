FROM golang:1.24.1

# RUN apt-get update && apt-get install -y make

WORKDIR /app
COPY . .

# RUN go install github.com/air-verse/air@latest
RUN go mod tidy
RUN go build -o main .

EXPOSE 3000
CMD ["/main"]
