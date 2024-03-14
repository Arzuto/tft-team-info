package controllers

import (
	"net/http"
	"strconv"
	"tft-team-info/database"
	"tft-team-info/repository"
	"tft-team-info/structs"

	"github.com/gin-gonic/gin"
)

func GetAllCharacter(c *gin.Context) {
	var (
		result gin.H
	)

	characters, err := repository.GetAllCharacter(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"Character Trait List": characters,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCharacter(c *gin.Context) {
	var character structs.Character

	err := c.ShouldBindJSON(&character)
	if err != nil {
		panic(err)
	}

	err = repository.InsertCharacter(database.DbConnection, character)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success add character",
	})
}

func UpdateCharacter(c *gin.Context) {
	var character structs.Character
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed convert id"})
		return
	}

	if err := c.ShouldBindJSON(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	character.ID = int64(id)

	err = repository.UpdateCharacter(database.DbConnection, character)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success update character",
	})
}

func DeleteCharacter(c *gin.Context) {
	var character structs.Character
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed convert id"})
		return
	}

	character.ID = int64(id)

	err = repository.DeleteCharacter(database.DbConnection, character)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "succes delete character",
	})
}
