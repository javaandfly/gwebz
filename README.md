# 一款基于gin+gorm+cli+viper+logz的web脚手架


## 入门
```bash
go get github.com/javaandfly/go-web-formwork
```
## Config
- 首先你需要创建你的配置文件,当前支持的配置文件类型,
    - 支持的文件类型 .yml , .yaml , .cfg , .json , .conf
- 然后需要自己构造你的配置文件中对应的字段的结构体如下所示
    ```go
    type DatabaseConfig struct {
	DriverName string `mapstructure:"driver_name"`
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	Database   string `mapstructure:"database"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	Charset    string `mapstructure:"charset"`
    }

    type Config struct {
	    DatabaseCfg DatabaseConfig `mapstructure:"db"`
    }
    ```
- 然后使用我们提供的解析方式解析,初始化
    ```go
    // 传入文件路径 支持绝对路径和相对路径，然后传入新建的结构体的指针
    err := ReadConfig("config_test.yaml", cfgPojo)
	if err != nil {
        //.........
	}
    ```
- 这个时候配置文件的所有信息都映射到结构体上了，全部可以使用


## DB
### 方式一 使用全局的连接池
- 初始化
     ```go
     // 参数: 数据库名称，地址，端口号，数据库名，用户名，用户密码，编码格式
    err := InitGlobalDB("mysql", "127.0.0.1", "3306", "test", "root", "123456", "utf8")
	if err != nil {
		panic(err)
	}
     ```
- 获取全局连接池
    ```go
    sqlPool := GetDB()
    ```
### 方式二 返回连接池对象，用户抉择如何使用
- 初始化 
    ```go
     // 参数: 数据库名称，地址，端口号，数据库名，用户名，用户密码，编码格式
    sqlPool,err := NewDB("mysql", "127.0.0.1", "3306", "test", "root", "123456", "utf8")
	if err != nil {
		panic(err)
	}
    ```
- 还支持设置连接池的属性

## Log

- log初始化 
    ```go
    // 参数: filepath 
    serverName := strings.Split(filepath.Base(os.Args[0]), ".")[0]
	serverMark := GetSvrmark(serverName)
    err := InitLog("test_log/", serverName, serverMark, func(str string) {})
	if err != nil {
		panic(err)
	}
    ```
- 使用
    - 用户等级
     ```go
        LogW("this is a log message %v",err)
     ```
    - Debug等级
     ```go
        LogD("this is a log message %v",err)
     ```   
    - 告警日志
     ```go
        LogError(err)
     ``` 
- 回调事件
    - 可以设置回调事件 默认日志等级在 2 3 4 会触发
    ```go
    LogSetCallback(func(s string) {
		LogW("回调函数被调用了")
	})
    ```
- log的参数都是可以设计的

## IOC

- ioc初始化 
    ```go
       RunContainerIOC(httpClint,InitDB,Initlog...)
     ``` 
     - 使用说明
        使用RunContainerIOC的方式启动我们的程序，使用依赖注入的方式；
        第一个参数传入http监听类，后续参数可以传入初始化的方法
        详情可看方法注释 使用的是fx依赖反转框架


## Redis
 - 初始化
    ```go
    err := InitGlobalRedisClient("127.0.0.1:6379", "", "123456", 1)
	if err != nil {
    panic(err)
	}
    ```
- 跟DB一样两种方式 一种全局 一种获取实例