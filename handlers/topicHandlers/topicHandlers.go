package topichandlers

import (
	Topic "backend/database/topic"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTopicByIdHandler(c *gin.Context){
	id := c.Query("tid")
	tid,err := strconv.ParseInt(id,10,64)
	if err!=nil{
		c.JSON(200,gin.H{
			"code":404,
			"msg":"题目不存在或未发布",
		})
		return
	}
	topic,err := Topic.GetTopicById(tid,nil)
	if err!=nil {
		c.JSON(200,gin.H{
			"code":404,
			"msg":"数据库错误",
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"data":topic,
	})
}

func GetPublishedTpoicHandler(c *gin.Context){
	topics,err := Topic.GetTopicListByPublishStatus(1,nil)
	if err!=nil {
		c.JSON(200,gin.H{
			"code":404,
			"msg":"数据库错误",
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"data":topics,
	})
}

func GetAllTopicHandler(c *gin.Context){
	topics,err := Topic.GetTopicList(nil)
	if err!=nil {
		c.JSON(200,gin.H{
			"code":404,
			"msg":"数据库错误",
		})
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"data":topics,
	})
}

func SetTopicPassedStatusHandler(c *gin.Context){
	id := c.PostForm("id")
	status := c.PostForm("status")
	tid,err1 := strconv.ParseInt(id,10,64)
	passedStatus,err2 := strconv.ParseInt(status,10,8)
	if err1!=nil || err2 != nil{
		c.JSON(200,gin.H{
			"code":404,
			"msg":"题目不存在",
		})
		return
	}
	var topic Topic.Topic = Topic.Topic{
		Id:tid,
		Passed: int8(passedStatus),
	}
	err := Topic.SetPassedStatus(topic,nil)
	if err != nil {
		c.JSON(200,gin.H{
			"code":404,
			"msg":"数据库错误",
		})
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":"设置成功",
	})
}
func SetTopicPublishedStatusHandler(c *gin.Context){
	id := c.PostForm("id")
	status := c.PostForm("status")
	tid,err1 := strconv.ParseInt(id,10,64)
	passedStatus,err2 := strconv.ParseInt(status,10,8)
	if err1!=nil || err2 != nil{
		c.JSON(200,gin.H{
			"code":404,
			"msg":"题目不存在",
		})
		return
	}
	var topic Topic.Topic = Topic.Topic{
		Id:tid,
		Published: int8(passedStatus),
	}
	err := Topic.SetPublishedStatus(topic,nil)
	if err != nil {
		c.JSON(200,gin.H{
			"code":404,
			"msg":"数据库错误",
		})
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":"设置成功",
	})
}

