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
	GongBasicFields           map[*GongBasicField]any
	GongBasicFields_mapString map[string]*GongBasicField

	OnAfterGongBasicFieldCreateCallback OnAfterCreateInterface[GongBasicField]
	OnAfterGongBasicFieldUpdateCallback OnAfterUpdateInterface[GongBasicField]
	OnAfterGongBasicFieldDeleteCallback OnAfterDeleteInterface[GongBasicField]
	OnAfterGongBasicFieldReadCallback   OnAfterReadInterface[GongBasicField]

	GongEnums           map[*GongEnum]any
	GongEnums_mapString map[string]*GongEnum

	OnAfterGongEnumCreateCallback OnAfterCreateInterface[GongEnum]
	OnAfterGongEnumUpdateCallback OnAfterUpdateInterface[GongEnum]
	OnAfterGongEnumDeleteCallback OnAfterDeleteInterface[GongEnum]
	OnAfterGongEnumReadCallback   OnAfterReadInterface[GongEnum]

	GongEnumValues           map[*GongEnumValue]any
	GongEnumValues_mapString map[string]*GongEnumValue

	OnAfterGongEnumValueCreateCallback OnAfterCreateInterface[GongEnumValue]
	OnAfterGongEnumValueUpdateCallback OnAfterUpdateInterface[GongEnumValue]
	OnAfterGongEnumValueDeleteCallback OnAfterDeleteInterface[GongEnumValue]
	OnAfterGongEnumValueReadCallback   OnAfterReadInterface[GongEnumValue]

	GongLinks           map[*GongLink]any
	GongLinks_mapString map[string]*GongLink

	OnAfterGongLinkCreateCallback OnAfterCreateInterface[GongLink]
	OnAfterGongLinkUpdateCallback OnAfterUpdateInterface[GongLink]
	OnAfterGongLinkDeleteCallback OnAfterDeleteInterface[GongLink]
	OnAfterGongLinkReadCallback   OnAfterReadInterface[GongLink]

	GongNotes           map[*GongNote]any
	GongNotes_mapString map[string]*GongNote

	OnAfterGongNoteCreateCallback OnAfterCreateInterface[GongNote]
	OnAfterGongNoteUpdateCallback OnAfterUpdateInterface[GongNote]
	OnAfterGongNoteDeleteCallback OnAfterDeleteInterface[GongNote]
	OnAfterGongNoteReadCallback   OnAfterReadInterface[GongNote]

	GongStructs           map[*GongStruct]any
	GongStructs_mapString map[string]*GongStruct

	OnAfterGongStructCreateCallback OnAfterCreateInterface[GongStruct]
	OnAfterGongStructUpdateCallback OnAfterUpdateInterface[GongStruct]
	OnAfterGongStructDeleteCallback OnAfterDeleteInterface[GongStruct]
	OnAfterGongStructReadCallback   OnAfterReadInterface[GongStruct]

	GongTimeFields           map[*GongTimeField]any
	GongTimeFields_mapString map[string]*GongTimeField

	OnAfterGongTimeFieldCreateCallback OnAfterCreateInterface[GongTimeField]
	OnAfterGongTimeFieldUpdateCallback OnAfterUpdateInterface[GongTimeField]
	OnAfterGongTimeFieldDeleteCallback OnAfterDeleteInterface[GongTimeField]
	OnAfterGongTimeFieldReadCallback   OnAfterReadInterface[GongTimeField]

	Metas           map[*Meta]any
	Metas_mapString map[string]*Meta

	OnAfterMetaCreateCallback OnAfterCreateInterface[Meta]
	OnAfterMetaUpdateCallback OnAfterUpdateInterface[Meta]
	OnAfterMetaDeleteCallback OnAfterDeleteInterface[Meta]
	OnAfterMetaReadCallback   OnAfterReadInterface[Meta]

	MetaReferences           map[*MetaReference]any
	MetaReferences_mapString map[string]*MetaReference

	OnAfterMetaReferenceCreateCallback OnAfterCreateInterface[MetaReference]
	OnAfterMetaReferenceUpdateCallback OnAfterUpdateInterface[MetaReference]
	OnAfterMetaReferenceDeleteCallback OnAfterDeleteInterface[MetaReference]
	OnAfterMetaReferenceReadCallback   OnAfterReadInterface[MetaReference]

	ModelPkgs           map[*ModelPkg]any
	ModelPkgs_mapString map[string]*ModelPkg

	OnAfterModelPkgCreateCallback OnAfterCreateInterface[ModelPkg]
	OnAfterModelPkgUpdateCallback OnAfterUpdateInterface[ModelPkg]
	OnAfterModelPkgDeleteCallback OnAfterDeleteInterface[ModelPkg]
	OnAfterModelPkgReadCallback   OnAfterReadInterface[ModelPkg]

	PointerToGongStructFields           map[*PointerToGongStructField]any
	PointerToGongStructFields_mapString map[string]*PointerToGongStructField

	OnAfterPointerToGongStructFieldCreateCallback OnAfterCreateInterface[PointerToGongStructField]
	OnAfterPointerToGongStructFieldUpdateCallback OnAfterUpdateInterface[PointerToGongStructField]
	OnAfterPointerToGongStructFieldDeleteCallback OnAfterDeleteInterface[PointerToGongStructField]
	OnAfterPointerToGongStructFieldReadCallback   OnAfterReadInterface[PointerToGongStructField]

	SliceOfPointerToGongStructFields           map[*SliceOfPointerToGongStructField]any
	SliceOfPointerToGongStructFields_mapString map[string]*SliceOfPointerToGongStructField

	OnAfterSliceOfPointerToGongStructFieldCreateCallback OnAfterCreateInterface[SliceOfPointerToGongStructField]
	OnAfterSliceOfPointerToGongStructFieldUpdateCallback OnAfterUpdateInterface[SliceOfPointerToGongStructField]
	OnAfterSliceOfPointerToGongStructFieldDeleteCallback OnAfterDeleteInterface[SliceOfPointerToGongStructField]
	OnAfterSliceOfPointerToGongStructFieldReadCallback   OnAfterReadInterface[SliceOfPointerToGongStructField]

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
	CommitGongBasicField(gongbasicfield *GongBasicField)
	CheckoutGongBasicField(gongbasicfield *GongBasicField)
	CommitGongEnum(gongenum *GongEnum)
	CheckoutGongEnum(gongenum *GongEnum)
	CommitGongEnumValue(gongenumvalue *GongEnumValue)
	CheckoutGongEnumValue(gongenumvalue *GongEnumValue)
	CommitGongLink(gonglink *GongLink)
	CheckoutGongLink(gonglink *GongLink)
	CommitGongNote(gongnote *GongNote)
	CheckoutGongNote(gongnote *GongNote)
	CommitGongStruct(gongstruct *GongStruct)
	CheckoutGongStruct(gongstruct *GongStruct)
	CommitGongTimeField(gongtimefield *GongTimeField)
	CheckoutGongTimeField(gongtimefield *GongTimeField)
	CommitMeta(meta *Meta)
	CheckoutMeta(meta *Meta)
	CommitMetaReference(metareference *MetaReference)
	CheckoutMetaReference(metareference *MetaReference)
	CommitModelPkg(modelpkg *ModelPkg)
	CheckoutModelPkg(modelpkg *ModelPkg)
	CommitPointerToGongStructField(pointertogongstructfield *PointerToGongStructField)
	CheckoutPointerToGongStructField(pointertogongstructfield *PointerToGongStructField)
	CommitSliceOfPointerToGongStructField(sliceofpointertogongstructfield *SliceOfPointerToGongStructField)
	CheckoutSliceOfPointerToGongStructField(sliceofpointertogongstructfield *SliceOfPointerToGongStructField)
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
		GongBasicFields:           make(map[*GongBasicField]any),
		GongBasicFields_mapString: make(map[string]*GongBasicField),

		GongEnums:           make(map[*GongEnum]any),
		GongEnums_mapString: make(map[string]*GongEnum),

		GongEnumValues:           make(map[*GongEnumValue]any),
		GongEnumValues_mapString: make(map[string]*GongEnumValue),

		GongLinks:           make(map[*GongLink]any),
		GongLinks_mapString: make(map[string]*GongLink),

		GongNotes:           make(map[*GongNote]any),
		GongNotes_mapString: make(map[string]*GongNote),

		GongStructs:           make(map[*GongStruct]any),
		GongStructs_mapString: make(map[string]*GongStruct),

		GongTimeFields:           make(map[*GongTimeField]any),
		GongTimeFields_mapString: make(map[string]*GongTimeField),

		Metas:           make(map[*Meta]any),
		Metas_mapString: make(map[string]*Meta),

		MetaReferences:           make(map[*MetaReference]any),
		MetaReferences_mapString: make(map[string]*MetaReference),

		ModelPkgs:           make(map[*ModelPkg]any),
		ModelPkgs_mapString: make(map[string]*ModelPkg),

		PointerToGongStructFields:           make(map[*PointerToGongStructField]any),
		PointerToGongStructFields_mapString: make(map[string]*PointerToGongStructField),

		SliceOfPointerToGongStructFields:           make(map[*SliceOfPointerToGongStructField]any),
		SliceOfPointerToGongStructFields_mapString: make(map[string]*SliceOfPointerToGongStructField),

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
	stage.Map_GongStructName_InstancesNb["GongBasicField"] = len(stage.GongBasicFields)
	stage.Map_GongStructName_InstancesNb["GongEnum"] = len(stage.GongEnums)
	stage.Map_GongStructName_InstancesNb["GongEnumValue"] = len(stage.GongEnumValues)
	stage.Map_GongStructName_InstancesNb["GongLink"] = len(stage.GongLinks)
	stage.Map_GongStructName_InstancesNb["GongNote"] = len(stage.GongNotes)
	stage.Map_GongStructName_InstancesNb["GongStruct"] = len(stage.GongStructs)
	stage.Map_GongStructName_InstancesNb["GongTimeField"] = len(stage.GongTimeFields)
	stage.Map_GongStructName_InstancesNb["Meta"] = len(stage.Metas)
	stage.Map_GongStructName_InstancesNb["MetaReference"] = len(stage.MetaReferences)
	stage.Map_GongStructName_InstancesNb["ModelPkg"] = len(stage.ModelPkgs)
	stage.Map_GongStructName_InstancesNb["PointerToGongStructField"] = len(stage.PointerToGongStructFields)
	stage.Map_GongStructName_InstancesNb["SliceOfPointerToGongStructField"] = len(stage.SliceOfPointerToGongStructFields)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["GongBasicField"] = len(stage.GongBasicFields)
	stage.Map_GongStructName_InstancesNb["GongEnum"] = len(stage.GongEnums)
	stage.Map_GongStructName_InstancesNb["GongEnumValue"] = len(stage.GongEnumValues)
	stage.Map_GongStructName_InstancesNb["GongLink"] = len(stage.GongLinks)
	stage.Map_GongStructName_InstancesNb["GongNote"] = len(stage.GongNotes)
	stage.Map_GongStructName_InstancesNb["GongStruct"] = len(stage.GongStructs)
	stage.Map_GongStructName_InstancesNb["GongTimeField"] = len(stage.GongTimeFields)
	stage.Map_GongStructName_InstancesNb["Meta"] = len(stage.Metas)
	stage.Map_GongStructName_InstancesNb["MetaReference"] = len(stage.MetaReferences)
	stage.Map_GongStructName_InstancesNb["ModelPkg"] = len(stage.ModelPkgs)
	stage.Map_GongStructName_InstancesNb["PointerToGongStructField"] = len(stage.PointerToGongStructFields)
	stage.Map_GongStructName_InstancesNb["SliceOfPointerToGongStructField"] = len(stage.SliceOfPointerToGongStructFields)

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
// Stage puts gongbasicfield to the model stage
func (gongbasicfield *GongBasicField) Stage(stage *StageStruct) *GongBasicField {
	stage.GongBasicFields[gongbasicfield] = __member
	stage.GongBasicFields_mapString[gongbasicfield.Name] = gongbasicfield

	return gongbasicfield
}

// Unstage removes gongbasicfield off the model stage
func (gongbasicfield *GongBasicField) Unstage(stage *StageStruct) *GongBasicField {
	delete(stage.GongBasicFields, gongbasicfield)
	delete(stage.GongBasicFields_mapString, gongbasicfield.Name)
	return gongbasicfield
}

// commit gongbasicfield to the back repo (if it is already staged)
func (gongbasicfield *GongBasicField) Commit(stage *StageStruct) *GongBasicField {
	if _, ok := stage.GongBasicFields[gongbasicfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongBasicField(gongbasicfield)
		}
	}
	return gongbasicfield
}

func (gongbasicfield *GongBasicField) CommitVoid(stage *StageStruct) {
	gongbasicfield.Commit(stage)
}

// Checkout gongbasicfield to the back repo (if it is already staged)
func (gongbasicfield *GongBasicField) Checkout(stage *StageStruct) *GongBasicField {
	if _, ok := stage.GongBasicFields[gongbasicfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongBasicField(gongbasicfield)
		}
	}
	return gongbasicfield
}

// for satisfaction of GongStruct interface
func (gongbasicfield *GongBasicField) GetName() (res string) {
	return gongbasicfield.Name
}

// Stage puts gongenum to the model stage
func (gongenum *GongEnum) Stage(stage *StageStruct) *GongEnum {
	stage.GongEnums[gongenum] = __member
	stage.GongEnums_mapString[gongenum.Name] = gongenum

	return gongenum
}

// Unstage removes gongenum off the model stage
func (gongenum *GongEnum) Unstage(stage *StageStruct) *GongEnum {
	delete(stage.GongEnums, gongenum)
	delete(stage.GongEnums_mapString, gongenum.Name)
	return gongenum
}

// commit gongenum to the back repo (if it is already staged)
func (gongenum *GongEnum) Commit(stage *StageStruct) *GongEnum {
	if _, ok := stage.GongEnums[gongenum]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongEnum(gongenum)
		}
	}
	return gongenum
}

func (gongenum *GongEnum) CommitVoid(stage *StageStruct) {
	gongenum.Commit(stage)
}

// Checkout gongenum to the back repo (if it is already staged)
func (gongenum *GongEnum) Checkout(stage *StageStruct) *GongEnum {
	if _, ok := stage.GongEnums[gongenum]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongEnum(gongenum)
		}
	}
	return gongenum
}

// for satisfaction of GongStruct interface
func (gongenum *GongEnum) GetName() (res string) {
	return gongenum.Name
}

// Stage puts gongenumvalue to the model stage
func (gongenumvalue *GongEnumValue) Stage(stage *StageStruct) *GongEnumValue {
	stage.GongEnumValues[gongenumvalue] = __member
	stage.GongEnumValues_mapString[gongenumvalue.Name] = gongenumvalue

	return gongenumvalue
}

// Unstage removes gongenumvalue off the model stage
func (gongenumvalue *GongEnumValue) Unstage(stage *StageStruct) *GongEnumValue {
	delete(stage.GongEnumValues, gongenumvalue)
	delete(stage.GongEnumValues_mapString, gongenumvalue.Name)
	return gongenumvalue
}

// commit gongenumvalue to the back repo (if it is already staged)
func (gongenumvalue *GongEnumValue) Commit(stage *StageStruct) *GongEnumValue {
	if _, ok := stage.GongEnumValues[gongenumvalue]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongEnumValue(gongenumvalue)
		}
	}
	return gongenumvalue
}

func (gongenumvalue *GongEnumValue) CommitVoid(stage *StageStruct) {
	gongenumvalue.Commit(stage)
}

// Checkout gongenumvalue to the back repo (if it is already staged)
func (gongenumvalue *GongEnumValue) Checkout(stage *StageStruct) *GongEnumValue {
	if _, ok := stage.GongEnumValues[gongenumvalue]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongEnumValue(gongenumvalue)
		}
	}
	return gongenumvalue
}

// for satisfaction of GongStruct interface
func (gongenumvalue *GongEnumValue) GetName() (res string) {
	return gongenumvalue.Name
}

// Stage puts gonglink to the model stage
func (gonglink *GongLink) Stage(stage *StageStruct) *GongLink {
	stage.GongLinks[gonglink] = __member
	stage.GongLinks_mapString[gonglink.Name] = gonglink

	return gonglink
}

// Unstage removes gonglink off the model stage
func (gonglink *GongLink) Unstage(stage *StageStruct) *GongLink {
	delete(stage.GongLinks, gonglink)
	delete(stage.GongLinks_mapString, gonglink.Name)
	return gonglink
}

// commit gonglink to the back repo (if it is already staged)
func (gonglink *GongLink) Commit(stage *StageStruct) *GongLink {
	if _, ok := stage.GongLinks[gonglink]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongLink(gonglink)
		}
	}
	return gonglink
}

func (gonglink *GongLink) CommitVoid(stage *StageStruct) {
	gonglink.Commit(stage)
}

// Checkout gonglink to the back repo (if it is already staged)
func (gonglink *GongLink) Checkout(stage *StageStruct) *GongLink {
	if _, ok := stage.GongLinks[gonglink]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongLink(gonglink)
		}
	}
	return gonglink
}

// for satisfaction of GongStruct interface
func (gonglink *GongLink) GetName() (res string) {
	return gonglink.Name
}

// Stage puts gongnote to the model stage
func (gongnote *GongNote) Stage(stage *StageStruct) *GongNote {
	stage.GongNotes[gongnote] = __member
	stage.GongNotes_mapString[gongnote.Name] = gongnote

	return gongnote
}

// Unstage removes gongnote off the model stage
func (gongnote *GongNote) Unstage(stage *StageStruct) *GongNote {
	delete(stage.GongNotes, gongnote)
	delete(stage.GongNotes_mapString, gongnote.Name)
	return gongnote
}

// commit gongnote to the back repo (if it is already staged)
func (gongnote *GongNote) Commit(stage *StageStruct) *GongNote {
	if _, ok := stage.GongNotes[gongnote]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongNote(gongnote)
		}
	}
	return gongnote
}

func (gongnote *GongNote) CommitVoid(stage *StageStruct) {
	gongnote.Commit(stage)
}

// Checkout gongnote to the back repo (if it is already staged)
func (gongnote *GongNote) Checkout(stage *StageStruct) *GongNote {
	if _, ok := stage.GongNotes[gongnote]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongNote(gongnote)
		}
	}
	return gongnote
}

// for satisfaction of GongStruct interface
func (gongnote *GongNote) GetName() (res string) {
	return gongnote.Name
}

// Stage puts gongstruct to the model stage
func (gongstruct *GongStruct) Stage(stage *StageStruct) *GongStruct {
	stage.GongStructs[gongstruct] = __member
	stage.GongStructs_mapString[gongstruct.Name] = gongstruct

	return gongstruct
}

// Unstage removes gongstruct off the model stage
func (gongstruct *GongStruct) Unstage(stage *StageStruct) *GongStruct {
	delete(stage.GongStructs, gongstruct)
	delete(stage.GongStructs_mapString, gongstruct.Name)
	return gongstruct
}

// commit gongstruct to the back repo (if it is already staged)
func (gongstruct *GongStruct) Commit(stage *StageStruct) *GongStruct {
	if _, ok := stage.GongStructs[gongstruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongStruct(gongstruct)
		}
	}
	return gongstruct
}

func (gongstruct *GongStruct) CommitVoid(stage *StageStruct) {
	gongstruct.Commit(stage)
}

// Checkout gongstruct to the back repo (if it is already staged)
func (gongstruct *GongStruct) Checkout(stage *StageStruct) *GongStruct {
	if _, ok := stage.GongStructs[gongstruct]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongStruct(gongstruct)
		}
	}
	return gongstruct
}

// for satisfaction of GongStruct interface
func (gongstruct *GongStruct) GetName() (res string) {
	return gongstruct.Name
}

// Stage puts gongtimefield to the model stage
func (gongtimefield *GongTimeField) Stage(stage *StageStruct) *GongTimeField {
	stage.GongTimeFields[gongtimefield] = __member
	stage.GongTimeFields_mapString[gongtimefield.Name] = gongtimefield

	return gongtimefield
}

// Unstage removes gongtimefield off the model stage
func (gongtimefield *GongTimeField) Unstage(stage *StageStruct) *GongTimeField {
	delete(stage.GongTimeFields, gongtimefield)
	delete(stage.GongTimeFields_mapString, gongtimefield.Name)
	return gongtimefield
}

// commit gongtimefield to the back repo (if it is already staged)
func (gongtimefield *GongTimeField) Commit(stage *StageStruct) *GongTimeField {
	if _, ok := stage.GongTimeFields[gongtimefield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGongTimeField(gongtimefield)
		}
	}
	return gongtimefield
}

func (gongtimefield *GongTimeField) CommitVoid(stage *StageStruct) {
	gongtimefield.Commit(stage)
}

// Checkout gongtimefield to the back repo (if it is already staged)
func (gongtimefield *GongTimeField) Checkout(stage *StageStruct) *GongTimeField {
	if _, ok := stage.GongTimeFields[gongtimefield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGongTimeField(gongtimefield)
		}
	}
	return gongtimefield
}

// for satisfaction of GongStruct interface
func (gongtimefield *GongTimeField) GetName() (res string) {
	return gongtimefield.Name
}

// Stage puts meta to the model stage
func (meta *Meta) Stage(stage *StageStruct) *Meta {
	stage.Metas[meta] = __member
	stage.Metas_mapString[meta.Name] = meta

	return meta
}

// Unstage removes meta off the model stage
func (meta *Meta) Unstage(stage *StageStruct) *Meta {
	delete(stage.Metas, meta)
	delete(stage.Metas_mapString, meta.Name)
	return meta
}

// commit meta to the back repo (if it is already staged)
func (meta *Meta) Commit(stage *StageStruct) *Meta {
	if _, ok := stage.Metas[meta]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMeta(meta)
		}
	}
	return meta
}

func (meta *Meta) CommitVoid(stage *StageStruct) {
	meta.Commit(stage)
}

// Checkout meta to the back repo (if it is already staged)
func (meta *Meta) Checkout(stage *StageStruct) *Meta {
	if _, ok := stage.Metas[meta]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMeta(meta)
		}
	}
	return meta
}

// for satisfaction of GongStruct interface
func (meta *Meta) GetName() (res string) {
	return meta.Name
}

// Stage puts metareference to the model stage
func (metareference *MetaReference) Stage(stage *StageStruct) *MetaReference {
	stage.MetaReferences[metareference] = __member
	stage.MetaReferences_mapString[metareference.Name] = metareference

	return metareference
}

// Unstage removes metareference off the model stage
func (metareference *MetaReference) Unstage(stage *StageStruct) *MetaReference {
	delete(stage.MetaReferences, metareference)
	delete(stage.MetaReferences_mapString, metareference.Name)
	return metareference
}

// commit metareference to the back repo (if it is already staged)
func (metareference *MetaReference) Commit(stage *StageStruct) *MetaReference {
	if _, ok := stage.MetaReferences[metareference]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMetaReference(metareference)
		}
	}
	return metareference
}

func (metareference *MetaReference) CommitVoid(stage *StageStruct) {
	metareference.Commit(stage)
}

// Checkout metareference to the back repo (if it is already staged)
func (metareference *MetaReference) Checkout(stage *StageStruct) *MetaReference {
	if _, ok := stage.MetaReferences[metareference]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMetaReference(metareference)
		}
	}
	return metareference
}

// for satisfaction of GongStruct interface
func (metareference *MetaReference) GetName() (res string) {
	return metareference.Name
}

// Stage puts modelpkg to the model stage
func (modelpkg *ModelPkg) Stage(stage *StageStruct) *ModelPkg {
	stage.ModelPkgs[modelpkg] = __member
	stage.ModelPkgs_mapString[modelpkg.Name] = modelpkg

	return modelpkg
}

// Unstage removes modelpkg off the model stage
func (modelpkg *ModelPkg) Unstage(stage *StageStruct) *ModelPkg {
	delete(stage.ModelPkgs, modelpkg)
	delete(stage.ModelPkgs_mapString, modelpkg.Name)
	return modelpkg
}

// commit modelpkg to the back repo (if it is already staged)
func (modelpkg *ModelPkg) Commit(stage *StageStruct) *ModelPkg {
	if _, ok := stage.ModelPkgs[modelpkg]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitModelPkg(modelpkg)
		}
	}
	return modelpkg
}

func (modelpkg *ModelPkg) CommitVoid(stage *StageStruct) {
	modelpkg.Commit(stage)
}

// Checkout modelpkg to the back repo (if it is already staged)
func (modelpkg *ModelPkg) Checkout(stage *StageStruct) *ModelPkg {
	if _, ok := stage.ModelPkgs[modelpkg]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutModelPkg(modelpkg)
		}
	}
	return modelpkg
}

// for satisfaction of GongStruct interface
func (modelpkg *ModelPkg) GetName() (res string) {
	return modelpkg.Name
}

// Stage puts pointertogongstructfield to the model stage
func (pointertogongstructfield *PointerToGongStructField) Stage(stage *StageStruct) *PointerToGongStructField {
	stage.PointerToGongStructFields[pointertogongstructfield] = __member
	stage.PointerToGongStructFields_mapString[pointertogongstructfield.Name] = pointertogongstructfield

	return pointertogongstructfield
}

// Unstage removes pointertogongstructfield off the model stage
func (pointertogongstructfield *PointerToGongStructField) Unstage(stage *StageStruct) *PointerToGongStructField {
	delete(stage.PointerToGongStructFields, pointertogongstructfield)
	delete(stage.PointerToGongStructFields_mapString, pointertogongstructfield.Name)
	return pointertogongstructfield
}

// commit pointertogongstructfield to the back repo (if it is already staged)
func (pointertogongstructfield *PointerToGongStructField) Commit(stage *StageStruct) *PointerToGongStructField {
	if _, ok := stage.PointerToGongStructFields[pointertogongstructfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPointerToGongStructField(pointertogongstructfield)
		}
	}
	return pointertogongstructfield
}

func (pointertogongstructfield *PointerToGongStructField) CommitVoid(stage *StageStruct) {
	pointertogongstructfield.Commit(stage)
}

// Checkout pointertogongstructfield to the back repo (if it is already staged)
func (pointertogongstructfield *PointerToGongStructField) Checkout(stage *StageStruct) *PointerToGongStructField {
	if _, ok := stage.PointerToGongStructFields[pointertogongstructfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPointerToGongStructField(pointertogongstructfield)
		}
	}
	return pointertogongstructfield
}

// for satisfaction of GongStruct interface
func (pointertogongstructfield *PointerToGongStructField) GetName() (res string) {
	return pointertogongstructfield.Name
}

// Stage puts sliceofpointertogongstructfield to the model stage
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) Stage(stage *StageStruct) *SliceOfPointerToGongStructField {
	stage.SliceOfPointerToGongStructFields[sliceofpointertogongstructfield] = __member
	stage.SliceOfPointerToGongStructFields_mapString[sliceofpointertogongstructfield.Name] = sliceofpointertogongstructfield

	return sliceofpointertogongstructfield
}

// Unstage removes sliceofpointertogongstructfield off the model stage
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) Unstage(stage *StageStruct) *SliceOfPointerToGongStructField {
	delete(stage.SliceOfPointerToGongStructFields, sliceofpointertogongstructfield)
	delete(stage.SliceOfPointerToGongStructFields_mapString, sliceofpointertogongstructfield.Name)
	return sliceofpointertogongstructfield
}

// commit sliceofpointertogongstructfield to the back repo (if it is already staged)
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) Commit(stage *StageStruct) *SliceOfPointerToGongStructField {
	if _, ok := stage.SliceOfPointerToGongStructFields[sliceofpointertogongstructfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSliceOfPointerToGongStructField(sliceofpointertogongstructfield)
		}
	}
	return sliceofpointertogongstructfield
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) CommitVoid(stage *StageStruct) {
	sliceofpointertogongstructfield.Commit(stage)
}

// Checkout sliceofpointertogongstructfield to the back repo (if it is already staged)
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) Checkout(stage *StageStruct) *SliceOfPointerToGongStructField {
	if _, ok := stage.SliceOfPointerToGongStructFields[sliceofpointertogongstructfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSliceOfPointerToGongStructField(sliceofpointertogongstructfield)
		}
	}
	return sliceofpointertogongstructfield
}

// for satisfaction of GongStruct interface
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GetName() (res string) {
	return sliceofpointertogongstructfield.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMGongBasicField(GongBasicField *GongBasicField)
	CreateORMGongEnum(GongEnum *GongEnum)
	CreateORMGongEnumValue(GongEnumValue *GongEnumValue)
	CreateORMGongLink(GongLink *GongLink)
	CreateORMGongNote(GongNote *GongNote)
	CreateORMGongStruct(GongStruct *GongStruct)
	CreateORMGongTimeField(GongTimeField *GongTimeField)
	CreateORMMeta(Meta *Meta)
	CreateORMMetaReference(MetaReference *MetaReference)
	CreateORMModelPkg(ModelPkg *ModelPkg)
	CreateORMPointerToGongStructField(PointerToGongStructField *PointerToGongStructField)
	CreateORMSliceOfPointerToGongStructField(SliceOfPointerToGongStructField *SliceOfPointerToGongStructField)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMGongBasicField(GongBasicField *GongBasicField)
	DeleteORMGongEnum(GongEnum *GongEnum)
	DeleteORMGongEnumValue(GongEnumValue *GongEnumValue)
	DeleteORMGongLink(GongLink *GongLink)
	DeleteORMGongNote(GongNote *GongNote)
	DeleteORMGongStruct(GongStruct *GongStruct)
	DeleteORMGongTimeField(GongTimeField *GongTimeField)
	DeleteORMMeta(Meta *Meta)
	DeleteORMMetaReference(MetaReference *MetaReference)
	DeleteORMModelPkg(ModelPkg *ModelPkg)
	DeleteORMPointerToGongStructField(PointerToGongStructField *PointerToGongStructField)
	DeleteORMSliceOfPointerToGongStructField(SliceOfPointerToGongStructField *SliceOfPointerToGongStructField)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.GongBasicFields = make(map[*GongBasicField]any)
	stage.GongBasicFields_mapString = make(map[string]*GongBasicField)

	stage.GongEnums = make(map[*GongEnum]any)
	stage.GongEnums_mapString = make(map[string]*GongEnum)

	stage.GongEnumValues = make(map[*GongEnumValue]any)
	stage.GongEnumValues_mapString = make(map[string]*GongEnumValue)

	stage.GongLinks = make(map[*GongLink]any)
	stage.GongLinks_mapString = make(map[string]*GongLink)

	stage.GongNotes = make(map[*GongNote]any)
	stage.GongNotes_mapString = make(map[string]*GongNote)

	stage.GongStructs = make(map[*GongStruct]any)
	stage.GongStructs_mapString = make(map[string]*GongStruct)

	stage.GongTimeFields = make(map[*GongTimeField]any)
	stage.GongTimeFields_mapString = make(map[string]*GongTimeField)

	stage.Metas = make(map[*Meta]any)
	stage.Metas_mapString = make(map[string]*Meta)

	stage.MetaReferences = make(map[*MetaReference]any)
	stage.MetaReferences_mapString = make(map[string]*MetaReference)

	stage.ModelPkgs = make(map[*ModelPkg]any)
	stage.ModelPkgs_mapString = make(map[string]*ModelPkg)

	stage.PointerToGongStructFields = make(map[*PointerToGongStructField]any)
	stage.PointerToGongStructFields_mapString = make(map[string]*PointerToGongStructField)

	stage.SliceOfPointerToGongStructFields = make(map[*SliceOfPointerToGongStructField]any)
	stage.SliceOfPointerToGongStructFields_mapString = make(map[string]*SliceOfPointerToGongStructField)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.GongBasicFields = nil
	stage.GongBasicFields_mapString = nil

	stage.GongEnums = nil
	stage.GongEnums_mapString = nil

	stage.GongEnumValues = nil
	stage.GongEnumValues_mapString = nil

	stage.GongLinks = nil
	stage.GongLinks_mapString = nil

	stage.GongNotes = nil
	stage.GongNotes_mapString = nil

	stage.GongStructs = nil
	stage.GongStructs_mapString = nil

	stage.GongTimeFields = nil
	stage.GongTimeFields_mapString = nil

	stage.Metas = nil
	stage.Metas_mapString = nil

	stage.MetaReferences = nil
	stage.MetaReferences_mapString = nil

	stage.ModelPkgs = nil
	stage.ModelPkgs_mapString = nil

	stage.PointerToGongStructFields = nil
	stage.PointerToGongStructFields_mapString = nil

	stage.SliceOfPointerToGongStructFields = nil
	stage.SliceOfPointerToGongStructFields_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for gongbasicfield := range stage.GongBasicFields {
		gongbasicfield.Unstage(stage)
	}

	for gongenum := range stage.GongEnums {
		gongenum.Unstage(stage)
	}

	for gongenumvalue := range stage.GongEnumValues {
		gongenumvalue.Unstage(stage)
	}

	for gonglink := range stage.GongLinks {
		gonglink.Unstage(stage)
	}

	for gongnote := range stage.GongNotes {
		gongnote.Unstage(stage)
	}

	for gongstruct := range stage.GongStructs {
		gongstruct.Unstage(stage)
	}

	for gongtimefield := range stage.GongTimeFields {
		gongtimefield.Unstage(stage)
	}

	for meta := range stage.Metas {
		meta.Unstage(stage)
	}

	for metareference := range stage.MetaReferences {
		metareference.Unstage(stage)
	}

	for modelpkg := range stage.ModelPkgs {
		modelpkg.Unstage(stage)
	}

	for pointertogongstructfield := range stage.PointerToGongStructFields {
		pointertogongstructfield.Unstage(stage)
	}

	for sliceofpointertogongstructfield := range stage.SliceOfPointerToGongStructFields {
		sliceofpointertogongstructfield.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
	// insertion point for generic types
	GongBasicField | GongEnum | GongEnumValue | GongLink | GongNote | GongStruct | GongTimeField | Meta | MetaReference | ModelPkg | PointerToGongStructField | SliceOfPointerToGongStructField
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
	*GongBasicField | *GongEnum | *GongEnumValue | *GongLink | *GongNote | *GongStruct | *GongTimeField | *Meta | *MetaReference | *ModelPkg | *PointerToGongStructField | *SliceOfPointerToGongStructField
	GetName() string
	CommitVoid(*StageStruct)
}

type GongstructSet interface {
	map[any]any |
		// insertion point for generic types
		map[*GongBasicField]any |
		map[*GongEnum]any |
		map[*GongEnumValue]any |
		map[*GongLink]any |
		map[*GongNote]any |
		map[*GongStruct]any |
		map[*GongTimeField]any |
		map[*Meta]any |
		map[*MetaReference]any |
		map[*ModelPkg]any |
		map[*PointerToGongStructField]any |
		map[*SliceOfPointerToGongStructField]any |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

type GongstructMapString interface {
	map[any]any |
		// insertion point for generic types
		map[string]*GongBasicField |
		map[string]*GongEnum |
		map[string]*GongEnumValue |
		map[string]*GongLink |
		map[string]*GongNote |
		map[string]*GongStruct |
		map[string]*GongTimeField |
		map[string]*Meta |
		map[string]*MetaReference |
		map[string]*ModelPkg |
		map[string]*PointerToGongStructField |
		map[string]*SliceOfPointerToGongStructField |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*GongBasicField]any:
		return any(&stage.GongBasicFields).(*Type)
	case map[*GongEnum]any:
		return any(&stage.GongEnums).(*Type)
	case map[*GongEnumValue]any:
		return any(&stage.GongEnumValues).(*Type)
	case map[*GongLink]any:
		return any(&stage.GongLinks).(*Type)
	case map[*GongNote]any:
		return any(&stage.GongNotes).(*Type)
	case map[*GongStruct]any:
		return any(&stage.GongStructs).(*Type)
	case map[*GongTimeField]any:
		return any(&stage.GongTimeFields).(*Type)
	case map[*Meta]any:
		return any(&stage.Metas).(*Type)
	case map[*MetaReference]any:
		return any(&stage.MetaReferences).(*Type)
	case map[*ModelPkg]any:
		return any(&stage.ModelPkgs).(*Type)
	case map[*PointerToGongStructField]any:
		return any(&stage.PointerToGongStructFields).(*Type)
	case map[*SliceOfPointerToGongStructField]any:
		return any(&stage.SliceOfPointerToGongStructFields).(*Type)
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
	case map[string]*GongBasicField:
		return any(&stage.GongBasicFields_mapString).(*Type)
	case map[string]*GongEnum:
		return any(&stage.GongEnums_mapString).(*Type)
	case map[string]*GongEnumValue:
		return any(&stage.GongEnumValues_mapString).(*Type)
	case map[string]*GongLink:
		return any(&stage.GongLinks_mapString).(*Type)
	case map[string]*GongNote:
		return any(&stage.GongNotes_mapString).(*Type)
	case map[string]*GongStruct:
		return any(&stage.GongStructs_mapString).(*Type)
	case map[string]*GongTimeField:
		return any(&stage.GongTimeFields_mapString).(*Type)
	case map[string]*Meta:
		return any(&stage.Metas_mapString).(*Type)
	case map[string]*MetaReference:
		return any(&stage.MetaReferences_mapString).(*Type)
	case map[string]*ModelPkg:
		return any(&stage.ModelPkgs_mapString).(*Type)
	case map[string]*PointerToGongStructField:
		return any(&stage.PointerToGongStructFields_mapString).(*Type)
	case map[string]*SliceOfPointerToGongStructField:
		return any(&stage.SliceOfPointerToGongStructFields_mapString).(*Type)
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
	case GongBasicField:
		return any(&stage.GongBasicFields).(*map[*Type]any)
	case GongEnum:
		return any(&stage.GongEnums).(*map[*Type]any)
	case GongEnumValue:
		return any(&stage.GongEnumValues).(*map[*Type]any)
	case GongLink:
		return any(&stage.GongLinks).(*map[*Type]any)
	case GongNote:
		return any(&stage.GongNotes).(*map[*Type]any)
	case GongStruct:
		return any(&stage.GongStructs).(*map[*Type]any)
	case GongTimeField:
		return any(&stage.GongTimeFields).(*map[*Type]any)
	case Meta:
		return any(&stage.Metas).(*map[*Type]any)
	case MetaReference:
		return any(&stage.MetaReferences).(*map[*Type]any)
	case ModelPkg:
		return any(&stage.ModelPkgs).(*map[*Type]any)
	case PointerToGongStructField:
		return any(&stage.PointerToGongStructFields).(*map[*Type]any)
	case SliceOfPointerToGongStructField:
		return any(&stage.SliceOfPointerToGongStructFields).(*map[*Type]any)
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
	case *GongBasicField:
		return any(&stage.GongBasicFields).(*map[Type]any)
	case *GongEnum:
		return any(&stage.GongEnums).(*map[Type]any)
	case *GongEnumValue:
		return any(&stage.GongEnumValues).(*map[Type]any)
	case *GongLink:
		return any(&stage.GongLinks).(*map[Type]any)
	case *GongNote:
		return any(&stage.GongNotes).(*map[Type]any)
	case *GongStruct:
		return any(&stage.GongStructs).(*map[Type]any)
	case *GongTimeField:
		return any(&stage.GongTimeFields).(*map[Type]any)
	case *Meta:
		return any(&stage.Metas).(*map[Type]any)
	case *MetaReference:
		return any(&stage.MetaReferences).(*map[Type]any)
	case *ModelPkg:
		return any(&stage.ModelPkgs).(*map[Type]any)
	case *PointerToGongStructField:
		return any(&stage.PointerToGongStructFields).(*map[Type]any)
	case *SliceOfPointerToGongStructField:
		return any(&stage.SliceOfPointerToGongStructFields).(*map[Type]any)
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
	case GongBasicField:
		return any(&stage.GongBasicFields_mapString).(*map[string]*Type)
	case GongEnum:
		return any(&stage.GongEnums_mapString).(*map[string]*Type)
	case GongEnumValue:
		return any(&stage.GongEnumValues_mapString).(*map[string]*Type)
	case GongLink:
		return any(&stage.GongLinks_mapString).(*map[string]*Type)
	case GongNote:
		return any(&stage.GongNotes_mapString).(*map[string]*Type)
	case GongStruct:
		return any(&stage.GongStructs_mapString).(*map[string]*Type)
	case GongTimeField:
		return any(&stage.GongTimeFields_mapString).(*map[string]*Type)
	case Meta:
		return any(&stage.Metas_mapString).(*map[string]*Type)
	case MetaReference:
		return any(&stage.MetaReferences_mapString).(*map[string]*Type)
	case ModelPkg:
		return any(&stage.ModelPkgs_mapString).(*map[string]*Type)
	case PointerToGongStructField:
		return any(&stage.PointerToGongStructFields_mapString).(*map[string]*Type)
	case SliceOfPointerToGongStructField:
		return any(&stage.SliceOfPointerToGongStructFields_mapString).(*map[string]*Type)
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
	case GongBasicField:
		return any(&GongBasicField{
			// Initialisation of associations
			// field is initialized with an instance of GongEnum with the name of the field
			GongEnum: &GongEnum{Name: "GongEnum"},
		}).(*Type)
	case GongEnum:
		return any(&GongEnum{
			// Initialisation of associations
			// field is initialized with an instance of GongEnumValue with the name of the field
			GongEnumValues: []*GongEnumValue{{Name: "GongEnumValues"}},
		}).(*Type)
	case GongEnumValue:
		return any(&GongEnumValue{
			// Initialisation of associations
		}).(*Type)
	case GongLink:
		return any(&GongLink{
			// Initialisation of associations
		}).(*Type)
	case GongNote:
		return any(&GongNote{
			// Initialisation of associations
			// field is initialized with an instance of GongLink with the name of the field
			Links: []*GongLink{{Name: "Links"}},
		}).(*Type)
	case GongStruct:
		return any(&GongStruct{
			// Initialisation of associations
			// field is initialized with an instance of GongBasicField with the name of the field
			GongBasicFields: []*GongBasicField{{Name: "GongBasicFields"}},
			// field is initialized with an instance of GongTimeField with the name of the field
			GongTimeFields: []*GongTimeField{{Name: "GongTimeFields"}},
			// field is initialized with an instance of PointerToGongStructField with the name of the field
			PointerToGongStructFields: []*PointerToGongStructField{{Name: "PointerToGongStructFields"}},
			// field is initialized with an instance of SliceOfPointerToGongStructField with the name of the field
			SliceOfPointerToGongStructFields: []*SliceOfPointerToGongStructField{{Name: "SliceOfPointerToGongStructFields"}},
		}).(*Type)
	case GongTimeField:
		return any(&GongTimeField{
			// Initialisation of associations
		}).(*Type)
	case Meta:
		return any(&Meta{
			// Initialisation of associations
			// field is initialized with an instance of MetaReference with the name of the field
			MetaReferences: []*MetaReference{{Name: "MetaReferences"}},
		}).(*Type)
	case MetaReference:
		return any(&MetaReference{
			// Initialisation of associations
		}).(*Type)
	case ModelPkg:
		return any(&ModelPkg{
			// Initialisation of associations
		}).(*Type)
	case PointerToGongStructField:
		return any(&PointerToGongStructField{
			// Initialisation of associations
			// field is initialized with an instance of GongStruct with the name of the field
			GongStruct: &GongStruct{Name: "GongStruct"},
		}).(*Type)
	case SliceOfPointerToGongStructField:
		return any(&SliceOfPointerToGongStructField{
			// Initialisation of associations
			// field is initialized with an instance of GongStruct with the name of the field
			GongStruct: &GongStruct{Name: "GongStruct"},
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
	// reverse maps of direct associations of GongBasicField
	case GongBasicField:
		switch fieldname {
		// insertion point for per direct association field
		case "GongEnum":
			res := make(map[*GongEnum][]*GongBasicField)
			for gongbasicfield := range stage.GongBasicFields {
				if gongbasicfield.GongEnum != nil {
					gongenum_ := gongbasicfield.GongEnum
					var gongbasicfields []*GongBasicField
					_, ok := res[gongenum_]
					if ok {
						gongbasicfields = res[gongenum_]
					} else {
						gongbasicfields = make([]*GongBasicField, 0)
					}
					gongbasicfields = append(gongbasicfields, gongbasicfield)
					res[gongenum_] = gongbasicfields
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of GongEnum
	case GongEnum:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongEnumValue
	case GongEnumValue:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongLink
	case GongLink:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongNote
	case GongNote:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongStruct
	case GongStruct:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongTimeField
	case GongTimeField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Meta
	case Meta:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MetaReference
	case MetaReference:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ModelPkg
	case ModelPkg:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PointerToGongStructField
	case PointerToGongStructField:
		switch fieldname {
		// insertion point for per direct association field
		case "GongStruct":
			res := make(map[*GongStruct][]*PointerToGongStructField)
			for pointertogongstructfield := range stage.PointerToGongStructFields {
				if pointertogongstructfield.GongStruct != nil {
					gongstruct_ := pointertogongstructfield.GongStruct
					var pointertogongstructfields []*PointerToGongStructField
					_, ok := res[gongstruct_]
					if ok {
						pointertogongstructfields = res[gongstruct_]
					} else {
						pointertogongstructfields = make([]*PointerToGongStructField, 0)
					}
					pointertogongstructfields = append(pointertogongstructfields, pointertogongstructfield)
					res[gongstruct_] = pointertogongstructfields
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SliceOfPointerToGongStructField
	case SliceOfPointerToGongStructField:
		switch fieldname {
		// insertion point for per direct association field
		case "GongStruct":
			res := make(map[*GongStruct][]*SliceOfPointerToGongStructField)
			for sliceofpointertogongstructfield := range stage.SliceOfPointerToGongStructFields {
				if sliceofpointertogongstructfield.GongStruct != nil {
					gongstruct_ := sliceofpointertogongstructfield.GongStruct
					var sliceofpointertogongstructfields []*SliceOfPointerToGongStructField
					_, ok := res[gongstruct_]
					if ok {
						sliceofpointertogongstructfields = res[gongstruct_]
					} else {
						sliceofpointertogongstructfields = make([]*SliceOfPointerToGongStructField, 0)
					}
					sliceofpointertogongstructfields = append(sliceofpointertogongstructfields, sliceofpointertogongstructfield)
					res[gongstruct_] = sliceofpointertogongstructfields
				}
			}
			return any(res).(map[*End][]*Start)
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
	// reverse maps of direct associations of GongBasicField
	case GongBasicField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongEnum
	case GongEnum:
		switch fieldname {
		// insertion point for per direct association field
		case "GongEnumValues":
			res := make(map[*GongEnumValue]*GongEnum)
			for gongenum := range stage.GongEnums {
				for _, gongenumvalue_ := range gongenum.GongEnumValues {
					res[gongenumvalue_] = gongenum
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of GongEnumValue
	case GongEnumValue:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongLink
	case GongLink:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GongNote
	case GongNote:
		switch fieldname {
		// insertion point for per direct association field
		case "Links":
			res := make(map[*GongLink]*GongNote)
			for gongnote := range stage.GongNotes {
				for _, gonglink_ := range gongnote.Links {
					res[gonglink_] = gongnote
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of GongStruct
	case GongStruct:
		switch fieldname {
		// insertion point for per direct association field
		case "GongBasicFields":
			res := make(map[*GongBasicField]*GongStruct)
			for gongstruct := range stage.GongStructs {
				for _, gongbasicfield_ := range gongstruct.GongBasicFields {
					res[gongbasicfield_] = gongstruct
				}
			}
			return any(res).(map[*End]*Start)
		case "GongTimeFields":
			res := make(map[*GongTimeField]*GongStruct)
			for gongstruct := range stage.GongStructs {
				for _, gongtimefield_ := range gongstruct.GongTimeFields {
					res[gongtimefield_] = gongstruct
				}
			}
			return any(res).(map[*End]*Start)
		case "PointerToGongStructFields":
			res := make(map[*PointerToGongStructField]*GongStruct)
			for gongstruct := range stage.GongStructs {
				for _, pointertogongstructfield_ := range gongstruct.PointerToGongStructFields {
					res[pointertogongstructfield_] = gongstruct
				}
			}
			return any(res).(map[*End]*Start)
		case "SliceOfPointerToGongStructFields":
			res := make(map[*SliceOfPointerToGongStructField]*GongStruct)
			for gongstruct := range stage.GongStructs {
				for _, sliceofpointertogongstructfield_ := range gongstruct.SliceOfPointerToGongStructFields {
					res[sliceofpointertogongstructfield_] = gongstruct
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of GongTimeField
	case GongTimeField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Meta
	case Meta:
		switch fieldname {
		// insertion point for per direct association field
		case "MetaReferences":
			res := make(map[*MetaReference]*Meta)
			for meta := range stage.Metas {
				for _, metareference_ := range meta.MetaReferences {
					res[metareference_] = meta
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of MetaReference
	case MetaReference:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ModelPkg
	case ModelPkg:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of PointerToGongStructField
	case PointerToGongStructField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SliceOfPointerToGongStructField
	case SliceOfPointerToGongStructField:
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
	case GongBasicField:
		res = "GongBasicField"
	case GongEnum:
		res = "GongEnum"
	case GongEnumValue:
		res = "GongEnumValue"
	case GongLink:
		res = "GongLink"
	case GongNote:
		res = "GongNote"
	case GongStruct:
		res = "GongStruct"
	case GongTimeField:
		res = "GongTimeField"
	case Meta:
		res = "Meta"
	case MetaReference:
		res = "MetaReference"
	case ModelPkg:
		res = "ModelPkg"
	case PointerToGongStructField:
		res = "PointerToGongStructField"
	case SliceOfPointerToGongStructField:
		res = "SliceOfPointerToGongStructField"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case GongBasicField:
		res = []string{"Name", "BasicKindName", "GongEnum", "DeclaredType", "CompositeStructName", "Index", "IsDocLink"}
	case GongEnum:
		res = []string{"Name", "Type", "GongEnumValues"}
	case GongEnumValue:
		res = []string{"Name", "Value"}
	case GongLink:
		res = []string{"Name", "Recv", "ImportPath"}
	case GongNote:
		res = []string{"Name", "Body", "BodyHTML", "Links"}
	case GongStruct:
		res = []string{"Name", "GongBasicFields", "GongTimeFields", "PointerToGongStructFields", "SliceOfPointerToGongStructFields", "HasOnAfterUpdateSignature"}
	case GongTimeField:
		res = []string{"Name", "Index", "CompositeStructName"}
	case Meta:
		res = []string{"Name", "Text", "MetaReferences"}
	case MetaReference:
		res = []string{"Name"}
	case ModelPkg:
		res = []string{"Name", "PkgPath"}
	case PointerToGongStructField:
		res = []string{"Name", "GongStruct", "Index", "CompositeStructName"}
	case SliceOfPointerToGongStructField:
		res = []string{"Name", "GongStruct", "Index", "CompositeStructName"}
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *GongBasicField:
		res = []string{"Name", "BasicKindName", "GongEnum", "DeclaredType", "CompositeStructName", "Index", "IsDocLink"}
	case *GongEnum:
		res = []string{"Name", "Type", "GongEnumValues"}
	case *GongEnumValue:
		res = []string{"Name", "Value"}
	case *GongLink:
		res = []string{"Name", "Recv", "ImportPath"}
	case *GongNote:
		res = []string{"Name", "Body", "BodyHTML", "Links"}
	case *GongStruct:
		res = []string{"Name", "GongBasicFields", "GongTimeFields", "PointerToGongStructFields", "SliceOfPointerToGongStructFields", "HasOnAfterUpdateSignature"}
	case *GongTimeField:
		res = []string{"Name", "Index", "CompositeStructName"}
	case *Meta:
		res = []string{"Name", "Text", "MetaReferences"}
	case *MetaReference:
		res = []string{"Name"}
	case *ModelPkg:
		res = []string{"Name", "PkgPath"}
	case *PointerToGongStructField:
		res = []string{"Name", "GongStruct", "Index", "CompositeStructName"}
	case *SliceOfPointerToGongStructField:
		res = []string{"Name", "GongStruct", "Index", "CompositeStructName"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *GongBasicField:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "BasicKindName":
			res = inferedInstance.BasicKindName
		case "GongEnum":
			if inferedInstance.GongEnum != nil {
				res = inferedInstance.GongEnum.Name
			}
		case "DeclaredType":
			res = inferedInstance.DeclaredType
		case "CompositeStructName":
			res = inferedInstance.CompositeStructName
		case "Index":
			res = fmt.Sprintf("%d", inferedInstance.Index)
		case "IsDocLink":
			res = fmt.Sprintf("%t", inferedInstance.IsDocLink)
		}
	case *GongEnum:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Type":
			enum := inferedInstance.Type
			res = enum.ToCodeString()
		case "GongEnumValues":
			for idx, __instance__ := range inferedInstance.GongEnumValues {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *GongEnumValue:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value
		}
	case *GongLink:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Recv":
			res = inferedInstance.Recv
		case "ImportPath":
			res = inferedInstance.ImportPath
		}
	case *GongNote:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Body":
			res = inferedInstance.Body
		case "BodyHTML":
			res = inferedInstance.BodyHTML
		case "Links":
			for idx, __instance__ := range inferedInstance.Links {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *GongStruct:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "GongBasicFields":
			for idx, __instance__ := range inferedInstance.GongBasicFields {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "GongTimeFields":
			for idx, __instance__ := range inferedInstance.GongTimeFields {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "PointerToGongStructFields":
			for idx, __instance__ := range inferedInstance.PointerToGongStructFields {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "SliceOfPointerToGongStructFields":
			for idx, __instance__ := range inferedInstance.SliceOfPointerToGongStructFields {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "HasOnAfterUpdateSignature":
			res = fmt.Sprintf("%t", inferedInstance.HasOnAfterUpdateSignature)
		}
	case *GongTimeField:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Index":
			res = fmt.Sprintf("%d", inferedInstance.Index)
		case "CompositeStructName":
			res = inferedInstance.CompositeStructName
		}
	case *Meta:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Text":
			res = inferedInstance.Text
		case "MetaReferences":
			for idx, __instance__ := range inferedInstance.MetaReferences {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *MetaReference:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *ModelPkg:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "PkgPath":
			res = inferedInstance.PkgPath
		}
	case *PointerToGongStructField:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "GongStruct":
			if inferedInstance.GongStruct != nil {
				res = inferedInstance.GongStruct.Name
			}
		case "Index":
			res = fmt.Sprintf("%d", inferedInstance.Index)
		case "CompositeStructName":
			res = inferedInstance.CompositeStructName
		}
	case *SliceOfPointerToGongStructField:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "GongStruct":
			if inferedInstance.GongStruct != nil {
				res = inferedInstance.GongStruct.Name
			}
		case "Index":
			res = fmt.Sprintf("%d", inferedInstance.Index)
		case "CompositeStructName":
			res = inferedInstance.CompositeStructName
		}
	default:
		_ = inferedInstance	
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case GongBasicField:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "BasicKindName":
			res = inferedInstance.BasicKindName
		case "GongEnum":
			if inferedInstance.GongEnum != nil {
				res = inferedInstance.GongEnum.Name
			}
		case "DeclaredType":
			res = inferedInstance.DeclaredType
		case "CompositeStructName":
			res = inferedInstance.CompositeStructName
		case "Index":
			res = fmt.Sprintf("%d", inferedInstance.Index)
		case "IsDocLink":
			res = fmt.Sprintf("%t", inferedInstance.IsDocLink)
		}
	case GongEnum:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Type":
			enum := inferedInstance.Type
			res = enum.ToCodeString()
		case "GongEnumValues":
			for idx, __instance__ := range inferedInstance.GongEnumValues {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case GongEnumValue:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value
		}
	case GongLink:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Recv":
			res = inferedInstance.Recv
		case "ImportPath":
			res = inferedInstance.ImportPath
		}
	case GongNote:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Body":
			res = inferedInstance.Body
		case "BodyHTML":
			res = inferedInstance.BodyHTML
		case "Links":
			for idx, __instance__ := range inferedInstance.Links {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case GongStruct:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "GongBasicFields":
			for idx, __instance__ := range inferedInstance.GongBasicFields {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "GongTimeFields":
			for idx, __instance__ := range inferedInstance.GongTimeFields {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "PointerToGongStructFields":
			for idx, __instance__ := range inferedInstance.PointerToGongStructFields {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "SliceOfPointerToGongStructFields":
			for idx, __instance__ := range inferedInstance.SliceOfPointerToGongStructFields {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "HasOnAfterUpdateSignature":
			res = fmt.Sprintf("%t", inferedInstance.HasOnAfterUpdateSignature)
		}
	case GongTimeField:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Index":
			res = fmt.Sprintf("%d", inferedInstance.Index)
		case "CompositeStructName":
			res = inferedInstance.CompositeStructName
		}
	case Meta:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Text":
			res = inferedInstance.Text
		case "MetaReferences":
			for idx, __instance__ := range inferedInstance.MetaReferences {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case MetaReference:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case ModelPkg:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "PkgPath":
			res = inferedInstance.PkgPath
		}
	case PointerToGongStructField:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "GongStruct":
			if inferedInstance.GongStruct != nil {
				res = inferedInstance.GongStruct.Name
			}
		case "Index":
			res = fmt.Sprintf("%d", inferedInstance.Index)
		case "CompositeStructName":
			res = inferedInstance.CompositeStructName
		}
	case SliceOfPointerToGongStructField:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "GongStruct":
			if inferedInstance.GongStruct != nil {
				res = inferedInstance.GongStruct.Name
			}
		case "Index":
			res = fmt.Sprintf("%d", inferedInstance.Index)
		case "CompositeStructName":
			res = inferedInstance.CompositeStructName
		}
	default:
		_ = inferedInstance	
	}
	return
}

// Last line of the template
