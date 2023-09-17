// generated code - do not edit
package models

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *GongBasicField:
		// insertion point per field

	case *GongEnum:
		// insertion point per field
		if fieldName == "GongEnumValues" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*GongEnum) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*GongEnum)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.GongEnumValues).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.GongEnumValues = _inferedTypeInstance.GongEnumValues[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.GongEnumValues =
								append(_inferedTypeInstance.GongEnumValues, any(fieldInstance).(*GongEnumValue))
						}
					}
				}
			}
		}

	case *GongEnumValue:
		// insertion point per field

	case *GongLink:
		// insertion point per field

	case *GongNote:
		// insertion point per field
		if fieldName == "Links" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*GongNote) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*GongNote)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Links).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Links = _inferedTypeInstance.Links[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Links =
								append(_inferedTypeInstance.Links, any(fieldInstance).(*GongLink))
						}
					}
				}
			}
		}

	case *GongStruct:
		// insertion point per field
		if fieldName == "GongBasicFields" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*GongStruct) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*GongStruct)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.GongBasicFields).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.GongBasicFields = _inferedTypeInstance.GongBasicFields[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.GongBasicFields =
								append(_inferedTypeInstance.GongBasicFields, any(fieldInstance).(*GongBasicField))
						}
					}
				}
			}
		}
		if fieldName == "GongTimeFields" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*GongStruct) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*GongStruct)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.GongTimeFields).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.GongTimeFields = _inferedTypeInstance.GongTimeFields[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.GongTimeFields =
								append(_inferedTypeInstance.GongTimeFields, any(fieldInstance).(*GongTimeField))
						}
					}
				}
			}
		}
		if fieldName == "PointerToGongStructFields" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*GongStruct) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*GongStruct)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.PointerToGongStructFields).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.PointerToGongStructFields = _inferedTypeInstance.PointerToGongStructFields[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.PointerToGongStructFields =
								append(_inferedTypeInstance.PointerToGongStructFields, any(fieldInstance).(*PointerToGongStructField))
						}
					}
				}
			}
		}
		if fieldName == "SliceOfPointerToGongStructFields" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*GongStruct) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*GongStruct)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.SliceOfPointerToGongStructFields).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.SliceOfPointerToGongStructFields = _inferedTypeInstance.SliceOfPointerToGongStructFields[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.SliceOfPointerToGongStructFields =
								append(_inferedTypeInstance.SliceOfPointerToGongStructFields, any(fieldInstance).(*SliceOfPointerToGongStructField))
						}
					}
				}
			}
		}

	case *GongTimeField:
		// insertion point per field

	case *Meta:
		// insertion point per field
		if fieldName == "MetaReferences" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Meta) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Meta)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.MetaReferences).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.MetaReferences = _inferedTypeInstance.MetaReferences[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.MetaReferences =
								append(_inferedTypeInstance.MetaReferences, any(fieldInstance).(*MetaReference))
						}
					}
				}
			}
		}

	case *MetaReference:
		// insertion point per field

	case *ModelPkg:
		// insertion point per field

	case *PointerToGongStructField:
		// insertion point per field

	case *SliceOfPointerToGongStructField:
		// insertion point per field

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}
