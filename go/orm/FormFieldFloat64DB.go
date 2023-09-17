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
var dummy_FormFieldFloat64_sql sql.NullBool
var dummy_FormFieldFloat64_time time.Duration
var dummy_FormFieldFloat64_sort sort.Float64Slice

// FormFieldFloat64API is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model formfieldfloat64API
type FormFieldFloat64API struct {
	gorm.Model

	models.FormFieldFloat64

	// encoding of pointers
	FormFieldFloat64PointersEnconding
}

// FormFieldFloat64PointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type FormFieldFloat64PointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// FormFieldFloat64DB describes a formfieldfloat64 in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model formfieldfloat64DB
type FormFieldFloat64DB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field formfieldfloat64DB.Name
	Name_Data sql.NullString

	// Declation for basic field formfieldfloat64DB.Value
	Value_Data sql.NullFloat64

	// Declation for basic field formfieldfloat64DB.HasMinValidator
	// provide the sql storage for the boolan
	HasMinValidator_Data sql.NullBool

	// Declation for basic field formfieldfloat64DB.MinValue
	MinValue_Data sql.NullFloat64

	// Declation for basic field formfieldfloat64DB.HasMaxValidator
	// provide the sql storage for the boolan
	HasMaxValidator_Data sql.NullBool

	// Declation for basic field formfieldfloat64DB.MaxValue
	MaxValue_Data sql.NullFloat64
	// encoding of pointers
	FormFieldFloat64PointersEnconding
}

// FormFieldFloat64DBs arrays formfieldfloat64DBs
// swagger:response formfieldfloat64DBsResponse
type FormFieldFloat64DBs []FormFieldFloat64DB

// FormFieldFloat64DBResponse provides response
// swagger:response formfieldfloat64DBResponse
type FormFieldFloat64DBResponse struct {
	FormFieldFloat64DB
}

// FormFieldFloat64WOP is a FormFieldFloat64 without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type FormFieldFloat64WOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Value float64 `xlsx:"2"`

	HasMinValidator bool `xlsx:"3"`

	MinValue float64 `xlsx:"4"`

	HasMaxValidator bool `xlsx:"5"`

	MaxValue float64 `xlsx:"6"`
	// insertion for WOP pointer fields
}

var FormFieldFloat64_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Value",
	"HasMinValidator",
	"MinValue",
	"HasMaxValidator",
	"MaxValue",
}

type BackRepoFormFieldFloat64Struct struct {
	// stores FormFieldFloat64DB according to their gorm ID
	Map_FormFieldFloat64DBID_FormFieldFloat64DB map[uint]*FormFieldFloat64DB

	// stores FormFieldFloat64DB ID according to FormFieldFloat64 address
	Map_FormFieldFloat64Ptr_FormFieldFloat64DBID map[*models.FormFieldFloat64]uint

	// stores FormFieldFloat64 according to their gorm ID
	Map_FormFieldFloat64DBID_FormFieldFloat64Ptr map[uint]*models.FormFieldFloat64

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) GetStage() (stage *models.StageStruct) {
	stage = backRepoFormFieldFloat64.stage
	return
}

func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) GetDB() *gorm.DB {
	return backRepoFormFieldFloat64.db
}

// GetFormFieldFloat64DBFromFormFieldFloat64Ptr is a handy function to access the back repo instance from the stage instance
func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) GetFormFieldFloat64DBFromFormFieldFloat64Ptr(formfieldfloat64 *models.FormFieldFloat64) (formfieldfloat64DB *FormFieldFloat64DB) {
	id := backRepoFormFieldFloat64.Map_FormFieldFloat64Ptr_FormFieldFloat64DBID[formfieldfloat64]
	formfieldfloat64DB = backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64DB[id]
	return
}

// BackRepoFormFieldFloat64.CommitPhaseOne commits all staged instances of FormFieldFloat64 to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for formfieldfloat64 := range stage.FormFieldFloat64s {
		backRepoFormFieldFloat64.CommitPhaseOneInstance(formfieldfloat64)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, formfieldfloat64 := range backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64Ptr {
		if _, ok := stage.FormFieldFloat64s[formfieldfloat64]; !ok {
			backRepoFormFieldFloat64.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoFormFieldFloat64.CommitDeleteInstance commits deletion of FormFieldFloat64 to the BackRepo
func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) CommitDeleteInstance(id uint) (Error error) {

	formfieldfloat64 := backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64Ptr[id]

	// formfieldfloat64 is not staged anymore, remove formfieldfloat64DB
	formfieldfloat64DB := backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64DB[id]
	query := backRepoFormFieldFloat64.db.Unscoped().Delete(&formfieldfloat64DB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoFormFieldFloat64.Map_FormFieldFloat64Ptr_FormFieldFloat64DBID, formfieldfloat64)
	delete(backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64Ptr, id)
	delete(backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64DB, id)

	return
}

// BackRepoFormFieldFloat64.CommitPhaseOneInstance commits formfieldfloat64 staged instances of FormFieldFloat64 to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) CommitPhaseOneInstance(formfieldfloat64 *models.FormFieldFloat64) (Error error) {

	// check if the formfieldfloat64 is not commited yet
	if _, ok := backRepoFormFieldFloat64.Map_FormFieldFloat64Ptr_FormFieldFloat64DBID[formfieldfloat64]; ok {
		return
	}

	// initiate formfieldfloat64
	var formfieldfloat64DB FormFieldFloat64DB
	formfieldfloat64DB.CopyBasicFieldsFromFormFieldFloat64(formfieldfloat64)

	query := backRepoFormFieldFloat64.db.Create(&formfieldfloat64DB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoFormFieldFloat64.Map_FormFieldFloat64Ptr_FormFieldFloat64DBID[formfieldfloat64] = formfieldfloat64DB.ID
	backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64Ptr[formfieldfloat64DB.ID] = formfieldfloat64
	backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64DB[formfieldfloat64DB.ID] = &formfieldfloat64DB

	return
}

// BackRepoFormFieldFloat64.CommitPhaseTwo commits all staged instances of FormFieldFloat64 to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, formfieldfloat64 := range backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64Ptr {
		backRepoFormFieldFloat64.CommitPhaseTwoInstance(backRepo, idx, formfieldfloat64)
	}

	return
}

// BackRepoFormFieldFloat64.CommitPhaseTwoInstance commits {{structname }} of models.FormFieldFloat64 to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, formfieldfloat64 *models.FormFieldFloat64) (Error error) {

	// fetch matching formfieldfloat64DB
	if formfieldfloat64DB, ok := backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64DB[idx]; ok {

		formfieldfloat64DB.CopyBasicFieldsFromFormFieldFloat64(formfieldfloat64)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoFormFieldFloat64.db.Save(&formfieldfloat64DB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown FormFieldFloat64 intance %s", formfieldfloat64.Name))
		return err
	}

	return
}

// BackRepoFormFieldFloat64.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) CheckoutPhaseOne() (Error error) {

	formfieldfloat64DBArray := make([]FormFieldFloat64DB, 0)
	query := backRepoFormFieldFloat64.db.Find(&formfieldfloat64DBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	formfieldfloat64InstancesToBeRemovedFromTheStage := make(map[*models.FormFieldFloat64]any)
	for key, value := range backRepoFormFieldFloat64.stage.FormFieldFloat64s {
		formfieldfloat64InstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, formfieldfloat64DB := range formfieldfloat64DBArray {
		backRepoFormFieldFloat64.CheckoutPhaseOneInstance(&formfieldfloat64DB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		formfieldfloat64, ok := backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64Ptr[formfieldfloat64DB.ID]
		if ok {
			delete(formfieldfloat64InstancesToBeRemovedFromTheStage, formfieldfloat64)
		}
	}

	// remove from stage and back repo's 3 maps all formfieldfloat64s that are not in the checkout
	for formfieldfloat64 := range formfieldfloat64InstancesToBeRemovedFromTheStage {
		formfieldfloat64.Unstage(backRepoFormFieldFloat64.GetStage())

		// remove instance from the back repo 3 maps
		formfieldfloat64ID := backRepoFormFieldFloat64.Map_FormFieldFloat64Ptr_FormFieldFloat64DBID[formfieldfloat64]
		delete(backRepoFormFieldFloat64.Map_FormFieldFloat64Ptr_FormFieldFloat64DBID, formfieldfloat64)
		delete(backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64DB, formfieldfloat64ID)
		delete(backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64Ptr, formfieldfloat64ID)
	}

	return
}

// CheckoutPhaseOneInstance takes a formfieldfloat64DB that has been found in the DB, updates the backRepo and stages the
// models version of the formfieldfloat64DB
func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) CheckoutPhaseOneInstance(formfieldfloat64DB *FormFieldFloat64DB) (Error error) {

	formfieldfloat64, ok := backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64Ptr[formfieldfloat64DB.ID]
	if !ok {
		formfieldfloat64 = new(models.FormFieldFloat64)

		backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64Ptr[formfieldfloat64DB.ID] = formfieldfloat64
		backRepoFormFieldFloat64.Map_FormFieldFloat64Ptr_FormFieldFloat64DBID[formfieldfloat64] = formfieldfloat64DB.ID

		// append model store with the new element
		formfieldfloat64.Name = formfieldfloat64DB.Name_Data.String
		formfieldfloat64.Stage(backRepoFormFieldFloat64.GetStage())
	}
	formfieldfloat64DB.CopyBasicFieldsToFormFieldFloat64(formfieldfloat64)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	formfieldfloat64.Stage(backRepoFormFieldFloat64.GetStage())

	// preserve pointer to formfieldfloat64DB. Otherwise, pointer will is recycled and the map of pointers
	// Map_FormFieldFloat64DBID_FormFieldFloat64DB)[formfieldfloat64DB hold variable pointers
	formfieldfloat64DB_Data := *formfieldfloat64DB
	preservedPtrToFormFieldFloat64 := &formfieldfloat64DB_Data
	backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64DB[formfieldfloat64DB.ID] = preservedPtrToFormFieldFloat64

	return
}

// BackRepoFormFieldFloat64.CheckoutPhaseTwo Checkouts all staged instances of FormFieldFloat64 to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, formfieldfloat64DB := range backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64DB {
		backRepoFormFieldFloat64.CheckoutPhaseTwoInstance(backRepo, formfieldfloat64DB)
	}
	return
}

// BackRepoFormFieldFloat64.CheckoutPhaseTwoInstance Checkouts staged instances of FormFieldFloat64 to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, formfieldfloat64DB *FormFieldFloat64DB) (Error error) {

	formfieldfloat64 := backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64Ptr[formfieldfloat64DB.ID]
	_ = formfieldfloat64 // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitFormFieldFloat64 allows commit of a single formfieldfloat64 (if already staged)
func (backRepo *BackRepoStruct) CommitFormFieldFloat64(formfieldfloat64 *models.FormFieldFloat64) {
	backRepo.BackRepoFormFieldFloat64.CommitPhaseOneInstance(formfieldfloat64)
	if id, ok := backRepo.BackRepoFormFieldFloat64.Map_FormFieldFloat64Ptr_FormFieldFloat64DBID[formfieldfloat64]; ok {
		backRepo.BackRepoFormFieldFloat64.CommitPhaseTwoInstance(backRepo, id, formfieldfloat64)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitFormFieldFloat64 allows checkout of a single formfieldfloat64 (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutFormFieldFloat64(formfieldfloat64 *models.FormFieldFloat64) {
	// check if the formfieldfloat64 is staged
	if _, ok := backRepo.BackRepoFormFieldFloat64.Map_FormFieldFloat64Ptr_FormFieldFloat64DBID[formfieldfloat64]; ok {

		if id, ok := backRepo.BackRepoFormFieldFloat64.Map_FormFieldFloat64Ptr_FormFieldFloat64DBID[formfieldfloat64]; ok {
			var formfieldfloat64DB FormFieldFloat64DB
			formfieldfloat64DB.ID = id

			if err := backRepo.BackRepoFormFieldFloat64.db.First(&formfieldfloat64DB, id).Error; err != nil {
				log.Panicln("CheckoutFormFieldFloat64 : Problem with getting object with id:", id)
			}
			backRepo.BackRepoFormFieldFloat64.CheckoutPhaseOneInstance(&formfieldfloat64DB)
			backRepo.BackRepoFormFieldFloat64.CheckoutPhaseTwoInstance(backRepo, &formfieldfloat64DB)
		}
	}
}

// CopyBasicFieldsFromFormFieldFloat64
func (formfieldfloat64DB *FormFieldFloat64DB) CopyBasicFieldsFromFormFieldFloat64(formfieldfloat64 *models.FormFieldFloat64) {
	// insertion point for fields commit

	formfieldfloat64DB.Name_Data.String = formfieldfloat64.Name
	formfieldfloat64DB.Name_Data.Valid = true

	formfieldfloat64DB.Value_Data.Float64 = formfieldfloat64.Value
	formfieldfloat64DB.Value_Data.Valid = true

	formfieldfloat64DB.HasMinValidator_Data.Bool = formfieldfloat64.HasMinValidator
	formfieldfloat64DB.HasMinValidator_Data.Valid = true

	formfieldfloat64DB.MinValue_Data.Float64 = formfieldfloat64.MinValue
	formfieldfloat64DB.MinValue_Data.Valid = true

	formfieldfloat64DB.HasMaxValidator_Data.Bool = formfieldfloat64.HasMaxValidator
	formfieldfloat64DB.HasMaxValidator_Data.Valid = true

	formfieldfloat64DB.MaxValue_Data.Float64 = formfieldfloat64.MaxValue
	formfieldfloat64DB.MaxValue_Data.Valid = true
}

// CopyBasicFieldsFromFormFieldFloat64WOP
func (formfieldfloat64DB *FormFieldFloat64DB) CopyBasicFieldsFromFormFieldFloat64WOP(formfieldfloat64 *FormFieldFloat64WOP) {
	// insertion point for fields commit

	formfieldfloat64DB.Name_Data.String = formfieldfloat64.Name
	formfieldfloat64DB.Name_Data.Valid = true

	formfieldfloat64DB.Value_Data.Float64 = formfieldfloat64.Value
	formfieldfloat64DB.Value_Data.Valid = true

	formfieldfloat64DB.HasMinValidator_Data.Bool = formfieldfloat64.HasMinValidator
	formfieldfloat64DB.HasMinValidator_Data.Valid = true

	formfieldfloat64DB.MinValue_Data.Float64 = formfieldfloat64.MinValue
	formfieldfloat64DB.MinValue_Data.Valid = true

	formfieldfloat64DB.HasMaxValidator_Data.Bool = formfieldfloat64.HasMaxValidator
	formfieldfloat64DB.HasMaxValidator_Data.Valid = true

	formfieldfloat64DB.MaxValue_Data.Float64 = formfieldfloat64.MaxValue
	formfieldfloat64DB.MaxValue_Data.Valid = true
}

// CopyBasicFieldsToFormFieldFloat64
func (formfieldfloat64DB *FormFieldFloat64DB) CopyBasicFieldsToFormFieldFloat64(formfieldfloat64 *models.FormFieldFloat64) {
	// insertion point for checkout of basic fields (back repo to stage)
	formfieldfloat64.Name = formfieldfloat64DB.Name_Data.String
	formfieldfloat64.Value = formfieldfloat64DB.Value_Data.Float64
	formfieldfloat64.HasMinValidator = formfieldfloat64DB.HasMinValidator_Data.Bool
	formfieldfloat64.MinValue = formfieldfloat64DB.MinValue_Data.Float64
	formfieldfloat64.HasMaxValidator = formfieldfloat64DB.HasMaxValidator_Data.Bool
	formfieldfloat64.MaxValue = formfieldfloat64DB.MaxValue_Data.Float64
}

// CopyBasicFieldsToFormFieldFloat64WOP
func (formfieldfloat64DB *FormFieldFloat64DB) CopyBasicFieldsToFormFieldFloat64WOP(formfieldfloat64 *FormFieldFloat64WOP) {
	formfieldfloat64.ID = int(formfieldfloat64DB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	formfieldfloat64.Name = formfieldfloat64DB.Name_Data.String
	formfieldfloat64.Value = formfieldfloat64DB.Value_Data.Float64
	formfieldfloat64.HasMinValidator = formfieldfloat64DB.HasMinValidator_Data.Bool
	formfieldfloat64.MinValue = formfieldfloat64DB.MinValue_Data.Float64
	formfieldfloat64.HasMaxValidator = formfieldfloat64DB.HasMaxValidator_Data.Bool
	formfieldfloat64.MaxValue = formfieldfloat64DB.MaxValue_Data.Float64
}

// Backup generates a json file from a slice of all FormFieldFloat64DB instances in the backrepo
func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "FormFieldFloat64DB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FormFieldFloat64DB, 0)
	for _, formfieldfloat64DB := range backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64DB {
		forBackup = append(forBackup, formfieldfloat64DB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json FormFieldFloat64 ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json FormFieldFloat64 file", err.Error())
	}
}

// Backup generates a json file from a slice of all FormFieldFloat64DB instances in the backrepo
func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FormFieldFloat64DB, 0)
	for _, formfieldfloat64DB := range backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64DB {
		forBackup = append(forBackup, formfieldfloat64DB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("FormFieldFloat64")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&FormFieldFloat64_Fields, -1)
	for _, formfieldfloat64DB := range forBackup {

		var formfieldfloat64WOP FormFieldFloat64WOP
		formfieldfloat64DB.CopyBasicFieldsToFormFieldFloat64WOP(&formfieldfloat64WOP)

		row := sh.AddRow()
		row.WriteStruct(&formfieldfloat64WOP, -1)
	}
}

// RestoreXL from the "FormFieldFloat64" sheet all FormFieldFloat64DB instances
func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoFormFieldFloat64id_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["FormFieldFloat64"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoFormFieldFloat64.rowVisitorFormFieldFloat64)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) rowVisitorFormFieldFloat64(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var formfieldfloat64WOP FormFieldFloat64WOP
		row.ReadStruct(&formfieldfloat64WOP)

		// add the unmarshalled struct to the stage
		formfieldfloat64DB := new(FormFieldFloat64DB)
		formfieldfloat64DB.CopyBasicFieldsFromFormFieldFloat64WOP(&formfieldfloat64WOP)

		formfieldfloat64DB_ID_atBackupTime := formfieldfloat64DB.ID
		formfieldfloat64DB.ID = 0
		query := backRepoFormFieldFloat64.db.Create(formfieldfloat64DB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64DB[formfieldfloat64DB.ID] = formfieldfloat64DB
		BackRepoFormFieldFloat64id_atBckpTime_newID[formfieldfloat64DB_ID_atBackupTime] = formfieldfloat64DB.ID
	}
	return nil
}

// RestorePhaseOne read the file "FormFieldFloat64DB.json" in dirPath that stores an array
// of FormFieldFloat64DB and stores it in the database
// the map BackRepoFormFieldFloat64id_atBckpTime_newID is updated accordingly
func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoFormFieldFloat64id_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "FormFieldFloat64DB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json FormFieldFloat64 file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*FormFieldFloat64DB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_FormFieldFloat64DBID_FormFieldFloat64DB
	for _, formfieldfloat64DB := range forRestore {

		formfieldfloat64DB_ID_atBackupTime := formfieldfloat64DB.ID
		formfieldfloat64DB.ID = 0
		query := backRepoFormFieldFloat64.db.Create(formfieldfloat64DB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64DB[formfieldfloat64DB.ID] = formfieldfloat64DB
		BackRepoFormFieldFloat64id_atBckpTime_newID[formfieldfloat64DB_ID_atBackupTime] = formfieldfloat64DB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json FormFieldFloat64 file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<FormFieldFloat64>id_atBckpTime_newID
// to compute new index
func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) RestorePhaseTwo() {

	for _, formfieldfloat64DB := range backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64DB {

		// next line of code is to avert unused variable compilation error
		_ = formfieldfloat64DB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoFormFieldFloat64.db.Model(formfieldfloat64DB).Updates(*formfieldfloat64DB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// BackRepoFormFieldFloat64.ResetReversePointers commits all staged instances of FormFieldFloat64 to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, formfieldfloat64 := range backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64Ptr {
		backRepoFormFieldFloat64.ResetReversePointersInstance(backRepo, idx, formfieldfloat64)
	}

	return
}

func (backRepoFormFieldFloat64 *BackRepoFormFieldFloat64Struct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, astruct *models.FormFieldFloat64) (Error error) {

	// fetch matching formfieldfloat64DB
	if formfieldfloat64DB, ok := backRepoFormFieldFloat64.Map_FormFieldFloat64DBID_FormFieldFloat64DB[idx]; ok {
		_ = formfieldfloat64DB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoFormFieldFloat64id_atBckpTime_newID map[uint]uint
