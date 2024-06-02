# 后端文档

这里是后端代码文件夹。

## 文件结构

```bash
.
├─api                  # API 版本目录
│  └─v1               # v1 版本 API
│      ├─letstrans    # letstrans 模块 API
│      └─system       # system 模块 API
├─config               # 配置文件目录
├─core                 # 核心功能目录
│  └─internal          # 内部核心功能
├─global               # 全局变量和设置
│  └─internal          # 全局内部设置
├─initialize           # 初始化脚本和配置
│  └─internal          # 内部初始化配置
├─middleware           # 中间件目录
├─model                # 数据模型目录
│  ├─common            # 通用数据模型
│  │  ├─request        # 通用请求模型
│  │  └─response       # 通用响应模型
│  ├─example           # 示例数据模型
│  │  └─response       # 示例响应模型
│  ├─letstrans         # letstrans 模块数据模型
│  │  ├─request        # letstrans 请求模型
│  │  └─response       # letstrans 响应模型
│  └─system            # system 模块数据模型
│      ├─request       # system 请求模型
│      └─response      # system 响应模型
├─python_src           # Python 源码目录
│  └─document_processor # 文档处理器
├─router               # 路由配置目录
│  ├─letstrans         # letstrans 模块路由
│  └─system            # system 模块路由
├─service              # 服务层目录
│  ├─letstrans         # letstrans 模块服务
│  └─system            # system 模块服务
├─test_case            # 测试用例目录
├─uploads              # 上传文件目录
└─utils                # 工具类目录
    ├─document         # 文档工具
    ├─plugin           # 插件工具
    ├─timer            # 计时工具
    ├─translate        # 翻译工具
    └─upload           # 上传工具

```

| 文件夹             | 说明                | 描述                               |
| ------------------ | ------------------- | ---------------------------------- |
| .env               | 环境变量文件        | 存储项目的环境变量                 |
| .env.backup        | 环境变量备份文件    | 环境变量的备份版本                 |
| config-docker.yaml | Docker 配置文件     | Docker 容器的配置文件              |
| config.yaml        | 常规配置文件        | 项目的常规配置文件                 |
| docker-compose.yml | Docker Compose 文件 | 定义多容器 Docker 应用的服务       |
| Dockerfile         | Docker 镜像文件     | 包含创建 Docker 镜像的指令         |
| main.go            | 主程序文件          | 项目的主入口文件                   |
| README.md          | 项目说明文件        | 提供项目的基本信息和使用说明       |
| api/v1/letstrans   | letstrans 模块 API  | 包含 letstrans 模块相关的 API 文件 |
| api/v1/system      | system 模块 API     | 包含 system 模块相关的 API 文件    |
| config             | 配置文件目录        | 存储项目的配置文件                 |
| core               | 核心功能目录        | 包含项目的核心功能代码             |
| global             | 全局变量和设置      | 定义全局变量和设置                 |
| initialize         | 初始化脚本和配置    | 初始化项目所需的脚本和配置         |
| middleware         | 中间件目录          | 存放中间件代码                     |
| model              | 数据模型目录        | 定义数据模型的目录                 |
| python_src         | Python 源码目录     | 主要是文档处理服务器代码           |
| router             | 路由配置目录        | 定义项目的路由配置                 |
| service            | 服务层目录          | 存放项目的服务层代码               |
| test_case          | 测试用例目录        | 存放测试用例文件                   |
| uploads            | 上传文件目录        | 存放上传的文件                     |
| utils              | 工具类目录          | 存放各种工具类代码                 |

## 运行配置

### 密钥配置

复制备份的环境变量文件并进行修改，填入自己的火山引擎 AK 和 SK。此配置将在 Docker Compose 启动时被读取并使用。

```bash
cp .env.backup .env
```

### 数据库配置

本地需要安装 PgSQL，并建立数据库`lets_trans`。

用户名和密码在 `config.yaml`进行配置。

### 一键启动

```bash
docker compose up -d
```

