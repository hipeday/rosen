# Banner

> Application start banner.

## 语法

### `${}`

这个方式通过 `${}` 来引用变量，变量名是由[配置文件](https://github.com/hipeday/rosen/blob/main/conf/config.yaml)的键。

e.g.

```text
这个应用程序名字: ${application.name} 当前版本号是: ${application.version}

# 输出
这个应用程序名字: Rosen 当前版本号是: 1.0.0
```

### `%%`

> 支持变量列表

| 变量名 | 值 | Since |   描述    | 当前状态 |
|:---:|:-:|:-----:|:-------:|:----:|
| PID | - | 1.0.0 | 当前程序PID |  ✅   |

这个方式通过 `%%` 来引用环境变量，变量名是由环境变量的键。

e.g.

```text
当前应用程序的PID是: %PID%

# 输出
当前应用程序的PID是: 1234
```