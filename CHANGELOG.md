# 更新日志

> 此文档为`QQBot_go`项目的更新日志

## [0.7.1] - 2022

### 优化

- 优化了项目结构



## [0.7.0] - 2022-10-13

### 新增

- 添加`github.com/sirupsen/logrus`，项目全面使用log框架进行输出管理
- 日志输出文件保存到`logs`文件夹 (文件格式: 2022-10-12.log)
- 错误日志增加调用文件详细信息后单独保存 (文件格式: error-2022-10-12.log)
- 接收消息错误后重连
- 新增`CHANGLOG.md`文件，用于记录更新日志

### 变更

- 更改`QQBot\api`对http请求结果的处理

- 变更`.gitignore`文件

- 变更`/info`命令回复消息

  ```go
  var info = "本机器人由YGXB_net开发" +
  	"\nQQ:3040809965" +
  	"\n\n当前版本: " + data.Version +
  	"\n更新日志: https://gitee.com/YGXB-net/QQBot_go/blob/develop/CHANGELOG.md"
  ```

### 优化

- 优化websocket连接逻辑

- 优化打印接收到消息的格式