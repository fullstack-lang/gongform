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
var dummy_FormEditAssocButton_sql sql.NullBool
var dummy_FormEditAssocButton_time time.Duration
var dummy_FormEditAssocButton_sort sort.Float64Slice

// FormEditAssocButtonAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model formeditassocbuttonAPI
type FormEditAssocButtonAPI struct {
	gorm.Model

	models.FormEditAssocButton_WOP

	// encoding of pointers
	FormEditAssocButtonPointersEncoding FormEditAssocButtonPointersEncoding
}

// FormEditAssocButtonPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type FormEditAssocButtonPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// FormEditAssocButtonDB describes a formeditassocbutton in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model formeditassocbuttonDB
type FormEditAssocButtonDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field formeditassocbuttonDB.Name
	Name_Data sql.NullString

	// Declation for basic field formeditassocbuttonDB.Label
	Label_Data sql.NullString
	// encoding of pointers
	FormEditAssocButtonPointersEncoding
}

// FormEditAssocButtonDBs arrays formeditassocbuttonDBs
// swagger:response formeditassocbuttonDBsResponse
type FormEditAssocButtonDBs []FormEditAssocButtonDB

// FormEditAssocButtonDBResponse provides response
// swagger:response formeditassocbuttonDBResponse
type FormEditAssocButtonDBResponse struct {
	FormEditAssocButtonDB
}

// FormEditAssocButtonWOP is a FormEditAssocButton without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type FormEditAssocButtonWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Label string `xlsx:"2"`
	// insertion for WOP pointer fields
}

var FormEditAssocButton_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Label",
}

type BackRepoFormEditAssocButtonStruct struct {
	// stores FormEditAssocButtonDB according to their gorm ID
	Map_FormEditAssocButtonDBID_FormEditAssocButtonDB map[uint]*FormEditAssocButtonDB

	// stores FormEditAssocButtonDB ID according to FormEditAssocButton address
	Map_FormEditAssocButtonPtr_FormEditAssocButtonDBID map[*models.FormEditAssocButton]uint

	// stores FormEditAssocButton according to their gorm ID
	Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr map[uint]*models.FormEditAssocButton

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoFormEditAssocButton.stage
	return
}

func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) GetDB() *gorm.DB {
	return backRepoFormEditAssocButton.db
}

// GetFormEditAssocButtonDBFromFormEditAssocButtonPtr is a handy function to access the back repo instance from the stage instance
func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) GetFormEditAssocButtonDBFromFormEditAssocButtonPtr(formeditassocbutton *models.FormEditAssocButton) (formeditassocbuttonDB *FormEditAssocButtonDB) {
	id := backRepoFormEditAssocButton.Map_FormEditAssocButtonPtr_FormEditAssocButtonDBID[formeditassocbutton]
	formeditassocbuttonDB = backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonDB[id]
	return
}

// BackRepoFormEditAssocButton.CommitPhaseOne commits all staged instances of FormEditAssocButton to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for formeditassocbutton := range stage.FormEditAssocButtons {
		backRepoFormEditAssocButton.CommitPhaseOneInstance(formeditassocbutton)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, formeditassocbutton := range backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr {
		if _, ok := stage.FormEditAssocButtons[formeditassocbutton]; !ok {
			backRepoFormEditAssocButton.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoFormEditAssocButton.CommitDeleteInstance commits deletion of FormEditAssocButton to the BackRepo
func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) CommitDeleteInstance(id uint) (Error error) {

	formeditassocbutton := backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr[id]

	// formeditassocbutton is not staged anymore, remove formeditassocbuttonDB
	formeditassocbuttonDB := backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonDB[id]
	query := backRepoFormEditAssocButton.db.Unscoped().Delete(&formeditassocbuttonDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoFormEditAssocButton.Map_FormEditAssocButtonPtr_FormEditAssocButtonDBID, formeditassocbutton)
	delete(backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr, id)
	delete(backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonDB, id)

	return
}

// BackRepoFormEditAssocButton.CommitPhaseOneInstance commits formeditassocbutton staged instances of FormEditAssocButton to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) CommitPhaseOneInstance(formeditassocbutton *models.FormEditAssocButton) (Error error) {

	// check if the formeditassocbutton is not commited yet
	if _, ok := backRepoFormEditAssocButton.Map_FormEditAssocButtonPtr_FormEditAssocButtonDBID[formeditassocbutton]; ok {
		return
	}

	// initiate formeditassocbutton
	var formeditassocbuttonDB FormEditAssocButtonDB
	formeditassocbuttonDB.CopyBasicFieldsFromFormEditAssocButton(formeditassocbutton)

	query := backRepoFormEditAssocButton.db.Create(&formeditassocbuttonDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoFormEditAssocButton.Map_FormEditAssocButtonPtr_FormEditAssocButtonDBID[formeditassocbutton] = formeditassocbuttonDB.ID
	backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr[formeditassocbuttonDB.ID] = formeditassocbutton
	backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonDB[formeditassocbuttonDB.ID] = &formeditassocbuttonDB

	return
}

// BackRepoFormEditAssocButton.CommitPhaseTwo commits all staged instances of FormEditAssocButton to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, formeditassocbutton := range backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr {
		backRepoFormEditAssocButton.CommitPhaseTwoInstance(backRepo, idx, formeditassocbutton)
	}

	return
}

// BackRepoFormEditAssocButton.CommitPhaseTwoInstance commits {{structname }} of models.FormEditAssocButton to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, formeditassocbutton *models.FormEditAssocButton) (Error error) {

	// fetch matching formeditassocbuttonDB
	if formeditassocbuttonDB, ok := backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonDB[idx]; ok {

		formeditassocbuttonDB.CopyBasicFieldsFromFormEditAssocButton(formeditassocbutton)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoFormEditAssocButton.db.Save(&formeditassocbuttonDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown FormEditAssocButton intance %s", formeditassocbutton.Name))
		return err
	}

	return
}

// BackRepoFormEditAssocButton.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) CheckoutPhaseOne() (Error error) {

	formeditassocbuttonDBArray := make([]FormEditAssocButtonDB, 0)
	query := backRepoFormEditAssocButton.db.Find(&formeditassocbuttonDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	formeditassocbuttonInstancesToBeRemovedFromTheStage := make(map[*models.FormEditAssocButton]any)
	for key, value := range backRepoFormEditAssocButton.stage.FormEditAssocButtons {
		formeditassocbuttonInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, formeditassocbuttonDB := range formeditassocbuttonDBArray {
		backRepoFormEditAssocButton.CheckoutPhaseOneInstance(&formeditassocbuttonDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		formeditassocbutton, ok := backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr[formeditassocbuttonDB.ID]
		if ok {
			delete(formeditassocbuttonInstancesToBeRemovedFromTheStage, formeditassocbutton)
		}
	}

	// remove from stage and back repo's 3 maps all formeditassocbuttons that are not in the checkout
	for formeditassocbutton := range formeditassocbuttonInstancesToBeRemovedFromTheStage {
		formeditassocbutton.Unstage(backRepoFormEditAssocButton.GetStage())

		// remove instance from the back repo 3 maps
		formeditassocbuttonID := backRepoFormEditAssocButton.Map_FormEditAssocButtonPtr_FormEditAssocButtonDBID[formeditassocbutton]
		delete(backRepoFormEditAssocButton.Map_FormEditAssocButtonPtr_FormEditAssocButtonDBID, formeditassocbutton)
		delete(backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonDB, formeditassocbuttonID)
		delete(backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr, formeditassocbuttonID)
	}

	return
}

// CheckoutPhaseOneInstance takes a formeditassocbuttonDB that has been found in the DB, updates the backRepo and stages the
// models version of the formeditassocbuttonDB
func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) CheckoutPhaseOneInstance(formeditassocbuttonDB *FormEditAssocButtonDB) (Error error) {

	formeditassocbutton, ok := backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr[formeditassocbuttonDB.ID]
	if !ok {
		formeditassocbutton = new(models.FormEditAssocButton)

		backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr[formeditassocbuttonDB.ID] = formeditassocbutton
		backRepoFormEditAssocButton.Map_FormEditAssocButtonPtr_FormEditAssocButtonDBID[formeditassocbutton] = formeditassocbuttonDB.ID

		// append model store with the new element
		formeditassocbutton.Name = formeditassocbuttonDB.Name_Data.String
		formeditassocbutton.Stage(backRepoFormEditAssocButton.GetStage())
	}
	formeditassocbuttonDB.CopyBasicFieldsToFormEditAssocButton(formeditassocbutton)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	formeditassocbutton.Stage(backRepoFormEditAssocButton.GetStage())

	// preserve pointer to formeditassocbuttonDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_FormEditAssocButtonDBID_FormEditAssocButtonDB)[formeditassocbuttonDB hold variable pointers
	formeditassocbuttonDB_Data := *formeditassocbuttonDB
	preservedPtrToFormEditAssocButton := &formeditassocbuttonDB_Data
	backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonDB[formeditassocbuttonDB.ID] = preservedPtrToFormEditAssocButton

	return
}

// BackRepoFormEditAssocButton.CheckoutPhaseTwo Checkouts all staged instances of FormEditAssocButton to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, formeditassocbuttonDB := range backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonDB {
		backRepoFormEditAssocButton.CheckoutPhaseTwoInstance(backRepo, formeditassocbuttonDB)
	}
	return
}

// BackRepoFormEditAssocButton.CheckoutPhaseTwoInstance Checkouts staged instances of FormEditAssocButton to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, formeditassocbuttonDB *FormEditAssocButtonDB) (Error error) {

	formeditassocbutton := backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr[formeditassocbuttonDB.ID]
	_ = formeditassocbutton // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitFormEditAssocButton allows commit of a single formeditassocbutton (if already staged)
func (backRepo *BackRepoStruct) CommitFormEditAssocButton(formeditassocbutton *models.FormEditAssocButton) {
	backRepo.BackRepoFormEditAssocButton.CommitPhaseOneInstance(formeditassocbutton)
	if id, ok := backRepo.BackRepoFormEditAssocButton.Map_FormEditAssocButtonPtr_FormEditAssocButtonDBID[formeditassocbutton]; ok {
		backRepo.BackRepoFormEditAssocButton.CommitPhaseTwoInstance(backRepo, id, formeditassocbutton)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitFormEditAssocButton allows checkout of a single formeditassocbutton (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutFormEditAssocButton(formeditassocbutton *models.FormEditAssocButton) {
	// check if the formeditassocbutton is staged
	if _, ok := backRepo.BackRepoFormEditAssocButton.Map_FormEditAssocButtonPtr_FormEditAssocButtonDBID[formeditassocbutton]; ok {

		if id, ok := backRepo.BackRepoFormEditAssocButton.Map_FormEditAssocButtonPtr_FormEditAssocButtonDBID[formeditassocbutton]; ok {
			var formeditassocbuttonDB FormEditAssocButtonDB
			formeditassocbuttonDB.ID = id

			if err := backRepo.BackRepoFormEditAssocButton.db.First(&formeditassocbuttonDB, id).Error; err != nil {
				log.Fatalln("CheckoutFormEditAssocButton : Problem with getting object with id:", id)
			}
			backRepo.BackRepoFormEditAssocButton.CheckoutPhaseOneInstance(&formeditassocbuttonDB)
			backRepo.BackRepoFormEditAssocButton.CheckoutPhaseTwoInstance(backRepo, &formeditassocbuttonDB)
		}
	}
}

// CopyBasicFieldsFromFormEditAssocButton
func (formeditassocbuttonDB *FormEditAssocButtonDB) CopyBasicFieldsFromFormEditAssocButton(formeditassocbutton *models.FormEditAssocButton) {
	// insertion point for fields commit

	formeditassocbuttonDB.Name_Data.String = formeditassocbutton.Name
	formeditassocbuttonDB.Name_Data.Valid = true

	formeditassocbuttonDB.Label_Data.String = formeditassocbutton.Label
	formeditassocbuttonDB.Label_Data.Valid = true
}

// CopyBasicFieldsFromFormEditAssocButton_WOP
func (formeditassocbuttonDB *FormEditAssocButtonDB) CopyBasicFieldsFromFormEditAssocButton_WOP(formeditassocbutton *models.FormEditAssocButton_WOP) {
	// insertion point for fields commit

	formeditassocbuttonDB.Name_Data.String = formeditassocbutton.Name
	formeditassocbuttonDB.Name_Data.Valid = true

	formeditassocbuttonDB.Label_Data.String = formeditassocbutton.Label
	formeditassocbuttonDB.Label_Data.Valid = true
}

// CopyBasicFieldsFromFormEditAssocButtonWOP
func (formeditassocbuttonDB *FormEditAssocButtonDB) CopyBasicFieldsFromFormEditAssocButtonWOP(formeditassocbutton *FormEditAssocButtonWOP) {
	// insertion point for fields commit

	formeditassocbuttonDB.Name_Data.String = formeditassocbutton.Name
	formeditassocbuttonDB.Name_Data.Valid = true

	formeditassocbuttonDB.Label_Data.String = formeditassocbutton.Label
	formeditassocbuttonDB.Label_Data.Valid = true
}

// CopyBasicFieldsToFormEditAssocButton
func (formeditassocbuttonDB *FormEditAssocButtonDB) CopyBasicFieldsToFormEditAssocButton(formeditassocbutton *models.FormEditAssocButton) {
	// insertion point for checkout of basic fields (back repo to stage)
	formeditassocbutton.Name = formeditassocbuttonDB.Name_Data.String
	formeditassocbutton.Label = formeditassocbuttonDB.Label_Data.String
}

// CopyBasicFieldsToFormEditAssocButton_WOP
func (formeditassocbuttonDB *FormEditAssocButtonDB) CopyBasicFieldsToFormEditAssocButton_WOP(formeditassocbutton *models.FormEditAssocButton_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	formeditassocbutton.Name = formeditassocbuttonDB.Name_Data.String
	formeditassocbutton.Label = formeditassocbuttonDB.Label_Data.String
}

// CopyBasicFieldsToFormEditAssocButtonWOP
func (formeditassocbuttonDB *FormEditAssocButtonDB) CopyBasicFieldsToFormEditAssocButtonWOP(formeditassocbutton *FormEditAssocButtonWOP) {
	formeditassocbutton.ID = int(formeditassocbuttonDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	formeditassocbutton.Name = formeditassocbuttonDB.Name_Data.String
	formeditassocbutton.Label = formeditassocbuttonDB.Label_Data.String
}

// Backup generates a json file from a slice of all FormEditAssocButtonDB instances in the backrepo
func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "FormEditAssocButtonDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FormEditAssocButtonDB, 0)
	for _, formeditassocbuttonDB := range backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonDB {
		forBackup = append(forBackup, formeditassocbuttonDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json FormEditAssocButton ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json FormEditAssocButton file", err.Error())
	}
}

// Backup generates a json file from a slice of all FormEditAssocButtonDB instances in the backrepo
func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FormEditAssocButtonDB, 0)
	for _, formeditassocbuttonDB := range backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonDB {
		forBackup = append(forBackup, formeditassocbuttonDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("FormEditAssocButton")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&FormEditAssocButton_Fields, -1)
	for _, formeditassocbuttonDB := range forBackup {

		var formeditassocbuttonWOP FormEditAssocButtonWOP
		formeditassocbuttonDB.CopyBasicFieldsToFormEditAssocButtonWOP(&formeditassocbuttonWOP)

		row := sh.AddRow()
		row.WriteStruct(&formeditassocbuttonWOP, -1)
	}
}

// RestoreXL from the "FormEditAssocButton" sheet all FormEditAssocButtonDB instances
func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoFormEditAssocButtonid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["FormEditAssocButton"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoFormEditAssocButton.rowVisitorFormEditAssocButton)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) rowVisitorFormEditAssocButton(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var formeditassocbuttonWOP FormEditAssocButtonWOP
		row.ReadStruct(&formeditassocbuttonWOP)

		// add the unmarshalled struct to the stage
		formeditassocbuttonDB := new(FormEditAssocButtonDB)
		formeditassocbuttonDB.CopyBasicFieldsFromFormEditAssocButtonWOP(&formeditassocbuttonWOP)

		formeditassocbuttonDB_ID_atBackupTime := formeditassocbuttonDB.ID
		formeditassocbuttonDB.ID = 0
		query := backRepoFormEditAssocButton.db.Create(formeditassocbuttonDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonDB[formeditassocbuttonDB.ID] = formeditassocbuttonDB
		BackRepoFormEditAssocButtonid_atBckpTime_newID[formeditassocbuttonDB_ID_atBackupTime] = formeditassocbuttonDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "FormEditAssocButtonDB.json" in dirPath that stores an array
// of FormEditAssocButtonDB and stores it in the database
// the map BackRepoFormEditAssocButtonid_atBckpTime_newID is updated accordingly
func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoFormEditAssocButtonid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "FormEditAssocButtonDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json FormEditAssocButton file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*FormEditAssocButtonDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_FormEditAssocButtonDBID_FormEditAssocButtonDB
	for _, formeditassocbuttonDB := range forRestore {

		formeditassocbuttonDB_ID_atBackupTime := formeditassocbuttonDB.ID
		formeditassocbuttonDB.ID = 0
		query := backRepoFormEditAssocButton.db.Create(formeditassocbuttonDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonDB[formeditassocbuttonDB.ID] = formeditassocbuttonDB
		BackRepoFormEditAssocButtonid_atBckpTime_newID[formeditassocbuttonDB_ID_atBackupTime] = formeditassocbuttonDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json FormEditAssocButton file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<FormEditAssocButton>id_atBckpTime_newID
// to compute new index
func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) RestorePhaseTwo() {

	for _, formeditassocbuttonDB := range backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonDB {

		// next line of code is to avert unused variable compilation error
		_ = formeditassocbuttonDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoFormEditAssocButton.db.Model(formeditassocbuttonDB).Updates(*formeditassocbuttonDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoFormEditAssocButton.ResetReversePointers commits all staged instances of FormEditAssocButton to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, formeditassocbutton := range backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr {
		backRepoFormEditAssocButton.ResetReversePointersInstance(backRepo, idx, formeditassocbutton)
	}

	return
}

func (backRepoFormEditAssocButton *BackRepoFormEditAssocButtonStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, astruct *models.FormEditAssocButton) (Error error) {

	// fetch matching formeditassocbuttonDB
	if formeditassocbuttonDB, ok := backRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonDB[idx]; ok {
		_ = formeditassocbuttonDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoFormEditAssocButtonid_atBckpTime_newID map[uint]uint
