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