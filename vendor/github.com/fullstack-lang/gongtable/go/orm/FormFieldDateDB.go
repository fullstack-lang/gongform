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
var dummy_FormFieldDate_sql sql.NullBool
var dummy_FormFieldDate_time time.Duration
var dummy_FormFieldDate_sort sort.Float64Slice

// FormFieldDateAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model formfielddateAPI
type FormFieldDateAPI struct {
	gorm.Model

	models.FormFieldDate_WOP

	// encoding of pointers
	FormFieldDatePointersEncoding FormFieldDatePointersEncoding
}

// FormFieldDatePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type FormFieldDatePointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// FormFieldDateDB describes a formfielddate in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model formfielddateDB
type FormFieldDateDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field formfielddateDB.Name
	Name_Data sql.NullString

	// Declation for basic field formfielddateDB.Value
	Value_Data sql.NullTime
	// encoding of pointers
	FormFieldDatePointersEncoding
}

// FormFieldDateDBs arrays formfielddateDBs
// swagger:response formfielddateDBsResponse
type FormFieldDateDBs []FormFieldDateDB

// FormFieldDateDBResponse provides response
// swagger:response formfielddateDBResponse
type FormFieldDateDBResponse struct {
	FormFieldDateDB
}

// FormFieldDateWOP is a FormFieldDate without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type FormFieldDateWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Value time.Time `xlsx:"2"`
	// insertion for WOP pointer fields
}

var FormFieldDate_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Value",
}

type BackRepoFormFieldDateStruct struct {
	// stores FormFieldDateDB according to their gorm ID
	Map_FormFieldDateDBID_FormFieldDateDB map[uint]*FormFieldDateDB

	// stores FormFieldDateDB ID according to FormFieldDate address
	Map_FormFieldDatePtr_FormFieldDateDBID map[*models.FormFieldDate]uint

	// stores FormFieldDate according to their gorm ID
	Map_FormFieldDateDBID_FormFieldDatePtr map[uint]*models.FormFieldDate

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoFormFieldDate.stage
	return
}

func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) GetDB() *gorm.DB {
	return backRepoFormFieldDate.db
}

// GetFormFieldDateDBFromFormFieldDatePtr is a handy function to access the back repo instance from the stage instance
func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) GetFormFieldDateDBFromFormFieldDatePtr(formfielddate *models.FormFieldDate) (formfielddateDB *FormFieldDateDB) {
	id := backRepoFormFieldDate.Map_FormFieldDatePtr_FormFieldDateDBID[formfielddate]
	formfielddateDB = backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDateDB[id]
	return
}

// BackRepoFormFieldDate.CommitPhaseOne commits all staged instances of FormFieldDate to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for formfielddate := range stage.FormFieldDates {
		backRepoFormFieldDate.CommitPhaseOneInstance(formfielddate)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, formfielddate := range backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDatePtr {
		if _, ok := stage.FormFieldDates[formfielddate]; !ok {
			backRepoFormFieldDate.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoFormFieldDate.CommitDeleteInstance commits deletion of FormFieldDate to the BackRepo
func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) CommitDeleteInstance(id uint) (Error error) {

	formfielddate := backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDatePtr[id]

	// formfielddate is not staged anymore, remove formfielddateDB
	formfielddateDB := backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDateDB[id]
	query := backRepoFormFieldDate.db.Unscoped().Delete(&formfielddateDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoFormFieldDate.Map_FormFieldDatePtr_FormFieldDateDBID, formfielddate)
	delete(backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDatePtr, id)
	delete(backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDateDB, id)

	return
}

// BackRepoFormFieldDate.CommitPhaseOneInstance commits formfielddate staged instances of FormFieldDate to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) CommitPhaseOneInstance(formfielddate *models.FormFieldDate) (Error error) {

	// check if the formfielddate is not commited yet
	if _, ok := backRepoFormFieldDate.Map_FormFieldDatePtr_FormFieldDateDBID[formfielddate]; ok {
		return
	}

	// initiate formfielddate
	var formfielddateDB FormFieldDateDB
	formfielddateDB.CopyBasicFieldsFromFormFieldDate(formfielddate)

	query := backRepoFormFieldDate.db.Create(&formfielddateDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoFormFieldDate.Map_FormFieldDatePtr_FormFieldDateDBID[formfielddate] = formfielddateDB.ID
	backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDatePtr[formfielddateDB.ID] = formfielddate
	backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDateDB[formfielddateDB.ID] = &formfielddateDB

	return
}

// BackRepoFormFieldDate.CommitPhaseTwo commits all staged instances of FormFieldDate to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, formfielddate := range backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDatePtr {
		backRepoFormFieldDate.CommitPhaseTwoInstance(backRepo, idx, formfielddate)
	}

	return
}

// BackRepoFormFieldDate.CommitPhaseTwoInstance commits {{structname }} of models.FormFieldDate to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, formfielddate *models.FormFieldDate) (Error error) {

	// fetch matching formfielddateDB
	if formfielddateDB, ok := backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDateDB[idx]; ok {

		formfielddateDB.CopyBasicFieldsFromFormFieldDate(formfielddate)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoFormFieldDate.db.Save(&formfielddateDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown FormFieldDate intance %s", formfielddate.Name))
		return err
	}

	return
}

// BackRepoFormFieldDate.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) CheckoutPhaseOne() (Error error) {

	formfielddateDBArray := make([]FormFieldDateDB, 0)
	query := backRepoFormFieldDate.db.Find(&formfielddateDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	formfielddateInstancesToBeRemovedFromTheStage := make(map[*models.FormFieldDate]any)
	for key, value := range backRepoFormFieldDate.stage.FormFieldDates {
		formfielddateInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, formfielddateDB := range formfielddateDBArray {
		backRepoFormFieldDate.CheckoutPhaseOneInstance(&formfielddateDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		formfielddate, ok := backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDatePtr[formfielddateDB.ID]
		if ok {
			delete(formfielddateInstancesToBeRemovedFromTheStage, formfielddate)
		}
	}

	// remove from stage and back repo's 3 maps all formfielddates that are not in the checkout
	for formfielddate := range formfielddateInstancesToBeRemovedFromTheStage {
		formfielddate.Unstage(backRepoFormFieldDate.GetStage())

		// remove instance from the back repo 3 maps
		formfielddateID := backRepoFormFieldDate.Map_FormFieldDatePtr_FormFieldDateDBID[formfielddate]
		delete(backRepoFormFieldDate.Map_FormFieldDatePtr_FormFieldDateDBID, formfielddate)
		delete(backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDateDB, formfielddateID)
		delete(backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDatePtr, formfielddateID)
	}

	return
}

// CheckoutPhaseOneInstance takes a formfielddateDB that has been found in the DB, updates the backRepo and stages the
// models version of the formfielddateDB
func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) CheckoutPhaseOneInstance(formfielddateDB *FormFieldDateDB) (Error error) {

	formfielddate, ok := backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDatePtr[formfielddateDB.ID]
	if !ok {
		formfielddate = new(models.FormFieldDate)

		backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDatePtr[formfielddateDB.ID] = formfielddate
		backRepoFormFieldDate.Map_FormFieldDatePtr_FormFieldDateDBID[formfielddate] = formfielddateDB.ID

		// append model store with the new element
		formfielddate.Name = formfielddateDB.Name_Data.String
		formfielddate.Stage(backRepoFormFieldDate.GetStage())
	}
	formfielddateDB.CopyBasicFieldsToFormFieldDate(formfielddate)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	formfielddate.Stage(backRepoFormFieldDate.GetStage())

	// preserve pointer to formfielddateDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_FormFieldDateDBID_FormFieldDateDB)[formfielddateDB hold variable pointers
	formfielddateDB_Data := *formfielddateDB
	preservedPtrToFormFieldDate := &formfielddateDB_Data
	backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDateDB[formfielddateDB.ID] = preservedPtrToFormFieldDate

	return
}

// BackRepoFormFieldDate.CheckoutPhaseTwo Checkouts all staged instances of FormFieldDate to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, formfielddateDB := range backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDateDB {
		backRepoFormFieldDate.CheckoutPhaseTwoInstance(backRepo, formfielddateDB)
	}
	return
}

// BackRepoFormFieldDate.CheckoutPhaseTwoInstance Checkouts staged instances of FormFieldDate to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, formfielddateDB *FormFieldDateDB) (Error error) {

	formfielddate := backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDatePtr[formfielddateDB.ID]
	_ = formfielddate // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitFormFieldDate allows commit of a single formfielddate (if already staged)
func (backRepo *BackRepoStruct) CommitFormFieldDate(formfielddate *models.FormFieldDate) {
	backRepo.BackRepoFormFieldDate.CommitPhaseOneInstance(formfielddate)
	if id, ok := backRepo.BackRepoFormFieldDate.Map_FormFieldDatePtr_FormFieldDateDBID[formfielddate]; ok {
		backRepo.BackRepoFormFieldDate.CommitPhaseTwoInstance(backRepo, id, formfielddate)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitFormFieldDate allows checkout of a single formfielddate (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutFormFieldDate(formfielddate *models.FormFieldDate) {
	// check if the formfielddate is staged
	if _, ok := backRepo.BackRepoFormFieldDate.Map_FormFieldDatePtr_FormFieldDateDBID[formfielddate]; ok {

		if id, ok := backRepo.BackRepoFormFieldDate.Map_FormFieldDatePtr_FormFieldDateDBID[formfielddate]; ok {
			var formfielddateDB FormFieldDateDB
			formfielddateDB.ID = id

			if err := backRepo.BackRepoFormFieldDate.db.First(&formfielddateDB, id).Error; err != nil {
				log.Fatalln("CheckoutFormFieldDate : Problem with getting object with id:", id)
			}
			backRepo.BackRepoFormFieldDate.CheckoutPhaseOneInstance(&formfielddateDB)
			backRepo.BackRepoFormFieldDate.CheckoutPhaseTwoInstance(backRepo, &formfielddateDB)
		}
	}
}

// CopyBasicFieldsFromFormFieldDate
func (formfielddateDB *FormFieldDateDB) CopyBasicFieldsFromFormFieldDate(formfielddate *models.FormFieldDate) {
	// insertion point for fields commit

	formfielddateDB.Name_Data.String = formfielddate.Name
	formfielddateDB.Name_Data.Valid = true

	formfielddateDB.Value_Data.Time = formfielddate.Value
	formfielddateDB.Value_Data.Valid = true
}

// CopyBasicFieldsFromFormFieldDate_WOP
func (formfielddateDB *FormFieldDateDB) CopyBasicFieldsFromFormFieldDate_WOP(formfielddate *models.FormFieldDate_WOP) {
	// insertion point for fields commit

	formfielddateDB.Name_Data.String = formfielddate.Name
	formfielddateDB.Name_Data.Valid = true

	formfielddateDB.Value_Data.Time = formfielddate.Value
	formfielddateDB.Value_Data.Valid = true
}

// CopyBasicFieldsFromFormFieldDateWOP
func (formfielddateDB *FormFieldDateDB) CopyBasicFieldsFromFormFieldDateWOP(formfielddate *FormFieldDateWOP) {
	// insertion point for fields commit

	formfielddateDB.Name_Data.String = formfielddate.Name
	formfielddateDB.Name_Data.Valid = true

	formfielddateDB.Value_Data.Time = formfielddate.Value
	formfielddateDB.Value_Data.Valid = true
}

// CopyBasicFieldsToFormFieldDate
func (formfielddateDB *FormFieldDateDB) CopyBasicFieldsToFormFieldDate(formfielddate *models.FormFieldDate) {
	// insertion point for checkout of basic fields (back repo to stage)
	formfielddate.Name = formfielddateDB.Name_Data.String
	formfielddate.Value = formfielddateDB.Value_Data.Time
}

// CopyBasicFieldsToFormFieldDate_WOP
func (formfielddateDB *FormFieldDateDB) CopyBasicFieldsToFormFieldDate_WOP(formfielddate *models.FormFieldDate_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	formfielddate.Name = formfielddateDB.Name_Data.String
	formfielddate.Value = formfielddateDB.Value_Data.Time
}

// CopyBasicFieldsToFormFieldDateWOP
func (formfielddateDB *FormFieldDateDB) CopyBasicFieldsToFormFieldDateWOP(formfielddate *FormFieldDateWOP) {
	formfielddate.ID = int(formfielddateDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	formfielddate.Name = formfielddateDB.Name_Data.String
	formfielddate.Value = formfielddateDB.Value_Data.Time
}

// Backup generates a json file from a slice of all FormFieldDateDB instances in the backrepo
func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "FormFieldDateDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FormFieldDateDB, 0)
	for _, formfielddateDB := range backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDateDB {
		forBackup = append(forBackup, formfielddateDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json FormFieldDate ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json FormFieldDate file", err.Error())
	}
}

// Backup generates a json file from a slice of all FormFieldDateDB instances in the backrepo
func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FormFieldDateDB, 0)
	for _, formfielddateDB := range backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDateDB {
		forBackup = append(forBackup, formfielddateDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("FormFieldDate")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&FormFieldDate_Fields, -1)
	for _, formfielddateDB := range forBackup {

		var formfielddateWOP FormFieldDateWOP
		formfielddateDB.CopyBasicFieldsToFormFieldDateWOP(&formfielddateWOP)

		row := sh.AddRow()
		row.WriteStruct(&formfielddateWOP, -1)
	}
}

// RestoreXL from the "FormFieldDate" sheet all FormFieldDateDB instances
func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoFormFieldDateid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["FormFieldDate"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoFormFieldDate.rowVisitorFormFieldDate)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) rowVisitorFormFieldDate(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var formfielddateWOP FormFieldDateWOP
		row.ReadStruct(&formfielddateWOP)

		// add the unmarshalled struct to the stage
		formfielddateDB := new(FormFieldDateDB)
		formfielddateDB.CopyBasicFieldsFromFormFieldDateWOP(&formfielddateWOP)

		formfielddateDB_ID_atBackupTime := formfielddateDB.ID
		formfielddateDB.ID = 0
		query := backRepoFormFieldDate.db.Create(formfielddateDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDateDB[formfielddateDB.ID] = formfielddateDB
		BackRepoFormFieldDateid_atBckpTime_newID[formfielddateDB_ID_atBackupTime] = formfielddateDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "FormFieldDateDB.json" in dirPath that stores an array
// of FormFieldDateDB and stores it in the database
// the map BackRepoFormFieldDateid_atBckpTime_newID is updated accordingly
func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoFormFieldDateid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "FormFieldDateDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json FormFieldDate file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*FormFieldDateDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_FormFieldDateDBID_FormFieldDateDB
	for _, formfielddateDB := range forRestore {

		formfielddateDB_ID_atBackupTime := formfielddateDB.ID
		formfielddateDB.ID = 0
		query := backRepoFormFieldDate.db.Create(formfielddateDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDateDB[formfielddateDB.ID] = formfielddateDB
		BackRepoFormFieldDateid_atBckpTime_newID[formfielddateDB_ID_atBackupTime] = formfielddateDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json FormFieldDate file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<FormFieldDate>id_atBckpTime_newID
// to compute new index
func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) RestorePhaseTwo() {

	for _, formfielddateDB := range backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDateDB {

		// next line of code is to avert unused variable compilation error
		_ = formfielddateDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoFormFieldDate.db.Model(formfielddateDB).Updates(*formfielddateDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoFormFieldDate.ResetReversePointers commits all staged instances of FormFieldDate to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, formfielddate := range backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDatePtr {
		backRepoFormFieldDate.ResetReversePointersInstance(backRepo, idx, formfielddate)
	}

	return
}

func (backRepoFormFieldDate *BackRepoFormFieldDateStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, astruct *models.FormFieldDate) (Error error) {

	// fetch matching formfielddateDB
	if formfielddateDB, ok := backRepoFormFieldDate.Map_FormFieldDateDBID_FormFieldDateDB[idx]; ok {
		_ = formfielddateDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoFormFieldDateid_atBckpTime_newID map[uint]uint
