package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skr-ninja/models"
)

// Save Certificate
func SaveCertDeatil(c *gin.Context) {

	var err error

	cert, err := models.ExtrctCert()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": cert})
}

func GetCertDeatil(c *gin.Context) {

	var err error

	var data []models.Certificate
	models.DB.Find(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}

// Save Certificate
func UsingFileCert(c *gin.Context) {

	form, _ := c.MultipartForm()
	files := form.File["file"]
   
	for _, file := range files {
		log.Println(file.Filename)
		err := c.SaveUploadedFile(file, "./temp/"+file.Filename)
		if err != nil {
			log.Fatal(err)
		}
	}
	var err error
      
	cert, err := models.UsingFile(files[0].Filename)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": cert})
}
