### 多环境配置文件

####  `GIN_MODE` 三种值

分别是 debug（开发模式）、release（生产模式）、test（测试模式）

```go
// 部分源码
const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)
```

默认为 debug 值：

```go
// 部分源码
func SetMode(value string) {
   if value == "" {
      value = DebugMode
   }
   //...
}
```

#### 设置 GIN_MODE

##### 第一种：CODE

在 main.go 文件中：

```go
gin.SetMode(gin.DebugMode)
gin.SetMode(gin.ReleaseMode)
gin.SetMode(gin.TestMode)
```

##### 第二种：env

执行 `go run main.go` 命令前，设置环境变量：

**Windows：**

```cmd
set GIN_MODE=release
echo %GIN_MODE%
go run main.go
```

**Linux：**

```bash
export GIN_MODE=release
echo $GIN_MODE
go run main.go
```

**注意：**两种都设置的时候，以 code 中设置的为准。

#### 配置文件

三种配置文件：

application-[GIN_MODE].yml

自动根据 `GIN_MODE` 的值加载相对应的配置文件。

