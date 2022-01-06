package auth

import (
	User "backend/database/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"errors"
)
func judgeAuth(a interface{},auth ...string) bool {
	if len(auth) == 0 {
		return true
	}
	var tmp_a string;
	if a == nil {
		return false
	}else{
		tmp_a = a.(string)
	}
	for _, v := range auth {
		if v == tmp_a {
			return true
		}
	}
	return false
}
func getAuthSession(c *gin.Context) sessions.Session{
	return sessions.DefaultMany(c,"auth")
}
func AuthCheck(authList ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := getAuthSession(c)
		auth := session.Get("auth")
		if judgeAuth(auth,authList...) {
			c.Next()
		}else{
			c.AbortWithStatusJSON(200,gin.H{
				"code":401,
				"msg":"无资源获取权限",
			})
		}
	}
}

func Login(c *gin.Context,username,password string) bool{
	user,err := User.GetUserByUsername(username,nil)
	if err != nil {
		return false
	}
	if user.Password == password {
		if user.Available != 1 {
			return false
		}
		session := getAuthSession(c)
		session.Set("id",user.Id)
		session.Set("auth",user.Auth)
		session.Set("info_id",user.InfoId)
		session.Save()
		return true
	}else{
		return false
	}
}
func Logout(c *gin.Context){
	session := getAuthSession(c)
	session.Clear()
	session.Save()
}
func LoginInfo(c *gin.Context) (gin.H,error){
	session := getAuthSession(c)
	id := session.Get("id")
	var err error = nil
	if id == nil {
		err = errors.New("账户未登录")
	}
	return gin.H{
		"id":id,
		"auth":session.Get("auth"),
		"info_id":session.Get("info_id"),
	},err
}