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
var __GongStructShape__dummysDeclaration__ models.GongStructShape
var __GongStructShape_time__dummyDeclaration time.Duration

var mutexGongStructShape sync.Mutex

// An GongStructShapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongStructShape updateGongStructShape deleteGongStructShape
type GongStructShapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongStructShapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongStructShape updateGongStructShape
type GongStructShapeInput struct {
	// The GongStructShape to submit or modify
	// in: body
	GongStructShape *orm.GongStructShapeAPI
}

// GetGongStructShapes
//
// swagger:route GET /gongstructshapes gongstructshapes getGongStructShapes
//
// # Get all gongstructshapes
//
// Responses:
// default: genericError
//
//	200: gongstructshapeDBResponse
func (controller *Controller) GetGongStructShapes(c *gin.Context) {

	// source slice
	var gongstructshapeDBs []orm.GongStructShapeDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongStructShapes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoGongStructShape.GetDB()

	query := db.Find(&gongstructshapeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongstructshapeAPIs := make([]orm.GongStructShapeAPI, 0)

	// for each gongstructshape, update fields from the database nullable fields
	for idx := range gongstructshapeDBs {
		gongstructshapeDB := &gongstructshapeDBs[idx]
		_ = gongstructshapeDB
		var gongstructshapeAPI orm.GongStructShapeAPI

		// insertion point for updating fields
		gongstructshapeAPI.ID = gongstructshapeDB.ID
		gongstructshapeDB.CopyBasicFieldsToGongStructShape(&gongstructshapeAPI.GongStructShape)
		gongstructshapeAPI.GongStructShapePointersEnconding = gongstructshapeDB.GongStructShapePointersEnconding
		gongstructshapeAPIs = append(gongstructshapeAPIs, gongstructshapeAPI)
	}

	c.JSON(http.StatusOK, gongstructshapeAPIs)
}

// PostGongStructShape
//
// swagger:route POST /gongstructshapes gongstructshapes postGongStructShape
//
// Creates a gongstructshape
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGongStructShape(c *gin.Context) {

	mutexGongStructShape.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGongStructShapes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoGongStructShape.GetDB()

	// Validate input
	var input orm.GongStructShapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongstructshape
	gongstructshapeDB := orm.GongStructShapeDB{}
	gongstructshapeDB.GongStructShapePointersEnconding = input.GongStructShapePointersEnconding
	gongstructshapeDB.CopyBasicFieldsFromGongStructShape(&input.GongStructShape)

	query := db.Create(&gongstructshapeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGongStructShape.CheckoutPhaseOneInstance(&gongstructshapeDB)
	gongstructshape := backRepo.BackRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr[gongstructshapeDB.ID]

	if gongstructshape != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gongstructshape)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongstructshapeDB)

	mutexGongStructShape.Unlock()
}

// GetGongStructShape
//
// swagger:route GET /gongstructshapes/{ID} gongstructshapes getGongStructShape
//
// Gets the details for a gongstructshape.
//
// Responses:
// default: genericError
//
//	200: gongstructshapeDBResponse
func (controller *Controller) GetGongStructShape(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongStructShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoGongStructShape.GetDB()

	// Get gongstructshapeDB in DB
	var gongstructshapeDB orm.GongStructShapeDB
	if err := db.First(&gongstructshapeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongstructshapeAPI orm.GongStructShapeAPI
	gongstructshapeAPI.ID = gongstructshapeDB.ID
	gongstructshapeAPI.GongStructShapePointersEnconding = gongstructshapeDB.GongStructShapePointersEnconding
	gongstructshapeDB.CopyBasicFieldsToGongStructShape(&gongstructshapeAPI.GongStructShape)

	c.JSON(http.StatusOK, gongstructshapeAPI)
}

// UpdateGongStructShape
//
// swagger:route PATCH /gongstructshapes/{ID} gongstructshapes updateGongStructShape
//
// # Update a gongstructshape
//
// Responses:
// default: genericError
//
//	200: gongstructshapeDBResponse
func (controller *Controller) UpdateGongStructShape(c *gin.Context) {

	mutexGongStructShape.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGongStructShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoGongStructShape.GetDB()

	// Validate input
	var input orm.GongStructShapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gongstructshapeDB orm.GongStructShapeDB

	// fetch the gongstructshape
	query := db.First(&gongstructshapeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gongstructshapeDB.CopyBasicFieldsFromGongStructShape(&input.GongStructShape)
	gongstructshapeDB.GongStructShapePointersEnconding = input.GongStructShapePointersEnconding

	query = db.Model(&gongstructshapeDB).Updates(gongstructshapeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongstructshapeNew := new(models.GongStructShape)
	gongstructshapeDB.CopyBasicFieldsToGongStructShape(gongstructshapeNew)

	// get stage instance from DB instance, and call callback function
	gongstructshapeOld := backRepo.BackRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr[gongstructshapeDB.ID]
	if gongstructshapeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), gongstructshapeOld, gongstructshapeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongstructshapeDB
	c.JSON(http.StatusOK, gongstructshapeDB)

	mutexGongStructShape.Unlock()
}

// DeleteGongStructShape
//
// swagger:route DELETE /gongstructshapes/{ID} gongstructshapes deleteGongStructShape
//
// # Delete a gongstructshape
//
// default: genericError
//
//	200: gongstructshapeDBResponse
func (controller *Controller) DeleteGongStructShape(c *gin.Context) {

	mutexGongStructShape.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGongStructShape", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoGongStructShape.GetDB()

	// Get model if exist
	var gongstructshapeDB orm.GongStructShapeDB
	if err := db.First(&gongstructshapeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&gongstructshapeDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongstructshapeDeleted := new(models.GongStructShape)
	gongstructshapeDB.CopyBasicFieldsToGongStructShape(gongstructshapeDeleted)

	// get stage instance from DB instance, and call callback function
	gongstructshapeStaged := backRepo.BackRepoGongStructShape.Map_GongStructShapeDBID_GongStructShapePtr[gongstructshapeDB.ID]
	if gongstructshapeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), gongstructshapeStaged, gongstructshapeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexGongStructShape.Unlock()
}
