package main

import (
	"log"
	"server/db"
	"server/db/sqlc"
	"server/internal/document"
	"server/internal/user"
	"server/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	queries := sqlc.New(dbConn.GetDB())

	userRep := sqlc.NewUserRepository(queries)
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	documentRep := sqlc.NewDocumentRepository(queries, dbConn.GetDB())
	documentSvc := document.NewService(documentRep)
	documentHandler := document.NewHandler(documentSvc)

	router.InitRouter(userHandler, documentHandler)
	router.Start("0.0.0.0:8080")
}
