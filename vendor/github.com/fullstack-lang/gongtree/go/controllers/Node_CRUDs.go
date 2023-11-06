// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongtree/go/models"
	"github.com/fullstack-lang/gongtree/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Node__dummysDeclaration__ models.Node
var __Node_time__dummyDeclaration time.Duration

var mutexNode sync.Mutex

// An NodeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getNode updateNode deleteNode
type NodeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// NodeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postNode updateNode
type NodeInput struct {
	// The Node to submit or modify
	// in: body
	Node *orm.NodeAPI
}

// GetNodes
//
// swagger:route GET /nodes nodes getNodes
//
// # Get all nodes
//
// Responses:
// default: genericError
//
//	200: nodeDBResponse
func (controller *Controller) GetNodes(c *gin.Context) {

	// source slice
	var nodeDBs []orm.NodeDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNodes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNode.GetDB()

	query := db.Find(&nodeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	nodeAPIs := make([]orm.NodeAPI, 0)

	// for each node, update fields from the database nullable fields
	for idx := range nodeDBs {
		nodeDB := &nodeDBs[idx]
		_ = nodeDB
		var nodeAPI orm.NodeAPI

		// insertion point for updating fields
		nodeAPI.ID = nodeDB.ID
		nodeDB.CopyBasicFieldsToNode_WOP(&nodeAPI.Node_WOP)
		nodeAPI.NodePointersEncoding = nodeDB.NodePointersEncoding
		nodeAPIs = append(nodeAPIs, nodeAPI)
	}

	c.JSON(http.StatusOK, nodeAPIs)
}

// PostNode
//
// swagger:route POST /nodes nodes postNode
//
// Creates a node
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostNode(c *gin.Context) {

	mutexNode.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostNodes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNode.GetDB()

	// Validate input
	var input orm.NodeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create node
	nodeDB := orm.NodeDB{}
	nodeDB.NodePointersEncoding = input.NodePointersEncoding
	nodeDB.CopyBasicFieldsFromNode_WOP(&input.Node_WOP)

	query := db.Create(&nodeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoNode.CheckoutPhaseOneInstance(&nodeDB)
	node := backRepo.BackRepoNode.Map_NodeDBID_NodePtr[nodeDB.ID]

	if node != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), node)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, nodeDB)

	mutexNode.Unlock()
}

// GetNode
//
// swagger:route GET /nodes/{ID} nodes getNode
//
// Gets the details for a node.
//
// Responses:
// default: genericError
//
//	200: nodeDBResponse
func (controller *Controller) GetNode(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNode", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNode.GetDB()

	// Get nodeDB in DB
	var nodeDB orm.NodeDB
	if err := db.First(&nodeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var nodeAPI orm.NodeAPI
	nodeAPI.ID = nodeDB.ID
	nodeAPI.NodePointersEncoding = nodeDB.NodePointersEncoding
	nodeDB.CopyBasicFieldsToNode_WOP(&nodeAPI.Node_WOP)

	c.JSON(http.StatusOK, nodeAPI)
}

// UpdateNode
//
// swagger:route PATCH /nodes/{ID} nodes updateNode
//
// # Update a node
//
// Responses:
// default: genericError
//
//	200: nodeDBResponse
func (controller *Controller) UpdateNode(c *gin.Context) {

	mutexNode.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateNode", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNode.GetDB()

	// Validate input
	var input orm.NodeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var nodeDB orm.NodeDB

	// fetch the node
	query := db.First(&nodeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	nodeDB.CopyBasicFieldsFromNode_WOP(&input.Node_WOP)
	nodeDB.NodePointersEncoding = input.NodePointersEncoding

	query = db.Model(&nodeDB).Updates(nodeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	nodeNew := new(models.Node)
	nodeDB.CopyBasicFieldsToNode(nodeNew)

	// get stage instance from DB instance, and call callback function
	nodeOld := backRepo.BackRepoNode.Map_NodeDBID_NodePtr[nodeDB.ID]
	if nodeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), nodeOld, nodeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the nodeDB
	c.JSON(http.StatusOK, nodeDB)

	mutexNode.Unlock()
}

// DeleteNode
//
// swagger:route DELETE /nodes/{ID} nodes deleteNode
//
// # Delete a node
//
// default: genericError
//
//	200: nodeDBResponse
func (controller *Controller) DeleteNode(c *gin.Context) {

	mutexNode.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteNode", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongtree/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNode.GetDB()

	// Get model if exist
	var nodeDB orm.NodeDB
	if err := db.First(&nodeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&nodeDB)

	// get an instance (not staged) from DB instance, and call callback function
	nodeDeleted := new(models.Node)
	nodeDB.CopyBasicFieldsToNode(nodeDeleted)

	// get stage instance from DB instance, and call callback function
	nodeStaged := backRepo.BackRepoNode.Map_NodeDBID_NodePtr[nodeDB.ID]
	if nodeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), nodeStaged, nodeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexNode.Unlock()
}
