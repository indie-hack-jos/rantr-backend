package services

import (
	"context"
	"log"
	"rantr/config"
	"rantr/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostService struct {
	client *mongo.Client
}

type PaginatedResponse struct {
    Posts      []models.Post `json:"posts"`
    TotalPages  int    `json:"total_pages"`
    CurrentPage int    `json:"current_page"`
}

func NewPostService() *PostService {
	client := config.GetMongoClient()
	return &PostService{client: client}
}

func (s *PostService) CreatePost (post models.Post) (*mongo.InsertOneResult, error) {
	collection := s.client.Database(config.GetEnv("MONGO_DB_NAME")).Collection("posts")
	result, err := collection.InsertOne(context.TODO(), post)

	if err != nil {
		return nil, err
	}


	return result, nil
}

func (s *PostService) GetPosts (pageNumber int, pageSize int) (PaginatedResponse, error) {
	collection := s.client.Database(config.GetEnv("MONGO_DB_NAME")).Collection("posts")
    skip := (pageNumber - 1) * pageSize

    findOptions := options.Find()
    findOptions.SetSkip(int64(skip))
    findOptions.SetLimit(int64(pageSize))

	cursor, err := collection.Find(context.Background(), bson.M{}, findOptions)

    if err != nil {
		log.Fatal(err)
		return PaginatedResponse{}, err
    }

    defer cursor.Close(context.Background())

	var posts []models.Post
    for cursor.Next(context.Background()) {
        var post models.Post
        if err := cursor.Decode(&post); err != nil {
            return PaginatedResponse{}, err
        }
        posts = append(posts, post)
    }

	totalDocuments, err := collection.CountDocuments(context.Background(), bson.M{})
    if err != nil {
        return PaginatedResponse{}, err
    }

	totalPages := int(totalDocuments) / pageSize
    if int(totalDocuments)%pageSize != 0 {
        totalPages++
    }

	return PaginatedResponse{
        Posts:       posts,
        TotalPages:  totalPages,
        CurrentPage: int(pageNumber),
    }, nil
}

