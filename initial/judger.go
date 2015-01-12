package initial

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/duguying/judger/client"
)

// connect judger and login
func InitJudger() {
	host := beego.AppConfig.String("judgerhost")
	port, err := beego.AppConfig.Int("judgerport")
	pass := beego.AppConfig.String("judgerpass")

	if err != nil {
		port = 1004
	}

	client.New(host, port)
	loginInfo := fmt.Sprintf("{\"action\":\"login\",\"password\":\"%s\"}\003", pass)
	client.J.Request(loginInfo)
}
