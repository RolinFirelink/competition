# 每日一灯博客

一个轻量级的个人博客系统，采用Vue.js开发，具有简洁的Typora风格设计。

## 项目概述

这是一个专为格斗游戏视频解说、IT开发、街霸6视频剪辑分享等内容创作者设计的个人博客系统。

## 功能特性

### 前台功能 (/user)
- **关于我**: 展示头像、个人签名和专注领域
- **特别鸣谢**: 展示鸣谢对象信息
- **对局投稿**: 提供对局投稿表单，支持三种对局类型
- **交流群**: 展示微信群和QQ群信息，支持图片点击放大
- **支持主包**: 展示支持主播的相关信息
- **粉丝福利**: 展示粉丝福利政策
- **必看视频**: 展示重要视频链接列表

### 后台功能 (/meiriyidengdehoutai)
- **投稿管理**: 按时间倒序展示所有投稿对局数据

## 技术栈

- **前端框架**: Vue.js 3.x
- **路由**: Vue Router 4.x
- **样式**: 原生CSS，采用Typora风格设计
- **数据管理**: Vue的响应式数据管理，无后端依赖

## 项目结构

```
light_blog_real/
├── index.html                 # 主页面
├── package.json              # 项目依赖
├── vite.config.js            # Vite配置
├── src/
│   ├── main.js               # 应用入口
│   ├── App.vue               # 根组件
│   ├── router/
│   │   └── index.js          # 路由配置
│   ├── components/
│   │   ├── UserFront.vue     # 前台页面组件
│   │   ├── AdminBack.vue     # 后台页面组件
│   │   ├── AboutMe.vue       # 关于我组件
│   │   ├── Acknowledgments.vue # 特别鸣谢组件
│   │   ├── SubmissionForm.vue # 对局投稿组件
│   │   ├── Community.vue     # 交流群组件
│   │   ├── Support.vue       # 支持主包组件
│   │   ├── FanBenefits.vue   # 粉丝福利组件
│   │   ├── MustWatch.vue     # 必看视频组件
│   │   └── ImageModal.vue    # 图片模态框组件
│   └── assets/
│       ├── styles/
│       │   └── main.css      # 全局样式
│       └── images/           # 静态图片资源
└── README.md                 # 项目说明文档
```

## 安装和运行

1. 安装依赖：
```bash
npm install
```

2. 启动开发服务器：
```bash
npm run dev
```

3. 构建生产版本：
```bash
npm run build
```

## 使用说明

### 前台访问
- 访问地址：`http://localhost:5173/user`
- 功能：浏览博客内容，提交对局投稿

### 后台访问
- 访问地址：`http://localhost:5173/meiriyidengdehoutai`
- 功能：查看投稿数据

## 对局投稿说明

### 对局类型
- 超唐对局
- 精彩对局
- 上大师了!

### 投稿要求
- 对局类型和录像ID为必填项
- 对局描述为选填项，但建议填写以便UP主快速定位内容
- 粉丝福利投稿需选择"上大师了!"类型，并在描述中注明抖音或B站ID

## 自定义配置

### 图片资源
- 头像图片：放置在 `src/assets/images/` 目录下
- OSS图片链接：在相应组件中直接替换图片URL

### 内容更新
- 特别鸣谢：在 `Acknowledgments.vue` 组件中硬编码添加
- 必看视频：在 `MustWatch.vue` 组件中硬编码添加
- 交流群和支持主包图片：在相应组件中替换OSS链接

## 设计特色

- **Typora风格**: 简洁、优雅的界面设计
- **响应式布局**: 适配不同屏幕尺寸
- **轻量级**: 无复杂依赖，快速加载
- **无后端**: 纯前端实现，数据通过Vue响应式管理

## 备案信息

网站底部显示备案号：粤ICP备2025451803号-1
链接至：http://beian.miit.gov.cn/

## 开发说明

本项目采用纯前端实现，所有数据通过Vue的响应式系统管理，无需后端服务器和数据库。投稿数据存储在内存中，页面刷新后会重置。如需持久化存储，可考虑使用localStorage或集成轻量级后端服务。 