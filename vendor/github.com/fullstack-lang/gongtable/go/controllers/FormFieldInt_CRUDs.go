// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongtable/go/models"
	"github.com/fullstack-lang/gongtable/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __FormFieldInt__dummysDeclaration__ models.FormFieldInt
var __FormFieldInt_time__dummyDeclaration time.Duration

var mutexFormFieldInt sync.Mutex

// An FormFieldIntID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFormFieldInt updateFormFieldInt deleteFormFieldInt
type FormFieldIntID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FormFieldIntInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFormFieldInt updateFormFieldInt
type FormFieldIntInput struct {
	// The FormFieldInt to submit or modify
	// in: body
	FormFieldInt *orm.FormFieldIntAPI
}

// GetFormFieldInts
//
// swagger:route GET /formfieldints formfieldints getFormFieldInts
//
// # Get all formfieldints
//
// Responses:
// default: genericError
//
//	200: formfieldintDBResponse
func (controller *Controller) GetFormFieldInts(c *gin.Context) {

	// source slice
	var formfieldintDBs []orm.FormFieldIntDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldInts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoFormFieldInt.GetDB()

	query := db.Find(&formfieldintDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formfieldintAPIs := make([]orm.FormFieldIntAPI, 0)

	// for each formfieldint, update fields from the database nullable fields
	for idx := range formfieldintDBs {
		formfieldintDB := &formfieldintDBs[idx]
		_ = formfieldintDB
		var formfieldintAPI orm.FormFieldIntAPI

		// insertion point for updating fields
		formfieldintAPI.ID = formfieldintDB.ID
		formfieldintDB.CopyBasicFieldsToFormFieldInt(&formfieldintAPI.FormFieldInt)
		formfieldintAPI.FormFieldIntPointersEnconding = formfieldintDB.FormFieldIntPointersEnconding
		formfieldintAPIs = append(formfieldintAPIs, formfieldintAPI)
	}

	c.JSON(http.StatusOK, formfieldintAPIs)
}

// PostFormFieldInt
//
// swagger:route POST /formfieldints formfieldints postFormFieldInt
//
// Creates a formfieldint
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFormFieldInt(c *gin.Context) {

	mutexFormFieldInt.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormFieldInts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoFormFieldInt.GetDB()

	// Validate input
	var input orm.FormFieldIntAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create formfieldint
	formfieldintDB := orm.FormFieldIntDB{}
	formfieldintDB.FormFieldIntPointersEnconding = input.FormFieldIntPointersEnconding
	formfieldintDB.CopyBasicFieldsFromFormFieldInt(&input.FormFieldInt)

	query := db.Create(&formfieldintDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFormFieldInt.CheckoutPhaseOneInstance(&formfieldintDB)
	formfieldint := backRepo.BackRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntPtr[formfieldintDB.ID]

	if formfieldint != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), formfieldint)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, formfieldintDB)

	mutexFormFieldInt.Unlock()
}

// GetFormFieldInt
//
// swagger:route GET /formfieldints/{ID} formfieldints getFormFieldInt
//
// Gets the details for a formfieldint.
//
// Responses:
// default: genericError
//
//	200: formfieldintDBResponse
func (controller *Controller) GetFormFieldInt(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormFieldInt", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoFormFieldInt.GetDB()

	// Get formfieldintDB in DB
	var formfieldintDB orm.FormFieldIntDB
	if err := db.First(&formfieldintDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var formfieldintAPI orm.FormFieldIntAPI
	formfieldintAPI.ID = formfieldintDB.ID
	formfieldintAPI.FormFieldIntPointersEnconding = formfieldintDB.FormFieldIntPointersEnconding
	formfieldintDB.CopyBasicFieldsToFormFieldInt(&formfieldintAPI.FormFieldInt)

	c.JSON(http.StatusOK, formfieldintAPI)
}

// UpdateFormFieldInt
//
// swagger:route PATCH /formfieldints/{ID} formfieldints updateFormFieldInt
//
// # Update a formfieldint
//
// Responses:
// default: genericError
//
//	200: formfieldintDBResponse
func (controller *Controller) UpdateFormFieldInt(c *gin.Context) {

	mutexFormFieldInt.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFormFieldInt", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoFormFieldInt.GetDB()

	// Validate input
	var input orm.FormFieldIntAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formfieldintDB orm.FormFieldIntDB

	// fetch the formfieldint
	query := db.First(&formfieldintDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formfieldintDB.CopyBasicFieldsFromFormFieldInt(&input.FormFieldInt)
	formfieldintDB.FormFieldIntPointersEnconding = input.FormFieldIntPointersEnconding

	query = db.Model(&formfieldintDB).Updates(formfieldintDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formfieldintNew := new(models.FormFieldInt)
	formfieldintDB.CopyBasicFieldsToFormFieldInt(formfieldintNew)

	// get stage instance from DB instance, and call callback function
	formfieldintOld := backRepo.BackRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntPtr[formfieldintDB.ID]
	if formfieldintOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), formfieldintOld, formfieldintNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formfieldintDB
	c.JSON(http.StatusOK, formfieldintDB)

	mutexFormFieldInt.Unlock()
}

// DeleteFormFieldInt
//
// swagger:route DELETE /formfieldints/{ID} formfieldints deleteFormFieldInt
//
// # Delete a formfieldint
//
// default: genericError
//
//	200: formfieldintDBResponse
func (controller *Controller) DeleteFormFieldInt(c *gin.Context) {

	mutexFormFieldInt.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormFieldInt", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoFormFieldInt.GetDB()

	// Get model if exist
	var formfieldintDB orm.FormFieldIntDB
	if err := db.First(&formfieldintDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&formfieldintDB)

	// get an instance (not staged) from DB instance, and call callback function
	formfieldintDeleted := new(models.FormFieldInt)
	formfieldintDB.CopyBasicFieldsToFormFieldInt(formfieldintDeleted)

	// get stage instance from DB instance, and call callback function
	formfieldintStaged := backRepo.BackRepoFormFieldInt.Map_FormFieldIntDBID_FormFieldIntPtr[formfieldintDB.ID]
	if formfieldintStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), formfieldintStaged, formfieldintDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexFormFieldInt.Unlock()
}
