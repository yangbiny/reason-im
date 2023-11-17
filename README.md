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

## support

会有一些 基础包的提供，在 [Github](https://github.com/yangbiny/reason-commons)

## Feature

- 群发 功能，可以通过 群发的方式 发送 数据，实现类型 微信的 群聊
- 发布订阅 功能，可以通过 发布订阅的方式 发送 数据，实现 类似于 微信的 公众号
  - 可以监听 订阅的消息
  - 订阅 管理者 发布的时候，会通知 所有的订阅者
  - 如何 维护 大数据量的 订阅 消息？
- 朋友圈
    - 朋友圈发布
    - 朋友圈发布通知
- 远程文件上传
