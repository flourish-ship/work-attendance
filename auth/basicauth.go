package auth

import (
	"github.com/iris-contrib/middleware/basicauth"
	"github.com/kataras/iris"
)

func Authentication(userName, passWord string) iris.HandlerFunc {
	return basicauth.Default(map[string]string{"username": userName, "password": passWord})
}
