package users

import (
	"context"

	"database/sql"
	"errors"

	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}

func DBInstance() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=dheker dbname=GestionRH sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close() // Close the connection before returning the error
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL!")

	return db, nil
}

var db, err = DBInstance()

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
		User_type:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(), // Refresh token expires in 7 days
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	signedToken, err = token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Println("Error generating token:", err)
		return "", "", err
	}

	signedRefreshToken, err = refreshToken.SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Println("Error generating refresh token:", err)
		return "", "", err
	}

	return signedToken, signedRefreshToken, nil
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, `
    UPDATE users 
    SET token = $1, 
        refresh_token = $2, 
        updated_at = $3 
    WHERE user_id = $4`,
		signedToken, signedRefreshToken, time.Now(), userId)

	if err != nil {
		log.Println("Error updating tokens:", err)
		return err
	}

	return nil
}

// GenerateSalt génère un sel aléatoire sécurisé

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "the token is invalid"
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token is expired"
		return
	}

	return claims, ""
}
func CheckUserType(c *gin.Context, role string) error {
	// Retrieve user's role from context
	userRole, exists := c.Get("user_type")
	if !exists {
		return errors.New("user role not found in context")
	}

	// Check if user's role matches the required role
	if userRole != role {
		return errors.New("unauthorized access")
	}

	return nil
}
