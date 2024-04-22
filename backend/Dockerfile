# 第一阶段：构建阶段
FROM golang:1.22.1 AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod ./
COPY go.sum ./
ENV GOPROXY=https://goproxy.cn
# 下载依赖
RUN  go mod download

# 复制源代码文件
COPY . .

# 编译项目，创建二进制文件
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp

#CMD ["./myapp"]

## 第二阶段：运行阶段
FROM alpine:latest

# 设置工作目录
WORKDIR /

# 从构建器镜像中复制二进制文件到基础镜像
COPY --from=builder /app/myapp .

# 添加运行时所需的任何额外库
RUN apk --no-cache add ca-certificates

# 声明运行时容器提供服务的端口
EXPOSE 8080

# 指定容器启动时运行的命令
CMD ["./myapp"]