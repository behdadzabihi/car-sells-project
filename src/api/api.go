package api

import (
	"fmt"
	"golang-project/src/api/middlewares"
	"golang-project/src/api/routers"
	"golang-project/src/api/validations"
	"golang-project/src/config"
	"golang-project/src/docs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()

	RegisterRoutes(r,cfg)
	RegisterValidation()
	RegisterSwagger(r,cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}
}


func RegisterValidation(){
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok{
		val.RegisterValidation("mobile",validations.IranianMobileNumberValidator,true)
		val.RegisterValidation("password",validations.PasswordValidator,true)
	}

}

func RegisterSwagger(r *gin.Engine,cfg *config.Config){
	docs.SwaggerInfo.Title="golang web api"
	docs.SwaggerInfo.Description="golang web api"
	docs.SwaggerInfo.Version="1.0"
	docs.SwaggerInfo.BasePath="/api"
	docs.SwaggerInfo.Host=fmt.Sprintf("localhost:%s",cfg.Server.Port)
	docs.SwaggerInfo.Schemes=[]string{"http"}
	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
}