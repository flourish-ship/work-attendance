package model

import (
	"github.com/flourish-ship/mongo-cli/services/mongo"
	"github.com/flourish-ship/work-attendance/consts"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id         bson.ObjectId `json:"id"`
	UserName   string        `json:"userName,omitempty"`   //姓名
	Gender     string        `json:"Gender,omitempty"`     //性别
	NickName   string        `json:"nickName,omitempty"`   //外号
	PinYin     string        `json:"pinyin,omitempty"`     //简拼
	FullPinYin string        `json:"fullPinYin,omitempty"` //全拼
}

func (self *User) GetMgoInfo() (string, string, string) {

	s := consts.SESSIONNAME
	d := consts.DBNAME
	c := "user"

	return d, c, s
}

func (self *User) GetId() bson.ObjectId {
	return self.Id
}

func GetUser(id bson.ObjectId) (user *User, err error) {
	user = &User{}
	o := mongo.NewOrm()
	err = o.Find(new(User)).Filter("_id", mongo.Equal, id).One(user)
	return
}

func CreateUser(user *User) error {
	o := mongo.NewOrm()
	return o.Insert(user)
}
