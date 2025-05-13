package main

import (
	"log"
	"github.com/ariel-madrid/agent-service-go/controller"
	"github.com/ariel-madrid/agent-service-go/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	utils.LoadEnv()
	gin.SetMode(os.Getenv("GIN_MODE"))
	app := gin.Default()
	gin.DefaultWriter = io.MultiWriter(os.Stdout, log.Writer())
	app.Use(middleware.CorsMiddleware())

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Servicio no encontrado."})
	})
	routes.InitRoutes(app)
	http.ListenAndServe(os.Getenv("ADDR"), app)
}
