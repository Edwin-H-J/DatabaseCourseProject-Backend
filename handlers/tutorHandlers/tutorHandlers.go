package tutorhandlers

import (
	Process "backend/database/process"
	Selected "backend/database/selected"
	Student "backend/database/student"
	Topic "backend/database/topic"
	Tutor "backend/database/tutor"
	DBUtil "backend/database/util"
	Auth "backend/middleware/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTopicHandler(c *gin.Context) {
	var topic Topic.Topic
	c.Bind(&topic)
	Info, err := Auth.LoginInfo(c)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "您未登录或无权限",
		})
		return
	}
	topic.Tutor.Id = Info["info_id"].(int64)
	topic, err = Topic.CreateTopic(topic, nil)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "创建失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"msg":   "创建成功",
		"topic": topic,
	})
}

func GetUserTopic(c *gin.Context) {
	Info, err := Auth.LoginInfo(c)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "您未登录或无权限",
		})
		return
	}
	topics, err := Topic.GetTopicListByTutorId(Info["info_id"].(int64), nil)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "创建成功",
		"topics": topics,
	})
}

func GetAllTutorConfirmSelectedHandler(c *gin.Context) {
	selections, err := Selected.GetTutorConfirmSelected(nil)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	for i := range selections {
		selections[i].Student, err = Student.GetStudentById(selections[i].Sid, nil)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 404,
				"msg":  "数据库异常",
			})
			return
		}
		selections[i].Topic, err = Topic.GetTopicById(selections[i].Tid, nil)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 404,
				"msg":  "数据库异常",
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": selections,
	})
}
func GetTutorConfirmSelectedHandler(c *gin.Context) {
	Info, err := Auth.LoginInfo(c)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "您未登录或无权限",
		})
		return
	}
	selections, err := Selected.GetSelectedByTutorId(Info["info_id"].(int64), nil)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	for i := range selections {
		selections[i].Student, err = Student.GetStudentById(selections[i].Sid, nil)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 404,
				"msg":  "数据库异常",
			})
			return
		}
		selections[i].Topic, err = Topic.GetTopicById(selections[i].Tid, nil)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 404,
				"msg":  "数据库异常",
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": selections,
	})
}

func SetTutorConfirmSelectedHandler(c *gin.Context) {
	tid, err1 := strconv.ParseInt(c.PostForm("tid"), 10, 64)
	sid, err2 := strconv.ParseInt(c.PostForm("sid"), 10, 64)
	status, err3 := strconv.ParseInt(c.PostForm("status"), 10, 8)
	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据异常",
		})
		return
	}
	err := Selected.SetTutorComfirm(tid, sid, int8(status), nil)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "确认完成",
	})
}

func SetManagerConfirmSelectedHandler(c *gin.Context) {
	tid, err1 := strconv.ParseInt(c.PostForm("tid"), 10, 64)
	sid, err2 := strconv.ParseInt(c.PostForm("sid"), 10, 64)
	status, err3 := strconv.ParseInt(c.PostForm("status"), 10, 8)
	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据异常",
		})
		return
	}
	err := Selected.SetManagerComfirm(tid, sid, int8(status), nil)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "确认完成",
	})
}

func SetSelectedPublishedHandler(c *gin.Context) {
	tid, err1 := strconv.ParseInt(c.PostForm("tid"), 10, 64)
	sid, err2 := strconv.ParseInt(c.PostForm("sid"), 10, 64)
	if err1 != nil || err2 != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据异常",
		})
		return
	}
	tx, err := DBUtil.StartTx()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	id, err := Process.CreateProcess(tx)
	if err != nil {
		tx.Rollback()
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	err = Selected.SetPublished(tid, sid, id, 1, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	tx.Commit()
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "发布成功",
	})
}
func SetProcessProp(c *gin.Context) {
	pid, err1 := strconv.ParseInt(c.PostForm("pid"), 10, 64)
	status, err2 := strconv.ParseInt(c.PostForm("status"), 10, 64)
	prop := c.PostForm("prop")
	detail := c.PostForm("detail")
	if err1 != nil && err2 != nil && detail != "" {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据异常",
		})
		return
	}
	if !checkSetAuth(c, pid) {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "权限异常",
		})
		return
	}
	tx, err := DBUtil.StartTx()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	err = Process.UpdateProp(prop, detail, pid, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	err = Process.UpdateStatus(int(status), pid, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "操作成功",
	})
}

func GetProcess(c *gin.Context) {
	pid, err1 := strconv.ParseInt(c.Query("pid"), 10, 64)
	if err1 != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据异常",
		})
		return
	}
	if !checkSetAuth(c, pid) {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "权限异常",
		})
		return
	}
	process, err := Process.GetProcessByPid(pid, nil)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": process,
	})
}

func checkSetAuth(c *gin.Context, pid int64) bool {
	Info, err := Auth.LoginInfo(c)
	if err != nil {
		return false
	}
	tid, _, err := Process.LookUpSidAndTidByPid(pid, nil)
	if err != nil {
		return false
	}
	if tid != Info["info_id"].(int64) {
		return false
	} else {
		return true
	}
}
func SetProcessStatus(c *gin.Context) {
	pid, err1 := strconv.ParseInt(c.PostForm("pid"), 10, 64)
	status, err2 := strconv.ParseInt(c.PostForm("status"), 10, 64)
	if err1 != nil && err2 != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据异常",
		})
		return
	}
	if !checkSetAuth(c, pid) {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "权限异常",
		})
		return
	}
	tx, err := DBUtil.StartTx()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	err = Process.UpdateStatus(int(status), pid, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "操作成功",
	})
}
func ModifyTopicHandler(c *gin.Context) {
	var topic Topic.Topic
	c.Bind(&topic)
	Info, err := Auth.LoginInfo(c)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "您未登录或无权限",
		})
		return
	}
	tid, err := strconv.ParseInt(c.PostForm("tid"), 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "题目号错误",
		})
		return
	}
	oldTopic,err :=Topic.GetTopicById(tid,nil)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	if oldTopic.Tutor.Id != Info["info_id"].(int64){
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "您未登录或无权限",
		})
		return
	}
	if  topic.Passed == 1{
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "该选题已不可修改",
		})
		return
	}
	topic.Id = tid
	topic.Tutor.Id = Info["info_id"].(int64)
	err = Topic.UpdateTopic(topic, nil)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "修改失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"msg":   "创建成功",
		"topic": topic,
	})
}
func UpdatePersonalInfo(c *gin.Context) {
	var tutor Tutor.Tutor
	c.Bind(&tutor)
	err := Tutor.UpdateTutor(tutor, nil)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "操作成功",
		})
	}
}
