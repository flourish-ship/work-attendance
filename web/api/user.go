package api

import (
	"github.com/flourish-ship/work-attendance/consts"
	"github.com/flourish-ship/work-attendance/web/model"
	"github.com/flourish-ship/work-attendance/web/service"

	//"github.com/flourish-ship/work-attendance/auth"
	"gopkg.in/kataras/iris.v6"
)

type UserApi struct {
}

func (userApi *UserApi) Register(api *iris.Framework) {

	user := api.Party(consts.URL_PRFIX + "/user")
	{
		user.Get("", GetUser)
		user.Post("/singin", SingIn)
		user.Post("", Create)
	}
}

func GetUser(ctx *iris.Context) {
	//u.JSON(iris.StatusOK, user)
}

// 登陆
func SingIn(u *iris.Context) {

	//user := model.User{}
	//
	//if err := json.Unmarshal(u.Request.Body(), &user); err != nil {
	//	u.JSON(errors.BADREQUEST_CODE, response.DefaultResponse(consts.BAD_PARAMS_ERROR))
	//	return
	//}

	//auth.Authentication(user.UserName, user.PassWord)
	//authentication := auth.Authentication(user.UserName, user.PassWord)
	//u.Session().Set("authUser", authentication)
}

func Create(ctx *iris.Context) {
	user := &model.User{}
	err := ctx.ReadJSON(user)
	if ErrorHandler(ctx, err) {
		return
	}

	if ErrorHandler(ctx, service.CreateUser(user)) {
		return
	}
	ctx.JSON(iris.StatusOK, user)
}
