
#### 使用
---
* `Dokcer`（默认支持部署方式）
    1. `make initdb` // 初始化数据库等资源（`docker docker-compose`需要正确安装）
    2. `make run` // 运行服务
* DIY打包
    1. `make build` // 打包资源
* 文档
    1. `make api` // 生成api资源
    2. `/api/roles/service/v1` 下生成 `roles.swagger.json` 可直接导入`ApiFox`等工具

#### 功能介绍
* 目的：包装后台管理项目用户管理/权限管理/角色管理/路由与角色与权限关联管理等功能，开发管理后台只需关心自己系统本身业务即可。
* 功能
    * 后台用户管理
    * 角色管理
    * 权限管理
    * 路由管理
    * 鉴权功能（用户操作你的业务数据时通过 调用本服务 进行鉴权）
