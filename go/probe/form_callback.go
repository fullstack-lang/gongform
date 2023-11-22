// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongform/go/models"
	"github.com/fullstack-lang/gongform/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__CheckBoxFormCallback(
	checkbox *models.CheckBox,
	probe *Probe,
) (checkboxFormCallback *CheckBoxFormCallback) {
	checkboxFormCallback = new(CheckBoxFormCallback)
	checkboxFormCallback.probe = probe
	checkboxFormCallback.checkbox = checkbox

	checkboxFormCallback.CreationMode = (checkbox == nil)

	return
}

type CheckBoxFormCallback struct {
	checkbox *models.CheckBox

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (checkboxFormCallback *CheckBoxFormCallback) OnSave() {

	log.Println("CheckBoxFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	checkboxFormCallback.probe.formStage.Checkout()

	if checkboxFormCallback.checkbox == nil {
		checkboxFormCallback.checkbox = new(models.CheckBox).Stage(checkboxFormCallback.probe.stageOfInterest)
	}
	checkbox_ := checkboxFormCallback.checkbox
	_ = checkbox_

	// get the formGroup
	formGroup := checkboxFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(checkbox_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(checkbox_.Value), formDiv)
		case "FormDiv:CheckBoxs":
			// we need to retrieve the field owner before the change
			var pastFormDivOwner *models.FormDiv
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormDiv"
			rf.Fieldname = "CheckBoxs"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				checkboxFormCallback.probe.stageOfInterest,
				checkboxFormCallback.probe.backRepoOfInterest,
				checkbox_,
				&rf)

			if reverseFieldOwner != nil {
				pastFormDivOwner = reverseFieldOwner.(*models.FormDiv)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastFormDivOwner != nil {
					idx := slices.Index(pastFormDivOwner.CheckBoxs, checkbox_)
					pastFormDivOwner.CheckBoxs = slices.Delete(pastFormDivOwner.CheckBoxs, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _formdiv := range *models.GetGongstructInstancesSet[models.FormDiv](checkboxFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _formdiv.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newFormDivOwner := _formdiv // we have a match
						if pastFormDivOwner != nil {
							if newFormDivOwner != pastFormDivOwner {
								idx := slices.Index(pastFormDivOwner.CheckBoxs, checkbox_)
								pastFormDivOwner.CheckBoxs = slices.Delete(pastFormDivOwner.CheckBoxs, idx, idx+1)
								newFormDivOwner.CheckBoxs = append(newFormDivOwner.CheckBoxs, checkbox_)
							}
						} else {
							newFormDivOwner.CheckBoxs = append(newFormDivOwner.CheckBoxs, checkbox_)
						}
					}
				}
			}
		}
	}

	checkboxFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.CheckBox](
		checkboxFormCallback.probe,
	)
	checkboxFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if checkboxFormCallback.CreationMode {
		checkboxFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__CheckBoxFormCallback(
				nil,
				checkboxFormCallback.probe,
			),
		}).Stage(checkboxFormCallback.probe.formStage)
		checkbox := new(models.CheckBox)
		FillUpForm(checkbox, newFormGroup, checkboxFormCallback.probe)
		checkboxFormCallback.probe.formStage.Commit()
	}

	fillUpTree(checkboxFormCallback.probe)
}
func __gong__New__FormDivFormCallback(
	formdiv *models.FormDiv,
	probe *Probe,
) (formdivFormCallback *FormDivFormCallback) {
	formdivFormCallback = new(FormDivFormCallback)
	formdivFormCallback.probe = probe
	formdivFormCallback.formdiv = formdiv

	formdivFormCallback.CreationMode = (formdiv == nil)

	return
}

type FormDivFormCallback struct {
	formdiv *models.FormDiv

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (formdivFormCallback *FormDivFormCallback) OnSave() {

	log.Println("FormDivFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formdivFormCallback.probe.formStage.Checkout()

	if formdivFormCallback.formdiv == nil {
		formdivFormCallback.formdiv = new(models.FormDiv).Stage(formdivFormCallback.probe.stageOfInterest)
	}
	formdiv_ := formdivFormCallback.formdiv
	_ = formdiv_

	// get the formGroup
	formGroup := formdivFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formdiv_.Name), formDiv)
		case "FormEditAssocButton":
			FormDivSelectFieldToField(&(formdiv_.FormEditAssocButton), formdivFormCallback.probe.stageOfInterest, formDiv)
		case "FormSortAssocButton":
			FormDivSelectFieldToField(&(formdiv_.FormSortAssocButton), formdivFormCallback.probe.stageOfInterest, formDiv)
		case "FormGroup:FormDivs":
			// we need to retrieve the field owner before the change
			var pastFormGroupOwner *models.FormGroup
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormGroup"
			rf.Fieldname = "FormDivs"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				formdivFormCallback.probe.stageOfInterest,
				formdivFormCallback.probe.backRepoOfInterest,
				formdiv_,
				&rf)

			if reverseFieldOwner != nil {
				pastFormGroupOwner = reverseFieldOwner.(*models.FormGroup)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastFormGroupOwner != nil {
					idx := slices.Index(pastFormGroupOwner.FormDivs, formdiv_)
					pastFormGroupOwner.FormDivs = slices.Delete(pastFormGroupOwner.FormDivs, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _formgroup := range *models.GetGongstructInstancesSet[models.FormGroup](formdivFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _formgroup.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newFormGroupOwner := _formgroup // we have a match
						if pastFormGroupOwner != nil {
							if newFormGroupOwner != pastFormGroupOwner {
								idx := slices.Index(pastFormGroupOwner.FormDivs, formdiv_)
								pastFormGroupOwner.FormDivs = slices.Delete(pastFormGroupOwner.FormDivs, idx, idx+1)
								newFormGroupOwner.FormDivs = append(newFormGroupOwner.FormDivs, formdiv_)
							}
						} else {
							newFormGroupOwner.FormDivs = append(newFormGroupOwner.FormDivs, formdiv_)
						}
					}
				}
			}
		}
	}

	formdivFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.FormDiv](
		formdivFormCallback.probe,
	)
	formdivFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formdivFormCallback.CreationMode {
		formdivFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormDivFormCallback(
				nil,
				formdivFormCallback.probe,
			),
		}).Stage(formdivFormCallback.probe.formStage)
		formdiv := new(models.FormDiv)
		FillUpForm(formdiv, newFormGroup, formdivFormCallback.probe)
		formdivFormCallback.probe.formStage.Commit()
	}

	fillUpTree(formdivFormCallback.probe)
}
func __gong__New__FormEditAssocButtonFormCallback(
	formeditassocbutton *models.FormEditAssocButton,
	probe *Probe,
) (formeditassocbuttonFormCallback *FormEditAssocButtonFormCallback) {
	formeditassocbuttonFormCallback = new(FormEditAssocButtonFormCallback)
	formeditassocbuttonFormCallback.probe = probe
	formeditassocbuttonFormCallback.formeditassocbutton = formeditassocbutton

	formeditassocbuttonFormCallback.CreationMode = (formeditassocbutton == nil)

	return
}

type FormEditAssocButtonFormCallback struct {
	formeditassocbutton *models.FormEditAssocButton

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (formeditassocbuttonFormCallback *FormEditAssocButtonFormCallback) OnSave() {

	log.Println("FormEditAssocButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formeditassocbuttonFormCallback.probe.formStage.Checkout()

	if formeditassocbuttonFormCallback.formeditassocbutton == nil {
		formeditassocbuttonFormCallback.formeditassocbutton = new(models.FormEditAssocButton).Stage(formeditassocbuttonFormCallback.probe.stageOfInterest)
	}
	formeditassocbutton_ := formeditassocbuttonFormCallback.formeditassocbutton
	_ = formeditassocbutton_

	// get the formGroup
	formGroup := formeditassocbuttonFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formeditassocbutton_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formeditassocbutton_.Label), formDiv)
		}
	}

	formeditassocbuttonFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.FormEditAssocButton](
		formeditassocbuttonFormCallback.probe,
	)
	formeditassocbuttonFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formeditassocbuttonFormCallback.CreationMode {
		formeditassocbuttonFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormEditAssocButtonFormCallback(
				nil,
				formeditassocbuttonFormCallback.probe,
			),
		}).Stage(formeditassocbuttonFormCallback.probe.formStage)
		formeditassocbutton := new(models.FormEditAssocButton)
		FillUpForm(formeditassocbutton, newFormGroup, formeditassocbuttonFormCallback.probe)
		formeditassocbuttonFormCallback.probe.formStage.Commit()
	}

	fillUpTree(formeditassocbuttonFormCallback.probe)
}
func __gong__New__FormFieldFormCallback(
	formfield *models.FormField,
	probe *Probe,
) (formfieldFormCallback *FormFieldFormCallback) {
	formfieldFormCallback = new(FormFieldFormCallback)
	formfieldFormCallback.probe = probe
	formfieldFormCallback.formfield = formfield

	formfieldFormCallback.CreationMode = (formfield == nil)

	return
}

type FormFieldFormCallback struct {
	formfield *models.FormField

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (formfieldFormCallback *FormFieldFormCallback) OnSave() {

	log.Println("FormFieldFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldFormCallback.probe.formStage.Checkout()

	if formfieldFormCallback.formfield == nil {
		formfieldFormCallback.formfield = new(models.FormField).Stage(formfieldFormCallback.probe.stageOfInterest)
	}
	formfield_ := formfieldFormCallback.formfield
	_ = formfield_

	// get the formGroup
	formGroup := formfieldFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
			FormDivSelectFieldToField(&(formfield_.FormFieldString), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldFloat64":
			FormDivSelectFieldToField(&(formfield_.FormFieldFloat64), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldInt":
			FormDivSelectFieldToField(&(formfield_.FormFieldInt), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldDate":
			FormDivSelectFieldToField(&(formfield_.FormFieldDate), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldTime":
			FormDivSelectFieldToField(&(formfield_.FormFieldTime), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldDateTime":
			FormDivSelectFieldToField(&(formfield_.FormFieldDateTime), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "FormFieldSelect":
			FormDivSelectFieldToField(&(formfield_.FormFieldSelect), formfieldFormCallback.probe.stageOfInterest, formDiv)
		case "HasBespokeWidth":
			FormDivBasicFieldToField(&(formfield_.HasBespokeWidth), formDiv)
		case "BespokeWidthPx":
			FormDivBasicFieldToField(&(formfield_.BespokeWidthPx), formDiv)
		case "FormDiv:FormFields":
			// we need to retrieve the field owner before the change
			var pastFormDivOwner *models.FormDiv
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormDiv"
			rf.Fieldname = "FormFields"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				formfieldFormCallback.probe.stageOfInterest,
				formfieldFormCallback.probe.backRepoOfInterest,
				formfield_,
				&rf)

			if reverseFieldOwner != nil {
				pastFormDivOwner = reverseFieldOwner.(*models.FormDiv)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastFormDivOwner != nil {
					idx := slices.Index(pastFormDivOwner.FormFields, formfield_)
					pastFormDivOwner.FormFields = slices.Delete(pastFormDivOwner.FormFields, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _formdiv := range *models.GetGongstructInstancesSet[models.FormDiv](formfieldFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _formdiv.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newFormDivOwner := _formdiv // we have a match
						if pastFormDivOwner != nil {
							if newFormDivOwner != pastFormDivOwner {
								idx := slices.Index(pastFormDivOwner.FormFields, formfield_)
								pastFormDivOwner.FormFields = slices.Delete(pastFormDivOwner.FormFields, idx, idx+1)
								newFormDivOwner.FormFields = append(newFormDivOwner.FormFields, formfield_)
							}
						} else {
							newFormDivOwner.FormFields = append(newFormDivOwner.FormFields, formfield_)
						}
					}
				}
			}
		}
	}

	formfieldFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.FormField](
		formfieldFormCallback.probe,
	)
	formfieldFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldFormCallback.CreationMode {
		formfieldFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormFieldFormCallback(
				nil,
				formfieldFormCallback.probe,
			),
		}).Stage(formfieldFormCallback.probe.formStage)
		formfield := new(models.FormField)
		FillUpForm(formfield, newFormGroup, formfieldFormCallback.probe)
		formfieldFormCallback.probe.formStage.Commit()
	}

	fillUpTree(formfieldFormCallback.probe)
}
func __gong__New__FormFieldDateFormCallback(
	formfielddate *models.FormFieldDate,
	probe *Probe,
) (formfielddateFormCallback *FormFieldDateFormCallback) {
	formfielddateFormCallback = new(FormFieldDateFormCallback)
	formfielddateFormCallback.probe = probe
	formfielddateFormCallback.formfielddate = formfielddate

	formfielddateFormCallback.CreationMode = (formfielddate == nil)

	return
}

type FormFieldDateFormCallback struct {
	formfielddate *models.FormFieldDate

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (formfielddateFormCallback *FormFieldDateFormCallback) OnSave() {

	log.Println("FormFieldDateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfielddateFormCallback.probe.formStage.Checkout()

	if formfielddateFormCallback.formfielddate == nil {
		formfielddateFormCallback.formfielddate = new(models.FormFieldDate).Stage(formfielddateFormCallback.probe.stageOfInterest)
	}
	formfielddate_ := formfielddateFormCallback.formfielddate
	_ = formfielddate_

	// get the formGroup
	formGroup := formfielddateFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfielddate_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfielddate_.Value), formDiv)
		}
	}

	formfielddateFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.FormFieldDate](
		formfielddateFormCallback.probe,
	)
	formfielddateFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formfielddateFormCallback.CreationMode {
		formfielddateFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormFieldDateFormCallback(
				nil,
				formfielddateFormCallback.probe,
			),
		}).Stage(formfielddateFormCallback.probe.formStage)
		formfielddate := new(models.FormFieldDate)
		FillUpForm(formfielddate, newFormGroup, formfielddateFormCallback.probe)
		formfielddateFormCallback.probe.formStage.Commit()
	}

	fillUpTree(formfielddateFormCallback.probe)
}
func __gong__New__FormFieldDateTimeFormCallback(
	formfielddatetime *models.FormFieldDateTime,
	probe *Probe,
) (formfielddatetimeFormCallback *FormFieldDateTimeFormCallback) {
	formfielddatetimeFormCallback = new(FormFieldDateTimeFormCallback)
	formfielddatetimeFormCallback.probe = probe
	formfielddatetimeFormCallback.formfielddatetime = formfielddatetime

	formfielddatetimeFormCallback.CreationMode = (formfielddatetime == nil)

	return
}

type FormFieldDateTimeFormCallback struct {
	formfielddatetime *models.FormFieldDateTime

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (formfielddatetimeFormCallback *FormFieldDateTimeFormCallback) OnSave() {

	log.Println("FormFieldDateTimeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfielddatetimeFormCallback.probe.formStage.Checkout()

	if formfielddatetimeFormCallback.formfielddatetime == nil {
		formfielddatetimeFormCallback.formfielddatetime = new(models.FormFieldDateTime).Stage(formfielddatetimeFormCallback.probe.stageOfInterest)
	}
	formfielddatetime_ := formfielddatetimeFormCallback.formfielddatetime
	_ = formfielddatetime_

	// get the formGroup
	formGroup := formfielddatetimeFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfielddatetime_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfielddatetime_.Value), formDiv)
		}
	}

	formfielddatetimeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.FormFieldDateTime](
		formfielddatetimeFormCallback.probe,
	)
	formfielddatetimeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formfielddatetimeFormCallback.CreationMode {
		formfielddatetimeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormFieldDateTimeFormCallback(
				nil,
				formfielddatetimeFormCallback.probe,
			),
		}).Stage(formfielddatetimeFormCallback.probe.formStage)
		formfielddatetime := new(models.FormFieldDateTime)
		FillUpForm(formfielddatetime, newFormGroup, formfielddatetimeFormCallback.probe)
		formfielddatetimeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(formfielddatetimeFormCallback.probe)
}
func __gong__New__FormFieldFloat64FormCallback(
	formfieldfloat64 *models.FormFieldFloat64,
	probe *Probe,
) (formfieldfloat64FormCallback *FormFieldFloat64FormCallback) {
	formfieldfloat64FormCallback = new(FormFieldFloat64FormCallback)
	formfieldfloat64FormCallback.probe = probe
	formfieldfloat64FormCallback.formfieldfloat64 = formfieldfloat64

	formfieldfloat64FormCallback.CreationMode = (formfieldfloat64 == nil)

	return
}

type FormFieldFloat64FormCallback struct {
	formfieldfloat64 *models.FormFieldFloat64

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (formfieldfloat64FormCallback *FormFieldFloat64FormCallback) OnSave() {

	log.Println("FormFieldFloat64FormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldfloat64FormCallback.probe.formStage.Checkout()

	if formfieldfloat64FormCallback.formfieldfloat64 == nil {
		formfieldfloat64FormCallback.formfieldfloat64 = new(models.FormFieldFloat64).Stage(formfieldfloat64FormCallback.probe.stageOfInterest)
	}
	formfieldfloat64_ := formfieldfloat64FormCallback.formfieldfloat64
	_ = formfieldfloat64_

	// get the formGroup
	formGroup := formfieldfloat64FormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	formfieldfloat64FormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.FormFieldFloat64](
		formfieldfloat64FormCallback.probe,
	)
	formfieldfloat64FormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldfloat64FormCallback.CreationMode {
		formfieldfloat64FormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormFieldFloat64FormCallback(
				nil,
				formfieldfloat64FormCallback.probe,
			),
		}).Stage(formfieldfloat64FormCallback.probe.formStage)
		formfieldfloat64 := new(models.FormFieldFloat64)
		FillUpForm(formfieldfloat64, newFormGroup, formfieldfloat64FormCallback.probe)
		formfieldfloat64FormCallback.probe.formStage.Commit()
	}

	fillUpTree(formfieldfloat64FormCallback.probe)
}
func __gong__New__FormFieldIntFormCallback(
	formfieldint *models.FormFieldInt,
	probe *Probe,
) (formfieldintFormCallback *FormFieldIntFormCallback) {
	formfieldintFormCallback = new(FormFieldIntFormCallback)
	formfieldintFormCallback.probe = probe
	formfieldintFormCallback.formfieldint = formfieldint

	formfieldintFormCallback.CreationMode = (formfieldint == nil)

	return
}

type FormFieldIntFormCallback struct {
	formfieldint *models.FormFieldInt

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (formfieldintFormCallback *FormFieldIntFormCallback) OnSave() {

	log.Println("FormFieldIntFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldintFormCallback.probe.formStage.Checkout()

	if formfieldintFormCallback.formfieldint == nil {
		formfieldintFormCallback.formfieldint = new(models.FormFieldInt).Stage(formfieldintFormCallback.probe.stageOfInterest)
	}
	formfieldint_ := formfieldintFormCallback.formfieldint
	_ = formfieldint_

	// get the formGroup
	formGroup := formfieldintFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	formfieldintFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.FormFieldInt](
		formfieldintFormCallback.probe,
	)
	formfieldintFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldintFormCallback.CreationMode {
		formfieldintFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormFieldIntFormCallback(
				nil,
				formfieldintFormCallback.probe,
			),
		}).Stage(formfieldintFormCallback.probe.formStage)
		formfieldint := new(models.FormFieldInt)
		FillUpForm(formfieldint, newFormGroup, formfieldintFormCallback.probe)
		formfieldintFormCallback.probe.formStage.Commit()
	}

	fillUpTree(formfieldintFormCallback.probe)
}
func __gong__New__FormFieldSelectFormCallback(
	formfieldselect *models.FormFieldSelect,
	probe *Probe,
) (formfieldselectFormCallback *FormFieldSelectFormCallback) {
	formfieldselectFormCallback = new(FormFieldSelectFormCallback)
	formfieldselectFormCallback.probe = probe
	formfieldselectFormCallback.formfieldselect = formfieldselect

	formfieldselectFormCallback.CreationMode = (formfieldselect == nil)

	return
}

type FormFieldSelectFormCallback struct {
	formfieldselect *models.FormFieldSelect

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (formfieldselectFormCallback *FormFieldSelectFormCallback) OnSave() {

	log.Println("FormFieldSelectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldselectFormCallback.probe.formStage.Checkout()

	if formfieldselectFormCallback.formfieldselect == nil {
		formfieldselectFormCallback.formfieldselect = new(models.FormFieldSelect).Stage(formfieldselectFormCallback.probe.stageOfInterest)
	}
	formfieldselect_ := formfieldselectFormCallback.formfieldselect
	_ = formfieldselect_

	// get the formGroup
	formGroup := formfieldselectFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldselect_.Name), formDiv)
		case "Value":
			FormDivSelectFieldToField(&(formfieldselect_.Value), formfieldselectFormCallback.probe.stageOfInterest, formDiv)
		case "CanBeEmpty":
			FormDivBasicFieldToField(&(formfieldselect_.CanBeEmpty), formDiv)
		}
	}

	formfieldselectFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.FormFieldSelect](
		formfieldselectFormCallback.probe,
	)
	formfieldselectFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldselectFormCallback.CreationMode {
		formfieldselectFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormFieldSelectFormCallback(
				nil,
				formfieldselectFormCallback.probe,
			),
		}).Stage(formfieldselectFormCallback.probe.formStage)
		formfieldselect := new(models.FormFieldSelect)
		FillUpForm(formfieldselect, newFormGroup, formfieldselectFormCallback.probe)
		formfieldselectFormCallback.probe.formStage.Commit()
	}

	fillUpTree(formfieldselectFormCallback.probe)
}
func __gong__New__FormFieldStringFormCallback(
	formfieldstring *models.FormFieldString,
	probe *Probe,
) (formfieldstringFormCallback *FormFieldStringFormCallback) {
	formfieldstringFormCallback = new(FormFieldStringFormCallback)
	formfieldstringFormCallback.probe = probe
	formfieldstringFormCallback.formfieldstring = formfieldstring

	formfieldstringFormCallback.CreationMode = (formfieldstring == nil)

	return
}

type FormFieldStringFormCallback struct {
	formfieldstring *models.FormFieldString

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (formfieldstringFormCallback *FormFieldStringFormCallback) OnSave() {

	log.Println("FormFieldStringFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldstringFormCallback.probe.formStage.Checkout()

	if formfieldstringFormCallback.formfieldstring == nil {
		formfieldstringFormCallback.formfieldstring = new(models.FormFieldString).Stage(formfieldstringFormCallback.probe.stageOfInterest)
	}
	formfieldstring_ := formfieldstringFormCallback.formfieldstring
	_ = formfieldstring_

	// get the formGroup
	formGroup := formfieldstringFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldstring_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfieldstring_.Value), formDiv)
		case "IsTextArea":
			FormDivBasicFieldToField(&(formfieldstring_.IsTextArea), formDiv)
		}
	}

	formfieldstringFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.FormFieldString](
		formfieldstringFormCallback.probe,
	)
	formfieldstringFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldstringFormCallback.CreationMode {
		formfieldstringFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormFieldStringFormCallback(
				nil,
				formfieldstringFormCallback.probe,
			),
		}).Stage(formfieldstringFormCallback.probe.formStage)
		formfieldstring := new(models.FormFieldString)
		FillUpForm(formfieldstring, newFormGroup, formfieldstringFormCallback.probe)
		formfieldstringFormCallback.probe.formStage.Commit()
	}

	fillUpTree(formfieldstringFormCallback.probe)
}
func __gong__New__FormFieldTimeFormCallback(
	formfieldtime *models.FormFieldTime,
	probe *Probe,
) (formfieldtimeFormCallback *FormFieldTimeFormCallback) {
	formfieldtimeFormCallback = new(FormFieldTimeFormCallback)
	formfieldtimeFormCallback.probe = probe
	formfieldtimeFormCallback.formfieldtime = formfieldtime

	formfieldtimeFormCallback.CreationMode = (formfieldtime == nil)

	return
}

type FormFieldTimeFormCallback struct {
	formfieldtime *models.FormFieldTime

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (formfieldtimeFormCallback *FormFieldTimeFormCallback) OnSave() {

	log.Println("FormFieldTimeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formfieldtimeFormCallback.probe.formStage.Checkout()

	if formfieldtimeFormCallback.formfieldtime == nil {
		formfieldtimeFormCallback.formfieldtime = new(models.FormFieldTime).Stage(formfieldtimeFormCallback.probe.stageOfInterest)
	}
	formfieldtime_ := formfieldtimeFormCallback.formfieldtime
	_ = formfieldtime_

	// get the formGroup
	formGroup := formfieldtimeFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formfieldtime_.Name), formDiv)
		case "Value":
			FormDivBasicFieldToField(&(formfieldtime_.Value), formDiv)
		case "Step":
			FormDivBasicFieldToField(&(formfieldtime_.Step), formDiv)
		}
	}

	formfieldtimeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.FormFieldTime](
		formfieldtimeFormCallback.probe,
	)
	formfieldtimeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formfieldtimeFormCallback.CreationMode {
		formfieldtimeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormFieldTimeFormCallback(
				nil,
				formfieldtimeFormCallback.probe,
			),
		}).Stage(formfieldtimeFormCallback.probe.formStage)
		formfieldtime := new(models.FormFieldTime)
		FillUpForm(formfieldtime, newFormGroup, formfieldtimeFormCallback.probe)
		formfieldtimeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(formfieldtimeFormCallback.probe)
}
func __gong__New__FormGroupFormCallback(
	formgroup *models.FormGroup,
	probe *Probe,
) (formgroupFormCallback *FormGroupFormCallback) {
	formgroupFormCallback = new(FormGroupFormCallback)
	formgroupFormCallback.probe = probe
	formgroupFormCallback.formgroup = formgroup

	formgroupFormCallback.CreationMode = (formgroup == nil)

	return
}

type FormGroupFormCallback struct {
	formgroup *models.FormGroup

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (formgroupFormCallback *FormGroupFormCallback) OnSave() {

	log.Println("FormGroupFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formgroupFormCallback.probe.formStage.Checkout()

	if formgroupFormCallback.formgroup == nil {
		formgroupFormCallback.formgroup = new(models.FormGroup).Stage(formgroupFormCallback.probe.stageOfInterest)
	}
	formgroup_ := formgroupFormCallback.formgroup
	_ = formgroup_

	// get the formGroup
	formGroup := formgroupFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formgroup_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formgroup_.Label), formDiv)
		}
	}

	formgroupFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.FormGroup](
		formgroupFormCallback.probe,
	)
	formgroupFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formgroupFormCallback.CreationMode {
		formgroupFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormGroupFormCallback(
				nil,
				formgroupFormCallback.probe,
			),
		}).Stage(formgroupFormCallback.probe.formStage)
		formgroup := new(models.FormGroup)
		FillUpForm(formgroup, newFormGroup, formgroupFormCallback.probe)
		formgroupFormCallback.probe.formStage.Commit()
	}

	fillUpTree(formgroupFormCallback.probe)
}
func __gong__New__FormSortAssocButtonFormCallback(
	formsortassocbutton *models.FormSortAssocButton,
	probe *Probe,
) (formsortassocbuttonFormCallback *FormSortAssocButtonFormCallback) {
	formsortassocbuttonFormCallback = new(FormSortAssocButtonFormCallback)
	formsortassocbuttonFormCallback.probe = probe
	formsortassocbuttonFormCallback.formsortassocbutton = formsortassocbutton

	formsortassocbuttonFormCallback.CreationMode = (formsortassocbutton == nil)

	return
}

type FormSortAssocButtonFormCallback struct {
	formsortassocbutton *models.FormSortAssocButton

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (formsortassocbuttonFormCallback *FormSortAssocButtonFormCallback) OnSave() {

	log.Println("FormSortAssocButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	formsortassocbuttonFormCallback.probe.formStage.Checkout()

	if formsortassocbuttonFormCallback.formsortassocbutton == nil {
		formsortassocbuttonFormCallback.formsortassocbutton = new(models.FormSortAssocButton).Stage(formsortassocbuttonFormCallback.probe.stageOfInterest)
	}
	formsortassocbutton_ := formsortassocbuttonFormCallback.formsortassocbutton
	_ = formsortassocbutton_

	// get the formGroup
	formGroup := formsortassocbuttonFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(formsortassocbutton_.Name), formDiv)
		case "Label":
			FormDivBasicFieldToField(&(formsortassocbutton_.Label), formDiv)
		}
	}

	formsortassocbuttonFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.FormSortAssocButton](
		formsortassocbuttonFormCallback.probe,
	)
	formsortassocbuttonFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if formsortassocbuttonFormCallback.CreationMode {
		formsortassocbuttonFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__FormSortAssocButtonFormCallback(
				nil,
				formsortassocbuttonFormCallback.probe,
			),
		}).Stage(formsortassocbuttonFormCallback.probe.formStage)
		formsortassocbutton := new(models.FormSortAssocButton)
		FillUpForm(formsortassocbutton, newFormGroup, formsortassocbuttonFormCallback.probe)
		formsortassocbuttonFormCallback.probe.formStage.Commit()
	}

	fillUpTree(formsortassocbuttonFormCallback.probe)
}
func __gong__New__OptionFormCallback(
	option *models.Option,
	probe *Probe,
) (optionFormCallback *OptionFormCallback) {
	optionFormCallback = new(OptionFormCallback)
	optionFormCallback.probe = probe
	optionFormCallback.option = option

	optionFormCallback.CreationMode = (option == nil)

	return
}

type OptionFormCallback struct {
	option *models.Option

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (optionFormCallback *OptionFormCallback) OnSave() {

	log.Println("OptionFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	optionFormCallback.probe.formStage.Checkout()

	if optionFormCallback.option == nil {
		optionFormCallback.option = new(models.Option).Stage(optionFormCallback.probe.stageOfInterest)
	}
	option_ := optionFormCallback.option
	_ = option_

	// get the formGroup
	formGroup := optionFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(option_.Name), formDiv)
		case "FormFieldSelect:Options":
			// we need to retrieve the field owner before the change
			var pastFormFieldSelectOwner *models.FormFieldSelect
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormFieldSelect"
			rf.Fieldname = "Options"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				optionFormCallback.probe.stageOfInterest,
				optionFormCallback.probe.backRepoOfInterest,
				option_,
				&rf)

			if reverseFieldOwner != nil {
				pastFormFieldSelectOwner = reverseFieldOwner.(*models.FormFieldSelect)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastFormFieldSelectOwner != nil {
					idx := slices.Index(pastFormFieldSelectOwner.Options, option_)
					pastFormFieldSelectOwner.Options = slices.Delete(pastFormFieldSelectOwner.Options, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _formfieldselect := range *models.GetGongstructInstancesSet[models.FormFieldSelect](optionFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _formfieldselect.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newFormFieldSelectOwner := _formfieldselect // we have a match
						if pastFormFieldSelectOwner != nil {
							if newFormFieldSelectOwner != pastFormFieldSelectOwner {
								idx := slices.Index(pastFormFieldSelectOwner.Options, option_)
								pastFormFieldSelectOwner.Options = slices.Delete(pastFormFieldSelectOwner.Options, idx, idx+1)
								newFormFieldSelectOwner.Options = append(newFormFieldSelectOwner.Options, option_)
							}
						} else {
							newFormFieldSelectOwner.Options = append(newFormFieldSelectOwner.Options, option_)
						}
					}
				}
			}
		}
	}

	optionFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Option](
		optionFormCallback.probe,
	)
	optionFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if optionFormCallback.CreationMode {
		optionFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__OptionFormCallback(
				nil,
				optionFormCallback.probe,
			),
		}).Stage(optionFormCallback.probe.formStage)
		option := new(models.Option)
		FillUpForm(option, newFormGroup, optionFormCallback.probe)
		optionFormCallback.probe.formStage.Commit()
	}

	fillUpTree(optionFormCallback.probe)
}
