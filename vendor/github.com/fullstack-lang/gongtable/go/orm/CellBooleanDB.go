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

	"github.com/fullstack-lang/gongtable/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_CellBoolean_sql sql.NullBool
var dummy_CellBoolean_time time.Duration
var dummy_CellBoolean_sort sort.Float64Slice

// CellBooleanAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model cellbooleanAPI
type CellBooleanAPI struct {
	gorm.Model

	models.CellBoolean

	// encoding of pointers
	CellBooleanPointersEnconding
}

// CellBooleanPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type CellBooleanPointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// CellBooleanDB describes a cellboolean in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model cellbooleanDB
type CellBooleanDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field cellbooleanDB.Name
	Name_Data sql.NullString

	// Declation for basic field cellbooleanDB.Value
	// provide the sql storage for the boolan
	Value_Data sql.NullBool
	// encoding of pointers
	CellBooleanPointersEnconding
}

// CellBooleanDBs arrays cellbooleanDBs
// swagger:response cellbooleanDBsResponse
type CellBooleanDBs []CellBooleanDB

// CellBooleanDBResponse provides response
// swagger:response cellbooleanDBResponse
type CellBooleanDBResponse struct {
	CellBooleanDB
}

// CellBooleanWOP is a CellBoolean without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type CellBooleanWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Value bool `xlsx:"2"`
	// insertion for WOP pointer fields
}

var CellBoolean_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Value",
}

type BackRepoCellBooleanStruct struct {
	// stores CellBooleanDB according to their gorm ID
	Map_CellBooleanDBID_CellBooleanDB map[uint]*CellBooleanDB

	// stores CellBooleanDB ID according to CellBoolean address
	Map_CellBooleanPtr_CellBooleanDBID map[*models.CellBoolean]uint

	// stores CellBoolean according to their gorm ID
	Map_CellBooleanDBID_CellBooleanPtr map[uint]*models.CellBoolean

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoCellBoolean *BackRepoCellBooleanStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoCellBoolean.stage
	return
}

func (backRepoCellBoolean *BackRepoCellBooleanStruct) GetDB() *gorm.DB {
	return backRepoCellBoolean.db
}

// GetCellBooleanDBFromCellBooleanPtr is a handy function to access the back repo instance from the stage instance
func (backRepoCellBoolean *BackRepoCellBooleanStruct) GetCellBooleanDBFromCellBooleanPtr(cellboolean *models.CellBoolean) (cellbooleanDB *CellBooleanDB) {
	id := backRepoCellBoolean.Map_CellBooleanPtr_CellBooleanDBID[cellboolean]
	cellbooleanDB = backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanDB[id]
	return
}

// BackRepoCellBoolean.CommitPhaseOne commits all staged instances of CellBoolean to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoCellBoolean *BackRepoCellBooleanStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for cellboolean := range stage.CellBooleans {
		backRepoCellBoolean.CommitPhaseOneInstance(cellboolean)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, cellboolean := range backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanPtr {
		if _, ok := stage.CellBooleans[cellboolean]; !ok {
			backRepoCellBoolean.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoCellBoolean.CommitDeleteInstance commits deletion of CellBoolean to the BackRepo
func (backRepoCellBoolean *BackRepoCellBooleanStruct) CommitDeleteInstance(id uint) (Error error) {

	cellboolean := backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanPtr[id]

	// cellboolean is not staged anymore, remove cellbooleanDB
	cellbooleanDB := backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanDB[id]
	query := backRepoCellBoolean.db.Unscoped().Delete(&cellbooleanDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoCellBoolean.Map_CellBooleanPtr_CellBooleanDBID, cellboolean)
	delete(backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanPtr, id)
	delete(backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanDB, id)

	return
}

// BackRepoCellBoolean.CommitPhaseOneInstance commits cellboolean staged instances of CellBoolean to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoCellBoolean *BackRepoCellBooleanStruct) CommitPhaseOneInstance(cellboolean *models.CellBoolean) (Error error) {

	// check if the cellboolean is not commited yet
	if _, ok := backRepoCellBoolean.Map_CellBooleanPtr_CellBooleanDBID[cellboolean]; ok {
		return
	}

	// initiate cellboolean
	var cellbooleanDB CellBooleanDB
	cellbooleanDB.CopyBasicFieldsFromCellBoolean(cellboolean)

	query := backRepoCellBoolean.db.Create(&cellbooleanDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoCellBoolean.Map_CellBooleanPtr_CellBooleanDBID[cellboolean] = cellbooleanDB.ID
	backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanPtr[cellbooleanDB.ID] = cellboolean
	backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanDB[cellbooleanDB.ID] = &cellbooleanDB

	return
}

// BackRepoCellBoolean.CommitPhaseTwo commits all staged instances of CellBoolean to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCellBoolean *BackRepoCellBooleanStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, cellboolean := range backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanPtr {
		backRepoCellBoolean.CommitPhaseTwoInstance(backRepo, idx, cellboolean)
	}

	return
}

// BackRepoCellBoolean.CommitPhaseTwoInstance commits {{structname }} of models.CellBoolean to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCellBoolean *BackRepoCellBooleanStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, cellboolean *models.CellBoolean) (Error error) {

	// fetch matching cellbooleanDB
	if cellbooleanDB, ok := backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanDB[idx]; ok {

		cellbooleanDB.CopyBasicFieldsFromCellBoolean(cellboolean)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoCellBoolean.db.Save(&cellbooleanDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown CellBoolean intance %s", cellboolean.Name))
		return err
	}

	return
}

// BackRepoCellBoolean.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoCellBoolean *BackRepoCellBooleanStruct) CheckoutPhaseOne() (Error error) {

	cellbooleanDBArray := make([]CellBooleanDB, 0)
	query := backRepoCellBoolean.db.Find(&cellbooleanDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	cellbooleanInstancesToBeRemovedFromTheStage := make(map[*models.CellBoolean]any)
	for key, value := range backRepoCellBoolean.stage.CellBooleans {
		cellbooleanInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, cellbooleanDB := range cellbooleanDBArray {
		backRepoCellBoolean.CheckoutPhaseOneInstance(&cellbooleanDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		cellboolean, ok := backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanPtr[cellbooleanDB.ID]
		if ok {
			delete(cellbooleanInstancesToBeRemovedFromTheStage, cellboolean)
		}
	}

	// remove from stage and back repo's 3 maps all cellbooleans that are not in the checkout
	for cellboolean := range cellbooleanInstancesToBeRemovedFromTheStage {
		cellboolean.Unstage(backRepoCellBoolean.GetStage())

		// remove instance from the back repo 3 maps
		cellbooleanID := backRepoCellBoolean.Map_CellBooleanPtr_CellBooleanDBID[cellboolean]
		delete(backRepoCellBoolean.Map_CellBooleanPtr_CellBooleanDBID, cellboolean)
		delete(backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanDB, cellbooleanID)
		delete(backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanPtr, cellbooleanID)
	}

	return
}

// CheckoutPhaseOneInstance takes a cellbooleanDB that has been found in the DB, updates the backRepo and stages the
// models version of the cellbooleanDB
func (backRepoCellBoolean *BackRepoCellBooleanStruct) CheckoutPhaseOneInstance(cellbooleanDB *CellBooleanDB) (Error error) {

	cellboolean, ok := backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanPtr[cellbooleanDB.ID]
	if !ok {
		cellboolean = new(models.CellBoolean)

		backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanPtr[cellbooleanDB.ID] = cellboolean
		backRepoCellBoolean.Map_CellBooleanPtr_CellBooleanDBID[cellboolean] = cellbooleanDB.ID

		// append model store with the new element
		cellboolean.Name = cellbooleanDB.Name_Data.String
		cellboolean.Stage(backRepoCellBoolean.GetStage())
	}
	cellbooleanDB.CopyBasicFieldsToCellBoolean(cellboolean)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	cellboolean.Stage(backRepoCellBoolean.GetStage())

	// preserve pointer to cellbooleanDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_CellBooleanDBID_CellBooleanDB)[cellbooleanDB hold variable pointers
	cellbooleanDB_Data := *cellbooleanDB
	preservedPtrToCellBoolean := &cellbooleanDB_Data
	backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanDB[cellbooleanDB.ID] = preservedPtrToCellBoolean

	return
}

// BackRepoCellBoolean.CheckoutPhaseTwo Checkouts all staged instances of CellBoolean to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCellBoolean *BackRepoCellBooleanStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, cellbooleanDB := range backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanDB {
		backRepoCellBoolean.CheckoutPhaseTwoInstance(backRepo, cellbooleanDB)
	}
	return
}

// BackRepoCellBoolean.CheckoutPhaseTwoInstance Checkouts staged instances of CellBoolean to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCellBoolean *BackRepoCellBooleanStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, cellbooleanDB *CellBooleanDB) (Error error) {

	cellboolean := backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanPtr[cellbooleanDB.ID]
	_ = cellboolean // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitCellBoolean allows commit of a single cellboolean (if already staged)
func (backRepo *BackRepoStruct) CommitCellBoolean(cellboolean *models.CellBoolean) {
	backRepo.BackRepoCellBoolean.CommitPhaseOneInstance(cellboolean)
	if id, ok := backRepo.BackRepoCellBoolean.Map_CellBooleanPtr_CellBooleanDBID[cellboolean]; ok {
		backRepo.BackRepoCellBoolean.CommitPhaseTwoInstance(backRepo, id, cellboolean)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitCellBoolean allows checkout of a single cellboolean (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutCellBoolean(cellboolean *models.CellBoolean) {
	// check if the cellboolean is staged
	if _, ok := backRepo.BackRepoCellBoolean.Map_CellBooleanPtr_CellBooleanDBID[cellboolean]; ok {

		if id, ok := backRepo.BackRepoCellBoolean.Map_CellBooleanPtr_CellBooleanDBID[cellboolean]; ok {
			var cellbooleanDB CellBooleanDB
			cellbooleanDB.ID = id

			if err := backRepo.BackRepoCellBoolean.db.First(&cellbooleanDB, id).Error; err != nil {
				log.Panicln("CheckoutCellBoolean : Problem with getting object with id:", id)
			}
			backRepo.BackRepoCellBoolean.CheckoutPhaseOneInstance(&cellbooleanDB)
			backRepo.BackRepoCellBoolean.CheckoutPhaseTwoInstance(backRepo, &cellbooleanDB)
		}
	}
}

// CopyBasicFieldsFromCellBoolean
func (cellbooleanDB *CellBooleanDB) CopyBasicFieldsFromCellBoolean(cellboolean *models.CellBoolean) {
	// insertion point for fields commit

	cellbooleanDB.Name_Data.String = cellboolean.Name
	cellbooleanDB.Name_Data.Valid = true

	cellbooleanDB.Value_Data.Bool = cellboolean.Value
	cellbooleanDB.Value_Data.Valid = true
}

// CopyBasicFieldsFromCellBooleanWOP
func (cellbooleanDB *CellBooleanDB) CopyBasicFieldsFromCellBooleanWOP(cellboolean *CellBooleanWOP) {
	// insertion point for fields commit

	cellbooleanDB.Name_Data.String = cellboolean.Name
	cellbooleanDB.Name_Data.Valid = true

	cellbooleanDB.Value_Data.Bool = cellboolean.Value
	cellbooleanDB.Value_Data.Valid = true
}

// CopyBasicFieldsToCellBoolean
func (cellbooleanDB *CellBooleanDB) CopyBasicFieldsToCellBoolean(cellboolean *models.CellBoolean) {
	// insertion point for checkout of basic fields (back repo to stage)
	cellboolean.Name = cellbooleanDB.Name_Data.String
	cellboolean.Value = cellbooleanDB.Value_Data.Bool
}

// CopyBasicFieldsToCellBooleanWOP
func (cellbooleanDB *CellBooleanDB) CopyBasicFieldsToCellBooleanWOP(cellboolean *CellBooleanWOP) {
	cellboolean.ID = int(cellbooleanDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	cellboolean.Name = cellbooleanDB.Name_Data.String
	cellboolean.Value = cellbooleanDB.Value_Data.Bool
}

// Backup generates a json file from a slice of all CellBooleanDB instances in the backrepo
func (backRepoCellBoolean *BackRepoCellBooleanStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "CellBooleanDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*CellBooleanDB, 0)
	for _, cellbooleanDB := range backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanDB {
		forBackup = append(forBackup, cellbooleanDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json CellBoolean ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json CellBoolean file", err.Error())
	}
}

// Backup generates a json file from a slice of all CellBooleanDB instances in the backrepo
func (backRepoCellBoolean *BackRepoCellBooleanStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*CellBooleanDB, 0)
	for _, cellbooleanDB := range backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanDB {
		forBackup = append(forBackup, cellbooleanDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("CellBoolean")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&CellBoolean_Fields, -1)
	for _, cellbooleanDB := range forBackup {

		var cellbooleanWOP CellBooleanWOP
		cellbooleanDB.CopyBasicFieldsToCellBooleanWOP(&cellbooleanWOP)

		row := sh.AddRow()
		row.WriteStruct(&cellbooleanWOP, -1)
	}
}

// RestoreXL from the "CellBoolean" sheet all CellBooleanDB instances
func (backRepoCellBoolean *BackRepoCellBooleanStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoCellBooleanid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["CellBoolean"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoCellBoolean.rowVisitorCellBoolean)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoCellBoolean *BackRepoCellBooleanStruct) rowVisitorCellBoolean(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var cellbooleanWOP CellBooleanWOP
		row.ReadStruct(&cellbooleanWOP)

		// add the unmarshalled struct to the stage
		cellbooleanDB := new(CellBooleanDB)
		cellbooleanDB.CopyBasicFieldsFromCellBooleanWOP(&cellbooleanWOP)

		cellbooleanDB_ID_atBackupTime := cellbooleanDB.ID
		cellbooleanDB.ID = 0
		query := backRepoCellBoolean.db.Create(cellbooleanDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanDB[cellbooleanDB.ID] = cellbooleanDB
		BackRepoCellBooleanid_atBckpTime_newID[cellbooleanDB_ID_atBackupTime] = cellbooleanDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "CellBooleanDB.json" in dirPath that stores an array
// of CellBooleanDB and stores it in the database
// the map BackRepoCellBooleanid_atBckpTime_newID is updated accordingly
func (backRepoCellBoolean *BackRepoCellBooleanStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoCellBooleanid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "CellBooleanDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json CellBoolean file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*CellBooleanDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_CellBooleanDBID_CellBooleanDB
	for _, cellbooleanDB := range forRestore {

		cellbooleanDB_ID_atBackupTime := cellbooleanDB.ID
		cellbooleanDB.ID = 0
		query := backRepoCellBoolean.db.Create(cellbooleanDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanDB[cellbooleanDB.ID] = cellbooleanDB
		BackRepoCellBooleanid_atBckpTime_newID[cellbooleanDB_ID_atBackupTime] = cellbooleanDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json CellBoolean file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<CellBoolean>id_atBckpTime_newID
// to compute new index
func (backRepoCellBoolean *BackRepoCellBooleanStruct) RestorePhaseTwo() {

	for _, cellbooleanDB := range backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanDB {

		// next line of code is to avert unused variable compilation error
		_ = cellbooleanDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoCellBoolean.db.Model(cellbooleanDB).Updates(*cellbooleanDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// BackRepoCellBoolean.ResetReversePointers commits all staged instances of CellBoolean to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCellBoolean *BackRepoCellBooleanStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, cellboolean := range backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanPtr {
		backRepoCellBoolean.ResetReversePointersInstance(backRepo, idx, cellboolean)
	}

	return
}

func (backRepoCellBoolean *BackRepoCellBooleanStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, astruct *models.CellBoolean) (Error error) {

	// fetch matching cellbooleanDB
	if cellbooleanDB, ok := backRepoCellBoolean.Map_CellBooleanDBID_CellBooleanDB[idx]; ok {
		_ = cellbooleanDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoCellBooleanid_atBckpTime_newID map[uint]uint
