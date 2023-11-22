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
var dummy_FormFieldString_sql sql.NullBool
var dummy_FormFieldString_time time.Duration
var dummy_FormFieldString_sort sort.Float64Slice

// FormFieldStringAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model formfieldstringAPI
type FormFieldStringAPI struct {
	gorm.Model

	models.FormFieldString_WOP

	// encoding of pointers
	FormFieldStringPointersEncoding FormFieldStringPointersEncoding
}

// FormFieldStringPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type FormFieldStringPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// FormFieldStringDB describes a formfieldstring in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model formfieldstringDB
type FormFieldStringDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field formfieldstringDB.Name
	Name_Data sql.NullString

	// Declation for basic field formfieldstringDB.Value
	Value_Data sql.NullString

	// Declation for basic field formfieldstringDB.IsTextArea
	// provide the sql storage for the boolan
	IsTextArea_Data sql.NullBool
	// encoding of pointers
	FormFieldStringPointersEncoding
}

// FormFieldStringDBs arrays formfieldstringDBs
// swagger:response formfieldstringDBsResponse
type FormFieldStringDBs []FormFieldStringDB

// FormFieldStringDBResponse provides response
// swagger:response formfieldstringDBResponse
type FormFieldStringDBResponse struct {
	FormFieldStringDB
}

// FormFieldStringWOP is a FormFieldString without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type FormFieldStringWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Value string `xlsx:"2"`

	IsTextArea bool `xlsx:"3"`
	// insertion for WOP pointer fields
}

var FormFieldString_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Value",
	"IsTextArea",
}

type BackRepoFormFieldStringStruct struct {
	// stores FormFieldStringDB according to their gorm ID
	Map_FormFieldStringDBID_FormFieldStringDB map[uint]*FormFieldStringDB

	// stores FormFieldStringDB ID according to FormFieldString address
	Map_FormFieldStringPtr_FormFieldStringDBID map[*models.FormFieldString]uint

	// stores FormFieldString according to their gorm ID
	Map_FormFieldStringDBID_FormFieldStringPtr map[uint]*models.FormFieldString

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoFormFieldString *BackRepoFormFieldStringStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoFormFieldString.stage
	return
}

func (backRepoFormFieldString *BackRepoFormFieldStringStruct) GetDB() *gorm.DB {
	return backRepoFormFieldString.db
}

// GetFormFieldStringDBFromFormFieldStringPtr is a handy function to access the back repo instance from the stage instance
func (backRepoFormFieldString *BackRepoFormFieldStringStruct) GetFormFieldStringDBFromFormFieldStringPtr(formfieldstring *models.FormFieldString) (formfieldstringDB *FormFieldStringDB) {
	id := backRepoFormFieldString.Map_FormFieldStringPtr_FormFieldStringDBID[formfieldstring]
	formfieldstringDB = backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringDB[id]
	return
}

// BackRepoFormFieldString.CommitPhaseOne commits all staged instances of FormFieldString to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormFieldString *BackRepoFormFieldStringStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for formfieldstring := range stage.FormFieldStrings {
		backRepoFormFieldString.CommitPhaseOneInstance(formfieldstring)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, formfieldstring := range backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringPtr {
		if _, ok := stage.FormFieldStrings[formfieldstring]; !ok {
			backRepoFormFieldString.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoFormFieldString.CommitDeleteInstance commits deletion of FormFieldString to the BackRepo
func (backRepoFormFieldString *BackRepoFormFieldStringStruct) CommitDeleteInstance(id uint) (Error error) {

	formfieldstring := backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringPtr[id]

	// formfieldstring is not staged anymore, remove formfieldstringDB
	formfieldstringDB := backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringDB[id]
	query := backRepoFormFieldString.db.Unscoped().Delete(&formfieldstringDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoFormFieldString.Map_FormFieldStringPtr_FormFieldStringDBID, formfieldstring)
	delete(backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringPtr, id)
	delete(backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringDB, id)

	return
}

// BackRepoFormFieldString.CommitPhaseOneInstance commits formfieldstring staged instances of FormFieldString to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormFieldString *BackRepoFormFieldStringStruct) CommitPhaseOneInstance(formfieldstring *models.FormFieldString) (Error error) {

	// check if the formfieldstring is not commited yet
	if _, ok := backRepoFormFieldString.Map_FormFieldStringPtr_FormFieldStringDBID[formfieldstring]; ok {
		return
	}

	// initiate formfieldstring
	var formfieldstringDB FormFieldStringDB
	formfieldstringDB.CopyBasicFieldsFromFormFieldString(formfieldstring)

	query := backRepoFormFieldString.db.Create(&formfieldstringDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoFormFieldString.Map_FormFieldStringPtr_FormFieldStringDBID[formfieldstring] = formfieldstringDB.ID
	backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringPtr[formfieldstringDB.ID] = formfieldstring
	backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringDB[formfieldstringDB.ID] = &formfieldstringDB

	return
}

// BackRepoFormFieldString.CommitPhaseTwo commits all staged instances of FormFieldString to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldString *BackRepoFormFieldStringStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, formfieldstring := range backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringPtr {
		backRepoFormFieldString.CommitPhaseTwoInstance(backRepo, idx, formfieldstring)
	}

	return
}

// BackRepoFormFieldString.CommitPhaseTwoInstance commits {{structname }} of models.FormFieldString to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldString *BackRepoFormFieldStringStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, formfieldstring *models.FormFieldString) (Error error) {

	// fetch matching formfieldstringDB
	if formfieldstringDB, ok := backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringDB[idx]; ok {

		formfieldstringDB.CopyBasicFieldsFromFormFieldString(formfieldstring)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoFormFieldString.db.Save(&formfieldstringDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown FormFieldString intance %s", formfieldstring.Name))
		return err
	}

	return
}

// BackRepoFormFieldString.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoFormFieldString *BackRepoFormFieldStringStruct) CheckoutPhaseOne() (Error error) {

	formfieldstringDBArray := make([]FormFieldStringDB, 0)
	query := backRepoFormFieldString.db.Find(&formfieldstringDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	formfieldstringInstancesToBeRemovedFromTheStage := make(map[*models.FormFieldString]any)
	for key, value := range backRepoFormFieldString.stage.FormFieldStrings {
		formfieldstringInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, formfieldstringDB := range formfieldstringDBArray {
		backRepoFormFieldString.CheckoutPhaseOneInstance(&formfieldstringDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		formfieldstring, ok := backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringPtr[formfieldstringDB.ID]
		if ok {
			delete(formfieldstringInstancesToBeRemovedFromTheStage, formfieldstring)
		}
	}

	// remove from stage and back repo's 3 maps all formfieldstrings that are not in the checkout
	for formfieldstring := range formfieldstringInstancesToBeRemovedFromTheStage {
		formfieldstring.Unstage(backRepoFormFieldString.GetStage())

		// remove instance from the back repo 3 maps
		formfieldstringID := backRepoFormFieldString.Map_FormFieldStringPtr_FormFieldStringDBID[formfieldstring]
		delete(backRepoFormFieldString.Map_FormFieldStringPtr_FormFieldStringDBID, formfieldstring)
		delete(backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringDB, formfieldstringID)
		delete(backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringPtr, formfieldstringID)
	}

	return
}

// CheckoutPhaseOneInstance takes a formfieldstringDB that has been found in the DB, updates the backRepo and stages the
// models version of the formfieldstringDB
func (backRepoFormFieldString *BackRepoFormFieldStringStruct) CheckoutPhaseOneInstance(formfieldstringDB *FormFieldStringDB) (Error error) {

	formfieldstring, ok := backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringPtr[formfieldstringDB.ID]
	if !ok {
		formfieldstring = new(models.FormFieldString)

		backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringPtr[formfieldstringDB.ID] = formfieldstring
		backRepoFormFieldString.Map_FormFieldStringPtr_FormFieldStringDBID[formfieldstring] = formfieldstringDB.ID

		// append model store with the new element
		formfieldstring.Name = formfieldstringDB.Name_Data.String
		formfieldstring.Stage(backRepoFormFieldString.GetStage())
	}
	formfieldstringDB.CopyBasicFieldsToFormFieldString(formfieldstring)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	formfieldstring.Stage(backRepoFormFieldString.GetStage())

	// preserve pointer to formfieldstringDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_FormFieldStringDBID_FormFieldStringDB)[formfieldstringDB hold variable pointers
	formfieldstringDB_Data := *formfieldstringDB
	preservedPtrToFormFieldString := &formfieldstringDB_Data
	backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringDB[formfieldstringDB.ID] = preservedPtrToFormFieldString

	return
}

// BackRepoFormFieldString.CheckoutPhaseTwo Checkouts all staged instances of FormFieldString to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldString *BackRepoFormFieldStringStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, formfieldstringDB := range backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringDB {
		backRepoFormFieldString.CheckoutPhaseTwoInstance(backRepo, formfieldstringDB)
	}
	return
}

// BackRepoFormFieldString.CheckoutPhaseTwoInstance Checkouts staged instances of FormFieldString to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldString *BackRepoFormFieldStringStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, formfieldstringDB *FormFieldStringDB) (Error error) {

	formfieldstring := backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringPtr[formfieldstringDB.ID]

	formfieldstringDB.DecodePointers(backRepo, formfieldstring)

	return
}

func (formfieldstringDB *FormFieldStringDB) DecodePointers(backRepo *BackRepoStruct, formfieldstring *models.FormFieldString) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitFormFieldString allows commit of a single formfieldstring (if already staged)
func (backRepo *BackRepoStruct) CommitFormFieldString(formfieldstring *models.FormFieldString) {
	backRepo.BackRepoFormFieldString.CommitPhaseOneInstance(formfieldstring)
	if id, ok := backRepo.BackRepoFormFieldString.Map_FormFieldStringPtr_FormFieldStringDBID[formfieldstring]; ok {
		backRepo.BackRepoFormFieldString.CommitPhaseTwoInstance(backRepo, id, formfieldstring)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitFormFieldString allows checkout of a single formfieldstring (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutFormFieldString(formfieldstring *models.FormFieldString) {
	// check if the formfieldstring is staged
	if _, ok := backRepo.BackRepoFormFieldString.Map_FormFieldStringPtr_FormFieldStringDBID[formfieldstring]; ok {

		if id, ok := backRepo.BackRepoFormFieldString.Map_FormFieldStringPtr_FormFieldStringDBID[formfieldstring]; ok {
			var formfieldstringDB FormFieldStringDB
			formfieldstringDB.ID = id

			if err := backRepo.BackRepoFormFieldString.db.First(&formfieldstringDB, id).Error; err != nil {
				log.Fatalln("CheckoutFormFieldString : Problem with getting object with id:", id)
			}
			backRepo.BackRepoFormFieldString.CheckoutPhaseOneInstance(&formfieldstringDB)
			backRepo.BackRepoFormFieldString.CheckoutPhaseTwoInstance(backRepo, &formfieldstringDB)
		}
	}
}

// CopyBasicFieldsFromFormFieldString
func (formfieldstringDB *FormFieldStringDB) CopyBasicFieldsFromFormFieldString(formfieldstring *models.FormFieldString) {
	// insertion point for fields commit

	formfieldstringDB.Name_Data.String = formfieldstring.Name
	formfieldstringDB.Name_Data.Valid = true

	formfieldstringDB.Value_Data.String = formfieldstring.Value
	formfieldstringDB.Value_Data.Valid = true

	formfieldstringDB.IsTextArea_Data.Bool = formfieldstring.IsTextArea
	formfieldstringDB.IsTextArea_Data.Valid = true
}

// CopyBasicFieldsFromFormFieldString_WOP
func (formfieldstringDB *FormFieldStringDB) CopyBasicFieldsFromFormFieldString_WOP(formfieldstring *models.FormFieldString_WOP) {
	// insertion point for fields commit

	formfieldstringDB.Name_Data.String = formfieldstring.Name
	formfieldstringDB.Name_Data.Valid = true

	formfieldstringDB.Value_Data.String = formfieldstring.Value
	formfieldstringDB.Value_Data.Valid = true

	formfieldstringDB.IsTextArea_Data.Bool = formfieldstring.IsTextArea
	formfieldstringDB.IsTextArea_Data.Valid = true
}

// CopyBasicFieldsFromFormFieldStringWOP
func (formfieldstringDB *FormFieldStringDB) CopyBasicFieldsFromFormFieldStringWOP(formfieldstring *FormFieldStringWOP) {
	// insertion point for fields commit

	formfieldstringDB.Name_Data.String = formfieldstring.Name
	formfieldstringDB.Name_Data.Valid = true

	formfieldstringDB.Value_Data.String = formfieldstring.Value
	formfieldstringDB.Value_Data.Valid = true

	formfieldstringDB.IsTextArea_Data.Bool = formfieldstring.IsTextArea
	formfieldstringDB.IsTextArea_Data.Valid = true
}

// CopyBasicFieldsToFormFieldString
func (formfieldstringDB *FormFieldStringDB) CopyBasicFieldsToFormFieldString(formfieldstring *models.FormFieldString) {
	// insertion point for checkout of basic fields (back repo to stage)
	formfieldstring.Name = formfieldstringDB.Name_Data.String
	formfieldstring.Value = formfieldstringDB.Value_Data.String
	formfieldstring.IsTextArea = formfieldstringDB.IsTextArea_Data.Bool
}

// CopyBasicFieldsToFormFieldString_WOP
func (formfieldstringDB *FormFieldStringDB) CopyBasicFieldsToFormFieldString_WOP(formfieldstring *models.FormFieldString_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	formfieldstring.Name = formfieldstringDB.Name_Data.String
	formfieldstring.Value = formfieldstringDB.Value_Data.String
	formfieldstring.IsTextArea = formfieldstringDB.IsTextArea_Data.Bool
}

// CopyBasicFieldsToFormFieldStringWOP
func (formfieldstringDB *FormFieldStringDB) CopyBasicFieldsToFormFieldStringWOP(formfieldstring *FormFieldStringWOP) {
	formfieldstring.ID = int(formfieldstringDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	formfieldstring.Name = formfieldstringDB.Name_Data.String
	formfieldstring.Value = formfieldstringDB.Value_Data.String
	formfieldstring.IsTextArea = formfieldstringDB.IsTextArea_Data.Bool
}

// Backup generates a json file from a slice of all FormFieldStringDB instances in the backrepo
func (backRepoFormFieldString *BackRepoFormFieldStringStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "FormFieldStringDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FormFieldStringDB, 0)
	for _, formfieldstringDB := range backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringDB {
		forBackup = append(forBackup, formfieldstringDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json FormFieldString ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json FormFieldString file", err.Error())
	}
}

// Backup generates a json file from a slice of all FormFieldStringDB instances in the backrepo
func (backRepoFormFieldString *BackRepoFormFieldStringStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FormFieldStringDB, 0)
	for _, formfieldstringDB := range backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringDB {
		forBackup = append(forBackup, formfieldstringDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("FormFieldString")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&FormFieldString_Fields, -1)
	for _, formfieldstringDB := range forBackup {

		var formfieldstringWOP FormFieldStringWOP
		formfieldstringDB.CopyBasicFieldsToFormFieldStringWOP(&formfieldstringWOP)

		row := sh.AddRow()
		row.WriteStruct(&formfieldstringWOP, -1)
	}
}

// RestoreXL from the "FormFieldString" sheet all FormFieldStringDB instances
func (backRepoFormFieldString *BackRepoFormFieldStringStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoFormFieldStringid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["FormFieldString"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoFormFieldString.rowVisitorFormFieldString)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoFormFieldString *BackRepoFormFieldStringStruct) rowVisitorFormFieldString(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var formfieldstringWOP FormFieldStringWOP
		row.ReadStruct(&formfieldstringWOP)

		// add the unmarshalled struct to the stage
		formfieldstringDB := new(FormFieldStringDB)
		formfieldstringDB.CopyBasicFieldsFromFormFieldStringWOP(&formfieldstringWOP)

		formfieldstringDB_ID_atBackupTime := formfieldstringDB.ID
		formfieldstringDB.ID = 0
		query := backRepoFormFieldString.db.Create(formfieldstringDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringDB[formfieldstringDB.ID] = formfieldstringDB
		BackRepoFormFieldStringid_atBckpTime_newID[formfieldstringDB_ID_atBackupTime] = formfieldstringDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "FormFieldStringDB.json" in dirPath that stores an array
// of FormFieldStringDB and stores it in the database
// the map BackRepoFormFieldStringid_atBckpTime_newID is updated accordingly
func (backRepoFormFieldString *BackRepoFormFieldStringStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoFormFieldStringid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "FormFieldStringDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json FormFieldString file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*FormFieldStringDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_FormFieldStringDBID_FormFieldStringDB
	for _, formfieldstringDB := range forRestore {

		formfieldstringDB_ID_atBackupTime := formfieldstringDB.ID
		formfieldstringDB.ID = 0
		query := backRepoFormFieldString.db.Create(formfieldstringDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringDB[formfieldstringDB.ID] = formfieldstringDB
		BackRepoFormFieldStringid_atBckpTime_newID[formfieldstringDB_ID_atBackupTime] = formfieldstringDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json FormFieldString file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<FormFieldString>id_atBckpTime_newID
// to compute new index
func (backRepoFormFieldString *BackRepoFormFieldStringStruct) RestorePhaseTwo() {

	for _, formfieldstringDB := range backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringDB {

		// next line of code is to avert unused variable compilation error
		_ = formfieldstringDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoFormFieldString.db.Model(formfieldstringDB).Updates(*formfieldstringDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoFormFieldString.ResetReversePointers commits all staged instances of FormFieldString to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldString *BackRepoFormFieldStringStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, formfieldstring := range backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringPtr {
		backRepoFormFieldString.ResetReversePointersInstance(backRepo, idx, formfieldstring)
	}

	return
}

func (backRepoFormFieldString *BackRepoFormFieldStringStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, formfieldstring *models.FormFieldString) (Error error) {

	// fetch matching formfieldstringDB
	if formfieldstringDB, ok := backRepoFormFieldString.Map_FormFieldStringDBID_FormFieldStringDB[idx]; ok {
		_ = formfieldstringDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoFormFieldStringid_atBckpTime_newID map[uint]uint
