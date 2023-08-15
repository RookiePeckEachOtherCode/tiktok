FROM golang:alpine AS builder
ENV GOPROXY=https://goproxy.cn 
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main .
FROM alpine:latest  
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add ffmpeg
RUN mkdir -p static/assets/avatar
COPY ./static/assets/avatar/* /static/assets/avatar/
COPY --from=builder /app/main /main
EXPOSE 8080 
CMD ["/main"]
