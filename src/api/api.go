package api

import (
	"fmt"
	"golang-project/src/api/middlewares"
	"golang-project/src/api/routers"
	"golang-project/src/api/validations"
	"golang-project/src/config"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok{
		val.RegisterValidation("mobile",validations.IranianMobileNumberValidator,true)
		val.RegisterValidation("password",validations.PasswordValidator,true)
	}
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
