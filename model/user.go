package model

type User struct {
	UserName   string `json:"userName,omitempty"`   //姓名
	Gender     string `json:"Gender,omitempty"`     //性别
	NickName   string `json:"nickName,omitempty"`   //外号
	PinYin     string `json:"pinyin,omitempty"`     //简拼
	FullPinYin string `json:"fullPinYin,omitempty"` //全拼
}
