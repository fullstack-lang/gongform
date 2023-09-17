// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongform/go/models"
)

func NewCellDeleteIconImpl[T models.Gongstruct](
	Instance *T,
	playground *Playground,
) (cellDeleteIconImpl *CellDeleteIconImpl[T]) {
	cellDeleteIconImpl = new(CellDeleteIconImpl[T])
	cellDeleteIconImpl.Instance = Instance
	cellDeleteIconImpl.playground = playground
	return
}

type CellDeleteIconImpl[T models.Gongstruct] struct {
	Instance   *T
	playground *Playground
}

func (cellDeleteIconImpl *CellDeleteIconImpl[T]) CellIconUpdated(stage *gongtable.StageStruct,
	row, updatedCellIcon *gongtable.CellIcon) {
	log.Println("CellIconUpdate: CellIconUpdated", updatedCellIcon.Name)

	switch instancesTyped := any(cellDeleteIconImpl.Instance).(type) {
	// insertion point
	case *models.CheckBox:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.FormDiv:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.FormEditAssocButton:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.FormField:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.FormFieldDate:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.FormFieldDateTime:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.FormFieldFloat64:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.FormFieldInt:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.FormFieldSelect:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.FormFieldString:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.FormFieldTime:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.FormGroup:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.FormSortAssocButton:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Option:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.playground.stageOfInterest.Commit()

	fillUpTable[T](cellDeleteIconImpl.playground)
	fillUpTree(cellDeleteIconImpl.playground)
	cellDeleteIconImpl.playground.tableStage.Commit()
}

