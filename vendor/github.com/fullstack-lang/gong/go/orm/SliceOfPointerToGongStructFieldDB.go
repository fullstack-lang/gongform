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

	"github.com/fullstack-lang/gong/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_SliceOfPointerToGongStructField_sql sql.NullBool
var dummy_SliceOfPointerToGongStructField_time time.Duration
var dummy_SliceOfPointerToGongStructField_sort sort.Float64Slice

// SliceOfPointerToGongStructFieldAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model sliceofpointertogongstructfieldAPI
type SliceOfPointerToGongStructFieldAPI struct {
	gorm.Model

	models.SliceOfPointerToGongStructField

	// encoding of pointers
	SliceOfPointerToGongStructFieldPointersEnconding
}

// SliceOfPointerToGongStructFieldPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SliceOfPointerToGongStructFieldPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// field GongStruct is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	GongStructID sql.NullInt64

	// Implementation of a reverse ID for field GongStruct{}.SliceOfPointerToGongStructFields []*SliceOfPointerToGongStructField
	GongStruct_SliceOfPointerToGongStructFieldsDBID sql.NullInt64

	// implementation of the index of the withing the slice
	GongStruct_SliceOfPointerToGongStructFieldsDBID_Index sql.NullInt64
}

// SliceOfPointerToGongStructFieldDB describes a sliceofpointertogongstructfield in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model sliceofpointertogongstructfieldDB
type SliceOfPointerToGongStructFieldDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field sliceofpointertogongstructfieldDB.Name
	Name_Data sql.NullString

	// Declation for basic field sliceofpointertogongstructfieldDB.Index
	Index_Data sql.NullInt64

	// Declation for basic field sliceofpointertogongstructfieldDB.CompositeStructName
	CompositeStructName_Data sql.NullString
	// encoding of pointers
	SliceOfPointerToGongStructFieldPointersEnconding
}

// SliceOfPointerToGongStructFieldDBs arrays sliceofpointertogongstructfieldDBs
// swagger:response sliceofpointertogongstructfieldDBsResponse
type SliceOfPointerToGongStructFieldDBs []SliceOfPointerToGongStructFieldDB

// SliceOfPointerToGongStructFieldDBResponse provides response
// swagger:response sliceofpointertogongstructfieldDBResponse
type SliceOfPointerToGongStructFieldDBResponse struct {
	SliceOfPointerToGongStructFieldDB
}

// SliceOfPointerToGongStructFieldWOP is a SliceOfPointerToGongStructField without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SliceOfPointerToGongStructFieldWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Index int `xlsx:"2"`

	CompositeStructName string `xlsx:"3"`
	// insertion for WOP pointer fields
}

var SliceOfPointerToGongStructField_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Index",
	"CompositeStructName",
}

type BackRepoSliceOfPointerToGongStructFieldStruct struct {
	// stores SliceOfPointerToGongStructFieldDB according to their gorm ID
	Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB map[uint]*SliceOfPointerToGongStructFieldDB

	// stores SliceOfPointerToGongStructFieldDB ID according to SliceOfPointerToGongStructField address
	Map_SliceOfPointerToGongStructFieldPtr_SliceOfPointerToGongStructFieldDBID map[*models.SliceOfPointerToGongStructField]uint

	// stores SliceOfPointerToGongStructField according to their gorm ID
	Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldPtr map[uint]*models.SliceOfPointerToGongStructField

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoSliceOfPointerToGongStructField.stage
	return
}

func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) GetDB() *gorm.DB {
	return backRepoSliceOfPointerToGongStructField.db
}

// GetSliceOfPointerToGongStructFieldDBFromSliceOfPointerToGongStructFieldPtr is a handy function to access the back repo instance from the stage instance
func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) GetSliceOfPointerToGongStructFieldDBFromSliceOfPointerToGongStructFieldPtr(sliceofpointertogongstructfield *models.SliceOfPointerToGongStructField) (sliceofpointertogongstructfieldDB *SliceOfPointerToGongStructFieldDB) {
	id := backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldPtr_SliceOfPointerToGongStructFieldDBID[sliceofpointertogongstructfield]
	sliceofpointertogongstructfieldDB = backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB[id]
	return
}

// BackRepoSliceOfPointerToGongStructField.CommitPhaseOne commits all staged instances of SliceOfPointerToGongStructField to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for sliceofpointertogongstructfield := range stage.SliceOfPointerToGongStructFields {
		backRepoSliceOfPointerToGongStructField.CommitPhaseOneInstance(sliceofpointertogongstructfield)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, sliceofpointertogongstructfield := range backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldPtr {
		if _, ok := stage.SliceOfPointerToGongStructFields[sliceofpointertogongstructfield]; !ok {
			backRepoSliceOfPointerToGongStructField.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSliceOfPointerToGongStructField.CommitDeleteInstance commits deletion of SliceOfPointerToGongStructField to the BackRepo
func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) CommitDeleteInstance(id uint) (Error error) {

	sliceofpointertogongstructfield := backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldPtr[id]

	// sliceofpointertogongstructfield is not staged anymore, remove sliceofpointertogongstructfieldDB
	sliceofpointertogongstructfieldDB := backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB[id]
	query := backRepoSliceOfPointerToGongStructField.db.Unscoped().Delete(&sliceofpointertogongstructfieldDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldPtr_SliceOfPointerToGongStructFieldDBID, sliceofpointertogongstructfield)
	delete(backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldPtr, id)
	delete(backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB, id)

	return
}

// BackRepoSliceOfPointerToGongStructField.CommitPhaseOneInstance commits sliceofpointertogongstructfield staged instances of SliceOfPointerToGongStructField to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) CommitPhaseOneInstance(sliceofpointertogongstructfield *models.SliceOfPointerToGongStructField) (Error error) {

	// check if the sliceofpointertogongstructfield is not commited yet
	if _, ok := backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldPtr_SliceOfPointerToGongStructFieldDBID[sliceofpointertogongstructfield]; ok {
		return
	}

	// initiate sliceofpointertogongstructfield
	var sliceofpointertogongstructfieldDB SliceOfPointerToGongStructFieldDB
	sliceofpointertogongstructfieldDB.CopyBasicFieldsFromSliceOfPointerToGongStructField(sliceofpointertogongstructfield)

	query := backRepoSliceOfPointerToGongStructField.db.Create(&sliceofpointertogongstructfieldDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldPtr_SliceOfPointerToGongStructFieldDBID[sliceofpointertogongstructfield] = sliceofpointertogongstructfieldDB.ID
	backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldPtr[sliceofpointertogongstructfieldDB.ID] = sliceofpointertogongstructfield
	backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB[sliceofpointertogongstructfieldDB.ID] = &sliceofpointertogongstructfieldDB

	return
}

// BackRepoSliceOfPointerToGongStructField.CommitPhaseTwo commits all staged instances of SliceOfPointerToGongStructField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, sliceofpointertogongstructfield := range backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldPtr {
		backRepoSliceOfPointerToGongStructField.CommitPhaseTwoInstance(backRepo, idx, sliceofpointertogongstructfield)
	}

	return
}

// BackRepoSliceOfPointerToGongStructField.CommitPhaseTwoInstance commits {{structname }} of models.SliceOfPointerToGongStructField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, sliceofpointertogongstructfield *models.SliceOfPointerToGongStructField) (Error error) {

	// fetch matching sliceofpointertogongstructfieldDB
	if sliceofpointertogongstructfieldDB, ok := backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB[idx]; ok {

		sliceofpointertogongstructfieldDB.CopyBasicFieldsFromSliceOfPointerToGongStructField(sliceofpointertogongstructfield)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value sliceofpointertogongstructfield.GongStruct translates to updating the sliceofpointertogongstructfield.GongStructID
		sliceofpointertogongstructfieldDB.GongStructID.Valid = true // allow for a 0 value (nil association)
		if sliceofpointertogongstructfield.GongStruct != nil {
			if GongStructId, ok := backRepo.BackRepoGongStruct.Map_GongStructPtr_GongStructDBID[sliceofpointertogongstructfield.GongStruct]; ok {
				sliceofpointertogongstructfieldDB.GongStructID.Int64 = int64(GongStructId)
				sliceofpointertogongstructfieldDB.GongStructID.Valid = true
			}
		} else {
			sliceofpointertogongstructfieldDB.GongStructID.Int64 = 0
			sliceofpointertogongstructfieldDB.GongStructID.Valid = true
		}

		query := backRepoSliceOfPointerToGongStructField.db.Save(&sliceofpointertogongstructfieldDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown SliceOfPointerToGongStructField intance %s", sliceofpointertogongstructfield.Name))
		return err
	}

	return
}

// BackRepoSliceOfPointerToGongStructField.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) CheckoutPhaseOne() (Error error) {

	sliceofpointertogongstructfieldDBArray := make([]SliceOfPointerToGongStructFieldDB, 0)
	query := backRepoSliceOfPointerToGongStructField.db.Find(&sliceofpointertogongstructfieldDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	sliceofpointertogongstructfieldInstancesToBeRemovedFromTheStage := make(map[*models.SliceOfPointerToGongStructField]any)
	for key, value := range backRepoSliceOfPointerToGongStructField.stage.SliceOfPointerToGongStructFields {
		sliceofpointertogongstructfieldInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, sliceofpointertogongstructfieldDB := range sliceofpointertogongstructfieldDBArray {
		backRepoSliceOfPointerToGongStructField.CheckoutPhaseOneInstance(&sliceofpointertogongstructfieldDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		sliceofpointertogongstructfield, ok := backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldPtr[sliceofpointertogongstructfieldDB.ID]
		if ok {
			delete(sliceofpointertogongstructfieldInstancesToBeRemovedFromTheStage, sliceofpointertogongstructfield)
		}
	}

	// remove from stage and back repo's 3 maps all sliceofpointertogongstructfields that are not in the checkout
	for sliceofpointertogongstructfield := range sliceofpointertogongstructfieldInstancesToBeRemovedFromTheStage {
		sliceofpointertogongstructfield.Unstage(backRepoSliceOfPointerToGongStructField.GetStage())

		// remove instance from the back repo 3 maps
		sliceofpointertogongstructfieldID := backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldPtr_SliceOfPointerToGongStructFieldDBID[sliceofpointertogongstructfield]
		delete(backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldPtr_SliceOfPointerToGongStructFieldDBID, sliceofpointertogongstructfield)
		delete(backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB, sliceofpointertogongstructfieldID)
		delete(backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldPtr, sliceofpointertogongstructfieldID)
	}

	return
}

// CheckoutPhaseOneInstance takes a sliceofpointertogongstructfieldDB that has been found in the DB, updates the backRepo and stages the
// models version of the sliceofpointertogongstructfieldDB
func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) CheckoutPhaseOneInstance(sliceofpointertogongstructfieldDB *SliceOfPointerToGongStructFieldDB) (Error error) {

	sliceofpointertogongstructfield, ok := backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldPtr[sliceofpointertogongstructfieldDB.ID]
	if !ok {
		sliceofpointertogongstructfield = new(models.SliceOfPointerToGongStructField)

		backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldPtr[sliceofpointertogongstructfieldDB.ID] = sliceofpointertogongstructfield
		backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldPtr_SliceOfPointerToGongStructFieldDBID[sliceofpointertogongstructfield] = sliceofpointertogongstructfieldDB.ID

		// append model store with the new element
		sliceofpointertogongstructfield.Name = sliceofpointertogongstructfieldDB.Name_Data.String
		sliceofpointertogongstructfield.Stage(backRepoSliceOfPointerToGongStructField.GetStage())
	}
	sliceofpointertogongstructfieldDB.CopyBasicFieldsToSliceOfPointerToGongStructField(sliceofpointertogongstructfield)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	sliceofpointertogongstructfield.Stage(backRepoSliceOfPointerToGongStructField.GetStage())

	// preserve pointer to sliceofpointertogongstructfieldDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB)[sliceofpointertogongstructfieldDB hold variable pointers
	sliceofpointertogongstructfieldDB_Data := *sliceofpointertogongstructfieldDB
	preservedPtrToSliceOfPointerToGongStructField := &sliceofpointertogongstructfieldDB_Data
	backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB[sliceofpointertogongstructfieldDB.ID] = preservedPtrToSliceOfPointerToGongStructField

	return
}

// BackRepoSliceOfPointerToGongStructField.CheckoutPhaseTwo Checkouts all staged instances of SliceOfPointerToGongStructField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, sliceofpointertogongstructfieldDB := range backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB {
		backRepoSliceOfPointerToGongStructField.CheckoutPhaseTwoInstance(backRepo, sliceofpointertogongstructfieldDB)
	}
	return
}

// BackRepoSliceOfPointerToGongStructField.CheckoutPhaseTwoInstance Checkouts staged instances of SliceOfPointerToGongStructField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, sliceofpointertogongstructfieldDB *SliceOfPointerToGongStructFieldDB) (Error error) {

	sliceofpointertogongstructfield := backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldPtr[sliceofpointertogongstructfieldDB.ID]
	_ = sliceofpointertogongstructfield // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// GongStruct field
	sliceofpointertogongstructfield.GongStruct = nil
	if sliceofpointertogongstructfieldDB.GongStructID.Int64 != 0 {
		sliceofpointertogongstructfield.GongStruct = backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr[uint(sliceofpointertogongstructfieldDB.GongStructID.Int64)]
	}
	return
}

// CommitSliceOfPointerToGongStructField allows commit of a single sliceofpointertogongstructfield (if already staged)
func (backRepo *BackRepoStruct) CommitSliceOfPointerToGongStructField(sliceofpointertogongstructfield *models.SliceOfPointerToGongStructField) {
	backRepo.BackRepoSliceOfPointerToGongStructField.CommitPhaseOneInstance(sliceofpointertogongstructfield)
	if id, ok := backRepo.BackRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldPtr_SliceOfPointerToGongStructFieldDBID[sliceofpointertogongstructfield]; ok {
		backRepo.BackRepoSliceOfPointerToGongStructField.CommitPhaseTwoInstance(backRepo, id, sliceofpointertogongstructfield)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitSliceOfPointerToGongStructField allows checkout of a single sliceofpointertogongstructfield (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSliceOfPointerToGongStructField(sliceofpointertogongstructfield *models.SliceOfPointerToGongStructField) {
	// check if the sliceofpointertogongstructfield is staged
	if _, ok := backRepo.BackRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldPtr_SliceOfPointerToGongStructFieldDBID[sliceofpointertogongstructfield]; ok {

		if id, ok := backRepo.BackRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldPtr_SliceOfPointerToGongStructFieldDBID[sliceofpointertogongstructfield]; ok {
			var sliceofpointertogongstructfieldDB SliceOfPointerToGongStructFieldDB
			sliceofpointertogongstructfieldDB.ID = id

			if err := backRepo.BackRepoSliceOfPointerToGongStructField.db.First(&sliceofpointertogongstructfieldDB, id).Error; err != nil {
				log.Panicln("CheckoutSliceOfPointerToGongStructField : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSliceOfPointerToGongStructField.CheckoutPhaseOneInstance(&sliceofpointertogongstructfieldDB)
			backRepo.BackRepoSliceOfPointerToGongStructField.CheckoutPhaseTwoInstance(backRepo, &sliceofpointertogongstructfieldDB)
		}
	}
}

// CopyBasicFieldsFromSliceOfPointerToGongStructField
func (sliceofpointertogongstructfieldDB *SliceOfPointerToGongStructFieldDB) CopyBasicFieldsFromSliceOfPointerToGongStructField(sliceofpointertogongstructfield *models.SliceOfPointerToGongStructField) {
	// insertion point for fields commit

	sliceofpointertogongstructfieldDB.Name_Data.String = sliceofpointertogongstructfield.Name
	sliceofpointertogongstructfieldDB.Name_Data.Valid = true

	sliceofpointertogongstructfieldDB.Index_Data.Int64 = int64(sliceofpointertogongstructfield.Index)
	sliceofpointertogongstructfieldDB.Index_Data.Valid = true

	sliceofpointertogongstructfieldDB.CompositeStructName_Data.String = sliceofpointertogongstructfield.CompositeStructName
	sliceofpointertogongstructfieldDB.CompositeStructName_Data.Valid = true
}

// CopyBasicFieldsFromSliceOfPointerToGongStructFieldWOP
func (sliceofpointertogongstructfieldDB *SliceOfPointerToGongStructFieldDB) CopyBasicFieldsFromSliceOfPointerToGongStructFieldWOP(sliceofpointertogongstructfield *SliceOfPointerToGongStructFieldWOP) {
	// insertion point for fields commit

	sliceofpointertogongstructfieldDB.Name_Data.String = sliceofpointertogongstructfield.Name
	sliceofpointertogongstructfieldDB.Name_Data.Valid = true

	sliceofpointertogongstructfieldDB.Index_Data.Int64 = int64(sliceofpointertogongstructfield.Index)
	sliceofpointertogongstructfieldDB.Index_Data.Valid = true

	sliceofpointertogongstructfieldDB.CompositeStructName_Data.String = sliceofpointertogongstructfield.CompositeStructName
	sliceofpointertogongstructfieldDB.CompositeStructName_Data.Valid = true
}

// CopyBasicFieldsToSliceOfPointerToGongStructField
func (sliceofpointertogongstructfieldDB *SliceOfPointerToGongStructFieldDB) CopyBasicFieldsToSliceOfPointerToGongStructField(sliceofpointertogongstructfield *models.SliceOfPointerToGongStructField) {
	// insertion point for checkout of basic fields (back repo to stage)
	sliceofpointertogongstructfield.Name = sliceofpointertogongstructfieldDB.Name_Data.String
	sliceofpointertogongstructfield.Index = int(sliceofpointertogongstructfieldDB.Index_Data.Int64)
	sliceofpointertogongstructfield.CompositeStructName = sliceofpointertogongstructfieldDB.CompositeStructName_Data.String
}

// CopyBasicFieldsToSliceOfPointerToGongStructFieldWOP
func (sliceofpointertogongstructfieldDB *SliceOfPointerToGongStructFieldDB) CopyBasicFieldsToSliceOfPointerToGongStructFieldWOP(sliceofpointertogongstructfield *SliceOfPointerToGongStructFieldWOP) {
	sliceofpointertogongstructfield.ID = int(sliceofpointertogongstructfieldDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	sliceofpointertogongstructfield.Name = sliceofpointertogongstructfieldDB.Name_Data.String
	sliceofpointertogongstructfield.Index = int(sliceofpointertogongstructfieldDB.Index_Data.Int64)
	sliceofpointertogongstructfield.CompositeStructName = sliceofpointertogongstructfieldDB.CompositeStructName_Data.String
}

// Backup generates a json file from a slice of all SliceOfPointerToGongStructFieldDB instances in the backrepo
func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SliceOfPointerToGongStructFieldDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SliceOfPointerToGongStructFieldDB, 0)
	for _, sliceofpointertogongstructfieldDB := range backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB {
		forBackup = append(forBackup, sliceofpointertogongstructfieldDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json SliceOfPointerToGongStructField ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json SliceOfPointerToGongStructField file", err.Error())
	}
}

// Backup generates a json file from a slice of all SliceOfPointerToGongStructFieldDB instances in the backrepo
func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SliceOfPointerToGongStructFieldDB, 0)
	for _, sliceofpointertogongstructfieldDB := range backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB {
		forBackup = append(forBackup, sliceofpointertogongstructfieldDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("SliceOfPointerToGongStructField")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&SliceOfPointerToGongStructField_Fields, -1)
	for _, sliceofpointertogongstructfieldDB := range forBackup {

		var sliceofpointertogongstructfieldWOP SliceOfPointerToGongStructFieldWOP
		sliceofpointertogongstructfieldDB.CopyBasicFieldsToSliceOfPointerToGongStructFieldWOP(&sliceofpointertogongstructfieldWOP)

		row := sh.AddRow()
		row.WriteStruct(&sliceofpointertogongstructfieldWOP, -1)
	}
}

// RestoreXL from the "SliceOfPointerToGongStructField" sheet all SliceOfPointerToGongStructFieldDB instances
func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSliceOfPointerToGongStructFieldid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["SliceOfPointerToGongStructField"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSliceOfPointerToGongStructField.rowVisitorSliceOfPointerToGongStructField)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) rowVisitorSliceOfPointerToGongStructField(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var sliceofpointertogongstructfieldWOP SliceOfPointerToGongStructFieldWOP
		row.ReadStruct(&sliceofpointertogongstructfieldWOP)

		// add the unmarshalled struct to the stage
		sliceofpointertogongstructfieldDB := new(SliceOfPointerToGongStructFieldDB)
		sliceofpointertogongstructfieldDB.CopyBasicFieldsFromSliceOfPointerToGongStructFieldWOP(&sliceofpointertogongstructfieldWOP)

		sliceofpointertogongstructfieldDB_ID_atBackupTime := sliceofpointertogongstructfieldDB.ID
		sliceofpointertogongstructfieldDB.ID = 0
		query := backRepoSliceOfPointerToGongStructField.db.Create(sliceofpointertogongstructfieldDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB[sliceofpointertogongstructfieldDB.ID] = sliceofpointertogongstructfieldDB
		BackRepoSliceOfPointerToGongStructFieldid_atBckpTime_newID[sliceofpointertogongstructfieldDB_ID_atBackupTime] = sliceofpointertogongstructfieldDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SliceOfPointerToGongStructFieldDB.json" in dirPath that stores an array
// of SliceOfPointerToGongStructFieldDB and stores it in the database
// the map BackRepoSliceOfPointerToGongStructFieldid_atBckpTime_newID is updated accordingly
func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSliceOfPointerToGongStructFieldid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SliceOfPointerToGongStructFieldDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json SliceOfPointerToGongStructField file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SliceOfPointerToGongStructFieldDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB
	for _, sliceofpointertogongstructfieldDB := range forRestore {

		sliceofpointertogongstructfieldDB_ID_atBackupTime := sliceofpointertogongstructfieldDB.ID
		sliceofpointertogongstructfieldDB.ID = 0
		query := backRepoSliceOfPointerToGongStructField.db.Create(sliceofpointertogongstructfieldDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB[sliceofpointertogongstructfieldDB.ID] = sliceofpointertogongstructfieldDB
		BackRepoSliceOfPointerToGongStructFieldid_atBckpTime_newID[sliceofpointertogongstructfieldDB_ID_atBackupTime] = sliceofpointertogongstructfieldDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json SliceOfPointerToGongStructField file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<SliceOfPointerToGongStructField>id_atBckpTime_newID
// to compute new index
func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) RestorePhaseTwo() {

	for _, sliceofpointertogongstructfieldDB := range backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB {

		// next line of code is to avert unused variable compilation error
		_ = sliceofpointertogongstructfieldDB

		// insertion point for reindexing pointers encoding
		// reindexing GongStruct field
		if sliceofpointertogongstructfieldDB.GongStructID.Int64 != 0 {
			sliceofpointertogongstructfieldDB.GongStructID.Int64 = int64(BackRepoGongStructid_atBckpTime_newID[uint(sliceofpointertogongstructfieldDB.GongStructID.Int64)])
			sliceofpointertogongstructfieldDB.GongStructID.Valid = true
		}

		// This reindex sliceofpointertogongstructfield.SliceOfPointerToGongStructFields
		if sliceofpointertogongstructfieldDB.GongStruct_SliceOfPointerToGongStructFieldsDBID.Int64 != 0 {
			sliceofpointertogongstructfieldDB.GongStruct_SliceOfPointerToGongStructFieldsDBID.Int64 =
				int64(BackRepoGongStructid_atBckpTime_newID[uint(sliceofpointertogongstructfieldDB.GongStruct_SliceOfPointerToGongStructFieldsDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoSliceOfPointerToGongStructField.db.Model(sliceofpointertogongstructfieldDB).Updates(*sliceofpointertogongstructfieldDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// BackRepoSliceOfPointerToGongStructField.ResetReversePointers commits all staged instances of SliceOfPointerToGongStructField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, sliceofpointertogongstructfield := range backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldPtr {
		backRepoSliceOfPointerToGongStructField.ResetReversePointersInstance(backRepo, idx, sliceofpointertogongstructfield)
	}

	return
}

func (backRepoSliceOfPointerToGongStructField *BackRepoSliceOfPointerToGongStructFieldStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, astruct *models.SliceOfPointerToGongStructField) (Error error) {

	// fetch matching sliceofpointertogongstructfieldDB
	if sliceofpointertogongstructfieldDB, ok := backRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldDB[idx]; ok {
		_ = sliceofpointertogongstructfieldDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		if sliceofpointertogongstructfieldDB.GongStruct_SliceOfPointerToGongStructFieldsDBID.Int64 != 0 {
			sliceofpointertogongstructfieldDB.GongStruct_SliceOfPointerToGongStructFieldsDBID.Int64 = 0
			sliceofpointertogongstructfieldDB.GongStruct_SliceOfPointerToGongStructFieldsDBID.Valid = true

			// save the reset
			if q := backRepoSliceOfPointerToGongStructField.db.Save(sliceofpointertogongstructfieldDB); q.Error != nil {
				return q.Error
			}
		}
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSliceOfPointerToGongStructFieldid_atBckpTime_newID map[uint]uint
