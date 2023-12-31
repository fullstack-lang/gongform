// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const marshallRes = `package {{PackageName}}

import (
	"time"

	"{{ModelsPackageName}}"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_{{databaseName}} models.StageStruct
var ___dummy__Time_{{databaseName}} time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_{{databaseName}} map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["{{databaseName}}"] = {{databaseName}}Injection
// }

// {{databaseName}}Injection will stage objects of database "{{databaseName}}"
func {{databaseName}}Injection(stage *models.StageStruct) {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}

`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

const StringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const NumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const PointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const SliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const TimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *StageStruct) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Println("filename of marshall output is " + name)
	newBase := filepath.Base(file.Name())

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(newBase, ".go", ""))
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField 

	// insertion initialization of objects to stage
	map_CheckBox_Identifiers := make(map[*CheckBox]string)
	_ = map_CheckBox_Identifiers

	checkboxOrdered := []*CheckBox{}
	for checkbox := range stage.CheckBoxs {
		checkboxOrdered = append(checkboxOrdered, checkbox)
	}
	sort.Slice(checkboxOrdered[:], func(i, j int) bool {
		return checkboxOrdered[i].Name < checkboxOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of CheckBox"
	for idx, checkbox := range checkboxOrdered {

		id = generatesIdentifier("CheckBox", idx, checkbox.Name)
		map_CheckBox_Identifiers[checkbox] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "CheckBox")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", checkbox.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// CheckBox values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(checkbox.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", checkbox.Value))
		initializerStatements += setValueField

	}

	map_FormDiv_Identifiers := make(map[*FormDiv]string)
	_ = map_FormDiv_Identifiers

	formdivOrdered := []*FormDiv{}
	for formdiv := range stage.FormDivs {
		formdivOrdered = append(formdivOrdered, formdiv)
	}
	sort.Slice(formdivOrdered[:], func(i, j int) bool {
		return formdivOrdered[i].Name < formdivOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of FormDiv"
	for idx, formdiv := range formdivOrdered {

		id = generatesIdentifier("FormDiv", idx, formdiv.Name)
		map_FormDiv_Identifiers[formdiv] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormDiv")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formdiv.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// FormDiv values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formdiv.Name))
		initializerStatements += setValueField

	}

	map_FormEditAssocButton_Identifiers := make(map[*FormEditAssocButton]string)
	_ = map_FormEditAssocButton_Identifiers

	formeditassocbuttonOrdered := []*FormEditAssocButton{}
	for formeditassocbutton := range stage.FormEditAssocButtons {
		formeditassocbuttonOrdered = append(formeditassocbuttonOrdered, formeditassocbutton)
	}
	sort.Slice(formeditassocbuttonOrdered[:], func(i, j int) bool {
		return formeditassocbuttonOrdered[i].Name < formeditassocbuttonOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of FormEditAssocButton"
	for idx, formeditassocbutton := range formeditassocbuttonOrdered {

		id = generatesIdentifier("FormEditAssocButton", idx, formeditassocbutton.Name)
		map_FormEditAssocButton_Identifiers[formeditassocbutton] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormEditAssocButton")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formeditassocbutton.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// FormEditAssocButton values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formeditassocbutton.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Label")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formeditassocbutton.Label))
		initializerStatements += setValueField

	}

	map_FormField_Identifiers := make(map[*FormField]string)
	_ = map_FormField_Identifiers

	formfieldOrdered := []*FormField{}
	for formfield := range stage.FormFields {
		formfieldOrdered = append(formfieldOrdered, formfield)
	}
	sort.Slice(formfieldOrdered[:], func(i, j int) bool {
		return formfieldOrdered[i].Name < formfieldOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of FormField"
	for idx, formfield := range formfieldOrdered {

		id = generatesIdentifier("FormField", idx, formfield.Name)
		map_FormField_Identifiers[formfield] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormField")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formfield.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// FormField values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfield.Name))
		initializerStatements += setValueField

		if formfield.InputTypeEnum != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "InputTypeEnum")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+formfield.InputTypeEnum.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Label")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfield.Label))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Placeholder")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfield.Placeholder))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasBespokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfield.HasBespokeWidth))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BespokeWidthPx")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", formfield.BespokeWidthPx))
		initializerStatements += setValueField

	}

	map_FormFieldDate_Identifiers := make(map[*FormFieldDate]string)
	_ = map_FormFieldDate_Identifiers

	formfielddateOrdered := []*FormFieldDate{}
	for formfielddate := range stage.FormFieldDates {
		formfielddateOrdered = append(formfielddateOrdered, formfielddate)
	}
	sort.Slice(formfielddateOrdered[:], func(i, j int) bool {
		return formfielddateOrdered[i].Name < formfielddateOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of FormFieldDate"
	for idx, formfielddate := range formfielddateOrdered {

		id = generatesIdentifier("FormFieldDate", idx, formfielddate.Name)
		map_FormFieldDate_Identifiers[formfielddate] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldDate")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formfielddate.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// FormFieldDate values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfielddate.Name))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", formfielddate.Value.String())
		initializerStatements += setValueField

	}

	map_FormFieldDateTime_Identifiers := make(map[*FormFieldDateTime]string)
	_ = map_FormFieldDateTime_Identifiers

	formfielddatetimeOrdered := []*FormFieldDateTime{}
	for formfielddatetime := range stage.FormFieldDateTimes {
		formfielddatetimeOrdered = append(formfielddatetimeOrdered, formfielddatetime)
	}
	sort.Slice(formfielddatetimeOrdered[:], func(i, j int) bool {
		return formfielddatetimeOrdered[i].Name < formfielddatetimeOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of FormFieldDateTime"
	for idx, formfielddatetime := range formfielddatetimeOrdered {

		id = generatesIdentifier("FormFieldDateTime", idx, formfielddatetime.Name)
		map_FormFieldDateTime_Identifiers[formfielddatetime] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldDateTime")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formfielddatetime.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// FormFieldDateTime values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfielddatetime.Name))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", formfielddatetime.Value.String())
		initializerStatements += setValueField

	}

	map_FormFieldFloat64_Identifiers := make(map[*FormFieldFloat64]string)
	_ = map_FormFieldFloat64_Identifiers

	formfieldfloat64Ordered := []*FormFieldFloat64{}
	for formfieldfloat64 := range stage.FormFieldFloat64s {
		formfieldfloat64Ordered = append(formfieldfloat64Ordered, formfieldfloat64)
	}
	sort.Slice(formfieldfloat64Ordered[:], func(i, j int) bool {
		return formfieldfloat64Ordered[i].Name < formfieldfloat64Ordered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of FormFieldFloat64"
	for idx, formfieldfloat64 := range formfieldfloat64Ordered {

		id = generatesIdentifier("FormFieldFloat64", idx, formfieldfloat64.Name)
		map_FormFieldFloat64_Identifiers[formfieldfloat64] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldFloat64")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formfieldfloat64.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// FormFieldFloat64 values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfieldfloat64.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", formfieldfloat64.Value))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasMinValidator")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldfloat64.HasMinValidator))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MinValue")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", formfieldfloat64.MinValue))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasMaxValidator")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldfloat64.HasMaxValidator))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxValue")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", formfieldfloat64.MaxValue))
		initializerStatements += setValueField

	}

	map_FormFieldInt_Identifiers := make(map[*FormFieldInt]string)
	_ = map_FormFieldInt_Identifiers

	formfieldintOrdered := []*FormFieldInt{}
	for formfieldint := range stage.FormFieldInts {
		formfieldintOrdered = append(formfieldintOrdered, formfieldint)
	}
	sort.Slice(formfieldintOrdered[:], func(i, j int) bool {
		return formfieldintOrdered[i].Name < formfieldintOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of FormFieldInt"
	for idx, formfieldint := range formfieldintOrdered {

		id = generatesIdentifier("FormFieldInt", idx, formfieldint.Name)
		map_FormFieldInt_Identifiers[formfieldint] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldInt")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formfieldint.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// FormFieldInt values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfieldint.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", formfieldint.Value))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasMinValidator")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldint.HasMinValidator))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MinValue")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", formfieldint.MinValue))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasMaxValidator")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldint.HasMaxValidator))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxValue")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", formfieldint.MaxValue))
		initializerStatements += setValueField

	}

	map_FormFieldSelect_Identifiers := make(map[*FormFieldSelect]string)
	_ = map_FormFieldSelect_Identifiers

	formfieldselectOrdered := []*FormFieldSelect{}
	for formfieldselect := range stage.FormFieldSelects {
		formfieldselectOrdered = append(formfieldselectOrdered, formfieldselect)
	}
	sort.Slice(formfieldselectOrdered[:], func(i, j int) bool {
		return formfieldselectOrdered[i].Name < formfieldselectOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of FormFieldSelect"
	for idx, formfieldselect := range formfieldselectOrdered {

		id = generatesIdentifier("FormFieldSelect", idx, formfieldselect.Name)
		map_FormFieldSelect_Identifiers[formfieldselect] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldSelect")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formfieldselect.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// FormFieldSelect values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfieldselect.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CanBeEmpty")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldselect.CanBeEmpty))
		initializerStatements += setValueField

	}

	map_FormFieldString_Identifiers := make(map[*FormFieldString]string)
	_ = map_FormFieldString_Identifiers

	formfieldstringOrdered := []*FormFieldString{}
	for formfieldstring := range stage.FormFieldStrings {
		formfieldstringOrdered = append(formfieldstringOrdered, formfieldstring)
	}
	sort.Slice(formfieldstringOrdered[:], func(i, j int) bool {
		return formfieldstringOrdered[i].Name < formfieldstringOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of FormFieldString"
	for idx, formfieldstring := range formfieldstringOrdered {

		id = generatesIdentifier("FormFieldString", idx, formfieldstring.Name)
		map_FormFieldString_Identifiers[formfieldstring] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldString")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formfieldstring.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// FormFieldString values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfieldstring.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfieldstring.Value))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsTextArea")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldstring.IsTextArea))
		initializerStatements += setValueField

	}

	map_FormFieldTime_Identifiers := make(map[*FormFieldTime]string)
	_ = map_FormFieldTime_Identifiers

	formfieldtimeOrdered := []*FormFieldTime{}
	for formfieldtime := range stage.FormFieldTimes {
		formfieldtimeOrdered = append(formfieldtimeOrdered, formfieldtime)
	}
	sort.Slice(formfieldtimeOrdered[:], func(i, j int) bool {
		return formfieldtimeOrdered[i].Name < formfieldtimeOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of FormFieldTime"
	for idx, formfieldtime := range formfieldtimeOrdered {

		id = generatesIdentifier("FormFieldTime", idx, formfieldtime.Name)
		map_FormFieldTime_Identifiers[formfieldtime] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormFieldTime")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formfieldtime.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// FormFieldTime values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfieldtime.Name))
		initializerStatements += setValueField

		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", formfieldtime.Value.String())
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Step")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", formfieldtime.Step))
		initializerStatements += setValueField

	}

	map_FormGroup_Identifiers := make(map[*FormGroup]string)
	_ = map_FormGroup_Identifiers

	formgroupOrdered := []*FormGroup{}
	for formgroup := range stage.FormGroups {
		formgroupOrdered = append(formgroupOrdered, formgroup)
	}
	sort.Slice(formgroupOrdered[:], func(i, j int) bool {
		return formgroupOrdered[i].Name < formgroupOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of FormGroup"
	for idx, formgroup := range formgroupOrdered {

		id = generatesIdentifier("FormGroup", idx, formgroup.Name)
		map_FormGroup_Identifiers[formgroup] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormGroup")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formgroup.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// FormGroup values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formgroup.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Label")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formgroup.Label))
		initializerStatements += setValueField

	}

	map_FormSortAssocButton_Identifiers := make(map[*FormSortAssocButton]string)
	_ = map_FormSortAssocButton_Identifiers

	formsortassocbuttonOrdered := []*FormSortAssocButton{}
	for formsortassocbutton := range stage.FormSortAssocButtons {
		formsortassocbuttonOrdered = append(formsortassocbuttonOrdered, formsortassocbutton)
	}
	sort.Slice(formsortassocbuttonOrdered[:], func(i, j int) bool {
		return formsortassocbuttonOrdered[i].Name < formsortassocbuttonOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of FormSortAssocButton"
	for idx, formsortassocbutton := range formsortassocbuttonOrdered {

		id = generatesIdentifier("FormSortAssocButton", idx, formsortassocbutton.Name)
		map_FormSortAssocButton_Identifiers[formsortassocbutton] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FormSortAssocButton")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", formsortassocbutton.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// FormSortAssocButton values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formsortassocbutton.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Label")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formsortassocbutton.Label))
		initializerStatements += setValueField

	}

	map_Option_Identifiers := make(map[*Option]string)
	_ = map_Option_Identifiers

	optionOrdered := []*Option{}
	for option := range stage.Options {
		optionOrdered = append(optionOrdered, option)
	}
	sort.Slice(optionOrdered[:], func(i, j int) bool {
		return optionOrdered[i].Name < optionOrdered[j].Name
	})
	identifiersDecl += "\n\n	// Declarations of staged instances of Option"
	for idx, option := range optionOrdered {

		id = generatesIdentifier("Option", idx, option.Name)
		map_Option_Identifiers[option] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Option")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", option.Name)
		identifiersDecl += decl

		initializerStatements += "\n\n	// Option values setup"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(option.Name))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	for idx, checkbox := range checkboxOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("CheckBox", idx, checkbox.Name)
		map_CheckBox_Identifiers[checkbox] = id

		// Initialisation of values
	}

	for idx, formdiv := range formdivOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("FormDiv", idx, formdiv.Name)
		map_FormDiv_Identifiers[formdiv] = id

		// Initialisation of values
		for _, _formfield := range formdiv.FormFields {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormFields")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_FormField_Identifiers[_formfield])
			pointersInitializesStatements += setPointerField
		}

		for _, _checkbox := range formdiv.CheckBoxs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CheckBoxs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_CheckBox_Identifiers[_checkbox])
			pointersInitializesStatements += setPointerField
		}

		if formdiv.FormEditAssocButton != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormEditAssocButton")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_FormEditAssocButton_Identifiers[formdiv.FormEditAssocButton])
			pointersInitializesStatements += setPointerField
		}

		if formdiv.FormSortAssocButton != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormSortAssocButton")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_FormSortAssocButton_Identifiers[formdiv.FormSortAssocButton])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, formeditassocbutton := range formeditassocbuttonOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("FormEditAssocButton", idx, formeditassocbutton.Name)
		map_FormEditAssocButton_Identifiers[formeditassocbutton] = id

		// Initialisation of values
	}

	for idx, formfield := range formfieldOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("FormField", idx, formfield.Name)
		map_FormField_Identifiers[formfield] = id

		// Initialisation of values
		if formfield.FormFieldString != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormFieldString")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_FormFieldString_Identifiers[formfield.FormFieldString])
			pointersInitializesStatements += setPointerField
		}

		if formfield.FormFieldFloat64 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormFieldFloat64")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_FormFieldFloat64_Identifiers[formfield.FormFieldFloat64])
			pointersInitializesStatements += setPointerField
		}

		if formfield.FormFieldInt != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormFieldInt")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_FormFieldInt_Identifiers[formfield.FormFieldInt])
			pointersInitializesStatements += setPointerField
		}

		if formfield.FormFieldDate != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormFieldDate")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_FormFieldDate_Identifiers[formfield.FormFieldDate])
			pointersInitializesStatements += setPointerField
		}

		if formfield.FormFieldTime != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormFieldTime")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_FormFieldTime_Identifiers[formfield.FormFieldTime])
			pointersInitializesStatements += setPointerField
		}

		if formfield.FormFieldDateTime != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormFieldDateTime")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_FormFieldDateTime_Identifiers[formfield.FormFieldDateTime])
			pointersInitializesStatements += setPointerField
		}

		if formfield.FormFieldSelect != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormFieldSelect")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_FormFieldSelect_Identifiers[formfield.FormFieldSelect])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, formfielddate := range formfielddateOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("FormFieldDate", idx, formfielddate.Name)
		map_FormFieldDate_Identifiers[formfielddate] = id

		// Initialisation of values
	}

	for idx, formfielddatetime := range formfielddatetimeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("FormFieldDateTime", idx, formfielddatetime.Name)
		map_FormFieldDateTime_Identifiers[formfielddatetime] = id

		// Initialisation of values
	}

	for idx, formfieldfloat64 := range formfieldfloat64Ordered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("FormFieldFloat64", idx, formfieldfloat64.Name)
		map_FormFieldFloat64_Identifiers[formfieldfloat64] = id

		// Initialisation of values
	}

	for idx, formfieldint := range formfieldintOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("FormFieldInt", idx, formfieldint.Name)
		map_FormFieldInt_Identifiers[formfieldint] = id

		// Initialisation of values
	}

	for idx, formfieldselect := range formfieldselectOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("FormFieldSelect", idx, formfieldselect.Name)
		map_FormFieldSelect_Identifiers[formfieldselect] = id

		// Initialisation of values
		if formfieldselect.Value != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Value")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Option_Identifiers[formfieldselect.Value])
			pointersInitializesStatements += setPointerField
		}

		for _, _option := range formfieldselect.Options {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Options")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Option_Identifiers[_option])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, formfieldstring := range formfieldstringOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("FormFieldString", idx, formfieldstring.Name)
		map_FormFieldString_Identifiers[formfieldstring] = id

		// Initialisation of values
	}

	for idx, formfieldtime := range formfieldtimeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("FormFieldTime", idx, formfieldtime.Name)
		map_FormFieldTime_Identifiers[formfieldtime] = id

		// Initialisation of values
	}

	for idx, formgroup := range formgroupOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("FormGroup", idx, formgroup.Name)
		map_FormGroup_Identifiers[formgroup] = id

		// Initialisation of values
		for _, _formdiv := range formgroup.FormDivs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormDivs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_FormDiv_Identifiers[_formdiv])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, formsortassocbutton := range formsortassocbuttonOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("FormSortAssocButton", idx, formsortassocbutton.Name)
		map_FormSortAssocButton_Identifiers[formsortassocbutton] = id

		// Initialisation of values
	}

	for idx, option := range optionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Option", idx, option.Name)
		map_Option_Identifiers[option] = id

		// Initialisation of values
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar ___dummy__%s_%s %s.StageStruct",
				stage.MetaPackageImportAlias,
				strings.ReplaceAll(filepath.Base(name), ".go", ""),
				stage.MetaPackageImportAlias))

		var entries string

		// regenerate the map of doc link renaming
		// the key and value are set to the value because
		// if it has been renamed, this is the new value that matters
		valuesOrdered := make([]GONG__Identifier, 0)
		for _, value := range stage.Map_DocLink_Renaming {
			valuesOrdered = append(valuesOrdered, value)
		}
		sort.Slice(valuesOrdered[:], func(i, j int) bool {
			return valuesOrdered[i].Ident < valuesOrdered[j].Ident
		})
		for _, value := range valuesOrdered {

			// get the number of points in the value to find if it is a field
			// or a struct

			switch value.Type {
			case GONG__ENUM_CAST_INT:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident)
			case GONG__ENUM_CAST_STRING:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident)
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries += fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier)
			case GONG__IDENTIFIER_CONST:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident)
			case GONG__STRUCT_INSTANCE:
				entries += fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident)
			}
		}

		res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
	}

	fmt.Fprintln(file, res)
}

// unique identifier per struct
func generatesIdentifier(gongStructName string, idx int, instanceName string) (identifier string) {

	identifier = instanceName
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(instanceName, "_")

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}
