package models

import (
	// "errors"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/duguying/ojsite/utils"
	// "github.com/gogather/com"
	"fmt"
	"github.com/duguying/judger/client"
	"github.com/gogather/com/log"
	"html"
	"time"
)

type Submissions struct {
	Id         int
	Pid        int
	Uid        int
	Type       string
	Language   string
	Code       string
	Judger     string
	Status     string
	BuildLog   string
	RunLog     string
	SubmitTime time.Time
	JudgeTime  time.Time
}

// add submission
func (this *Submissions) Add(pid int, uid int, ptype string, language string, code string, judger string) (int64, error) {
	code = html.EscapeString(code)

	o := orm.NewOrm()
	var subm Submissions
	subm.Pid = pid
	subm.Uid = uid
	subm.Type = ptype
	subm.Language = language
	subm.Code = code
	subm.Judger = judger
	subm.SubmitTime = time.Now()
	subm.JudgeTime = time.Now()

	id, err := o.Insert(&subm)

	if err != nil {
		return id, err
	}

	msg := utils.MsgPack(map[string]interface{}{
		"action":   "task_add",
		"sid":      "randomstring",
		"id":       id,
		"time":     time.Now(),
		"language": language,
		"code":     code,
	})

	fmt.Println(msg)

	client.J.Request(msg)

	return id, err
}

// update submission status
func UpdateSubmissionStatus(id int, status string) error {
	o := orm.NewOrm()
	var subm Submissions
	subm.Id = id
	err := o.Read(&subm, "Id")

	if err != nil {
		log.Warnf("记录[%d]不存在\n", id)
		return err
	} else {
		subm.Status = status
		if _, err := o.Update(&subm); err != nil {
			log.Warnln("状态更新失败")
			return err
		} else {
			return nil
		}
	}
}
