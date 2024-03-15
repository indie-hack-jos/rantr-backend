package services

import (
	"context"
	// "fmt"
	"rantr/config"
	"rantr/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	client *mongo.Client
}

func NewUserService() *UserService {
	client := config.DbConfig()
	return &UserService{client: client}
}

func (s *UserService) CreateUser(user models.User) (string, error) {
	jwtSecretKey := []byte(config.GetEnv("JWT_SECRET_KEY"))

	collection := s.client.Database(config.GetEnv("MONGO_DB_NAME")).Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtSecretKey)
	// if err != nil {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
	//     return
	// }
	return tokenString, err
}

func (s *UserService) GenerateJwtToken(username string) (string, error) {
	jwtSecretKey := []byte(config.GetEnv("JWT_SECRET_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtSecretKey)

	return tokenString, err
}
func (s *UserService) FindOne(filter bson.M) (bson.M, error) {
	var result bson.M

	collection := s.client.Database(config.GetEnv("MONGO_DB_NAME")).Collection("users")
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		// if err == mongo.ErrNoDocuments {
		// 	// No document was found
		// 	return nil, "No document found"
		// }
		// An error occurred
		return nil, err
	}

	return result, nil
}

// Implement GetUser, UpdateUser, and DeleteUser similarly
