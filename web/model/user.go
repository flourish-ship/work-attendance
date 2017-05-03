package model

import "github.com/flourish-ship/work-attendance/web/tools"

type User struct {
	Id         int64  `json:"id,omitempty" xorm:"pk autoincr"`
	UserName   string `json:"userName,omitempty"`   //姓名
	Gender     string `json:"Gender,omitempty"`     //性别
	NickName   string `json:"nickName,omitempty"`   //昵称
	PinYin     string `json:"pinyin,omitempty"`     //简拼
	FullPinYin string `json:"fullPinYin,omitempty"` //全拼
	PassWord   string `json:"passWord,omitempty"`
	CreatedAt  int64  `json:"createdAt" xorm:"created"`
	UpdatedAt  int64  `json:"updatedAt" xorm:"updated"`
	Salt_Hash  string `json:"-"`
	Salt       string `json:"-"`
}

func (user *User) BeforeInsert() {
	user.FullPinYin = tools.FullPinYin(user.UserName)
}
