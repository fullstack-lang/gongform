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
var dummy_FormGroup_sql sql.NullBool
var dummy_FormGroup_time time.Duration
var dummy_FormGroup_sort sort.Float64Slice

// FormGroupAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model formgroupAPI
type FormGroupAPI struct {
	gorm.Model

	models.FormGroup

	// encoding of pointers
	FormGroupPointersEnconding
}

// FormGroupPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type FormGroupPointersEnconding struct {
	// insertion for pointer fields encoding declaration
}

// FormGroupDB describes a formgroup in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model formgroupDB
type FormGroupDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field formgroupDB.Name
	Name_Data sql.NullString

	// Declation for basic field formgroupDB.Label
	Label_Data sql.NullString
	// encoding of pointers
	FormGroupPointersEnconding
}

// FormGroupDBs arrays formgroupDBs
// swagger:response formgroupDBsResponse
type FormGroupDBs []FormGroupDB

// FormGroupDBResponse provides response
// swagger:response formgroupDBResponse
type FormGroupDBResponse struct {
	FormGroupDB
}

// FormGroupWOP is a FormGroup without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type FormGroupWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Label string `xlsx:"2"`
	// insertion for WOP pointer fields
}

var FormGroup_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Label",
}

type BackRepoFormGroupStruct struct {
	// stores FormGroupDB according to their gorm ID
	Map_FormGroupDBID_FormGroupDB map[uint]*FormGroupDB

	// stores FormGroupDB ID according to FormGroup address
	Map_FormGroupPtr_FormGroupDBID map[*models.FormGroup]uint

	// stores FormGroup according to their gorm ID
	Map_FormGroupDBID_FormGroupPtr map[uint]*models.FormGroup

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoFormGroup *BackRepoFormGroupStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoFormGroup.stage
	return
}

func (backRepoFormGroup *BackRepoFormGroupStruct) GetDB() *gorm.DB {
	return backRepoFormGroup.db
}

// GetFormGroupDBFromFormGroupPtr is a handy function to access the back repo instance from the stage instance
func (backRepoFormGroup *BackRepoFormGroupStruct) GetFormGroupDBFromFormGroupPtr(formgroup *models.FormGroup) (formgroupDB *FormGroupDB) {
	id := backRepoFormGroup.Map_FormGroupPtr_FormGroupDBID[formgroup]
	formgroupDB = backRepoFormGroup.Map_FormGroupDBID_FormGroupDB[id]
	return
}

// BackRepoFormGroup.CommitPhaseOne commits all staged instances of FormGroup to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormGroup *BackRepoFormGroupStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for formgroup := range stage.FormGroups {
		backRepoFormGroup.CommitPhaseOneInstance(formgroup)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, formgroup := range backRepoFormGroup.Map_FormGroupDBID_FormGroupPtr {
		if _, ok := stage.FormGroups[formgroup]; !ok {
			backRepoFormGroup.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoFormGroup.CommitDeleteInstance commits deletion of FormGroup to the BackRepo
func (backRepoFormGroup *BackRepoFormGroupStruct) CommitDeleteInstance(id uint) (Error error) {

	formgroup := backRepoFormGroup.Map_FormGroupDBID_FormGroupPtr[id]

	// formgroup is not staged anymore, remove formgroupDB
	formgroupDB := backRepoFormGroup.Map_FormGroupDBID_FormGroupDB[id]
	query := backRepoFormGroup.db.Unscoped().Delete(&formgroupDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoFormGroup.Map_FormGroupPtr_FormGroupDBID, formgroup)
	delete(backRepoFormGroup.Map_FormGroupDBID_FormGroupPtr, id)
	delete(backRepoFormGroup.Map_FormGroupDBID_FormGroupDB, id)

	return
}

// BackRepoFormGroup.CommitPhaseOneInstance commits formgroup staged instances of FormGroup to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormGroup *BackRepoFormGroupStruct) CommitPhaseOneInstance(formgroup *models.FormGroup) (Error error) {

	// check if the formgroup is not commited yet
	if _, ok := backRepoFormGroup.Map_FormGroupPtr_FormGroupDBID[formgroup]; ok {
		return
	}

	// initiate formgroup
	var formgroupDB FormGroupDB
	formgroupDB.CopyBasicFieldsFromFormGroup(formgroup)

	query := backRepoFormGroup.db.Create(&formgroupDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoFormGroup.Map_FormGroupPtr_FormGroupDBID[formgroup] = formgroupDB.ID
	backRepoFormGroup.Map_FormGroupDBID_FormGroupPtr[formgroupDB.ID] = formgroup
	backRepoFormGroup.Map_FormGroupDBID_FormGroupDB[formgroupDB.ID] = &formgroupDB

	return
}

// BackRepoFormGroup.CommitPhaseTwo commits all staged instances of FormGroup to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormGroup *BackRepoFormGroupStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, formgroup := range backRepoFormGroup.Map_FormGroupDBID_FormGroupPtr {
		backRepoFormGroup.CommitPhaseTwoInstance(backRepo, idx, formgroup)
	}

	return
}

// BackRepoFormGroup.CommitPhaseTwoInstance commits {{structname }} of models.FormGroup to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormGroup *BackRepoFormGroupStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, formgroup *models.FormGroup) (Error error) {

	// fetch matching formgroupDB
	if formgroupDB, ok := backRepoFormGroup.Map_FormGroupDBID_FormGroupDB[idx]; ok {

		formgroupDB.CopyBasicFieldsFromFormGroup(formgroup)

		// insertion point for translating pointers encodings into actual pointers
		// This loop encodes the slice of pointers formgroup.FormDivs into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, formdivAssocEnd := range formgroup.FormDivs {

			// get the back repo instance at the association end
			formdivAssocEnd_DB :=
				backRepo.BackRepoFormDiv.GetFormDivDBFromFormDivPtr(formdivAssocEnd)

			// encode reverse pointer in the association end back repo instance
			formdivAssocEnd_DB.FormGroup_FormDivsDBID.Int64 = int64(formgroupDB.ID)
			formdivAssocEnd_DB.FormGroup_FormDivsDBID.Valid = true
			formdivAssocEnd_DB.FormGroup_FormDivsDBID_Index.Int64 = int64(idx)
			formdivAssocEnd_DB.FormGroup_FormDivsDBID_Index.Valid = true
			if q := backRepoFormGroup.db.Save(formdivAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		query := backRepoFormGroup.db.Save(&formgroupDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown FormGroup intance %s", formgroup.Name))
		return err
	}

	return
}

// BackRepoFormGroup.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoFormGroup *BackRepoFormGroupStruct) CheckoutPhaseOne() (Error error) {

	formgroupDBArray := make([]FormGroupDB, 0)
	query := backRepoFormGroup.db.Find(&formgroupDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	formgroupInstancesToBeRemovedFromTheStage := make(map[*models.FormGroup]any)
	for key, value := range backRepoFormGroup.stage.FormGroups {
		formgroupInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, formgroupDB := range formgroupDBArray {
		backRepoFormGroup.CheckoutPhaseOneInstance(&formgroupDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		formgroup, ok := backRepoFormGroup.Map_FormGroupDBID_FormGroupPtr[formgroupDB.ID]
		if ok {
			delete(formgroupInstancesToBeRemovedFromTheStage, formgroup)
		}
	}

	// remove from stage and back repo's 3 maps all formgroups that are not in the checkout
	for formgroup := range formgroupInstancesToBeRemovedFromTheStage {
		formgroup.Unstage(backRepoFormGroup.GetStage())

		// remove instance from the back repo 3 maps
		formgroupID := backRepoFormGroup.Map_FormGroupPtr_FormGroupDBID[formgroup]
		delete(backRepoFormGroup.Map_FormGroupPtr_FormGroupDBID, formgroup)
		delete(backRepoFormGroup.Map_FormGroupDBID_FormGroupDB, formgroupID)
		delete(backRepoFormGroup.Map_FormGroupDBID_FormGroupPtr, formgroupID)
	}

	return
}

// CheckoutPhaseOneInstance takes a formgroupDB that has been found in the DB, updates the backRepo and stages the
// models version of the formgroupDB
func (backRepoFormGroup *BackRepoFormGroupStruct) CheckoutPhaseOneInstance(formgroupDB *FormGroupDB) (Error error) {

	formgroup, ok := backRepoFormGroup.Map_FormGroupDBID_FormGroupPtr[formgroupDB.ID]
	if !ok {
		formgroup = new(models.FormGroup)

		backRepoFormGroup.Map_FormGroupDBID_FormGroupPtr[formgroupDB.ID] = formgroup
		backRepoFormGroup.Map_FormGroupPtr_FormGroupDBID[formgroup] = formgroupDB.ID

		// append model store with the new element
		formgroup.Name = formgroupDB.Name_Data.String
		formgroup.Stage(backRepoFormGroup.GetStage())
	}
	formgroupDB.CopyBasicFieldsToFormGroup(formgroup)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	formgroup.Stage(backRepoFormGroup.GetStage())

	// preserve pointer to formgroupDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_FormGroupDBID_FormGroupDB)[formgroupDB hold variable pointers
	formgroupDB_Data := *formgroupDB
	preservedPtrToFormGroup := &formgroupDB_Data
	backRepoFormGroup.Map_FormGroupDBID_FormGroupDB[formgroupDB.ID] = preservedPtrToFormGroup

	return
}

// BackRepoFormGroup.CheckoutPhaseTwo Checkouts all staged instances of FormGroup to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormGroup *BackRepoFormGroupStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, formgroupDB := range backRepoFormGroup.Map_FormGroupDBID_FormGroupDB {
		backRepoFormGroup.CheckoutPhaseTwoInstance(backRepo, formgroupDB)
	}
	return
}

// BackRepoFormGroup.CheckoutPhaseTwoInstance Checkouts staged instances of FormGroup to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormGroup *BackRepoFormGroupStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, formgroupDB *FormGroupDB) (Error error) {

	formgroup := backRepoFormGroup.Map_FormGroupDBID_FormGroupPtr[formgroupDB.ID]
	_ = formgroup // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// This loop redeem formgroup.FormDivs in the stage from the encode in the back repo
	// It parses all FormDivDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	formgroup.FormDivs = formgroup.FormDivs[:0]
	// 2. loop all instances in the type in the association end
	for _, formdivDB_AssocEnd := range backRepo.BackRepoFormDiv.Map_FormDivDBID_FormDivDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if formdivDB_AssocEnd.FormGroup_FormDivsDBID.Int64 == int64(formgroupDB.ID) {
			// 4. fetch the associated instance in the stage
			formdiv_AssocEnd := backRepo.BackRepoFormDiv.Map_FormDivDBID_FormDivPtr[formdivDB_AssocEnd.ID]
			// 5. append it the association slice
			formgroup.FormDivs = append(formgroup.FormDivs, formdiv_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(formgroup.FormDivs, func(i, j int) bool {
		formdivDB_i_ID := backRepo.BackRepoFormDiv.Map_FormDivPtr_FormDivDBID[formgroup.FormDivs[i]]
		formdivDB_j_ID := backRepo.BackRepoFormDiv.Map_FormDivPtr_FormDivDBID[formgroup.FormDivs[j]]

		formdivDB_i := backRepo.BackRepoFormDiv.Map_FormDivDBID_FormDivDB[formdivDB_i_ID]
		formdivDB_j := backRepo.BackRepoFormDiv.Map_FormDivDBID_FormDivDB[formdivDB_j_ID]

		return formdivDB_i.FormGroup_FormDivsDBID_Index.Int64 < formdivDB_j.FormGroup_FormDivsDBID_Index.Int64
	})

	return
}

// CommitFormGroup allows commit of a single formgroup (if already staged)
func (backRepo *BackRepoStruct) CommitFormGroup(formgroup *models.FormGroup) {
	backRepo.BackRepoFormGroup.CommitPhaseOneInstance(formgroup)
	if id, ok := backRepo.BackRepoFormGroup.Map_FormGroupPtr_FormGroupDBID[formgroup]; ok {
		backRepo.BackRepoFormGroup.CommitPhaseTwoInstance(backRepo, id, formgroup)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitFormGroup allows checkout of a single formgroup (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutFormGroup(formgroup *models.FormGroup) {
	// check if the formgroup is staged
	if _, ok := backRepo.BackRepoFormGroup.Map_FormGroupPtr_FormGroupDBID[formgroup]; ok {

		if id, ok := backRepo.BackRepoFormGroup.Map_FormGroupPtr_FormGroupDBID[formgroup]; ok {
			var formgroupDB FormGroupDB
			formgroupDB.ID = id

			if err := backRepo.BackRepoFormGroup.db.First(&formgroupDB, id).Error; err != nil {
				log.Panicln("CheckoutFormGroup : Problem with getting object with id:", id)
			}
			backRepo.BackRepoFormGroup.CheckoutPhaseOneInstance(&formgroupDB)
			backRepo.BackRepoFormGroup.CheckoutPhaseTwoInstance(backRepo, &formgroupDB)
		}
	}
}

// CopyBasicFieldsFromFormGroup
func (formgroupDB *FormGroupDB) CopyBasicFieldsFromFormGroup(formgroup *models.FormGroup) {
	// insertion point for fields commit

	formgroupDB.Name_Data.String = formgroup.Name
	formgroupDB.Name_Data.Valid = true

	formgroupDB.Label_Data.String = formgroup.Label
	formgroupDB.Label_Data.Valid = true
}

// CopyBasicFieldsFromFormGroupWOP
func (formgroupDB *FormGroupDB) CopyBasicFieldsFromFormGroupWOP(formgroup *FormGroupWOP) {
	// insertion point for fields commit

	formgroupDB.Name_Data.String = formgroup.Name
	formgroupDB.Name_Data.Valid = true

	formgroupDB.Label_Data.String = formgroup.Label
	formgroupDB.Label_Data.Valid = true
}

// CopyBasicFieldsToFormGroup
func (formgroupDB *FormGroupDB) CopyBasicFieldsToFormGroup(formgroup *models.FormGroup) {
	// insertion point for checkout of basic fields (back repo to stage)
	formgroup.Name = formgroupDB.Name_Data.String
	formgroup.Label = formgroupDB.Label_Data.String
}

// CopyBasicFieldsToFormGroupWOP
func (formgroupDB *FormGroupDB) CopyBasicFieldsToFormGroupWOP(formgroup *FormGroupWOP) {
	formgroup.ID = int(formgroupDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	formgroup.Name = formgroupDB.Name_Data.String
	formgroup.Label = formgroupDB.Label_Data.String
}

// Backup generates a json file from a slice of all FormGroupDB instances in the backrepo
func (backRepoFormGroup *BackRepoFormGroupStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "FormGroupDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FormGroupDB, 0)
	for _, formgroupDB := range backRepoFormGroup.Map_FormGroupDBID_FormGroupDB {
		forBackup = append(forBackup, formgroupDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json FormGroup ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json FormGroup file", err.Error())
	}
}

// Backup generates a json file from a slice of all FormGroupDB instances in the backrepo
func (backRepoFormGroup *BackRepoFormGroupStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FormGroupDB, 0)
	for _, formgroupDB := range backRepoFormGroup.Map_FormGroupDBID_FormGroupDB {
		forBackup = append(forBackup, formgroupDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("FormGroup")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&FormGroup_Fields, -1)
	for _, formgroupDB := range forBackup {

		var formgroupWOP FormGroupWOP
		formgroupDB.CopyBasicFieldsToFormGroupWOP(&formgroupWOP)

		row := sh.AddRow()
		row.WriteStruct(&formgroupWOP, -1)
	}
}

// RestoreXL from the "FormGroup" sheet all FormGroupDB instances
func (backRepoFormGroup *BackRepoFormGroupStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoFormGroupid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["FormGroup"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoFormGroup.rowVisitorFormGroup)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoFormGroup *BackRepoFormGroupStruct) rowVisitorFormGroup(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var formgroupWOP FormGroupWOP
		row.ReadStruct(&formgroupWOP)

		// add the unmarshalled struct to the stage
		formgroupDB := new(FormGroupDB)
		formgroupDB.CopyBasicFieldsFromFormGroupWOP(&formgroupWOP)

		formgroupDB_ID_atBackupTime := formgroupDB.ID
		formgroupDB.ID = 0
		query := backRepoFormGroup.db.Create(formgroupDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoFormGroup.Map_FormGroupDBID_FormGroupDB[formgroupDB.ID] = formgroupDB
		BackRepoFormGroupid_atBckpTime_newID[formgroupDB_ID_atBackupTime] = formgroupDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "FormGroupDB.json" in dirPath that stores an array
// of FormGroupDB and stores it in the database
// the map BackRepoFormGroupid_atBckpTime_newID is updated accordingly
func (backRepoFormGroup *BackRepoFormGroupStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoFormGroupid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "FormGroupDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json FormGroup file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*FormGroupDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_FormGroupDBID_FormGroupDB
	for _, formgroupDB := range forRestore {

		formgroupDB_ID_atBackupTime := formgroupDB.ID
		formgroupDB.ID = 0
		query := backRepoFormGroup.db.Create(formgroupDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoFormGroup.Map_FormGroupDBID_FormGroupDB[formgroupDB.ID] = formgroupDB
		BackRepoFormGroupid_atBckpTime_newID[formgroupDB_ID_atBackupTime] = formgroupDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json FormGroup file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<FormGroup>id_atBckpTime_newID
// to compute new index
func (backRepoFormGroup *BackRepoFormGroupStruct) RestorePhaseTwo() {

	for _, formgroupDB := range backRepoFormGroup.Map_FormGroupDBID_FormGroupDB {

		// next line of code is to avert unused variable compilation error
		_ = formgroupDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoFormGroup.db.Model(formgroupDB).Updates(*formgroupDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// BackRepoFormGroup.ResetReversePointers commits all staged instances of FormGroup to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormGroup *BackRepoFormGroupStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, formgroup := range backRepoFormGroup.Map_FormGroupDBID_FormGroupPtr {
		backRepoFormGroup.ResetReversePointersInstance(backRepo, idx, formgroup)
	}

	return
}

func (backRepoFormGroup *BackRepoFormGroupStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, astruct *models.FormGroup) (Error error) {

	// fetch matching formgroupDB
	if formgroupDB, ok := backRepoFormGroup.Map_FormGroupDBID_FormGroupDB[idx]; ok {
		_ = formgroupDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoFormGroupid_atBckpTime_newID map[uint]uint
