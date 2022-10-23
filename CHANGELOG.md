# 更新日志

> 此文档为`QQBot_go`项目的更新日志

## [0.7.6] - 2022-10-23

### 变更

- 更改对于消息切片长度不足的处理

```go
//因为切片会出现长度不足，所以会抛出异常
defer func() { recover() }()
if message[0:1] == "/" || message[0:3] == "／" {
    httpapi.Send_group_msg(group_id, "命令输入错误或没有此命令\n请输入 /help 查看帮助")
} else if strings.Index(message, "[CQ:at,qq=2700154874]") != -1 {
    httpapi.Send_group_msg(group_id, fmt.Sprintf("[CQ:at,qq=%s] 叫你爸爸干嘛？", user_id))
}
```

### 新增

- 新增Group中对表情包命令的处理，通过计算文件的MD5值来判断

### 优化

- 优化 `service/handle_order/Group.go` 的代码，添加一堆注释



## [0.7.5] - 2022-10-19

### 新增

- 增加@机器人时回复：叫你爸爸干嘛？

### 优化

- 优化整体项目结构，使项目更好开发（应该）



## [0.7.4] - 2022-10-19

### 修复

- 修复严重bug，处理命令时，对 `message` 切片会出现长度不足导致程序退出

```go
//文件: service/handle_order/Group.go:46
default:
		if message[0:1] == "/" || message[0:3] == "／" {
			httpapi.Send_group_msg(group_id, "命令输入错误或没有此命令\n请输入 /help 查看帮助")
		}
	}
```



## [0.7.3] - 2022-10-18

### 变更

- 更改 `/help` 命令回复消息

```go
var help_info = "----------帮助信息----------" +
	"\n\n/help 获取帮助" +
	"\n/info 获取机器人信息" +
	"\n\n/dk 进行打卡"
```

### 移除

- 移除了对命令触发的限制，可以使用 `／` 作为命令触发符号



## [0.7.2] - 2022-10-17

### 新增

- 增加打卡方式

### 变更

- 更改命令处理逻辑



## [0.7.1] - 2022-10-14

### 变更

- 变更`/info`命令回复消息

```go
var info = "本机器人由YGXB_net开发" +
	"\nQQ:3040809965" +
	"\n\n当前版本: " + base.Version +
	"\n更新日志: https://gitee.com/YGXB-net/QQBot_go/blob/master/CHANGELOG.md"
```

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