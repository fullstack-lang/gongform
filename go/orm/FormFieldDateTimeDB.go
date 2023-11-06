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

	"github.com/fullstack-lang/gongform/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_FormFieldDateTime_sql sql.NullBool
var dummy_FormFieldDateTime_time time.Duration
var dummy_FormFieldDateTime_sort sort.Float64Slice

// FormFieldDateTimeAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model formfielddatetimeAPI
type FormFieldDateTimeAPI struct {
	gorm.Model

	models.FormFieldDateTime_WOP

	// encoding of pointers
	FormFieldDateTimePointersEncoding FormFieldDateTimePointersEncoding
}

// FormFieldDateTimePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type FormFieldDateTimePointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// FormFieldDateTimeDB describes a formfielddatetime in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model formfielddatetimeDB
type FormFieldDateTimeDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field formfielddatetimeDB.Name
	Name_Data sql.NullString

	// Declation for basic field formfielddatetimeDB.Value
	Value_Data sql.NullTime
	// encoding of pointers
	FormFieldDateTimePointersEncoding
}

// FormFieldDateTimeDBs arrays formfielddatetimeDBs
// swagger:response formfielddatetimeDBsResponse
type FormFieldDateTimeDBs []FormFieldDateTimeDB

// FormFieldDateTimeDBResponse provides response
// swagger:response formfielddatetimeDBResponse
type FormFieldDateTimeDBResponse struct {
	FormFieldDateTimeDB
}

// FormFieldDateTimeWOP is a FormFieldDateTime without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type FormFieldDateTimeWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Value time.Time `xlsx:"2"`
	// insertion for WOP pointer fields
}

var FormFieldDateTime_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Value",
}

type BackRepoFormFieldDateTimeStruct struct {
	// stores FormFieldDateTimeDB according to their gorm ID
	Map_FormFieldDateTimeDBID_FormFieldDateTimeDB map[uint]*FormFieldDateTimeDB

	// stores FormFieldDateTimeDB ID according to FormFieldDateTime address
	Map_FormFieldDateTimePtr_FormFieldDateTimeDBID map[*models.FormFieldDateTime]uint

	// stores FormFieldDateTime according to their gorm ID
	Map_FormFieldDateTimeDBID_FormFieldDateTimePtr map[uint]*models.FormFieldDateTime

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoFormFieldDateTime.stage
	return
}

func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) GetDB() *gorm.DB {
	return backRepoFormFieldDateTime.db
}

// GetFormFieldDateTimeDBFromFormFieldDateTimePtr is a handy function to access the back repo instance from the stage instance
func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) GetFormFieldDateTimeDBFromFormFieldDateTimePtr(formfielddatetime *models.FormFieldDateTime) (formfielddatetimeDB *FormFieldDateTimeDB) {
	id := backRepoFormFieldDateTime.Map_FormFieldDateTimePtr_FormFieldDateTimeDBID[formfielddatetime]
	formfielddatetimeDB = backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimeDB[id]
	return
}

// BackRepoFormFieldDateTime.CommitPhaseOne commits all staged instances of FormFieldDateTime to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for formfielddatetime := range stage.FormFieldDateTimes {
		backRepoFormFieldDateTime.CommitPhaseOneInstance(formfielddatetime)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, formfielddatetime := range backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimePtr {
		if _, ok := stage.FormFieldDateTimes[formfielddatetime]; !ok {
			backRepoFormFieldDateTime.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoFormFieldDateTime.CommitDeleteInstance commits deletion of FormFieldDateTime to the BackRepo
func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) CommitDeleteInstance(id uint) (Error error) {

	formfielddatetime := backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimePtr[id]

	// formfielddatetime is not staged anymore, remove formfielddatetimeDB
	formfielddatetimeDB := backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimeDB[id]
	query := backRepoFormFieldDateTime.db.Unscoped().Delete(&formfielddatetimeDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoFormFieldDateTime.Map_FormFieldDateTimePtr_FormFieldDateTimeDBID, formfielddatetime)
	delete(backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimePtr, id)
	delete(backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimeDB, id)

	return
}

// BackRepoFormFieldDateTime.CommitPhaseOneInstance commits formfielddatetime staged instances of FormFieldDateTime to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) CommitPhaseOneInstance(formfielddatetime *models.FormFieldDateTime) (Error error) {

	// check if the formfielddatetime is not commited yet
	if _, ok := backRepoFormFieldDateTime.Map_FormFieldDateTimePtr_FormFieldDateTimeDBID[formfielddatetime]; ok {
		return
	}

	// initiate formfielddatetime
	var formfielddatetimeDB FormFieldDateTimeDB
	formfielddatetimeDB.CopyBasicFieldsFromFormFieldDateTime(formfielddatetime)

	query := backRepoFormFieldDateTime.db.Create(&formfielddatetimeDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoFormFieldDateTime.Map_FormFieldDateTimePtr_FormFieldDateTimeDBID[formfielddatetime] = formfielddatetimeDB.ID
	backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimePtr[formfielddatetimeDB.ID] = formfielddatetime
	backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimeDB[formfielddatetimeDB.ID] = &formfielddatetimeDB

	return
}

// BackRepoFormFieldDateTime.CommitPhaseTwo commits all staged instances of FormFieldDateTime to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, formfielddatetime := range backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimePtr {
		backRepoFormFieldDateTime.CommitPhaseTwoInstance(backRepo, idx, formfielddatetime)
	}

	return
}

// BackRepoFormFieldDateTime.CommitPhaseTwoInstance commits {{structname }} of models.FormFieldDateTime to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, formfielddatetime *models.FormFieldDateTime) (Error error) {

	// fetch matching formfielddatetimeDB
	if formfielddatetimeDB, ok := backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimeDB[idx]; ok {

		formfielddatetimeDB.CopyBasicFieldsFromFormFieldDateTime(formfielddatetime)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoFormFieldDateTime.db.Save(&formfielddatetimeDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown FormFieldDateTime intance %s", formfielddatetime.Name))
		return err
	}

	return
}

// BackRepoFormFieldDateTime.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) CheckoutPhaseOne() (Error error) {

	formfielddatetimeDBArray := make([]FormFieldDateTimeDB, 0)
	query := backRepoFormFieldDateTime.db.Find(&formfielddatetimeDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	formfielddatetimeInstancesToBeRemovedFromTheStage := make(map[*models.FormFieldDateTime]any)
	for key, value := range backRepoFormFieldDateTime.stage.FormFieldDateTimes {
		formfielddatetimeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, formfielddatetimeDB := range formfielddatetimeDBArray {
		backRepoFormFieldDateTime.CheckoutPhaseOneInstance(&formfielddatetimeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		formfielddatetime, ok := backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimePtr[formfielddatetimeDB.ID]
		if ok {
			delete(formfielddatetimeInstancesToBeRemovedFromTheStage, formfielddatetime)
		}
	}

	// remove from stage and back repo's 3 maps all formfielddatetimes that are not in the checkout
	for formfielddatetime := range formfielddatetimeInstancesToBeRemovedFromTheStage {
		formfielddatetime.Unstage(backRepoFormFieldDateTime.GetStage())

		// remove instance from the back repo 3 maps
		formfielddatetimeID := backRepoFormFieldDateTime.Map_FormFieldDateTimePtr_FormFieldDateTimeDBID[formfielddatetime]
		delete(backRepoFormFieldDateTime.Map_FormFieldDateTimePtr_FormFieldDateTimeDBID, formfielddatetime)
		delete(backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimeDB, formfielddatetimeID)
		delete(backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimePtr, formfielddatetimeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a formfielddatetimeDB that has been found in the DB, updates the backRepo and stages the
// models version of the formfielddatetimeDB
func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) CheckoutPhaseOneInstance(formfielddatetimeDB *FormFieldDateTimeDB) (Error error) {

	formfielddatetime, ok := backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimePtr[formfielddatetimeDB.ID]
	if !ok {
		formfielddatetime = new(models.FormFieldDateTime)

		backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimePtr[formfielddatetimeDB.ID] = formfielddatetime
		backRepoFormFieldDateTime.Map_FormFieldDateTimePtr_FormFieldDateTimeDBID[formfielddatetime] = formfielddatetimeDB.ID

		// append model store with the new element
		formfielddatetime.Name = formfielddatetimeDB.Name_Data.String
		formfielddatetime.Stage(backRepoFormFieldDateTime.GetStage())
	}
	formfielddatetimeDB.CopyBasicFieldsToFormFieldDateTime(formfielddatetime)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	formfielddatetime.Stage(backRepoFormFieldDateTime.GetStage())

	// preserve pointer to formfielddatetimeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_FormFieldDateTimeDBID_FormFieldDateTimeDB)[formfielddatetimeDB hold variable pointers
	formfielddatetimeDB_Data := *formfielddatetimeDB
	preservedPtrToFormFieldDateTime := &formfielddatetimeDB_Data
	backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimeDB[formfielddatetimeDB.ID] = preservedPtrToFormFieldDateTime

	return
}

// BackRepoFormFieldDateTime.CheckoutPhaseTwo Checkouts all staged instances of FormFieldDateTime to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, formfielddatetimeDB := range backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimeDB {
		backRepoFormFieldDateTime.CheckoutPhaseTwoInstance(backRepo, formfielddatetimeDB)
	}
	return
}

// BackRepoFormFieldDateTime.CheckoutPhaseTwoInstance Checkouts staged instances of FormFieldDateTime to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, formfielddatetimeDB *FormFieldDateTimeDB) (Error error) {

	formfielddatetime := backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimePtr[formfielddatetimeDB.ID]
	_ = formfielddatetime // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitFormFieldDateTime allows commit of a single formfielddatetime (if already staged)
func (backRepo *BackRepoStruct) CommitFormFieldDateTime(formfielddatetime *models.FormFieldDateTime) {
	backRepo.BackRepoFormFieldDateTime.CommitPhaseOneInstance(formfielddatetime)
	if id, ok := backRepo.BackRepoFormFieldDateTime.Map_FormFieldDateTimePtr_FormFieldDateTimeDBID[formfielddatetime]; ok {
		backRepo.BackRepoFormFieldDateTime.CommitPhaseTwoInstance(backRepo, id, formfielddatetime)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitFormFieldDateTime allows checkout of a single formfielddatetime (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutFormFieldDateTime(formfielddatetime *models.FormFieldDateTime) {
	// check if the formfielddatetime is staged
	if _, ok := backRepo.BackRepoFormFieldDateTime.Map_FormFieldDateTimePtr_FormFieldDateTimeDBID[formfielddatetime]; ok {

		if id, ok := backRepo.BackRepoFormFieldDateTime.Map_FormFieldDateTimePtr_FormFieldDateTimeDBID[formfielddatetime]; ok {
			var formfielddatetimeDB FormFieldDateTimeDB
			formfielddatetimeDB.ID = id

			if err := backRepo.BackRepoFormFieldDateTime.db.First(&formfielddatetimeDB, id).Error; err != nil {
				log.Fatalln("CheckoutFormFieldDateTime : Problem with getting object with id:", id)
			}
			backRepo.BackRepoFormFieldDateTime.CheckoutPhaseOneInstance(&formfielddatetimeDB)
			backRepo.BackRepoFormFieldDateTime.CheckoutPhaseTwoInstance(backRepo, &formfielddatetimeDB)
		}
	}
}

// CopyBasicFieldsFromFormFieldDateTime
func (formfielddatetimeDB *FormFieldDateTimeDB) CopyBasicFieldsFromFormFieldDateTime(formfielddatetime *models.FormFieldDateTime) {
	// insertion point for fields commit

	formfielddatetimeDB.Name_Data.String = formfielddatetime.Name
	formfielddatetimeDB.Name_Data.Valid = true

	formfielddatetimeDB.Value_Data.Time = formfielddatetime.Value
	formfielddatetimeDB.Value_Data.Valid = true
}

// CopyBasicFieldsFromFormFieldDateTime_WOP
func (formfielddatetimeDB *FormFieldDateTimeDB) CopyBasicFieldsFromFormFieldDateTime_WOP(formfielddatetime *models.FormFieldDateTime_WOP) {
	// insertion point for fields commit

	formfielddatetimeDB.Name_Data.String = formfielddatetime.Name
	formfielddatetimeDB.Name_Data.Valid = true

	formfielddatetimeDB.Value_Data.Time = formfielddatetime.Value
	formfielddatetimeDB.Value_Data.Valid = true
}

// CopyBasicFieldsFromFormFieldDateTimeWOP
func (formfielddatetimeDB *FormFieldDateTimeDB) CopyBasicFieldsFromFormFieldDateTimeWOP(formfielddatetime *FormFieldDateTimeWOP) {
	// insertion point for fields commit

	formfielddatetimeDB.Name_Data.String = formfielddatetime.Name
	formfielddatetimeDB.Name_Data.Valid = true

	formfielddatetimeDB.Value_Data.Time = formfielddatetime.Value
	formfielddatetimeDB.Value_Data.Valid = true
}

// CopyBasicFieldsToFormFieldDateTime
func (formfielddatetimeDB *FormFieldDateTimeDB) CopyBasicFieldsToFormFieldDateTime(formfielddatetime *models.FormFieldDateTime) {
	// insertion point for checkout of basic fields (back repo to stage)
	formfielddatetime.Name = formfielddatetimeDB.Name_Data.String
	formfielddatetime.Value = formfielddatetimeDB.Value_Data.Time
}

// CopyBasicFieldsToFormFieldDateTime_WOP
func (formfielddatetimeDB *FormFieldDateTimeDB) CopyBasicFieldsToFormFieldDateTime_WOP(formfielddatetime *models.FormFieldDateTime_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	formfielddatetime.Name = formfielddatetimeDB.Name_Data.String
	formfielddatetime.Value = formfielddatetimeDB.Value_Data.Time
}

// CopyBasicFieldsToFormFieldDateTimeWOP
func (formfielddatetimeDB *FormFieldDateTimeDB) CopyBasicFieldsToFormFieldDateTimeWOP(formfielddatetime *FormFieldDateTimeWOP) {
	formfielddatetime.ID = int(formfielddatetimeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	formfielddatetime.Name = formfielddatetimeDB.Name_Data.String
	formfielddatetime.Value = formfielddatetimeDB.Value_Data.Time
}

// Backup generates a json file from a slice of all FormFieldDateTimeDB instances in the backrepo
func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "FormFieldDateTimeDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FormFieldDateTimeDB, 0)
	for _, formfielddatetimeDB := range backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimeDB {
		forBackup = append(forBackup, formfielddatetimeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json FormFieldDateTime ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json FormFieldDateTime file", err.Error())
	}
}

// Backup generates a json file from a slice of all FormFieldDateTimeDB instances in the backrepo
func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FormFieldDateTimeDB, 0)
	for _, formfielddatetimeDB := range backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimeDB {
		forBackup = append(forBackup, formfielddatetimeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("FormFieldDateTime")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&FormFieldDateTime_Fields, -1)
	for _, formfielddatetimeDB := range forBackup {

		var formfielddatetimeWOP FormFieldDateTimeWOP
		formfielddatetimeDB.CopyBasicFieldsToFormFieldDateTimeWOP(&formfielddatetimeWOP)

		row := sh.AddRow()
		row.WriteStruct(&formfielddatetimeWOP, -1)
	}
}

// RestoreXL from the "FormFieldDateTime" sheet all FormFieldDateTimeDB instances
func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoFormFieldDateTimeid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["FormFieldDateTime"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoFormFieldDateTime.rowVisitorFormFieldDateTime)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) rowVisitorFormFieldDateTime(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var formfielddatetimeWOP FormFieldDateTimeWOP
		row.ReadStruct(&formfielddatetimeWOP)

		// add the unmarshalled struct to the stage
		formfielddatetimeDB := new(FormFieldDateTimeDB)
		formfielddatetimeDB.CopyBasicFieldsFromFormFieldDateTimeWOP(&formfielddatetimeWOP)

		formfielddatetimeDB_ID_atBackupTime := formfielddatetimeDB.ID
		formfielddatetimeDB.ID = 0
		query := backRepoFormFieldDateTime.db.Create(formfielddatetimeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimeDB[formfielddatetimeDB.ID] = formfielddatetimeDB
		BackRepoFormFieldDateTimeid_atBckpTime_newID[formfielddatetimeDB_ID_atBackupTime] = formfielddatetimeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "FormFieldDateTimeDB.json" in dirPath that stores an array
// of FormFieldDateTimeDB and stores it in the database
// the map BackRepoFormFieldDateTimeid_atBckpTime_newID is updated accordingly
func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoFormFieldDateTimeid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "FormFieldDateTimeDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json FormFieldDateTime file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*FormFieldDateTimeDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_FormFieldDateTimeDBID_FormFieldDateTimeDB
	for _, formfielddatetimeDB := range forRestore {

		formfielddatetimeDB_ID_atBackupTime := formfielddatetimeDB.ID
		formfielddatetimeDB.ID = 0
		query := backRepoFormFieldDateTime.db.Create(formfielddatetimeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimeDB[formfielddatetimeDB.ID] = formfielddatetimeDB
		BackRepoFormFieldDateTimeid_atBckpTime_newID[formfielddatetimeDB_ID_atBackupTime] = formfielddatetimeDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json FormFieldDateTime file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<FormFieldDateTime>id_atBckpTime_newID
// to compute new index
func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) RestorePhaseTwo() {

	for _, formfielddatetimeDB := range backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimeDB {

		// next line of code is to avert unused variable compilation error
		_ = formfielddatetimeDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoFormFieldDateTime.db.Model(formfielddatetimeDB).Updates(*formfielddatetimeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoFormFieldDateTime.ResetReversePointers commits all staged instances of FormFieldDateTime to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, formfielddatetime := range backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimePtr {
		backRepoFormFieldDateTime.ResetReversePointersInstance(backRepo, idx, formfielddatetime)
	}

	return
}

func (backRepoFormFieldDateTime *BackRepoFormFieldDateTimeStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, astruct *models.FormFieldDateTime) (Error error) {

	// fetch matching formfielddatetimeDB
	if formfielddatetimeDB, ok := backRepoFormFieldDateTime.Map_FormFieldDateTimeDBID_FormFieldDateTimeDB[idx]; ok {
		_ = formfielddatetimeDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoFormFieldDateTimeid_atBckpTime_newID map[uint]uint
