package controller

import (
	"net/http"
	"repo/internal/services"
	"github.com/gin-gonic/gin"
)

func GetLocales(c *gin.Context) {

	locales, err := services.GetDatosLocales()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, locales)
}	
