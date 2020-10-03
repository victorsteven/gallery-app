package handler

import (
	"fmt"
	"gallery/server/domain"
	"gallery/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type service struct {
	imgDomain domain.ImageService
}

func NewHandlerService(imgDomain domain.ImageService) *service {
	return &service{imgDomain: imgDomain}
}

func (s *service) ListImages(c *gin.Context) {

	images := utils.GetImages()

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"images":   images,
	})
}

func (s *service) SaveImageInfo(c *gin.Context) {

	var request *domain.ImageInfo

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"body":   "please provide valid inputs",
		})
		return
	}

	if err := utils.ValidateInputs(*request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"body":   err.Error(),
		})
		return
	}

	info, err := s.imgDomain.SaveImageInfo(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"body":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"body": info,
	})
}

//This function enable us to check the tag status
func (s *service) GetImageInfo(c *gin.Context) {

	requestId := c.Param("imageId")
	imageId, err := strconv.ParseUint(requestId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"body":   "Invalid image id provided",
		})
	}

	img, _ := s.imgDomain.GetImageInfo(imageId)

	if img != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"body": true,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"body": false,
	})
}

func (s *service) DeleteImageInfo(c *gin.Context) {

	requestId := c.Param("imageId")

	fmt.Println("the requested id: ", requestId)

	imageId, err := strconv.ParseUint(requestId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"body":   "Invalid id provided",
		})
		return
	}

	err = s.imgDomain.DeleteImageInfo(imageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"body":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"body": "success",
	})
}


