// roles.go
package roles

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// InitializeDB initializes the database instance
func InitializeDB(database *gorm.DB) {
	db = database
}

// GetRoles retrieves all roles from the database.
// c *gin.Context contient la reponse hhtp request du serveur 
func GetRoles(c *gin.Context) {
	// db doit etre initializer
	//check
	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection not initialized"})
		return
	}

	var roles []Role
	db.Find(&roles)
	// retour en reponse json
	c.JSON(200, roles)
}

// CreateRole creates a new role in the database.
func CreateRole(c *gin.Context) {
	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection not initialized"})
		return
	}

	var role Role
	if err := c.BindJSON(&role); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Vérification de la contrainte de clé primaire
	var existingRole Role
	if db.Where("id = ?", role.ID).First(&existingRole).RecordNotFound() {
		// L'enregistrement avec cet ID n'existe pas, créer la nouvelle entité
		db.Create(&role)
		c.JSON(201, role)
	} else {
		// L'enregistrement avec cet ID existe déjà, renvoyer une erreur
		c.JSON(409, gin.H{"error": "Role ID already exists"})
	}
}

// UpdateRole updates an existing role in the database.
func UpdateRole(c *gin.Context) {
	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection not initialized"})
		return
	}

	// Get role ID from URL parameter
	roleID := c.Param("id")

	var role Role
	// Find the role by ID
	if db.First(&role, roleID).RecordNotFound() {
		c.JSON(404, gin.H{"error": "Role not found"})
		return
	}

	// Bind JSON data to role
	if err := c.BindJSON(&role); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Save the role
	db.Save(&role)
	c.JSON(200, role)
}

// DeleteRole deletes a role from the database.
/* func DeleteRole(c *gin.Context) {
	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection not initialized"})
		return
	}

	var role Role
	roleID := c.Param("id")

	if db.First(&role, roleID).RecordNotFound() {
		c.JSON(404, gin.H{"error": "Role not found"})
		return
	}

	db.Delete(&role)
	c.Status(204)
} */

// GetRoleByID retrieves a single role by its ID from the database.

func GetRoleByID(c *gin.Context) {
	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection not initialized"})
		return
	}

	// Get role ID from URL parameter
	roleID := c.Param("id")

	var role Role
	// Find the role by ID
	if err := db.First(&role, roleID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Role not found"})
		return
	}

	// Return the role
	c.JSON(200, role)
}

// DeleteRole deletes a role from the database.
func DeleteRole(c *gin.Context) {
	if db == nil {
		c.JSON(500, gin.H{"error": "Database connection not initialized"})
		return
	}

	// Get role ID from URL parameter
	roleID := c.Param("id")

	var role Role
	// Find the role by ID
	if err := db.First(&role, roleID).Error; err != nil {
		// Check if the error is due to record not found
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(404, gin.H{"error": "Role not found"})
			return
		}
		// Handle other errors
		c.JSON(500, gin.H{"error": "Error while querying the database"})
		return
	}

	// Delete the role
	if err := db.Delete(&role).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error while deleting the role"})
		return
	}

	c.Status(204)
}
