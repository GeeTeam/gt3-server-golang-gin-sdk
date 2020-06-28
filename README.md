# gt3-server-golang-gin-sdk

## 示例部署环境
条目|说明
----|----
操作系统|ubuntu 16.04.6 lts
golang版本|1.13.4 linux/amd64
gin版本|1.6.3

> **注意**，go.mod是golang1.11版本新引入的官方包管理工具，go1.13已经默认开启

## 部署流程

### 下载sdk demo
```
git clone https://github.com/GeeTeam/gt3-server-golang-gin-sdk.git
```

### 配置密钥，修改请求参数
> 配置密钥

从[极验管理后台](https://auth.geetest.com/login/)获取公钥（id）和私钥（key）, 并在代码中配置。配置文件的相对路径如下：
```
controllers/geetest_config.go
```

> 修改请求参数（可选）

名称|说明
----|------
user_id|客户端用户的唯一标识，作用于提供进阶数据分析服务，可在register和validate接口传入，不传入也不影响验证服务的使用；若担心用户信息风险，可作预处理(如哈希处理)再提供到极验
client_type|客户端类型，web：电脑上的浏览器；h5：手机上的浏览器，包括移动应用内完全内置的web_view；native：通过原生sdk植入app应用的方式；unknown：未知
ip_address|客户端请求sdk服务器的ip地址

### 关键文件说明
名称|说明|相对路径
----|----|----
main.go|项目入口，静态资源、中间件、路由设置等|
geetest_controller.go|接口请求控制器，主要处理验证初始化和二次验证接口请求|controllers/
geetest_config.go|配置id和key|controllers/
geetest_lib.go|核心sdk，处理各种业务|controllers/sdk/
geetest_lib_result.go|核心sdk返回数据的包装对象|controllers/sdk/
index.html|demo示例首页|static/
go.mod|依赖管理配置文件|

### 运行demo
```
cd gt3-server-golang-gin-sdk
sudo go run main.go
```
在浏览器中访问`http://localhost:8000`即可看到demo界面。

## 发布日志

### tag：20200701
- 统一各语言sdk标准
- 版本：golang-gin:3.1.0


