package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id         bson.ObjectId `json:"id,omitempty"`
	UserName   string        `json:"userName,omitempty" form:"userName"` //姓名
	Gender     string        `json:"Gender,omitempty"`                   //性别
	NickName   string        `json:"nickName,omitempty"`                 //昵称
	PinYin     string        `json:"pinyin,omitempty"`                   //简拼
	FullPinYin string        `json:"fullPinYin,omitempty"`               //全拼
	PassWord   string        `json:"passWord,omitempty" form:"passWord"`
	Salt_Hash  string        `json:"-"`
	Salt       string        `json:"-"`
}
