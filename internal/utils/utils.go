package utils

import (
	"encoding/json"
	"net/http"
	"fmt"
	"os"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
)

func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(w http.ResponseWriter, data map[string] interface{})  {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if bearerToken == "" { //Token is missing, returns with error code 403 Unauthorized
		keys, ok := r.URL.Query()["Baerer"]
		if !ok || len(keys[0]) < 1 { //Token is missing, returns with error code 403 Unauthorized
			return ""
		} else {
			bearerToken = "Baerer " + keys[0] //Grab the token from the header
		}
	}

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(r *http.Request) (string, error) {

	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid := claims["UserId"]
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%v", uid), nil
	}
	return "", nil
}