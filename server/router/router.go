package router

import (
	"time"

	"server/internal/document"
	"server/internal/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, documentHandler *document.Handler) {
	r = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Origin", "Content-Length", "Content-Type", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
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
