// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongform/go/models"
	"github.com/fullstack-lang/gongform/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.CheckBox:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormDiv"
			rf.Fieldname = "CheckBoxs"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.FormDiv),
					"CheckBoxs",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.FormDiv, *models.CheckBox](
					nil,
					"CheckBoxs",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.FormDiv:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("FormFields", instanceWithInferedType, &instanceWithInferedType.FormFields, formGroup, probe)
		AssociationSliceToForm("CheckBoxs", instanceWithInferedType, &instanceWithInferedType.CheckBoxs, formGroup, probe)
		AssociationFieldToForm("FormEditAssocButton", instanceWithInferedType.FormEditAssocButton, formGroup, probe)
		AssociationFieldToForm("FormSortAssocButton", instanceWithInferedType.FormSortAssocButton, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormGroup"
			rf.Fieldname = "FormDivs"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.FormGroup),
					"FormDivs",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.FormGroup, *models.FormDiv](
					nil,
					"FormDivs",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.FormEditAssocButton:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.FormField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("InputTypeEnum", instanceWithInferedType.InputTypeEnum, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placeholder", instanceWithInferedType.Placeholder, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("FormFieldString", instanceWithInferedType.FormFieldString, formGroup, probe)
		AssociationFieldToForm("FormFieldFloat64", instanceWithInferedType.FormFieldFloat64, formGroup, probe)
		AssociationFieldToForm("FormFieldInt", instanceWithInferedType.FormFieldInt, formGroup, probe)
		AssociationFieldToForm("FormFieldDate", instanceWithInferedType.FormFieldDate, formGroup, probe)
		AssociationFieldToForm("FormFieldTime", instanceWithInferedType.FormFieldTime, formGroup, probe)
		AssociationFieldToForm("FormFieldDateTime", instanceWithInferedType.FormFieldDateTime, formGroup, probe)
		AssociationFieldToForm("FormFieldSelect", instanceWithInferedType.FormFieldSelect, formGroup, probe)
		BasicFieldtoForm("HasBespokeWidth", instanceWithInferedType.HasBespokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BespokeWidthPx", instanceWithInferedType.BespokeWidthPx, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormDiv"
			rf.Fieldname = "FormFields"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.FormDiv),
					"FormFields",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.FormDiv, *models.FormField](
					nil,
					"FormFields",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.FormFieldDate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.FormFieldDateTime:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.FormFieldFloat64:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasMinValidator", instanceWithInferedType.HasMinValidator, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MinValue", instanceWithInferedType.MinValue, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasMaxValidator", instanceWithInferedType.HasMaxValidator, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MaxValue", instanceWithInferedType.MaxValue, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.FormFieldInt:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasMinValidator", instanceWithInferedType.HasMinValidator, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MinValue", instanceWithInferedType.MinValue, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasMaxValidator", instanceWithInferedType.HasMaxValidator, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MaxValue", instanceWithInferedType.MaxValue, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.FormFieldSelect:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Value", instanceWithInferedType.Value, formGroup, probe)
		AssociationSliceToForm("Options", instanceWithInferedType, &instanceWithInferedType.Options, formGroup, probe)
		BasicFieldtoForm("CanBeEmpty", instanceWithInferedType.CanBeEmpty, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.FormFieldString:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsTextArea", instanceWithInferedType.IsTextArea, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.FormFieldTime:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Step", instanceWithInferedType.Step, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.FormGroup:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("FormDivs", instanceWithInferedType, &instanceWithInferedType.FormDivs, formGroup, probe)

	case *models.FormSortAssocButton:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Option:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "FormFieldSelect"
			rf.Fieldname = "Options"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.FormFieldSelect),
					"Options",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.FormFieldSelect, *models.Option](
					nil,
					"Options",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	default:
		_ = instanceWithInferedType
	}
}
