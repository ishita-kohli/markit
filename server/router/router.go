package router

import (
	"server/internal/document"
	"server/internal/user"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, documentHandler *document.Handler) {
	r = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://127.0.0.1:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	ar := r.Group("/", userHandler.AuthTokenMiddleware())

	ar.GET("/users", userHandler.Getuserlist)

	ar.POST("/document", documentHandler.CreateDocument)
	ar.GET("/document", documentHandler.Listdocuments)
	ar.GET("/document/:id", documentHandler.GetDocumentByID)
	ar.PATCH("/document/:id", documentHandler.UpdateDocument)
	ar.PATCH("/document/:id/share", documentHandler.ShareDocument)

}

func Start(addr string) error {
	return r.Run(addr)
}
