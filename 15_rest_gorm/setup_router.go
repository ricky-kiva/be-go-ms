package main

import (
	"15_rest_gorm/handler"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func setupRouter(r *gin.Engine, db *gorm.DB, err error) {
	if err != nil {
		panic("DB connection failed")
	}
	fmt.Println("DB connection success")

	migrateDB(db)

	r.POST("/aot", func(ctx *gin.Context) { handler.PostHandler(ctx, db) })
	r.GET("/aot", func(ctx *gin.Context) { handler.GetAllHandler(ctx, db) })
}
