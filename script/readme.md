### script 脚本



```shell
go run main.go --config=configFilePath --cmd=cmdCode

# cmd `po` | `repository` | `dependency` | `all`
```


```yaml
 # mysql config
  Mysql:
    Usage: "default"
    RunMode: "debug"
    DSN: "root:root@tcp(127.0.0.1:3307)"
    Database: "coffee"
    Prefix: "system_"
    MaxIdleConn: 10
    MaxOpenConn: 10
    MaxLifeTime: 10
  # table prefix
  Prefix: "yunka_system_"
  # table name
  Table: "yunka_system_api"
  # source file path
  SourceFile: "./tmpl/dependency/model.go"
  # target path
  TargetPath: "./new/dependency"
  # import po path
  NewPoPath: "weicai.zhao.io/script/new/po"
  # tmpl file, import po path
  OldPoPath: "weicai.zhao.io/script/tmpl/po"
```