// generated code - do not edit
package probe

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongform/go/models"
)

const __dummmy__time = time.Nanosecond

// insertion point
func NewCheckBoxFormCallback(
	checkbox *models.CheckBox,
	playground *Playground,
) (checkboxFormCallback *CheckBoxFormCallback) {
	checkboxFormCallback = new(CheckBoxFormCallback)
	checkboxFormCallback.playground = playground
	checkboxFormCallback.checkbox = checkbox

	checkboxFormCallback.CreationMode = (checkbox == nil)

	return
}

type CheckBoxFormCallback struct {
	checkbox *models.CheckBox

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (checkboxFormCallback *CheckBoxFormCallback) OnSave() {

	log.Println("CheckBoxFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	checkboxFormCallback.playground.formStage.Checkout()

	if checkboxFormCallback.checkbox == nil {
		checkboxFormCallback.checkbox = new(models.CheckBox).Stage(checkboxFormCallback.playground.stageOfInterest)
	}
	checkbox_ := checkboxFormCallback.checkbox
	_ = checkbox_

	// get the formGroup
	formGroup := checkboxFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(checkbox_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(checkbox_.Value), formDiv)
		}
	}

	checkboxFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.CheckBox](
		checkboxFormCallback.playground,
	)
	checkboxFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if checkboxFormCallback.CreationMode {
		checkboxFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewCheckBoxFormCallback(
				nil,
				checkboxFormCallback.playground,
			),
		}).Stage(checkboxFormCallback.playground.formStage)
		checkbox := new(models.CheckBox)
		FillUpForm(checkbox, newFormGroup, checkboxFormCallback.playground)
		checkboxFormCallback.playground.formStage.Commit()
	}

}
func NewFormDivFormCallback(
	formdiv *models.FormDiv,
	playground *Playground,
) (formdivFormCallback *FormDivFormCallback) {
	formdivFormCallback = new(FormDivFormCallback)
	formdivFormCallback.playground = playground
	formdivFormCallback.formdiv = formdiv

	formdivFormCallback.CreationMode = (formdiv == nil)

	return
}

type FormDivFormCallback struct {
	formdiv *models.FormDiv

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (formdivFormCallback *FormDivFormCallback) OnSave() {

	log.Println("FormDivFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formdivFormCallback.playground.formStage.Checkout()

	if formdivFormCallback.formdiv == nil {
		formdivFormCallback.formdiv = new(models.FormDiv).Stage(formdivFormCallback.playground.stageOfInterest)
	}
	formdiv_ := formdivFormCallback.formdiv
	_ = formdiv_

	// get the formGroup
	formGroup := formdivFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formdiv_.Name), formDiv)
		case "FormEditAssocButton":
			FormDivSelectFieldToField(&(formdiv_.FormEditAssocButton), formdivFormCallback.playground.stageOfInterest, formDiv)
		case "FormSortAssocButton":
			FormDivSelectFieldToField(&(formdiv_.FormSortAssocButton), formdivFormCallback.playground.stageOfInterest, formDiv)
		}
	}

	formdivFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormDiv](
		formdivFormCallback.playground,
	)
	formdivFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if formdivFormCallback.CreationMode {
		formdivFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewFormDivFormCallback(
				nil,
				formdivFormCallback.playground,
			),
		}).Stage(formdivFormCallback.playground.formStage)
		formdiv := new(models.FormDiv)
		FillUpForm(formdiv, newFormGroup, formdivFormCallback.playground)
		formdivFormCallback.playground.formStage.Commit()
	}

}
func NewFormEditAssocButtonFormCallback(
	formeditassocbutton *models.FormEditAssocButton,
	playground *Playground,
) (formeditassocbuttonFormCallback *FormEditAssocButtonFormCallback) {
	formeditassocbuttonFormCallback = new(FormEditAssocButtonFormCallback)
	formeditassocbuttonFormCallback.playground = playground
	formeditassocbuttonFormCallback.formeditassocbutton = formeditassocbutton

	formeditassocbuttonFormCallback.CreationMode = (formeditassocbutton == nil)

	return
}

type FormEditAssocButtonFormCallback struct {
	formeditassocbutton *models.FormEditAssocButton

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (formeditassocbuttonFormCallback *FormEditAssocButtonFormCallback) OnSave() {

	log.Println("FormEditAssocButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formeditassocbuttonFormCallback.playground.formStage.Checkout()

	if formeditassocbuttonFormCallback.formeditassocbutton == nil {
		formeditassocbuttonFormCallback.formeditassocbutton = new(models.FormEditAssocButton).Stage(formeditassocbuttonFormCallback.playground.stageOfInterest)
	}
	formeditassocbutton_ := formeditassocbuttonFormCallback.formeditassocbutton
	_ = formeditassocbutton_

	// get the formGroup
	formGroup := formeditassocbuttonFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formeditassocbutton_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formeditassocbutton_.Label), formDiv)
		}
	}

	formeditassocbuttonFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormEditAssocButton](
		formeditassocbuttonFormCallback.playground,
	)
	formeditassocbuttonFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if formeditassocbuttonFormCallback.CreationMode {
		formeditassocbuttonFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewFormEditAssocButtonFormCallback(
				nil,
				formeditassocbuttonFormCallback.playground,
			),
		}).Stage(formeditassocbuttonFormCallback.playground.formStage)
		formeditassocbutton := new(models.FormEditAssocButton)
		FillUpForm(formeditassocbutton, newFormGroup, formeditassocbuttonFormCallback.playground)
		formeditassocbuttonFormCallback.playground.formStage.Commit()
	}

}
func NewFormFieldFormCallback(
	formfield *models.FormField,
	playground *Playground,
) (formfieldFormCallback *FormFieldFormCallback) {
	formfieldFormCallback = new(FormFieldFormCallback)
	formfieldFormCallback.playground = playground
	formfieldFormCallback.formfield = formfield

	formfieldFormCallback.CreationMode = (formfield == nil)

	return
}

type FormFieldFormCallback struct {
	formfield *models.FormField

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (formfieldFormCallback *FormFieldFormCallback) OnSave() {

	log.Println("FormFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldFormCallback.playground.formStage.Checkout()

	if formfieldFormCallback.formfield == nil {
		formfieldFormCallback.formfield = new(models.FormField).Stage(formfieldFormCallback.playground.stageOfInterest)
	}
	formfield_ := formfieldFormCallback.formfield
	_ = formfield_

	// get the formGroup
	formGroup := formfieldFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfield_.Name), formDiv)
		case "InputTypeEnum":
			FormDivEnumStringFieldToField(&(formfield_.InputTypeEnum), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formfield_.Label), formDiv)
		case "Placeholder":
			FormDivBasicFieldToField(&(formfield_.Placeholder), formDiv)
		case "FormFieldString":
			FormDivSelectFieldToField(&(formfield_.FormFieldString), formfieldFormCallback.playground.stageOfInterest, formDiv)
		case "FormFieldFloat64":
			FormDivSelectFieldToField(&(formfield_.FormFieldFloat64), formfieldFormCallback.playground.stageOfInterest, formDiv)
		case "FormFieldInt":
			FormDivSelectFieldToField(&(formfield_.FormFieldInt), formfieldFormCallback.playground.stageOfInterest, formDiv)
		case "FormFieldDate":
			FormDivSelectFieldToField(&(formfield_.FormFieldDate), formfieldFormCallback.playground.stageOfInterest, formDiv)
		case "FormFieldTime":
			FormDivSelectFieldToField(&(formfield_.FormFieldTime), formfieldFormCallback.playground.stageOfInterest, formDiv)
		case "FormFieldDateTime":
			FormDivSelectFieldToField(&(formfield_.FormFieldDateTime), formfieldFormCallback.playground.stageOfInterest, formDiv)
		case "FormFieldSelect":
			FormDivSelectFieldToField(&(formfield_.FormFieldSelect), formfieldFormCallback.playground.stageOfInterest, formDiv)
		case "HasBespokeWidth":
			FormDivBasicFieldToField(&(formfield_.HasBespokeWidth), formDiv)
		case "BespokeWidthPx":
			FormDivBasicFieldToField(&(formfield_.BespokeWidthPx), formDiv)
		}
	}

	formfieldFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormField](
		formfieldFormCallback.playground,
	)
	formfieldFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldFormCallback.CreationMode {
		formfieldFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldFormCallback(
				nil,
				formfieldFormCallback.playground,
			),
		}).Stage(formfieldFormCallback.playground.formStage)
		formfield := new(models.FormField)
		FillUpForm(formfield, newFormGroup, formfieldFormCallback.playground)
		formfieldFormCallback.playground.formStage.Commit()
	}

}
func NewFormFieldDateFormCallback(
	formfielddate *models.FormFieldDate,
	playground *Playground,
) (formfielddateFormCallback *FormFieldDateFormCallback) {
	formfielddateFormCallback = new(FormFieldDateFormCallback)
	formfielddateFormCallback.playground = playground
	formfielddateFormCallback.formfielddate = formfielddate

	formfielddateFormCallback.CreationMode = (formfielddate == nil)

	return
}

type FormFieldDateFormCallback struct {
	formfielddate *models.FormFieldDate

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (formfielddateFormCallback *FormFieldDateFormCallback) OnSave() {

	log.Println("FormFieldDateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfielddateFormCallback.playground.formStage.Checkout()

	if formfielddateFormCallback.formfielddate == nil {
		formfielddateFormCallback.formfielddate = new(models.FormFieldDate).Stage(formfielddateFormCallback.playground.stageOfInterest)
	}
	formfielddate_ := formfielddateFormCallback.formfielddate
	_ = formfielddate_

	// get the formGroup
	formGroup := formfielddateFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfielddate_.Name), formDiv)
		}
	}

	formfielddateFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormFieldDate](
		formfielddateFormCallback.playground,
	)
	formfielddateFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if formfielddateFormCallback.CreationMode {
		formfielddateFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldDateFormCallback(
				nil,
				formfielddateFormCallback.playground,
			),
		}).Stage(formfielddateFormCallback.playground.formStage)
		formfielddate := new(models.FormFieldDate)
		FillUpForm(formfielddate, newFormGroup, formfielddateFormCallback.playground)
		formfielddateFormCallback.playground.formStage.Commit()
	}

}
func NewFormFieldDateTimeFormCallback(
	formfielddatetime *models.FormFieldDateTime,
	playground *Playground,
) (formfielddatetimeFormCallback *FormFieldDateTimeFormCallback) {
	formfielddatetimeFormCallback = new(FormFieldDateTimeFormCallback)
	formfielddatetimeFormCallback.playground = playground
	formfielddatetimeFormCallback.formfielddatetime = formfielddatetime

	formfielddatetimeFormCallback.CreationMode = (formfielddatetime == nil)

	return
}

type FormFieldDateTimeFormCallback struct {
	formfielddatetime *models.FormFieldDateTime

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (formfielddatetimeFormCallback *FormFieldDateTimeFormCallback) OnSave() {

	log.Println("FormFieldDateTimeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfielddatetimeFormCallback.playground.formStage.Checkout()

	if formfielddatetimeFormCallback.formfielddatetime == nil {
		formfielddatetimeFormCallback.formfielddatetime = new(models.FormFieldDateTime).Stage(formfielddatetimeFormCallback.playground.stageOfInterest)
	}
	formfielddatetime_ := formfielddatetimeFormCallback.formfielddatetime
	_ = formfielddatetime_

	// get the formGroup
	formGroup := formfielddatetimeFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfielddatetime_.Name), formDiv)
		}
	}

	formfielddatetimeFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormFieldDateTime](
		formfielddatetimeFormCallback.playground,
	)
	formfielddatetimeFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if formfielddatetimeFormCallback.CreationMode {
		formfielddatetimeFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldDateTimeFormCallback(
				nil,
				formfielddatetimeFormCallback.playground,
			),
		}).Stage(formfielddatetimeFormCallback.playground.formStage)
		formfielddatetime := new(models.FormFieldDateTime)
		FillUpForm(formfielddatetime, newFormGroup, formfielddatetimeFormCallback.playground)
		formfielddatetimeFormCallback.playground.formStage.Commit()
	}

}
func NewFormFieldFloat64FormCallback(
	formfieldfloat64 *models.FormFieldFloat64,
	playground *Playground,
) (formfieldfloat64FormCallback *FormFieldFloat64FormCallback) {
	formfieldfloat64FormCallback = new(FormFieldFloat64FormCallback)
	formfieldfloat64FormCallback.playground = playground
	formfieldfloat64FormCallback.formfieldfloat64 = formfieldfloat64

	formfieldfloat64FormCallback.CreationMode = (formfieldfloat64 == nil)

	return
}

type FormFieldFloat64FormCallback struct {
	formfieldfloat64 *models.FormFieldFloat64

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (formfieldfloat64FormCallback *FormFieldFloat64FormCallback) OnSave() {

	log.Println("FormFieldFloat64FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldfloat64FormCallback.playground.formStage.Checkout()

	if formfieldfloat64FormCallback.formfieldfloat64 == nil {
		formfieldfloat64FormCallback.formfieldfloat64 = new(models.FormFieldFloat64).Stage(formfieldfloat64FormCallback.playground.stageOfInterest)
	}
	formfieldfloat64_ := formfieldfloat64FormCallback.formfieldfloat64
	_ = formfieldfloat64_

	// get the formGroup
	formGroup := formfieldfloat64FormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldfloat64_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfieldfloat64_.Value), formDiv)
		case "HasMinValidator":
			FormDivBasicFieldToField(&(formfieldfloat64_.HasMinValidator), formDiv)
		case "MinValue":
			FormDivBasicFieldToField(&(formfieldfloat64_.MinValue), formDiv)
		case "HasMaxValidator":
			FormDivBasicFieldToField(&(formfieldfloat64_.HasMaxValidator), formDiv)
		case "MaxValue":
			FormDivBasicFieldToField(&(formfieldfloat64_.MaxValue), formDiv)
		}
	}

	formfieldfloat64FormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormFieldFloat64](
		formfieldfloat64FormCallback.playground,
	)
	formfieldfloat64FormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldfloat64FormCallback.CreationMode {
		formfieldfloat64FormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldFloat64FormCallback(
				nil,
				formfieldfloat64FormCallback.playground,
			),
		}).Stage(formfieldfloat64FormCallback.playground.formStage)
		formfieldfloat64 := new(models.FormFieldFloat64)
		FillUpForm(formfieldfloat64, newFormGroup, formfieldfloat64FormCallback.playground)
		formfieldfloat64FormCallback.playground.formStage.Commit()
	}

}
func NewFormFieldIntFormCallback(
	formfieldint *models.FormFieldInt,
	playground *Playground,
) (formfieldintFormCallback *FormFieldIntFormCallback) {
	formfieldintFormCallback = new(FormFieldIntFormCallback)
	formfieldintFormCallback.playground = playground
	formfieldintFormCallback.formfieldint = formfieldint

	formfieldintFormCallback.CreationMode = (formfieldint == nil)

	return
}

type FormFieldIntFormCallback struct {
	formfieldint *models.FormFieldInt

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (formfieldintFormCallback *FormFieldIntFormCallback) OnSave() {

	log.Println("FormFieldIntFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldintFormCallback.playground.formStage.Checkout()

	if formfieldintFormCallback.formfieldint == nil {
		formfieldintFormCallback.formfieldint = new(models.FormFieldInt).Stage(formfieldintFormCallback.playground.stageOfInterest)
	}
	formfieldint_ := formfieldintFormCallback.formfieldint
	_ = formfieldint_

	// get the formGroup
	formGroup := formfieldintFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldint_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfieldint_.Value), formDiv)
		case "HasMinValidator":
			FormDivBasicFieldToField(&(formfieldint_.HasMinValidator), formDiv)
		case "MinValue":
			FormDivBasicFieldToField(&(formfieldint_.MinValue), formDiv)
		case "HasMaxValidator":
			FormDivBasicFieldToField(&(formfieldint_.HasMaxValidator), formDiv)
		case "MaxValue":
			FormDivBasicFieldToField(&(formfieldint_.MaxValue), formDiv)
		}
	}

	formfieldintFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormFieldInt](
		formfieldintFormCallback.playground,
	)
	formfieldintFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldintFormCallback.CreationMode {
		formfieldintFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldIntFormCallback(
				nil,
				formfieldintFormCallback.playground,
			),
		}).Stage(formfieldintFormCallback.playground.formStage)
		formfieldint := new(models.FormFieldInt)
		FillUpForm(formfieldint, newFormGroup, formfieldintFormCallback.playground)
		formfieldintFormCallback.playground.formStage.Commit()
	}

}
func NewFormFieldSelectFormCallback(
	formfieldselect *models.FormFieldSelect,
	playground *Playground,
) (formfieldselectFormCallback *FormFieldSelectFormCallback) {
	formfieldselectFormCallback = new(FormFieldSelectFormCallback)
	formfieldselectFormCallback.playground = playground
	formfieldselectFormCallback.formfieldselect = formfieldselect

	formfieldselectFormCallback.CreationMode = (formfieldselect == nil)

	return
}

type FormFieldSelectFormCallback struct {
	formfieldselect *models.FormFieldSelect

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (formfieldselectFormCallback *FormFieldSelectFormCallback) OnSave() {

	log.Println("FormFieldSelectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldselectFormCallback.playground.formStage.Checkout()

	if formfieldselectFormCallback.formfieldselect == nil {
		formfieldselectFormCallback.formfieldselect = new(models.FormFieldSelect).Stage(formfieldselectFormCallback.playground.stageOfInterest)
	}
	formfieldselect_ := formfieldselectFormCallback.formfieldselect
	_ = formfieldselect_

	// get the formGroup
	formGroup := formfieldselectFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldselect_.Name), formDiv)
		case "Value":
			FormDivSelectFieldToField(&(formfieldselect_.Value), formfieldselectFormCallback.playground.stageOfInterest, formDiv)
		case "CanBeEmpty":
			FormDivBasicFieldToField(&(formfieldselect_.CanBeEmpty), formDiv)
		}
	}

	formfieldselectFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormFieldSelect](
		formfieldselectFormCallback.playground,
	)
	formfieldselectFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldselectFormCallback.CreationMode {
		formfieldselectFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldSelectFormCallback(
				nil,
				formfieldselectFormCallback.playground,
			),
		}).Stage(formfieldselectFormCallback.playground.formStage)
		formfieldselect := new(models.FormFieldSelect)
		FillUpForm(formfieldselect, newFormGroup, formfieldselectFormCallback.playground)
		formfieldselectFormCallback.playground.formStage.Commit()
	}

}
func NewFormFieldStringFormCallback(
	formfieldstring *models.FormFieldString,
	playground *Playground,
) (formfieldstringFormCallback *FormFieldStringFormCallback) {
	formfieldstringFormCallback = new(FormFieldStringFormCallback)
	formfieldstringFormCallback.playground = playground
	formfieldstringFormCallback.formfieldstring = formfieldstring

	formfieldstringFormCallback.CreationMode = (formfieldstring == nil)

	return
}

type FormFieldStringFormCallback struct {
	formfieldstring *models.FormFieldString

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (formfieldstringFormCallback *FormFieldStringFormCallback) OnSave() {

	log.Println("FormFieldStringFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldstringFormCallback.playground.formStage.Checkout()

	if formfieldstringFormCallback.formfieldstring == nil {
		formfieldstringFormCallback.formfieldstring = new(models.FormFieldString).Stage(formfieldstringFormCallback.playground.stageOfInterest)
	}
	formfieldstring_ := formfieldstringFormCallback.formfieldstring
	_ = formfieldstring_

	// get the formGroup
	formGroup := formfieldstringFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldstring_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfieldstring_.Value), formDiv)
		}
	}

	formfieldstringFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormFieldString](
		formfieldstringFormCallback.playground,
	)
	formfieldstringFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldstringFormCallback.CreationMode {
		formfieldstringFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldStringFormCallback(
				nil,
				formfieldstringFormCallback.playground,
			),
		}).Stage(formfieldstringFormCallback.playground.formStage)
		formfieldstring := new(models.FormFieldString)
		FillUpForm(formfieldstring, newFormGroup, formfieldstringFormCallback.playground)
		formfieldstringFormCallback.playground.formStage.Commit()
	}

}
func NewFormFieldTimeFormCallback(
	formfieldtime *models.FormFieldTime,
	playground *Playground,
) (formfieldtimeFormCallback *FormFieldTimeFormCallback) {
	formfieldtimeFormCallback = new(FormFieldTimeFormCallback)
	formfieldtimeFormCallback.playground = playground
	formfieldtimeFormCallback.formfieldtime = formfieldtime

	formfieldtimeFormCallback.CreationMode = (formfieldtime == nil)

	return
}

type FormFieldTimeFormCallback struct {
	formfieldtime *models.FormFieldTime

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (formfieldtimeFormCallback *FormFieldTimeFormCallback) OnSave() {

	log.Println("FormFieldTimeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldtimeFormCallback.playground.formStage.Checkout()

	if formfieldtimeFormCallback.formfieldtime == nil {
		formfieldtimeFormCallback.formfieldtime = new(models.FormFieldTime).Stage(formfieldtimeFormCallback.playground.stageOfInterest)
	}
	formfieldtime_ := formfieldtimeFormCallback.formfieldtime
	_ = formfieldtime_

	// get the formGroup
	formGroup := formfieldtimeFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldtime_.Name), formDiv)
		case "Step":
			FormDivBasicFieldToField(&(formfieldtime_.Step), formDiv)
		}
	}

	formfieldtimeFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormFieldTime](
		formfieldtimeFormCallback.playground,
	)
	formfieldtimeFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldtimeFormCallback.CreationMode {
		formfieldtimeFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewFormFieldTimeFormCallback(
				nil,
				formfieldtimeFormCallback.playground,
			),
		}).Stage(formfieldtimeFormCallback.playground.formStage)
		formfieldtime := new(models.FormFieldTime)
		FillUpForm(formfieldtime, newFormGroup, formfieldtimeFormCallback.playground)
		formfieldtimeFormCallback.playground.formStage.Commit()
	}

}
func NewFormGroupFormCallback(
	formgroup *models.FormGroup,
	playground *Playground,
) (formgroupFormCallback *FormGroupFormCallback) {
	formgroupFormCallback = new(FormGroupFormCallback)
	formgroupFormCallback.playground = playground
	formgroupFormCallback.formgroup = formgroup

	formgroupFormCallback.CreationMode = (formgroup == nil)

	return
}

type FormGroupFormCallback struct {
	formgroup *models.FormGroup

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (formgroupFormCallback *FormGroupFormCallback) OnSave() {

	log.Println("FormGroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formgroupFormCallback.playground.formStage.Checkout()

	if formgroupFormCallback.formgroup == nil {
		formgroupFormCallback.formgroup = new(models.FormGroup).Stage(formgroupFormCallback.playground.stageOfInterest)
	}
	formgroup_ := formgroupFormCallback.formgroup
	_ = formgroup_

	// get the formGroup
	formGroup := formgroupFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formgroup_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formgroup_.Label), formDiv)
		}
	}

	formgroupFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormGroup](
		formgroupFormCallback.playground,
	)
	formgroupFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if formgroupFormCallback.CreationMode {
		formgroupFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewFormGroupFormCallback(
				nil,
				formgroupFormCallback.playground,
			),
		}).Stage(formgroupFormCallback.playground.formStage)
		formgroup := new(models.FormGroup)
		FillUpForm(formgroup, newFormGroup, formgroupFormCallback.playground)
		formgroupFormCallback.playground.formStage.Commit()
	}

}
func NewFormSortAssocButtonFormCallback(
	formsortassocbutton *models.FormSortAssocButton,
	playground *Playground,
) (formsortassocbuttonFormCallback *FormSortAssocButtonFormCallback) {
	formsortassocbuttonFormCallback = new(FormSortAssocButtonFormCallback)
	formsortassocbuttonFormCallback.playground = playground
	formsortassocbuttonFormCallback.formsortassocbutton = formsortassocbutton

	formsortassocbuttonFormCallback.CreationMode = (formsortassocbutton == nil)

	return
}

type FormSortAssocButtonFormCallback struct {
	formsortassocbutton *models.FormSortAssocButton

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (formsortassocbuttonFormCallback *FormSortAssocButtonFormCallback) OnSave() {

	log.Println("FormSortAssocButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formsortassocbuttonFormCallback.playground.formStage.Checkout()

	if formsortassocbuttonFormCallback.formsortassocbutton == nil {
		formsortassocbuttonFormCallback.formsortassocbutton = new(models.FormSortAssocButton).Stage(formsortassocbuttonFormCallback.playground.stageOfInterest)
	}
	formsortassocbutton_ := formsortassocbuttonFormCallback.formsortassocbutton
	_ = formsortassocbutton_

	// get the formGroup
	formGroup := formsortassocbuttonFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formsortassocbutton_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formsortassocbutton_.Label), formDiv)
		}
	}

	formsortassocbuttonFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.FormSortAssocButton](
		formsortassocbuttonFormCallback.playground,
	)
	formsortassocbuttonFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if formsortassocbuttonFormCallback.CreationMode {
		formsortassocbuttonFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewFormSortAssocButtonFormCallback(
				nil,
				formsortassocbuttonFormCallback.playground,
			),
		}).Stage(formsortassocbuttonFormCallback.playground.formStage)
		formsortassocbutton := new(models.FormSortAssocButton)
		FillUpForm(formsortassocbutton, newFormGroup, formsortassocbuttonFormCallback.playground)
		formsortassocbuttonFormCallback.playground.formStage.Commit()
	}

}
func NewOptionFormCallback(
	option *models.Option,
	playground *Playground,
) (optionFormCallback *OptionFormCallback) {
	optionFormCallback = new(OptionFormCallback)
	optionFormCallback.playground = playground
	optionFormCallback.option = option

	optionFormCallback.CreationMode = (option == nil)

	return
}

type OptionFormCallback struct {
	option *models.Option

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (optionFormCallback *OptionFormCallback) OnSave() {

	log.Println("OptionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	optionFormCallback.playground.formStage.Checkout()

	if optionFormCallback.option == nil {
		optionFormCallback.option = new(models.Option).Stage(optionFormCallback.playground.stageOfInterest)
	}
	option_ := optionFormCallback.option
	_ = option_

	// get the formGroup
	formGroup := optionFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(option_.Name), formDiv)
		}
	}

	optionFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Option](
		optionFormCallback.playground,
	)
	optionFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if optionFormCallback.CreationMode {
		optionFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: NewOptionFormCallback(
				nil,
				optionFormCallback.playground,
			),
		}).Stage(optionFormCallback.playground.formStage)
		option := new(models.Option)
		FillUpForm(option, newFormGroup, optionFormCallback.playground)
		optionFormCallback.playground.formStage.Commit()
	}

}
