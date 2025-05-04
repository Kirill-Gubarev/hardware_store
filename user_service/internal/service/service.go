package service

import (
	"github.com/gin-gonic/gin"
	"user_service/internal/utils"
	"user_service/internal/db"
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

	r.GET("/api/roles/:id", GetRole)
	r.GET("/api/roles", GetRoles)
	r.POST("/api/roles", CreateRole)
	r.DELETE("/api/roles/:id", DeleteRole)

	r.GET("/api/users/:id", GetUser)
	r.GET("/api/users", GetUsers)
	r.POST("/api/users", CreateUser)
	r.DELETE("/api/users/:id", DeleteUser)

	r.POST("/api/login", LoginSession)
	r.POST("/api/logout", LogoutSession)
	r.POST("/api/verify", VerifySession)

	r.Run(cfg.Host + ":" + cfg.Port)
}
