FROM golang:latest

# 设置工作目录
WORKDIR /app

# 复制Go应用程序到容器中
COPY . .

# 构建Go应用程序
RUN go build -o main .

# 暴露应用程序的端口
EXPOSE 8080

# 启动应用程序
CMD ["./main"]
