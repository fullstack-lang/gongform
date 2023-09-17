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
var __Layer__dummysDeclaration__ models.Layer
var __Layer_time__dummyDeclaration time.Duration

var mutexLayer sync.Mutex

// An LayerID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLayer updateLayer deleteLayer
type LayerID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LayerInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLayer updateLayer
type LayerInput struct {
	// The Layer to submit or modify
	// in: body
	Layer *orm.LayerAPI
}

// GetLayers
//
// swagger:route GET /layers layers getLayers
//
// # Get all layers
//
// Responses:
// default: genericError
//
//	200: layerDBResponse
func (controller *Controller) GetLayers(c *gin.Context) {

	// source slice
	var layerDBs []orm.LayerDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLayers", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoLayer.GetDB()

	query := db.Find(&layerDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	layerAPIs := make([]orm.LayerAPI, 0)

	// for each layer, update fields from the database nullable fields
	for idx := range layerDBs {
		layerDB := &layerDBs[idx]
		_ = layerDB
		var layerAPI orm.LayerAPI

		// insertion point for updating fields
		layerAPI.ID = layerDB.ID
		layerDB.CopyBasicFieldsToLayer(&layerAPI.Layer)
		layerAPI.LayerPointersEnconding = layerDB.LayerPointersEnconding
		layerAPIs = append(layerAPIs, layerAPI)
	}

	c.JSON(http.StatusOK, layerAPIs)
}

// PostLayer
//
// swagger:route POST /layers layers postLayer
//
// Creates a layer
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLayer(c *gin.Context) {

	mutexLayer.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLayers", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoLayer.GetDB()

	// Validate input
	var input orm.LayerAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create layer
	layerDB := orm.LayerDB{}
	layerDB.LayerPointersEnconding = input.LayerPointersEnconding
	layerDB.CopyBasicFieldsFromLayer(&input.Layer)

	query := db.Create(&layerDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLayer.CheckoutPhaseOneInstance(&layerDB)
	layer := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[layerDB.ID]

	if layer != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), layer)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, layerDB)

	mutexLayer.Unlock()
}

// GetLayer
//
// swagger:route GET /layers/{ID} layers getLayer
//
// Gets the details for a layer.
//
// Responses:
// default: genericError
//
//	200: layerDBResponse
func (controller *Controller) GetLayer(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLayer", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoLayer.GetDB()

	// Get layerDB in DB
	var layerDB orm.LayerDB
	if err := db.First(&layerDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var layerAPI orm.LayerAPI
	layerAPI.ID = layerDB.ID
	layerAPI.LayerPointersEnconding = layerDB.LayerPointersEnconding
	layerDB.CopyBasicFieldsToLayer(&layerAPI.Layer)

	c.JSON(http.StatusOK, layerAPI)
}

// UpdateLayer
//
// swagger:route PATCH /layers/{ID} layers updateLayer
//
// # Update a layer
//
// Responses:
// default: genericError
//
//	200: layerDBResponse
func (controller *Controller) UpdateLayer(c *gin.Context) {

	mutexLayer.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLayer", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoLayer.GetDB()

	// Validate input
	var input orm.LayerAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var layerDB orm.LayerDB

	// fetch the layer
	query := db.First(&layerDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	layerDB.CopyBasicFieldsFromLayer(&input.Layer)
	layerDB.LayerPointersEnconding = input.LayerPointersEnconding

	query = db.Model(&layerDB).Updates(layerDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	layerNew := new(models.Layer)
	layerDB.CopyBasicFieldsToLayer(layerNew)

	// get stage instance from DB instance, and call callback function
	layerOld := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[layerDB.ID]
	if layerOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), layerOld, layerNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the layerDB
	c.JSON(http.StatusOK, layerDB)

	mutexLayer.Unlock()
}

// DeleteLayer
//
// swagger:route DELETE /layers/{ID} layers deleteLayer
//
// # Delete a layer
//
// default: genericError
//
//	200: layerDBResponse
func (controller *Controller) DeleteLayer(c *gin.Context) {

	mutexLayer.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLayer", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoLayer.GetDB()

	// Get model if exist
	var layerDB orm.LayerDB
	if err := db.First(&layerDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&layerDB)

	// get an instance (not staged) from DB instance, and call callback function
	layerDeleted := new(models.Layer)
	layerDB.CopyBasicFieldsToLayer(layerDeleted)

	// get stage instance from DB instance, and call callback function
	layerStaged := backRepo.BackRepoLayer.Map_LayerDBID_LayerPtr[layerDB.ID]
	if layerStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), layerStaged, layerDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexLayer.Unlock()
}
