// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongform/go/models"
)

func FillUpFormFromGongstructName(
	playground *Playground,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = "New"
	} else {
		prefix = "Update"
	}

	switch gongstructName {
	// insertion point
	case "CheckBox":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " CheckBox Form",
			OnSave: NewCheckBoxFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		checkbox := new(models.CheckBox)
		FillUpForm(checkbox, formGroup, playground)
	case "FormDiv":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormDiv Form",
			OnSave: NewFormDivFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formdiv := new(models.FormDiv)
		FillUpForm(formdiv, formGroup, playground)
	case "FormEditAssocButton":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormEditAssocButton Form",
			OnSave: NewFormEditAssocButtonFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formeditassocbutton := new(models.FormEditAssocButton)
		FillUpForm(formeditassocbutton, formGroup, playground)
	case "FormField":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormField Form",
			OnSave: NewFormFieldFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formfield := new(models.FormField)
		FillUpForm(formfield, formGroup, playground)
	case "FormFieldDate":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldDate Form",
			OnSave: NewFormFieldDateFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formfielddate := new(models.FormFieldDate)
		FillUpForm(formfielddate, formGroup, playground)
	case "FormFieldDateTime":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldDateTime Form",
			OnSave: NewFormFieldDateTimeFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formfielddatetime := new(models.FormFieldDateTime)
		FillUpForm(formfielddatetime, formGroup, playground)
	case "FormFieldFloat64":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldFloat64 Form",
			OnSave: NewFormFieldFloat64FormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formfieldfloat64 := new(models.FormFieldFloat64)
		FillUpForm(formfieldfloat64, formGroup, playground)
	case "FormFieldInt":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldInt Form",
			OnSave: NewFormFieldIntFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formfieldint := new(models.FormFieldInt)
		FillUpForm(formfieldint, formGroup, playground)
	case "FormFieldSelect":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldSelect Form",
			OnSave: NewFormFieldSelectFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formfieldselect := new(models.FormFieldSelect)
		FillUpForm(formfieldselect, formGroup, playground)
	case "FormFieldString":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldString Form",
			OnSave: NewFormFieldStringFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formfieldstring := new(models.FormFieldString)
		FillUpForm(formfieldstring, formGroup, playground)
	case "FormFieldTime":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormFieldTime Form",
			OnSave: NewFormFieldTimeFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formfieldtime := new(models.FormFieldTime)
		FillUpForm(formfieldtime, formGroup, playground)
	case "FormGroup":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormGroup Form",
			OnSave: NewFormGroupFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formgroup := new(models.FormGroup)
		FillUpForm(formgroup, formGroup, playground)
	case "FormSortAssocButton":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " FormSortAssocButton Form",
			OnSave: NewFormSortAssocButtonFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		formsortassocbutton := new(models.FormSortAssocButton)
		FillUpForm(formsortassocbutton, formGroup, playground)
	case "Option":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Option Form",
			OnSave: NewOptionFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		option := new(models.Option)
		FillUpForm(option, formGroup, playground)
	}
	formStage.Commit()
}
