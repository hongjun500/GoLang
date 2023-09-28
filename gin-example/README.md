# gin 框架使用示例

------
详见 [example](./restful-server/example) 目录下的示例代码

## 备忘录

### init

- 初始化一个 gin 引擎 `engine := gin.Default()` 默认使用了 Logger 和 Recovery
  中间件，如果不需要可以使用 `engine := gin.New()`
- `engine.Run()` 默认监听 `8080` 端口，可以使用 `engine.Run(":8081")` 指定端口
- `engine.RunTLS()` 启动一个 HTTPS 服务，需要传入证书和密钥文件
- 后续待补充

### restful api

> `gin` 的上下文 `context`

- 显示绑定请求的数据 `context.ShouldBindWith(&user, binding.Form)`, `binding.Form`
  表示绑定的数据来源，可以是 `form`、`json`、`xml`、`yaml`、`protobuf`、`msgpack` 等
- 绑定 `Content-Type=multipart/form-data` 的表单数据 `context.ShouldBindWith(&user, binding.FormMultipart)`
  ，或者使用 `context.PostForm("key")` 获取表单数据
- 当 `Content-Type: application/x-www-form-urlencoded` 时另外一种获取请求参数的方式

> 请求的 `url` 中带有 `id`和 `page`,`message` 和 `name` 在 `body` 中

```go
id := context.Query("id")
page := context.DefaultQuery("page", "0") // 设置默认值
name := context.PostForm("name")
message := context.PostForm("message")
```

- 只绑定 `url` 中的参数而忽略 `body` 中的参数 `context.ShouldBindQuery(&user)`
- 设置 Cookie `context.SetCookie("user", "gin", 3600, "/", "localhost", false, true)`
- 获取 Cookie `cookie, err := context.Cookie("user")`
- 绑定 `GET` 或者 `POST` 的请求参数到对应结构体 `context.ShouldBind(&user)`
- 