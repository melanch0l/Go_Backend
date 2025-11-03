package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const appSecretKey = "admin1234"

// func init() {
// 	loadEnv()
// }

//	func loadEnv() {
//		err := godotenv.Load(".env")
//		if err != nil {
//			// Use log.Fatalf to ensure the program stops if the key can't be loaded
//			log.Fatalf("Error loading .env file: %v", err)
//		}
//		appSecretKey = os.Getenv("SECRET_KEY")
//		if appSecretKey == "" {
//			log.Fatal("SECRET_KEY environment variable not set. Cannot run application.")
//		}
//	}
func GenerateToken(email string, userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), //valid for 2 hrs
	})
	return token.SignedString([]byte(appSecretKey))
}
func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC) //check method
		if !ok {
			return nil, errors.New("wrong signing method")
		}
		return []byte(appSecretKey), nil
	})
	if err != nil {
		return 0, errors.New("could not parse token")
	}
	isValidToken := parsedToken.Valid
	if !isValidToken {
		return 0, errors.New("invalid token")
	}
	claim, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claim")

	}
	// email := claim["email"].(string)
	userId := int64(claim["userId"].(float64))
	return userId, nil
}
