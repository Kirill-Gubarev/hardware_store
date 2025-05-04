package service

import (
	"github.com/gin-gonic/gin"
	"address_service/internal/utils"
	"address_service/internal/db"
	"log"
)

var APIKey string

func Run(){
	utils.LogInit()

	err := db.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	
	routesRun()
}

func routesRun(){
	cfg, err := utils.GetServiceConfig()
	if err != nil{
		log.Fatalln(err)
	}

	var r = gin.Default()

	r.GET("/api/suggest", SuggestAddress)
	r.GET("/api/address/:id", GetAddress)
	r.GET("/api/distance", GetDistance)
	r.POST("/api/address", CreateAddress)

	APIKey = cfg.APIKey
	r.Run(cfg.Host + ":" + cfg.Port)
}
