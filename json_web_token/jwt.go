package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/prometheus/common/log"
)

func main() {
	fmt.Println("main!")

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// env variable

	token, err := encryptToken()
	if err != nil {
		logs.Warn("Error create token!")
	}

	bearerToken := "Bearer " + token
	fmt.Println("bearerToken")
	fmt.Println(bearerToken)
	fmt.Println("bearerToken")
	result, errAuthorize := authorizeToken(bearerToken)
	if errAuthorize {
		fmt.Println(errAuthorize)
	}

	fmt.Println("result is %v", result)
}

// encryptToken
func encryptToken() (string, error) {
	var err error

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user"] = "PT_indomarco_prismatama"
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

// authorizeToken ..
func authorizeToken(header string) (string, bool) {
	fmt.Println(">>> AuthorizeToken <<<")

	// decode auth token
	res, decodeAuthTokenSuccessful := decodeAuthToken(header)
	if !decodeAuthTokenSuccessful {
		return "customer", false
	}
	fmt.Println(res)
	return "customer", true
}

// decodeAuthToken ..
func decodeAuthToken(header string) (string, bool) {
	fmt.Println(">>> Decode Auth Token - Function <<<")

	var granted bool

	authToken, getAuthTokenSuccessful := getAuthToken(header)
	if !getAuthTokenSuccessful {
		return "", false
	}

	// decode auth token
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	fmt.Println("begin jwt.ParseWithClaims token")
	fmt.Println(token)
	fmt.Println(token.Valid)
	fmt.Println("end jwt.ParseWithClaims token")
	if err != nil {
		logs.Warn("Error ParseWithClaims %v", err.Error())
		return "", false
	}
	if !token.Valid {
		log.Warn("Warning token invalid!")
		return "", false
	}

	// checking claims
	if claims["authorized"] == true && claims["user"] == os.Getenv("GRANTED_USER") {
		granted = true
	} else {
		granted = false
	}

	fmt.Println("access granted is ==> %v", granted)
	res := fmt.Sprintf("access granted is ==> %v", granted)
	return res, true
}

// getAuthToken ..
func getAuthToken(header string) (string, bool) {
	fmt.Println(">>> Get Auth Token - Function <<<")

	if strings.TrimSpace(header) == "" {
		logs.Critical("Missing authorization token")
		return "", false
	}

	if !strings.HasPrefix(header, "Bearer ") {
		logs.Critical("Invalid bearer authorization token")
		return "", false
	}

	token := strings.TrimPrefix(header, "Bearer ")
	fmt.Println("begin of token")
	fmt.Println(token)
	fmt.Println("end of token")
	return token, true
}
