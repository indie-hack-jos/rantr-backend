package controllers

import (
	"fmt"
	"net/http"
	"rantr/models"
	"rantr/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Username string `bson:"username" binding:"required"`
	Password string `bson:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		if err.Error() == "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag\nKey: 'User.Username' Error:Field validation for 'Username' failed on the 'required' tag" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username or Email is required"})
			return
		}

		if err.Error() == "Key: 'User.Username' Error:Field validation for 'Username' failed on the 'required' tag" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
			return
		}

		// if err.Error() == "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag" {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		// 	return
		// }

		if err.Error() == "Key: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = string(hashedPassword)
	userService := services.NewUserService()
	token, err := userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Created Successfully", "token": token, "data": map[string]interface{}{}})
}

func LoginUser(c *gin.Context) {
	var user Login
	if err := c.ShouldBindJSON(&user); err != nil {

		if err.Error() == "Key: 'User.Username' Error:Field validation for 'Username' failed on the 'required' tag" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
			return
		}

		if err.Error() == "Key: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := services.NewUserService()
	data, err := userService.FindOne(bson.M{"username": user.Username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// userPassword, ok := data["password"].(string)
	// if !ok {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Password does not exists"})
	// 	return
	// }

	// fmt.Println("USER ENCRYPTED PASSWORD", userPassword)
	password, ok := data["password"].(string)
	if !ok {
		fmt.Println("Password not found or not a string")
		return
	}

	// username, ok := data["username"].(string)
	// if !ok {
	// 	fmt.Println("Username not found or not a string")
	// 	return
	// }

	// id, ok := data["id"].(string)
	// if !ok {
	// 	fmt.Println("Id not found or not a string")
	// 	return
	// }

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.NewUserService().GenerateJwtToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// keys to filter out
	keysToFilter := []string{"password"}

	filteredData := make(bson.M)

	for key, value := range data {
		// Check if the key is in the list of keys to filter out
		filtered := false
		for _, filterKey := range keysToFilter {
			if key == filterKey {
				filtered = true
				break
			}
		}

		// If the key is not in the list of keys to filter out, add it to the filteredData map
		if !filtered {
			filteredData[key] = value
		}
	}
	// responseData := map[string]interface{}{
	// 	"id":       data["id"].(string),
	// 	"username": data["username"].(string),
	// }

	c.JSON(http.StatusOK, gin.H{"status": "Logged in Successfully", "token": token, "data": filteredData})
}

func AuthUser(c *gin.Context) {
	claim, exists := c.Get("user")
	if !exists {
		// Handle the case where "claims" does not exist in the context
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Retrieved Successfully", "data": map[string]interface{}{}})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Retrieved Successfully", "data": claim})
}

// Implement GetUser, UpdateUser, and DeleteUser similarly
