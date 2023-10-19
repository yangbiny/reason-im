# reason-im

全部使用 go 语言 实现的 轻量级 IM 服务。目前仅支持 登录、注册、邀请朋友 和 创建 websocket 链接

自己封装了 MySQL 工具，可以更方便快捷的 使用MySQL， 对 logger 工具 进行封装，以便打印更详细的日志信息

## 项目结构

- cmd
    - 项目的启动入口
- config
    - 配置信息
- internal
    - 内部 实现。不对外暴露
    - 主要是 API 和 service
    - utils 为 封装的工具包
- pkg
    - 对外暴露的信息