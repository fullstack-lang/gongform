// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongform/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	FillUpNamedFormFromGongstruct[T](instance, probe, formStage, gongtable.FormGroupDefaultName.ToString())

}

func FillUpNamedFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.CheckBox:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "CheckBox Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CheckBoxFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormDiv:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "FormDiv Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormDivFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormEditAssocButton:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "FormEditAssocButton Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormEditAssocButtonFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormField:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "FormField Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormFieldDate:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "FormFieldDate Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldDateFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormFieldDateTime:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "FormFieldDateTime Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldDateTimeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormFieldFloat64:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "FormFieldFloat64 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldFloat64FormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormFieldInt:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "FormFieldInt Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldIntFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormFieldSelect:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "FormFieldSelect Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldSelectFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormFieldString:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "FormFieldString Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldStringFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormFieldTime:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "FormFieldTime Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormFieldTimeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormGroup:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "FormGroup Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormGroupFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.FormSortAssocButton:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "FormSortAssocButton Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FormSortAssocButtonFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Option:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Option Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OptionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
