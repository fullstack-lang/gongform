// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongsvg/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __RectLinkLink__dummysDeclaration__ models.RectLinkLink
var __RectLinkLink_time__dummyDeclaration time.Duration

var mutexRectLinkLink sync.Mutex

// An RectLinkLinkID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRectLinkLink updateRectLinkLink deleteRectLinkLink
type RectLinkLinkID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RectLinkLinkInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRectLinkLink updateRectLinkLink
type RectLinkLinkInput struct {
	// The RectLinkLink to submit or modify
	// in: body
	RectLinkLink *orm.RectLinkLinkAPI
}

// GetRectLinkLinks
//
// swagger:route GET /rectlinklinks rectlinklinks getRectLinkLinks
//
// # Get all rectlinklinks
//
// Responses:
// default: genericError
//
//	200: rectlinklinkDBResponse
func (controller *Controller) GetRectLinkLinks(c *gin.Context) {

	// source slice
	var rectlinklinkDBs []orm.RectLinkLinkDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRectLinkLinks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRectLinkLink.GetDB()

	query := db.Find(&rectlinklinkDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	rectlinklinkAPIs := make([]orm.RectLinkLinkAPI, 0)

	// for each rectlinklink, update fields from the database nullable fields
	for idx := range rectlinklinkDBs {
		rectlinklinkDB := &rectlinklinkDBs[idx]
		_ = rectlinklinkDB
		var rectlinklinkAPI orm.RectLinkLinkAPI

		// insertion point for updating fields
		rectlinklinkAPI.ID = rectlinklinkDB.ID
		rectlinklinkDB.CopyBasicFieldsToRectLinkLink_WOP(&rectlinklinkAPI.RectLinkLink_WOP)
		rectlinklinkAPI.RectLinkLinkPointersEncoding = rectlinklinkDB.RectLinkLinkPointersEncoding
		rectlinklinkAPIs = append(rectlinklinkAPIs, rectlinklinkAPI)
	}

	c.JSON(http.StatusOK, rectlinklinkAPIs)
}

// PostRectLinkLink
//
// swagger:route POST /rectlinklinks rectlinklinks postRectLinkLink
//
// Creates a rectlinklink
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRectLinkLink(c *gin.Context) {

	mutexRectLinkLink.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRectLinkLinks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRectLinkLink.GetDB()

	// Validate input
	var input orm.RectLinkLinkAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create rectlinklink
	rectlinklinkDB := orm.RectLinkLinkDB{}
	rectlinklinkDB.RectLinkLinkPointersEncoding = input.RectLinkLinkPointersEncoding
	rectlinklinkDB.CopyBasicFieldsFromRectLinkLink_WOP(&input.RectLinkLink_WOP)

	query := db.Create(&rectlinklinkDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRectLinkLink.CheckoutPhaseOneInstance(&rectlinklinkDB)
	rectlinklink := backRepo.BackRepoRectLinkLink.Map_RectLinkLinkDBID_RectLinkLinkPtr[rectlinklinkDB.ID]

	if rectlinklink != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), rectlinklink)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, rectlinklinkDB)

	mutexRectLinkLink.Unlock()
}

// GetRectLinkLink
//
// swagger:route GET /rectlinklinks/{ID} rectlinklinks getRectLinkLink
//
// Gets the details for a rectlinklink.
//
// Responses:
// default: genericError
//
//	200: rectlinklinkDBResponse
func (controller *Controller) GetRectLinkLink(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRectLinkLink", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRectLinkLink.GetDB()

	// Get rectlinklinkDB in DB
	var rectlinklinkDB orm.RectLinkLinkDB
	if err := db.First(&rectlinklinkDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var rectlinklinkAPI orm.RectLinkLinkAPI
	rectlinklinkAPI.ID = rectlinklinkDB.ID
	rectlinklinkAPI.RectLinkLinkPointersEncoding = rectlinklinkDB.RectLinkLinkPointersEncoding
	rectlinklinkDB.CopyBasicFieldsToRectLinkLink_WOP(&rectlinklinkAPI.RectLinkLink_WOP)

	c.JSON(http.StatusOK, rectlinklinkAPI)
}

// UpdateRectLinkLink
//
// swagger:route PATCH /rectlinklinks/{ID} rectlinklinks updateRectLinkLink
//
// # Update a rectlinklink
//
// Responses:
// default: genericError
//
//	200: rectlinklinkDBResponse
func (controller *Controller) UpdateRectLinkLink(c *gin.Context) {

	mutexRectLinkLink.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRectLinkLink", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRectLinkLink.GetDB()

	// Validate input
	var input orm.RectLinkLinkAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rectlinklinkDB orm.RectLinkLinkDB

	// fetch the rectlinklink
	query := db.First(&rectlinklinkDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rectlinklinkDB.CopyBasicFieldsFromRectLinkLink_WOP(&input.RectLinkLink_WOP)
	rectlinklinkDB.RectLinkLinkPointersEncoding = input.RectLinkLinkPointersEncoding

	query = db.Model(&rectlinklinkDB).Updates(rectlinklinkDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rectlinklinkNew := new(models.RectLinkLink)
	rectlinklinkDB.CopyBasicFieldsToRectLinkLink(rectlinklinkNew)

	// redeem pointers
	rectlinklinkDB.DecodePointers(backRepo, rectlinklinkNew)

	// get stage instance from DB instance, and call callback function
	rectlinklinkOld := backRepo.BackRepoRectLinkLink.Map_RectLinkLinkDBID_RectLinkLinkPtr[rectlinklinkDB.ID]
	if rectlinklinkOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), rectlinklinkOld, rectlinklinkNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rectlinklinkDB
	c.JSON(http.StatusOK, rectlinklinkDB)

	mutexRectLinkLink.Unlock()
}

// DeleteRectLinkLink
//
// swagger:route DELETE /rectlinklinks/{ID} rectlinklinks deleteRectLinkLink
//
// # Delete a rectlinklink
//
// default: genericError
//
//	200: rectlinklinkDBResponse
func (controller *Controller) DeleteRectLinkLink(c *gin.Context) {

	mutexRectLinkLink.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRectLinkLink", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongsvg/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRectLinkLink.GetDB()

	// Get model if exist
	var rectlinklinkDB orm.RectLinkLinkDB
	if err := db.First(&rectlinklinkDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&rectlinklinkDB)

	// get an instance (not staged) from DB instance, and call callback function
	rectlinklinkDeleted := new(models.RectLinkLink)
	rectlinklinkDB.CopyBasicFieldsToRectLinkLink(rectlinklinkDeleted)

	// get stage instance from DB instance, and call callback function
	rectlinklinkStaged := backRepo.BackRepoRectLinkLink.Map_RectLinkLinkDBID_RectLinkLinkPtr[rectlinklinkDB.ID]
	if rectlinklinkStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), rectlinklinkStaged, rectlinklinkDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexRectLinkLink.Unlock()
}
