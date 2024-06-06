# 使用官方 Node.js 20 镜像
FROM node:20

# 设置工作目录
WORKDIR /app

# 全局安装 pnpm
RUN npm install -g pnpm

# 复制项目文件到工作目录
COPY . .

# 安装项目依赖
RUN pnpm install

# 构建项目
RUN pnpm build

# 由于只是构建，不需要启动服务，所以不设置 CMD 或 ENTRYPOINT
# docker build -t my-frontend-build .
# docker run --rm -v $(pwd)/dist:/app/dist my-frontend-build
