# dsp-template
dsp 模版

# Go 项目结构规范

https://www.jianshu.com/p/4726b9ac5fb1

## /cmd 

该项目的主程序.

每个程序目录的名字应该和可执行文件的名字保持一致 (比如 /cmd/myapp).

## /internal

程序和库的私有代码. 这里的代码都是你不希望被别的应用和库所引用的.

把你真正的应用代码放在 /internal/app 目录，把你的应用间共享的代码放在 /internal/pkg 目录

## /pkg

可以被其他外部应用引用的代码

## /vendor

应用的依赖

## /api

OpenAPI/Swagger 规范, JSON schema 文件, 协议定义文件.

## /web

Web 应用标准组件: 静态 Web 资源, 服务端模板, 单页应用.

## /configs

配置文件模板或者默认配置.

在这里放置你的 confd 或者 consul-template 模板文件.

## /init

系统初始化 (systemd, upstart, sysv) 及进程管理/监控 (runit, supervisord) 配置.

## /scripts

执行各种构建, 安装, 分析等其他操作的脚本.

这些脚本要保持根级别的 Makefile 小而简单

## /build

打包及持续集成.

将 cloud (AMI), container (Docker), OS (deb, rpm, pkg) 包配置放在 /build/package 目录下.

将 CI (travis, circle, drone) 配置和脚本放在 /build/ci 目录.

## /deployments

IaaS, Paas, 系统, 容器编排的部署配置和模板 

## /test

额外的外部测试软件和测试数据. 

## /docs

用户及设计文档

## /tools

项目的支持工具. 注意, 这些工具可以引入 /pkg 和 /internal 目录的代码.

## /examples

应用或者库的示例文件.

## /third_party

外部辅助工具, forked 代码, 以及其他第三方工具

## /assets

其他和你的代码仓库一起的资源文件 