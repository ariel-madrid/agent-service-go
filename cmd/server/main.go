package main

import (
	"log"
	"definir/internal/controller"

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
