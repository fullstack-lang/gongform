// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gongsvg/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_LinkAnchoredText_sql sql.NullBool
var dummy_LinkAnchoredText_time time.Duration
var dummy_LinkAnchoredText_sort sort.Float64Slice

// LinkAnchoredTextAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model linkanchoredtextAPI
type LinkAnchoredTextAPI struct {
	gorm.Model

	models.LinkAnchoredText

	// encoding of pointers
	LinkAnchoredTextPointersEnconding
}

// LinkAnchoredTextPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type LinkAnchoredTextPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// Implementation of a reverse ID for field Link{}.TextAtArrowEnd []*LinkAnchoredText
	Link_TextAtArrowEndDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Link_TextAtArrowEndDBID_Index sql.NullInt64

	// Implementation of a reverse ID for field Link{}.TextAtArrowStart []*LinkAnchoredText
	Link_TextAtArrowStartDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Link_TextAtArrowStartDBID_Index sql.NullInt64
}

// LinkAnchoredTextDB describes a linkanchoredtext in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model linkanchoredtextDB
type LinkAnchoredTextDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field linkanchoredtextDB.Name
	Name_Data sql.NullString

	// Declation for basic field linkanchoredtextDB.Content
	Content_Data sql.NullString

	// Declation for basic field linkanchoredtextDB.X_Offset
	X_Offset_Data sql.NullFloat64

	// Declation for basic field linkanchoredtextDB.Y_Offset
	Y_Offset_Data sql.NullFloat64

	// Declation for basic field linkanchoredtextDB.FontWeight
	FontWeight_Data sql.NullString

	// Declation for basic field linkanchoredtextDB.Color
	Color_Data sql.NullString

	// Declation for basic field linkanchoredtextDB.FillOpacity
	FillOpacity_Data sql.NullFloat64

	// Declation for basic field linkanchoredtextDB.Stroke
	Stroke_Data sql.NullString

	// Declation for basic field linkanchoredtextDB.StrokeWidth
	StrokeWidth_Data sql.NullFloat64

	// Declation for basic field linkanchoredtextDB.StrokeDashArray
	StrokeDashArray_Data sql.NullString

	// Declation for basic field linkanchoredtextDB.StrokeDashArrayWhenSelected
	StrokeDashArrayWhenSelected_Data sql.NullString

	// Declation for basic field linkanchoredtextDB.Transform
	Transform_Data sql.NullString
	// encoding of pointers
	LinkAnchoredTextPointersEnconding
}

// LinkAnchoredTextDBs arrays linkanchoredtextDBs
// swagger:response linkanchoredtextDBsResponse
type LinkAnchoredTextDBs []LinkAnchoredTextDB

// LinkAnchoredTextDBResponse provides response
// swagger:response linkanchoredtextDBResponse
type LinkAnchoredTextDBResponse struct {
	LinkAnchoredTextDB
}

// LinkAnchoredTextWOP is a LinkAnchoredText without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type LinkAnchoredTextWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Content string `xlsx:"2"`

	X_Offset float64 `xlsx:"3"`

	Y_Offset float64 `xlsx:"4"`

	FontWeight string `xlsx:"5"`

	Color string `xlsx:"6"`

	FillOpacity float64 `xlsx:"7"`

	Stroke string `xlsx:"8"`

	StrokeWidth float64 `xlsx:"9"`

	StrokeDashArray string `xlsx:"10"`

	StrokeDashArrayWhenSelected string `xlsx:"11"`

	Transform string `xlsx:"12"`
	// insertion for WOP pointer fields
}

var LinkAnchoredText_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Content",
	"X_Offset",
	"Y_Offset",
	"FontWeight",
	"Color",
	"FillOpacity",
	"Stroke",
	"StrokeWidth",
	"StrokeDashArray",
	"StrokeDashArrayWhenSelected",
	"Transform",
}

type BackRepoLinkAnchoredTextStruct struct {
	// stores LinkAnchoredTextDB according to their gorm ID
	Map_LinkAnchoredTextDBID_LinkAnchoredTextDB map[uint]*LinkAnchoredTextDB

	// stores LinkAnchoredTextDB ID according to LinkAnchoredText address
	Map_LinkAnchoredTextPtr_LinkAnchoredTextDBID map[*models.LinkAnchoredText]uint

	// stores LinkAnchoredText according to their gorm ID
	Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr map[uint]*models.LinkAnchoredText

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoLinkAnchoredText.stage
	return
}

func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) GetDB() *gorm.DB {
	return backRepoLinkAnchoredText.db
}

// GetLinkAnchoredTextDBFromLinkAnchoredTextPtr is a handy function to access the back repo instance from the stage instance
func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) GetLinkAnchoredTextDBFromLinkAnchoredTextPtr(linkanchoredtext *models.LinkAnchoredText) (linkanchoredtextDB *LinkAnchoredTextDB) {
	id := backRepoLinkAnchoredText.Map_LinkAnchoredTextPtr_LinkAnchoredTextDBID[linkanchoredtext]
	linkanchoredtextDB = backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextDB[id]
	return
}

// BackRepoLinkAnchoredText.CommitPhaseOne commits all staged instances of LinkAnchoredText to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for linkanchoredtext := range stage.LinkAnchoredTexts {
		backRepoLinkAnchoredText.CommitPhaseOneInstance(linkanchoredtext)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, linkanchoredtext := range backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr {
		if _, ok := stage.LinkAnchoredTexts[linkanchoredtext]; !ok {
			backRepoLinkAnchoredText.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoLinkAnchoredText.CommitDeleteInstance commits deletion of LinkAnchoredText to the BackRepo
func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) CommitDeleteInstance(id uint) (Error error) {

	linkanchoredtext := backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr[id]

	// linkanchoredtext is not staged anymore, remove linkanchoredtextDB
	linkanchoredtextDB := backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextDB[id]
	query := backRepoLinkAnchoredText.db.Unscoped().Delete(&linkanchoredtextDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoLinkAnchoredText.Map_LinkAnchoredTextPtr_LinkAnchoredTextDBID, linkanchoredtext)
	delete(backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr, id)
	delete(backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextDB, id)

	return
}

// BackRepoLinkAnchoredText.CommitPhaseOneInstance commits linkanchoredtext staged instances of LinkAnchoredText to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) CommitPhaseOneInstance(linkanchoredtext *models.LinkAnchoredText) (Error error) {

	// check if the linkanchoredtext is not commited yet
	if _, ok := backRepoLinkAnchoredText.Map_LinkAnchoredTextPtr_LinkAnchoredTextDBID[linkanchoredtext]; ok {
		return
	}

	// initiate linkanchoredtext
	var linkanchoredtextDB LinkAnchoredTextDB
	linkanchoredtextDB.CopyBasicFieldsFromLinkAnchoredText(linkanchoredtext)

	query := backRepoLinkAnchoredText.db.Create(&linkanchoredtextDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoLinkAnchoredText.Map_LinkAnchoredTextPtr_LinkAnchoredTextDBID[linkanchoredtext] = linkanchoredtextDB.ID
	backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr[linkanchoredtextDB.ID] = linkanchoredtext
	backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextDB[linkanchoredtextDB.ID] = &linkanchoredtextDB

	return
}

// BackRepoLinkAnchoredText.CommitPhaseTwo commits all staged instances of LinkAnchoredText to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, linkanchoredtext := range backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr {
		backRepoLinkAnchoredText.CommitPhaseTwoInstance(backRepo, idx, linkanchoredtext)
	}

	return
}

// BackRepoLinkAnchoredText.CommitPhaseTwoInstance commits {{structname }} of models.LinkAnchoredText to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, linkanchoredtext *models.LinkAnchoredText) (Error error) {

	// fetch matching linkanchoredtextDB
	if linkanchoredtextDB, ok := backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextDB[idx]; ok {

		linkanchoredtextDB.CopyBasicFieldsFromLinkAnchoredText(linkanchoredtext)

		// insertion point for translating pointers encodings into actual pointers
		// This loop encodes the slice of pointers linkanchoredtext.Animates into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, animateAssocEnd := range linkanchoredtext.Animates {

			// get the back repo instance at the association end
			animateAssocEnd_DB :=
				backRepo.BackRepoAnimate.GetAnimateDBFromAnimatePtr(animateAssocEnd)

			// encode reverse pointer in the association end back repo instance
			animateAssocEnd_DB.LinkAnchoredText_AnimatesDBID.Int64 = int64(linkanchoredtextDB.ID)
			animateAssocEnd_DB.LinkAnchoredText_AnimatesDBID.Valid = true
			animateAssocEnd_DB.LinkAnchoredText_AnimatesDBID_Index.Int64 = int64(idx)
			animateAssocEnd_DB.LinkAnchoredText_AnimatesDBID_Index.Valid = true
			if q := backRepoLinkAnchoredText.db.Save(animateAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		query := backRepoLinkAnchoredText.db.Save(&linkanchoredtextDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown LinkAnchoredText intance %s", linkanchoredtext.Name))
		return err
	}

	return
}

// BackRepoLinkAnchoredText.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) CheckoutPhaseOne() (Error error) {

	linkanchoredtextDBArray := make([]LinkAnchoredTextDB, 0)
	query := backRepoLinkAnchoredText.db.Find(&linkanchoredtextDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	linkanchoredtextInstancesToBeRemovedFromTheStage := make(map[*models.LinkAnchoredText]any)
	for key, value := range backRepoLinkAnchoredText.stage.LinkAnchoredTexts {
		linkanchoredtextInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, linkanchoredtextDB := range linkanchoredtextDBArray {
		backRepoLinkAnchoredText.CheckoutPhaseOneInstance(&linkanchoredtextDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		linkanchoredtext, ok := backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr[linkanchoredtextDB.ID]
		if ok {
			delete(linkanchoredtextInstancesToBeRemovedFromTheStage, linkanchoredtext)
		}
	}

	// remove from stage and back repo's 3 maps all linkanchoredtexts that are not in the checkout
	for linkanchoredtext := range linkanchoredtextInstancesToBeRemovedFromTheStage {
		linkanchoredtext.Unstage(backRepoLinkAnchoredText.GetStage())

		// remove instance from the back repo 3 maps
		linkanchoredtextID := backRepoLinkAnchoredText.Map_LinkAnchoredTextPtr_LinkAnchoredTextDBID[linkanchoredtext]
		delete(backRepoLinkAnchoredText.Map_LinkAnchoredTextPtr_LinkAnchoredTextDBID, linkanchoredtext)
		delete(backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextDB, linkanchoredtextID)
		delete(backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr, linkanchoredtextID)
	}

	return
}

// CheckoutPhaseOneInstance takes a linkanchoredtextDB that has been found in the DB, updates the backRepo and stages the
// models version of the linkanchoredtextDB
func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) CheckoutPhaseOneInstance(linkanchoredtextDB *LinkAnchoredTextDB) (Error error) {

	linkanchoredtext, ok := backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr[linkanchoredtextDB.ID]
	if !ok {
		linkanchoredtext = new(models.LinkAnchoredText)

		backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr[linkanchoredtextDB.ID] = linkanchoredtext
		backRepoLinkAnchoredText.Map_LinkAnchoredTextPtr_LinkAnchoredTextDBID[linkanchoredtext] = linkanchoredtextDB.ID

		// append model store with the new element
		linkanchoredtext.Name = linkanchoredtextDB.Name_Data.String
		linkanchoredtext.Stage(backRepoLinkAnchoredText.GetStage())
	}
	linkanchoredtextDB.CopyBasicFieldsToLinkAnchoredText(linkanchoredtext)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	linkanchoredtext.Stage(backRepoLinkAnchoredText.GetStage())

	// preserve pointer to linkanchoredtextDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_LinkAnchoredTextDBID_LinkAnchoredTextDB)[linkanchoredtextDB hold variable pointers
	linkanchoredtextDB_Data := *linkanchoredtextDB
	preservedPtrToLinkAnchoredText := &linkanchoredtextDB_Data
	backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextDB[linkanchoredtextDB.ID] = preservedPtrToLinkAnchoredText

	return
}

// BackRepoLinkAnchoredText.CheckoutPhaseTwo Checkouts all staged instances of LinkAnchoredText to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, linkanchoredtextDB := range backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextDB {
		backRepoLinkAnchoredText.CheckoutPhaseTwoInstance(backRepo, linkanchoredtextDB)
	}
	return
}

// BackRepoLinkAnchoredText.CheckoutPhaseTwoInstance Checkouts staged instances of LinkAnchoredText to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, linkanchoredtextDB *LinkAnchoredTextDB) (Error error) {

	linkanchoredtext := backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextPtr[linkanchoredtextDB.ID]
	_ = linkanchoredtext // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// This loop redeem linkanchoredtext.Animates in the stage from the encode in the back repo
	// It parses all AnimateDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	linkanchoredtext.Animates = linkanchoredtext.Animates[:0]
	// 2. loop all instances in the type in the association end
	for _, animateDB_AssocEnd := range backRepo.BackRepoAnimate.Map_AnimateDBID_AnimateDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if animateDB_AssocEnd.LinkAnchoredText_AnimatesDBID.Int64 == int64(linkanchoredtextDB.ID) {
			// 4. fetch the associated instance in the stage
			animate_AssocEnd := backRepo.BackRepoAnimate.Map_AnimateDBID_AnimatePtr[animateDB_AssocEnd.ID]
			// 5. append it the association slice
			linkanchoredtext.Animates = append(linkanchoredtext.Animates, animate_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(linkanchoredtext.Animates, func(i, j int) bool {
		animateDB_i_ID := backRepo.BackRepoAnimate.Map_AnimatePtr_AnimateDBID[linkanchoredtext.Animates[i]]
		animateDB_j_ID := backRepo.BackRepoAnimate.Map_AnimatePtr_AnimateDBID[linkanchoredtext.Animates[j]]

		animateDB_i := backRepo.BackRepoAnimate.Map_AnimateDBID_AnimateDB[animateDB_i_ID]
		animateDB_j := backRepo.BackRepoAnimate.Map_AnimateDBID_AnimateDB[animateDB_j_ID]

		return animateDB_i.LinkAnchoredText_AnimatesDBID_Index.Int64 < animateDB_j.LinkAnchoredText_AnimatesDBID_Index.Int64
	})

	return
}

// CommitLinkAnchoredText allows commit of a single linkanchoredtext (if already staged)
func (backRepo *BackRepoStruct) CommitLinkAnchoredText(linkanchoredtext *models.LinkAnchoredText) {
	backRepo.BackRepoLinkAnchoredText.CommitPhaseOneInstance(linkanchoredtext)
	if id, ok := backRepo.BackRepoLinkAnchoredText.Map_LinkAnchoredTextPtr_LinkAnchoredTextDBID[linkanchoredtext]; ok {
		backRepo.BackRepoLinkAnchoredText.CommitPhaseTwoInstance(backRepo, id, linkanchoredtext)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitLinkAnchoredText allows checkout of a single linkanchoredtext (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutLinkAnchoredText(linkanchoredtext *models.LinkAnchoredText) {
	// check if the linkanchoredtext is staged
	if _, ok := backRepo.BackRepoLinkAnchoredText.Map_LinkAnchoredTextPtr_LinkAnchoredTextDBID[linkanchoredtext]; ok {

		if id, ok := backRepo.BackRepoLinkAnchoredText.Map_LinkAnchoredTextPtr_LinkAnchoredTextDBID[linkanchoredtext]; ok {
			var linkanchoredtextDB LinkAnchoredTextDB
			linkanchoredtextDB.ID = id

			if err := backRepo.BackRepoLinkAnchoredText.db.First(&linkanchoredtextDB, id).Error; err != nil {
				log.Panicln("CheckoutLinkAnchoredText : Problem with getting object with id:", id)
			}
			backRepo.BackRepoLinkAnchoredText.CheckoutPhaseOneInstance(&linkanchoredtextDB)
			backRepo.BackRepoLinkAnchoredText.CheckoutPhaseTwoInstance(backRepo, &linkanchoredtextDB)
		}
	}
}

// CopyBasicFieldsFromLinkAnchoredText
func (linkanchoredtextDB *LinkAnchoredTextDB) CopyBasicFieldsFromLinkAnchoredText(linkanchoredtext *models.LinkAnchoredText) {
	// insertion point for fields commit

	linkanchoredtextDB.Name_Data.String = linkanchoredtext.Name
	linkanchoredtextDB.Name_Data.Valid = true

	linkanchoredtextDB.Content_Data.String = linkanchoredtext.Content
	linkanchoredtextDB.Content_Data.Valid = true

	linkanchoredtextDB.X_Offset_Data.Float64 = linkanchoredtext.X_Offset
	linkanchoredtextDB.X_Offset_Data.Valid = true

	linkanchoredtextDB.Y_Offset_Data.Float64 = linkanchoredtext.Y_Offset
	linkanchoredtextDB.Y_Offset_Data.Valid = true

	linkanchoredtextDB.FontWeight_Data.String = linkanchoredtext.FontWeight
	linkanchoredtextDB.FontWeight_Data.Valid = true

	linkanchoredtextDB.Color_Data.String = linkanchoredtext.Color
	linkanchoredtextDB.Color_Data.Valid = true

	linkanchoredtextDB.FillOpacity_Data.Float64 = linkanchoredtext.FillOpacity
	linkanchoredtextDB.FillOpacity_Data.Valid = true

	linkanchoredtextDB.Stroke_Data.String = linkanchoredtext.Stroke
	linkanchoredtextDB.Stroke_Data.Valid = true

	linkanchoredtextDB.StrokeWidth_Data.Float64 = linkanchoredtext.StrokeWidth
	linkanchoredtextDB.StrokeWidth_Data.Valid = true

	linkanchoredtextDB.StrokeDashArray_Data.String = linkanchoredtext.StrokeDashArray
	linkanchoredtextDB.StrokeDashArray_Data.Valid = true

	linkanchoredtextDB.StrokeDashArrayWhenSelected_Data.String = linkanchoredtext.StrokeDashArrayWhenSelected
	linkanchoredtextDB.StrokeDashArrayWhenSelected_Data.Valid = true

	linkanchoredtextDB.Transform_Data.String = linkanchoredtext.Transform
	linkanchoredtextDB.Transform_Data.Valid = true
}

// CopyBasicFieldsFromLinkAnchoredTextWOP
func (linkanchoredtextDB *LinkAnchoredTextDB) CopyBasicFieldsFromLinkAnchoredTextWOP(linkanchoredtext *LinkAnchoredTextWOP) {
	// insertion point for fields commit

	linkanchoredtextDB.Name_Data.String = linkanchoredtext.Name
	linkanchoredtextDB.Name_Data.Valid = true

	linkanchoredtextDB.Content_Data.String = linkanchoredtext.Content
	linkanchoredtextDB.Content_Data.Valid = true

	linkanchoredtextDB.X_Offset_Data.Float64 = linkanchoredtext.X_Offset
	linkanchoredtextDB.X_Offset_Data.Valid = true

	linkanchoredtextDB.Y_Offset_Data.Float64 = linkanchoredtext.Y_Offset
	linkanchoredtextDB.Y_Offset_Data.Valid = true

	linkanchoredtextDB.FontWeight_Data.String = linkanchoredtext.FontWeight
	linkanchoredtextDB.FontWeight_Data.Valid = true

	linkanchoredtextDB.Color_Data.String = linkanchoredtext.Color
	linkanchoredtextDB.Color_Data.Valid = true

	linkanchoredtextDB.FillOpacity_Data.Float64 = linkanchoredtext.FillOpacity
	linkanchoredtextDB.FillOpacity_Data.Valid = true

	linkanchoredtextDB.Stroke_Data.String = linkanchoredtext.Stroke
	linkanchoredtextDB.Stroke_Data.Valid = true

	linkanchoredtextDB.StrokeWidth_Data.Float64 = linkanchoredtext.StrokeWidth
	linkanchoredtextDB.StrokeWidth_Data.Valid = true

	linkanchoredtextDB.StrokeDashArray_Data.String = linkanchoredtext.StrokeDashArray
	linkanchoredtextDB.StrokeDashArray_Data.Valid = true

	linkanchoredtextDB.StrokeDashArrayWhenSelected_Data.String = linkanchoredtext.StrokeDashArrayWhenSelected
	linkanchoredtextDB.StrokeDashArrayWhenSelected_Data.Valid = true

	linkanchoredtextDB.Transform_Data.String = linkanchoredtext.Transform
	linkanchoredtextDB.Transform_Data.Valid = true
}

// CopyBasicFieldsToLinkAnchoredText
func (linkanchoredtextDB *LinkAnchoredTextDB) CopyBasicFieldsToLinkAnchoredText(linkanchoredtext *models.LinkAnchoredText) {
	// insertion point for checkout of basic fields (back repo to stage)
	linkanchoredtext.Name = linkanchoredtextDB.Name_Data.String
	linkanchoredtext.Content = linkanchoredtextDB.Content_Data.String
	linkanchoredtext.X_Offset = linkanchoredtextDB.X_Offset_Data.Float64
	linkanchoredtext.Y_Offset = linkanchoredtextDB.Y_Offset_Data.Float64
	linkanchoredtext.FontWeight = linkanchoredtextDB.FontWeight_Data.String
	linkanchoredtext.Color = linkanchoredtextDB.Color_Data.String
	linkanchoredtext.FillOpacity = linkanchoredtextDB.FillOpacity_Data.Float64
	linkanchoredtext.Stroke = linkanchoredtextDB.Stroke_Data.String
	linkanchoredtext.StrokeWidth = linkanchoredtextDB.StrokeWidth_Data.Float64
	linkanchoredtext.StrokeDashArray = linkanchoredtextDB.StrokeDashArray_Data.String
	linkanchoredtext.StrokeDashArrayWhenSelected = linkanchoredtextDB.StrokeDashArrayWhenSelected_Data.String
	linkanchoredtext.Transform = linkanchoredtextDB.Transform_Data.String
}

// CopyBasicFieldsToLinkAnchoredTextWOP
func (linkanchoredtextDB *LinkAnchoredTextDB) CopyBasicFieldsToLinkAnchoredTextWOP(linkanchoredtext *LinkAnchoredTextWOP) {
	linkanchoredtext.ID = int(linkanchoredtextDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	linkanchoredtext.Name = linkanchoredtextDB.Name_Data.String
	linkanchoredtext.Content = linkanchoredtextDB.Content_Data.String
	linkanchoredtext.X_Offset = linkanchoredtextDB.X_Offset_Data.Float64
	linkanchoredtext.Y_Offset = linkanchoredtextDB.Y_Offset_Data.Float64
	linkanchoredtext.FontWeight = linkanchoredtextDB.FontWeight_Data.String
	linkanchoredtext.Color = linkanchoredtextDB.Color_Data.String
	linkanchoredtext.FillOpacity = linkanchoredtextDB.FillOpacity_Data.Float64
	linkanchoredtext.Stroke = linkanchoredtextDB.Stroke_Data.String
	linkanchoredtext.StrokeWidth = linkanchoredtextDB.StrokeWidth_Data.Float64
	linkanchoredtext.StrokeDashArray = linkanchoredtextDB.StrokeDashArray_Data.String
	linkanchoredtext.StrokeDashArrayWhenSelected = linkanchoredtextDB.StrokeDashArrayWhenSelected_Data.String
	linkanchoredtext.Transform = linkanchoredtextDB.Transform_Data.String
}

// Backup generates a json file from a slice of all LinkAnchoredTextDB instances in the backrepo
func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "LinkAnchoredTextDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LinkAnchoredTextDB, 0)
	for _, linkanchoredtextDB := range backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextDB {
		forBackup = append(forBackup, linkanchoredtextDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json LinkAnchoredText ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json LinkAnchoredText file", err.Error())
	}
}

// Backup generates a json file from a slice of all LinkAnchoredTextDB instances in the backrepo
func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LinkAnchoredTextDB, 0)
	for _, linkanchoredtextDB := range backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextDB {
		forBackup = append(forBackup, linkanchoredtextDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("LinkAnchoredText")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&LinkAnchoredText_Fields, -1)
	for _, linkanchoredtextDB := range forBackup {

		var linkanchoredtextWOP LinkAnchoredTextWOP
		linkanchoredtextDB.CopyBasicFieldsToLinkAnchoredTextWOP(&linkanchoredtextWOP)

		row := sh.AddRow()
		row.WriteStruct(&linkanchoredtextWOP, -1)
	}
}

// RestoreXL from the "LinkAnchoredText" sheet all LinkAnchoredTextDB instances
func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoLinkAnchoredTextid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["LinkAnchoredText"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoLinkAnchoredText.rowVisitorLinkAnchoredText)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) rowVisitorLinkAnchoredText(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var linkanchoredtextWOP LinkAnchoredTextWOP
		row.ReadStruct(&linkanchoredtextWOP)

		// add the unmarshalled struct to the stage
		linkanchoredtextDB := new(LinkAnchoredTextDB)
		linkanchoredtextDB.CopyBasicFieldsFromLinkAnchoredTextWOP(&linkanchoredtextWOP)

		linkanchoredtextDB_ID_atBackupTime := linkanchoredtextDB.ID
		linkanchoredtextDB.ID = 0
		query := backRepoLinkAnchoredText.db.Create(linkanchoredtextDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextDB[linkanchoredtextDB.ID] = linkanchoredtextDB
		BackRepoLinkAnchoredTextid_atBckpTime_newID[linkanchoredtextDB_ID_atBackupTime] = linkanchoredtextDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "LinkAnchoredTextDB.json" in dirPath that stores an array
// of LinkAnchoredTextDB and stores it in the database
// the map BackRepoLinkAnchoredTextid_atBckpTime_newID is updated accordingly
func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoLinkAnchoredTextid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "LinkAnchoredTextDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json LinkAnchoredText file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*LinkAnchoredTextDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_LinkAnchoredTextDBID_LinkAnchoredTextDB
	for _, linkanchoredtextDB := range forRestore {

		linkanchoredtextDB_ID_atBackupTime := linkanchoredtextDB.ID
		linkanchoredtextDB.ID = 0
		query := backRepoLinkAnchoredText.db.Create(linkanchoredtextDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextDB[linkanchoredtextDB.ID] = linkanchoredtextDB
		BackRepoLinkAnchoredTextid_atBckpTime_newID[linkanchoredtextDB_ID_atBackupTime] = linkanchoredtextDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json LinkAnchoredText file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<LinkAnchoredText>id_atBckpTime_newID
// to compute new index
func (backRepoLinkAnchoredText *BackRepoLinkAnchoredTextStruct) RestorePhaseTwo() {

	for _, linkanchoredtextDB := range backRepoLinkAnchoredText.Map_LinkAnchoredTextDBID_LinkAnchoredTextDB {

		// next line of code is to avert unused variable compilation error
		_ = linkanchoredtextDB

		// insertion point for reindexing pointers encoding
		// This reindex linkanchoredtext.TextAtArrowEnd
		if linkanchoredtextDB.Link_TextAtArrowEndDBID.Int64 != 0 {
			linkanchoredtextDB.Link_TextAtArrowEndDBID.Int64 =
				int64(BackRepoLinkid_atBckpTime_newID[uint(linkanchoredtextDB.Link_TextAtArrowEndDBID.Int64)])
		}

		// This reindex linkanchoredtext.TextAtArrowStart
		if linkanchoredtextDB.Link_TextAtArrowStartDBID.Int64 != 0 {
			linkanchoredtextDB.Link_TextAtArrowStartDBID.Int64 =
				int64(BackRepoLinkid_atBckpTime_newID[uint(linkanchoredtextDB.Link_TextAtArrowStartDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoLinkAnchoredText.db.Model(linkanchoredtextDB).Updates(*linkanchoredtextDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoLinkAnchoredTextid_atBckpTime_newID map[uint]uint
