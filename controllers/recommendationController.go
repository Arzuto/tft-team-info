package controllers

import (
	"net/http"
	"strconv"
	"tft-team-info/database"
	"tft-team-info/repository"
	"tft-team-info/structs"

	"github.com/gin-gonic/gin"
)

func GetAllRecommendation(c *gin.Context) {
	recommendations, err := repository.GetAllRecommendation(database.DbConnection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"recommendations": recommendations})
}

func InsertRecommendation(c *gin.Context) {
	var recommendation structs.Recommendation

	err := c.ShouldBindJSON(&recommendation)
	if err != nil {
		panic(err)
	}

	err = repository.InsertRecommendation(database.DbConnection, recommendation)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success add recommendation",
	})
}

func UpdateRecommendation(c *gin.Context) {
	var recommendation structs.Recommendation
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed convert id"})
		return
	}

	if err := c.ShouldBindJSON(&recommendation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recommendation.ID = int64(id)

	err = repository.UpdateRecommendation(database.DbConnection, recommendation)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success update recommendation",
	})
}

func DeleteRecommendation(c *gin.Context) {
	var recommendation structs.Recommendation
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed convert id"})
		return
	}

	recommendation.ID = int64(id)

	err = repository.DeleteRecommendation(database.DbConnection, recommendation)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "succes delete recommendation",
	})
}
