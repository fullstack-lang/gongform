// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __NoteShapeLink__dummysDeclaration__ models.NoteShapeLink
var __NoteShapeLink_time__dummyDeclaration time.Duration

var mutexNoteShapeLink sync.Mutex

// An NoteShapeLinkID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getNoteShapeLink updateNoteShapeLink deleteNoteShapeLink
type NoteShapeLinkID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// NoteShapeLinkInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postNoteShapeLink updateNoteShapeLink
type NoteShapeLinkInput struct {
	// The NoteShapeLink to submit or modify
	// in: body
	NoteShapeLink *orm.NoteShapeLinkAPI
}

// GetNoteShapeLinks
//
// swagger:route GET /noteshapelinks noteshapelinks getNoteShapeLinks
//
// # Get all noteshapelinks
//
// Responses:
// default: genericError
//
//	200: noteshapelinkDBResponse
func (controller *Controller) GetNoteShapeLinks(c *gin.Context) {

	// source slice
	var noteshapelinkDBs []orm.NoteShapeLinkDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNoteShapeLinks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoNoteShapeLink.GetDB()

	query := db.Find(&noteshapelinkDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	noteshapelinkAPIs := make([]orm.NoteShapeLinkAPI, 0)

	// for each noteshapelink, update fields from the database nullable fields
	for idx := range noteshapelinkDBs {
		noteshapelinkDB := &noteshapelinkDBs[idx]
		_ = noteshapelinkDB
		var noteshapelinkAPI orm.NoteShapeLinkAPI

		// insertion point for updating fields
		noteshapelinkAPI.ID = noteshapelinkDB.ID
		noteshapelinkDB.CopyBasicFieldsToNoteShapeLink(&noteshapelinkAPI.NoteShapeLink)
		noteshapelinkAPI.NoteShapeLinkPointersEnconding = noteshapelinkDB.NoteShapeLinkPointersEnconding
		noteshapelinkAPIs = append(noteshapelinkAPIs, noteshapelinkAPI)
	}

	c.JSON(http.StatusOK, noteshapelinkAPIs)
}

// PostNoteShapeLink
//
// swagger:route POST /noteshapelinks noteshapelinks postNoteShapeLink
//
// Creates a noteshapelink
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostNoteShapeLink(c *gin.Context) {

	mutexNoteShapeLink.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostNoteShapeLinks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoNoteShapeLink.GetDB()

	// Validate input
	var input orm.NoteShapeLinkAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create noteshapelink
	noteshapelinkDB := orm.NoteShapeLinkDB{}
	noteshapelinkDB.NoteShapeLinkPointersEnconding = input.NoteShapeLinkPointersEnconding
	noteshapelinkDB.CopyBasicFieldsFromNoteShapeLink(&input.NoteShapeLink)

	query := db.Create(&noteshapelinkDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoNoteShapeLink.CheckoutPhaseOneInstance(&noteshapelinkDB)
	noteshapelink := backRepo.BackRepoNoteShapeLink.Map_NoteShapeLinkDBID_NoteShapeLinkPtr[noteshapelinkDB.ID]

	if noteshapelink != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), noteshapelink)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, noteshapelinkDB)

	mutexNoteShapeLink.Unlock()
}

// GetNoteShapeLink
//
// swagger:route GET /noteshapelinks/{ID} noteshapelinks getNoteShapeLink
//
// Gets the details for a noteshapelink.
//
// Responses:
// default: genericError
//
//	200: noteshapelinkDBResponse
func (controller *Controller) GetNoteShapeLink(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNoteShapeLink", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoNoteShapeLink.GetDB()

	// Get noteshapelinkDB in DB
	var noteshapelinkDB orm.NoteShapeLinkDB
	if err := db.First(&noteshapelinkDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var noteshapelinkAPI orm.NoteShapeLinkAPI
	noteshapelinkAPI.ID = noteshapelinkDB.ID
	noteshapelinkAPI.NoteShapeLinkPointersEnconding = noteshapelinkDB.NoteShapeLinkPointersEnconding
	noteshapelinkDB.CopyBasicFieldsToNoteShapeLink(&noteshapelinkAPI.NoteShapeLink)

	c.JSON(http.StatusOK, noteshapelinkAPI)
}

// UpdateNoteShapeLink
//
// swagger:route PATCH /noteshapelinks/{ID} noteshapelinks updateNoteShapeLink
//
// # Update a noteshapelink
//
// Responses:
// default: genericError
//
//	200: noteshapelinkDBResponse
func (controller *Controller) UpdateNoteShapeLink(c *gin.Context) {

	mutexNoteShapeLink.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateNoteShapeLink", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoNoteShapeLink.GetDB()

	// Validate input
	var input orm.NoteShapeLinkAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var noteshapelinkDB orm.NoteShapeLinkDB

	// fetch the noteshapelink
	query := db.First(&noteshapelinkDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	noteshapelinkDB.CopyBasicFieldsFromNoteShapeLink(&input.NoteShapeLink)
	noteshapelinkDB.NoteShapeLinkPointersEnconding = input.NoteShapeLinkPointersEnconding

	query = db.Model(&noteshapelinkDB).Updates(noteshapelinkDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	noteshapelinkNew := new(models.NoteShapeLink)
	noteshapelinkDB.CopyBasicFieldsToNoteShapeLink(noteshapelinkNew)

	// get stage instance from DB instance, and call callback function
	noteshapelinkOld := backRepo.BackRepoNoteShapeLink.Map_NoteShapeLinkDBID_NoteShapeLinkPtr[noteshapelinkDB.ID]
	if noteshapelinkOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), noteshapelinkOld, noteshapelinkNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the noteshapelinkDB
	c.JSON(http.StatusOK, noteshapelinkDB)

	mutexNoteShapeLink.Unlock()
}

// DeleteNoteShapeLink
//
// swagger:route DELETE /noteshapelinks/{ID} noteshapelinks deleteNoteShapeLink
//
// # Delete a noteshapelink
//
// default: genericError
//
//	200: noteshapelinkDBResponse
func (controller *Controller) DeleteNoteShapeLink(c *gin.Context) {

	mutexNoteShapeLink.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteNoteShapeLink", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoNoteShapeLink.GetDB()

	// Get model if exist
	var noteshapelinkDB orm.NoteShapeLinkDB
	if err := db.First(&noteshapelinkDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&noteshapelinkDB)

	// get an instance (not staged) from DB instance, and call callback function
	noteshapelinkDeleted := new(models.NoteShapeLink)
	noteshapelinkDB.CopyBasicFieldsToNoteShapeLink(noteshapelinkDeleted)

	// get stage instance from DB instance, and call callback function
	noteshapelinkStaged := backRepo.BackRepoNoteShapeLink.Map_NoteShapeLinkDBID_NoteShapeLinkPtr[noteshapelinkDB.ID]
	if noteshapelinkStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), noteshapelinkStaged, noteshapelinkDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexNoteShapeLink.Unlock()
}
