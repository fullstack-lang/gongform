// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gongform/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	playground *Playground
}

func NewTreeNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	playground *Playground,
) (nodeImplGongstruct *TreeNodeImplGongstruct) {

	nodeImplGongstruct = new(TreeNodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.playground = playground
	return
}

func (nodeImplGongstruct *TreeNodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "CheckBox" {
		fillUpTable[models.CheckBox](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormDiv" {
		fillUpTable[models.FormDiv](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormEditAssocButton" {
		fillUpTable[models.FormEditAssocButton](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormField" {
		fillUpTable[models.FormField](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldDate" {
		fillUpTable[models.FormFieldDate](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldDateTime" {
		fillUpTable[models.FormFieldDateTime](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldFloat64" {
		fillUpTable[models.FormFieldFloat64](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldInt" {
		fillUpTable[models.FormFieldInt](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldSelect" {
		fillUpTable[models.FormFieldSelect](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldString" {
		fillUpTable[models.FormFieldString](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormFieldTime" {
		fillUpTable[models.FormFieldTime](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormGroup" {
		fillUpTable[models.FormGroup](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "FormSortAssocButton" {
		fillUpTable[models.FormSortAssocButton](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Option" {
		fillUpTable[models.Option](nodeImplGongstruct.playground)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.playground.tableStage.Commit()
}
