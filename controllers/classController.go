package controllers

import (
	"net/http"
	"strconv"
	"tft-team-info/database"
	"tft-team-info/repository"
	"tft-team-info/structs"

	"github.com/gin-gonic/gin"
)

func GetAllClass(c *gin.Context) {
	var (
		result gin.H
	)

	classs, err := repository.GetAllClass(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"Class Trait List": classs,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertClass(c *gin.Context) {
	var class structs.Class

	err := c.ShouldBindJSON(&class)
	if err != nil {
		panic(err)
	}

	err = repository.InsertClass(database.DbConnection, class)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success add class",
	})
}

func UpdateClass(c *gin.Context) {
	var class structs.Class
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed convert id"})
		return
	}

	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	class.ID = int64(id)

	err = repository.UpdateClass(database.DbConnection, class)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success update class",
	})
}

func DeleteClass(c *gin.Context) {
	var class structs.Class
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed convert id"})
		return
	}

	class.ID = int64(id)

	err = repository.DeleteClass(database.DbConnection, class)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "succes delete class",
	})
}
