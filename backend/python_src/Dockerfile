# 使用官方 Python 镜像作为基础镜像
FROM python:3.9-slim

# 设置工作目录
WORKDIR /app

# 复制 requirements.txt 并安装 Python 依赖
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# 复制应用程序代码
COPY  document_processor/ /app/

# 暴露应用程序运行的端口（根据实际应用端口号调整）
EXPOSE 5000
RUN echo "=== Checking /app contents ===" && ls -al /app

# 设置容器启动时运行的命令（根据你的 Python 应用启动方式调整）
CMD ["python", "doc_server.py"]