package controller

import (
	"c2fit-hw-backend/databases"
	"c2fit-hw-backend/models"
	"c2fit-hw-backend/session"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(c *gin.Context) {
	var user models.User
	users := databases.MyDB.Database.Collection("user")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := user.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var result bson.M
	err := users.FindOne(context.TODO(), bson.M{"Email": user.Email}).Decode(&result)
	// fmt.Println(result, result == nil)
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Your Email already exists"})
		return
	}
	// tmp := sha256.Sum256([]byte(user.Password))[:]
	// var tmp []byte
	h := sha256.New()
	h.Write([]byte(user.Password))
	user.HashPasswd = base64.URLEncoding.EncodeToString(h.Sum(nil))

	_, err = users.InsertOne(context.TODO(), user)
	if err != nil {
		log.Panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User created successfully",
		"data": gin.H{
			"Email": user.Email,
		},
	})
}

func UserLogin(c *gin.Context) {
	var userLogin models.UserLogIn

	users := databases.MyDB.Database.Collection("user")

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err := users.FindOne(context.TODO(), bson.M{"Email": userLogin.Email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			c.JSON(http.StatusBadRequest, gin.H{"error": "Your Email or Password is incorrect"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h := sha256.New()
	h.Write([]byte(userLogin.Password))
	if user.HashPasswd != base64.URLEncoding.EncodeToString(h.Sum(nil)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Your Email or Password is incorrect"})
		return
	}

	if err := session.LoginSession(c, &user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, user)
}

func WhoAmI(c *gin.Context) {
	if !session.HasLoggedIn(c) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not logged in"})
	}
	userID := session.GetUserId(c)
	c.JSON(http.StatusOK, gin.H{"user": userID})
}

func UserLogOut(c *gin.Context) {
	session.LogoutSession(c)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
