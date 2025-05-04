package service

import (
	"github.com/gin-gonic/gin"
	"product_service/internal/utils"
	"product_service/internal/db"
	"log"
)

func Run(){
	utils.LogInit()

	err := db.Connect()
	if err != nil{
		log.Fatalln(err)
	}
	defer db.Close()

	routesRun()
}

func routesRun(){
	cfg, err := utils.GetServiceConfig()
	if err != nil{
		log.Println(err)
	}

	var r = gin.Default()

	r.GET("/api/units/:id", GetUnit)
	r.GET("/api/units", GetUnits)
	r.POST("/api/units", CreateUnit)
	r.DELETE("/api/units/:id", DeleteUnit)

	r.GET("/api/products/:id", GetProduct)
	r.GET("/api/products", GetProducts)
	r.POST("/api/products", CreateProduct)
	r.DELETE("/api/products/:id", DeleteProduct)

	r.Run(cfg.Host + ":" + cfg.Port)
}
