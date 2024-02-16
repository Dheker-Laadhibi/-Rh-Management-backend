package users

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"database/sql"

	helper "github.com/NourSammari/jwt/helpers"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	company "Test/company" 
)

func InitializeDB(database *sql.DB) {
	db = database
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("email or password is incorrect")
		check = false
	}
	return check, msg
}

/* func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		var validate = validator.New()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		if db == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection is nil"})
			return
		}

		// Check if email exists
		var emailExists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", user.Email).Scan(&emailExists)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking for the email"})
			return
		}
		if emailExists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email already exists"})
			return
		}

		// Check if phone exists
		var phoneExists bool
		err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE phone = $1)", user.Phone).Scan(&phoneExists)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking for the phone number"})
			return
		}
		if phoneExists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this phone number already exists"})
			return
		}

		password := HashPassword(user.Password)

		// Insert user into database
		stmt, err := db.Prepare("INSERT INTO users (email, password, phone, first_name, last_name, user_type, created_at, updated_at, user_id, token, refresh_token) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10,$11)")
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while preparing statement"})
			return
		}
		defer stmt.Close()

		// Generate a UUID for User_id
		User_id := uuid.New()
		user.User_id = User_id.String()

		// Generate tokens
		token, refreshToken, err := helper.GenerateAllTokens(user.Email, user.First_name, user.Last_name, user.User_type, user.User_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while generating tokens"})
			return
		}

		user.Token = token
		user.Refresh_token = refreshToken

		_, err = stmt.Exec(user.Email, password, user.Phone, user.First_name, user.Last_name, user.User_type, time.Now(), time.Now(), user.User_id, user.Token, user.Refresh_token)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while inserting user"})
			return
		}

		// Update tokens in the database
		err = helper.UpdateAllTokens(token, refreshToken, user.User_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating tokens"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	}
} */

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		var validate = validator.New()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		if db == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection is nil"})
			return
		}

		// Check if email exists
		var emailExists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", user.Email).Scan(&emailExists)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking for the email"})
			return
		}
		if emailExists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this email already exists"})
			return
		}

		// Check if phone exists
		var phoneExists bool
		err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE phone = $1)", user.Phone).Scan(&phoneExists)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking for the phone number"})
			return
		}
		if phoneExists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this phone number already exists"})
			return
		}

		password := HashPassword(user.Password)

		// Insert user into database
		stmt, err := db.Prepare("INSERT INTO users (email, password, phone, first_name, last_name, user_type, created_at, updated_at, user_id, token, refresh_token) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10,$11)")
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while preparing statement"})
			return
		}
		defer stmt.Close()

		// Generate a UUID for User_id
		User_id := uuid.New()
		user.User_id = User_id.String()

		// Generate tokens
		token, refreshToken, err := helper.GenerateAllTokens(user.Email, user.First_name, user.Last_name, user.User_type, user.User_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while generating tokens"})
			return
		}

		user.Token = token
		user.Refresh_token = refreshToken

		_, err = stmt.Exec(user.Email, password, user.Phone, user.First_name, user.Last_name, user.User_type, time.Now(), time.Now(), user.User_id, user.Token, user.Refresh_token)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while inserting user"})
			return
		}

		// Update tokens in the database
		err = helper.UpdateAllTokens(token, refreshToken, user.User_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating tokens"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	}
}

/* func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		var foundUser User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("Attempting login for email: %s", user.Email)

		err := db.QueryRow("SELECT id, email , password , phone , first_name, last_name,user_type,created_at, updated_at , token , refresh_token FROM users WHERE Email = $1", user.Email).Scan(
			&foundUser.ID,
			&foundUser.Email,
			&foundUser.Password,
			&foundUser.Phone,
			&foundUser.First_name,
			&foundUser.Last_name,
			&foundUser.User_type,
			&foundUser.Created_at,
			&foundUser.Updated_at,
			&foundUser.Token,
			&foundUser.Refresh_token,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(user.Password, foundUser.Password)
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, refreshToken, _ := helper.GenerateAllTokens(foundUser.Email, foundUser.First_name, foundUser.Last_name, foundUser.User_type, foundUser.User_id)
		helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)

		c.JSON(http.StatusOK, foundUser)
	}
}
*/

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		var foundUser User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Printf("Attempting login for email: %s", user.Email)

		err := db.QueryRow("SELECT id, email , password , phone , first_name, last_name,user_type,created_at, updated_at , token , refresh_token FROM users WHERE Email = $1", user.Email).Scan(
			&foundUser.ID,
			&foundUser.Email,
			&foundUser.Password,
			&foundUser.Phone,
			&foundUser.First_name,
			&foundUser.Last_name,
			&foundUser.User_type,
			&foundUser.Created_at,
			&foundUser.Updated_at,
			&foundUser.Token,
			&foundUser.Refresh_token,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
			return
		}

		passwordIsValid, msg := VerifyPassword(user.Password, foundUser.Password)
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		// Générer et mettre à jour les jetons
		token, refreshToken, _ := helper.GenerateAllTokens(foundUser.Email, foundUser.First_name, foundUser.Last_name, foundUser.User_type, foundUser.User_id)
		helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)

		// Renvoyer les jetons dans la réponse
		c.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": refreshToken})
	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {


		
		/*  if err := helper.CheckUserType(c,"ADMIN"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}  */

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage

		rows, err := db.Query("SELECT COUNT(*) AS total_count FROM users")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while counting user items"})
			return
		}
		defer rows.Close()

		var totalCount int
		if rows.Next() {
			if err := rows.Scan(&totalCount); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while scanning total count"})
				return
			}
		}

		rows, err = db.Query("SELECT * FROM users OFFSET $1 LIMIT $2", startIndex, recordPerPage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing user items"})
			return
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.Email, &user.First_name, &user.Last_name, &user.User_type, &user.Created_at, &user.Updated_at); err != nil {
				log.Printf("Error scanning row: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while scanning user"})
				return
			}

			users = append(users, user)
		}

		c.JSON(http.StatusOK, gin.H{"total_count": totalCount, "user_items": users})
	}
}

func GetUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")

        var user User
        err := db.QueryRow("SELECT ID, Email, First_name, Last_name, User_type, Created_at, Updated_at FROM users WHERE user_id = $1", id).
            Scan(&user.ID, &user.Email, &user.First_name, &user.Last_name, &user.User_type, &user.Created_at, &user.Updated_at)
        if err != nil {
            log.Printf("Error scanning row: %v", err)
            if err == sql.ErrNoRows {
                c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
                return
            }
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while retrieving user"})
            return
        }

        c.JSON(http.StatusOK, user)
    }
}
func UpdateUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")

        // Get user data from request body
        var updateUser User
        if err := c.BindJSON(&updateUser); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Update user information in the database
        result, err := db.Exec("UPDATE users SET email = $1, phone = $2, first_name = $3, last_name = $4, user_type = $5, updated_at = $6 WHERE user_id = $7",
            updateUser.Email, updateUser.Phone, updateUser.First_name, updateUser.Last_name, updateUser.User_type, time.Now(), id)
        if err != nil {
            log.Printf("Error updating user: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating user"})
            return
        }

        rowsAffected, err := result.RowsAffected()
        if err != nil {
            log.Printf("Error getting rows affected: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating user"})
            return
        }

        if rowsAffected == 0 {
            c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
    }
}
func DeleteUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")

        result, err := db.Exec("DELETE FROM users WHERE user_id = $1", id)
        if err != nil {
            log.Printf("Error deleting user: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while deleting user"})
            return
        }

        rowsAffected, err := result.RowsAffected()
        if err != nil {
            log.Printf("Error getting rows affected: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while deleting user"})
            return
        }

        if rowsAffected == 0 {
            c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
    }
}
func CreateUserWithCompany(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Lire les données de la requête JSON
        var userData User
        if err := c.ShouldBindJSON(&userData); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }

        // Vérifier si l'ID de l'entreprise est fourni
        companyId, err := strconv.ParseUint(c.Query("company_id"), 10, 64)
        if err != nil {
            c.JSON(400, gin.H{"error": "company_id must be provided"})
            return
        }

        // Convertir companyId en int
        companyIdInt := int(companyId)

        // Vérifier si l'entreprise existe
        var company company.Company
        if err := db.First(&company, companyIdInt).Error; err != nil {
            c.JSON(404, gin.H{"error": "company not found"})
            return
        }

        // Assigner l'ID de l'entreprise à l'utilisateur
        userData.CompanyId = companyIdInt

        // Créer l'utilisateur dans la base de données
        if err := db.Create(&userData).Error; err != nil {
            c.JSON(500, gin.H{"error": "failed to create user"})
            return
        }

        c.JSON(200, userData)
    }
}
