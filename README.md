# vertical
基础组件操作 （不断补充更新中）

```
go get -u github.com/joycastle/vertical@main 
```

### 使用demo 
```
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/joycastle/matching-story/app"
    
    vertical_config "github.com/joycastle/vertical/config"
    vertical_conn "github.com/joycastle/vertical/connector"
    vertical_gin "github.com/joycastle/vertical/gin"
    vertical_logger "github.com/joycastle/vertical/logger"
)

func main() {

    conf_path := "conf/test"

    //初始化配置
    vertical_config.InitYmalConfig(conf_path)

    //初始化日志
    vertical_logger.InitLogger(vertical_config.C_Log)

    //初始化mysql连接
    vertical_conn.InitMysqlConn(vertical_config.C_Mysql)

    //初始化redis连接
    vertical_conn.InitRedisConn(vertical_config.C_Redis)

    //初始化gin框架
    vertical_gin.InitGinServer(gin.ReleaseMode)

    app.InitMiddware()
    app.InitControler()

    vertical_gin.StartGin(vertical_config.C_Gin, 4001)
}
```


### 配置（yaml）
```
  常用配置：（直接解析）
    ├── gin.yaml
    ├── log.yaml
    ├── mysql.yaml
    ├── redis.yaml
    
  自定义配置
    var gcf GrpcConf
    config.RegisterParser("grpc.yaml"， &grpc) 
    ------
    config.RegisterParser("配置文件名"， ”解析结构体“) 
  
  初始化配置
  config.InitYmalConfig(conf_path)   conf_path 配置文件文件夹
```

### logger使用
```
log.yaml
  error: ./error.log
  run: ./run.log-*-*-*
  mysql: ./mysql.log-*-*
  {yourlog}: {logerfile}
  (需要什么日志追加)

规则：
./error.log 日志不拆分
./run.log-*-*-* 日志按照年月日拆分
./mysql.log-*-* 日志按照年月拆分

星号代表日期拆分规则分隔符
*  * *  *  *
年 月 日 时 分


1.初始化：logger.InitLogger(vertical_config.C_Log)
2.设置颜色：打开（默认）：log.EnableColor()  关闭 log.DisableColor() 
3.调用：logger.GetLogger("error").Debug()|Info()|Warn()|Fatal() | Debugf()|Infof()|Warnf()|Fatalf()
```

### Mysql
```
mysql.yaml:
db_game:
  Master:
    Dsn: root:123456@tcp(127.0.0.1:3306)/db_game?charset=utf8mb4&parseTime=True&timeOut=10s
    MaxIdle: 16
    MaxOpen: 64
    MaxLifeTime: 86400s
  Slave:
    Dsn: root:123456@tcp(127.0.0.1:3306)/db_game?charset=utf8mb4&parseTime=True&timeOut=10s
    MaxIdle: 16
    MaxOpen: 64
    MaxLifeTime: 86400s
    
db_game2:
  Master:
    Dsn: root:123456@tcp(127.0.0.1:3306)/db_game?charset=utf8mb4&parseTime=True&timeOut=10s
    MaxIdle: 16
    MaxOpen: 64
    MaxLifeTime: 86400s
  Slave:
    Dsn: root:123456@tcp(127.0.0.1:3306)/db_game?charset=utf8mb4&parseTime=True&timeOut=10s
    MaxIdle: 16
    MaxOpen: 64
    MaxLifeTime: 86400s

1.初始化  connector.InitMysqlConn(config.C_Mysql)
2.调用 
    type User Struct
    var u User
    god := orm.NewGormSwitch("db_game")
  
    #1.使用已有方法、默认会主从选择（gorm方法重写, 参考：orm/gorm.go）
    god.First(&u) 
    
    #2.使用原生gorm,根据读写选择主从.
    god.Master.First(&u)
    god.Slave.First(&u)
```

### Redis
```
redis.yaml
default:
  Addr: 127.0.0.1:6379,127.0.0.1:6379,127.0.0.1:6379
  Password: 123456
  MaxActive: 64
  MaxIdle: 16
  IdleTimeout: 240s
  ConnectTimeout: 10s
  ReadTimeout: 500ms
  WriteTimeout: 500ms
  TestInterval: 60s
redis2:
  Addr: 127.0.0.1:6379,127.0.0.1:6379,127.0.0.1:6379
  Password: 123456
  MaxActive: 64
  MaxIdle: 16
  IdleTimeout: 240s
  ConnectTimeout: 10s
  ReadTimeout: 500ms
  WriteTimeout: 500ms
  TestInterval: 60s
 
 1.初始化 : connector.InitRedisConn(vertical_config.C_Redis)
 2.调用： redis.Rds_GetSting("redis2", "redis_key") 参考 /redis文件下操作
```

### 目录结构
```
.
├── README.md
├── config                #配置文件解析
│   ├── app.go
│   ├── gin.yaml
│   ├── log.yaml
│   ├── mysql.yaml
│   ├── redis.yaml
│   ├── test.yaml
│   ├── yaml.go
│   └── yaml_test.go
├── connector           #连接器connector
│   ├── mysql.go
│   ├── mysql_test.go
│   ├── redis.go
│   └── redis_test.go
├── gin                 #gin
│   ├── gin.go
│   └── gin_test.go
├── go.mod
├── go.sum
├── log                 #日志基础模块
│   ├── log.go
│   ├── log_file.go
│   └── log_test.go
├── logger              #日志上层应用
│   └── logger.go
├── orm                 #数据库
│   ├── gorm.go
│   └── gorm_test.go
├── redis               #缓存
│   ├── redis_hash.go
│   ├── redis_hash_test.go
│   ├── redis_key.go
│   └── redis_key_test.go
└── util                工具集
    ├── file.go
    ├── file_test.go
    └── ud_time.go
```
