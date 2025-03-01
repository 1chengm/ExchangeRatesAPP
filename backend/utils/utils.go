package utils

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
	"fmt"
)

func HashPassword(pwd string) (string , error){
	hash ,err := bcrypt.GenerateFromPassword([]byte(pwd), 12)

	return string(hash),err
}
func GenerateJWT(username string)(string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"username":username,
		"exp":time.Now().Add(time.Hour * 72).Unix(),
	})
	SignedToken,err := token.SignedString([]byte("secret"))
	return "Bearer " + SignedToken,err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ParseJWT(tokenString string)(string, error){
	if len(tokenString)>7 && tokenString[:7] == "Bearer "{
		tokenString = tokenString[7:]
	}
	token ,err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,fmt.Errorf("unexpected signing method: %v",t.Header["alg"])
		}
		return []byte("secret"),nil
	})

	if err != nil {
		return "",err
	}
	if claims , ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return "",fmt.Errorf("username not found")
		}
		return username,nil
	}
	return "", err
}