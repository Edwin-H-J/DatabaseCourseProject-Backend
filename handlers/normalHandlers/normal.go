package normal

import (
	"bytes"
	"os/exec"
)
import (
	Process "backend/database/process"
	Selected "backend/database/selected"
	Student "backend/database/student"
	Topic "backend/database/topic"
	Tutor "backend/database/tutor"
	User "backend/database/user"
	DBUtil "backend/database/util"
	Auth "backend/middleware/auth"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "用户名或密码为空",
		})
	} else {
		if Auth.Login(c, username, password) {
			info, err := Auth.LoginInfo(c)
			if err != nil {
				c.JSON(200, gin.H{
					"code": 404,
					"msg":  "登录失败,请检查是否开启cookie",
				})
				return
			}
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "登录成功",
				"info": info,
			})
		} else {
			c.JSON(200, gin.H{
				"code": 401,
				"msg":  "登录失败,账户或密码错误或账户被禁用",
			})
		}
	}
}

func RegisterHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	name := c.PostForm("name")
	auth := c.PostForm("auth")
	if username == "" || password == "" || name == "" || auth == "" {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "信息不全",
		})
		return
	}
	tx, err := DBUtil.StartTx()
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库错误",
			"err":  err,
		})
		return
	}
	var info_id int64
	if auth == "student" {
		var student Student.Student
		c.Bind(&student)
		student, err := Student.CreateStudent(student, tx)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			c.JSON(200, gin.H{
				"code": 404,
				"msg":  "数据库错误",
				"err":  err,
			})
			return
		}
		info_id = student.Id
	} else if auth == "tutor" || auth == "manager" {
		var tutor Tutor.Tutor
		c.Bind(&tutor)
		tutor, err := Tutor.CreateTutor(tutor, tx)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			c.JSON(200, gin.H{
				"code": 404,
				"msg":  "数据错误，已有相同用户",
				"err":  err,
			})
			return
		}
		info_id = tutor.Id
	}
	user := User.User{Id: 0, Username: username, Password: password, Auth: auth, InfoId: info_id}
	user, err = User.CreateUser(user, tx)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库错误",
			"err":  err,
		})
		return
	}
	err = tx.Commit()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库错误",
			"err":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "创建成功,请等待审核",
	})
}
func CheckLoginHandle(c *gin.Context) {
	info, err := Auth.LoginInfo(c)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "未登录,请检查是否开启cookie",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"info": info,
	})
}
func LogoutHandle(c *gin.Context) {
	Auth.Logout(c)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "登出成功",
	})
}

func LookupBaseInfo(c *gin.Context) {
	info, err := Auth.LoginInfo(c)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "未登录,请检查是否开启cookie",
		})
		return
	}
	data := gin.H{
		"code": 200,
	}
	auth := info["auth"].(string)
	if auth == "student" {
		student, err := Student.GetStudentById(info["info_id"].(int64), nil)
		if err != nil {
			c.JSON(404, gin.H{
				"code": 404,
				"msg":  "数据库错误",
			})
			return
		}
		data["student"] = student
		selected, err := Selected.GetSelectedBySid(info["info_id"].(int64), nil)
		if err != nil {
			c.JSON(200, data)
			return
		}
		data["selected"] = selected
		if selected.ProcessId != nil {
			process, err := Process.GetProcessByPid(selected.ProcessId.(int64), nil)
			if err == nil {
				data["process"] = process.ProcessStatus
			}
		}
		c.JSON(200, data)
		return
	} else if auth == "tutor" || auth == "manager" {
		tutor, err := Tutor.GetTutorById(info["info_id"].(int64), nil)
		if err != nil {
			c.JSON(404, gin.H{
				"code": 404,
				"msg":  "数据库错误",
			})
			return
		}
		data["tutor"] = tutor
		selecteds, err := Selected.GetSelectedByTutorId(info["info_id"].(int64), nil)
		if err == nil {
			data["selecteds"] = selecteds
		}
		if auth == "manager" {
			ConfirmSelects, err := Selected.GetTutorConfirmSelected(nil)
			if err == nil {
				data["ConfirmSelects"] = ConfirmSelects
			}
			ConfirmTopics, err := Topic.GetTopicList(nil)
			if err == nil {
				data["ConfirmTopics"] = ConfirmTopics
			}

		}
		c.JSON(200, data)
		return

	} else if auth == "admin" {
		users, err := User.GetAllUser(nil)
		if err != nil {
			c.JSON(404, gin.H{
				"code": 404,
				"msg":  "数据库错误",
			})
		}
		data["users"] = users
		c.JSON(200, data)
		return
	}

}

func LookUpPersonalInfo(c *gin.Context) {
	info, err := Auth.LoginInfo(c)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "未登录,请检查是否开启cookie",
		})
		return
	}
	var uId int64
	uid, err := strconv.ParseInt(c.Query("uid"), 10, 64)
	if err != nil {
		uId = info["id"].(int64)
	} else {
		uId = uid
	}
	user, err := User.GetUserById(uId, nil)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "未登录,请检查是否开启cookie",
		})
		return
	}
	user.Password = ""
	auth := user.Auth
	data := gin.H{
		"code": 200,
		"user": user,
	}
	if auth == "student" {
		student, err := Student.GetStudentById(user.InfoId, nil)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 404,
				"msg":  "数据库错误",
			})
			return
		}
		data["info"] = gin.H{
			"id":          student.Id,
			"name":        student.Name,
			"sex":         student.Sex,
			"major":       student.Major,
			"class":       student.Class,
			"phoneNumber": student.PhoneNumber,
			"Email":       student.Email,
			"remark":      student.Remark,
		}
	} else if auth == "tutor" || auth == "manager" {
		tutor, err := Tutor.GetTutorById(user.InfoId, nil)
		if err != nil {
			c.JSON(404, gin.H{
				"code": 404,
				"msg":  "数据库错误",
			})
			return
		}
		data["info"] = gin.H{
			"id":                    tutor.Id,
			"name":                  tutor.Name,
			"sex":                   tutor.Sex,
			"birthday":              tutor.Birthday,
			"EducationalBackground": tutor.EducationalBackground,
			"title":                 tutor.Title,
			"ResearchDirection":     tutor.ResearchDirection,
			"PhoneNumber":           tutor.PhoneNumber,
			"Email":                 tutor.Email,
		}
	}
	c.JSON(200, data)
}

func ModifyPassword(c *gin.Context) {
	info, err := Auth.LoginInfo(c)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "未登录,请检查是否开启cookie",
		})
		return
	}
	uid, err := strconv.ParseInt(c.PostForm("uid"), 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "请提供uid",
		})
	}
	if info["id"].(int64) != uid && info["auth"].(string) != "admin" {
		c.JSON(200, gin.H{
			"code": 401,
			"msg":  "您无权限",
		})
	}
	err = User.UpdatePassword(uid, c.PostForm("password"), nil)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "修改失败",
		})
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})

}

func GetUserList(c *gin.Context) {
	users, err := User.GetAllUser(nil)
	if err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "数据库错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": users,
	})
}

func SetUserStatus(c *gin.Context){
	uid, err1 := strconv.ParseInt(c.PostForm("uid"), 10, 64)
	status, err2 := strconv.ParseInt(c.PostForm("status"), 10, 64)
	if err1 != nil || err2 != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "uid错误",
		})
	}
	err := User.UpdateUserStatus(uid,status,nil)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库错误",
		})
	}
	c.JSON(200, gin.H{
			"code": 200,
			"msg":  "操作成功",
		})
}


func Backup(c *gin.Context){
	var cmd = exec.Command("mysqldump", "-uroot", "-pjh6679386" ,"graduate_manage")
	var out bytes.Buffer
	cmd.Stdout = &out
	err:=cmd.Run()
	if err!=nil{
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "备份失败",
		})
	}
	c.Writer.WriteString(out.String())
}
func SetUserAuth(c *gin.Context){
	uid, err1 := strconv.ParseInt(c.PostForm("uid"), 10, 64)
	status, err2 := strconv.ParseInt(c.PostForm("status"), 10, 64)
	if err1 != nil || err2 != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "uid错误",
		})
	}
	var err error
	if status == 1{
		err = User.UpdateUserManager(uid,nil)
	}else{
		err = User.UpdateUserTutor(uid,nil)
	}
	
	if err != nil {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库错误",
		})
	}
	c.JSON(200, gin.H{
			"code": 200,
			"msg":  "操作成功",
		})
}
func LookupUserId(c *gin.Context){
	sid, err1 := strconv.ParseInt(c.Query("sid"), 10, 64)
	var id int64
	if err1 != nil{
		tid, err1 := strconv.ParseInt(c.Query("tid"), 10, 64)
		if err1 != nil{
			c.JSON(200, gin.H{
			"code": 404,
			"msg":  "uid错误",
		})
		}
		id,err1 = User.LookupUserIdByTid(tid,nil)
		if err1 != nil{
			c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库错误",
		})
	}
	}else{
		id,err1 = User.LookupUserIdBySid(sid,nil)
		if err1 != nil{
			c.JSON(200, gin.H{
			"code": 404,
			"msg":  "数据库错误",
		})
		}
	}
	c.JSON(200, gin.H{
			"code": 200,
			"msg":  "操作成功",
			"id":id,
		})
	

}