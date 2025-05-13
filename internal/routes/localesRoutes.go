package routes

import (
	"github.com/ariel-madrid/agent-service-go/controller"
	"github.com/gin-gonic/gin"
)

// Estas son las rutas del backend - Los endpoints a los que el Frontend va a consultar po informacion
func InitLocalesRoutes(r *gin.Engine) {
	LocalesGroup := r.Group("/api/v1")

	LocalesGroup.GET("/locales", controller.GetLocales)
	
}