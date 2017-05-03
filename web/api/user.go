package api

import (
	"github.com/flourish-ship/work-attendance/consts"
	"github.com/flourish-ship/work-attendance/web/model"
	"github.com/flourish-ship/work-attendance/web/service"

	//"github.com/flourish-ship/work-attendance/auth"

	"github.com/kataras/iris"
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

func GetUser(u *iris.Context) {
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

func Create(u *iris.Context) {
	user := &model.User{}
	err := u.ReadJSON(user)
	if ErrorHandler(u, err) {
		return
	}

	if ErrorHandler(u, service.CreateUser(user)) {
		return
	}
	u.JSON(iris.StatusOK, user)
}
