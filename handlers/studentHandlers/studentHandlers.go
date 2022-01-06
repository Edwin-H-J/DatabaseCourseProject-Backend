package studenthandlers

import (
	Process "backend/database/process"
	Selected "backend/database/selected"
	Auth "backend/middleware/auth"
	Student "backend/database/student"
	"strconv"
	DBUtil "backend/database/util"
	"github.com/gin-gonic/gin"
)

func SelectTopicHandler(c *gin.Context) {
	Info, err := Auth.LoginInfo(c)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "您未登录或无权限",
		})
		return
	}
	_, err = Selected.GetSelectedBySid(Info["info_id"].(int64), nil)
	if err == nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "已存在选择",
		})
		return
	}
	tid, err := strconv.ParseInt(c.PostForm("tid"), 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "题目编号有误",
		})
		return
	}
	err = Selected.CreateSelected(tid, Info["info_id"].(int64), nil)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "数据库异常",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "选择成功",
	})
}
func CancelSelectTopicHandler(c *gin.Context) {
	Info, err := Auth.LoginInfo(c)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "您未登录或无权限",
		})
		return
	}
	selected, err := Selected.GetSelectedBySid(Info["info_id"].(int64), nil)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	if selected.TutorCheck != 0 && selected.ManagerCheck != 0 && selected.Published!=0{
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "已被上级确认，不可取消",
		})
		return
	}
	err = Selected.DeleteSelected(selected.Tid, selected.Sid, nil)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "取消成功",
	})
}

func GetSelectedBySidHandler(c *gin.Context){
	Info, err := Auth.LoginInfo(c)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "您未登录或无权限",
		})
		return
	}
	selected, err := Selected.GetSelectedBySid(Info["info_id"].(int64), nil)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "未选择",
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"data":selected,
	})
}

func GetMyProcess(c *gin.Context){
	Info, err := Auth.LoginInfo(c)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "您未登录或无权限",
		})
		return
	}
	selected,err := Selected.GetSelectedBySid(Info["info_id"].(int64),nil)
	if err!=nil || selected.ProcessId == nil{
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "无论文进程",
		})
		return
	}
	process,err := Process.GetProcessByPid(selected.ProcessId.(int64),nil)
	if err!=nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库错误",
		})
		return
	}
	c.JSON(200, gin.H{
			"code": 200,
			"data":  process,
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
	err = Process.UpdateProp(prop, detail,pid, tx)
	if err != nil {
		tx.Rollback()
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return
	}
	err = Process.UpdateStatus(int(status),pid, tx)
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

func checkSetAuth(c *gin.Context, pid int64) bool {
	Info, err := Auth.LoginInfo(c)
	if err != nil {
		return false
	}
	_, sid, err := Process.LookUpSidAndTidByPid(pid, nil)
	if err != nil {
		return false
	}
	if sid != Info["info_id"].(int64) {
		return false
	} else {
		return true
	}
}

func UpdatePersonalInfo(c *gin.Context) {
	var student Student.Student
	c.Bind(&student)
	err := Student.UpdateStudent(student,nil)
	if err != nil{
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库异常",
		})
		return 
	}else{
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "操作成功",
		})
	}
}