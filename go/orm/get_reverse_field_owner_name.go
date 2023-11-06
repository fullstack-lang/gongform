// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongform/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.CheckBox:
		tmp := GetInstanceDBFromInstance[models.CheckBox, CheckBoxDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			case "CheckBoxs":
				if _formdiv, ok := stage.FormDiv_CheckBoxs_reverseMap[inst]; ok {
					res = _formdiv.Name
				}
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormDiv:
		tmp := GetInstanceDBFromInstance[models.FormDiv, FormDivDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			case "FormDivs":
				if _formgroup, ok := stage.FormGroup_FormDivs_reverseMap[inst]; ok {
					res = _formgroup.Name
				}
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormEditAssocButton:
		tmp := GetInstanceDBFromInstance[models.FormEditAssocButton, FormEditAssocButtonDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormField:
		tmp := GetInstanceDBFromInstance[models.FormField, FormFieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			case "FormFields":
				if _formdiv, ok := stage.FormDiv_FormFields_reverseMap[inst]; ok {
					res = _formdiv.Name
				}
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldDate:
		tmp := GetInstanceDBFromInstance[models.FormFieldDate, FormFieldDateDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldDateTime:
		tmp := GetInstanceDBFromInstance[models.FormFieldDateTime, FormFieldDateTimeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldFloat64:
		tmp := GetInstanceDBFromInstance[models.FormFieldFloat64, FormFieldFloat64DB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldInt:
		tmp := GetInstanceDBFromInstance[models.FormFieldInt, FormFieldIntDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldSelect:
		tmp := GetInstanceDBFromInstance[models.FormFieldSelect, FormFieldSelectDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldString:
		tmp := GetInstanceDBFromInstance[models.FormFieldString, FormFieldStringDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldTime:
		tmp := GetInstanceDBFromInstance[models.FormFieldTime, FormFieldTimeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormGroup:
		tmp := GetInstanceDBFromInstance[models.FormGroup, FormGroupDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormSortAssocButton:
		tmp := GetInstanceDBFromInstance[models.FormSortAssocButton, FormSortAssocButtonDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.Option:
		tmp := GetInstanceDBFromInstance[models.Option, OptionDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			case "Options":
				if _formfieldselect, ok := stage.FormFieldSelect_Options_reverseMap[inst]; ok {
					res = _formfieldselect.Name
				}
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.CheckBox:
		tmp := GetInstanceDBFromInstance[models.CheckBox, CheckBoxDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			case "CheckBoxs":
				res = stage.FormDiv_CheckBoxs_reverseMap[inst]
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormDiv:
		tmp := GetInstanceDBFromInstance[models.FormDiv, FormDivDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			case "FormDivs":
				res = stage.FormGroup_FormDivs_reverseMap[inst]
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormEditAssocButton:
		tmp := GetInstanceDBFromInstance[models.FormEditAssocButton, FormEditAssocButtonDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormField:
		tmp := GetInstanceDBFromInstance[models.FormField, FormFieldDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			case "FormFields":
				res = stage.FormDiv_FormFields_reverseMap[inst]
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldDate:
		tmp := GetInstanceDBFromInstance[models.FormFieldDate, FormFieldDateDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldDateTime:
		tmp := GetInstanceDBFromInstance[models.FormFieldDateTime, FormFieldDateTimeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldFloat64:
		tmp := GetInstanceDBFromInstance[models.FormFieldFloat64, FormFieldFloat64DB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldInt:
		tmp := GetInstanceDBFromInstance[models.FormFieldInt, FormFieldIntDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldSelect:
		tmp := GetInstanceDBFromInstance[models.FormFieldSelect, FormFieldSelectDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldString:
		tmp := GetInstanceDBFromInstance[models.FormFieldString, FormFieldStringDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormFieldTime:
		tmp := GetInstanceDBFromInstance[models.FormFieldTime, FormFieldTimeDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormGroup:
		tmp := GetInstanceDBFromInstance[models.FormGroup, FormGroupDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.FormSortAssocButton:
		tmp := GetInstanceDBFromInstance[models.FormSortAssocButton, FormSortAssocButtonDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	case *models.Option:
		tmp := GetInstanceDBFromInstance[models.Option, OptionDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "CheckBox":
			switch reverseField.Fieldname {
			}
		case "FormDiv":
			switch reverseField.Fieldname {
			}
		case "FormEditAssocButton":
			switch reverseField.Fieldname {
			}
		case "FormField":
			switch reverseField.Fieldname {
			}
		case "FormFieldDate":
			switch reverseField.Fieldname {
			}
		case "FormFieldDateTime":
			switch reverseField.Fieldname {
			}
		case "FormFieldFloat64":
			switch reverseField.Fieldname {
			}
		case "FormFieldInt":
			switch reverseField.Fieldname {
			}
		case "FormFieldSelect":
			switch reverseField.Fieldname {
			case "Options":
				res = stage.FormFieldSelect_Options_reverseMap[inst]
			}
		case "FormFieldString":
			switch reverseField.Fieldname {
			}
		case "FormFieldTime":
			switch reverseField.Fieldname {
			}
		case "FormGroup":
			switch reverseField.Fieldname {
			}
		case "FormSortAssocButton":
			switch reverseField.Fieldname {
			}
		case "Option":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return res
}
