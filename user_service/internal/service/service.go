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

	r.GET("/roles/:id", GetRole)
	r.GET("/roles", GetRoles)
	r.POST("/roles", CreateRole)
	r.DELETE("/roles/:id", DeleteRole)

	r.GET("/users/:id", GetUser)
	r.GET("/users", GetUsers)
	r.POST("/users", CreateUser)
	r.DELETE("/users/:id", DeleteUser)

	r.POST("/login", LoginSession)
	r.POST("/logout", LogoutSession)
	r.POST("/verify", VerifySession)

	r.Run(cfg.Host + ":" + cfg.Port)
}
