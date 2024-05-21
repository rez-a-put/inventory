package utils

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	projectDirName := "inventory"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		os.Exit(1)
	}
}

// GetEnvByKey : to get data from env files
func GetEnvByKey(key string) string {
	return os.Getenv(key)
}

// GetErrorMsg : to return error message
func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Please input right format of email"
	case "max":
		return "Maximum " + fe.Param() + " characters"
	case "min":
		return "Minimum " + fe.Param() + " characters"
	}
	return "Unknown error"
}

// VerifyPassword : to verify inputted password
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GenerateToken : to generate token data
func GenerateToken(user_id int64) (string, error) {
	var err error

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

// ExtractTokenID : to return token id data
func ExtractTokenID(ctx *gin.Context) (jwt.MapClaims, uint, error) {
	tokenString := extractToken(ctx)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return nil, 0, err
		}
		return claims, uint(uid), nil
	}
	return nil, 0, nil
}

func TokenValid(c *gin.Context) (err error) {
	tokenString := extractToken(c)
	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}

	return nil
}

// extractToken : to extract token from request handler authorization
func extractToken(ctx *gin.Context) string {
	// token := ctx.Query("token")
	// if token != "" {
	// 	return token
	// }
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
