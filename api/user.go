package api

import (
	"github.com/flourish-ship/work-attendance/consts"
	"github.com/flourish-ship/work-attendance/model"
	"github.com/kataras/iris"
)

type UserApi struct {
}

func (userApi *UserApi) Register(api *iris.Framework) {
	user := api.Party(consts.URL_PRFIX)
	{
		user.Get("/user", GetUser)
	}
}

func GetUser(u *iris.Context) {
	user := model.User{UserName: "mojo-zd", Gender: "male", NickName: "mojo"}
	u.JSON(iris.StatusOK, user)
}

//func (u UserApi) Post() {
//	user := model.User{}
//	if err := u.Context.ReadForm(&user); err != nil {
//		fmt.Printf("%+v\n", err)
//		u.JSON(errors.BADREQUEST_CODE, err)
//		return
//	}
//	u.JSON(iris.StatusOK, user)
//}
