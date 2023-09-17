// generated code - do not edit
package models

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

	// insertion point for definition of arrays registering instances
	CheckBoxs           map[*CheckBox]any
	CheckBoxs_mapString map[string]*CheckBox

	OnAfterCheckBoxCreateCallback OnAfterCreateInterface[CheckBox]
	OnAfterCheckBoxUpdateCallback OnAfterUpdateInterface[CheckBox]
	OnAfterCheckBoxDeleteCallback OnAfterDeleteInterface[CheckBox]
	OnAfterCheckBoxReadCallback   OnAfterReadInterface[CheckBox]

	FormDivs           map[*FormDiv]any
	FormDivs_mapString map[string]*FormDiv

	OnAfterFormDivCreateCallback OnAfterCreateInterface[FormDiv]
	OnAfterFormDivUpdateCallback OnAfterUpdateInterface[FormDiv]
	OnAfterFormDivDeleteCallback OnAfterDeleteInterface[FormDiv]
	OnAfterFormDivReadCallback   OnAfterReadInterface[FormDiv]

	FormEditAssocButtons           map[*FormEditAssocButton]any
	FormEditAssocButtons_mapString map[string]*FormEditAssocButton

	OnAfterFormEditAssocButtonCreateCallback OnAfterCreateInterface[FormEditAssocButton]
	OnAfterFormEditAssocButtonUpdateCallback OnAfterUpdateInterface[FormEditAssocButton]
	OnAfterFormEditAssocButtonDeleteCallback OnAfterDeleteInterface[FormEditAssocButton]
	OnAfterFormEditAssocButtonReadCallback   OnAfterReadInterface[FormEditAssocButton]

	FormFields           map[*FormField]any
	FormFields_mapString map[string]*FormField

	OnAfterFormFieldCreateCallback OnAfterCreateInterface[FormField]
	OnAfterFormFieldUpdateCallback OnAfterUpdateInterface[FormField]
	OnAfterFormFieldDeleteCallback OnAfterDeleteInterface[FormField]
	OnAfterFormFieldReadCallback   OnAfterReadInterface[FormField]

	FormFieldDates           map[*FormFieldDate]any
	FormFieldDates_mapString map[string]*FormFieldDate

	OnAfterFormFieldDateCreateCallback OnAfterCreateInterface[FormFieldDate]
	OnAfterFormFieldDateUpdateCallback OnAfterUpdateInterface[FormFieldDate]
	OnAfterFormFieldDateDeleteCallback OnAfterDeleteInterface[FormFieldDate]
	OnAfterFormFieldDateReadCallback   OnAfterReadInterface[FormFieldDate]

	FormFieldDateTimes           map[*FormFieldDateTime]any
	FormFieldDateTimes_mapString map[string]*FormFieldDateTime

	OnAfterFormFieldDateTimeCreateCallback OnAfterCreateInterface[FormFieldDateTime]
	OnAfterFormFieldDateTimeUpdateCallback OnAfterUpdateInterface[FormFieldDateTime]
	OnAfterFormFieldDateTimeDeleteCallback OnAfterDeleteInterface[FormFieldDateTime]
	OnAfterFormFieldDateTimeReadCallback   OnAfterReadInterface[FormFieldDateTime]

	FormFieldFloat64s           map[*FormFieldFloat64]any
	FormFieldFloat64s_mapString map[string]*FormFieldFloat64

	OnAfterFormFieldFloat64CreateCallback OnAfterCreateInterface[FormFieldFloat64]
	OnAfterFormFieldFloat64UpdateCallback OnAfterUpdateInterface[FormFieldFloat64]
	OnAfterFormFieldFloat64DeleteCallback OnAfterDeleteInterface[FormFieldFloat64]
	OnAfterFormFieldFloat64ReadCallback   OnAfterReadInterface[FormFieldFloat64]

	FormFieldInts           map[*FormFieldInt]any
	FormFieldInts_mapString map[string]*FormFieldInt

	OnAfterFormFieldIntCreateCallback OnAfterCreateInterface[FormFieldInt]
	OnAfterFormFieldIntUpdateCallback OnAfterUpdateInterface[FormFieldInt]
	OnAfterFormFieldIntDeleteCallback OnAfterDeleteInterface[FormFieldInt]
	OnAfterFormFieldIntReadCallback   OnAfterReadInterface[FormFieldInt]

	FormFieldSelects           map[*FormFieldSelect]any
	FormFieldSelects_mapString map[string]*FormFieldSelect

	OnAfterFormFieldSelectCreateCallback OnAfterCreateInterface[FormFieldSelect]
	OnAfterFormFieldSelectUpdateCallback OnAfterUpdateInterface[FormFieldSelect]
	OnAfterFormFieldSelectDeleteCallback OnAfterDeleteInterface[FormFieldSelect]
	OnAfterFormFieldSelectReadCallback   OnAfterReadInterface[FormFieldSelect]

	FormFieldStrings           map[*FormFieldString]any
	FormFieldStrings_mapString map[string]*FormFieldString

	OnAfterFormFieldStringCreateCallback OnAfterCreateInterface[FormFieldString]
	OnAfterFormFieldStringUpdateCallback OnAfterUpdateInterface[FormFieldString]
	OnAfterFormFieldStringDeleteCallback OnAfterDeleteInterface[FormFieldString]
	OnAfterFormFieldStringReadCallback   OnAfterReadInterface[FormFieldString]

	FormFieldTimes           map[*FormFieldTime]any
	FormFieldTimes_mapString map[string]*FormFieldTime

	OnAfterFormFieldTimeCreateCallback OnAfterCreateInterface[FormFieldTime]
	OnAfterFormFieldTimeUpdateCallback OnAfterUpdateInterface[FormFieldTime]
	OnAfterFormFieldTimeDeleteCallback OnAfterDeleteInterface[FormFieldTime]
	OnAfterFormFieldTimeReadCallback   OnAfterReadInterface[FormFieldTime]

	FormGroups           map[*FormGroup]any
	FormGroups_mapString map[string]*FormGroup

	OnAfterFormGroupCreateCallback OnAfterCreateInterface[FormGroup]
	OnAfterFormGroupUpdateCallback OnAfterUpdateInterface[FormGroup]
	OnAfterFormGroupDeleteCallback OnAfterDeleteInterface[FormGroup]
	OnAfterFormGroupReadCallback   OnAfterReadInterface[FormGroup]

	FormSortAssocButtons           map[*FormSortAssocButton]any
	FormSortAssocButtons_mapString map[string]*FormSortAssocButton

	OnAfterFormSortAssocButtonCreateCallback OnAfterCreateInterface[FormSortAssocButton]
	OnAfterFormSortAssocButtonUpdateCallback OnAfterUpdateInterface[FormSortAssocButton]
	OnAfterFormSortAssocButtonDeleteCallback OnAfterDeleteInterface[FormSortAssocButton]
	OnAfterFormSortAssocButtonReadCallback   OnAfterReadInterface[FormSortAssocButton]

	Options           map[*Option]any
	Options_mapString map[string]*Option

	OnAfterOptionCreateCallback OnAfterCreateInterface[Option]
	OnAfterOptionUpdateCallback OnAfterUpdateInterface[Option]
	OnAfterOptionDeleteCallback OnAfterDeleteInterface[Option]
	OnAfterOptionReadCallback   OnAfterReadInterface[Option]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitCheckBox(checkbox *CheckBox)
	CheckoutCheckBox(checkbox *CheckBox)
	CommitFormDiv(formdiv *FormDiv)
	CheckoutFormDiv(formdiv *FormDiv)
	CommitFormEditAssocButton(formeditassocbutton *FormEditAssocButton)
	CheckoutFormEditAssocButton(formeditassocbutton *FormEditAssocButton)
	CommitFormField(formfield *FormField)
	CheckoutFormField(formfield *FormField)
	CommitFormFieldDate(formfielddate *FormFieldDate)
	CheckoutFormFieldDate(formfielddate *FormFieldDate)
	CommitFormFieldDateTime(formfielddatetime *FormFieldDateTime)
	CheckoutFormFieldDateTime(formfielddatetime *FormFieldDateTime)
	CommitFormFieldFloat64(formfieldfloat64 *FormFieldFloat64)
	CheckoutFormFieldFloat64(formfieldfloat64 *FormFieldFloat64)
	CommitFormFieldInt(formfieldint *FormFieldInt)
	CheckoutFormFieldInt(formfieldint *FormFieldInt)
	CommitFormFieldSelect(formfieldselect *FormFieldSelect)
	CheckoutFormFieldSelect(formfieldselect *FormFieldSelect)
	CommitFormFieldString(formfieldstring *FormFieldString)
	CheckoutFormFieldString(formfieldstring *FormFieldString)
	CommitFormFieldTime(formfieldtime *FormFieldTime)
	CheckoutFormFieldTime(formfieldtime *FormFieldTime)
	CommitFormGroup(formgroup *FormGroup)
	CheckoutFormGroup(formgroup *FormGroup)
	CommitFormSortAssocButton(formsortassocbutton *FormSortAssocButton)
	CheckoutFormSortAssocButton(formsortassocbutton *FormSortAssocButton)
	CommitOption(option *Option)
	CheckoutOption(option *Option)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

var _stage *StageStruct

var once sync.Once

func GetDefaultStage() *StageStruct {
	once.Do(func() {
		_stage = NewStage("")
	})
	return _stage
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		CheckBoxs:           make(map[*CheckBox]any),
		CheckBoxs_mapString: make(map[string]*CheckBox),

		FormDivs:           make(map[*FormDiv]any),
		FormDivs_mapString: make(map[string]*FormDiv),

		FormEditAssocButtons:           make(map[*FormEditAssocButton]any),
		FormEditAssocButtons_mapString: make(map[string]*FormEditAssocButton),

		FormFields:           make(map[*FormField]any),
		FormFields_mapString: make(map[string]*FormField),

		FormFieldDates:           make(map[*FormFieldDate]any),
		FormFieldDates_mapString: make(map[string]*FormFieldDate),

		FormFieldDateTimes:           make(map[*FormFieldDateTime]any),
		FormFieldDateTimes_mapString: make(map[string]*FormFieldDateTime),

		FormFieldFloat64s:           make(map[*FormFieldFloat64]any),
		FormFieldFloat64s_mapString: make(map[string]*FormFieldFloat64),

		FormFieldInts:           make(map[*FormFieldInt]any),
		FormFieldInts_mapString: make(map[string]*FormFieldInt),

		FormFieldSelects:           make(map[*FormFieldSelect]any),
		FormFieldSelects_mapString: make(map[string]*FormFieldSelect),

		FormFieldStrings:           make(map[*FormFieldString]any),
		FormFieldStrings_mapString: make(map[string]*FormFieldString),

		FormFieldTimes:           make(map[*FormFieldTime]any),
		FormFieldTimes_mapString: make(map[string]*FormFieldTime),

		FormGroups:           make(map[*FormGroup]any),
		FormGroups_mapString: make(map[string]*FormGroup),

		FormSortAssocButtons:           make(map[*FormSortAssocButton]any),
		FormSortAssocButtons_mapString: make(map[string]*FormSortAssocButton),

		Options:           make(map[*Option]any),
		Options_mapString: make(map[string]*Option),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["CheckBox"] = len(stage.CheckBoxs)
	stage.Map_GongStructName_InstancesNb["FormDiv"] = len(stage.FormDivs)
	stage.Map_GongStructName_InstancesNb["FormEditAssocButton"] = len(stage.FormEditAssocButtons)
	stage.Map_GongStructName_InstancesNb["FormField"] = len(stage.FormFields)
	stage.Map_GongStructName_InstancesNb["FormFieldDate"] = len(stage.FormFieldDates)
	stage.Map_GongStructName_InstancesNb["FormFieldDateTime"] = len(stage.FormFieldDateTimes)
	stage.Map_GongStructName_InstancesNb["FormFieldFloat64"] = len(stage.FormFieldFloat64s)
	stage.Map_GongStructName_InstancesNb["FormFieldInt"] = len(stage.FormFieldInts)
	stage.Map_GongStructName_InstancesNb["FormFieldSelect"] = len(stage.FormFieldSelects)
	stage.Map_GongStructName_InstancesNb["FormFieldString"] = len(stage.FormFieldStrings)
	stage.Map_GongStructName_InstancesNb["FormFieldTime"] = len(stage.FormFieldTimes)
	stage.Map_GongStructName_InstancesNb["FormGroup"] = len(stage.FormGroups)
	stage.Map_GongStructName_InstancesNb["FormSortAssocButton"] = len(stage.FormSortAssocButtons)
	stage.Map_GongStructName_InstancesNb["Option"] = len(stage.Options)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["CheckBox"] = len(stage.CheckBoxs)
	stage.Map_GongStructName_InstancesNb["FormDiv"] = len(stage.FormDivs)
	stage.Map_GongStructName_InstancesNb["FormEditAssocButton"] = len(stage.FormEditAssocButtons)
	stage.Map_GongStructName_InstancesNb["FormField"] = len(stage.FormFields)
	stage.Map_GongStructName_InstancesNb["FormFieldDate"] = len(stage.FormFieldDates)
	stage.Map_GongStructName_InstancesNb["FormFieldDateTime"] = len(stage.FormFieldDateTimes)
	stage.Map_GongStructName_InstancesNb["FormFieldFloat64"] = len(stage.FormFieldFloat64s)
	stage.Map_GongStructName_InstancesNb["FormFieldInt"] = len(stage.FormFieldInts)
	stage.Map_GongStructName_InstancesNb["FormFieldSelect"] = len(stage.FormFieldSelects)
	stage.Map_GongStructName_InstancesNb["FormFieldString"] = len(stage.FormFieldStrings)
	stage.Map_GongStructName_InstancesNb["FormFieldTime"] = len(stage.FormFieldTimes)
	stage.Map_GongStructName_InstancesNb["FormGroup"] = len(stage.FormGroups)
	stage.Map_GongStructName_InstancesNb["FormSortAssocButton"] = len(stage.FormSortAssocButtons)
	stage.Map_GongStructName_InstancesNb["Option"] = len(stage.Options)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts checkbox to the model stage
func (checkbox *CheckBox) Stage(stage *StageStruct) *CheckBox {
	stage.CheckBoxs[checkbox] = __member
	stage.CheckBoxs_mapString[checkbox.Name] = checkbox

	return checkbox
}

// Unstage removes checkbox off the model stage
func (checkbox *CheckBox) Unstage(stage *StageStruct) *CheckBox {
	delete(stage.CheckBoxs, checkbox)
	delete(stage.CheckBoxs_mapString, checkbox.Name)
	return checkbox
}

// UnstageVoid removes checkbox off the model stage
func (checkbox *CheckBox) UnstageVoid(stage *StageStruct) {
	delete(stage.CheckBoxs, checkbox)
	delete(stage.CheckBoxs_mapString, checkbox.Name)
}

// commit checkbox to the back repo (if it is already staged)
func (checkbox *CheckBox) Commit(stage *StageStruct) *CheckBox {
	if _, ok := stage.CheckBoxs[checkbox]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCheckBox(checkbox)
		}
	}
	return checkbox
}

func (checkbox *CheckBox) CommitVoid(stage *StageStruct) {
	checkbox.Commit(stage)
}

// Checkout checkbox to the back repo (if it is already staged)
func (checkbox *CheckBox) Checkout(stage *StageStruct) *CheckBox {
	if _, ok := stage.CheckBoxs[checkbox]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCheckBox(checkbox)
		}
	}
	return checkbox
}

// for satisfaction of GongStruct interface
func (checkbox *CheckBox) GetName() (res string) {
	return checkbox.Name
}

// Stage puts formdiv to the model stage
func (formdiv *FormDiv) Stage(stage *StageStruct) *FormDiv {
	stage.FormDivs[formdiv] = __member
	stage.FormDivs_mapString[formdiv.Name] = formdiv

	return formdiv
}

// Unstage removes formdiv off the model stage
func (formdiv *FormDiv) Unstage(stage *StageStruct) *FormDiv {
	delete(stage.FormDivs, formdiv)
	delete(stage.FormDivs_mapString, formdiv.Name)
	return formdiv
}

// UnstageVoid removes formdiv off the model stage
func (formdiv *FormDiv) UnstageVoid(stage *StageStruct) {
	delete(stage.FormDivs, formdiv)
	delete(stage.FormDivs_mapString, formdiv.Name)
}

// commit formdiv to the back repo (if it is already staged)
func (formdiv *FormDiv) Commit(stage *StageStruct) *FormDiv {
	if _, ok := stage.FormDivs[formdiv]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormDiv(formdiv)
		}
	}
	return formdiv
}

func (formdiv *FormDiv) CommitVoid(stage *StageStruct) {
	formdiv.Commit(stage)
}

// Checkout formdiv to the back repo (if it is already staged)
func (formdiv *FormDiv) Checkout(stage *StageStruct) *FormDiv {
	if _, ok := stage.FormDivs[formdiv]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormDiv(formdiv)
		}
	}
	return formdiv
}

// for satisfaction of GongStruct interface
func (formdiv *FormDiv) GetName() (res string) {
	return formdiv.Name
}

// Stage puts formeditassocbutton to the model stage
func (formeditassocbutton *FormEditAssocButton) Stage(stage *StageStruct) *FormEditAssocButton {
	stage.FormEditAssocButtons[formeditassocbutton] = __member
	stage.FormEditAssocButtons_mapString[formeditassocbutton.Name] = formeditassocbutton

	return formeditassocbutton
}

// Unstage removes formeditassocbutton off the model stage
func (formeditassocbutton *FormEditAssocButton) Unstage(stage *StageStruct) *FormEditAssocButton {
	delete(stage.FormEditAssocButtons, formeditassocbutton)
	delete(stage.FormEditAssocButtons_mapString, formeditassocbutton.Name)
	return formeditassocbutton
}

// UnstageVoid removes formeditassocbutton off the model stage
func (formeditassocbutton *FormEditAssocButton) UnstageVoid(stage *StageStruct) {
	delete(stage.FormEditAssocButtons, formeditassocbutton)
	delete(stage.FormEditAssocButtons_mapString, formeditassocbutton.Name)
}

// commit formeditassocbutton to the back repo (if it is already staged)
func (formeditassocbutton *FormEditAssocButton) Commit(stage *StageStruct) *FormEditAssocButton {
	if _, ok := stage.FormEditAssocButtons[formeditassocbutton]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormEditAssocButton(formeditassocbutton)
		}
	}
	return formeditassocbutton
}

func (formeditassocbutton *FormEditAssocButton) CommitVoid(stage *StageStruct) {
	formeditassocbutton.Commit(stage)
}

// Checkout formeditassocbutton to the back repo (if it is already staged)
func (formeditassocbutton *FormEditAssocButton) Checkout(stage *StageStruct) *FormEditAssocButton {
	if _, ok := stage.FormEditAssocButtons[formeditassocbutton]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormEditAssocButton(formeditassocbutton)
		}
	}
	return formeditassocbutton
}

// for satisfaction of GongStruct interface
func (formeditassocbutton *FormEditAssocButton) GetName() (res string) {
	return formeditassocbutton.Name
}

// Stage puts formfield to the model stage
func (formfield *FormField) Stage(stage *StageStruct) *FormField {
	stage.FormFields[formfield] = __member
	stage.FormFields_mapString[formfield.Name] = formfield

	return formfield
}

// Unstage removes formfield off the model stage
func (formfield *FormField) Unstage(stage *StageStruct) *FormField {
	delete(stage.FormFields, formfield)
	delete(stage.FormFields_mapString, formfield.Name)
	return formfield
}

// UnstageVoid removes formfield off the model stage
func (formfield *FormField) UnstageVoid(stage *StageStruct) {
	delete(stage.FormFields, formfield)
	delete(stage.FormFields_mapString, formfield.Name)
}

// commit formfield to the back repo (if it is already staged)
func (formfield *FormField) Commit(stage *StageStruct) *FormField {
	if _, ok := stage.FormFields[formfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormField(formfield)
		}
	}
	return formfield
}

func (formfield *FormField) CommitVoid(stage *StageStruct) {
	formfield.Commit(stage)
}

// Checkout formfield to the back repo (if it is already staged)
func (formfield *FormField) Checkout(stage *StageStruct) *FormField {
	if _, ok := stage.FormFields[formfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormField(formfield)
		}
	}
	return formfield
}

// for satisfaction of GongStruct interface
func (formfield *FormField) GetName() (res string) {
	return formfield.Name
}

// Stage puts formfielddate to the model stage
func (formfielddate *FormFieldDate) Stage(stage *StageStruct) *FormFieldDate {
	stage.FormFieldDates[formfielddate] = __member
	stage.FormFieldDates_mapString[formfielddate.Name] = formfielddate

	return formfielddate
}

// Unstage removes formfielddate off the model stage
func (formfielddate *FormFieldDate) Unstage(stage *StageStruct) *FormFieldDate {
	delete(stage.FormFieldDates, formfielddate)
	delete(stage.FormFieldDates_mapString, formfielddate.Name)
	return formfielddate
}

// UnstageVoid removes formfielddate off the model stage
func (formfielddate *FormFieldDate) UnstageVoid(stage *StageStruct) {
	delete(stage.FormFieldDates, formfielddate)
	delete(stage.FormFieldDates_mapString, formfielddate.Name)
}

// commit formfielddate to the back repo (if it is already staged)
func (formfielddate *FormFieldDate) Commit(stage *StageStruct) *FormFieldDate {
	if _, ok := stage.FormFieldDates[formfielddate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldDate(formfielddate)
		}
	}
	return formfielddate
}

func (formfielddate *FormFieldDate) CommitVoid(stage *StageStruct) {
	formfielddate.Commit(stage)
}

// Checkout formfielddate to the back repo (if it is already staged)
func (formfielddate *FormFieldDate) Checkout(stage *StageStruct) *FormFieldDate {
	if _, ok := stage.FormFieldDates[formfielddate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldDate(formfielddate)
		}
	}
	return formfielddate
}

// for satisfaction of GongStruct interface
func (formfielddate *FormFieldDate) GetName() (res string) {
	return formfielddate.Name
}

// Stage puts formfielddatetime to the model stage
func (formfielddatetime *FormFieldDateTime) Stage(stage *StageStruct) *FormFieldDateTime {
	stage.FormFieldDateTimes[formfielddatetime] = __member
	stage.FormFieldDateTimes_mapString[formfielddatetime.Name] = formfielddatetime

	return formfielddatetime
}

// Unstage removes formfielddatetime off the model stage
func (formfielddatetime *FormFieldDateTime) Unstage(stage *StageStruct) *FormFieldDateTime {
	delete(stage.FormFieldDateTimes, formfielddatetime)
	delete(stage.FormFieldDateTimes_mapString, formfielddatetime.Name)
	return formfielddatetime
}

// UnstageVoid removes formfielddatetime off the model stage
func (formfielddatetime *FormFieldDateTime) UnstageVoid(stage *StageStruct) {
	delete(stage.FormFieldDateTimes, formfielddatetime)
	delete(stage.FormFieldDateTimes_mapString, formfielddatetime.Name)
}

// commit formfielddatetime to the back repo (if it is already staged)
func (formfielddatetime *FormFieldDateTime) Commit(stage *StageStruct) *FormFieldDateTime {
	if _, ok := stage.FormFieldDateTimes[formfielddatetime]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldDateTime(formfielddatetime)
		}
	}
	return formfielddatetime
}

func (formfielddatetime *FormFieldDateTime) CommitVoid(stage *StageStruct) {
	formfielddatetime.Commit(stage)
}

// Checkout formfielddatetime to the back repo (if it is already staged)
func (formfielddatetime *FormFieldDateTime) Checkout(stage *StageStruct) *FormFieldDateTime {
	if _, ok := stage.FormFieldDateTimes[formfielddatetime]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldDateTime(formfielddatetime)
		}
	}
	return formfielddatetime
}

// for satisfaction of GongStruct interface
func (formfielddatetime *FormFieldDateTime) GetName() (res string) {
	return formfielddatetime.Name
}

// Stage puts formfieldfloat64 to the model stage
func (formfieldfloat64 *FormFieldFloat64) Stage(stage *StageStruct) *FormFieldFloat64 {
	stage.FormFieldFloat64s[formfieldfloat64] = __member
	stage.FormFieldFloat64s_mapString[formfieldfloat64.Name] = formfieldfloat64

	return formfieldfloat64
}

// Unstage removes formfieldfloat64 off the model stage
func (formfieldfloat64 *FormFieldFloat64) Unstage(stage *StageStruct) *FormFieldFloat64 {
	delete(stage.FormFieldFloat64s, formfieldfloat64)
	delete(stage.FormFieldFloat64s_mapString, formfieldfloat64.Name)
	return formfieldfloat64
}

// UnstageVoid removes formfieldfloat64 off the model stage
func (formfieldfloat64 *FormFieldFloat64) UnstageVoid(stage *StageStruct) {
	delete(stage.FormFieldFloat64s, formfieldfloat64)
	delete(stage.FormFieldFloat64s_mapString, formfieldfloat64.Name)
}

// commit formfieldfloat64 to the back repo (if it is already staged)
func (formfieldfloat64 *FormFieldFloat64) Commit(stage *StageStruct) *FormFieldFloat64 {
	if _, ok := stage.FormFieldFloat64s[formfieldfloat64]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldFloat64(formfieldfloat64)
		}
	}
	return formfieldfloat64
}

func (formfieldfloat64 *FormFieldFloat64) CommitVoid(stage *StageStruct) {
	formfieldfloat64.Commit(stage)
}

// Checkout formfieldfloat64 to the back repo (if it is already staged)
func (formfieldfloat64 *FormFieldFloat64) Checkout(stage *StageStruct) *FormFieldFloat64 {
	if _, ok := stage.FormFieldFloat64s[formfieldfloat64]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldFloat64(formfieldfloat64)
		}
	}
	return formfieldfloat64
}

// for satisfaction of GongStruct interface
func (formfieldfloat64 *FormFieldFloat64) GetName() (res string) {
	return formfieldfloat64.Name
}

// Stage puts formfieldint to the model stage
func (formfieldint *FormFieldInt) Stage(stage *StageStruct) *FormFieldInt {
	stage.FormFieldInts[formfieldint] = __member
	stage.FormFieldInts_mapString[formfieldint.Name] = formfieldint

	return formfieldint
}

// Unstage removes formfieldint off the model stage
func (formfieldint *FormFieldInt) Unstage(stage *StageStruct) *FormFieldInt {
	delete(stage.FormFieldInts, formfieldint)
	delete(stage.FormFieldInts_mapString, formfieldint.Name)
	return formfieldint
}

// UnstageVoid removes formfieldint off the model stage
func (formfieldint *FormFieldInt) UnstageVoid(stage *StageStruct) {
	delete(stage.FormFieldInts, formfieldint)
	delete(stage.FormFieldInts_mapString, formfieldint.Name)
}

// commit formfieldint to the back repo (if it is already staged)
func (formfieldint *FormFieldInt) Commit(stage *StageStruct) *FormFieldInt {
	if _, ok := stage.FormFieldInts[formfieldint]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldInt(formfieldint)
		}
	}
	return formfieldint
}

func (formfieldint *FormFieldInt) CommitVoid(stage *StageStruct) {
	formfieldint.Commit(stage)
}

// Checkout formfieldint to the back repo (if it is already staged)
func (formfieldint *FormFieldInt) Checkout(stage *StageStruct) *FormFieldInt {
	if _, ok := stage.FormFieldInts[formfieldint]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldInt(formfieldint)
		}
	}
	return formfieldint
}

// for satisfaction of GongStruct interface
func (formfieldint *FormFieldInt) GetName() (res string) {
	return formfieldint.Name
}

// Stage puts formfieldselect to the model stage
func (formfieldselect *FormFieldSelect) Stage(stage *StageStruct) *FormFieldSelect {
	stage.FormFieldSelects[formfieldselect] = __member
	stage.FormFieldSelects_mapString[formfieldselect.Name] = formfieldselect

	return formfieldselect
}

// Unstage removes formfieldselect off the model stage
func (formfieldselect *FormFieldSelect) Unstage(stage *StageStruct) *FormFieldSelect {
	delete(stage.FormFieldSelects, formfieldselect)
	delete(stage.FormFieldSelects_mapString, formfieldselect.Name)
	return formfieldselect
}

// UnstageVoid removes formfieldselect off the model stage
func (formfieldselect *FormFieldSelect) UnstageVoid(stage *StageStruct) {
	delete(stage.FormFieldSelects, formfieldselect)
	delete(stage.FormFieldSelects_mapString, formfieldselect.Name)
}

// commit formfieldselect to the back repo (if it is already staged)
func (formfieldselect *FormFieldSelect) Commit(stage *StageStruct) *FormFieldSelect {
	if _, ok := stage.FormFieldSelects[formfieldselect]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldSelect(formfieldselect)
		}
	}
	return formfieldselect
}

func (formfieldselect *FormFieldSelect) CommitVoid(stage *StageStruct) {
	formfieldselect.Commit(stage)
}

// Checkout formfieldselect to the back repo (if it is already staged)
func (formfieldselect *FormFieldSelect) Checkout(stage *StageStruct) *FormFieldSelect {
	if _, ok := stage.FormFieldSelects[formfieldselect]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldSelect(formfieldselect)
		}
	}
	return formfieldselect
}

// for satisfaction of GongStruct interface
func (formfieldselect *FormFieldSelect) GetName() (res string) {
	return formfieldselect.Name
}

// Stage puts formfieldstring to the model stage
func (formfieldstring *FormFieldString) Stage(stage *StageStruct) *FormFieldString {
	stage.FormFieldStrings[formfieldstring] = __member
	stage.FormFieldStrings_mapString[formfieldstring.Name] = formfieldstring

	return formfieldstring
}

// Unstage removes formfieldstring off the model stage
func (formfieldstring *FormFieldString) Unstage(stage *StageStruct) *FormFieldString {
	delete(stage.FormFieldStrings, formfieldstring)
	delete(stage.FormFieldStrings_mapString, formfieldstring.Name)
	return formfieldstring
}

// UnstageVoid removes formfieldstring off the model stage
func (formfieldstring *FormFieldString) UnstageVoid(stage *StageStruct) {
	delete(stage.FormFieldStrings, formfieldstring)
	delete(stage.FormFieldStrings_mapString, formfieldstring.Name)
}

// commit formfieldstring to the back repo (if it is already staged)
func (formfieldstring *FormFieldString) Commit(stage *StageStruct) *FormFieldString {
	if _, ok := stage.FormFieldStrings[formfieldstring]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldString(formfieldstring)
		}
	}
	return formfieldstring
}

func (formfieldstring *FormFieldString) CommitVoid(stage *StageStruct) {
	formfieldstring.Commit(stage)
}

// Checkout formfieldstring to the back repo (if it is already staged)
func (formfieldstring *FormFieldString) Checkout(stage *StageStruct) *FormFieldString {
	if _, ok := stage.FormFieldStrings[formfieldstring]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldString(formfieldstring)
		}
	}
	return formfieldstring
}

// for satisfaction of GongStruct interface
func (formfieldstring *FormFieldString) GetName() (res string) {
	return formfieldstring.Name
}

// Stage puts formfieldtime to the model stage
func (formfieldtime *FormFieldTime) Stage(stage *StageStruct) *FormFieldTime {
	stage.FormFieldTimes[formfieldtime] = __member
	stage.FormFieldTimes_mapString[formfieldtime.Name] = formfieldtime

	return formfieldtime
}

// Unstage removes formfieldtime off the model stage
func (formfieldtime *FormFieldTime) Unstage(stage *StageStruct) *FormFieldTime {
	delete(stage.FormFieldTimes, formfieldtime)
	delete(stage.FormFieldTimes_mapString, formfieldtime.Name)
	return formfieldtime
}

// UnstageVoid removes formfieldtime off the model stage
func (formfieldtime *FormFieldTime) UnstageVoid(stage *StageStruct) {
	delete(stage.FormFieldTimes, formfieldtime)
	delete(stage.FormFieldTimes_mapString, formfieldtime.Name)
}

// commit formfieldtime to the back repo (if it is already staged)
func (formfieldtime *FormFieldTime) Commit(stage *StageStruct) *FormFieldTime {
	if _, ok := stage.FormFieldTimes[formfieldtime]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldTime(formfieldtime)
		}
	}
	return formfieldtime
}

func (formfieldtime *FormFieldTime) CommitVoid(stage *StageStruct) {
	formfieldtime.Commit(stage)
}

// Checkout formfieldtime to the back repo (if it is already staged)
func (formfieldtime *FormFieldTime) Checkout(stage *StageStruct) *FormFieldTime {
	if _, ok := stage.FormFieldTimes[formfieldtime]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldTime(formfieldtime)
		}
	}
	return formfieldtime
}

// for satisfaction of GongStruct interface
func (formfieldtime *FormFieldTime) GetName() (res string) {
	return formfieldtime.Name
}

// Stage puts formgroup to the model stage
func (formgroup *FormGroup) Stage(stage *StageStruct) *FormGroup {
	stage.FormGroups[formgroup] = __member
	stage.FormGroups_mapString[formgroup.Name] = formgroup

	return formgroup
}

// Unstage removes formgroup off the model stage
func (formgroup *FormGroup) Unstage(stage *StageStruct) *FormGroup {
	delete(stage.FormGroups, formgroup)
	delete(stage.FormGroups_mapString, formgroup.Name)
	return formgroup
}

// UnstageVoid removes formgroup off the model stage
func (formgroup *FormGroup) UnstageVoid(stage *StageStruct) {
	delete(stage.FormGroups, formgroup)
	delete(stage.FormGroups_mapString, formgroup.Name)
}

// commit formgroup to the back repo (if it is already staged)
func (formgroup *FormGroup) Commit(stage *StageStruct) *FormGroup {
	if _, ok := stage.FormGroups[formgroup]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormGroup(formgroup)
		}
	}
	return formgroup
}

func (formgroup *FormGroup) CommitVoid(stage *StageStruct) {
	formgroup.Commit(stage)
}

// Checkout formgroup to the back repo (if it is already staged)
func (formgroup *FormGroup) Checkout(stage *StageStruct) *FormGroup {
	if _, ok := stage.FormGroups[formgroup]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormGroup(formgroup)
		}
	}
	return formgroup
}

// for satisfaction of GongStruct interface
func (formgroup *FormGroup) GetName() (res string) {
	return formgroup.Name
}

// Stage puts formsortassocbutton to the model stage
func (formsortassocbutton *FormSortAssocButton) Stage(stage *StageStruct) *FormSortAssocButton {
	stage.FormSortAssocButtons[formsortassocbutton] = __member
	stage.FormSortAssocButtons_mapString[formsortassocbutton.Name] = formsortassocbutton

	return formsortassocbutton
}

// Unstage removes formsortassocbutton off the model stage
func (formsortassocbutton *FormSortAssocButton) Unstage(stage *StageStruct) *FormSortAssocButton {
	delete(stage.FormSortAssocButtons, formsortassocbutton)
	delete(stage.FormSortAssocButtons_mapString, formsortassocbutton.Name)
	return formsortassocbutton
}

// UnstageVoid removes formsortassocbutton off the model stage
func (formsortassocbutton *FormSortAssocButton) UnstageVoid(stage *StageStruct) {
	delete(stage.FormSortAssocButtons, formsortassocbutton)
	delete(stage.FormSortAssocButtons_mapString, formsortassocbutton.Name)
}

// commit formsortassocbutton to the back repo (if it is already staged)
func (formsortassocbutton *FormSortAssocButton) Commit(stage *StageStruct) *FormSortAssocButton {
	if _, ok := stage.FormSortAssocButtons[formsortassocbutton]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormSortAssocButton(formsortassocbutton)
		}
	}
	return formsortassocbutton
}

func (formsortassocbutton *FormSortAssocButton) CommitVoid(stage *StageStruct) {
	formsortassocbutton.Commit(stage)
}

// Checkout formsortassocbutton to the back repo (if it is already staged)
func (formsortassocbutton *FormSortAssocButton) Checkout(stage *StageStruct) *FormSortAssocButton {
	if _, ok := stage.FormSortAssocButtons[formsortassocbutton]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormSortAssocButton(formsortassocbutton)
		}
	}
	return formsortassocbutton
}

// for satisfaction of GongStruct interface
func (formsortassocbutton *FormSortAssocButton) GetName() (res string) {
	return formsortassocbutton.Name
}

// Stage puts option to the model stage
func (option *Option) Stage(stage *StageStruct) *Option {
	stage.Options[option] = __member
	stage.Options_mapString[option.Name] = option

	return option
}

// Unstage removes option off the model stage
func (option *Option) Unstage(stage *StageStruct) *Option {
	delete(stage.Options, option)
	delete(stage.Options_mapString, option.Name)
	return option
}

// UnstageVoid removes option off the model stage
func (option *Option) UnstageVoid(stage *StageStruct) {
	delete(stage.Options, option)
	delete(stage.Options_mapString, option.Name)
}

// commit option to the back repo (if it is already staged)
func (option *Option) Commit(stage *StageStruct) *Option {
	if _, ok := stage.Options[option]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitOption(option)
		}
	}
	return option
}

func (option *Option) CommitVoid(stage *StageStruct) {
	option.Commit(stage)
}

// Checkout option to the back repo (if it is already staged)
func (option *Option) Checkout(stage *StageStruct) *Option {
	if _, ok := stage.Options[option]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutOption(option)
		}
	}
	return option
}

// for satisfaction of GongStruct interface
func (option *Option) GetName() (res string) {
	return option.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMCheckBox(CheckBox *CheckBox)
	CreateORMFormDiv(FormDiv *FormDiv)
	CreateORMFormEditAssocButton(FormEditAssocButton *FormEditAssocButton)
	CreateORMFormField(FormField *FormField)
	CreateORMFormFieldDate(FormFieldDate *FormFieldDate)
	CreateORMFormFieldDateTime(FormFieldDateTime *FormFieldDateTime)
	CreateORMFormFieldFloat64(FormFieldFloat64 *FormFieldFloat64)
	CreateORMFormFieldInt(FormFieldInt *FormFieldInt)
	CreateORMFormFieldSelect(FormFieldSelect *FormFieldSelect)
	CreateORMFormFieldString(FormFieldString *FormFieldString)
	CreateORMFormFieldTime(FormFieldTime *FormFieldTime)
	CreateORMFormGroup(FormGroup *FormGroup)
	CreateORMFormSortAssocButton(FormSortAssocButton *FormSortAssocButton)
	CreateORMOption(Option *Option)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMCheckBox(CheckBox *CheckBox)
	DeleteORMFormDiv(FormDiv *FormDiv)
	DeleteORMFormEditAssocButton(FormEditAssocButton *FormEditAssocButton)
	DeleteORMFormField(FormField *FormField)
	DeleteORMFormFieldDate(FormFieldDate *FormFieldDate)
	DeleteORMFormFieldDateTime(FormFieldDateTime *FormFieldDateTime)
	DeleteORMFormFieldFloat64(FormFieldFloat64 *FormFieldFloat64)
	DeleteORMFormFieldInt(FormFieldInt *FormFieldInt)
	DeleteORMFormFieldSelect(FormFieldSelect *FormFieldSelect)
	DeleteORMFormFieldString(FormFieldString *FormFieldString)
	DeleteORMFormFieldTime(FormFieldTime *FormFieldTime)
	DeleteORMFormGroup(FormGroup *FormGroup)
	DeleteORMFormSortAssocButton(FormSortAssocButton *FormSortAssocButton)
	DeleteORMOption(Option *Option)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.CheckBoxs = make(map[*CheckBox]any)
	stage.CheckBoxs_mapString = make(map[string]*CheckBox)

	stage.FormDivs = make(map[*FormDiv]any)
	stage.FormDivs_mapString = make(map[string]*FormDiv)

	stage.FormEditAssocButtons = make(map[*FormEditAssocButton]any)
	stage.FormEditAssocButtons_mapString = make(map[string]*FormEditAssocButton)

	stage.FormFields = make(map[*FormField]any)
	stage.FormFields_mapString = make(map[string]*FormField)

	stage.FormFieldDates = make(map[*FormFieldDate]any)
	stage.FormFieldDates_mapString = make(map[string]*FormFieldDate)

	stage.FormFieldDateTimes = make(map[*FormFieldDateTime]any)
	stage.FormFieldDateTimes_mapString = make(map[string]*FormFieldDateTime)

	stage.FormFieldFloat64s = make(map[*FormFieldFloat64]any)
	stage.FormFieldFloat64s_mapString = make(map[string]*FormFieldFloat64)

	stage.FormFieldInts = make(map[*FormFieldInt]any)
	stage.FormFieldInts_mapString = make(map[string]*FormFieldInt)

	stage.FormFieldSelects = make(map[*FormFieldSelect]any)
	stage.FormFieldSelects_mapString = make(map[string]*FormFieldSelect)

	stage.FormFieldStrings = make(map[*FormFieldString]any)
	stage.FormFieldStrings_mapString = make(map[string]*FormFieldString)

	stage.FormFieldTimes = make(map[*FormFieldTime]any)
	stage.FormFieldTimes_mapString = make(map[string]*FormFieldTime)

	stage.FormGroups = make(map[*FormGroup]any)
	stage.FormGroups_mapString = make(map[string]*FormGroup)

	stage.FormSortAssocButtons = make(map[*FormSortAssocButton]any)
	stage.FormSortAssocButtons_mapString = make(map[string]*FormSortAssocButton)

	stage.Options = make(map[*Option]any)
	stage.Options_mapString = make(map[string]*Option)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.CheckBoxs = nil
	stage.CheckBoxs_mapString = nil

	stage.FormDivs = nil
	stage.FormDivs_mapString = nil

	stage.FormEditAssocButtons = nil
	stage.FormEditAssocButtons_mapString = nil

	stage.FormFields = nil
	stage.FormFields_mapString = nil

	stage.FormFieldDates = nil
	stage.FormFieldDates_mapString = nil

	stage.FormFieldDateTimes = nil
	stage.FormFieldDateTimes_mapString = nil

	stage.FormFieldFloat64s = nil
	stage.FormFieldFloat64s_mapString = nil

	stage.FormFieldInts = nil
	stage.FormFieldInts_mapString = nil

	stage.FormFieldSelects = nil
	stage.FormFieldSelects_mapString = nil

	stage.FormFieldStrings = nil
	stage.FormFieldStrings_mapString = nil

	stage.FormFieldTimes = nil
	stage.FormFieldTimes_mapString = nil

	stage.FormGroups = nil
	stage.FormGroups_mapString = nil

	stage.FormSortAssocButtons = nil
	stage.FormSortAssocButtons_mapString = nil

	stage.Options = nil
	stage.Options_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for checkbox := range stage.CheckBoxs {
		checkbox.Unstage(stage)
	}

	for formdiv := range stage.FormDivs {
		formdiv.Unstage(stage)
	}

	for formeditassocbutton := range stage.FormEditAssocButtons {
		formeditassocbutton.Unstage(stage)
	}

	for formfield := range stage.FormFields {
		formfield.Unstage(stage)
	}

	for formfielddate := range stage.FormFieldDates {
		formfielddate.Unstage(stage)
	}

	for formfielddatetime := range stage.FormFieldDateTimes {
		formfielddatetime.Unstage(stage)
	}

	for formfieldfloat64 := range stage.FormFieldFloat64s {
		formfieldfloat64.Unstage(stage)
	}

	for formfieldint := range stage.FormFieldInts {
		formfieldint.Unstage(stage)
	}

	for formfieldselect := range stage.FormFieldSelects {
		formfieldselect.Unstage(stage)
	}

	for formfieldstring := range stage.FormFieldStrings {
		formfieldstring.Unstage(stage)
	}

	for formfieldtime := range stage.FormFieldTimes {
		formfieldtime.Unstage(stage)
	}

	for formgroup := range stage.FormGroups {
		formgroup.Unstage(stage)
	}

	for formsortassocbutton := range stage.FormSortAssocButtons {
		formsortassocbutton.Unstage(stage)
	}

	for option := range stage.Options {
		option.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
	// insertion point for generic types
	CheckBox | FormDiv | FormEditAssocButton | FormField | FormFieldDate | FormFieldDateTime | FormFieldFloat64 | FormFieldInt | FormFieldSelect | FormFieldString | FormFieldTime | FormGroup | FormSortAssocButton | Option
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	// insertion point for generic types
	*CheckBox | *FormDiv | *FormEditAssocButton | *FormField | *FormFieldDate | *FormFieldDateTime | *FormFieldFloat64 | *FormFieldInt | *FormFieldSelect | *FormFieldString | *FormFieldTime | *FormGroup | *FormSortAssocButton | *Option
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
}

type GongstructSet interface {
	map[any]any |
		// insertion point for generic types
		map[*CheckBox]any |
		map[*FormDiv]any |
		map[*FormEditAssocButton]any |
		map[*FormField]any |
		map[*FormFieldDate]any |
		map[*FormFieldDateTime]any |
		map[*FormFieldFloat64]any |
		map[*FormFieldInt]any |
		map[*FormFieldSelect]any |
		map[*FormFieldString]any |
		map[*FormFieldTime]any |
		map[*FormGroup]any |
		map[*FormSortAssocButton]any |
		map[*Option]any |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

type GongstructMapString interface {
	map[any]any |
		// insertion point for generic types
		map[string]*CheckBox |
		map[string]*FormDiv |
		map[string]*FormEditAssocButton |
		map[string]*FormField |
		map[string]*FormFieldDate |
		map[string]*FormFieldDateTime |
		map[string]*FormFieldFloat64 |
		map[string]*FormFieldInt |
		map[string]*FormFieldSelect |
		map[string]*FormFieldString |
		map[string]*FormFieldTime |
		map[string]*FormGroup |
		map[string]*FormSortAssocButton |
		map[string]*Option |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*CheckBox]any:
		return any(&stage.CheckBoxs).(*Type)
	case map[*FormDiv]any:
		return any(&stage.FormDivs).(*Type)
	case map[*FormEditAssocButton]any:
		return any(&stage.FormEditAssocButtons).(*Type)
	case map[*FormField]any:
		return any(&stage.FormFields).(*Type)
	case map[*FormFieldDate]any:
		return any(&stage.FormFieldDates).(*Type)
	case map[*FormFieldDateTime]any:
		return any(&stage.FormFieldDateTimes).(*Type)
	case map[*FormFieldFloat64]any:
		return any(&stage.FormFieldFloat64s).(*Type)
	case map[*FormFieldInt]any:
		return any(&stage.FormFieldInts).(*Type)
	case map[*FormFieldSelect]any:
		return any(&stage.FormFieldSelects).(*Type)
	case map[*FormFieldString]any:
		return any(&stage.FormFieldStrings).(*Type)
	case map[*FormFieldTime]any:
		return any(&stage.FormFieldTimes).(*Type)
	case map[*FormGroup]any:
		return any(&stage.FormGroups).(*Type)
	case map[*FormSortAssocButton]any:
		return any(&stage.FormSortAssocButtons).(*Type)
	case map[*Option]any:
		return any(&stage.Options).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*CheckBox:
		return any(&stage.CheckBoxs_mapString).(*Type)
	case map[string]*FormDiv:
		return any(&stage.FormDivs_mapString).(*Type)
	case map[string]*FormEditAssocButton:
		return any(&stage.FormEditAssocButtons_mapString).(*Type)
	case map[string]*FormField:
		return any(&stage.FormFields_mapString).(*Type)
	case map[string]*FormFieldDate:
		return any(&stage.FormFieldDates_mapString).(*Type)
	case map[string]*FormFieldDateTime:
		return any(&stage.FormFieldDateTimes_mapString).(*Type)
	case map[string]*FormFieldFloat64:
		return any(&stage.FormFieldFloat64s_mapString).(*Type)
	case map[string]*FormFieldInt:
		return any(&stage.FormFieldInts_mapString).(*Type)
	case map[string]*FormFieldSelect:
		return any(&stage.FormFieldSelects_mapString).(*Type)
	case map[string]*FormFieldString:
		return any(&stage.FormFieldStrings_mapString).(*Type)
	case map[string]*FormFieldTime:
		return any(&stage.FormFieldTimes_mapString).(*Type)
	case map[string]*FormGroup:
		return any(&stage.FormGroups_mapString).(*Type)
	case map[string]*FormSortAssocButton:
		return any(&stage.FormSortAssocButtons_mapString).(*Type)
	case map[string]*Option:
		return any(&stage.Options_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case CheckBox:
		return any(&stage.CheckBoxs).(*map[*Type]any)
	case FormDiv:
		return any(&stage.FormDivs).(*map[*Type]any)
	case FormEditAssocButton:
		return any(&stage.FormEditAssocButtons).(*map[*Type]any)
	case FormField:
		return any(&stage.FormFields).(*map[*Type]any)
	case FormFieldDate:
		return any(&stage.FormFieldDates).(*map[*Type]any)
	case FormFieldDateTime:
		return any(&stage.FormFieldDateTimes).(*map[*Type]any)
	case FormFieldFloat64:
		return any(&stage.FormFieldFloat64s).(*map[*Type]any)
	case FormFieldInt:
		return any(&stage.FormFieldInts).(*map[*Type]any)
	case FormFieldSelect:
		return any(&stage.FormFieldSelects).(*map[*Type]any)
	case FormFieldString:
		return any(&stage.FormFieldStrings).(*map[*Type]any)
	case FormFieldTime:
		return any(&stage.FormFieldTimes).(*map[*Type]any)
	case FormGroup:
		return any(&stage.FormGroups).(*map[*Type]any)
	case FormSortAssocButton:
		return any(&stage.FormSortAssocButtons).(*map[*Type]any)
	case Option:
		return any(&stage.Options).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *CheckBox:
		return any(&stage.CheckBoxs).(*map[Type]any)
	case *FormDiv:
		return any(&stage.FormDivs).(*map[Type]any)
	case *FormEditAssocButton:
		return any(&stage.FormEditAssocButtons).(*map[Type]any)
	case *FormField:
		return any(&stage.FormFields).(*map[Type]any)
	case *FormFieldDate:
		return any(&stage.FormFieldDates).(*map[Type]any)
	case *FormFieldDateTime:
		return any(&stage.FormFieldDateTimes).(*map[Type]any)
	case *FormFieldFloat64:
		return any(&stage.FormFieldFloat64s).(*map[Type]any)
	case *FormFieldInt:
		return any(&stage.FormFieldInts).(*map[Type]any)
	case *FormFieldSelect:
		return any(&stage.FormFieldSelects).(*map[Type]any)
	case *FormFieldString:
		return any(&stage.FormFieldStrings).(*map[Type]any)
	case *FormFieldTime:
		return any(&stage.FormFieldTimes).(*map[Type]any)
	case *FormGroup:
		return any(&stage.FormGroups).(*map[Type]any)
	case *FormSortAssocButton:
		return any(&stage.FormSortAssocButtons).(*map[Type]any)
	case *Option:
		return any(&stage.Options).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case CheckBox:
		return any(&stage.CheckBoxs_mapString).(*map[string]*Type)
	case FormDiv:
		return any(&stage.FormDivs_mapString).(*map[string]*Type)
	case FormEditAssocButton:
		return any(&stage.FormEditAssocButtons_mapString).(*map[string]*Type)
	case FormField:
		return any(&stage.FormFields_mapString).(*map[string]*Type)
	case FormFieldDate:
		return any(&stage.FormFieldDates_mapString).(*map[string]*Type)
	case FormFieldDateTime:
		return any(&stage.FormFieldDateTimes_mapString).(*map[string]*Type)
	case FormFieldFloat64:
		return any(&stage.FormFieldFloat64s_mapString).(*map[string]*Type)
	case FormFieldInt:
		return any(&stage.FormFieldInts_mapString).(*map[string]*Type)
	case FormFieldSelect:
		return any(&stage.FormFieldSelects_mapString).(*map[string]*Type)
	case FormFieldString:
		return any(&stage.FormFieldStrings_mapString).(*map[string]*Type)
	case FormFieldTime:
		return any(&stage.FormFieldTimes_mapString).(*map[string]*Type)
	case FormGroup:
		return any(&stage.FormGroups_mapString).(*map[string]*Type)
	case FormSortAssocButton:
		return any(&stage.FormSortAssocButtons_mapString).(*map[string]*Type)
	case Option:
		return any(&stage.Options_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case CheckBox:
		return any(&CheckBox{
			// Initialisation of associations
		}).(*Type)
	case FormDiv:
		return any(&FormDiv{
			// Initialisation of associations
			// field is initialized with an instance of FormField with the name of the field
			FormFields: []*FormField{{Name: "FormFields"}},
			// field is initialized with an instance of CheckBox with the name of the field
			CheckBoxs: []*CheckBox{{Name: "CheckBoxs"}},
			// field is initialized with an instance of FormEditAssocButton with the name of the field
			FormEditAssocButton: &FormEditAssocButton{Name: "FormEditAssocButton"},
			// field is initialized with an instance of FormSortAssocButton with the name of the field
			FormSortAssocButton: &FormSortAssocButton{Name: "FormSortAssocButton"},
		}).(*Type)
	case FormEditAssocButton:
		return any(&FormEditAssocButton{
			// Initialisation of associations
		}).(*Type)
	case FormField:
		return any(&FormField{
			// Initialisation of associations
			// field is initialized with an instance of FormFieldString with the name of the field
			FormFieldString: &FormFieldString{Name: "FormFieldString"},
			// field is initialized with an instance of FormFieldFloat64 with the name of the field
			FormFieldFloat64: &FormFieldFloat64{Name: "FormFieldFloat64"},
			// field is initialized with an instance of FormFieldInt with the name of the field
			FormFieldInt: &FormFieldInt{Name: "FormFieldInt"},
			// field is initialized with an instance of FormFieldDate with the name of the field
			FormFieldDate: &FormFieldDate{Name: "FormFieldDate"},
			// field is initialized with an instance of FormFieldTime with the name of the field
			FormFieldTime: &FormFieldTime{Name: "FormFieldTime"},
			// field is initialized with an instance of FormFieldDateTime with the name of the field
			FormFieldDateTime: &FormFieldDateTime{Name: "FormFieldDateTime"},
			// field is initialized with an instance of FormFieldSelect with the name of the field
			FormFieldSelect: &FormFieldSelect{Name: "FormFieldSelect"},
		}).(*Type)
	case FormFieldDate:
		return any(&FormFieldDate{
			// Initialisation of associations
		}).(*Type)
	case FormFieldDateTime:
		return any(&FormFieldDateTime{
			// Initialisation of associations
		}).(*Type)
	case FormFieldFloat64:
		return any(&FormFieldFloat64{
			// Initialisation of associations
		}).(*Type)
	case FormFieldInt:
		return any(&FormFieldInt{
			// Initialisation of associations
		}).(*Type)
	case FormFieldSelect:
		return any(&FormFieldSelect{
			// Initialisation of associations
			// field is initialized with an instance of Option with the name of the field
			Value: &Option{Name: "Value"},
			// field is initialized with an instance of Option with the name of the field
			Options: []*Option{{Name: "Options"}},
		}).(*Type)
	case FormFieldString:
		return any(&FormFieldString{
			// Initialisation of associations
		}).(*Type)
	case FormFieldTime:
		return any(&FormFieldTime{
			// Initialisation of associations
		}).(*Type)
	case FormGroup:
		return any(&FormGroup{
			// Initialisation of associations
			// field is initialized with an instance of FormDiv with the name of the field
			FormDivs: []*FormDiv{{Name: "FormDivs"}},
		}).(*Type)
	case FormSortAssocButton:
		return any(&FormSortAssocButton{
			// Initialisation of associations
		}).(*Type)
	case Option:
		return any(&Option{
			// Initialisation of associations
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of CheckBox
	case CheckBox:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormDiv
	case FormDiv:
		switch fieldname {
		// insertion point for per direct association field
		case "FormEditAssocButton":
			res := make(map[*FormEditAssocButton][]*FormDiv)
			for formdiv := range stage.FormDivs {
				if formdiv.FormEditAssocButton != nil {
					formeditassocbutton_ := formdiv.FormEditAssocButton
					var formdivs []*FormDiv
					_, ok := res[formeditassocbutton_]
					if ok {
						formdivs = res[formeditassocbutton_]
					} else {
						formdivs = make([]*FormDiv, 0)
					}
					formdivs = append(formdivs, formdiv)
					res[formeditassocbutton_] = formdivs
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormSortAssocButton":
			res := make(map[*FormSortAssocButton][]*FormDiv)
			for formdiv := range stage.FormDivs {
				if formdiv.FormSortAssocButton != nil {
					formsortassocbutton_ := formdiv.FormSortAssocButton
					var formdivs []*FormDiv
					_, ok := res[formsortassocbutton_]
					if ok {
						formdivs = res[formsortassocbutton_]
					} else {
						formdivs = make([]*FormDiv, 0)
					}
					formdivs = append(formdivs, formdiv)
					res[formsortassocbutton_] = formdivs
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of FormEditAssocButton
	case FormEditAssocButton:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormField
	case FormField:
		switch fieldname {
		// insertion point for per direct association field
		case "FormFieldString":
			res := make(map[*FormFieldString][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldString != nil {
					formfieldstring_ := formfield.FormFieldString
					var formfields []*FormField
					_, ok := res[formfieldstring_]
					if ok {
						formfields = res[formfieldstring_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfieldstring_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldFloat64":
			res := make(map[*FormFieldFloat64][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldFloat64 != nil {
					formfieldfloat64_ := formfield.FormFieldFloat64
					var formfields []*FormField
					_, ok := res[formfieldfloat64_]
					if ok {
						formfields = res[formfieldfloat64_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfieldfloat64_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldInt":
			res := make(map[*FormFieldInt][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldInt != nil {
					formfieldint_ := formfield.FormFieldInt
					var formfields []*FormField
					_, ok := res[formfieldint_]
					if ok {
						formfields = res[formfieldint_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfieldint_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldDate":
			res := make(map[*FormFieldDate][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldDate != nil {
					formfielddate_ := formfield.FormFieldDate
					var formfields []*FormField
					_, ok := res[formfielddate_]
					if ok {
						formfields = res[formfielddate_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfielddate_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldTime":
			res := make(map[*FormFieldTime][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldTime != nil {
					formfieldtime_ := formfield.FormFieldTime
					var formfields []*FormField
					_, ok := res[formfieldtime_]
					if ok {
						formfields = res[formfieldtime_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfieldtime_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldDateTime":
			res := make(map[*FormFieldDateTime][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldDateTime != nil {
					formfielddatetime_ := formfield.FormFieldDateTime
					var formfields []*FormField
					_, ok := res[formfielddatetime_]
					if ok {
						formfields = res[formfielddatetime_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfielddatetime_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldSelect":
			res := make(map[*FormFieldSelect][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldSelect != nil {
					formfieldselect_ := formfield.FormFieldSelect
					var formfields []*FormField
					_, ok := res[formfieldselect_]
					if ok {
						formfields = res[formfieldselect_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfieldselect_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of FormFieldDate
	case FormFieldDate:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldDateTime
	case FormFieldDateTime:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldFloat64
	case FormFieldFloat64:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldInt
	case FormFieldInt:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldSelect
	case FormFieldSelect:
		switch fieldname {
		// insertion point for per direct association field
		case "Value":
			res := make(map[*Option][]*FormFieldSelect)
			for formfieldselect := range stage.FormFieldSelects {
				if formfieldselect.Value != nil {
					option_ := formfieldselect.Value
					var formfieldselects []*FormFieldSelect
					_, ok := res[option_]
					if ok {
						formfieldselects = res[option_]
					} else {
						formfieldselects = make([]*FormFieldSelect, 0)
					}
					formfieldselects = append(formfieldselects, formfieldselect)
					res[option_] = formfieldselects
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of FormFieldString
	case FormFieldString:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldTime
	case FormFieldTime:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormGroup
	case FormGroup:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormSortAssocButton
	case FormSortAssocButton:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Option
	case Option:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of CheckBox
	case CheckBox:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormDiv
	case FormDiv:
		switch fieldname {
		// insertion point for per direct association field
		case "FormFields":
			res := make(map[*FormField]*FormDiv)
			for formdiv := range stage.FormDivs {
				for _, formfield_ := range formdiv.FormFields {
					res[formfield_] = formdiv
				}
			}
			return any(res).(map[*End]*Start)
		case "CheckBoxs":
			res := make(map[*CheckBox]*FormDiv)
			for formdiv := range stage.FormDivs {
				for _, checkbox_ := range formdiv.CheckBoxs {
					res[checkbox_] = formdiv
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of FormEditAssocButton
	case FormEditAssocButton:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormField
	case FormField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldDate
	case FormFieldDate:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldDateTime
	case FormFieldDateTime:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldFloat64
	case FormFieldFloat64:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldInt
	case FormFieldInt:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldSelect
	case FormFieldSelect:
		switch fieldname {
		// insertion point for per direct association field
		case "Options":
			res := make(map[*Option]*FormFieldSelect)
			for formfieldselect := range stage.FormFieldSelects {
				for _, option_ := range formfieldselect.Options {
					res[option_] = formfieldselect
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of FormFieldString
	case FormFieldString:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldTime
	case FormFieldTime:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormGroup
	case FormGroup:
		switch fieldname {
		// insertion point for per direct association field
		case "FormDivs":
			res := make(map[*FormDiv]*FormGroup)
			for formgroup := range stage.FormGroups {
				for _, formdiv_ := range formgroup.FormDivs {
					res[formdiv_] = formgroup
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of FormSortAssocButton
	case FormSortAssocButton:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Option
	case Option:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case CheckBox:
		res = "CheckBox"
	case FormDiv:
		res = "FormDiv"
	case FormEditAssocButton:
		res = "FormEditAssocButton"
	case FormField:
		res = "FormField"
	case FormFieldDate:
		res = "FormFieldDate"
	case FormFieldDateTime:
		res = "FormFieldDateTime"
	case FormFieldFloat64:
		res = "FormFieldFloat64"
	case FormFieldInt:
		res = "FormFieldInt"
	case FormFieldSelect:
		res = "FormFieldSelect"
	case FormFieldString:
		res = "FormFieldString"
	case FormFieldTime:
		res = "FormFieldTime"
	case FormGroup:
		res = "FormGroup"
	case FormSortAssocButton:
		res = "FormSortAssocButton"
	case Option:
		res = "Option"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case CheckBox:
		res = []string{"Name", "Value"}
	case FormDiv:
		res = []string{"Name", "FormFields", "CheckBoxs", "FormEditAssocButton", "FormSortAssocButton"}
	case FormEditAssocButton:
		res = []string{"Name", "Label"}
	case FormField:
		res = []string{"Name", "InputTypeEnum", "Label", "Placeholder", "FormFieldString", "FormFieldFloat64", "FormFieldInt", "FormFieldDate", "FormFieldTime", "FormFieldDateTime", "FormFieldSelect", "HasBespokeWidth", "BespokeWidthPx"}
	case FormFieldDate:
		res = []string{"Name", "Value"}
	case FormFieldDateTime:
		res = []string{"Name", "Value"}
	case FormFieldFloat64:
		res = []string{"Name", "Value", "HasMinValidator", "MinValue", "HasMaxValidator", "MaxValue"}
	case FormFieldInt:
		res = []string{"Name", "Value", "HasMinValidator", "MinValue", "HasMaxValidator", "MaxValue"}
	case FormFieldSelect:
		res = []string{"Name", "Value", "Options", "CanBeEmpty"}
	case FormFieldString:
		res = []string{"Name", "Value"}
	case FormFieldTime:
		res = []string{"Name", "Value", "Step"}
	case FormGroup:
		res = []string{"Name", "Label", "FormDivs"}
	case FormSortAssocButton:
		res = []string{"Name", "Label"}
	case Option:
		res = []string{"Name"}
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *CheckBox:
		res = []string{"Name", "Value"}
	case *FormDiv:
		res = []string{"Name", "FormFields", "CheckBoxs", "FormEditAssocButton", "FormSortAssocButton"}
	case *FormEditAssocButton:
		res = []string{"Name", "Label"}
	case *FormField:
		res = []string{"Name", "InputTypeEnum", "Label", "Placeholder", "FormFieldString", "FormFieldFloat64", "FormFieldInt", "FormFieldDate", "FormFieldTime", "FormFieldDateTime", "FormFieldSelect", "HasBespokeWidth", "BespokeWidthPx"}
	case *FormFieldDate:
		res = []string{"Name", "Value"}
	case *FormFieldDateTime:
		res = []string{"Name", "Value"}
	case *FormFieldFloat64:
		res = []string{"Name", "Value", "HasMinValidator", "MinValue", "HasMaxValidator", "MaxValue"}
	case *FormFieldInt:
		res = []string{"Name", "Value", "HasMinValidator", "MinValue", "HasMaxValidator", "MaxValue"}
	case *FormFieldSelect:
		res = []string{"Name", "Value", "Options", "CanBeEmpty"}
	case *FormFieldString:
		res = []string{"Name", "Value"}
	case *FormFieldTime:
		res = []string{"Name", "Value", "Step"}
	case *FormGroup:
		res = []string{"Name", "Label", "FormDivs"}
	case *FormSortAssocButton:
		res = []string{"Name", "Label"}
	case *Option:
		res = []string{"Name"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *CheckBox:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%t", inferedInstance.Value)
		}
	case *FormDiv:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "FormFields":
			for idx, __instance__ := range inferedInstance.FormFields {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "CheckBoxs":
			for idx, __instance__ := range inferedInstance.CheckBoxs {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "FormEditAssocButton":
			if inferedInstance.FormEditAssocButton != nil {
				res = inferedInstance.FormEditAssocButton.Name
			}
		case "FormSortAssocButton":
			if inferedInstance.FormSortAssocButton != nil {
				res = inferedInstance.FormSortAssocButton.Name
			}
		}
	case *FormEditAssocButton:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Label":
			res = inferedInstance.Label
		}
	case *FormField:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "InputTypeEnum":
			enum := inferedInstance.InputTypeEnum
			res = enum.ToCodeString()
		case "Label":
			res = inferedInstance.Label
		case "Placeholder":
			res = inferedInstance.Placeholder
		case "FormFieldString":
			if inferedInstance.FormFieldString != nil {
				res = inferedInstance.FormFieldString.Name
			}
		case "FormFieldFloat64":
			if inferedInstance.FormFieldFloat64 != nil {
				res = inferedInstance.FormFieldFloat64.Name
			}
		case "FormFieldInt":
			if inferedInstance.FormFieldInt != nil {
				res = inferedInstance.FormFieldInt.Name
			}
		case "FormFieldDate":
			if inferedInstance.FormFieldDate != nil {
				res = inferedInstance.FormFieldDate.Name
			}
		case "FormFieldTime":
			if inferedInstance.FormFieldTime != nil {
				res = inferedInstance.FormFieldTime.Name
			}
		case "FormFieldDateTime":
			if inferedInstance.FormFieldDateTime != nil {
				res = inferedInstance.FormFieldDateTime.Name
			}
		case "FormFieldSelect":
			if inferedInstance.FormFieldSelect != nil {
				res = inferedInstance.FormFieldSelect.Name
			}
		case "HasBespokeWidth":
			res = fmt.Sprintf("%t", inferedInstance.HasBespokeWidth)
		case "BespokeWidthPx":
			res = fmt.Sprintf("%d", inferedInstance.BespokeWidthPx)
		}
	case *FormFieldDate:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value.String()
		}
	case *FormFieldDateTime:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value.String()
		}
	case *FormFieldFloat64:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%f", inferedInstance.Value)
		case "HasMinValidator":
			res = fmt.Sprintf("%t", inferedInstance.HasMinValidator)
		case "MinValue":
			res = fmt.Sprintf("%f", inferedInstance.MinValue)
		case "HasMaxValidator":
			res = fmt.Sprintf("%t", inferedInstance.HasMaxValidator)
		case "MaxValue":
			res = fmt.Sprintf("%f", inferedInstance.MaxValue)
		}
	case *FormFieldInt:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%d", inferedInstance.Value)
		case "HasMinValidator":
			res = fmt.Sprintf("%t", inferedInstance.HasMinValidator)
		case "MinValue":
			res = fmt.Sprintf("%d", inferedInstance.MinValue)
		case "HasMaxValidator":
			res = fmt.Sprintf("%t", inferedInstance.HasMaxValidator)
		case "MaxValue":
			res = fmt.Sprintf("%d", inferedInstance.MaxValue)
		}
	case *FormFieldSelect:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			if inferedInstance.Value != nil {
				res = inferedInstance.Value.Name
			}
		case "Options":
			for idx, __instance__ := range inferedInstance.Options {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "CanBeEmpty":
			res = fmt.Sprintf("%t", inferedInstance.CanBeEmpty)
		}
	case *FormFieldString:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value
		}
	case *FormFieldTime:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value.String()
		case "Step":
			res = fmt.Sprintf("%f", inferedInstance.Step)
		}
	case *FormGroup:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Label":
			res = inferedInstance.Label
		case "FormDivs":
			for idx, __instance__ := range inferedInstance.FormDivs {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *FormSortAssocButton:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Label":
			res = inferedInstance.Label
		}
	case *Option:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	default:
		_ = inferedInstance	
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case CheckBox:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%t", inferedInstance.Value)
		}
	case FormDiv:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "FormFields":
			for idx, __instance__ := range inferedInstance.FormFields {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "CheckBoxs":
			for idx, __instance__ := range inferedInstance.CheckBoxs {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "FormEditAssocButton":
			if inferedInstance.FormEditAssocButton != nil {
				res = inferedInstance.FormEditAssocButton.Name
			}
		case "FormSortAssocButton":
			if inferedInstance.FormSortAssocButton != nil {
				res = inferedInstance.FormSortAssocButton.Name
			}
		}
	case FormEditAssocButton:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Label":
			res = inferedInstance.Label
		}
	case FormField:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "InputTypeEnum":
			enum := inferedInstance.InputTypeEnum
			res = enum.ToCodeString()
		case "Label":
			res = inferedInstance.Label
		case "Placeholder":
			res = inferedInstance.Placeholder
		case "FormFieldString":
			if inferedInstance.FormFieldString != nil {
				res = inferedInstance.FormFieldString.Name
			}
		case "FormFieldFloat64":
			if inferedInstance.FormFieldFloat64 != nil {
				res = inferedInstance.FormFieldFloat64.Name
			}
		case "FormFieldInt":
			if inferedInstance.FormFieldInt != nil {
				res = inferedInstance.FormFieldInt.Name
			}
		case "FormFieldDate":
			if inferedInstance.FormFieldDate != nil {
				res = inferedInstance.FormFieldDate.Name
			}
		case "FormFieldTime":
			if inferedInstance.FormFieldTime != nil {
				res = inferedInstance.FormFieldTime.Name
			}
		case "FormFieldDateTime":
			if inferedInstance.FormFieldDateTime != nil {
				res = inferedInstance.FormFieldDateTime.Name
			}
		case "FormFieldSelect":
			if inferedInstance.FormFieldSelect != nil {
				res = inferedInstance.FormFieldSelect.Name
			}
		case "HasBespokeWidth":
			res = fmt.Sprintf("%t", inferedInstance.HasBespokeWidth)
		case "BespokeWidthPx":
			res = fmt.Sprintf("%d", inferedInstance.BespokeWidthPx)
		}
	case FormFieldDate:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value.String()
		}
	case FormFieldDateTime:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value.String()
		}
	case FormFieldFloat64:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%f", inferedInstance.Value)
		case "HasMinValidator":
			res = fmt.Sprintf("%t", inferedInstance.HasMinValidator)
		case "MinValue":
			res = fmt.Sprintf("%f", inferedInstance.MinValue)
		case "HasMaxValidator":
			res = fmt.Sprintf("%t", inferedInstance.HasMaxValidator)
		case "MaxValue":
			res = fmt.Sprintf("%f", inferedInstance.MaxValue)
		}
	case FormFieldInt:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%d", inferedInstance.Value)
		case "HasMinValidator":
			res = fmt.Sprintf("%t", inferedInstance.HasMinValidator)
		case "MinValue":
			res = fmt.Sprintf("%d", inferedInstance.MinValue)
		case "HasMaxValidator":
			res = fmt.Sprintf("%t", inferedInstance.HasMaxValidator)
		case "MaxValue":
			res = fmt.Sprintf("%d", inferedInstance.MaxValue)
		}
	case FormFieldSelect:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			if inferedInstance.Value != nil {
				res = inferedInstance.Value.Name
			}
		case "Options":
			for idx, __instance__ := range inferedInstance.Options {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "CanBeEmpty":
			res = fmt.Sprintf("%t", inferedInstance.CanBeEmpty)
		}
	case FormFieldString:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value
		}
	case FormFieldTime:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value.String()
		case "Step":
			res = fmt.Sprintf("%f", inferedInstance.Step)
		}
	case FormGroup:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Label":
			res = inferedInstance.Label
		case "FormDivs":
			for idx, __instance__ := range inferedInstance.FormDivs {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case FormSortAssocButton:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Label":
			res = inferedInstance.Label
		}
	case Option:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	default:
		_ = inferedInstance	
	}
	return
}

// Last line of the template
