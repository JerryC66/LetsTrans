# 第一阶段：构建阶段
FROM golang:1.22.1 AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

#CMD ["./myapp"]

## 第二阶段：运行阶段
FROM alpine:latest

# 设置工作目录
WORKDIR /

# 从构建器镜像中复制二进制文件到基础镜像
COPY --from=builder /app/server .
COPY --from=builder /app/config.yaml .

# 添加运行时所需的任何额外库
RUN apk --no-cache add ca-certificates

# 声明运行时容器提供服务的端口
EXPOSE 8098

# 指定容器启动时运行的命令
ENTRYPOINT ./server