package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"gt3-server-golang-gin-sdk/controllers/sdk"
)

// 验证初始化接口，GET请求
func FirstRegister(c *gin.Context) {
	/*
	   必传参数
	       digestmod 此版本sdk可支持md5、sha256、hmac-sha256，md5之外的算法需特殊配置的账号，联系极验客服
	   自定义参数,可选择添加
		   user_id 客户端用户的唯一标识，确定用户的唯一性；作用于提供进阶数据分析服务，可在register和validate接口传入，不传入也不影响验证服务的使用；若担心用户信息风险，可作预处理(如哈希处理)再提供到极验
		   client_type 客户端类型，web：电脑上的浏览器；h5：手机上的浏览器，包括移动应用内完全内置的web_view；native：通过原生sdk植入app应用的方式；unknown：未知
		   ip_address 客户端请求sdk服务器的ip地址
	*/
	gtLib := sdk.NewGeetestLib(GEETEST_ID, GEETEST_KEY)
	digestmod := "md5"
	userId := "test"
	params := map[string]string{
		"digestmod":   digestmod,
		"user_id":     userId,
		"client_type": "web",
		"ip_address":  "127.0.0.1",
	}
	result := gtLib.Register(digestmod, params)
	// 将结果状态写到session中，此处register接口存入session，后续validate接口会取出使用
	// 注意，此demo应用的session是单机模式，格外注意分布式环境下session的应用
	session := sessions.Default(c)
	session.Set(sdk.GEETEST_SERVER_STATUS_SESSION_KEY, result.Status)
	session.Set("userId", userId)
	session.Save()
	// 注意，不要更改返回的结构和值类型
	c.Header("Content-Type", "application/json;charset=UTF-8")
	c.String(http.StatusOK, result.Data)
}

// 二次验证接口，POST请求
func SecondValidate(c *gin.Context) {
	gtLib := sdk.NewGeetestLib(GEETEST_ID, GEETEST_KEY)
	challenge := c.PostForm(sdk.GEETEST_CHALLENGE)
	validate := c.PostForm(sdk.GEETEST_VALIDATE)
	seccode := c.PostForm(sdk.GEETEST_SECCODE)
	// session必须取出值，若取不出值，直接当做异常退出
	session := sessions.Default(c)
	status := session.Get(sdk.GEETEST_SERVER_STATUS_SESSION_KEY)
	userId := session.Get("userId")
	if status == nil {
		c.JSON(http.StatusOK, gin.H{"result": "fail", "version": sdk.VERSION, "msg": "session取key发生异常"})
		return
	}
	var result *sdk.GeetestLibResult
	if status.(int) == 1 {
		/*
		   自定义参数,可选择添加
		       user_id 客户端用户的唯一标识，确定用户的唯一性；作用于提供进阶数据分析服务，可在register和validate接口传入，不传入也不影响验证服务的使用；若担心用户信息风险，可作预处理(如哈希处理)再提供到极验
			   client_type 客户端类型，web：电脑上的浏览器；h5：手机上的浏览器，包括移动应用内完全内置的web_view；native：通过原生sdk植入app应用的方式；unknown：未知
			   ip_address 客户端请求sdk服务器的ip地址
		*/
		params := map[string]string{
			"user_id":     userId.(string),
			"client_type": "web",
			"ip_address":  "127.0.0.1",
		}
		result = gtLib.SuccessValidate(challenge, validate, seccode, params)
	} else {
		result = gtLib.FailValidate(challenge, validate, seccode)
	}
	// 注意，不要更改返回的结构和值类型
	if result.Status == 1 {
		c.JSON(http.StatusOK, gin.H{"result": "success", "version": sdk.VERSION})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "fail", "version": sdk.VERSION, "msg": result.Msg})
	}
}
