package routes

import (
	"TDList/api"
	"TDList/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret921122"))
	r.Use(sessions.Sessions("mySessions", store))
	v1 := r.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		authorized := v1.Group("/")
		authorized.Use(middleware.JWT())
		{
			authorized.POST("task", api.CreateTask)
			authorized.GET("task/:id", api.TaskDetail)
		}
	}
	return r
}
