package model

type User struct {
	Id         int64  `json:"id,omitempty" xorm:"pk autoincr"`
	UserName   string `json:"userName,omitempty"`   //姓名
	Gender     string `json:"Gender,omitempty"`     //性别
	NickName   string `json:"nickName,omitempty"`   //昵称
	PinYin     string `json:"pinyin,omitempty"`     //简拼
	FullPinYin string `json:"fullPinYin,omitempty"` //全拼
	PassWord   string `json:"passWord,omitempty"`
	Salt_Hash  string `json:"-"`
	Salt       string `json:"-"`
}
