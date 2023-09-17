// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongform/go/models"
)

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.CheckBox:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormDiv:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("FormFields", instanceWithInferedType, &instanceWithInferedType.FormFields, formGroup, playground)
		AssociationSliceToForm("CheckBoxs", instanceWithInferedType, &instanceWithInferedType.CheckBoxs, formGroup, playground)
		AssociationFieldToForm("FormEditAssocButton", instanceWithInferedType.FormEditAssocButton, formGroup, playground)
		AssociationFieldToForm("FormSortAssocButton", instanceWithInferedType.FormSortAssocButton, formGroup, playground)

	case *models.FormEditAssocButton:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		EnumTypeStringToForm("InputTypeEnum", instanceWithInferedType.InputTypeEnum, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Placeholder", instanceWithInferedType.Placeholder, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("FormFieldString", instanceWithInferedType.FormFieldString, formGroup, playground)
		AssociationFieldToForm("FormFieldFloat64", instanceWithInferedType.FormFieldFloat64, formGroup, playground)
		AssociationFieldToForm("FormFieldInt", instanceWithInferedType.FormFieldInt, formGroup, playground)
		AssociationFieldToForm("FormFieldDate", instanceWithInferedType.FormFieldDate, formGroup, playground)
		AssociationFieldToForm("FormFieldTime", instanceWithInferedType.FormFieldTime, formGroup, playground)
		AssociationFieldToForm("FormFieldDateTime", instanceWithInferedType.FormFieldDateTime, formGroup, playground)
		AssociationFieldToForm("FormFieldSelect", instanceWithInferedType.FormFieldSelect, formGroup, playground)
		BasicFieldtoForm("HasBespokeWidth", instanceWithInferedType.HasBespokeWidth, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("BespokeWidthPx", instanceWithInferedType.BespokeWidthPx, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormFieldDate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormFieldDateTime:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormFieldFloat64:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasMinValidator", instanceWithInferedType.HasMinValidator, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("MinValue", instanceWithInferedType.MinValue, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasMaxValidator", instanceWithInferedType.HasMaxValidator, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("MaxValue", instanceWithInferedType.MaxValue, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormFieldInt:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasMinValidator", instanceWithInferedType.HasMinValidator, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("MinValue", instanceWithInferedType.MinValue, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasMaxValidator", instanceWithInferedType.HasMaxValidator, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("MaxValue", instanceWithInferedType.MaxValue, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormFieldSelect:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationFieldToForm("Value", instanceWithInferedType.Value, formGroup, playground)
		AssociationSliceToForm("Options", instanceWithInferedType, &instanceWithInferedType.Options, formGroup, playground)
		BasicFieldtoForm("CanBeEmpty", instanceWithInferedType.CanBeEmpty, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormFieldString:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormFieldTime:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Step", instanceWithInferedType.Step, instanceWithInferedType, playground.formStage, formGroup)

	case *models.FormGroup:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("FormDivs", instanceWithInferedType, &instanceWithInferedType.FormDivs, formGroup, playground)

	case *models.FormSortAssocButton:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Option:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)

	default:
		_ = instanceWithInferedType
	}
}
