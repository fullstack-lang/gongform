// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongform/go/models"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
	int | CheckBoxDB | FormDivDB | FormEditAssocButtonDB | FormFieldDB | FormFieldDateDB | FormFieldDateTimeDB | FormFieldFloat64DB | FormFieldIntDB | FormFieldSelectDB | FormFieldStringDB | FormFieldTimeDB | FormGroupDB | FormSortAssocButtonDB | OptionDB
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.CheckBox:
		checkboxInstance := any(concreteInstance).(*models.CheckBox)
		ret2 := backRepo.BackRepoCheckBox.GetCheckBoxDBFromCheckBoxPtr(checkboxInstance)
		ret = any(ret2).(*T2)
	case *models.FormDiv:
		formdivInstance := any(concreteInstance).(*models.FormDiv)
		ret2 := backRepo.BackRepoFormDiv.GetFormDivDBFromFormDivPtr(formdivInstance)
		ret = any(ret2).(*T2)
	case *models.FormEditAssocButton:
		formeditassocbuttonInstance := any(concreteInstance).(*models.FormEditAssocButton)
		ret2 := backRepo.BackRepoFormEditAssocButton.GetFormEditAssocButtonDBFromFormEditAssocButtonPtr(formeditassocbuttonInstance)
		ret = any(ret2).(*T2)
	case *models.FormField:
		formfieldInstance := any(concreteInstance).(*models.FormField)
		ret2 := backRepo.BackRepoFormField.GetFormFieldDBFromFormFieldPtr(formfieldInstance)
		ret = any(ret2).(*T2)
	case *models.FormFieldDate:
		formfielddateInstance := any(concreteInstance).(*models.FormFieldDate)
		ret2 := backRepo.BackRepoFormFieldDate.GetFormFieldDateDBFromFormFieldDatePtr(formfielddateInstance)
		ret = any(ret2).(*T2)
	case *models.FormFieldDateTime:
		formfielddatetimeInstance := any(concreteInstance).(*models.FormFieldDateTime)
		ret2 := backRepo.BackRepoFormFieldDateTime.GetFormFieldDateTimeDBFromFormFieldDateTimePtr(formfielddatetimeInstance)
		ret = any(ret2).(*T2)
	case *models.FormFieldFloat64:
		formfieldfloat64Instance := any(concreteInstance).(*models.FormFieldFloat64)
		ret2 := backRepo.BackRepoFormFieldFloat64.GetFormFieldFloat64DBFromFormFieldFloat64Ptr(formfieldfloat64Instance)
		ret = any(ret2).(*T2)
	case *models.FormFieldInt:
		formfieldintInstance := any(concreteInstance).(*models.FormFieldInt)
		ret2 := backRepo.BackRepoFormFieldInt.GetFormFieldIntDBFromFormFieldIntPtr(formfieldintInstance)
		ret = any(ret2).(*T2)
	case *models.FormFieldSelect:
		formfieldselectInstance := any(concreteInstance).(*models.FormFieldSelect)
		ret2 := backRepo.BackRepoFormFieldSelect.GetFormFieldSelectDBFromFormFieldSelectPtr(formfieldselectInstance)
		ret = any(ret2).(*T2)
	case *models.FormFieldString:
		formfieldstringInstance := any(concreteInstance).(*models.FormFieldString)
		ret2 := backRepo.BackRepoFormFieldString.GetFormFieldStringDBFromFormFieldStringPtr(formfieldstringInstance)
		ret = any(ret2).(*T2)
	case *models.FormFieldTime:
		formfieldtimeInstance := any(concreteInstance).(*models.FormFieldTime)
		ret2 := backRepo.BackRepoFormFieldTime.GetFormFieldTimeDBFromFormFieldTimePtr(formfieldtimeInstance)
		ret = any(ret2).(*T2)
	case *models.FormGroup:
		formgroupInstance := any(concreteInstance).(*models.FormGroup)
		ret2 := backRepo.BackRepoFormGroup.GetFormGroupDBFromFormGroupPtr(formgroupInstance)
		ret = any(ret2).(*T2)
	case *models.FormSortAssocButton:
		formsortassocbuttonInstance := any(concreteInstance).(*models.FormSortAssocButton)
		ret2 := backRepo.BackRepoFormSortAssocButton.GetFormSortAssocButtonDBFromFormSortAssocButtonPtr(formsortassocbuttonInstance)
		ret = any(ret2).(*T2)
	case *models.Option:
		optionInstance := any(concreteInstance).(*models.Option)
		ret2 := backRepo.BackRepoOption.GetOptionDBFromOptionPtr(optionInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.CheckBox:
		tmp := GetInstanceDBFromInstance[models.CheckBox, CheckBoxDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormDiv:
		tmp := GetInstanceDBFromInstance[models.FormDiv, FormDivDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormEditAssocButton:
		tmp := GetInstanceDBFromInstance[models.FormEditAssocButton, FormEditAssocButtonDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormField:
		tmp := GetInstanceDBFromInstance[models.FormField, FormFieldDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormFieldDate:
		tmp := GetInstanceDBFromInstance[models.FormFieldDate, FormFieldDateDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormFieldDateTime:
		tmp := GetInstanceDBFromInstance[models.FormFieldDateTime, FormFieldDateTimeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormFieldFloat64:
		tmp := GetInstanceDBFromInstance[models.FormFieldFloat64, FormFieldFloat64DB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormFieldInt:
		tmp := GetInstanceDBFromInstance[models.FormFieldInt, FormFieldIntDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormFieldSelect:
		tmp := GetInstanceDBFromInstance[models.FormFieldSelect, FormFieldSelectDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormFieldString:
		tmp := GetInstanceDBFromInstance[models.FormFieldString, FormFieldStringDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormFieldTime:
		tmp := GetInstanceDBFromInstance[models.FormFieldTime, FormFieldTimeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormGroup:
		tmp := GetInstanceDBFromInstance[models.FormGroup, FormGroupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormSortAssocButton:
		tmp := GetInstanceDBFromInstance[models.FormSortAssocButton, FormSortAssocButtonDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Option:
		tmp := GetInstanceDBFromInstance[models.Option, OptionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.CheckBox:
		tmp := GetInstanceDBFromInstance[models.CheckBox, CheckBoxDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormDiv:
		tmp := GetInstanceDBFromInstance[models.FormDiv, FormDivDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormEditAssocButton:
		tmp := GetInstanceDBFromInstance[models.FormEditAssocButton, FormEditAssocButtonDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormField:
		tmp := GetInstanceDBFromInstance[models.FormField, FormFieldDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormFieldDate:
		tmp := GetInstanceDBFromInstance[models.FormFieldDate, FormFieldDateDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormFieldDateTime:
		tmp := GetInstanceDBFromInstance[models.FormFieldDateTime, FormFieldDateTimeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormFieldFloat64:
		tmp := GetInstanceDBFromInstance[models.FormFieldFloat64, FormFieldFloat64DB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormFieldInt:
		tmp := GetInstanceDBFromInstance[models.FormFieldInt, FormFieldIntDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormFieldSelect:
		tmp := GetInstanceDBFromInstance[models.FormFieldSelect, FormFieldSelectDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormFieldString:
		tmp := GetInstanceDBFromInstance[models.FormFieldString, FormFieldStringDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormFieldTime:
		tmp := GetInstanceDBFromInstance[models.FormFieldTime, FormFieldTimeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormGroup:
		tmp := GetInstanceDBFromInstance[models.FormGroup, FormGroupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.FormSortAssocButton:
		tmp := GetInstanceDBFromInstance[models.FormSortAssocButton, FormSortAssocButtonDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Option:
		tmp := GetInstanceDBFromInstance[models.Option, OptionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
