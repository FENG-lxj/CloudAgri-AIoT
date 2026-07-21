# 🌾 CloudAgri（云端田园）— 智慧农业管理平台

一个完整的智慧农业物联网（IoT）管理平台，集成了设备监控、AI 分析、视觉检测、作物产量预测与作物推荐等功能。采用前后端分离架构，适用于现代化智能农业场景。

---

## 📋 项目概述

CloudAgri 面向现代化农业需求，提供从设备数据采集、环境监测到 AI 智能决策的全链路解决方案。系统支持：

- 🌡️ 农业生态环境（AE）多指标实时监测（光照、pH、温度、湿度、氮磷钾、空气质量等）
- 🏭 粮仓环境（AS）监测（温度、湿度、CO₂、O₂）
- 📹 视频监控 + YOLO 视觉检测（病虫害识别、人物入侵、火灾检测）
- 🤖 AI 农业生产建议（对接大语言模型，基于实时环境数据生成管理建议）
- 🌱 作物产量预测（XGBoost 机器学习模型，支持 43 种作物）
- 🔍 智能作物推荐（基于土壤和环境特征，匹配最适合种植的作物）
- 📡 MQTT 协议设备通信（IoT 设备数据采集与控制）
- 📊 大数据可视化看板（历史数据趋势、异常告警）
- 👤 用户管理 + JWT 认证 + 超级管理员权限

---

## 🏗️ 系统架构

```
┌──────────────────────────────────────────────────────────────────┐
│                        前端 (PastoralCloud)                       │
│                      Vue 2 + Element UI + ECharts                 │
│                     http://localhost:8090                         │
└──────────────────────────┬───────────────────────────────────────┘
                           │ HTTP / REST API
                           ▼
┌──────────────────────────────────────────────────────────────────┐
│                   主后端服务 (IoTA_WebServer)                       │
│                      Go + Gin + GORM v2                           │
│                    http://localhost:20201                         │
│  ┌──────────┬───────────┬────────────┬──────────┬─────────────┐  │
│  │用户管理  │ AI推理服务 │ 大数据分析  │ 文件上传  │ JWT认证鉴权  │  │
│  └──────────┴───────────┴────────────┴──────────┴─────────────┘  │
└──────┬──────────┬──────────────┬──────────────┬──────────────────┘
       │          │              │              │
       ▼          ▼              ▼              ▼
┌──────────┐ ┌────────┐ ┌──────────┐ ┌──────────────┐
│  MySQL   │ │ Redis  │ │ InfluxDB │ │ 文件存储(本地) │
│ (关系型) │ │ (缓存) │ │ (时序库) │ │              │
└──────────┘ └────────┘ └──────────┘ └──────────────┘

┌──────────────────────────────────────────────────────────────────┐
│                   IoT设备服务 (IOTA_ZZ_server)                     │
│                      Go + MQTT + InfluxDB                         │
│  订阅 MQTT 主题 → 解析设备数据 → 存入 InfluxDB + MySQL            │
└──────────────────────────────────────────────────────────────────┘
                           ▲
                           │ MQTT 协议
┌──────────────────────┬───────────────────────────────────────────┐
│  视觉检测服务          │  AI 分析服务 (YUCE)                        │
│  (KONG/PastoralCloud) │  Python Flask + XGBoost + LLM              │
│  Python + YOLO +      │  ├─ Productionapp.py: 作物产量预测          │
│  OpenCV + FFmpeg      │  ├─ crop_recommender.py: 作物推荐           │
│  ├─ YOLO目标检测      │  └─ tianqi.py: 农业专家AI问答               │
│  ├─ RTMP视频推流      │                                            │
│  ├─ MQTT日志上报      │                                            │
│  └─ 检测图片上传API   │                                            │
└──────────────────────┴───────────────────────────────────────────┘
```

---

## 📁 项目目录结构

```
CloudAgri/
├── PastoralCloud/            # 🌐 前端 — Vue.js 2
│   ├── src/                  # 源代码
│   │   ├── AI_Agro/          # AI 农业助手模块
│   │   ├── Device_Management_Page/  # 设备管理页面
│   │   ├── Storage/          # 数据存储看板
│   │   ├── SuperAdmin/       # 超级管理员页面
│   │   └── img/              # 静态图片资源
│   ├── public/               # 静态资源
│   ├── package.json          # 依赖配置
│   ├── vue.config.js         # Vue CLI 配置
│   └── babel.config.js       # Babel 配置
│
├── IoTA_WebServer/           # 🔧 主后端服务 — Go (Gin)
│   ├── app/
│   │   ├── global/           # 全局变量与常量
│   │   ├── http/             # HTTP 中间件与验证器
│   │   ├── model/            # 数据模型 (MySQL/InfluxDB)
│   │   ├── service/          # 业务逻辑层
│   │   └── utils/            # 工具函数
│   ├── bootstrap/            # 应用引导初始化
│   ├── cmd/web/              # Web 服务入口
│   ├── config/               # 配置文件 (YAML)
│   ├── routers/              # 路由定义
│   ├── storage/              # 日志与上传文件存储
│   └── makefile              # 编译脚本
│
├── IOTA_ZZ_server/           # 📡 设备数据采集服务 — Go
│   ├── config/               # 配置文件
│   ├── mapper/               # 数据持久化层
│   ├── model/                # 数据模型
│   ├── pkg/                  # 基础设施包
│   │   ├── MyMQTT/           # MQTT 客户端
│   │   ├── InfluxDB/         # InfluxDB 客户端
│   │   └── MySQL/            # MySQL 客户端
│   ├── service/              # MQTT 消息处理服务
│   └── main.go               # 入口文件
│
├── YUCE/                     # 🤖 AI 预测推荐服务 — Python
│   ├── Productionapp.py      # 作物产量预测模型
│   ├── crop_recommender.py   # 作物智能推荐
│   ├── tianqi.py             # 农业专家 AI 问答
│   ├── combined_server.py    # 多服务进程管理
│   ├── crop_models/          # XGBoost 模型文件 (43种作物)
│   └── crop_data_china/      # 中国作物种植特征数据
│
├── KONG/                     # 📹 视觉检测推流服务 — Python
│   └── PastoralCloud/
│       ├── main.py           # 主程序 (YOLO检测+推流+MQTT)
│       └── config/           # 配置文件
│
├── iota.sql                  # 🗄️ 数据库初始化脚本
└── .gitignore                # Git 忽略规则
```

---

## 🚀 快速开始

### 环境要求

| 组件 | 版本/说明 |
|------|----------|
| Go | ≥ 1.19（IOTA_ZZ_server）, ≥ 1.23（IoTA_WebServer） |
| Python | ≥ 3.8 |
| Node.js | ≥ 14（推荐 16+） |
| MySQL | ≥ 5.7 |
| Redis | ≥ 6.0 |
| InfluxDB | ≥ 2.0 |
| MQTT Broker | EMQX / Mosquitto |
| FFmpeg | ≥ 4.0（视频检测模块） |
| 摄像头 | 视频检测模块需要摄像头设备 |
| YOLO 模型 | YOLOv8 训练权重文件 |

### 1. 数据库初始化

```bash
# 导入 SQL 脚本到 MySQL
mysql -u root -p < iota.sql
```

### 2. 前端启动 (PastoralCloud)

```bash
cd PastoralCloud
npm install
npm run serve
# 开发服务器运行在 http://localhost:8090
```

### 3. 主后端服务启动 (IoTA_WebServer)

```bash
cd IoTA_WebServer

# 修改配置文件
# config/config.yml     — 应用配置（JWT密钥、端口、Redis等）
# config/gorm_v2.yml    — 数据库连接配置

# 安装依赖并编译
go mod tidy
make build-web

# 或直接运行
go run ./cmd/web/main.go
# 服务运行在 http://localhost:20201
```

### 4. IoT 设备服务启动 (IOTA_ZZ_server)

```bash
cd IOTA_ZZ_server

# 修改 config/config.yaml 中的 MQTT/InfluxDB/MySQL 连接信息
# 确保 MQTT Broker 已启动

go run main.go
```

### 5. AI 预测推荐服务启动 (YUCE)

```bash
cd YUCE

# 安装依赖
pip install flask flask-cors xgboost numpy pandas openai dashscope

# 启动所有服务（产量预测 + 作物推荐）
python combined_server.py

# 或分别启动各服务
python Productionapp.py       # 产量预测 → 端口 8901
python crop_recommender.py    # 作物推荐 → 端口 8900
python tianqi.py              # AI问答   → 端口 4999
```

### 6. 视觉检测服务启动 (KONG)

```bash
cd KONG/PastoralCloud

# 安装依赖
pip install ultralytics opencv-python paho-mqtt requests

# 安装 YOLO 依赖
# 下载 YOLO 模型权重文件到运行目录
# 配置 main.py 中的参数（摄像头索引、RTMP推流地址、API地址等）

python main.py
```

---

## 🔌 API 接口概览

### Web 服务 (IoTA_WebServer, port 20201)

| 模块 | 路径 | 说明 |
|------|------|------|
| 用户 | `POST /IoTA/users/login` | 手机验证码登录 |
| 用户 | `POST /IoTA/users/captcha` | 发送手机验证码 |
| 用户 | `POST /IoTA/users/refreshtoken` | 刷新 JWT Token |
| 超级管理员 | `POST /IoTA/SuperAdmin/SuperAdminAddUser` | 添加管理用户 |
| 超级管理员 | `GET /IoTA/SuperAdmin/GetUserList` | 获取用户列表 |
| AI | `GET /IoTA/ai/ProposalAI_AE` | AI 农业环境分析建议 |
| AI | `GET /IoTA/ai/ProposalAI_AS` | AI 粮仓环境分析建议 |
| AI | `GET /IoTA/ai/AgricultureExpertAI` | AI 农业专家问答 |
| AI | `GET /IoTA/ai/DetectionInfo` | 检测信息查询 |
| AI | `GET /IoTA/ai/DetectionImage` | 检测图片获取 |
| AI | `GET /IoTA/ai/GetMonitorLogs` | AI 监控日志 |
| AI | `POST /IoTA/users/AiImageInfo` | AI检测图片上传 |
| 大数据 | `GET /IoTA/bigdata/GetAEdata` | 农业环境时序数据 |
| 大数据 | `GET /IoTA/bigdata/GetASdata` | 粮仓环境时序数据 |
| 大数据 | `GET /IoTA/bigdata/GetDevInfo` | 设备信息列表 |
| 大数据 | `GET /IoTA/bigdata/GetMemCPU` | 服务器资源监控 |
| 文件 | `POST /IoTA/upload/files` | 文件上传 |

### AI 服务 (YUCE)

| 服务 | 端口 | 路径 | 说明 |
|------|------|------|------|
| 产量预测 | 8901 | `POST /predict` | 根据土壤环境数据预测作物产量 |
| 产量预测 | 8901 | `GET /health` | 健康检查 |
| 作物推荐 | 8900 | `POST /recommend_crop` | 根据土地特征推荐最适作物 |
| 作物推荐 | 8900 | `GET /health` | 健康检查 |
| 作物推荐 | 8900 | `GET /available_crops` | 获取可用作物列表 |
| AI问答 | 4999 | `GET /IoTA/ai/AgricultureExpertAI` | 农业专家AI问答 |

---

## 🛠️ 技术栈详情

### 前端
- **Vue 2** — 渐进式 JavaScript 框架
- **Element UI** — 桌面端组件库
- **ECharts** — 数据可视化图表
- **Axios** — HTTP 客户端
- **Vue Router** + **Vuex** — 路由与状态管理
- **MQTT.js** — 浏览器端 MQTT 客户端
- **flv.js** — FLV 视频流播放
- **高德地图 API** — 地图定位与展示

### 后端 (Go)
- **Gin** — 高性能 HTTP Web 框架
- **GORM v2** — Go ORM 框架
- **Viper** — 配置管理（支持热加载）
- **Zap + Lumberjack** — 结构化日志 + 日志切割
- **dgrijalva/jwt-go** — JWT 令牌管理
- **go-playground/validator** — 请求参数验证
- **InfluxDB Client v2** — 时序数据库客户端
- **Paho MQTT** — MQTT 客户端
- **OpenAI Go SDK** — AI API 调用

### 后端 (Python)
- **Flask** — 轻量级 Web 框架
- **XGBoost** — 梯度提升树（产量预测模型）
- **OpenCV** — 计算机视觉
- **YOLOv8 (Ultralytics)** — 目标检测
- **Paho MQTT** — MQTT 客户端
- **FFmpeg** — 视频编解码与 RTMP 推流
- **阿里云百炼 (DashScope)** — 大语言模型 API

### 基础设施
- **MySQL** — 关系型数据存储
- **InfluxDB** — 时序数据存储（设备传感器数据）
- **Redis** — 缓存服务（Token 缓存、API Key 缓存）
- **MQTT (EMQX/Mosquitto)** — IoT 设备消息协议
- **Nginx + RTMP Module** — 视频流媒体服务器

---

## ⚙️ 配置说明

所有配置文件中的敏感信息（密码、Token、API Key 等）均使用占位符 `your-xxx-here`，部署时请替换为实际值：

### 需要配置的文件

| 文件 | 需要修改的配置项 |
|------|-----------------|
| `IoTA_WebServer/config/config.yml` | JWT签名密钥、Redis密码、InfluxDB Token、超级管理员手机号 |
| `IoTA_WebServer/config/gorm_v2.yml` | MySQL 数据库连接信息（Host/User/Pass） |
| `IOTA_ZZ_server/config/config.yaml` | MQTT Broker地址/用户名/密码、InfluxDB Token、MySQL密码 |
| `YUCE/Productionapp.py` | 阿里云百炼 API Key（`your-api-key-here`） |
| `YUCE/tianqi.py` | 阿里云百炼 API Key + App ID |
| `KONG/PastoralCloud/main.py` | JWT Token、API_URL、RTMP_URL、MODEL_PATH、摄像头参数 |

### 推荐配置方式

生产环境建议使用环境变量或外部配置中心管理敏感信息，而非硬编码在配置文件中。

---

## 🚧 部署建议

### 生产环境部署

1. **编译 Go 后端**
   ```bash
   cd IoTA_WebServer
   make build-web
   # 生成二进制文件 IoTA_WebServer
   ```

2. **构建前端**
   ```bash
   cd PastoralCloud
   npm run build
   # 产物在 PastoralCloud/dist/ 目录
   ```

3. **使用 Nginx 反向代理**，统一端口访问前端静态资源和后端 API

4. **配置 HTTPS**（推荐使用 Let's Encrypt 免费证书）

5. **使用 systemd/supervisord** 管理 Python/Go 服务进程

6. **设置防火墙**，仅开放必要的对外端口

### 推荐服务器配置
- CPU：4 核以上
- 内存：8GB 以上（含 YOLO 推理时建议 16GB+）
- 存储：50GB+（含时序数据库存储）
- GPU：视觉检测服务建议 NVIDIA GPU（可选）

---

## 📝 许可证

MIT License

---

## 🤝 贡献

欢迎提交 Issue 和 Pull Request。

---

## 📧 联系方式

如有问题或合作需求，请通过 Issue 或邮件 3265605854@qq.com 联系。

---

*Developed for modern intelligent agricultural IoT.*
