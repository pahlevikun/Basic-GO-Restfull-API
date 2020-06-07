package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"../structs"
)

// GetPerson : Get data by {{id}}
func (idb *InDB) GetPerson(context *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := context.Param("id")
	errors := idb.DB.Where("id = ?", id).First(&person).Error
	if errors != nil {
		result = gin.H{
			"result": errors.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}
	context.JSON(http.StatusOK, result)
}

// GetAllPerson : Get all person data
func (idb *InDB) GetAllPerson(context *gin.Context) {
	var (
		persons []structs.Person
		result  gin.H
	)

	idb.DB.Find(&persons)
	if len(persons) < 1 {
		result = gin.H{
			"result":  nil,
			"count":   0,
			"message": "succeed get all data",
		}
	} else {
		result = gin.H{
			"result":  persons,
			"count":   len(persons),
			"message": "succeed get all data",
		}
	}
	context.JSON(http.StatusOK, result)
}

// CreatePerson : Create single person data
func (idb *InDB) CreatePerson(context *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	firstName := context.PostForm("first_name")
	lastName := context.PostForm("last_name")
	person.FirstName = firstName
	person.LastName = lastName

	idb.DB.Create(&person)

	result = gin.H{
		"result":  person,
		"message": "succeed create data",
	}
	context.JSON(http.StatusOK, result)
}

// UpdatePerson : Update single person data with {{id}} as query
func (idb *InDB) UpdatePerson(context *gin.Context) {
	var (
		person        structs.Person
		updatedPerson structs.Person
		result        gin.H
	)

	id := context.Param("id")
	firstName := context.PostForm("first_name")
	lastName := context.PostForm("last_name")

	errors := idb.DB.First(&person, id).Error
	if errors != nil {
		result = gin.H{
			"result":  nil,
			"message": "data not found",
		}
	}

	updatedPerson.FirstName = firstName
	updatedPerson.LastName = lastName

	errors = idb.DB.Model(&person).Update(updatedPerson).Error

	if errors != nil {
		result = gin.H{
			"result":  nil,
			"message": "Failed updating data",
		}
	} else {
		result = gin.H{
			"result":  updatedPerson,
			"message": "Succeed updating data",
		}
	}

	context.JSON(http.StatusOK, result)
}

// DeletePerson : Delete single person data with {{id}} as query
func (idb *InDB) DeletePerson(context *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := context.Query("id")

	errors := idb.DB.First(&person, id).Error
	if errors != nil {
		result = gin.H{
			"result":  false,
			"message": "data not found",
		}
	}

	errors = idb.DB.Delete(&person).Error

	if errors != nil {
		result = gin.H{
			"result":  false,
			"message": "Failed delete data",
		}
	} else {
		result = gin.H{
			"result":  true,
			"message": "Succeed delete data",
		}
	}

	context.JSON(http.StatusOK, result)
}
