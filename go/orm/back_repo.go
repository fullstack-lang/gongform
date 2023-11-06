// do not modify, generated file
package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/fullstack-lang/gongform/go/models"

	"github.com/tealeg/xlsx/v3"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoCheckBox BackRepoCheckBoxStruct

	BackRepoFormDiv BackRepoFormDivStruct

	BackRepoFormEditAssocButton BackRepoFormEditAssocButtonStruct

	BackRepoFormField BackRepoFormFieldStruct

	BackRepoFormFieldDate BackRepoFormFieldDateStruct

	BackRepoFormFieldDateTime BackRepoFormFieldDateTimeStruct

	BackRepoFormFieldFloat64 BackRepoFormFieldFloat64Struct

	BackRepoFormFieldInt BackRepoFormFieldIntStruct

	BackRepoFormFieldSelect BackRepoFormFieldSelectStruct

	BackRepoFormFieldString BackRepoFormFieldStringStruct

	BackRepoFormFieldTime BackRepoFormFieldTimeStruct

	BackRepoFormGroup BackRepoFormGroupStruct

	BackRepoFormSortAssocButton BackRepoFormSortAssocButtonStruct

	BackRepoOption BackRepoOptionStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filename), gormConfig)

	// since testsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	if err != nil {
		panic("Failed to connect to database!")
	}

	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
	}

	err = db.AutoMigrate( // insertion point for reference to structs
		&CheckBoxDB{},
		&FormDivDB{},
		&FormEditAssocButtonDB{},
		&FormFieldDB{},
		&FormFieldDateDB{},
		&FormFieldDateTimeDB{},
		&FormFieldFloat64DB{},
		&FormFieldIntDB{},
		&FormFieldSelectDB{},
		&FormFieldStringDB{},
		&FormFieldTimeDB{},
		&FormGroupDB{},
		&FormSortAssocButtonDB{},
		&OptionDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/test/go")
	}

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoCheckBox = BackRepoCheckBoxStruct{
		Map_CheckBoxDBID_CheckBoxPtr: make(map[uint]*models.CheckBox, 0),
		Map_CheckBoxDBID_CheckBoxDB:  make(map[uint]*CheckBoxDB, 0),
		Map_CheckBoxPtr_CheckBoxDBID: make(map[*models.CheckBox]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFormDiv = BackRepoFormDivStruct{
		Map_FormDivDBID_FormDivPtr: make(map[uint]*models.FormDiv, 0),
		Map_FormDivDBID_FormDivDB:  make(map[uint]*FormDivDB, 0),
		Map_FormDivPtr_FormDivDBID: make(map[*models.FormDiv]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFormEditAssocButton = BackRepoFormEditAssocButtonStruct{
		Map_FormEditAssocButtonDBID_FormEditAssocButtonPtr: make(map[uint]*models.FormEditAssocButton, 0),
		Map_FormEditAssocButtonDBID_FormEditAssocButtonDB:  make(map[uint]*FormEditAssocButtonDB, 0),
		Map_FormEditAssocButtonPtr_FormEditAssocButtonDBID: make(map[*models.FormEditAssocButton]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFormField = BackRepoFormFieldStruct{
		Map_FormFieldDBID_FormFieldPtr: make(map[uint]*models.FormField, 0),
		Map_FormFieldDBID_FormFieldDB:  make(map[uint]*FormFieldDB, 0),
		Map_FormFieldPtr_FormFieldDBID: make(map[*models.FormField]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFormFieldDate = BackRepoFormFieldDateStruct{
		Map_FormFieldDateDBID_FormFieldDatePtr: make(map[uint]*models.FormFieldDate, 0),
		Map_FormFieldDateDBID_FormFieldDateDB:  make(map[uint]*FormFieldDateDB, 0),
		Map_FormFieldDatePtr_FormFieldDateDBID: make(map[*models.FormFieldDate]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFormFieldDateTime = BackRepoFormFieldDateTimeStruct{
		Map_FormFieldDateTimeDBID_FormFieldDateTimePtr: make(map[uint]*models.FormFieldDateTime, 0),
		Map_FormFieldDateTimeDBID_FormFieldDateTimeDB:  make(map[uint]*FormFieldDateTimeDB, 0),
		Map_FormFieldDateTimePtr_FormFieldDateTimeDBID: make(map[*models.FormFieldDateTime]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFormFieldFloat64 = BackRepoFormFieldFloat64Struct{
		Map_FormFieldFloat64DBID_FormFieldFloat64Ptr: make(map[uint]*models.FormFieldFloat64, 0),
		Map_FormFieldFloat64DBID_FormFieldFloat64DB:  make(map[uint]*FormFieldFloat64DB, 0),
		Map_FormFieldFloat64Ptr_FormFieldFloat64DBID: make(map[*models.FormFieldFloat64]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFormFieldInt = BackRepoFormFieldIntStruct{
		Map_FormFieldIntDBID_FormFieldIntPtr: make(map[uint]*models.FormFieldInt, 0),
		Map_FormFieldIntDBID_FormFieldIntDB:  make(map[uint]*FormFieldIntDB, 0),
		Map_FormFieldIntPtr_FormFieldIntDBID: make(map[*models.FormFieldInt]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFormFieldSelect = BackRepoFormFieldSelectStruct{
		Map_FormFieldSelectDBID_FormFieldSelectPtr: make(map[uint]*models.FormFieldSelect, 0),
		Map_FormFieldSelectDBID_FormFieldSelectDB:  make(map[uint]*FormFieldSelectDB, 0),
		Map_FormFieldSelectPtr_FormFieldSelectDBID: make(map[*models.FormFieldSelect]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFormFieldString = BackRepoFormFieldStringStruct{
		Map_FormFieldStringDBID_FormFieldStringPtr: make(map[uint]*models.FormFieldString, 0),
		Map_FormFieldStringDBID_FormFieldStringDB:  make(map[uint]*FormFieldStringDB, 0),
		Map_FormFieldStringPtr_FormFieldStringDBID: make(map[*models.FormFieldString]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFormFieldTime = BackRepoFormFieldTimeStruct{
		Map_FormFieldTimeDBID_FormFieldTimePtr: make(map[uint]*models.FormFieldTime, 0),
		Map_FormFieldTimeDBID_FormFieldTimeDB:  make(map[uint]*FormFieldTimeDB, 0),
		Map_FormFieldTimePtr_FormFieldTimeDBID: make(map[*models.FormFieldTime]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFormGroup = BackRepoFormGroupStruct{
		Map_FormGroupDBID_FormGroupPtr: make(map[uint]*models.FormGroup, 0),
		Map_FormGroupDBID_FormGroupDB:  make(map[uint]*FormGroupDB, 0),
		Map_FormGroupPtr_FormGroupDBID: make(map[*models.FormGroup]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFormSortAssocButton = BackRepoFormSortAssocButtonStruct{
		Map_FormSortAssocButtonDBID_FormSortAssocButtonPtr: make(map[uint]*models.FormSortAssocButton, 0),
		Map_FormSortAssocButtonDBID_FormSortAssocButtonDB:  make(map[uint]*FormSortAssocButtonDB, 0),
		Map_FormSortAssocButtonPtr_FormSortAssocButtonDBID: make(map[*models.FormSortAssocButton]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoOption = BackRepoOptionStruct{
		Map_OptionDBID_OptionPtr: make(map[uint]*models.Option, 0),
		Map_OptionDBID_OptionDB:  make(map[uint]*OptionDB, 0),
		Map_OptionPtr_OptionDBID: make(map[*models.Option]uint, 0),

		db:    db,
		stage: stage,
	}

	stage.BackRepo = backRepo
	backRepo.stage = stage

	return
}

func (backRepo *BackRepoStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepo.stage
	return
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromBackCallback != nil {
		backRepo.stage.OnInitCommitFromBackCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromFrontCallback != nil {
		backRepo.stage.OnInitCommitFromFrontCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoCheckBox.CommitPhaseOne(stage)
	backRepo.BackRepoFormDiv.CommitPhaseOne(stage)
	backRepo.BackRepoFormEditAssocButton.CommitPhaseOne(stage)
	backRepo.BackRepoFormField.CommitPhaseOne(stage)
	backRepo.BackRepoFormFieldDate.CommitPhaseOne(stage)
	backRepo.BackRepoFormFieldDateTime.CommitPhaseOne(stage)
	backRepo.BackRepoFormFieldFloat64.CommitPhaseOne(stage)
	backRepo.BackRepoFormFieldInt.CommitPhaseOne(stage)
	backRepo.BackRepoFormFieldSelect.CommitPhaseOne(stage)
	backRepo.BackRepoFormFieldString.CommitPhaseOne(stage)
	backRepo.BackRepoFormFieldTime.CommitPhaseOne(stage)
	backRepo.BackRepoFormGroup.CommitPhaseOne(stage)
	backRepo.BackRepoFormSortAssocButton.CommitPhaseOne(stage)
	backRepo.BackRepoOption.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoCheckBox.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFormDiv.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFormEditAssocButton.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFormField.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFormFieldDate.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFormFieldDateTime.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFormFieldFloat64.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFormFieldInt.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFormFieldSelect.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFormFieldString.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFormFieldTime.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFormGroup.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFormSortAssocButton.CommitPhaseTwo(backRepo)
	backRepo.BackRepoOption.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoCheckBox.CheckoutPhaseOne()
	backRepo.BackRepoFormDiv.CheckoutPhaseOne()
	backRepo.BackRepoFormEditAssocButton.CheckoutPhaseOne()
	backRepo.BackRepoFormField.CheckoutPhaseOne()
	backRepo.BackRepoFormFieldDate.CheckoutPhaseOne()
	backRepo.BackRepoFormFieldDateTime.CheckoutPhaseOne()
	backRepo.BackRepoFormFieldFloat64.CheckoutPhaseOne()
	backRepo.BackRepoFormFieldInt.CheckoutPhaseOne()
	backRepo.BackRepoFormFieldSelect.CheckoutPhaseOne()
	backRepo.BackRepoFormFieldString.CheckoutPhaseOne()
	backRepo.BackRepoFormFieldTime.CheckoutPhaseOne()
	backRepo.BackRepoFormGroup.CheckoutPhaseOne()
	backRepo.BackRepoFormSortAssocButton.CheckoutPhaseOne()
	backRepo.BackRepoOption.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoCheckBox.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFormDiv.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFormEditAssocButton.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFormField.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFormFieldDate.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFormFieldDateTime.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFormFieldFloat64.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFormFieldInt.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFormFieldSelect.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFormFieldString.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFormFieldTime.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFormGroup.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFormSortAssocButton.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoOption.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoCheckBox.Backup(dirPath)
	backRepo.BackRepoFormDiv.Backup(dirPath)
	backRepo.BackRepoFormEditAssocButton.Backup(dirPath)
	backRepo.BackRepoFormField.Backup(dirPath)
	backRepo.BackRepoFormFieldDate.Backup(dirPath)
	backRepo.BackRepoFormFieldDateTime.Backup(dirPath)
	backRepo.BackRepoFormFieldFloat64.Backup(dirPath)
	backRepo.BackRepoFormFieldInt.Backup(dirPath)
	backRepo.BackRepoFormFieldSelect.Backup(dirPath)
	backRepo.BackRepoFormFieldString.Backup(dirPath)
	backRepo.BackRepoFormFieldTime.Backup(dirPath)
	backRepo.BackRepoFormGroup.Backup(dirPath)
	backRepo.BackRepoFormSortAssocButton.Backup(dirPath)
	backRepo.BackRepoOption.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoCheckBox.BackupXL(file)
	backRepo.BackRepoFormDiv.BackupXL(file)
	backRepo.BackRepoFormEditAssocButton.BackupXL(file)
	backRepo.BackRepoFormField.BackupXL(file)
	backRepo.BackRepoFormFieldDate.BackupXL(file)
	backRepo.BackRepoFormFieldDateTime.BackupXL(file)
	backRepo.BackRepoFormFieldFloat64.BackupXL(file)
	backRepo.BackRepoFormFieldInt.BackupXL(file)
	backRepo.BackRepoFormFieldSelect.BackupXL(file)
	backRepo.BackRepoFormFieldString.BackupXL(file)
	backRepo.BackRepoFormFieldTime.BackupXL(file)
	backRepo.BackRepoFormGroup.BackupXL(file)
	backRepo.BackRepoFormSortAssocButton.BackupXL(file)
	backRepo.BackRepoOption.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	backRepo.stage.Commit()
	backRepo.stage.Reset()
	backRepo.stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoCheckBox.RestorePhaseOne(dirPath)
	backRepo.BackRepoFormDiv.RestorePhaseOne(dirPath)
	backRepo.BackRepoFormEditAssocButton.RestorePhaseOne(dirPath)
	backRepo.BackRepoFormField.RestorePhaseOne(dirPath)
	backRepo.BackRepoFormFieldDate.RestorePhaseOne(dirPath)
	backRepo.BackRepoFormFieldDateTime.RestorePhaseOne(dirPath)
	backRepo.BackRepoFormFieldFloat64.RestorePhaseOne(dirPath)
	backRepo.BackRepoFormFieldInt.RestorePhaseOne(dirPath)
	backRepo.BackRepoFormFieldSelect.RestorePhaseOne(dirPath)
	backRepo.BackRepoFormFieldString.RestorePhaseOne(dirPath)
	backRepo.BackRepoFormFieldTime.RestorePhaseOne(dirPath)
	backRepo.BackRepoFormGroup.RestorePhaseOne(dirPath)
	backRepo.BackRepoFormSortAssocButton.RestorePhaseOne(dirPath)
	backRepo.BackRepoOption.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoCheckBox.RestorePhaseTwo()
	backRepo.BackRepoFormDiv.RestorePhaseTwo()
	backRepo.BackRepoFormEditAssocButton.RestorePhaseTwo()
	backRepo.BackRepoFormField.RestorePhaseTwo()
	backRepo.BackRepoFormFieldDate.RestorePhaseTwo()
	backRepo.BackRepoFormFieldDateTime.RestorePhaseTwo()
	backRepo.BackRepoFormFieldFloat64.RestorePhaseTwo()
	backRepo.BackRepoFormFieldInt.RestorePhaseTwo()
	backRepo.BackRepoFormFieldSelect.RestorePhaseTwo()
	backRepo.BackRepoFormFieldString.RestorePhaseTwo()
	backRepo.BackRepoFormFieldTime.RestorePhaseTwo()
	backRepo.BackRepoFormGroup.RestorePhaseTwo()
	backRepo.BackRepoFormSortAssocButton.RestorePhaseTwo()
	backRepo.BackRepoOption.RestorePhaseTwo()

	backRepo.stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	backRepo.stage.Reset()

	// commit the cleaned stage
	backRepo.stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)
	_ = file

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoCheckBox.RestoreXLPhaseOne(file)
	backRepo.BackRepoFormDiv.RestoreXLPhaseOne(file)
	backRepo.BackRepoFormEditAssocButton.RestoreXLPhaseOne(file)
	backRepo.BackRepoFormField.RestoreXLPhaseOne(file)
	backRepo.BackRepoFormFieldDate.RestoreXLPhaseOne(file)
	backRepo.BackRepoFormFieldDateTime.RestoreXLPhaseOne(file)
	backRepo.BackRepoFormFieldFloat64.RestoreXLPhaseOne(file)
	backRepo.BackRepoFormFieldInt.RestoreXLPhaseOne(file)
	backRepo.BackRepoFormFieldSelect.RestoreXLPhaseOne(file)
	backRepo.BackRepoFormFieldString.RestoreXLPhaseOne(file)
	backRepo.BackRepoFormFieldTime.RestoreXLPhaseOne(file)
	backRepo.BackRepoFormGroup.RestoreXLPhaseOne(file)
	backRepo.BackRepoFormSortAssocButton.RestoreXLPhaseOne(file)
	backRepo.BackRepoOption.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}
