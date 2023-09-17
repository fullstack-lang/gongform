// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Meta__dummysDeclaration__ models.Meta
var __Meta_time__dummyDeclaration time.Duration

var mutexMeta sync.Mutex

// An MetaID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMeta updateMeta deleteMeta
type MetaID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MetaInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMeta updateMeta
type MetaInput struct {
	// The Meta to submit or modify
	// in: body
	Meta *orm.MetaAPI
}

// GetMetas
//
// swagger:route GET /metas metas getMetas
//
// # Get all metas
//
// Responses:
// default: genericError
//
//	200: metaDBResponse
func (controller *Controller) GetMetas(c *gin.Context) {

	// source slice
	var metaDBs []orm.MetaDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMetas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoMeta.GetDB()

	query := db.Find(&metaDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	metaAPIs := make([]orm.MetaAPI, 0)

	// for each meta, update fields from the database nullable fields
	for idx := range metaDBs {
		metaDB := &metaDBs[idx]
		_ = metaDB
		var metaAPI orm.MetaAPI

		// insertion point for updating fields
		metaAPI.ID = metaDB.ID
		metaDB.CopyBasicFieldsToMeta(&metaAPI.Meta)
		metaAPI.MetaPointersEnconding = metaDB.MetaPointersEnconding
		metaAPIs = append(metaAPIs, metaAPI)
	}

	c.JSON(http.StatusOK, metaAPIs)
}

// PostMeta
//
// swagger:route POST /metas metas postMeta
//
// Creates a meta
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMeta(c *gin.Context) {

	mutexMeta.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMetas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoMeta.GetDB()

	// Validate input
	var input orm.MetaAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create meta
	metaDB := orm.MetaDB{}
	metaDB.MetaPointersEnconding = input.MetaPointersEnconding
	metaDB.CopyBasicFieldsFromMeta(&input.Meta)

	query := db.Create(&metaDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMeta.CheckoutPhaseOneInstance(&metaDB)
	meta := backRepo.BackRepoMeta.Map_MetaDBID_MetaPtr[metaDB.ID]

	if meta != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), meta)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, metaDB)

	mutexMeta.Unlock()
}

// GetMeta
//
// swagger:route GET /metas/{ID} metas getMeta
//
// Gets the details for a meta.
//
// Responses:
// default: genericError
//
//	200: metaDBResponse
func (controller *Controller) GetMeta(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMeta", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoMeta.GetDB()

	// Get metaDB in DB
	var metaDB orm.MetaDB
	if err := db.First(&metaDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var metaAPI orm.MetaAPI
	metaAPI.ID = metaDB.ID
	metaAPI.MetaPointersEnconding = metaDB.MetaPointersEnconding
	metaDB.CopyBasicFieldsToMeta(&metaAPI.Meta)

	c.JSON(http.StatusOK, metaAPI)
}

// UpdateMeta
//
// swagger:route PATCH /metas/{ID} metas updateMeta
//
// # Update a meta
//
// Responses:
// default: genericError
//
//	200: metaDBResponse
func (controller *Controller) UpdateMeta(c *gin.Context) {

	mutexMeta.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMeta", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoMeta.GetDB()

	// Validate input
	var input orm.MetaAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var metaDB orm.MetaDB

	// fetch the meta
	query := db.First(&metaDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	metaDB.CopyBasicFieldsFromMeta(&input.Meta)
	metaDB.MetaPointersEnconding = input.MetaPointersEnconding

	query = db.Model(&metaDB).Updates(metaDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	metaNew := new(models.Meta)
	metaDB.CopyBasicFieldsToMeta(metaNew)

	// get stage instance from DB instance, and call callback function
	metaOld := backRepo.BackRepoMeta.Map_MetaDBID_MetaPtr[metaDB.ID]
	if metaOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), metaOld, metaNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the metaDB
	c.JSON(http.StatusOK, metaDB)

	mutexMeta.Unlock()
}

// DeleteMeta
//
// swagger:route DELETE /metas/{ID} metas deleteMeta
//
// # Delete a meta
//
// default: genericError
//
//	200: metaDBResponse
func (controller *Controller) DeleteMeta(c *gin.Context) {

	mutexMeta.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMeta", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	db := backRepo.BackRepoMeta.GetDB()

	// Get model if exist
	var metaDB orm.MetaDB
	if err := db.First(&metaDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&metaDB)

	// get an instance (not staged) from DB instance, and call callback function
	metaDeleted := new(models.Meta)
	metaDB.CopyBasicFieldsToMeta(metaDeleted)

	// get stage instance from DB instance, and call callback function
	metaStaged := backRepo.BackRepoMeta.Map_MetaDBID_MetaPtr[metaDB.ID]
	if metaStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), metaStaged, metaDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexMeta.Unlock()
}