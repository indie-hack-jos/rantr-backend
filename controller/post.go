package controllers

import (
	"net/http"
	"rantr/dto"
	"rantr/models"
	"rantr/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/twitchyliquid64/golang-asm/obj"
)


func CreatePost (c *gin.Context) {
	var post models.Post
	var input dto.CreatePostInput
	claim, exists := c.Get("user")

	if !exists {
		// Handle the case where "claims" does not exist in the context
		c.JSON(http.StatusOK, gin.H{"status": http.NotFound, "message": "Cannot find user"})
	}

    if err := c.ShouldBindJSON(&input); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	post.User = claim.(primitive.M)["_id"].(primitive.ObjectID)

	post.Content = input.Content

	postService := services.NewPostService()
	data, err := postService.CreatePost(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Post created successfully", "data": data})
}

func GetPosts(c *gin.Context) {
	pageStr := c.Query("page_number")
    pageSizeStr := c.Query("page_size")


	pageNumber, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid page number"})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid page size"})
        return
    }

	postService := services.NewPostService()
	data, err := postService.GetPosts(pageNumber, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Retrieved Posts Successfully", "data": data})

}