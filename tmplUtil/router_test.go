package tmplUtil

import (
	"go-utils/structUtil"
	"testing"
	"time"
)

type AdminUserLog struct {
	ID         int       `json:"id" xorm:"autoincr pk"`
	Username   string    `json:"username" xorm:"unique notnull"`
	Password   string    `json:"-" xorm:"notnull"`
	Ver        int       `json:"-" xorm:"default(1)"`
	RoleID     int       `json:"role_id"`
	Telegram   string    `json:"telegram"`
	TelegramID string    `json:"telegram_id"`
	State      bool      `json:"-" xorm:"default(1)"`
	UpdateTime time.Time `json:"-" xorm:"updated"`
	CreateTime time.Time `json:"-" xorm:"created"`
}

func TestNewRouter(t *testing.T) {
	r := NewRouter(structUtil.NewStruct(AdminUserLog{}, "用户信息"))

	err := r.Write("./router_test2.go1", RouterMethod{
		Method:  "get",
		Url:     ":id<int>",
		Handler: "handler.*.get",
	})
	if err != nil {
		panic(err)
	}
	//blocks := r.RouterBlock()
	//for i, block := range blocks {
	//	fmt.Println(i, "\n", block.Render())
	//}
}
func TestNewTemp(t *testing.T) {
	r := NewTemp(structUtil.NewStruct(AdminUserLog{}, "用户信息"))
	err := r.Write("./router_test2.go1", "./router_test3.go")
	if err != nil {
		panic(err)
	}
}
