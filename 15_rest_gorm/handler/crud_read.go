package handler

import (
	"15_rest_gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetAllHandler(c *gin.Context, db *gorm.DB) {
	var soldiers []models.TitanSoldier

	db.Find(&soldiers)
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   soldiers,
	})
}
