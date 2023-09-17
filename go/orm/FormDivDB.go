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
var dummy_FormDiv_sql sql.NullBool
var dummy_FormDiv_time time.Duration
var dummy_FormDiv_sort sort.Float64Slice

// FormDivAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model formdivAPI
type FormDivAPI struct {
	gorm.Model

	models.FormDiv

	// encoding of pointers
	FormDivPointersEnconding
}

// FormDivPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type FormDivPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// field FormEditAssocButton is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	FormEditAssocButtonID sql.NullInt64

	// field FormSortAssocButton is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	FormSortAssocButtonID sql.NullInt64

	// Implementation of a reverse ID for field FormGroup{}.FormDivs []*FormDiv
	FormGroup_FormDivsDBID sql.NullInt64

	// implementation of the index of the withing the slice
	FormGroup_FormDivsDBID_Index sql.NullInt64
}

// FormDivDB describes a formdiv in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model formdivDB
type FormDivDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field formdivDB.Name
	Name_Data sql.NullString
	// encoding of pointers
	FormDivPointersEnconding
}

// FormDivDBs arrays formdivDBs
// swagger:response formdivDBsResponse
type FormDivDBs []FormDivDB

// FormDivDBResponse provides response
// swagger:response formdivDBResponse
type FormDivDBResponse struct {
	FormDivDB
}

// FormDivWOP is a FormDiv without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type FormDivWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var FormDiv_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoFormDivStruct struct {
	// stores FormDivDB according to their gorm ID
	Map_FormDivDBID_FormDivDB map[uint]*FormDivDB

	// stores FormDivDB ID according to FormDiv address
	Map_FormDivPtr_FormDivDBID map[*models.FormDiv]uint

	// stores FormDiv according to their gorm ID
	Map_FormDivDBID_FormDivPtr map[uint]*models.FormDiv

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoFormDiv *BackRepoFormDivStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoFormDiv.stage
	return
}

func (backRepoFormDiv *BackRepoFormDivStruct) GetDB() *gorm.DB {
	return backRepoFormDiv.db
}

// GetFormDivDBFromFormDivPtr is a handy function to access the back repo instance from the stage instance
func (backRepoFormDiv *BackRepoFormDivStruct) GetFormDivDBFromFormDivPtr(formdiv *models.FormDiv) (formdivDB *FormDivDB) {
	id := backRepoFormDiv.Map_FormDivPtr_FormDivDBID[formdiv]
	formdivDB = backRepoFormDiv.Map_FormDivDBID_FormDivDB[id]
	return
}

// BackRepoFormDiv.CommitPhaseOne commits all staged instances of FormDiv to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormDiv *BackRepoFormDivStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for formdiv := range stage.FormDivs {
		backRepoFormDiv.CommitPhaseOneInstance(formdiv)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, formdiv := range backRepoFormDiv.Map_FormDivDBID_FormDivPtr {
		if _, ok := stage.FormDivs[formdiv]; !ok {
			backRepoFormDiv.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoFormDiv.CommitDeleteInstance commits deletion of FormDiv to the BackRepo
func (backRepoFormDiv *BackRepoFormDivStruct) CommitDeleteInstance(id uint) (Error error) {

	formdiv := backRepoFormDiv.Map_FormDivDBID_FormDivPtr[id]

	// formdiv is not staged anymore, remove formdivDB
	formdivDB := backRepoFormDiv.Map_FormDivDBID_FormDivDB[id]
	query := backRepoFormDiv.db.Unscoped().Delete(&formdivDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoFormDiv.Map_FormDivPtr_FormDivDBID, formdiv)
	delete(backRepoFormDiv.Map_FormDivDBID_FormDivPtr, id)
	delete(backRepoFormDiv.Map_FormDivDBID_FormDivDB, id)

	return
}

// BackRepoFormDiv.CommitPhaseOneInstance commits formdiv staged instances of FormDiv to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormDiv *BackRepoFormDivStruct) CommitPhaseOneInstance(formdiv *models.FormDiv) (Error error) {

	// check if the formdiv is not commited yet
	if _, ok := backRepoFormDiv.Map_FormDivPtr_FormDivDBID[formdiv]; ok {
		return
	}

	// initiate formdiv
	var formdivDB FormDivDB
	formdivDB.CopyBasicFieldsFromFormDiv(formdiv)

	query := backRepoFormDiv.db.Create(&formdivDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoFormDiv.Map_FormDivPtr_FormDivDBID[formdiv] = formdivDB.ID
	backRepoFormDiv.Map_FormDivDBID_FormDivPtr[formdivDB.ID] = formdiv
	backRepoFormDiv.Map_FormDivDBID_FormDivDB[formdivDB.ID] = &formdivDB

	return
}

// BackRepoFormDiv.CommitPhaseTwo commits all staged instances of FormDiv to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormDiv *BackRepoFormDivStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, formdiv := range backRepoFormDiv.Map_FormDivDBID_FormDivPtr {
		backRepoFormDiv.CommitPhaseTwoInstance(backRepo, idx, formdiv)
	}

	return
}

// BackRepoFormDiv.CommitPhaseTwoInstance commits {{structname }} of models.FormDiv to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormDiv *BackRepoFormDivStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, formdiv *models.FormDiv) (Error error) {

	// fetch matching formdivDB
	if formdivDB, ok := backRepoFormDiv.Map_FormDivDBID_FormDivDB[idx]; ok {

		formdivDB.CopyBasicFieldsFromFormDiv(formdiv)

		// insertion point for translating pointers encodings into actual pointers
		// This loop encodes the slice of pointers formdiv.FormFields into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, formfieldAssocEnd := range formdiv.FormFields {

			// get the back repo instance at the association end
			formfieldAssocEnd_DB :=
				backRepo.BackRepoFormField.GetFormFieldDBFromFormFieldPtr(formfieldAssocEnd)

			// encode reverse pointer in the association end back repo instance
			formfieldAssocEnd_DB.FormDiv_FormFieldsDBID.Int64 = int64(formdivDB.ID)
			formfieldAssocEnd_DB.FormDiv_FormFieldsDBID.Valid = true
			formfieldAssocEnd_DB.FormDiv_FormFieldsDBID_Index.Int64 = int64(idx)
			formfieldAssocEnd_DB.FormDiv_FormFieldsDBID_Index.Valid = true
			if q := backRepoFormDiv.db.Save(formfieldAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		// This loop encodes the slice of pointers formdiv.CheckBoxs into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, checkboxAssocEnd := range formdiv.CheckBoxs {

			// get the back repo instance at the association end
			checkboxAssocEnd_DB :=
				backRepo.BackRepoCheckBox.GetCheckBoxDBFromCheckBoxPtr(checkboxAssocEnd)

			// encode reverse pointer in the association end back repo instance
			checkboxAssocEnd_DB.FormDiv_CheckBoxsDBID.Int64 = int64(formdivDB.ID)
			checkboxAssocEnd_DB.FormDiv_CheckBoxsDBID.Valid = true
			checkboxAssocEnd_DB.FormDiv_CheckBoxsDBID_Index.Int64 = int64(idx)
			checkboxAssocEnd_DB.FormDiv_CheckBoxsDBID_Index.Valid = true
			if q := backRepoFormDiv.db.Save(checkboxAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		// commit pointer value formdiv.FormEditAssocButton translates to updating the formdiv.FormEditAssocButtonID
		formdivDB.FormEditAssocButtonID.Valid = true // allow for a 0 value (nil association)
		if formdiv.FormEditAssocButton != nil {
			if FormEditAssocButtonId, ok := backRepo.BackRepoFormEditAssocButton.Map_FormEditAssocButtonPtr_FormEditAssocButtonDBID[formdiv.FormEditAssocButton]; ok {
				formdivDB.FormEditAssocButtonID.Int64 = int64(FormEditAssocButtonId)
				formdivDB.FormEditAssocButtonID.Valid = true
			}
		} else {
			formdivDB.FormEditAssocButtonID.Int64 = 0
			formdivDB.FormEditAssocButtonID.Valid = true
		}

		// commit pointer value formdiv.FormSortAssocButton translates to updating the formdiv.FormSortAssocButtonID
		formdivDB.FormSortAssocButtonID.Valid = true // allow for a 0 value (nil association)
		if formdiv.FormSortAssocButton != nil {
			if FormSortAssocButtonId, ok := backRepo.BackRepoFormSortAssocButton.Map_FormSortAssocButtonPtr_FormSortAssocButtonDBID[formdiv.FormSortAssocButton]; ok {
				formdivDB.FormSortAssocButtonID.Int64 = int64(FormSortAssocButtonId)
				formdivDB.FormSortAssocButtonID.Valid = true
			}
		} else {
			formdivDB.FormSortAssocButtonID.Int64 = 0
			formdivDB.FormSortAssocButtonID.Valid = true
		}

		query := backRepoFormDiv.db.Save(&formdivDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown FormDiv intance %s", formdiv.Name))
		return err
	}

	return
}

// BackRepoFormDiv.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoFormDiv *BackRepoFormDivStruct) CheckoutPhaseOne() (Error error) {

	formdivDBArray := make([]FormDivDB, 0)
	query := backRepoFormDiv.db.Find(&formdivDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	formdivInstancesToBeRemovedFromTheStage := make(map[*models.FormDiv]any)
	for key, value := range backRepoFormDiv.stage.FormDivs {
		formdivInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, formdivDB := range formdivDBArray {
		backRepoFormDiv.CheckoutPhaseOneInstance(&formdivDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		formdiv, ok := backRepoFormDiv.Map_FormDivDBID_FormDivPtr[formdivDB.ID]
		if ok {
			delete(formdivInstancesToBeRemovedFromTheStage, formdiv)
		}
	}

	// remove from stage and back repo's 3 maps all formdivs that are not in the checkout
	for formdiv := range formdivInstancesToBeRemovedFromTheStage {
		formdiv.Unstage(backRepoFormDiv.GetStage())

		// remove instance from the back repo 3 maps
		formdivID := backRepoFormDiv.Map_FormDivPtr_FormDivDBID[formdiv]
		delete(backRepoFormDiv.Map_FormDivPtr_FormDivDBID, formdiv)
		delete(backRepoFormDiv.Map_FormDivDBID_FormDivDB, formdivID)
		delete(backRepoFormDiv.Map_FormDivDBID_FormDivPtr, formdivID)
	}

	return
}

// CheckoutPhaseOneInstance takes a formdivDB that has been found in the DB, updates the backRepo and stages the
// models version of the formdivDB
func (backRepoFormDiv *BackRepoFormDivStruct) CheckoutPhaseOneInstance(formdivDB *FormDivDB) (Error error) {

	formdiv, ok := backRepoFormDiv.Map_FormDivDBID_FormDivPtr[formdivDB.ID]
	if !ok {
		formdiv = new(models.FormDiv)

		backRepoFormDiv.Map_FormDivDBID_FormDivPtr[formdivDB.ID] = formdiv
		backRepoFormDiv.Map_FormDivPtr_FormDivDBID[formdiv] = formdivDB.ID

		// append model store with the new element
		formdiv.Name = formdivDB.Name_Data.String
		formdiv.Stage(backRepoFormDiv.GetStage())
	}
	formdivDB.CopyBasicFieldsToFormDiv(formdiv)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	formdiv.Stage(backRepoFormDiv.GetStage())

	// preserve pointer to formdivDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_FormDivDBID_FormDivDB)[formdivDB hold variable pointers
	formdivDB_Data := *formdivDB
	preservedPtrToFormDiv := &formdivDB_Data
	backRepoFormDiv.Map_FormDivDBID_FormDivDB[formdivDB.ID] = preservedPtrToFormDiv

	return
}

// BackRepoFormDiv.CheckoutPhaseTwo Checkouts all staged instances of FormDiv to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormDiv *BackRepoFormDivStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, formdivDB := range backRepoFormDiv.Map_FormDivDBID_FormDivDB {
		backRepoFormDiv.CheckoutPhaseTwoInstance(backRepo, formdivDB)
	}
	return
}

// BackRepoFormDiv.CheckoutPhaseTwoInstance Checkouts staged instances of FormDiv to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormDiv *BackRepoFormDivStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, formdivDB *FormDivDB) (Error error) {

	formdiv := backRepoFormDiv.Map_FormDivDBID_FormDivPtr[formdivDB.ID]
	_ = formdiv // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// This loop redeem formdiv.FormFields in the stage from the encode in the back repo
	// It parses all FormFieldDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	formdiv.FormFields = formdiv.FormFields[:0]
	// 2. loop all instances in the type in the association end
	for _, formfieldDB_AssocEnd := range backRepo.BackRepoFormField.Map_FormFieldDBID_FormFieldDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if formfieldDB_AssocEnd.FormDiv_FormFieldsDBID.Int64 == int64(formdivDB.ID) {
			// 4. fetch the associated instance in the stage
			formfield_AssocEnd := backRepo.BackRepoFormField.Map_FormFieldDBID_FormFieldPtr[formfieldDB_AssocEnd.ID]
			// 5. append it the association slice
			formdiv.FormFields = append(formdiv.FormFields, formfield_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(formdiv.FormFields, func(i, j int) bool {
		formfieldDB_i_ID := backRepo.BackRepoFormField.Map_FormFieldPtr_FormFieldDBID[formdiv.FormFields[i]]
		formfieldDB_j_ID := backRepo.BackRepoFormField.Map_FormFieldPtr_FormFieldDBID[formdiv.FormFields[j]]

		formfieldDB_i := backRepo.BackRepoFormField.Map_FormFieldDBID_FormFieldDB[formfieldDB_i_ID]
		formfieldDB_j := backRepo.BackRepoFormField.Map_FormFieldDBID_FormFieldDB[formfieldDB_j_ID]

		return formfieldDB_i.FormDiv_FormFieldsDBID_Index.Int64 < formfieldDB_j.FormDiv_FormFieldsDBID_Index.Int64
	})

	// This loop redeem formdiv.CheckBoxs in the stage from the encode in the back repo
	// It parses all CheckBoxDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	formdiv.CheckBoxs = formdiv.CheckBoxs[:0]
	// 2. loop all instances in the type in the association end
	for _, checkboxDB_AssocEnd := range backRepo.BackRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if checkboxDB_AssocEnd.FormDiv_CheckBoxsDBID.Int64 == int64(formdivDB.ID) {
			// 4. fetch the associated instance in the stage
			checkbox_AssocEnd := backRepo.BackRepoCheckBox.Map_CheckBoxDBID_CheckBoxPtr[checkboxDB_AssocEnd.ID]
			// 5. append it the association slice
			formdiv.CheckBoxs = append(formdiv.CheckBoxs, checkbox_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(formdiv.CheckBoxs, func(i, j int) bool {
		checkboxDB_i_ID := backRepo.BackRepoCheckBox.Map_CheckBoxPtr_CheckBoxDBID[formdiv.CheckBoxs[i]]
		checkboxDB_j_ID := backRepo.BackRepoCheckBox.Map_CheckBoxPtr_CheckBoxDBID[formdiv.CheckBoxs[j]]

		checkboxDB_i := backRepo.BackRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB[checkboxDB_i_ID]
		checkboxDB_j := backRepo.BackRepoCheckBox.Map_CheckBoxDBID_CheckBoxDB[checkboxDB_j_ID]

		return checkboxDB_i.FormDiv_CheckBoxsDBID_Index.Int64 < checkboxDB_j.FormDiv_CheckBoxsDBID_Index.Int64
	})

	// FormEditAssocButton field
	formdiv.FormEditAssocButton = nil
	if formdivDB.FormEditAssocButtonID.Int64 != 0 {
		formdiv.FormEditAssocButton = backRepo.BackRepoFormEditAssocButton.Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr[uint(formdivDB.FormEditAssocButtonID.Int64)]
	}
	// FormSortAssocButton field
	formdiv.FormSortAssocButton = nil
	if formdivDB.FormSortAssocButtonID.Int64 != 0 {
		formdiv.FormSortAssocButton = backRepo.BackRepoFormSortAssocButton.Map_FormSortAssocButtonDBID_FormSortAssocButtonPtr[uint(formdivDB.FormSortAssocButtonID.Int64)]
	}
	return
}

// CommitFormDiv allows commit of a single formdiv (if already staged)
func (backRepo *BackRepoStruct) CommitFormDiv(formdiv *models.FormDiv) {
	backRepo.BackRepoFormDiv.CommitPhaseOneInstance(formdiv)
	if id, ok := backRepo.BackRepoFormDiv.Map_FormDivPtr_FormDivDBID[formdiv]; ok {
		backRepo.BackRepoFormDiv.CommitPhaseTwoInstance(backRepo, id, formdiv)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitFormDiv allows checkout of a single formdiv (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutFormDiv(formdiv *models.FormDiv) {
	// check if the formdiv is staged
	if _, ok := backRepo.BackRepoFormDiv.Map_FormDivPtr_FormDivDBID[formdiv]; ok {

		if id, ok := backRepo.BackRepoFormDiv.Map_FormDivPtr_FormDivDBID[formdiv]; ok {
			var formdivDB FormDivDB
			formdivDB.ID = id

			if err := backRepo.BackRepoFormDiv.db.First(&formdivDB, id).Error; err != nil {
				log.Panicln("CheckoutFormDiv : Problem with getting object with id:", id)
			}
			backRepo.BackRepoFormDiv.CheckoutPhaseOneInstance(&formdivDB)
			backRepo.BackRepoFormDiv.CheckoutPhaseTwoInstance(backRepo, &formdivDB)
		}
	}
}

// CopyBasicFieldsFromFormDiv
func (formdivDB *FormDivDB) CopyBasicFieldsFromFormDiv(formdiv *models.FormDiv) {
	// insertion point for fields commit

	formdivDB.Name_Data.String = formdiv.Name
	formdivDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromFormDivWOP
func (formdivDB *FormDivDB) CopyBasicFieldsFromFormDivWOP(formdiv *FormDivWOP) {
	// insertion point for fields commit

	formdivDB.Name_Data.String = formdiv.Name
	formdivDB.Name_Data.Valid = true
}

// CopyBasicFieldsToFormDiv
func (formdivDB *FormDivDB) CopyBasicFieldsToFormDiv(formdiv *models.FormDiv) {
	// insertion point for checkout of basic fields (back repo to stage)
	formdiv.Name = formdivDB.Name_Data.String
}

// CopyBasicFieldsToFormDivWOP
func (formdivDB *FormDivDB) CopyBasicFieldsToFormDivWOP(formdiv *FormDivWOP) {
	formdiv.ID = int(formdivDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	formdiv.Name = formdivDB.Name_Data.String
}

// Backup generates a json file from a slice of all FormDivDB instances in the backrepo
func (backRepoFormDiv *BackRepoFormDivStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "FormDivDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FormDivDB, 0)
	for _, formdivDB := range backRepoFormDiv.Map_FormDivDBID_FormDivDB {
		forBackup = append(forBackup, formdivDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json FormDiv ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json FormDiv file", err.Error())
	}
}

// Backup generates a json file from a slice of all FormDivDB instances in the backrepo
func (backRepoFormDiv *BackRepoFormDivStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FormDivDB, 0)
	for _, formdivDB := range backRepoFormDiv.Map_FormDivDBID_FormDivDB {
		forBackup = append(forBackup, formdivDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("FormDiv")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&FormDiv_Fields, -1)
	for _, formdivDB := range forBackup {

		var formdivWOP FormDivWOP
		formdivDB.CopyBasicFieldsToFormDivWOP(&formdivWOP)

		row := sh.AddRow()
		row.WriteStruct(&formdivWOP, -1)
	}
}

// RestoreXL from the "FormDiv" sheet all FormDivDB instances
func (backRepoFormDiv *BackRepoFormDivStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoFormDivid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["FormDiv"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoFormDiv.rowVisitorFormDiv)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoFormDiv *BackRepoFormDivStruct) rowVisitorFormDiv(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var formdivWOP FormDivWOP
		row.ReadStruct(&formdivWOP)

		// add the unmarshalled struct to the stage
		formdivDB := new(FormDivDB)
		formdivDB.CopyBasicFieldsFromFormDivWOP(&formdivWOP)

		formdivDB_ID_atBackupTime := formdivDB.ID
		formdivDB.ID = 0
		query := backRepoFormDiv.db.Create(formdivDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoFormDiv.Map_FormDivDBID_FormDivDB[formdivDB.ID] = formdivDB
		BackRepoFormDivid_atBckpTime_newID[formdivDB_ID_atBackupTime] = formdivDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "FormDivDB.json" in dirPath that stores an array
// of FormDivDB and stores it in the database
// the map BackRepoFormDivid_atBckpTime_newID is updated accordingly
func (backRepoFormDiv *BackRepoFormDivStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoFormDivid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "FormDivDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json FormDiv file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*FormDivDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_FormDivDBID_FormDivDB
	for _, formdivDB := range forRestore {

		formdivDB_ID_atBackupTime := formdivDB.ID
		formdivDB.ID = 0
		query := backRepoFormDiv.db.Create(formdivDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoFormDiv.Map_FormDivDBID_FormDivDB[formdivDB.ID] = formdivDB
		BackRepoFormDivid_atBckpTime_newID[formdivDB_ID_atBackupTime] = formdivDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json FormDiv file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<FormDiv>id_atBckpTime_newID
// to compute new index
func (backRepoFormDiv *BackRepoFormDivStruct) RestorePhaseTwo() {

	for _, formdivDB := range backRepoFormDiv.Map_FormDivDBID_FormDivDB {

		// next line of code is to avert unused variable compilation error
		_ = formdivDB

		// insertion point for reindexing pointers encoding
		// reindexing FormEditAssocButton field
		if formdivDB.FormEditAssocButtonID.Int64 != 0 {
			formdivDB.FormEditAssocButtonID.Int64 = int64(BackRepoFormEditAssocButtonid_atBckpTime_newID[uint(formdivDB.FormEditAssocButtonID.Int64)])
			formdivDB.FormEditAssocButtonID.Valid = true
		}

		// reindexing FormSortAssocButton field
		if formdivDB.FormSortAssocButtonID.Int64 != 0 {
			formdivDB.FormSortAssocButtonID.Int64 = int64(BackRepoFormSortAssocButtonid_atBckpTime_newID[uint(formdivDB.FormSortAssocButtonID.Int64)])
			formdivDB.FormSortAssocButtonID.Valid = true
		}

		// This reindex formdiv.FormDivs
		if formdivDB.FormGroup_FormDivsDBID.Int64 != 0 {
			formdivDB.FormGroup_FormDivsDBID.Int64 =
				int64(BackRepoFormGroupid_atBckpTime_newID[uint(formdivDB.FormGroup_FormDivsDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoFormDiv.db.Model(formdivDB).Updates(*formdivDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// BackRepoFormDiv.ResetReversePointers commits all staged instances of FormDiv to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormDiv *BackRepoFormDivStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, formdiv := range backRepoFormDiv.Map_FormDivDBID_FormDivPtr {
		backRepoFormDiv.ResetReversePointersInstance(backRepo, idx, formdiv)
	}

	return
}

func (backRepoFormDiv *BackRepoFormDivStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, astruct *models.FormDiv) (Error error) {

	// fetch matching formdivDB
	if formdivDB, ok := backRepoFormDiv.Map_FormDivDBID_FormDivDB[idx]; ok {
		_ = formdivDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		if formdivDB.FormGroup_FormDivsDBID.Int64 != 0 {
			formdivDB.FormGroup_FormDivsDBID.Int64 = 0
			formdivDB.FormGroup_FormDivsDBID.Valid = true

			// save the reset
			if q := backRepoFormDiv.db.Save(formdivDB); q.Error != nil {
				return q.Error
			}
		}
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoFormDivid_atBckpTime_newID map[uint]uint
