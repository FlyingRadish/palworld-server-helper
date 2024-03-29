## 《幻兽帕鲁》 服务器助手

## 免责声明
本程序仅供参考和个人使用，不得用于商业目的。使用者应自行承担使用本程序的风险，对于因使用本程序所导致的任何直接或间接损失，我概不负责。本程序不保证其提供的信息准确、完整、适用或有效，也不对使用本程序可能引发的风险承担责任。

## 作用
- 定时检测服务器内存占用情况，如超出指定阈值，向游戏内玩家广播退出倒计时，踢出所有玩家后，调用指定脚本重启服务器。
- 服务器面板，提供如下功能
  - 查看运行状态、内存状态
  - 查看在线玩家列表
  - 游戏内广播
  - 发送RCON指令
  - 重启(直接重启/广播倒计时重启)

![面板示例](/panel-screeshot.png)

## 环境要求
Ubuntu x64

## 使用方法
### 1. 服务器配置开启Rcon
修改你的服务器配置文件，确保以下配置
```
RCONEnabled=True,RCONPort=25570
```
### 2. 下载release后解压
### 3. 配置
修改`helper_config.json`
```
{
    "ip": "127.0.0.1",                              #服务器Rcon的IP
    "port": 25570,                                  #服务器Rcon端口
    "password": "srv_pwd",                          #服务器Rcon密码
    "retryDelay": 10,                               #Rcon重试前等待时间，实测过短会导致再也连不上
    "retryCount": 3,                                #Rcon重试次数
    "rebootScriptPath": "/path/to/your/restart.sh",  #重启服务器的脚本路径
    "rebootSeconds": 60                             #重启倒计时，单位秒
    "oomThreshold": 70,                             #内存阈值，超出该值将重启
    "oomCheckIntervalSeconds": 5,                   #内存占用检查间隔，如每5s检查一次
    "playerCheckInterval": 5,                       #在线玩家列表检测间隔，如每5s检查一次
    "apiHost": "host_ip_or_domain",                 #API服务器地址，可以写服务器IP，或域名，不能设为127.0.0.1/localhost
    "apiPort": 8311                                 #API服务器端口
    "panelPath": "/path/to/panel"                   #前端面板路径(panel文件夹的路径)
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
