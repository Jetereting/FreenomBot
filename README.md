# FreenomBot Freenom机器人
freenom 中的域名在14天快过期的时候自动续期
<br>
The domain name in freenom is automatically renewed when it is about to expire in 14 days
---
# How to use it 如何使用

## Edit config.toml 编辑 config.toml 配置文件
会优先读取环境变量的 FreenomBot.toml 变量，如果没有则读取同目录下的 config.toml 文件
<br>
The FreedomBot.toml variable of the environment variable will be read preferentially. If not, the config.toml file in the same directory will be read
``` toml
# 浏览服务的自定义账号密码
[System]
Account = "admin"
Password = "123"
CronTiming = "23:30"

# 微信通知
[WeChatNotify]
Enable = true
# Daily true为每天发，false只在更新时发通知
Daily = true
CorpID = "wwxxxxxxxxxxxxxx84"
CorpSecret = "APIAKxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxU"
AgentID = "1xxxxx2"

# freenom 的账号密码
[[Accounts]]
Username = "xxxxxxxxxx@qq.com"
Password = "xxxxxxxxxxxxxx"

#[[Accounts]]
#Username = "xxxxxxxx@gmail.com"
#Password = "xxxxxxx"
```

## Launch FreenomBot 启动FreenomBot

``` sh
go run .
```
It will start http service on server, So you may check the status page of FrenomBot on http://localhost:8080
<br>
它将在服务器上启动http服务，因此您可以查看FrenomBot的状态页面： http://localhost:8080

