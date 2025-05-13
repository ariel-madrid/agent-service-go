package routes

import (
	"repo/controller"
	"repo/middleware"
	"github.com/gin-gonic/gin"
)

// Estas son las rutas del backend - Los endpoints a los que el Frontend va a consultar po informacion
func InitLocalesRoutes(r *gin.Engine) {
	AlumnoGroup := r.Group("/api/v1/buscador-estudiantes")

	AlumnoGroup.GET("/locales", controller.GetPlanesAlumno)
	
}