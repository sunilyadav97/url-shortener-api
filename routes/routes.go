package routes

import (
	"context"
	"net/http"
	"time"

	"url-shortener-api/database"
	"url-shortener-api/models"
	"url-shortener-api/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateShortUrlRequest is the request payload for creating a short URL.
type CreateShortUrlRequest struct {
	URL string `json:"url" binding:"required"`
}

// UpdateShortUrlRequest is the request payload for updating an existing short URL.
type UpdateShortUrlRequest struct {
	URL string `json:"url" binding:"required"`
}

// SetupRoutes sets up the API endpoints.
func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.POST("/createShortUrl", createShortUrl)
		api.GET("/:shortUrl", redirectShortUrl)
		api.PUT("/updateShortUrl/:shortUrl", updateShortUrl)
	}
}

// createShortUrl handles the creation of a short URL.
func createShortUrl(c *gin.Context) {
	var req CreateShortUrlRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Check if URL already exists
	var existing models.URLMapping
	collection := database.MI.MongoCollection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"longUrl": req.URL}).Decode(&existing)
	if err == nil {
		c.JSON(http.StatusCreated, gin.H{"shortUrl": existing.ShortURL})
		return
	}

	// Generate a unique short URL hash
	shortURL := utils.GenerateShortURL(req.URL)

	mapping := models.URLMapping{
		ID:        primitive.NewObjectID(),
		LongURL:   req.URL,
		ShortURL:  shortURL,
		CreatedAt: time.Now(),
	}

	_, err = collection.InsertOne(ctx, mapping)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while saving to DB"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"shortUrl": shortURL})
}

// redirectShortUrl retrieves the original URL and redirects the user.
func redirectShortUrl(c *gin.Context) {
	shortUrlParam := c.Param("shortUrl")
	collection := database.MI.MongoCollection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var mapping models.URLMapping
	err := collection.FindOne(ctx, bson.M{"shortUrl": shortUrlParam}).Decode(&mapping)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}
	// Redirect with status 301 (Moved Permanently)
	c.Redirect(http.StatusMovedPermanently, mapping.LongURL)
}

// updateShortUrl allows updating the destination URL for an existing short URL.
func updateShortUrl(c *gin.Context) {
	shortUrlParam := c.Param("shortUrl")
	var req UpdateShortUrlRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	collection := database.MI.MongoCollection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": bson.M{"longUrl": req.URL}}
	res, err := collection.UpdateOne(ctx, bson.M{"shortUrl": shortUrlParam}, update)
	if err != nil || res.ModifiedCount == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Successful", "message": "Short URL updated successfully"})
}
