package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skr-ninja/controllers"
	"github.com/skr-ninja/models"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.GET("/saveCertificate", controllers.SaveCertDeatil)
	public.GET("/getCertficate", controllers.GetCertDeatil)
	public.POST("/uploadCertfile", controllers.UsingFileCert)
	r.Run(":8084")

}
