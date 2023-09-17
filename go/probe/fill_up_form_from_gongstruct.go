// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongform/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, playground *Playground) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.CheckBox:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update CheckBox Form",
			OnSave: NewCheckBoxFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormDiv:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormDiv Form",
			OnSave: NewFormDivFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormEditAssocButton:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormEditAssocButton Form",
			OnSave: NewFormEditAssocButtonFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormField:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormField Form",
			OnSave: NewFormFieldFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormFieldDate:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormFieldDate Form",
			OnSave: NewFormFieldDateFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormFieldDateTime:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormFieldDateTime Form",
			OnSave: NewFormFieldDateTimeFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormFieldFloat64:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormFieldFloat64 Form",
			OnSave: NewFormFieldFloat64FormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormFieldInt:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormFieldInt Form",
			OnSave: NewFormFieldIntFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormFieldSelect:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormFieldSelect Form",
			OnSave: NewFormFieldSelectFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormFieldString:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormFieldString Form",
			OnSave: NewFormFieldStringFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormFieldTime:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormFieldTime Form",
			OnSave: NewFormFieldTimeFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormGroup:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormGroup Form",
			OnSave: NewFormGroupFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.FormSortAssocButton:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update FormSortAssocButton Form",
			OnSave: NewFormSortAssocButtonFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.Option:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Option Form",
			OnSave: NewOptionFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
