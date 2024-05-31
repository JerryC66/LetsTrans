# 后端文档

这里是后端代码文件夹。

## 文件结构

```shell

```

| 文件夹          | 说明         | 描述                     |
|--------------|------------|------------------------|
| `controller` | API 层      |                        |
| `--v1`       | V1 版本的 API |                        |
| `config`     | 配置包        | config.go 是综合体，其他是配置单元 |
| `global`     | 全局变量       | 数据库客户端                 |
|              |            |                        |

## 密钥配置

通过 `.env` 文件配置密钥。格式同 `.env.backup`。

AK 和 SK 分别是火山引擎的 Access Key 和 Secret Key。这一配置会在 docker compose 启动时读取。