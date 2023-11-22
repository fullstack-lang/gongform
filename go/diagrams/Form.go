package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongtable/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Form models.StageStruct
var ___dummy__Time_Form time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Form ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Form map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.Cell": &(ref_models.Cell{}),

	"ref_models.Cell.CellBool": (ref_models.Cell{}).CellBool,

	"ref_models.Cell.CellFloat64": (ref_models.Cell{}).CellFloat64,

	"ref_models.Cell.CellIcon": (ref_models.Cell{}).CellIcon,

	"ref_models.Cell.CellInt": (ref_models.Cell{}).CellInt,

	"ref_models.Cell.CellString": (ref_models.Cell{}).CellString,

	"ref_models.Cell.Name": (ref_models.Cell{}).Name,

	"ref_models.CellBoolean": &(ref_models.CellBoolean{}),

	"ref_models.CellBoolean.Name": (ref_models.CellBoolean{}).Name,

	"ref_models.CellBoolean.Value": (ref_models.CellBoolean{}).Value,

	"ref_models.CellFloat64": &(ref_models.CellFloat64{}),

	"ref_models.CellFloat64.Name": (ref_models.CellFloat64{}).Name,

	"ref_models.CellFloat64.Value": (ref_models.CellFloat64{}).Value,

	"ref_models.CellIcon": &(ref_models.CellIcon{}),

	"ref_models.CellIcon.Icon": (ref_models.CellIcon{}).Icon,

	"ref_models.CellIcon.Name": (ref_models.CellIcon{}).Name,

	"ref_models.CellInt": &(ref_models.CellInt{}),

	"ref_models.CellInt.Name": (ref_models.CellInt{}).Name,

	"ref_models.CellInt.Value": (ref_models.CellInt{}).Value,

	"ref_models.CellString": &(ref_models.CellString{}),

	"ref_models.CellString.Name": (ref_models.CellString{}).Name,

	"ref_models.CellString.Value": (ref_models.CellString{}).Value,

	"ref_models.CheckBox": &(ref_models.CheckBox{}),

	"ref_models.CheckBox.Name": (ref_models.CheckBox{}).Name,

	"ref_models.CheckBox.Value": (ref_models.CheckBox{}).Value,

	"ref_models.Color": ref_models.Color,

	"ref_models.Date": ref_models.Date,

	"ref_models.Datetime": ref_models.Datetime,

	"ref_models.DisplayedColumn": &(ref_models.DisplayedColumn{}),

	"ref_models.DisplayedColumn.Name": (ref_models.DisplayedColumn{}).Name,

	"ref_models.Email": ref_models.Email,

	"ref_models.File": ref_models.File,

	"ref_models.FormDiv": &(ref_models.FormDiv{}),

	"ref_models.FormDiv.CheckBoxs": (ref_models.FormDiv{}).CheckBoxs,

	"ref_models.FormDiv.FormEditAssocButton": (ref_models.FormDiv{}).FormEditAssocButton,

	"ref_models.FormDiv.FormFields": (ref_models.FormDiv{}).FormFields,

	"ref_models.FormDiv.Name": (ref_models.FormDiv{}).Name,

	"ref_models.FormEditAssocButton": &(ref_models.FormEditAssocButton{}),

	"ref_models.FormEditAssocButton.Label": (ref_models.FormEditAssocButton{}).Label,

	"ref_models.FormEditAssocButton.Name": (ref_models.FormEditAssocButton{}).Name,

	"ref_models.FormField": &(ref_models.FormField{}),

	"ref_models.FormField.FormFieldDate": (ref_models.FormField{}).FormFieldDate,

	"ref_models.FormField.FormFieldDateTime": (ref_models.FormField{}).FormFieldDateTime,

	"ref_models.FormField.FormFieldFloat64": (ref_models.FormField{}).FormFieldFloat64,

	"ref_models.FormField.FormFieldInt": (ref_models.FormField{}).FormFieldInt,

	"ref_models.FormField.FormFieldSelect": (ref_models.FormField{}).FormFieldSelect,

	"ref_models.FormField.FormFieldString": (ref_models.FormField{}).FormFieldString,

	"ref_models.FormField.FormFieldTime": (ref_models.FormField{}).FormFieldTime,

	"ref_models.FormField.InputTypeEnum": (ref_models.FormField{}).InputTypeEnum,

	"ref_models.FormField.Label": (ref_models.FormField{}).Label,

	"ref_models.FormField.Name": (ref_models.FormField{}).Name,

	"ref_models.FormField.Placeholder": (ref_models.FormField{}).Placeholder,

	"ref_models.FormFieldDate": &(ref_models.FormFieldDate{}),

	"ref_models.FormFieldDate.Name": (ref_models.FormFieldDate{}).Name,

	"ref_models.FormFieldDate.Value": (ref_models.FormFieldDate{}).Value,

	"ref_models.FormFieldDateTime": &(ref_models.FormFieldDateTime{}),

	"ref_models.FormFieldDateTime.Name": (ref_models.FormFieldDateTime{}).Name,

	"ref_models.FormFieldDateTime.Value": (ref_models.FormFieldDateTime{}).Value,

	"ref_models.FormFieldFloat64": &(ref_models.FormFieldFloat64{}),

	"ref_models.FormFieldFloat64.Name": (ref_models.FormFieldFloat64{}).Name,

	"ref_models.FormFieldFloat64.Value": (ref_models.FormFieldFloat64{}).Value,

	"ref_models.FormFieldInt": &(ref_models.FormFieldInt{}),

	"ref_models.FormFieldInt.Name": (ref_models.FormFieldInt{}).Name,

	"ref_models.FormFieldInt.Value": (ref_models.FormFieldInt{}).Value,

	"ref_models.FormFieldSelect": &(ref_models.FormFieldSelect{}),

	"ref_models.FormFieldSelect.Name": (ref_models.FormFieldSelect{}).Name,

	"ref_models.FormFieldSelect.Options": (ref_models.FormFieldSelect{}).Options,

	"ref_models.FormFieldSelect.Value": (ref_models.FormFieldSelect{}).Value,

	"ref_models.FormFieldString": &(ref_models.FormFieldString{}),

	"ref_models.FormFieldString.Name": (ref_models.FormFieldString{}).Name,

	"ref_models.FormFieldString.Value": (ref_models.FormFieldString{}).Value,

	"ref_models.FormFieldTime": &(ref_models.FormFieldTime{}),

	"ref_models.FormFieldTime.Name": (ref_models.FormFieldTime{}).Name,

	"ref_models.FormFieldTime.Step": (ref_models.FormFieldTime{}).Step,

	"ref_models.FormFieldTime.Value": (ref_models.FormFieldTime{}).Value,

	"ref_models.FormGroup": &(ref_models.FormGroup{}),

	"ref_models.FormGroup.FormDivs": (ref_models.FormGroup{}).FormDivs,

	"ref_models.FormGroup.Name": (ref_models.FormGroup{}).Name,

	"ref_models.Hidden": ref_models.Hidden,

	"ref_models.InputTypeEnum": ref_models.InputTypeEnum(""),

	"ref_models.Month": ref_models.Month,

	"ref_models.Number": ref_models.Number,

	"ref_models.Option": &(ref_models.Option{}),

	"ref_models.Option.Name": (ref_models.Option{}).Name,

	"ref_models.Password": ref_models.Password,

	"ref_models.Range": ref_models.Range,

	"ref_models.Row": &(ref_models.Row{}),

	"ref_models.Row.Cells": (ref_models.Row{}).Cells,

	"ref_models.Row.IsChecked": (ref_models.Row{}).IsChecked,

	"ref_models.Row.Name": (ref_models.Row{}).Name,

	"ref_models.Search": ref_models.Search,

	"ref_models.Table": &(ref_models.Table{}),

	"ref_models.Table.CanDragDropRows": (ref_models.Table{}).CanDragDropRows,

	"ref_models.Table.DisplayedColumns": (ref_models.Table{}).DisplayedColumns,

	"ref_models.Table.HasCheckableRows": (ref_models.Table{}).HasCheckableRows,

	"ref_models.Table.HasColumnSorting": (ref_models.Table{}).HasColumnSorting,

	"ref_models.Table.HasFiltering": (ref_models.Table{}).HasFiltering,

	"ref_models.Table.HasPaginator": (ref_models.Table{}).HasPaginator,

	"ref_models.Table.HasSaveButton": (ref_models.Table{}).HasSaveButton,

	"ref_models.Table.Name": (ref_models.Table{}).Name,

	"ref_models.Table.Rows": (ref_models.Table{}).Rows,

	"ref_models.Table.SavingInProgress": (ref_models.Table{}).SavingInProgress,

	"ref_models.TableExtraNameEnum": ref_models.TableExtraNameEnum(""),

	"ref_models.TableExtraPathEnum": ref_models.TableExtraPathEnum(""),

	"ref_models.TableSelectExtra": ref_models.StackNamePostFixForTableForAssociation,

	"ref_models.TableSelectExtraName": ref_models.TableSelectExtraName,

	"ref_models.Tel": ref_models.Tel,

	"ref_models.Text": ref_models.Text,

	"ref_models.Time": ref_models.Time,

	"ref_models.URL": ref_models.URL,

	"ref_models.Week": ref_models.Week,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Form"] = FormInjection
// }

// FormInjection will stage objects of database "Form"
func FormInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Form := (&models.Classdiagram{Name: `Form`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_InputTypeEnum := (&models.Field{Name: `InputTypeEnum`}).Stage(stage)
	__Field__000001_Label := (&models.Field{Name: `Label`}).Stage(stage)
	__Field__000002_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000003_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000004_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000005_Placeholder := (&models.Field{Name: `Placeholder`}).Stage(stage)
	__Field__000006_Step := (&models.Field{Name: `Step`}).Stage(stage)
	__Field__000007_Value := (&models.Field{Name: `Value`}).Stage(stage)
	__Field__000008_Value := (&models.Field{Name: `Value`}).Stage(stage)
	__Field__000009_Value := (&models.Field{Name: `Value`}).Stage(stage)
	__Field__000010_Value := (&models.Field{Name: `Value`}).Stage(stage)
	__Field__000011_Value := (&models.Field{Name: `Value`}).Stage(stage)
	__Field__000012_Value := (&models.Field{Name: `Value`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape
	__GongEnumShape__000000_Form_InputTypeEnum := (&models.GongEnumShape{Name: `Form-InputTypeEnum`}).Stage(stage)

	// Declarations of staged instances of GongEnumValueEntry
	__GongEnumValueEntry__000000_Color := (&models.GongEnumValueEntry{Name: `Color`}).Stage(stage)
	__GongEnumValueEntry__000001_Date := (&models.GongEnumValueEntry{Name: `Date`}).Stage(stage)
	__GongEnumValueEntry__000002_Datetime := (&models.GongEnumValueEntry{Name: `Datetime`}).Stage(stage)
	__GongEnumValueEntry__000003_Email := (&models.GongEnumValueEntry{Name: `Email`}).Stage(stage)
	__GongEnumValueEntry__000004_File := (&models.GongEnumValueEntry{Name: `File`}).Stage(stage)
	__GongEnumValueEntry__000005_Hidden := (&models.GongEnumValueEntry{Name: `Hidden`}).Stage(stage)
	__GongEnumValueEntry__000006_Month := (&models.GongEnumValueEntry{Name: `Month`}).Stage(stage)
	__GongEnumValueEntry__000007_Number := (&models.GongEnumValueEntry{Name: `Number`}).Stage(stage)
	__GongEnumValueEntry__000008_Password := (&models.GongEnumValueEntry{Name: `Password`}).Stage(stage)
	__GongEnumValueEntry__000009_Range := (&models.GongEnumValueEntry{Name: `Range`}).Stage(stage)
	__GongEnumValueEntry__000010_Search := (&models.GongEnumValueEntry{Name: `Search`}).Stage(stage)
	__GongEnumValueEntry__000011_Tel := (&models.GongEnumValueEntry{Name: `Tel`}).Stage(stage)
	__GongEnumValueEntry__000012_Text := (&models.GongEnumValueEntry{Name: `Text`}).Stage(stage)
	__GongEnumValueEntry__000013_Time := (&models.GongEnumValueEntry{Name: `Time`}).Stage(stage)
	__GongEnumValueEntry__000014_URL := (&models.GongEnumValueEntry{Name: `URL`}).Stage(stage)
	__GongEnumValueEntry__000015_Week := (&models.GongEnumValueEntry{Name: `Week`}).Stage(stage)

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_Form_CheckBox := (&models.GongStructShape{Name: `Form-CheckBox`}).Stage(stage)
	__GongStructShape__000001_Form_FormCell := (&models.GongStructShape{Name: `Form-FormCell`}).Stage(stage)
	__GongStructShape__000002_Form_FormCellFloat64 := (&models.GongStructShape{Name: `Form-FormCellFloat64`}).Stage(stage)
	__GongStructShape__000003_Form_FormCellInt := (&models.GongStructShape{Name: `Form-FormCellInt`}).Stage(stage)
	__GongStructShape__000004_Form_FormCellString := (&models.GongStructShape{Name: `Form-FormCellString`}).Stage(stage)
	__GongStructShape__000005_Form_FormDiv := (&models.GongStructShape{Name: `Form-FormDiv`}).Stage(stage)
	__GongStructShape__000006_Form_FormEditAssocButton := (&models.GongStructShape{Name: `Form-FormEditAssocButton`}).Stage(stage)
	__GongStructShape__000007_Form_FormFieldDate := (&models.GongStructShape{Name: `Form-FormFieldDate`}).Stage(stage)
	__GongStructShape__000008_Form_FormFieldDateTime := (&models.GongStructShape{Name: `Form-FormFieldDateTime`}).Stage(stage)
	__GongStructShape__000009_Form_FormFieldSelect := (&models.GongStructShape{Name: `Form-FormFieldSelect`}).Stage(stage)
	__GongStructShape__000010_Form_FormFieldTime := (&models.GongStructShape{Name: `Form-FormFieldTime`}).Stage(stage)
	__GongStructShape__000011_Form_FormGroup := (&models.GongStructShape{Name: `Form-FormGroup`}).Stage(stage)
	__GongStructShape__000012_Form_Option := (&models.GongStructShape{Name: `Form-Option`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_CheckBoxs := (&models.Link{Name: `CheckBoxs`}).Stage(stage)
	__Link__000001_FormCellFloat64 := (&models.Link{Name: `FormCellFloat64`}).Stage(stage)
	__Link__000002_FormCellInt := (&models.Link{Name: `FormCellInt`}).Stage(stage)
	__Link__000003_FormCellString := (&models.Link{Name: `FormCellString`}).Stage(stage)
	__Link__000004_FormDivs := (&models.Link{Name: `FormDivs`}).Stage(stage)
	__Link__000005_FormEditAssocButton := (&models.Link{Name: `FormEditAssocButton`}).Stage(stage)
	__Link__000006_FormFieldDate := (&models.Link{Name: `FormFieldDate`}).Stage(stage)
	__Link__000007_FormFieldDateTime := (&models.Link{Name: `FormFieldDateTime`}).Stage(stage)
	__Link__000008_FormFieldSelect := (&models.Link{Name: `FormFieldSelect`}).Stage(stage)
	__Link__000009_FormFieldTime := (&models.Link{Name: `FormFieldTime`}).Stage(stage)
	__Link__000010_FormFields := (&models.Link{Name: `FormFields`}).Stage(stage)
	__Link__000011_Options := (&models.Link{Name: `Options`}).Stage(stage)
	__Link__000012_Value := (&models.Link{Name: `Value`}).Stage(stage)

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_Form_CheckBox := (&models.Position{Name: `Pos-Form-CheckBox`}).Stage(stage)
	__Position__000001_Pos_Form_FormCell := (&models.Position{Name: `Pos-Form-FormCell`}).Stage(stage)
	__Position__000002_Pos_Form_FormCellFloat64 := (&models.Position{Name: `Pos-Form-FormCellFloat64`}).Stage(stage)
	__Position__000003_Pos_Form_FormCellInt := (&models.Position{Name: `Pos-Form-FormCellInt`}).Stage(stage)
	__Position__000004_Pos_Form_FormCellString := (&models.Position{Name: `Pos-Form-FormCellString`}).Stage(stage)
	__Position__000005_Pos_Form_FormDiv := (&models.Position{Name: `Pos-Form-FormDiv`}).Stage(stage)
	__Position__000006_Pos_Form_FormEditAssocButton := (&models.Position{Name: `Pos-Form-FormEditAssocButton`}).Stage(stage)
	__Position__000007_Pos_Form_FormFieldDate := (&models.Position{Name: `Pos-Form-FormFieldDate`}).Stage(stage)
	__Position__000008_Pos_Form_FormFieldDateTime := (&models.Position{Name: `Pos-Form-FormFieldDateTime`}).Stage(stage)
	__Position__000009_Pos_Form_FormFieldSelect := (&models.Position{Name: `Pos-Form-FormFieldSelect`}).Stage(stage)
	__Position__000010_Pos_Form_FormFieldTime := (&models.Position{Name: `Pos-Form-FormFieldTime`}).Stage(stage)
	__Position__000011_Pos_Form_FormGroup := (&models.Position{Name: `Pos-Form-FormGroup`}).Stage(stage)
	__Position__000012_Pos_Form_InputTypeEnum := (&models.Position{Name: `Pos-Form-InputTypeEnum`}).Stage(stage)
	__Position__000013_Pos_Form_Option := (&models.Position{Name: `Pos-Form-Option`}).Stage(stage)

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellFloat64 := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellFloat64`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellInt := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellInt`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellString := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellString`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldDate := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormCell and Form-FormFieldDate`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldDateTime := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormCell and Form-FormFieldDateTime`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldSelect := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormCell and Form-FormFieldSelect`}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldTime := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormCell and Form-FormFieldTime`}).Stage(stage)
	__Vertice__000007_Verticle_in_class_diagram_Form_in_middle_between_Form_FormDiv_and_Form_CheckBox := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormDiv and Form-CheckBox`}).Stage(stage)
	__Vertice__000008_Verticle_in_class_diagram_Form_in_middle_between_Form_FormDiv_and_Form_FormCell := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormDiv and Form-FormCell`}).Stage(stage)
	__Vertice__000009_Verticle_in_class_diagram_Form_in_middle_between_Form_FormDiv_and_Form_FormEditAssocButton := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormDiv and Form-FormEditAssocButton`}).Stage(stage)
	__Vertice__000010_Verticle_in_class_diagram_Form_in_middle_between_Form_FormFieldSelect_and_Form_Option := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormFieldSelect and Form-Option`}).Stage(stage)
	__Vertice__000011_Verticle_in_class_diagram_Form_in_middle_between_Form_FormFieldSelect_and_Form_Option := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormFieldSelect and Form-Option`}).Stage(stage)
	__Vertice__000012_Verticle_in_class_diagram_Form_in_middle_between_Form_FormGroup_and_Form_FormDiv := (&models.Vertice{Name: `Verticle in class diagram Form in middle between Form-FormGroup and Form-FormDiv`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Form.Name = `Form`
	__Classdiagram__000000_Form.IsInDrawMode = true

	// Field values setup
	__Field__000000_InputTypeEnum.Name = `InputTypeEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.InputTypeEnum]
	__Field__000000_InputTypeEnum.Identifier = `ref_models.FormField.InputTypeEnum`
	__Field__000000_InputTypeEnum.FieldTypeAsString = ``
	__Field__000000_InputTypeEnum.Structname = `FormField`
	__Field__000000_InputTypeEnum.Fieldtypename = `InputTypeEnum`

	// Field values setup
	__Field__000001_Label.Name = `Label`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.Label]
	__Field__000001_Label.Identifier = `ref_models.FormField.Label`
	__Field__000001_Label.FieldTypeAsString = ``
	__Field__000001_Label.Structname = `FormField`
	__Field__000001_Label.Fieldtypename = `string`

	// Field values setup
	__Field__000002_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.Name]
	__Field__000002_Name.Identifier = `ref_models.FormField.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `FormField`
	__Field__000002_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000003_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldSelect.Name]
	__Field__000003_Name.Identifier = `ref_models.FormFieldSelect.Name`
	__Field__000003_Name.FieldTypeAsString = ``
	__Field__000003_Name.Structname = `FormFieldSelect`
	__Field__000003_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000004_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormDiv.Name]
	__Field__000004_Name.Identifier = `ref_models.FormDiv.Name`
	__Field__000004_Name.FieldTypeAsString = ``
	__Field__000004_Name.Structname = `FormDiv`
	__Field__000004_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000005_Placeholder.Name = `Placeholder`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.Placeholder]
	__Field__000005_Placeholder.Identifier = `ref_models.FormField.Placeholder`
	__Field__000005_Placeholder.FieldTypeAsString = ``
	__Field__000005_Placeholder.Structname = `FormField`
	__Field__000005_Placeholder.Fieldtypename = `string`

	// Field values setup
	__Field__000006_Step.Name = `Step`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldTime.Step]
	__Field__000006_Step.Identifier = `ref_models.FormFieldTime.Step`
	__Field__000006_Step.FieldTypeAsString = ``
	__Field__000006_Step.Structname = `FormFieldTime`
	__Field__000006_Step.Fieldtypename = `float64`

	// Field values setup
	__Field__000007_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldTime.Value]
	__Field__000007_Value.Identifier = `ref_models.FormFieldTime.Value`
	__Field__000007_Value.FieldTypeAsString = ``
	__Field__000007_Value.Structname = `FormFieldTime`
	__Field__000007_Value.Fieldtypename = `Time`

	// Field values setup
	__Field__000008_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldDate.Value]
	__Field__000008_Value.Identifier = `ref_models.FormFieldDate.Value`
	__Field__000008_Value.FieldTypeAsString = ``
	__Field__000008_Value.Structname = `FormFieldDate`
	__Field__000008_Value.Fieldtypename = `Time`

	// Field values setup
	__Field__000009_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CheckBox.Value]
	__Field__000009_Value.Identifier = `ref_models.CheckBox.Value`
	__Field__000009_Value.FieldTypeAsString = ``
	__Field__000009_Value.Structname = `CheckBox`
	__Field__000009_Value.Fieldtypename = `bool`

	// Field values setup
	__Field__000010_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldInt.Value]
	__Field__000010_Value.Identifier = `ref_models.FormFieldInt.Value`
	__Field__000010_Value.FieldTypeAsString = ``
	__Field__000010_Value.Structname = `FormCellInt`
	__Field__000010_Value.Fieldtypename = `int`

	// Field values setup
	__Field__000011_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldFloat64.Value]
	__Field__000011_Value.Identifier = `ref_models.FormFieldFloat64.Value`
	__Field__000011_Value.FieldTypeAsString = ``
	__Field__000011_Value.Structname = `FormCellFloat64`
	__Field__000011_Value.Fieldtypename = `float64`

	// Field values setup
	__Field__000012_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldString.Value]
	__Field__000012_Value.Identifier = `ref_models.FormFieldString.Value`
	__Field__000012_Value.FieldTypeAsString = ``
	__Field__000012_Value.Structname = `FormCellString`
	__Field__000012_Value.Fieldtypename = `string`

	// GongEnumShape values setup
	__GongEnumShape__000000_Form_InputTypeEnum.Name = `Form-InputTypeEnum`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum]
	__GongEnumShape__000000_Form_InputTypeEnum.Identifier = `ref_models.InputTypeEnum`
	__GongEnumShape__000000_Form_InputTypeEnum.Width = 240.000000
	__GongEnumShape__000000_Form_InputTypeEnum.Height = 303.000000

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000000_Color.Name = `Color`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Color]
	__GongEnumValueEntry__000000_Color.Identifier = `ref_models.InputTypeEnum.Color`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000001_Date.Name = `Date`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Date]
	__GongEnumValueEntry__000001_Date.Identifier = `ref_models.InputTypeEnum.Date`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000002_Datetime.Name = `Datetime`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Datetime]
	__GongEnumValueEntry__000002_Datetime.Identifier = `ref_models.InputTypeEnum.Datetime`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000003_Email.Name = `Email`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Email]
	__GongEnumValueEntry__000003_Email.Identifier = `ref_models.InputTypeEnum.Email`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000004_File.Name = `File`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.File]
	__GongEnumValueEntry__000004_File.Identifier = `ref_models.InputTypeEnum.File`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000005_Hidden.Name = `Hidden`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Hidden]
	__GongEnumValueEntry__000005_Hidden.Identifier = `ref_models.InputTypeEnum.Hidden`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000006_Month.Name = `Month`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Month]
	__GongEnumValueEntry__000006_Month.Identifier = `ref_models.InputTypeEnum.Month`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000007_Number.Name = `Number`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Number]
	__GongEnumValueEntry__000007_Number.Identifier = `ref_models.InputTypeEnum.Number`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000008_Password.Name = `Password`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Password]
	__GongEnumValueEntry__000008_Password.Identifier = `ref_models.InputTypeEnum.Password`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000009_Range.Name = `Range`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Range]
	__GongEnumValueEntry__000009_Range.Identifier = `ref_models.InputTypeEnum.Range`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000010_Search.Name = `Search`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Search]
	__GongEnumValueEntry__000010_Search.Identifier = `ref_models.InputTypeEnum.Search`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000011_Tel.Name = `Tel`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Tel]
	__GongEnumValueEntry__000011_Tel.Identifier = `ref_models.InputTypeEnum.Tel`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000012_Text.Name = `Text`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Text]
	__GongEnumValueEntry__000012_Text.Identifier = `ref_models.InputTypeEnum.Text`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000013_Time.Name = `Time`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Time]
	__GongEnumValueEntry__000013_Time.Identifier = `ref_models.InputTypeEnum.Time`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000014_URL.Name = `URL`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.URL]
	__GongEnumValueEntry__000014_URL.Identifier = `ref_models.InputTypeEnum.URL`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000015_Week.Name = `Week`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.InputTypeEnum.Week]
	__GongEnumValueEntry__000015_Week.Identifier = `ref_models.InputTypeEnum.Week`

	// GongStructShape values setup
	__GongStructShape__000000_Form_CheckBox.Name = `Form-CheckBox`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CheckBox]
	__GongStructShape__000000_Form_CheckBox.Identifier = `ref_models.CheckBox`
	__GongStructShape__000000_Form_CheckBox.ShowNbInstances = true
	__GongStructShape__000000_Form_CheckBox.NbInstances = 0
	__GongStructShape__000000_Form_CheckBox.Width = 240.000000
	__GongStructShape__000000_Form_CheckBox.Height = 78.000000
	__GongStructShape__000000_Form_CheckBox.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_Form_FormCell.Name = `Form-FormCell`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField]
	__GongStructShape__000001_Form_FormCell.Identifier = `ref_models.FormField`
	__GongStructShape__000001_Form_FormCell.ShowNbInstances = true
	__GongStructShape__000001_Form_FormCell.NbInstances = 0
	__GongStructShape__000001_Form_FormCell.Width = 263.000000
	__GongStructShape__000001_Form_FormCell.Height = 123.000000
	__GongStructShape__000001_Form_FormCell.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_Form_FormCellFloat64.Name = `Form-FormCellFloat64`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldFloat64]
	__GongStructShape__000002_Form_FormCellFloat64.Identifier = `ref_models.FormFieldFloat64`
	__GongStructShape__000002_Form_FormCellFloat64.ShowNbInstances = true
	__GongStructShape__000002_Form_FormCellFloat64.NbInstances = 0
	__GongStructShape__000002_Form_FormCellFloat64.Width = 240.000000
	__GongStructShape__000002_Form_FormCellFloat64.Height = 78.000000
	__GongStructShape__000002_Form_FormCellFloat64.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_Form_FormCellInt.Name = `Form-FormCellInt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldInt]
	__GongStructShape__000003_Form_FormCellInt.Identifier = `ref_models.FormFieldInt`
	__GongStructShape__000003_Form_FormCellInt.ShowNbInstances = true
	__GongStructShape__000003_Form_FormCellInt.NbInstances = 0
	__GongStructShape__000003_Form_FormCellInt.Width = 240.000000
	__GongStructShape__000003_Form_FormCellInt.Height = 78.000000
	__GongStructShape__000003_Form_FormCellInt.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_Form_FormCellString.Name = `Form-FormCellString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldString]
	__GongStructShape__000004_Form_FormCellString.Identifier = `ref_models.FormFieldString`
	__GongStructShape__000004_Form_FormCellString.ShowNbInstances = true
	__GongStructShape__000004_Form_FormCellString.NbInstances = 0
	__GongStructShape__000004_Form_FormCellString.Width = 240.000000
	__GongStructShape__000004_Form_FormCellString.Height = 78.000000
	__GongStructShape__000004_Form_FormCellString.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000005_Form_FormDiv.Name = `Form-FormDiv`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormDiv]
	__GongStructShape__000005_Form_FormDiv.Identifier = `ref_models.FormDiv`
	__GongStructShape__000005_Form_FormDiv.ShowNbInstances = true
	__GongStructShape__000005_Form_FormDiv.NbInstances = 0
	__GongStructShape__000005_Form_FormDiv.Width = 240.000000
	__GongStructShape__000005_Form_FormDiv.Height = 78.000000
	__GongStructShape__000005_Form_FormDiv.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000006_Form_FormEditAssocButton.Name = `Form-FormEditAssocButton`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormEditAssocButton]
	__GongStructShape__000006_Form_FormEditAssocButton.Identifier = `ref_models.FormEditAssocButton`
	__GongStructShape__000006_Form_FormEditAssocButton.ShowNbInstances = false
	__GongStructShape__000006_Form_FormEditAssocButton.NbInstances = 0
	__GongStructShape__000006_Form_FormEditAssocButton.Width = 240.000000
	__GongStructShape__000006_Form_FormEditAssocButton.Height = 63.000000
	__GongStructShape__000006_Form_FormEditAssocButton.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000007_Form_FormFieldDate.Name = `Form-FormFieldDate`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldDate]
	__GongStructShape__000007_Form_FormFieldDate.Identifier = `ref_models.FormFieldDate`
	__GongStructShape__000007_Form_FormFieldDate.ShowNbInstances = true
	__GongStructShape__000007_Form_FormFieldDate.NbInstances = 0
	__GongStructShape__000007_Form_FormFieldDate.Width = 240.000000
	__GongStructShape__000007_Form_FormFieldDate.Height = 78.000000
	__GongStructShape__000007_Form_FormFieldDate.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000008_Form_FormFieldDateTime.Name = `Form-FormFieldDateTime`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldDateTime]
	__GongStructShape__000008_Form_FormFieldDateTime.Identifier = `ref_models.FormFieldDateTime`
	__GongStructShape__000008_Form_FormFieldDateTime.ShowNbInstances = true
	__GongStructShape__000008_Form_FormFieldDateTime.NbInstances = 0
	__GongStructShape__000008_Form_FormFieldDateTime.Width = 240.000000
	__GongStructShape__000008_Form_FormFieldDateTime.Height = 63.000000
	__GongStructShape__000008_Form_FormFieldDateTime.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000009_Form_FormFieldSelect.Name = `Form-FormFieldSelect`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldSelect]
	__GongStructShape__000009_Form_FormFieldSelect.Identifier = `ref_models.FormFieldSelect`
	__GongStructShape__000009_Form_FormFieldSelect.ShowNbInstances = false
	__GongStructShape__000009_Form_FormFieldSelect.NbInstances = 0
	__GongStructShape__000009_Form_FormFieldSelect.Width = 240.000000
	__GongStructShape__000009_Form_FormFieldSelect.Height = 78.000000
	__GongStructShape__000009_Form_FormFieldSelect.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000010_Form_FormFieldTime.Name = `Form-FormFieldTime`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldTime]
	__GongStructShape__000010_Form_FormFieldTime.Identifier = `ref_models.FormFieldTime`
	__GongStructShape__000010_Form_FormFieldTime.ShowNbInstances = true
	__GongStructShape__000010_Form_FormFieldTime.NbInstances = 0
	__GongStructShape__000010_Form_FormFieldTime.Width = 240.000000
	__GongStructShape__000010_Form_FormFieldTime.Height = 93.000000
	__GongStructShape__000010_Form_FormFieldTime.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000011_Form_FormGroup.Name = `Form-FormGroup`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormGroup]
	__GongStructShape__000011_Form_FormGroup.Identifier = `ref_models.FormGroup`
	__GongStructShape__000011_Form_FormGroup.ShowNbInstances = true
	__GongStructShape__000011_Form_FormGroup.NbInstances = 0
	__GongStructShape__000011_Form_FormGroup.Width = 240.000000
	__GongStructShape__000011_Form_FormGroup.Height = 63.000000
	__GongStructShape__000011_Form_FormGroup.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000012_Form_Option.Name = `Form-Option`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Option]
	__GongStructShape__000012_Form_Option.Identifier = `ref_models.Option`
	__GongStructShape__000012_Form_Option.ShowNbInstances = false
	__GongStructShape__000012_Form_Option.NbInstances = 0
	__GongStructShape__000012_Form_Option.Width = 240.000000
	__GongStructShape__000012_Form_Option.Height = 63.000000
	__GongStructShape__000012_Form_Option.IsSelected = false

	// Link values setup
	__Link__000000_CheckBoxs.Name = `CheckBoxs`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormDiv.CheckBoxs]
	__Link__000000_CheckBoxs.Identifier = `ref_models.FormDiv.CheckBoxs`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.CheckBox]
	__Link__000000_CheckBoxs.Fieldtypename = `ref_models.CheckBox`
	__Link__000000_CheckBoxs.FieldOffsetX = -95.000000
	__Link__000000_CheckBoxs.FieldOffsetY = -22.000000
	__Link__000000_CheckBoxs.TargetMultiplicity = models.MANY
	__Link__000000_CheckBoxs.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_CheckBoxs.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_CheckBoxs.SourceMultiplicity = models.ZERO_ONE
	__Link__000000_CheckBoxs.SourceMultiplicityOffsetX = 9.000000
	__Link__000000_CheckBoxs.SourceMultiplicityOffsetY = -11.000000
	__Link__000000_CheckBoxs.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_CheckBoxs.StartRatio = 0.391026
	__Link__000000_CheckBoxs.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_CheckBoxs.EndRatio = 0.500000
	__Link__000000_CheckBoxs.CornerOffsetRatio = 1.540007

	// Link values setup
	__Link__000001_FormCellFloat64.Name = `FormCellFloat64`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.FormFieldFloat64]
	__Link__000001_FormCellFloat64.Identifier = `ref_models.FormField.FormFieldFloat64`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldFloat64]
	__Link__000001_FormCellFloat64.Fieldtypename = `ref_models.FormFieldFloat64`
	__Link__000001_FormCellFloat64.FieldOffsetX = -131.000000
	__Link__000001_FormCellFloat64.FieldOffsetY = -21.000000
	__Link__000001_FormCellFloat64.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_FormCellFloat64.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_FormCellFloat64.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_FormCellFloat64.SourceMultiplicity = models.MANY
	__Link__000001_FormCellFloat64.SourceMultiplicityOffsetX = 27.000000
	__Link__000001_FormCellFloat64.SourceMultiplicityOffsetY = 28.000000
	__Link__000001_FormCellFloat64.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_FormCellFloat64.StartRatio = 0.500000
	__Link__000001_FormCellFloat64.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_FormCellFloat64.EndRatio = 0.500000
	__Link__000001_FormCellFloat64.CornerOffsetRatio = 1.731641

	// Link values setup
	__Link__000002_FormCellInt.Name = `FormCellInt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.FormFieldInt]
	__Link__000002_FormCellInt.Identifier = `ref_models.FormField.FormFieldInt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldInt]
	__Link__000002_FormCellInt.Fieldtypename = `ref_models.FormFieldInt`
	__Link__000002_FormCellInt.FieldOffsetX = -93.000000
	__Link__000002_FormCellInt.FieldOffsetY = -16.000000
	__Link__000002_FormCellInt.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_FormCellInt.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_FormCellInt.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_FormCellInt.SourceMultiplicity = models.MANY
	__Link__000002_FormCellInt.SourceMultiplicityOffsetX = 14.000000
	__Link__000002_FormCellInt.SourceMultiplicityOffsetY = 41.000000
	__Link__000002_FormCellInt.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_FormCellInt.StartRatio = 0.500000
	__Link__000002_FormCellInt.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_FormCellInt.EndRatio = 0.500000
	__Link__000002_FormCellInt.CornerOffsetRatio = 1.731641

	// Link values setup
	__Link__000003_FormCellString.Name = `FormCellString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.FormFieldString]
	__Link__000003_FormCellString.Identifier = `ref_models.FormField.FormFieldString`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldString]
	__Link__000003_FormCellString.Fieldtypename = `ref_models.FormFieldString`
	__Link__000003_FormCellString.FieldOffsetX = -119.000000
	__Link__000003_FormCellString.FieldOffsetY = -22.000000
	__Link__000003_FormCellString.TargetMultiplicity = models.ZERO_ONE
	__Link__000003_FormCellString.TargetMultiplicityOffsetX = -50.000000
	__Link__000003_FormCellString.TargetMultiplicityOffsetY = 16.000000
	__Link__000003_FormCellString.SourceMultiplicity = models.MANY
	__Link__000003_FormCellString.SourceMultiplicityOffsetX = 16.000000
	__Link__000003_FormCellString.SourceMultiplicityOffsetY = 28.000000
	__Link__000003_FormCellString.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_FormCellString.StartRatio = 0.500000
	__Link__000003_FormCellString.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_FormCellString.EndRatio = 0.500000
	__Link__000003_FormCellString.CornerOffsetRatio = 1.731641

	// Link values setup
	__Link__000004_FormDivs.Name = `FormDivs`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormGroup.FormDivs]
	__Link__000004_FormDivs.Identifier = `ref_models.FormGroup.FormDivs`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormDiv]
	__Link__000004_FormDivs.Fieldtypename = `ref_models.FormDiv`
	__Link__000004_FormDivs.FieldOffsetX = 18.000000
	__Link__000004_FormDivs.FieldOffsetY = -19.000000
	__Link__000004_FormDivs.TargetMultiplicity = models.MANY
	__Link__000004_FormDivs.TargetMultiplicityOffsetX = -42.000000
	__Link__000004_FormDivs.TargetMultiplicityOffsetY = -15.000000
	__Link__000004_FormDivs.SourceMultiplicity = models.ZERO_ONE
	__Link__000004_FormDivs.SourceMultiplicityOffsetX = -40.000000
	__Link__000004_FormDivs.SourceMultiplicityOffsetY = 23.000000
	__Link__000004_FormDivs.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000004_FormDivs.StartRatio = 0.381673
	__Link__000004_FormDivs.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000004_FormDivs.EndRatio = 0.435840
	__Link__000004_FormDivs.CornerOffsetRatio = 1.761905

	// Link values setup
	__Link__000005_FormEditAssocButton.Name = `FormEditAssocButton`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormDiv.FormEditAssocButton]
	__Link__000005_FormEditAssocButton.Identifier = `ref_models.FormDiv.FormEditAssocButton`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormEditAssocButton]
	__Link__000005_FormEditAssocButton.Fieldtypename = `ref_models.FormEditAssocButton`
	__Link__000005_FormEditAssocButton.FieldOffsetX = -177.000000
	__Link__000005_FormEditAssocButton.FieldOffsetY = -21.000000
	__Link__000005_FormEditAssocButton.TargetMultiplicity = models.ZERO_ONE
	__Link__000005_FormEditAssocButton.TargetMultiplicityOffsetX = -56.000000
	__Link__000005_FormEditAssocButton.TargetMultiplicityOffsetY = 27.000000
	__Link__000005_FormEditAssocButton.SourceMultiplicity = models.MANY
	__Link__000005_FormEditAssocButton.SourceMultiplicityOffsetX = 10.000000
	__Link__000005_FormEditAssocButton.SourceMultiplicityOffsetY = 24.000000
	__Link__000005_FormEditAssocButton.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_FormEditAssocButton.StartRatio = 0.544872
	__Link__000005_FormEditAssocButton.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_FormEditAssocButton.EndRatio = 0.500000
	__Link__000005_FormEditAssocButton.CornerOffsetRatio = 1.531673

	// Link values setup
	__Link__000006_FormFieldDate.Name = `FormFieldDate`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.FormFieldDate]
	__Link__000006_FormFieldDate.Identifier = `ref_models.FormField.FormFieldDate`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldDate]
	__Link__000006_FormFieldDate.Fieldtypename = `ref_models.FormFieldDate`
	__Link__000006_FormFieldDate.FieldOffsetX = -116.000000
	__Link__000006_FormFieldDate.FieldOffsetY = -22.000000
	__Link__000006_FormFieldDate.TargetMultiplicity = models.ZERO_ONE
	__Link__000006_FormFieldDate.TargetMultiplicityOffsetX = -49.000000
	__Link__000006_FormFieldDate.TargetMultiplicityOffsetY = 19.000000
	__Link__000006_FormFieldDate.SourceMultiplicity = models.MANY
	__Link__000006_FormFieldDate.SourceMultiplicityOffsetX = 37.000000
	__Link__000006_FormFieldDate.SourceMultiplicityOffsetY = 39.000000
	__Link__000006_FormFieldDate.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_FormFieldDate.StartRatio = 0.500000
	__Link__000006_FormFieldDate.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_FormFieldDate.EndRatio = 0.500000
	__Link__000006_FormFieldDate.CornerOffsetRatio = 1.731673

	// Link values setup
	__Link__000007_FormFieldDateTime.Name = `FormFieldDateTime`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.FormFieldDateTime]
	__Link__000007_FormFieldDateTime.Identifier = `ref_models.FormField.FormFieldDateTime`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldDateTime]
	__Link__000007_FormFieldDateTime.Fieldtypename = `ref_models.FormFieldDateTime`
	__Link__000007_FormFieldDateTime.FieldOffsetX = -152.000000
	__Link__000007_FormFieldDateTime.FieldOffsetY = -21.000000
	__Link__000007_FormFieldDateTime.TargetMultiplicity = models.ZERO_ONE
	__Link__000007_FormFieldDateTime.TargetMultiplicityOffsetX = -50.000000
	__Link__000007_FormFieldDateTime.TargetMultiplicityOffsetY = 16.000000
	__Link__000007_FormFieldDateTime.SourceMultiplicity = models.MANY
	__Link__000007_FormFieldDateTime.SourceMultiplicityOffsetX = 30.000000
	__Link__000007_FormFieldDateTime.SourceMultiplicityOffsetY = 56.000000
	__Link__000007_FormFieldDateTime.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000007_FormFieldDateTime.StartRatio = 0.500000
	__Link__000007_FormFieldDateTime.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000007_FormFieldDateTime.EndRatio = 0.500000
	__Link__000007_FormFieldDateTime.CornerOffsetRatio = 1.728523

	// Link values setup
	__Link__000008_FormFieldSelect.Name = `FormFieldSelect`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.FormFieldSelect]
	__Link__000008_FormFieldSelect.Identifier = `ref_models.FormField.FormFieldSelect`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldSelect]
	__Link__000008_FormFieldSelect.Fieldtypename = `ref_models.FormFieldSelect`
	__Link__000008_FormFieldSelect.FieldOffsetX = -126.000000
	__Link__000008_FormFieldSelect.FieldOffsetY = -21.000000
	__Link__000008_FormFieldSelect.TargetMultiplicity = models.ZERO_ONE
	__Link__000008_FormFieldSelect.TargetMultiplicityOffsetX = -50.000000
	__Link__000008_FormFieldSelect.TargetMultiplicityOffsetY = 16.000000
	__Link__000008_FormFieldSelect.SourceMultiplicity = models.MANY
	__Link__000008_FormFieldSelect.SourceMultiplicityOffsetX = 11.000000
	__Link__000008_FormFieldSelect.SourceMultiplicityOffsetY = 58.000000
	__Link__000008_FormFieldSelect.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000008_FormFieldSelect.StartRatio = 0.500000
	__Link__000008_FormFieldSelect.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000008_FormFieldSelect.EndRatio = 0.500000
	__Link__000008_FormFieldSelect.CornerOffsetRatio = 1.728523

	// Link values setup
	__Link__000009_FormFieldTime.Name = `FormFieldTime`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField.FormFieldTime]
	__Link__000009_FormFieldTime.Identifier = `ref_models.FormField.FormFieldTime`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldTime]
	__Link__000009_FormFieldTime.Fieldtypename = `ref_models.FormFieldTime`
	__Link__000009_FormFieldTime.FieldOffsetX = -123.000000
	__Link__000009_FormFieldTime.FieldOffsetY = -25.000000
	__Link__000009_FormFieldTime.TargetMultiplicity = models.ZERO_ONE
	__Link__000009_FormFieldTime.TargetMultiplicityOffsetX = -46.000000
	__Link__000009_FormFieldTime.TargetMultiplicityOffsetY = 24.000000
	__Link__000009_FormFieldTime.SourceMultiplicity = models.MANY
	__Link__000009_FormFieldTime.SourceMultiplicityOffsetX = 41.000000
	__Link__000009_FormFieldTime.SourceMultiplicityOffsetY = 26.000000
	__Link__000009_FormFieldTime.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000009_FormFieldTime.StartRatio = 0.500000
	__Link__000009_FormFieldTime.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000009_FormFieldTime.EndRatio = 0.500000
	__Link__000009_FormFieldTime.CornerOffsetRatio = 1.731673

	// Link values setup
	__Link__000010_FormFields.Name = `FormFields`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormDiv.FormFields]
	__Link__000010_FormFields.Identifier = `ref_models.FormDiv.FormFields`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormField]
	__Link__000010_FormFields.Fieldtypename = `ref_models.FormField`
	__Link__000010_FormFields.FieldOffsetX = 24.000000
	__Link__000010_FormFields.FieldOffsetY = -19.000000
	__Link__000010_FormFields.TargetMultiplicity = models.MANY
	__Link__000010_FormFields.TargetMultiplicityOffsetX = -32.000000
	__Link__000010_FormFields.TargetMultiplicityOffsetY = -5.000000
	__Link__000010_FormFields.SourceMultiplicity = models.ZERO_ONE
	__Link__000010_FormFields.SourceMultiplicityOffsetX = -35.000000
	__Link__000010_FormFields.SourceMultiplicityOffsetY = 25.000000
	__Link__000010_FormFields.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000010_FormFields.StartRatio = 0.440007
	__Link__000010_FormFields.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000010_FormFields.EndRatio = 0.424341
	__Link__000010_FormFields.CornerOffsetRatio = 2.150794

	// Link values setup
	__Link__000011_Options.Name = `Options`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldSelect.Options]
	__Link__000011_Options.Identifier = `ref_models.FormFieldSelect.Options`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Option]
	__Link__000011_Options.Fieldtypename = `ref_models.Option`
	__Link__000011_Options.FieldOffsetX = -57.000000
	__Link__000011_Options.FieldOffsetY = -30.000000
	__Link__000011_Options.TargetMultiplicity = models.MANY
	__Link__000011_Options.TargetMultiplicityOffsetX = -27.000000
	__Link__000011_Options.TargetMultiplicityOffsetY = -2.000000
	__Link__000011_Options.SourceMultiplicity = models.ZERO_ONE
	__Link__000011_Options.SourceMultiplicityOffsetX = 10.000000
	__Link__000011_Options.SourceMultiplicityOffsetY = -9.000000
	__Link__000011_Options.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000011_Options.StartRatio = 0.423077
	__Link__000011_Options.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000011_Options.EndRatio = 0.301587
	__Link__000011_Options.CornerOffsetRatio = 1.210840

	// Link values setup
	__Link__000012_Value.Name = `Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.FormFieldSelect.Value]
	__Link__000012_Value.Identifier = `ref_models.FormFieldSelect.Value`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Option]
	__Link__000012_Value.Fieldtypename = `ref_models.Option`
	__Link__000012_Value.FieldOffsetX = -46.000000
	__Link__000012_Value.FieldOffsetY = 26.000000
	__Link__000012_Value.TargetMultiplicity = models.ZERO_ONE
	__Link__000012_Value.TargetMultiplicityOffsetX = -41.000000
	__Link__000012_Value.TargetMultiplicityOffsetY = -8.000000
	__Link__000012_Value.SourceMultiplicity = models.MANY
	__Link__000012_Value.SourceMultiplicityOffsetX = 17.000000
	__Link__000012_Value.SourceMultiplicityOffsetY = -3.000000
	__Link__000012_Value.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000012_Value.StartRatio = 0.961538
	__Link__000012_Value.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000012_Value.EndRatio = 1.000000
	__Link__000012_Value.CornerOffsetRatio = 1.248340

	// Position values setup
	__Position__000000_Pos_Form_CheckBox.X = 643.000000
	__Position__000000_Pos_Form_CheckBox.Y = 39.000000
	__Position__000000_Pos_Form_CheckBox.Name = `Pos-Form-CheckBox`

	// Position values setup
	__Position__000001_Pos_Form_FormCell.X = 15.000000
	__Position__000001_Pos_Form_FormCell.Y = 434.000000
	__Position__000001_Pos_Form_FormCell.Name = `Pos-Form-FormCell`

	// Position values setup
	__Position__000002_Pos_Form_FormCellFloat64.X = 641.000000
	__Position__000002_Pos_Form_FormCellFloat64.Y = 389.000000
	__Position__000002_Pos_Form_FormCellFloat64.Name = `Pos-Form-FormCellFloat64`

	// Position values setup
	__Position__000003_Pos_Form_FormCellInt.X = 641.000000
	__Position__000003_Pos_Form_FormCellInt.Y = 284.000000
	__Position__000003_Pos_Form_FormCellInt.Name = `Pos-Form-FormCellInt`

	// Position values setup
	__Position__000004_Pos_Form_FormCellString.X = 643.000000
	__Position__000004_Pos_Form_FormCellString.Y = 156.000000
	__Position__000004_Pos_Form_FormCellString.Name = `Pos-Form-FormCellString`

	// Position values setup
	__Position__000005_Pos_Form_FormDiv.X = 17.000000
	__Position__000005_Pos_Form_FormDiv.Y = 241.000000
	__Position__000005_Pos_Form_FormDiv.Name = `Pos-Form-FormDiv`

	// Position values setup
	__Position__000006_Pos_Form_FormEditAssocButton.X = 660.000000
	__Position__000006_Pos_Form_FormEditAssocButton.Y = 1005.000000
	__Position__000006_Pos_Form_FormEditAssocButton.Name = `Pos-Form-FormEditAssocButton`

	// Position values setup
	__Position__000007_Pos_Form_FormFieldDate.X = 639.000000
	__Position__000007_Pos_Form_FormFieldDate.Y = 515.000000
	__Position__000007_Pos_Form_FormFieldDate.Name = `Pos-Form-FormFieldDate`

	// Position values setup
	__Position__000008_Pos_Form_FormFieldDateTime.X = 650.000000
	__Position__000008_Pos_Form_FormFieldDateTime.Y = 767.000000
	__Position__000008_Pos_Form_FormFieldDateTime.Name = `Pos-Form-FormFieldDateTime`

	// Position values setup
	__Position__000009_Pos_Form_FormFieldSelect.X = 649.000000
	__Position__000009_Pos_Form_FormFieldSelect.Y = 872.000000
	__Position__000009_Pos_Form_FormFieldSelect.Name = `Pos-Form-FormFieldSelect`

	// Position values setup
	__Position__000010_Pos_Form_FormFieldTime.X = 647.000000
	__Position__000010_Pos_Form_FormFieldTime.Y = 610.000000
	__Position__000010_Pos_Form_FormFieldTime.Name = `Pos-Form-FormFieldTime`

	// Position values setup
	__Position__000011_Pos_Form_FormGroup.X = 30.000000
	__Position__000011_Pos_Form_FormGroup.Y = 59.000000
	__Position__000011_Pos_Form_FormGroup.Name = `Pos-Form-FormGroup`

	// Position values setup
	__Position__000012_Pos_Form_InputTypeEnum.X = 28.000000
	__Position__000012_Pos_Form_InputTypeEnum.Y = 596.000000
	__Position__000012_Pos_Form_InputTypeEnum.Name = `Pos-Form-InputTypeEnum`

	// Position values setup
	__Position__000013_Pos_Form_Option.X = 1038.000000
	__Position__000013_Pos_Form_Option.Y = 884.000000
	__Position__000013_Pos_Form_Option.Name = `Pos-Form-Option`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellFloat64.X = 409.500000
	__Vertice__000000_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellFloat64.Y = 128.500000
	__Vertice__000000_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellFloat64.Name = `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellFloat64`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellInt.X = 413.000000
	__Vertice__000001_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellInt.Y = 108.500000
	__Vertice__000001_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellInt.Name = `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellInt`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellString.X = 435.500000
	__Vertice__000002_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellString.Y = 126.500000
	__Vertice__000002_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellString.Name = `Verticle in class diagram Form in middle between Form-FormCell and Form-FormCellString`

	// Vertice values setup
	__Vertice__000003_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldDate.X = 707.000000
	__Vertice__000003_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldDate.Y = 413.500000
	__Vertice__000003_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldDate.Name = `Verticle in class diagram Form in middle between Form-FormCell and Form-FormFieldDate`

	// Vertice values setup
	__Vertice__000004_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldDateTime.X = 727.000000
	__Vertice__000004_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldDateTime.Y = 662.000000
	__Vertice__000004_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldDateTime.Name = `Verticle in class diagram Form in middle between Form-FormCell and Form-FormFieldDateTime`

	// Vertice values setup
	__Vertice__000005_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldSelect.X = 726.500000
	__Vertice__000005_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldSelect.Y = 714.500000
	__Vertice__000005_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldSelect.Name = `Verticle in class diagram Form in middle between Form-FormCell and Form-FormFieldSelect`

	// Vertice values setup
	__Vertice__000006_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldTime.X = 711.000000
	__Vertice__000006_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldTime.Y = 461.000000
	__Vertice__000006_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldTime.Name = `Verticle in class diagram Form in middle between Form-FormCell and Form-FormFieldTime`

	// Vertice values setup
	__Vertice__000007_Verticle_in_class_diagram_Form_in_middle_between_Form_FormDiv_and_Form_CheckBox.X = 690.000000
	__Vertice__000007_Verticle_in_class_diagram_Form_in_middle_between_Form_FormDiv_and_Form_CheckBox.Y = 171.500000
	__Vertice__000007_Verticle_in_class_diagram_Form_in_middle_between_Form_FormDiv_and_Form_CheckBox.Name = `Verticle in class diagram Form in middle between Form-FormDiv and Form-CheckBox`

	// Vertice values setup
	__Vertice__000008_Verticle_in_class_diagram_Form_in_middle_between_Form_FormDiv_and_Form_FormCell.X = 385.500000
	__Vertice__000008_Verticle_in_class_diagram_Form_in_middle_between_Form_FormDiv_and_Form_FormCell.Y = 371.500000
	__Vertice__000008_Verticle_in_class_diagram_Form_in_middle_between_Form_FormDiv_and_Form_FormCell.Name = `Verticle in class diagram Form in middle between Form-FormDiv and Form-FormCell`

	// Vertice values setup
	__Vertice__000009_Verticle_in_class_diagram_Form_in_middle_between_Form_FormDiv_and_Form_FormEditAssocButton.X = 698.500000
	__Vertice__000009_Verticle_in_class_diagram_Form_in_middle_between_Form_FormDiv_and_Form_FormEditAssocButton.Y = 654.500000
	__Vertice__000009_Verticle_in_class_diagram_Form_in_middle_between_Form_FormDiv_and_Form_FormEditAssocButton.Name = `Verticle in class diagram Form in middle between Form-FormDiv and Form-FormEditAssocButton`

	// Vertice values setup
	__Vertice__000010_Verticle_in_class_diagram_Form_in_middle_between_Form_FormFieldSelect_and_Form_Option.X = 510.000000
	__Vertice__000010_Verticle_in_class_diagram_Form_in_middle_between_Form_FormFieldSelect_and_Form_Option.Y = 476.000000
	__Vertice__000010_Verticle_in_class_diagram_Form_in_middle_between_Form_FormFieldSelect_and_Form_Option.Name = `Verticle in class diagram Form in middle between Form-FormFieldSelect and Form-Option`

	// Vertice values setup
	__Vertice__000011_Verticle_in_class_diagram_Form_in_middle_between_Form_FormFieldSelect_and_Form_Option.X = 510.000000
	__Vertice__000011_Verticle_in_class_diagram_Form_in_middle_between_Form_FormFieldSelect_and_Form_Option.Y = 476.000000
	__Vertice__000011_Verticle_in_class_diagram_Form_in_middle_between_Form_FormFieldSelect_and_Form_Option.Name = `Verticle in class diagram Form in middle between Form-FormFieldSelect and Form-Option`

	// Vertice values setup
	__Vertice__000012_Verticle_in_class_diagram_Form_in_middle_between_Form_FormGroup_and_Form_FormDiv.X = 390.000000
	__Vertice__000012_Verticle_in_class_diagram_Form_in_middle_between_Form_FormGroup_and_Form_FormDiv.Y = 185.000000
	__Vertice__000012_Verticle_in_class_diagram_Form_in_middle_between_Form_FormGroup_and_Form_FormDiv.Name = `Verticle in class diagram Form in middle between Form-FormGroup and Form-FormDiv`

	// Setup of pointers
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000001_Form_FormCell)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000002_Form_FormCellFloat64)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000003_Form_FormCellInt)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000004_Form_FormCellString)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000010_Form_FormFieldTime)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000007_Form_FormFieldDate)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000011_Form_FormGroup)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000008_Form_FormFieldDateTime)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000005_Form_FormDiv)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000000_Form_CheckBox)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000006_Form_FormEditAssocButton)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000009_Form_FormFieldSelect)
	__Classdiagram__000000_Form.GongStructShapes = append(__Classdiagram__000000_Form.GongStructShapes, __GongStructShape__000012_Form_Option)
	__Classdiagram__000000_Form.GongEnumShapes = append(__Classdiagram__000000_Form.GongEnumShapes, __GongEnumShape__000000_Form_InputTypeEnum)
	__GongEnumShape__000000_Form_InputTypeEnum.Position = __Position__000012_Pos_Form_InputTypeEnum
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000012_Text)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000008_Password)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000007_Number)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000003_Email)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000011_Tel)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000001_Date)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000002_Datetime)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000013_Time)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000014_URL)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000010_Search)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000009_Range)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000000_Color)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000004_File)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000005_Hidden)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000006_Month)
	__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys = append(__GongEnumShape__000000_Form_InputTypeEnum.GongEnumValueEntrys, __GongEnumValueEntry__000015_Week)
	__GongStructShape__000000_Form_CheckBox.Position = __Position__000000_Pos_Form_CheckBox
	__GongStructShape__000000_Form_CheckBox.Fields = append(__GongStructShape__000000_Form_CheckBox.Fields, __Field__000009_Value)
	__GongStructShape__000001_Form_FormCell.Position = __Position__000001_Pos_Form_FormCell
	__GongStructShape__000001_Form_FormCell.Fields = append(__GongStructShape__000001_Form_FormCell.Fields, __Field__000002_Name)
	__GongStructShape__000001_Form_FormCell.Fields = append(__GongStructShape__000001_Form_FormCell.Fields, __Field__000000_InputTypeEnum)
	__GongStructShape__000001_Form_FormCell.Fields = append(__GongStructShape__000001_Form_FormCell.Fields, __Field__000001_Label)
	__GongStructShape__000001_Form_FormCell.Fields = append(__GongStructShape__000001_Form_FormCell.Fields, __Field__000005_Placeholder)
	__GongStructShape__000001_Form_FormCell.Links = append(__GongStructShape__000001_Form_FormCell.Links, __Link__000001_FormCellFloat64)
	__GongStructShape__000001_Form_FormCell.Links = append(__GongStructShape__000001_Form_FormCell.Links, __Link__000002_FormCellInt)
	__GongStructShape__000001_Form_FormCell.Links = append(__GongStructShape__000001_Form_FormCell.Links, __Link__000003_FormCellString)
	__GongStructShape__000001_Form_FormCell.Links = append(__GongStructShape__000001_Form_FormCell.Links, __Link__000009_FormFieldTime)
	__GongStructShape__000001_Form_FormCell.Links = append(__GongStructShape__000001_Form_FormCell.Links, __Link__000006_FormFieldDate)
	__GongStructShape__000001_Form_FormCell.Links = append(__GongStructShape__000001_Form_FormCell.Links, __Link__000007_FormFieldDateTime)
	__GongStructShape__000001_Form_FormCell.Links = append(__GongStructShape__000001_Form_FormCell.Links, __Link__000008_FormFieldSelect)
	__GongStructShape__000002_Form_FormCellFloat64.Position = __Position__000002_Pos_Form_FormCellFloat64
	__GongStructShape__000002_Form_FormCellFloat64.Fields = append(__GongStructShape__000002_Form_FormCellFloat64.Fields, __Field__000011_Value)
	__GongStructShape__000003_Form_FormCellInt.Position = __Position__000003_Pos_Form_FormCellInt
	__GongStructShape__000003_Form_FormCellInt.Fields = append(__GongStructShape__000003_Form_FormCellInt.Fields, __Field__000010_Value)
	__GongStructShape__000004_Form_FormCellString.Position = __Position__000004_Pos_Form_FormCellString
	__GongStructShape__000004_Form_FormCellString.Fields = append(__GongStructShape__000004_Form_FormCellString.Fields, __Field__000012_Value)
	__GongStructShape__000005_Form_FormDiv.Position = __Position__000005_Pos_Form_FormDiv
	__GongStructShape__000005_Form_FormDiv.Fields = append(__GongStructShape__000005_Form_FormDiv.Fields, __Field__000004_Name)
	__GongStructShape__000005_Form_FormDiv.Links = append(__GongStructShape__000005_Form_FormDiv.Links, __Link__000010_FormFields)
	__GongStructShape__000005_Form_FormDiv.Links = append(__GongStructShape__000005_Form_FormDiv.Links, __Link__000000_CheckBoxs)
	__GongStructShape__000005_Form_FormDiv.Links = append(__GongStructShape__000005_Form_FormDiv.Links, __Link__000005_FormEditAssocButton)
	__GongStructShape__000006_Form_FormEditAssocButton.Position = __Position__000006_Pos_Form_FormEditAssocButton
	__GongStructShape__000007_Form_FormFieldDate.Position = __Position__000007_Pos_Form_FormFieldDate
	__GongStructShape__000007_Form_FormFieldDate.Fields = append(__GongStructShape__000007_Form_FormFieldDate.Fields, __Field__000008_Value)
	__GongStructShape__000008_Form_FormFieldDateTime.Position = __Position__000008_Pos_Form_FormFieldDateTime
	__GongStructShape__000009_Form_FormFieldSelect.Position = __Position__000009_Pos_Form_FormFieldSelect
	__GongStructShape__000009_Form_FormFieldSelect.Fields = append(__GongStructShape__000009_Form_FormFieldSelect.Fields, __Field__000003_Name)
	__GongStructShape__000009_Form_FormFieldSelect.Links = append(__GongStructShape__000009_Form_FormFieldSelect.Links, __Link__000011_Options)
	__GongStructShape__000009_Form_FormFieldSelect.Links = append(__GongStructShape__000009_Form_FormFieldSelect.Links, __Link__000012_Value)
	__GongStructShape__000010_Form_FormFieldTime.Position = __Position__000010_Pos_Form_FormFieldTime
	__GongStructShape__000010_Form_FormFieldTime.Fields = append(__GongStructShape__000010_Form_FormFieldTime.Fields, __Field__000007_Value)
	__GongStructShape__000010_Form_FormFieldTime.Fields = append(__GongStructShape__000010_Form_FormFieldTime.Fields, __Field__000006_Step)
	__GongStructShape__000011_Form_FormGroup.Position = __Position__000011_Pos_Form_FormGroup
	__GongStructShape__000011_Form_FormGroup.Links = append(__GongStructShape__000011_Form_FormGroup.Links, __Link__000004_FormDivs)
	__GongStructShape__000012_Form_Option.Position = __Position__000013_Pos_Form_Option
	__Link__000000_CheckBoxs.Middlevertice = __Vertice__000007_Verticle_in_class_diagram_Form_in_middle_between_Form_FormDiv_and_Form_CheckBox
	__Link__000001_FormCellFloat64.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellFloat64
	__Link__000002_FormCellInt.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellInt
	__Link__000003_FormCellString.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormCellString
	__Link__000004_FormDivs.Middlevertice = __Vertice__000012_Verticle_in_class_diagram_Form_in_middle_between_Form_FormGroup_and_Form_FormDiv
	__Link__000005_FormEditAssocButton.Middlevertice = __Vertice__000009_Verticle_in_class_diagram_Form_in_middle_between_Form_FormDiv_and_Form_FormEditAssocButton
	__Link__000006_FormFieldDate.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldDate
	__Link__000007_FormFieldDateTime.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldDateTime
	__Link__000008_FormFieldSelect.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldSelect
	__Link__000009_FormFieldTime.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_Form_in_middle_between_Form_FormCell_and_Form_FormFieldTime
	__Link__000010_FormFields.Middlevertice = __Vertice__000008_Verticle_in_class_diagram_Form_in_middle_between_Form_FormDiv_and_Form_FormCell
	__Link__000011_Options.Middlevertice = __Vertice__000011_Verticle_in_class_diagram_Form_in_middle_between_Form_FormFieldSelect_and_Form_Option
	__Link__000012_Value.Middlevertice = __Vertice__000010_Verticle_in_class_diagram_Form_in_middle_between_Form_FormFieldSelect_and_Form_Option
}
