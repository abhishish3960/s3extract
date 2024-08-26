package controllers

import (
	"backend/models"
	"backend/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ContentController struct {
	s3Service *services.S3Service
}

func NewContentController(s3Service *services.S3Service) *ContentController {
	return &ContentController{s3Service: s3Service}
}

func (ctrl *ContentController) GetContents(c *gin.Context) {
	imagePrefix := "images/docx/"
	textPrefix := "text/docx/"

	imageKeys, err := ctrl.s3Service.ListFiles(imagePrefix)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list images"})
		return
	}

	textKeys, err := ctrl.s3Service.ListFiles(textPrefix)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list texts"})
		return
	}

	var images []models.ContentImageItem
	for _, key := range imageKeys {
		name := strings.TrimPrefix(key, imagePrefix)
		url, err := ctrl.s3Service.GeneratePresignedURL(key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate image URL"})
			return
		}
		images = append(images, models.ContentImageItem{Name: name, URL: url})
	}

	var texts []models.ContentTextItem
	for _, key := range textKeys {
		name := strings.TrimPrefix(key, textPrefix)

		// Read the content of the text file from S3
		content, err := ctrl.s3Service.GetFileContent(key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read text file content"})
			return
		}
		texts = append(texts, models.ContentTextItem{Name: name, Content: content})
	}

	c.JSON(http.StatusOK, gin.H{
		"images": images,
		"texts":  texts,
	})
}
