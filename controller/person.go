package controller

import (
	"mini_project_restapi/database"
	"mini_project_restapi/repository"
	"mini_project_restapi/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPerson(ctx *gin.Context) {
	var (
		result     gin.H
		statusCode int
	)

	persons, err := repository.GetAllPerson(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
		statusCode = http.StatusBadRequest
	} else {
		result = gin.H{
			"result": persons,
		}
		statusCode = http.StatusOK
	}

	ctx.JSON(statusCode, result)
}

func InsertPerson(ctx *gin.Context) {
	var (
		person structs.Person
	)

	err := ctx.BindJSON(&person)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": err,
		})
		return
	}
	err = repository.InsertPerson(database.DbConnection, person)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Person",
	})
	return
}

func UpdatePerson(ctx *gin.Context) {
	var person structs.Person
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := ctx.BindJSON(&person)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": err,
		})
		return
	}

	person.ID = int64(id)

	err = repository.UpdatePerson(database.DbConnection, person)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "Success Update Person",
	})
	return
}

func DeletePerson(ctx *gin.Context) {
	var person structs.Person
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": err,
		})
	}

	person.ID = int64(id)

	err = repository.DeletePerson(database.DbConnection, person)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Person",
	})
	return
}
