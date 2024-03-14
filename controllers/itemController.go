package controllers

import (
	"net/http"
	"strconv"
	"tft-team-info/database"
	"tft-team-info/repository"
	"tft-team-info/structs"

	"github.com/gin-gonic/gin"
)

func GetAllItem(c *gin.Context) {
	var (
		result gin.H
	)

	items, err := repository.GetAllItem(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"Item List": items,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertItem(c *gin.Context) {
	var item structs.Item

	err := c.ShouldBindJSON(&item)
	if err != nil {
		panic(err)
	}

	err = repository.InsertItem(database.DbConnection, item)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success add item",
	})
}

func UpdateItem(c *gin.Context) {
	var item structs.Item
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed convert id"})
		return
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.ID = int64(id)

	err = repository.UpdateItem(database.DbConnection, item)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success update item",
	})
}

func DeleteItem(c *gin.Context) {
	var item structs.Item
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed convert id"})
		return
	}

	item.ID = int64(id)

	err = repository.DeleteItem(database.DbConnection, item)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "succes delete item",
	})
}
