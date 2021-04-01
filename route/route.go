package route

import (
	"github.com/algol2302/go-admin-api/auth"
	"github.com/algol2302/go-admin-api/controller"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	authMiddleware, err := auth.SetupAuth()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to my Admin App")
	})

	v1 := router.Group("/v1")
	{
		v1.POST("/login", authMiddleware.LoginHandler)

		v1.POST("/register", controller.RegisterEndPoint)

		todo := v1.Group("todo")
		{
			todo.GET("/all", authMiddleware.MiddlewareFunc(), controller.FetchAllUsers)
			todo.GET("/get/:id", authMiddleware.MiddlewareFunc(), controller.FetchSingleUser)
			todo.PUT("/update/:id", authMiddleware.MiddlewareFunc(), controller.UpdateUser)
			todo.DELETE("/delete/:id", authMiddleware.MiddlewareFunc(), controller.DeleteUser)
		}
	}

	authorization := router.Group("/auth")
	authorization.GET("/refresh_token", authMiddleware.RefreshHandler)

	return router
}
