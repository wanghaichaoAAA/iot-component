# iot-component
数采仪、网关公用组件
#### 主要涵盖：
1. 公用对象
2. 公用工具类

##现在开始
- 安装
`https://github.com/wanghaichaoAAA/iot-component.git`

##文件分层
```
├── README.md
├── object                      公共对象
│   ├── common.go               公共常量
│   ├── hj212command.go         Hj212-2017命令
│   ├── hj212messageobject.go   Hj212-2017消息对象结构体及其方法
│   └── remotecontrolobj.go     远程配置交互结构体
└── utils                       工具类
    ├── calcutils.go            数学计算工具类          
    ├── crcutils.go             crc校验工具类
    ├── crcutils.go             crc校验工具类
    ├── idworkerutils.go        id生成器工具类
    ├── pageutils.go            分页参数工具类
    ├── resp.go                 gin通用返回对象
    ├── stringutils.go          字符串操作工具类
    └── timeutils.go            时间操作工具类│      
```