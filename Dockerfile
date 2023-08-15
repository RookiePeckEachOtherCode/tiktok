FROM golang:alpine AS builder
ENV GOPROXY=https://goproxy.cn 
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
FROM alpine:latest  
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add ffmpeg
COPY --from=builder /app/main /main
EXPOSE 8080 
CMD ["/main"]
