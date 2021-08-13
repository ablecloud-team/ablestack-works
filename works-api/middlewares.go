package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

var expirationTime = 5 * time.Minute

func SetHeader(c *gin.Context)  {
	//c.Header("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0, max-age=0")
	//c.Header("Last-Modified", time.Now().String())
	//c.Header("Pragma", "no-cache")
	//c.Header("Expires", "-1")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}

func createToken(userid string) (string,error){
	var JwtKey = []byte(os.Getenv("MoldSecretKey"))
	expirationTime := time.Now().Add(expirationTime)

	type WorksClaims struct{
		UserID string
		jwt.StandardClaims
	}

	claims := WorksClaims{
		userid,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func checkToken(c *gin.Context){
	var JwtKey = []byte(os.Getenv("MoldSecretKey"))
	token, err := c.Request.Cookie("access-token")
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"result": http.StatusUnauthorized,
			"message": "Authentication failed(token is None)",
			"error": err,
		})
		c.Abort()
		return
	}
	tknStr := token.Value
	if tknStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"message": "token is None",
			"error": err,
		})
		c.Abort()
		return
	}
	type WorksClaims struct{
		UserID string
		jwt.StandardClaims
	}
	at(time.Unix(0,0), func(){
		token, err := jwt.ParseWithClaims(tknStr, &WorksClaims{}, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
		if claims, ok := token.Claims.(*WorksClaims); ok && token.Valid {
			fmt.Printf("%v %v", claims.UserID, claims.StandardClaims.ExpiresAt)
			fmt.Println("Token check OK!!!")
			c.Set("cookie-user-id",claims.UserID)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"message": "token is expired",
				"error": err,
			})
			c.Abort()
			return
		}
	})
}

func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}
