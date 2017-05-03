package api

import (
	"github.com/flourish-ship/work-attendance/consts"
	"github.com/kataras/iris"
)

type SecretApi struct {
}

func (secretApi *SecretApi) Register(api *iris.Framework) {
	secret := api.Party(consts.URL_PRFIX + "/secret")
	{
		{
			secret.Get("", SecretFilter)
		}
	}
}

func SecretFilter(c *iris.Context) {
	// 验证登陆的用户是否合法
	//userName := c.Session().GetString("username")
	//passWord := c.Session().GetString("password")
}
