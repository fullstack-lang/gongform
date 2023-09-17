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
var __Rect__dummysDeclaration__ models.Rect
var __Rect_time__dummyDeclaration time.Duration

var mutexRect sync.Mutex

// An RectID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRect updateRect deleteRect
type RectID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RectInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRect updateRect
type RectInput struct {
	// The Rect to submit or modify
	// in: body
	Rect *orm.RectAPI
}

// GetRects
//
// swagger:route GET /rects rects getRects
//
// # Get all rects
//
// Responses:
// default: genericError
//
//	200: rectDBResponse
func (controller *Controller) GetRects(c *gin.Context) {

	// source slice
	var rectDBs []orm.RectDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRects", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoRect.GetDB()

	query := db.Find(&rectDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	rectAPIs := make([]orm.RectAPI, 0)

	// for each rect, update fields from the database nullable fields
	for idx := range rectDBs {
		rectDB := &rectDBs[idx]
		_ = rectDB
		var rectAPI orm.RectAPI

		// insertion point for updating fields
		rectAPI.ID = rectDB.ID
		rectDB.CopyBasicFieldsToRect(&rectAPI.Rect)
		rectAPI.RectPointersEnconding = rectDB.RectPointersEnconding
		rectAPIs = append(rectAPIs, rectAPI)
	}

	c.JSON(http.StatusOK, rectAPIs)
}

// PostRect
//
// swagger:route POST /rects rects postRect
//
// Creates a rect
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRect(c *gin.Context) {

	mutexRect.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRects", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoRect.GetDB()

	// Validate input
	var input orm.RectAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create rect
	rectDB := orm.RectDB{}
	rectDB.RectPointersEnconding = input.RectPointersEnconding
	rectDB.CopyBasicFieldsFromRect(&input.Rect)

	query := db.Create(&rectDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRect.CheckoutPhaseOneInstance(&rectDB)
	rect := backRepo.BackRepoRect.Map_RectDBID_RectPtr[rectDB.ID]

	if rect != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), rect)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, rectDB)

	mutexRect.Unlock()
}

// GetRect
//
// swagger:route GET /rects/{ID} rects getRect
//
// Gets the details for a rect.
//
// Responses:
// default: genericError
//
//	200: rectDBResponse
func (controller *Controller) GetRect(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRect", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoRect.GetDB()

	// Get rectDB in DB
	var rectDB orm.RectDB
	if err := db.First(&rectDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var rectAPI orm.RectAPI
	rectAPI.ID = rectDB.ID
	rectAPI.RectPointersEnconding = rectDB.RectPointersEnconding
	rectDB.CopyBasicFieldsToRect(&rectAPI.Rect)

	c.JSON(http.StatusOK, rectAPI)
}

// UpdateRect
//
// swagger:route PATCH /rects/{ID} rects updateRect
//
// # Update a rect
//
// Responses:
// default: genericError
//
//	200: rectDBResponse
func (controller *Controller) UpdateRect(c *gin.Context) {

	mutexRect.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRect", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoRect.GetDB()

	// Validate input
	var input orm.RectAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rectDB orm.RectDB

	// fetch the rect
	query := db.First(&rectDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rectDB.CopyBasicFieldsFromRect(&input.Rect)
	rectDB.RectPointersEnconding = input.RectPointersEnconding

	query = db.Model(&rectDB).Updates(rectDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rectNew := new(models.Rect)
	rectDB.CopyBasicFieldsToRect(rectNew)

	// get stage instance from DB instance, and call callback function
	rectOld := backRepo.BackRepoRect.Map_RectDBID_RectPtr[rectDB.ID]
	if rectOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), rectOld, rectNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rectDB
	c.JSON(http.StatusOK, rectDB)

	mutexRect.Unlock()
}

// DeleteRect
//
// swagger:route DELETE /rects/{ID} rects deleteRect
//
// # Delete a rect
//
// default: genericError
//
//	200: rectDBResponse
func (controller *Controller) DeleteRect(c *gin.Context) {

	mutexRect.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRect", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoRect.GetDB()

	// Get model if exist
	var rectDB orm.RectDB
	if err := db.First(&rectDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&rectDB)

	// get an instance (not staged) from DB instance, and call callback function
	rectDeleted := new(models.Rect)
	rectDB.CopyBasicFieldsToRect(rectDeleted)

	// get stage instance from DB instance, and call callback function
	rectStaged := backRepo.BackRepoRect.Map_RectDBID_RectPtr[rectDB.ID]
	if rectStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), rectStaged, rectDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexRect.Unlock()
}
