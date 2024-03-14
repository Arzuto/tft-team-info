package controllers

import (
	"net/http"
	"strconv"
	"tft-team-info/database"
	"tft-team-info/repository"
	"tft-team-info/structs"

	"github.com/gin-gonic/gin"
)

func GetAllOrigin(c *gin.Context) {
	var (
		result gin.H
	)

	origins, err := repository.GetAllOrigin(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"Origin Trait": origins,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertOrigin(c *gin.Context) {
	var origin structs.Origin

	err := c.ShouldBindJSON(&origin)
	if err != nil {
		panic(err)
	}

	err = repository.InsertOrigin(database.DbConnection, origin)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success add origin",
	})
}

func UpdateOrigin(c *gin.Context) {
	var origin structs.Origin
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed convert id"})
		return
	}

	if err := c.ShouldBindJSON(&origin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	origin.ID = int64(id)

	err = repository.UpdateOrigin(database.DbConnection, origin)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success update origin",
	})
}

func DeleteOrigin(c *gin.Context) {
	var origin structs.Origin
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed convert id"})
		return
	}

	origin.ID = int64(id)

	err = repository.DeleteOrigin(database.DbConnection, origin)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "succes delete origin",
	})
}
