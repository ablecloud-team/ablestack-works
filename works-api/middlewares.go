package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

var expirationTime = 5 * time.Minute
var allowUrl string

func SetHeader(c *gin.Context) {
	allowUrlList := []string{"http://localhost:8080", "http://172.16.0.85:8083", "http://172.16.0.89:8080", "http://172.16.0.85:8081"}
	for _, url := range allowUrlList {
		if c.Request.Header.Get("Origin") == url {
			allowUrl = url
			break
		}
	}
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", allowUrl)
	c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST, PUT, PATCH")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}

func createToken(userid string) (string, error) {
	var JwtKey = []byte(os.Getenv("MoldSecretKey"))
	expirationTime := time.Now().Add(expirationTime)

	type WorksClaims struct {
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

func checkToken(c *gin.Context) {
	result := map[string]interface{}{}
	var JwtKey = []byte(os.Getenv("MoldSecretKey"))

	tknStr := c.Request.Header.Get("Authorization")

	//token, err := c.Request.Cookie("access-token")
	if tknStr == "works" || tknStr == "" {
		result["status"] = http.StatusUnauthorized
		result["message"] = "token is None"
		result["errorCode"] = "9999"
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
		c.Abort()
		return
	}
	type WorksClaims struct {
		UserID string
		jwt.StandardClaims
	}
	log.WithFields(logrus.Fields{
		"serverVersion": Version,
	}).Infof("Starting application")
	log.Debugf("%v", tknStr)
	tokenString := tknStr[5:]
	log.Info(tokenString)
	at(time.Unix(0, 0), func() {
		token, err := jwt.ParseWithClaims(tokenString, &WorksClaims{}, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
		if err != nil {
			log.Errorf("The token value is invalid.\n")
			log.Errorf("%v\n", err)
			result["status"] = http.StatusUnauthorized
			result["message"] = "The token value is invalid."
			result["errorCode"] = "9998"
			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*WorksClaims); ok && token.Valid {
			log.WithFields(logrus.Fields{
				"claims.UserID": claims.UserID,
			}).Infof("Token check OK!!!\n")
			c.Set("cookie-user-id", claims.UserID)
			c.Next()
		} else {
			result["status"] = http.StatusUnauthorized
			result["message"] = "token is expired."
			result["errorCode"] = "9997"
			c.JSON(http.StatusOK, gin.H{
				"result": result,
			})
			c.Abort()
			return
		}
	})
	//c.Set("cookie-user-id","administrator")
}

func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}
