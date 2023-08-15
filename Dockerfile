FROM golang:alpine AS builder
ENV GOPROXY=https://goproxy.cn 
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main /main
EXPOSE 8080 
CMD ["/main"]
