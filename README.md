## 《幻兽帕鲁》 服务器助手

## 免责声明
本程序仅供参考和个人使用，不得用于商业目的。使用者应自行承担使用本程序的风险，对于因使用本程序所导致的任何直接或间接损失，我概不负责。本程序不保证其提供的信息准确、完整、适用或有效，也不对使用本程序可能引发的风险承担责任。

## 作用
定时检测服务器内存占用情况，如超出指定阈值，向游戏内玩家广播退出倒计时，执行`/Shutdown`方法将玩家踢下线，然后调用指定脚本重启服务器。

## 环境要求
Ubuntu x64

## 使用方法
### 1. 下载文件
### 2. 服务器配置开启Rcon
修改你的服务器配置文件，确保以下配置
```
RCONEnabled=True,RCONPort=25570
```
### 2. 配置
修改`helper_config.json`
```
{
	"serverIPAndPort": "127.0.0.1:25570",           #服务器Rcon的IP和密码
    "serverPassword": "srv_pwd",                    #服务器Rcon密码
    "rebootScriptPath": "/path/to/you/restart.sh",  #重启服务器的脚本路径
	"oomThreshold": 70,                             #内存阈值，超出该值将重启
	"checkIntervalSeconds": 5,                      #内存占用检查间隔，如每5s检查一次
	"rebootSeconds": 60                             #重启倒计时，单位秒
}
```
### 3. 创建后台服务
1. 创建`/etc/systemd/system/pal-server-helper.server`，设置程序路径及配置文件路径
```
[Unit]
Description=Pal server helper.
After=network.target

[Service]
Type=simple
ExecStart=/path/to/pal-server-helper -c /path/to/helper_config.json

[Install]
WantedBy=multi-user.target
```

2. 执行以下命令
```
sudo systemctl daemon-reload            #重新加载 systemd 配置
sudo systemctl enable pal-server-helper #开机自启
sudo systemctl start pal-server-helper  #启动服务
```
