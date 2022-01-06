package main
import (
    router "backend/router"
	session "backend/middleware/session"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(session.Session())
	basicGroup := r.Group("")
	router.SetRouter(basicGroup)
	r.Run(":8001")
}
