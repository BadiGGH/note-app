package handlers

import (
	docs "github.com/BadiGGH/note-app/docs"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetRoutes() *gin.Engine {
	r := gin.Default()

	// Swagger metadata
	docs.SwaggerInfo.Title = "Note App API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "Simple CRUD API for note management"
	docs.SwaggerInfo.BasePath = "/"

	// Routes
	r.GET("/notes", ReadAllNotesHandler)
	r.GET("/notes/:id", ReadNoteByIdHandler)
	r.POST("/notes", CreateNewNote)
	r.PUT("/notes/:id", UpdateNoteHandler)
	r.DELETE("/notes/:id", DeleteNoteByIdHandler)

	// Swagger UI endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
