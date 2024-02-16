package condidatures

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitializeDB(database *gorm.DB) {
	db = database
}

func GetCondidatures(c *gin.Context) {
	var condidatures []Condidature
	db.Find(&condidatures)
	c.JSON(200, condidatures)
}

func CreateCondidature(c *gin.Context) {
	var condidature Condidature
	if err := c.BindJSON(&condidature); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Create(&condidature)
	c.JSON(200, condidature)
}

func UpdateCondidature(c *gin.Context) {
	id := c.Param("id")
	var condidature Condidature
	if err := db.Where("idcondidature = ?", id).First(&condidature).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Candidature not found"})
		return
	}
	if err := c.BindJSON(&condidature); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Save(&condidature)
	c.JSON(200, condidature)
}

func DeleteCondidature(c *gin.Context) {
	id := c.Param("id")
	var condidature Condidature
	db.Where("idcondidature = ?", id).Delete(&condidature)
	c.JSON(200, gin.H{"message": "Candidature deleted"})
}

func GetCondidatureByID(c *gin.Context) {
	id := c.Param("id")
	var condidature Condidature
	if err := db.First(&condidature, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Candidature not found"})
		return
	}
	c.JSON(200, condidature)
}
