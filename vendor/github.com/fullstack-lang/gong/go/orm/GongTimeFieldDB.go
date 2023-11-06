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
var dummy_GongTimeField_sql sql.NullBool
var dummy_GongTimeField_time time.Duration
var dummy_GongTimeField_sort sort.Float64Slice

// GongTimeFieldAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model gongtimefieldAPI
type GongTimeFieldAPI struct {
	gorm.Model

	models.GongTimeField_WOP

	// encoding of pointers
	GongTimeFieldPointersEncoding GongTimeFieldPointersEncoding
}

// GongTimeFieldPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type GongTimeFieldPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// GongTimeFieldDB describes a gongtimefield in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model gongtimefieldDB
type GongTimeFieldDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field gongtimefieldDB.Name
	Name_Data sql.NullString

	// Declation for basic field gongtimefieldDB.Index
	Index_Data sql.NullInt64

	// Declation for basic field gongtimefieldDB.CompositeStructName
	CompositeStructName_Data sql.NullString
	// encoding of pointers
	GongTimeFieldPointersEncoding
}

// GongTimeFieldDBs arrays gongtimefieldDBs
// swagger:response gongtimefieldDBsResponse
type GongTimeFieldDBs []GongTimeFieldDB

// GongTimeFieldDBResponse provides response
// swagger:response gongtimefieldDBResponse
type GongTimeFieldDBResponse struct {
	GongTimeFieldDB
}

// GongTimeFieldWOP is a GongTimeField without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type GongTimeFieldWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Index int `xlsx:"2"`

	CompositeStructName string `xlsx:"3"`
	// insertion for WOP pointer fields
}

var GongTimeField_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Index",
	"CompositeStructName",
}

type BackRepoGongTimeFieldStruct struct {
	// stores GongTimeFieldDB according to their gorm ID
	Map_GongTimeFieldDBID_GongTimeFieldDB map[uint]*GongTimeFieldDB

	// stores GongTimeFieldDB ID according to GongTimeField address
	Map_GongTimeFieldPtr_GongTimeFieldDBID map[*models.GongTimeField]uint

	// stores GongTimeField according to their gorm ID
	Map_GongTimeFieldDBID_GongTimeFieldPtr map[uint]*models.GongTimeField

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoGongTimeField.stage
	return
}

func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) GetDB() *gorm.DB {
	return backRepoGongTimeField.db
}

// GetGongTimeFieldDBFromGongTimeFieldPtr is a handy function to access the back repo instance from the stage instance
func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) GetGongTimeFieldDBFromGongTimeFieldPtr(gongtimefield *models.GongTimeField) (gongtimefieldDB *GongTimeFieldDB) {
	id := backRepoGongTimeField.Map_GongTimeFieldPtr_GongTimeFieldDBID[gongtimefield]
	gongtimefieldDB = backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldDB[id]
	return
}

// BackRepoGongTimeField.CommitPhaseOne commits all staged instances of GongTimeField to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for gongtimefield := range stage.GongTimeFields {
		backRepoGongTimeField.CommitPhaseOneInstance(gongtimefield)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, gongtimefield := range backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldPtr {
		if _, ok := stage.GongTimeFields[gongtimefield]; !ok {
			backRepoGongTimeField.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoGongTimeField.CommitDeleteInstance commits deletion of GongTimeField to the BackRepo
func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) CommitDeleteInstance(id uint) (Error error) {

	gongtimefield := backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldPtr[id]

	// gongtimefield is not staged anymore, remove gongtimefieldDB
	gongtimefieldDB := backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldDB[id]
	query := backRepoGongTimeField.db.Unscoped().Delete(&gongtimefieldDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoGongTimeField.Map_GongTimeFieldPtr_GongTimeFieldDBID, gongtimefield)
	delete(backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldPtr, id)
	delete(backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldDB, id)

	return
}

// BackRepoGongTimeField.CommitPhaseOneInstance commits gongtimefield staged instances of GongTimeField to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) CommitPhaseOneInstance(gongtimefield *models.GongTimeField) (Error error) {

	// check if the gongtimefield is not commited yet
	if _, ok := backRepoGongTimeField.Map_GongTimeFieldPtr_GongTimeFieldDBID[gongtimefield]; ok {
		return
	}

	// initiate gongtimefield
	var gongtimefieldDB GongTimeFieldDB
	gongtimefieldDB.CopyBasicFieldsFromGongTimeField(gongtimefield)

	query := backRepoGongTimeField.db.Create(&gongtimefieldDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoGongTimeField.Map_GongTimeFieldPtr_GongTimeFieldDBID[gongtimefield] = gongtimefieldDB.ID
	backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldPtr[gongtimefieldDB.ID] = gongtimefield
	backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldDB[gongtimefieldDB.ID] = &gongtimefieldDB

	return
}

// BackRepoGongTimeField.CommitPhaseTwo commits all staged instances of GongTimeField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, gongtimefield := range backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldPtr {
		backRepoGongTimeField.CommitPhaseTwoInstance(backRepo, idx, gongtimefield)
	}

	return
}

// BackRepoGongTimeField.CommitPhaseTwoInstance commits {{structname }} of models.GongTimeField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, gongtimefield *models.GongTimeField) (Error error) {

	// fetch matching gongtimefieldDB
	if gongtimefieldDB, ok := backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldDB[idx]; ok {

		gongtimefieldDB.CopyBasicFieldsFromGongTimeField(gongtimefield)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoGongTimeField.db.Save(&gongtimefieldDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown GongTimeField intance %s", gongtimefield.Name))
		return err
	}

	return
}

// BackRepoGongTimeField.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) CheckoutPhaseOne() (Error error) {

	gongtimefieldDBArray := make([]GongTimeFieldDB, 0)
	query := backRepoGongTimeField.db.Find(&gongtimefieldDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	gongtimefieldInstancesToBeRemovedFromTheStage := make(map[*models.GongTimeField]any)
	for key, value := range backRepoGongTimeField.stage.GongTimeFields {
		gongtimefieldInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, gongtimefieldDB := range gongtimefieldDBArray {
		backRepoGongTimeField.CheckoutPhaseOneInstance(&gongtimefieldDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		gongtimefield, ok := backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldPtr[gongtimefieldDB.ID]
		if ok {
			delete(gongtimefieldInstancesToBeRemovedFromTheStage, gongtimefield)
		}
	}

	// remove from stage and back repo's 3 maps all gongtimefields that are not in the checkout
	for gongtimefield := range gongtimefieldInstancesToBeRemovedFromTheStage {
		gongtimefield.Unstage(backRepoGongTimeField.GetStage())

		// remove instance from the back repo 3 maps
		gongtimefieldID := backRepoGongTimeField.Map_GongTimeFieldPtr_GongTimeFieldDBID[gongtimefield]
		delete(backRepoGongTimeField.Map_GongTimeFieldPtr_GongTimeFieldDBID, gongtimefield)
		delete(backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldDB, gongtimefieldID)
		delete(backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldPtr, gongtimefieldID)
	}

	return
}

// CheckoutPhaseOneInstance takes a gongtimefieldDB that has been found in the DB, updates the backRepo and stages the
// models version of the gongtimefieldDB
func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) CheckoutPhaseOneInstance(gongtimefieldDB *GongTimeFieldDB) (Error error) {

	gongtimefield, ok := backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldPtr[gongtimefieldDB.ID]
	if !ok {
		gongtimefield = new(models.GongTimeField)

		backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldPtr[gongtimefieldDB.ID] = gongtimefield
		backRepoGongTimeField.Map_GongTimeFieldPtr_GongTimeFieldDBID[gongtimefield] = gongtimefieldDB.ID

		// append model store with the new element
		gongtimefield.Name = gongtimefieldDB.Name_Data.String
		gongtimefield.Stage(backRepoGongTimeField.GetStage())
	}
	gongtimefieldDB.CopyBasicFieldsToGongTimeField(gongtimefield)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	gongtimefield.Stage(backRepoGongTimeField.GetStage())

	// preserve pointer to gongtimefieldDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_GongTimeFieldDBID_GongTimeFieldDB)[gongtimefieldDB hold variable pointers
	gongtimefieldDB_Data := *gongtimefieldDB
	preservedPtrToGongTimeField := &gongtimefieldDB_Data
	backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldDB[gongtimefieldDB.ID] = preservedPtrToGongTimeField

	return
}

// BackRepoGongTimeField.CheckoutPhaseTwo Checkouts all staged instances of GongTimeField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, gongtimefieldDB := range backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldDB {
		backRepoGongTimeField.CheckoutPhaseTwoInstance(backRepo, gongtimefieldDB)
	}
	return
}

// BackRepoGongTimeField.CheckoutPhaseTwoInstance Checkouts staged instances of GongTimeField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, gongtimefieldDB *GongTimeFieldDB) (Error error) {

	gongtimefield := backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldPtr[gongtimefieldDB.ID]
	_ = gongtimefield // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitGongTimeField allows commit of a single gongtimefield (if already staged)
func (backRepo *BackRepoStruct) CommitGongTimeField(gongtimefield *models.GongTimeField) {
	backRepo.BackRepoGongTimeField.CommitPhaseOneInstance(gongtimefield)
	if id, ok := backRepo.BackRepoGongTimeField.Map_GongTimeFieldPtr_GongTimeFieldDBID[gongtimefield]; ok {
		backRepo.BackRepoGongTimeField.CommitPhaseTwoInstance(backRepo, id, gongtimefield)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitGongTimeField allows checkout of a single gongtimefield (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutGongTimeField(gongtimefield *models.GongTimeField) {
	// check if the gongtimefield is staged
	if _, ok := backRepo.BackRepoGongTimeField.Map_GongTimeFieldPtr_GongTimeFieldDBID[gongtimefield]; ok {

		if id, ok := backRepo.BackRepoGongTimeField.Map_GongTimeFieldPtr_GongTimeFieldDBID[gongtimefield]; ok {
			var gongtimefieldDB GongTimeFieldDB
			gongtimefieldDB.ID = id

			if err := backRepo.BackRepoGongTimeField.db.First(&gongtimefieldDB, id).Error; err != nil {
				log.Fatalln("CheckoutGongTimeField : Problem with getting object with id:", id)
			}
			backRepo.BackRepoGongTimeField.CheckoutPhaseOneInstance(&gongtimefieldDB)
			backRepo.BackRepoGongTimeField.CheckoutPhaseTwoInstance(backRepo, &gongtimefieldDB)
		}
	}
}

// CopyBasicFieldsFromGongTimeField
func (gongtimefieldDB *GongTimeFieldDB) CopyBasicFieldsFromGongTimeField(gongtimefield *models.GongTimeField) {
	// insertion point for fields commit

	gongtimefieldDB.Name_Data.String = gongtimefield.Name
	gongtimefieldDB.Name_Data.Valid = true

	gongtimefieldDB.Index_Data.Int64 = int64(gongtimefield.Index)
	gongtimefieldDB.Index_Data.Valid = true

	gongtimefieldDB.CompositeStructName_Data.String = gongtimefield.CompositeStructName
	gongtimefieldDB.CompositeStructName_Data.Valid = true
}

// CopyBasicFieldsFromGongTimeField_WOP
func (gongtimefieldDB *GongTimeFieldDB) CopyBasicFieldsFromGongTimeField_WOP(gongtimefield *models.GongTimeField_WOP) {
	// insertion point for fields commit

	gongtimefieldDB.Name_Data.String = gongtimefield.Name
	gongtimefieldDB.Name_Data.Valid = true

	gongtimefieldDB.Index_Data.Int64 = int64(gongtimefield.Index)
	gongtimefieldDB.Index_Data.Valid = true

	gongtimefieldDB.CompositeStructName_Data.String = gongtimefield.CompositeStructName
	gongtimefieldDB.CompositeStructName_Data.Valid = true
}

// CopyBasicFieldsFromGongTimeFieldWOP
func (gongtimefieldDB *GongTimeFieldDB) CopyBasicFieldsFromGongTimeFieldWOP(gongtimefield *GongTimeFieldWOP) {
	// insertion point for fields commit

	gongtimefieldDB.Name_Data.String = gongtimefield.Name
	gongtimefieldDB.Name_Data.Valid = true

	gongtimefieldDB.Index_Data.Int64 = int64(gongtimefield.Index)
	gongtimefieldDB.Index_Data.Valid = true

	gongtimefieldDB.CompositeStructName_Data.String = gongtimefield.CompositeStructName
	gongtimefieldDB.CompositeStructName_Data.Valid = true
}

// CopyBasicFieldsToGongTimeField
func (gongtimefieldDB *GongTimeFieldDB) CopyBasicFieldsToGongTimeField(gongtimefield *models.GongTimeField) {
	// insertion point for checkout of basic fields (back repo to stage)
	gongtimefield.Name = gongtimefieldDB.Name_Data.String
	gongtimefield.Index = int(gongtimefieldDB.Index_Data.Int64)
	gongtimefield.CompositeStructName = gongtimefieldDB.CompositeStructName_Data.String
}

// CopyBasicFieldsToGongTimeField_WOP
func (gongtimefieldDB *GongTimeFieldDB) CopyBasicFieldsToGongTimeField_WOP(gongtimefield *models.GongTimeField_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	gongtimefield.Name = gongtimefieldDB.Name_Data.String
	gongtimefield.Index = int(gongtimefieldDB.Index_Data.Int64)
	gongtimefield.CompositeStructName = gongtimefieldDB.CompositeStructName_Data.String
}

// CopyBasicFieldsToGongTimeFieldWOP
func (gongtimefieldDB *GongTimeFieldDB) CopyBasicFieldsToGongTimeFieldWOP(gongtimefield *GongTimeFieldWOP) {
	gongtimefield.ID = int(gongtimefieldDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	gongtimefield.Name = gongtimefieldDB.Name_Data.String
	gongtimefield.Index = int(gongtimefieldDB.Index_Data.Int64)
	gongtimefield.CompositeStructName = gongtimefieldDB.CompositeStructName_Data.String
}

// Backup generates a json file from a slice of all GongTimeFieldDB instances in the backrepo
func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "GongTimeFieldDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongTimeFieldDB, 0)
	for _, gongtimefieldDB := range backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldDB {
		forBackup = append(forBackup, gongtimefieldDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json GongTimeField ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json GongTimeField file", err.Error())
	}
}

// Backup generates a json file from a slice of all GongTimeFieldDB instances in the backrepo
func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongTimeFieldDB, 0)
	for _, gongtimefieldDB := range backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldDB {
		forBackup = append(forBackup, gongtimefieldDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("GongTimeField")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&GongTimeField_Fields, -1)
	for _, gongtimefieldDB := range forBackup {

		var gongtimefieldWOP GongTimeFieldWOP
		gongtimefieldDB.CopyBasicFieldsToGongTimeFieldWOP(&gongtimefieldWOP)

		row := sh.AddRow()
		row.WriteStruct(&gongtimefieldWOP, -1)
	}
}

// RestoreXL from the "GongTimeField" sheet all GongTimeFieldDB instances
func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoGongTimeFieldid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["GongTimeField"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoGongTimeField.rowVisitorGongTimeField)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) rowVisitorGongTimeField(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var gongtimefieldWOP GongTimeFieldWOP
		row.ReadStruct(&gongtimefieldWOP)

		// add the unmarshalled struct to the stage
		gongtimefieldDB := new(GongTimeFieldDB)
		gongtimefieldDB.CopyBasicFieldsFromGongTimeFieldWOP(&gongtimefieldWOP)

		gongtimefieldDB_ID_atBackupTime := gongtimefieldDB.ID
		gongtimefieldDB.ID = 0
		query := backRepoGongTimeField.db.Create(gongtimefieldDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldDB[gongtimefieldDB.ID] = gongtimefieldDB
		BackRepoGongTimeFieldid_atBckpTime_newID[gongtimefieldDB_ID_atBackupTime] = gongtimefieldDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "GongTimeFieldDB.json" in dirPath that stores an array
// of GongTimeFieldDB and stores it in the database
// the map BackRepoGongTimeFieldid_atBckpTime_newID is updated accordingly
func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoGongTimeFieldid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "GongTimeFieldDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json GongTimeField file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*GongTimeFieldDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_GongTimeFieldDBID_GongTimeFieldDB
	for _, gongtimefieldDB := range forRestore {

		gongtimefieldDB_ID_atBackupTime := gongtimefieldDB.ID
		gongtimefieldDB.ID = 0
		query := backRepoGongTimeField.db.Create(gongtimefieldDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldDB[gongtimefieldDB.ID] = gongtimefieldDB
		BackRepoGongTimeFieldid_atBckpTime_newID[gongtimefieldDB_ID_atBackupTime] = gongtimefieldDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json GongTimeField file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<GongTimeField>id_atBckpTime_newID
// to compute new index
func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) RestorePhaseTwo() {

	for _, gongtimefieldDB := range backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldDB {

		// next line of code is to avert unused variable compilation error
		_ = gongtimefieldDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoGongTimeField.db.Model(gongtimefieldDB).Updates(*gongtimefieldDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoGongTimeField.ResetReversePointers commits all staged instances of GongTimeField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, gongtimefield := range backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldPtr {
		backRepoGongTimeField.ResetReversePointersInstance(backRepo, idx, gongtimefield)
	}

	return
}

func (backRepoGongTimeField *BackRepoGongTimeFieldStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, astruct *models.GongTimeField) (Error error) {

	// fetch matching gongtimefieldDB
	if gongtimefieldDB, ok := backRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldDB[idx]; ok {
		_ = gongtimefieldDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoGongTimeFieldid_atBckpTime_newID map[uint]uint
