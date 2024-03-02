package xormUtil

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
}

func TestName(t *testing.T) {
	var query = BuilderQuery{
		"id":   In,
		"name": In,
	}
	db, err := NewMysqlEngine("apphub:123456@tcp(127.0.0.1:3306)/apphub?charset=utf8mb4&timeout=3s&parseTime=true&loc=Local", nil)
	if err != nil {
		panic(err)
	}
	b := Builder{}
	b.Set("id", "11,21,31,41")
	var data []AdminUser
	if err != nil {
		panic(err)
	}
	err = db.Limit(b.Limit(), (b.Page()-1)*b.Limit()).Where(b.BuilderWhere(query)).Find(&data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

}
