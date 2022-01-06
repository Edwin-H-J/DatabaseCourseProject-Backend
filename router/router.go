package route

import (
	"github.com/gin-gonic/gin"
	normal "backend/handlers/normalHandlers"
	topic "backend/handlers/topicHandlers"
	Auth "backend/middleware/auth"
	tutor "backend/handlers/tutorHandlers"
	Student "backend/handlers/studentHandlers"
)

func SetRouter(r *gin.RouterGroup){
	r.POST("/login",normal.LoginHandler)
	r.POST("/register",normal.RegisterHandler)
	r.GET("/checkLogin",normal.CheckLoginHandle)
	r.GET("/logout",normal.LogoutHandle)
	r.GET("/getTopic",Auth.AuthCheck("student","tutor","manager"),topic.GetTopicByIdHandler)
	r.GET("/getAllTopic",Auth.AuthCheck("manager"),topic.GetTopicByIdHandler)
	r.GET("/getPublishedTopic",Auth.AuthCheck("manager"),topic.GetTopicByIdHandler)
	r.POST("/Topic",Auth.AuthCheck("tutor","manager"),tutor.CreateTopicHandler)
	r.GET("/myTopicList",Auth.AuthCheck("tutor","manager"),tutor.GetUserTopic)
	r.GET("/Topic",Auth.AuthCheck("student","tutor","manager"),topic.GetPublishedTpoicHandler)
	r.GET("/AllTopic",Auth.AuthCheck("manager"),topic.GetAllTopicHandler)
	r.POST("/setTopicPassedStatus",Auth.AuthCheck("manager"),topic.SetTopicPassedStatusHandler)
	r.POST("/setTopicPublishedStatus",Auth.AuthCheck("manager"),topic.SetTopicPublishedStatusHandler)
	r.POST("/selectTopicHandler",Auth.AuthCheck("student"),Student.SelectTopicHandler)
	r.POST("/cancelSelectTopicHandler",Auth.AuthCheck("student"),Student.CancelSelectTopicHandler)
	r.GET("/getMySelected",Auth.AuthCheck("student"),Student.GetSelectedBySidHandler)
	r.GET("/GetAllTutorConfirmSelected",Auth.AuthCheck("manager"),tutor.GetAllTutorConfirmSelectedHandler)
	r.GET("/GetTutorConfirmSelected",Auth.AuthCheck("tutor","manager"),tutor.GetTutorConfirmSelectedHandler)
	r.POST("/tutorConfirm",Auth.AuthCheck("tutor","manager"),tutor.SetTutorConfirmSelectedHandler)
	r.POST("/managerConfirm",Auth.AuthCheck("manager"),tutor.SetManagerConfirmSelectedHandler)
	r.POST("/publishSelected",Auth.AuthCheck("manager"),tutor.SetSelectedPublishedHandler)
	r.GET("/GetMyProcess",Auth.AuthCheck("student"),Student.GetMyProcess)
	r.GET("/GetProcess",Auth.AuthCheck("tutor","manager"),tutor.GetProcess)
	r.POST("/TutorSetProcessProp",Auth.AuthCheck("tutor","manager"),tutor.SetProcessProp)
	r.POST("/StudentSetProcessProp",Auth.AuthCheck("student"),Student.SetProcessProp)
	r.POST("/SetProcessStatus",Auth.AuthCheck("tutor","manager"),tutor.SetProcessStatus)
	r.GET("/LookupBaseInfo",normal.LookupBaseInfo)
	r.GET("/LookUpPersonalInfo",normal.LookUpPersonalInfo)
	r.POST("/modifyStudentInfo",Auth.AuthCheck("student","admin"),Student.UpdatePersonalInfo)
	r.POST("/modifyTutorInfo",Auth.AuthCheck("admin","tutor","manager"),tutor.UpdatePersonalInfo)
	r.POST("/ModifyPassword",Auth.AuthCheck("student","admin","tutor","manager"),normal.ModifyPassword)
	r.GET("/GetUserList",Auth.AuthCheck("admin"),normal.GetUserList)
	r.POST("/SetUserStatus",Auth.AuthCheck("admin"),normal.SetUserStatus)
	r.POST("/ModifyTopicHandler",Auth.AuthCheck("tutor","manager"),tutor.ModifyTopicHandler)
	r.GET("/Backup",Auth.AuthCheck("admin"),normal.Backup)
	r.POST("/SetTutorAuth",Auth.AuthCheck("admin"),normal.SetUserAuth)
	r.GET("/LookupUserId",normal.LookupUserId)
}