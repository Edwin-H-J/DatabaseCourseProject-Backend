package session

import (
	"log"
	"os"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)
var store redis.Store;

func Session() gin.HandlerFunc{
	sessionNames := []string{"auth", "normal"}
	return sessions.SessionsMany(sessionNames, store)
}
func init(){
	var err error;
	store, err = redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}