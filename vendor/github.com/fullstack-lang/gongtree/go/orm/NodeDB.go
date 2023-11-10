// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gongtree/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Node_sql sql.NullBool
var dummy_Node_time time.Duration
var dummy_Node_sort sort.Float64Slice

// NodeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model nodeAPI
type NodeAPI struct {
	gorm.Model

	models.Node_WOP

	// encoding of pointers
	NodePointersEncoding NodePointersEncoding
}

// NodePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type NodePointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Children is a slice of pointers to another Struct (optional or 0..1)
	Children IntSlice `gorm:"type:TEXT"`

	// field Buttons is a slice of pointers to another Struct (optional or 0..1)
	Buttons IntSlice `gorm:"type:TEXT"`
}

// NodeDB describes a node in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model nodeDB
type NodeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field nodeDB.Name
	Name_Data sql.NullString

	// Declation for basic field nodeDB.BackgroundColor
	BackgroundColor_Data sql.NullString

	// Declation for basic field nodeDB.IsExpanded
	// provide the sql storage for the boolan
	IsExpanded_Data sql.NullBool

	// Declation for basic field nodeDB.HasCheckboxButton
	// provide the sql storage for the boolan
	HasCheckboxButton_Data sql.NullBool

	// Declation for basic field nodeDB.IsChecked
	// provide the sql storage for the boolan
	IsChecked_Data sql.NullBool

	// Declation for basic field nodeDB.IsCheckboxDisabled
	// provide the sql storage for the boolan
	IsCheckboxDisabled_Data sql.NullBool

	// Declation for basic field nodeDB.IsInEditMode
	// provide the sql storage for the boolan
	IsInEditMode_Data sql.NullBool

	// Declation for basic field nodeDB.IsNodeClickable
	// provide the sql storage for the boolan
	IsNodeClickable_Data sql.NullBool

	// Declation for basic field nodeDB.IsWithPreceedingIcon
	// provide the sql storage for the boolan
	IsWithPreceedingIcon_Data sql.NullBool

	// Declation for basic field nodeDB.PreceedingIcon
	PreceedingIcon_Data sql.NullString
	// encoding of pointers
	NodePointersEncoding
}

// NodeDBs arrays nodeDBs
// swagger:response nodeDBsResponse
type NodeDBs []NodeDB

// NodeDBResponse provides response
// swagger:response nodeDBResponse
type NodeDBResponse struct {
	NodeDB
}

// NodeWOP is a Node without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type NodeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	BackgroundColor string `xlsx:"2"`

	IsExpanded bool `xlsx:"3"`

	HasCheckboxButton bool `xlsx:"4"`

	IsChecked bool `xlsx:"5"`

	IsCheckboxDisabled bool `xlsx:"6"`

	IsInEditMode bool `xlsx:"7"`

	IsNodeClickable bool `xlsx:"8"`

	IsWithPreceedingIcon bool `xlsx:"9"`

	PreceedingIcon string `xlsx:"10"`
	// insertion for WOP pointer fields
}

var Node_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"BackgroundColor",
	"IsExpanded",
	"HasCheckboxButton",
	"IsChecked",
	"IsCheckboxDisabled",
	"IsInEditMode",
	"IsNodeClickable",
	"IsWithPreceedingIcon",
	"PreceedingIcon",
}

type BackRepoNodeStruct struct {
	// stores NodeDB according to their gorm ID
	Map_NodeDBID_NodeDB map[uint]*NodeDB

	// stores NodeDB ID according to Node address
	Map_NodePtr_NodeDBID map[*models.Node]uint

	// stores Node according to their gorm ID
	Map_NodeDBID_NodePtr map[uint]*models.Node

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoNode *BackRepoNodeStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoNode.stage
	return
}

func (backRepoNode *BackRepoNodeStruct) GetDB() *gorm.DB {
	return backRepoNode.db
}

// GetNodeDBFromNodePtr is a handy function to access the back repo instance from the stage instance
func (backRepoNode *BackRepoNodeStruct) GetNodeDBFromNodePtr(node *models.Node) (nodeDB *NodeDB) {
	id := backRepoNode.Map_NodePtr_NodeDBID[node]
	nodeDB = backRepoNode.Map_NodeDBID_NodeDB[id]
	return
}

// BackRepoNode.CommitPhaseOne commits all staged instances of Node to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoNode *BackRepoNodeStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for node := range stage.Nodes {
		backRepoNode.CommitPhaseOneInstance(node)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, node := range backRepoNode.Map_NodeDBID_NodePtr {
		if _, ok := stage.Nodes[node]; !ok {
			backRepoNode.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoNode.CommitDeleteInstance commits deletion of Node to the BackRepo
func (backRepoNode *BackRepoNodeStruct) CommitDeleteInstance(id uint) (Error error) {

	node := backRepoNode.Map_NodeDBID_NodePtr[id]

	// node is not staged anymore, remove nodeDB
	nodeDB := backRepoNode.Map_NodeDBID_NodeDB[id]
	query := backRepoNode.db.Unscoped().Delete(&nodeDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoNode.Map_NodePtr_NodeDBID, node)
	delete(backRepoNode.Map_NodeDBID_NodePtr, id)
	delete(backRepoNode.Map_NodeDBID_NodeDB, id)

	return
}

// BackRepoNode.CommitPhaseOneInstance commits node staged instances of Node to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoNode *BackRepoNodeStruct) CommitPhaseOneInstance(node *models.Node) (Error error) {

	// check if the node is not commited yet
	if _, ok := backRepoNode.Map_NodePtr_NodeDBID[node]; ok {
		return
	}

	// initiate node
	var nodeDB NodeDB
	nodeDB.CopyBasicFieldsFromNode(node)

	query := backRepoNode.db.Create(&nodeDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoNode.Map_NodePtr_NodeDBID[node] = nodeDB.ID
	backRepoNode.Map_NodeDBID_NodePtr[nodeDB.ID] = node
	backRepoNode.Map_NodeDBID_NodeDB[nodeDB.ID] = &nodeDB

	return
}

// BackRepoNode.CommitPhaseTwo commits all staged instances of Node to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNode *BackRepoNodeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, node := range backRepoNode.Map_NodeDBID_NodePtr {
		backRepoNode.CommitPhaseTwoInstance(backRepo, idx, node)
	}

	return
}

// BackRepoNode.CommitPhaseTwoInstance commits {{structname }} of models.Node to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNode *BackRepoNodeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, node *models.Node) (Error error) {

	// fetch matching nodeDB
	if nodeDB, ok := backRepoNode.Map_NodeDBID_NodeDB[idx]; ok {

		nodeDB.CopyBasicFieldsFromNode(node)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		nodeDB.NodePointersEncoding.Children = make([]int, 0)
		// 2. encode
		for _, nodeAssocEnd := range node.Children {
			nodeAssocEnd_DB :=
				backRepo.BackRepoNode.GetNodeDBFromNodePtr(nodeAssocEnd)
			nodeDB.NodePointersEncoding.Children =
				append(nodeDB.NodePointersEncoding.Children, int(nodeAssocEnd_DB.ID))
		}

		// 1. reset
		nodeDB.NodePointersEncoding.Buttons = make([]int, 0)
		// 2. encode
		for _, buttonAssocEnd := range node.Buttons {
			buttonAssocEnd_DB :=
				backRepo.BackRepoButton.GetButtonDBFromButtonPtr(buttonAssocEnd)
			nodeDB.NodePointersEncoding.Buttons =
				append(nodeDB.NodePointersEncoding.Buttons, int(buttonAssocEnd_DB.ID))
		}

		query := backRepoNode.db.Save(&nodeDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Node intance %s", node.Name))
		return err
	}

	return
}

// BackRepoNode.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoNode *BackRepoNodeStruct) CheckoutPhaseOne() (Error error) {

	nodeDBArray := make([]NodeDB, 0)
	query := backRepoNode.db.Find(&nodeDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	nodeInstancesToBeRemovedFromTheStage := make(map[*models.Node]any)
	for key, value := range backRepoNode.stage.Nodes {
		nodeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, nodeDB := range nodeDBArray {
		backRepoNode.CheckoutPhaseOneInstance(&nodeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		node, ok := backRepoNode.Map_NodeDBID_NodePtr[nodeDB.ID]
		if ok {
			delete(nodeInstancesToBeRemovedFromTheStage, node)
		}
	}

	// remove from stage and back repo's 3 maps all nodes that are not in the checkout
	for node := range nodeInstancesToBeRemovedFromTheStage {
		node.Unstage(backRepoNode.GetStage())

		// remove instance from the back repo 3 maps
		nodeID := backRepoNode.Map_NodePtr_NodeDBID[node]
		delete(backRepoNode.Map_NodePtr_NodeDBID, node)
		delete(backRepoNode.Map_NodeDBID_NodeDB, nodeID)
		delete(backRepoNode.Map_NodeDBID_NodePtr, nodeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a nodeDB that has been found in the DB, updates the backRepo and stages the
// models version of the nodeDB
func (backRepoNode *BackRepoNodeStruct) CheckoutPhaseOneInstance(nodeDB *NodeDB) (Error error) {

	node, ok := backRepoNode.Map_NodeDBID_NodePtr[nodeDB.ID]
	if !ok {
		node = new(models.Node)

		backRepoNode.Map_NodeDBID_NodePtr[nodeDB.ID] = node
		backRepoNode.Map_NodePtr_NodeDBID[node] = nodeDB.ID

		// append model store with the new element
		node.Name = nodeDB.Name_Data.String
		node.Stage(backRepoNode.GetStage())
	}
	nodeDB.CopyBasicFieldsToNode(node)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	node.Stage(backRepoNode.GetStage())

	// preserve pointer to nodeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_NodeDBID_NodeDB)[nodeDB hold variable pointers
	nodeDB_Data := *nodeDB
	preservedPtrToNode := &nodeDB_Data
	backRepoNode.Map_NodeDBID_NodeDB[nodeDB.ID] = preservedPtrToNode

	return
}

// BackRepoNode.CheckoutPhaseTwo Checkouts all staged instances of Node to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNode *BackRepoNodeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, nodeDB := range backRepoNode.Map_NodeDBID_NodeDB {
		backRepoNode.CheckoutPhaseTwoInstance(backRepo, nodeDB)
	}
	return
}

// BackRepoNode.CheckoutPhaseTwoInstance Checkouts staged instances of Node to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNode *BackRepoNodeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, nodeDB *NodeDB) (Error error) {

	node := backRepoNode.Map_NodeDBID_NodePtr[nodeDB.ID]
	_ = node // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// This loop redeem node.Children in the stage from the encode in the back repo
	// It parses all NodeDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	node.Children = node.Children[:0]
	for _, _Nodeid := range nodeDB.NodePointersEncoding.Children {
		node.Children = append(node.Children, backRepo.BackRepoNode.Map_NodeDBID_NodePtr[uint(_Nodeid)])
	}

	// This loop redeem node.Buttons in the stage from the encode in the back repo
	// It parses all ButtonDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	node.Buttons = node.Buttons[:0]
	for _, _Buttonid := range nodeDB.NodePointersEncoding.Buttons {
		node.Buttons = append(node.Buttons, backRepo.BackRepoButton.Map_ButtonDBID_ButtonPtr[uint(_Buttonid)])
	}

	return
}

// CommitNode allows commit of a single node (if already staged)
func (backRepo *BackRepoStruct) CommitNode(node *models.Node) {
	backRepo.BackRepoNode.CommitPhaseOneInstance(node)
	if id, ok := backRepo.BackRepoNode.Map_NodePtr_NodeDBID[node]; ok {
		backRepo.BackRepoNode.CommitPhaseTwoInstance(backRepo, id, node)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitNode allows checkout of a single node (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutNode(node *models.Node) {
	// check if the node is staged
	if _, ok := backRepo.BackRepoNode.Map_NodePtr_NodeDBID[node]; ok {

		if id, ok := backRepo.BackRepoNode.Map_NodePtr_NodeDBID[node]; ok {
			var nodeDB NodeDB
			nodeDB.ID = id

			if err := backRepo.BackRepoNode.db.First(&nodeDB, id).Error; err != nil {
				log.Fatalln("CheckoutNode : Problem with getting object with id:", id)
			}
			backRepo.BackRepoNode.CheckoutPhaseOneInstance(&nodeDB)
			backRepo.BackRepoNode.CheckoutPhaseTwoInstance(backRepo, &nodeDB)
		}
	}
}

// CopyBasicFieldsFromNode
func (nodeDB *NodeDB) CopyBasicFieldsFromNode(node *models.Node) {
	// insertion point for fields commit

	nodeDB.Name_Data.String = node.Name
	nodeDB.Name_Data.Valid = true

	nodeDB.BackgroundColor_Data.String = node.BackgroundColor
	nodeDB.BackgroundColor_Data.Valid = true

	nodeDB.IsExpanded_Data.Bool = node.IsExpanded
	nodeDB.IsExpanded_Data.Valid = true

	nodeDB.HasCheckboxButton_Data.Bool = node.HasCheckboxButton
	nodeDB.HasCheckboxButton_Data.Valid = true

	nodeDB.IsChecked_Data.Bool = node.IsChecked
	nodeDB.IsChecked_Data.Valid = true

	nodeDB.IsCheckboxDisabled_Data.Bool = node.IsCheckboxDisabled
	nodeDB.IsCheckboxDisabled_Data.Valid = true

	nodeDB.IsInEditMode_Data.Bool = node.IsInEditMode
	nodeDB.IsInEditMode_Data.Valid = true

	nodeDB.IsNodeClickable_Data.Bool = node.IsNodeClickable
	nodeDB.IsNodeClickable_Data.Valid = true

	nodeDB.IsWithPreceedingIcon_Data.Bool = node.IsWithPreceedingIcon
	nodeDB.IsWithPreceedingIcon_Data.Valid = true

	nodeDB.PreceedingIcon_Data.String = node.PreceedingIcon
	nodeDB.PreceedingIcon_Data.Valid = true
}

// CopyBasicFieldsFromNode_WOP
func (nodeDB *NodeDB) CopyBasicFieldsFromNode_WOP(node *models.Node_WOP) {
	// insertion point for fields commit

	nodeDB.Name_Data.String = node.Name
	nodeDB.Name_Data.Valid = true

	nodeDB.BackgroundColor_Data.String = node.BackgroundColor
	nodeDB.BackgroundColor_Data.Valid = true

	nodeDB.IsExpanded_Data.Bool = node.IsExpanded
	nodeDB.IsExpanded_Data.Valid = true

	nodeDB.HasCheckboxButton_Data.Bool = node.HasCheckboxButton
	nodeDB.HasCheckboxButton_Data.Valid = true

	nodeDB.IsChecked_Data.Bool = node.IsChecked
	nodeDB.IsChecked_Data.Valid = true

	nodeDB.IsCheckboxDisabled_Data.Bool = node.IsCheckboxDisabled
	nodeDB.IsCheckboxDisabled_Data.Valid = true

	nodeDB.IsInEditMode_Data.Bool = node.IsInEditMode
	nodeDB.IsInEditMode_Data.Valid = true

	nodeDB.IsNodeClickable_Data.Bool = node.IsNodeClickable
	nodeDB.IsNodeClickable_Data.Valid = true

	nodeDB.IsWithPreceedingIcon_Data.Bool = node.IsWithPreceedingIcon
	nodeDB.IsWithPreceedingIcon_Data.Valid = true

	nodeDB.PreceedingIcon_Data.String = node.PreceedingIcon
	nodeDB.PreceedingIcon_Data.Valid = true
}

// CopyBasicFieldsFromNodeWOP
func (nodeDB *NodeDB) CopyBasicFieldsFromNodeWOP(node *NodeWOP) {
	// insertion point for fields commit

	nodeDB.Name_Data.String = node.Name
	nodeDB.Name_Data.Valid = true

	nodeDB.BackgroundColor_Data.String = node.BackgroundColor
	nodeDB.BackgroundColor_Data.Valid = true

	nodeDB.IsExpanded_Data.Bool = node.IsExpanded
	nodeDB.IsExpanded_Data.Valid = true

	nodeDB.HasCheckboxButton_Data.Bool = node.HasCheckboxButton
	nodeDB.HasCheckboxButton_Data.Valid = true

	nodeDB.IsChecked_Data.Bool = node.IsChecked
	nodeDB.IsChecked_Data.Valid = true

	nodeDB.IsCheckboxDisabled_Data.Bool = node.IsCheckboxDisabled
	nodeDB.IsCheckboxDisabled_Data.Valid = true

	nodeDB.IsInEditMode_Data.Bool = node.IsInEditMode
	nodeDB.IsInEditMode_Data.Valid = true

	nodeDB.IsNodeClickable_Data.Bool = node.IsNodeClickable
	nodeDB.IsNodeClickable_Data.Valid = true

	nodeDB.IsWithPreceedingIcon_Data.Bool = node.IsWithPreceedingIcon
	nodeDB.IsWithPreceedingIcon_Data.Valid = true

	nodeDB.PreceedingIcon_Data.String = node.PreceedingIcon
	nodeDB.PreceedingIcon_Data.Valid = true
}

// CopyBasicFieldsToNode
func (nodeDB *NodeDB) CopyBasicFieldsToNode(node *models.Node) {
	// insertion point for checkout of basic fields (back repo to stage)
	node.Name = nodeDB.Name_Data.String
	node.BackgroundColor = nodeDB.BackgroundColor_Data.String
	node.IsExpanded = nodeDB.IsExpanded_Data.Bool
	node.HasCheckboxButton = nodeDB.HasCheckboxButton_Data.Bool
	node.IsChecked = nodeDB.IsChecked_Data.Bool
	node.IsCheckboxDisabled = nodeDB.IsCheckboxDisabled_Data.Bool
	node.IsInEditMode = nodeDB.IsInEditMode_Data.Bool
	node.IsNodeClickable = nodeDB.IsNodeClickable_Data.Bool
	node.IsWithPreceedingIcon = nodeDB.IsWithPreceedingIcon_Data.Bool
	node.PreceedingIcon = nodeDB.PreceedingIcon_Data.String
}

// CopyBasicFieldsToNode_WOP
func (nodeDB *NodeDB) CopyBasicFieldsToNode_WOP(node *models.Node_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	node.Name = nodeDB.Name_Data.String
	node.BackgroundColor = nodeDB.BackgroundColor_Data.String
	node.IsExpanded = nodeDB.IsExpanded_Data.Bool
	node.HasCheckboxButton = nodeDB.HasCheckboxButton_Data.Bool
	node.IsChecked = nodeDB.IsChecked_Data.Bool
	node.IsCheckboxDisabled = nodeDB.IsCheckboxDisabled_Data.Bool
	node.IsInEditMode = nodeDB.IsInEditMode_Data.Bool
	node.IsNodeClickable = nodeDB.IsNodeClickable_Data.Bool
	node.IsWithPreceedingIcon = nodeDB.IsWithPreceedingIcon_Data.Bool
	node.PreceedingIcon = nodeDB.PreceedingIcon_Data.String
}

// CopyBasicFieldsToNodeWOP
func (nodeDB *NodeDB) CopyBasicFieldsToNodeWOP(node *NodeWOP) {
	node.ID = int(nodeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	node.Name = nodeDB.Name_Data.String
	node.BackgroundColor = nodeDB.BackgroundColor_Data.String
	node.IsExpanded = nodeDB.IsExpanded_Data.Bool
	node.HasCheckboxButton = nodeDB.HasCheckboxButton_Data.Bool
	node.IsChecked = nodeDB.IsChecked_Data.Bool
	node.IsCheckboxDisabled = nodeDB.IsCheckboxDisabled_Data.Bool
	node.IsInEditMode = nodeDB.IsInEditMode_Data.Bool
	node.IsNodeClickable = nodeDB.IsNodeClickable_Data.Bool
	node.IsWithPreceedingIcon = nodeDB.IsWithPreceedingIcon_Data.Bool
	node.PreceedingIcon = nodeDB.PreceedingIcon_Data.String
}

// Backup generates a json file from a slice of all NodeDB instances in the backrepo
func (backRepoNode *BackRepoNodeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "NodeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*NodeDB, 0)
	for _, nodeDB := range backRepoNode.Map_NodeDBID_NodeDB {
		forBackup = append(forBackup, nodeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Node ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Node file", err.Error())
	}
}

// Backup generates a json file from a slice of all NodeDB instances in the backrepo
func (backRepoNode *BackRepoNodeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*NodeDB, 0)
	for _, nodeDB := range backRepoNode.Map_NodeDBID_NodeDB {
		forBackup = append(forBackup, nodeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Node")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Node_Fields, -1)
	for _, nodeDB := range forBackup {

		var nodeWOP NodeWOP
		nodeDB.CopyBasicFieldsToNodeWOP(&nodeWOP)

		row := sh.AddRow()
		row.WriteStruct(&nodeWOP, -1)
	}
}

// RestoreXL from the "Node" sheet all NodeDB instances
func (backRepoNode *BackRepoNodeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoNodeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Node"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoNode.rowVisitorNode)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoNode *BackRepoNodeStruct) rowVisitorNode(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var nodeWOP NodeWOP
		row.ReadStruct(&nodeWOP)

		// add the unmarshalled struct to the stage
		nodeDB := new(NodeDB)
		nodeDB.CopyBasicFieldsFromNodeWOP(&nodeWOP)

		nodeDB_ID_atBackupTime := nodeDB.ID
		nodeDB.ID = 0
		query := backRepoNode.db.Create(nodeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoNode.Map_NodeDBID_NodeDB[nodeDB.ID] = nodeDB
		BackRepoNodeid_atBckpTime_newID[nodeDB_ID_atBackupTime] = nodeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "NodeDB.json" in dirPath that stores an array
// of NodeDB and stores it in the database
// the map BackRepoNodeid_atBckpTime_newID is updated accordingly
func (backRepoNode *BackRepoNodeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoNodeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "NodeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Node file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*NodeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_NodeDBID_NodeDB
	for _, nodeDB := range forRestore {

		nodeDB_ID_atBackupTime := nodeDB.ID
		nodeDB.ID = 0
		query := backRepoNode.db.Create(nodeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoNode.Map_NodeDBID_NodeDB[nodeDB.ID] = nodeDB
		BackRepoNodeid_atBckpTime_newID[nodeDB_ID_atBackupTime] = nodeDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Node file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Node>id_atBckpTime_newID
// to compute new index
func (backRepoNode *BackRepoNodeStruct) RestorePhaseTwo() {

	for _, nodeDB := range backRepoNode.Map_NodeDBID_NodeDB {

		// next line of code is to avert unused variable compilation error
		_ = nodeDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoNode.db.Model(nodeDB).Updates(*nodeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoNode.ResetReversePointers commits all staged instances of Node to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNode *BackRepoNodeStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, node := range backRepoNode.Map_NodeDBID_NodePtr {
		backRepoNode.ResetReversePointersInstance(backRepo, idx, node)
	}

	return
}

func (backRepoNode *BackRepoNodeStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, astruct *models.Node) (Error error) {

	// fetch matching nodeDB
	if nodeDB, ok := backRepoNode.Map_NodeDBID_NodeDB[idx]; ok {
		_ = nodeDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoNodeid_atBckpTime_newID map[uint]uint
