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
	case *Button:
		// insertion point per field

	case *Node:
		// insertion point per field
		if fieldName == "Children" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Node) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Node)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Children).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Children = _inferedTypeInstance.Children[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Children =
								append(_inferedTypeInstance.Children, any(fieldInstance).(*Node))
						}
					}
				}
			}
		}
		if fieldName == "Buttons" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Node) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Node)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Buttons).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Buttons = _inferedTypeInstance.Buttons[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Buttons =
								append(_inferedTypeInstance.Buttons, any(fieldInstance).(*Button))
						}
					}
				}
			}
		}

	case *Tree:
		// insertion point per field
		if fieldName == "RootNodes" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Tree) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Tree)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.RootNodes).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.RootNodes = _inferedTypeInstance.RootNodes[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.RootNodes =
								append(_inferedTypeInstance.RootNodes, any(fieldInstance).(*Node))
						}
					}
				}
			}
		}

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}
