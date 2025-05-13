package controller

import (
	"net/http"
	"github.com/ariel-madrid/agent-service-go/services"
	"github.com/gin-gonic/gin"
)

func GetLocales(c *gin.Context) {

	locales, err := services.GetDatosLocales()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, locales)
}	
