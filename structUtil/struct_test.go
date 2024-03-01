package structUtil

import (
	"fmt"
	"testing"
	"time"
)

type AdminUser struct {
	ID         int       `json:"id" form:"id" xorm:"pk autoincr"`
	Username   string    `json:"username" form:"username" xorm:"unique comment(用户名)"`
	Password   string    `json:"password" form:"password" xorm:"password comment(密码)"`
	GroupID    int       `json:"group_id" form:"group_id" xorm:"comment(用户组)"`
	Telegram   string    `json:"telegram" form:"telegram" xorm:"comment(电报)"`
	TelegramID int64     `json:"telegram_id" form:"telegram_id" xorm:"comment(电报ID)"`
	Version    int       `json:"version" form:"version" xorm:"default(1) comment(版本)"`
	State      bool      `json:"state" form:"state" xorm:"comment(状态)"`
	LoginTime  time.Time `json:"login_time" form:"login_time" xorm:"comment(上次登录时间)"`
	CreateTime time.Time `json:"create_time" form:"create_time" xorm:"created comment(创建时间)"`
	Test
}
type Test struct {
	Version    int       `json:"version" form:"version" xorm:"default(1) comment(版本)"`
	State      bool      `json:"state" form:"state" xorm:"comment(状态)"`
	LoginTime  time.Time `json:"login_time" form:"login_time" xorm:"comment(上次登录时间)"`
	CreateTime time.Time `json:"create_time" form:"create_time" xorm:"created comment(创建时间)"`
}

func TestNewStruct(t *testing.T) {
	s := NewStruct(AdminUser{})
	for i, field := range s.Fields {
		fmt.Println(i, field.Names().Snake("id"), field.Type().String(), field.Tag().Match("comment\\((.*?)\\)"))
	}
}
