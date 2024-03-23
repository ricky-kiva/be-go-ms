package handler

import (
	"15_rest_gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func PostHandler(c *gin.Context, db *gorm.DB) {
	var newSoldier models.TitanSoldier

	// `Bind` automatically bind the 'request body' sent by 'request' to `newSoldier` struct
	if c.Bind(&newSoldier) == nil {
		db.Create(&newSoldier)
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   newSoldier,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Bind error",
		})
	}
}
