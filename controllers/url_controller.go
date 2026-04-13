package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-short/models"
	"github.com/go-short/services"
	"github.com/go-short/utils"
)

type Request struct {
	URL string `json:"url" binding:"required,url"`
}

func CreateShortURL(c *gin.Context) {
	var req Request

	if err := c.ShouldBindJSON(&req); err != nil || req.URL == "" {
		requestError := utils.FormatValidationError(err)
		var message strings.Builder
		for _, v := range requestError {
			fmt.Fprint(&message, v.Msg+". ")
		}
		c.JSON(http.StatusBadRequest, models.CustomResponseModel{
			Status:  "error",
			Message: message.String(),
		})
		return
	}

	url, err := services.CreateURL(req.URL)
	if err != nil {
		code, response := utils.FormatDBError(err)
		c.JSON(code, models.CustomResponseModel{
			Status:  "error",
			Message: response["msg"],
		})
		return
	}

	redirectUrl := os.Getenv("RETURN_URL")

	c.JSON(http.StatusCreated, models.CustomResponseModel{
		Status: "success",
		Data: map[string]any{
			"code":         url.ShortCode,
			"original_url": req.URL,
			"short_url":    redirectUrl + url.ShortCode,
			"created_at":   url.CreatedAt,
			"updated_at":   url.UpdatedAt,
		},
		Message: "Short URL created successfully",
	})
}

func RedirectURL(c *gin.Context) {
	code := c.Param("code")

	url, err := services.GetURL(code)
	if err != nil {
		c.JSON(http.StatusNotFound, models.CustomResponseModel{
			Status:  "error",
			Message: "Data not found",
		})
		return
	}

	services.IncrementClicks(code)
	c.Redirect(http.StatusMovedPermanently, url.URL)
}

func GetURL(c *gin.Context) {
	code := c.Param("code")

	url, err := services.GetURL(code)
	if err != nil {
		code, response := utils.FormatDBError(err)
		c.JSON(code, models.CustomResponseModel{
			Status:  "error",
			Message: response["msg"],
		})
		return
	}

	c.JSON(http.StatusOK, models.CustomResponseModel{
		Status:  "success",
		Message: "Short URL retrieved successfully",
		Data:    url,
	})
}

func UpdateURL(c *gin.Context) {
	code := c.Param("code")

	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload paylad must be not empty"})
		return
	}

	err := services.UpdateURL(code, req.URL)
	if err != nil {
		code, response := utils.FormatDBError(err)
		c.JSON(code, models.CustomResponseModel{
			Status:  "error",
			Message: response["msg"],
		})
		return
	}

	c.JSON(http.StatusOK, models.CustomResponseModel{
		Status:  "success",
		Message: "Short URL updated successfully",
	})
}

func DeleteURL(c *gin.Context) {
	code := c.Param("code")

	err := services.DeleteURL(code)
	if err != nil {
		code, response := utils.FormatDBError(err)
		c.JSON(code, models.CustomResponseModel{
			Status:  "error",
			Message: response["msg"],
		})
		return
	}

	c.JSON(http.StatusOK, models.CustomResponseModel{
		Status:  "success",
		Message: "Short URL deleted successfully",
	})
}

func GetStats(c *gin.Context) {
	code := c.Param("code")

	url, err := services.GetURL(code)
	if err != nil {
		code, response := utils.FormatDBError(err)
		c.JSON(code, models.CustomResponseModel{
			Status:  "error",
			Message: response["msg"],
		})
		return
	}

	c.JSON(http.StatusOK, models.CustomResponseModel{
		Data:    url,
		Message: "Stats retrieved successfully",
		Status:  "success",
	})
}
