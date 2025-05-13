package routes

import (
	"github.com/ariel-madrid/agent-service-go/controller"
	"github.com/ariel-madrid/agent-service-go/middleware"
	"github.com/gin-gonic/gin"
)

// Estas son las rutas del backend - Los endpoints a los que el Frontend va a consultar po informacion
func InitLocalesRoutes(r *gin.Engine) {
	AlumnoGroup := r.Group("/api/v1/buscador-estudiantes")

	AlumnoGroup.GET("/locales", controller.GetPlanesAlumno)
	
}