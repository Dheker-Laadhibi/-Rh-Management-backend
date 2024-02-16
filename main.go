package main

import (
	"Test/condidatures"
	"Test/notifications"
	"Test/roles"
	"Test/users"

	//"test/users"
	"Test/company"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func main() {
	// Connect to PostgreSQL database
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=GestionRH password=dheker sslmode=disable")

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()
	// Initialize the database connection in the roles package
	roles.InitializeDB(db)
	company.InitializeDB(db)
	notifications.InitializeDB(db)
	condidatures.InitializeDB(db)
	router := gin.Default()

	// Auto-migration pour créer automatiquement la table si elle n'existe pas
	db.AutoMigrate(&roles.Role{})

	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&condidatures.Condidature{})
	// Role Routes

	v1 := router.Group("/api")
	{
		Roles := v1.Group("/Roles")
		{

			Roles.GET("/GetALL", roles.GetRoles)            // Add forward slash
			Roles.POST("/Create", roles.CreateRole)         // Add forward slash
			Roles.PUT("/update/:id", roles.UpdateRole)      // Add forward slash
			Roles.GET("/GetOneRole/:id", roles.GetRoleByID) // Add forward slash
			Roles.DELETE("/delete/:id", roles.DeleteRole)   // Add forward slash

		}

		companies := v1.Group("/companies")
		{
			companies.GET("", company.GetCompanies)
			companies.GET("/:id", company.GetCompanyByID)
			companies.POST("", company.CreateCompany)
			companies.PUT("/:id", company.UpdateCompany)
			companies.DELETE("/:id", company.DeleteCompany)
		}

		notif := v1.Group("/notifications")
		{
			notif.GET("", notifications.GetNotifications)
			notif.GET("/:id", notifications.GetNotificationByID)
			notif.POST("", notifications.CreateNotification)
			notif.PUT("/:id", notifications.UpdateNotification)
			notif.DELETE("/:id", notifications.DeleteNotification)
		}

		u := v1.Group("/users")
		{

			u.POST("", users.Signup())
			u.POST("/user", users.Login())
			u.GET("", users.GetUsers())
			u.GET("/:id", users.GetUser())
			u.PUT("update/:id", users.UpdateUser())
			u.DELETE("delete/:id", users.DeleteUser())
		}

		Condidature := v1.Group("/Condidature")
{
    Condidature.GET("/GetALLc", condidatures.GetCondidatures)
    Condidature.POST("/Createc", condidatures.CreateCondidature)
    Condidature.PUT("/updatec/:id", condidatures.UpdateCondidature)
    Condidature.GET("/GetOneCondidaturec/:id", condidatures.GetCondidatureByID)
    Condidature.DELETE("/deletec/:id", condidatures.DeleteCondidature)
}


	}

	// nour work

	db.AutoMigrate(&company.Company{})
	db.AutoMigrate(&notifications.Notification{})

	//end of nour
	router.Run(":8080")
}
