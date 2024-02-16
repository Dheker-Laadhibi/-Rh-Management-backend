package notifications

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitializeDB(database *gorm.DB) {
	db = database
}

func GetNotifications(c *gin.Context) {
	var notifications []Notification
	db.Find(&notifications)
	c.JSON(200, notifications)
}

func CreateNotification(c *gin.Context) {
	var notification Notification
	if err := c.BindJSON(&notification); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Create(&notification)
	c.JSON(200, notification)
}

func UpdateNotification(c *gin.Context) {
	id := c.Param("id")
	var notification Notification
	if err := db.Where("idnotif = ?", id).First(&notification).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Notification not found"})
		return
	}
	if err := c.BindJSON(&notification); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Save(&notification)
	c.JSON(200, notification)
}

func DeleteNotification(c *gin.Context) {
	id := c.Param("id")
	var notification Notification
	db.Where("idnotif = ?", id).Delete(&notification)
	c.JSON(200, gin.H{"message": "Notification deleted"})
}

func GetNotificationByID(c *gin.Context) {
	NotificationID := c.Param("id")
	var notification Notification
	if err := db.First(&notification, "idnotif = ?", NotificationID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Notification not found"})
		return
	}
	c.JSON(200, notification)
}
