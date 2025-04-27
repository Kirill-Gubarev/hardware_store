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

	r.GET("/units/:id", GetUnit)
	r.GET("/units", GetUnits)
	r.POST("/units", CreateUnit)
	r.DELETE("/units/:id", DeleteUnit)

	r.GET("/products/:id", GetProduct)
	r.GET("/products", GetProducts)
	r.POST("/products", CreateProduct)
	r.DELETE("/products/:id", DeleteProduct)

	r.Run(cfg.Host + ":" + cfg.Port)
}
