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
var dummy_Point_sql sql.NullBool
var dummy_Point_time time.Duration
var dummy_Point_sort sort.Float64Slice

// PointAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model pointAPI
type PointAPI struct {
	gorm.Model

	models.Point

	// encoding of pointers
	PointPointersEnconding
}

// PointPointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type PointPointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// Implementation of a reverse ID for field Link{}.ControlPoints []*Point
	Link_ControlPointsDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Link_ControlPointsDBID_Index sql.NullInt64
}

// PointDB describes a point in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model pointDB
type PointDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field pointDB.Name
	Name_Data sql.NullString

	// Declation for basic field pointDB.X
	X_Data sql.NullFloat64

	// Declation for basic field pointDB.Y
	Y_Data sql.NullFloat64
	// encoding of pointers
	PointPointersEnconding
}

// PointDBs arrays pointDBs
// swagger:response pointDBsResponse
type PointDBs []PointDB

// PointDBResponse provides response
// swagger:response pointDBResponse
type PointDBResponse struct {
	PointDB
}

// PointWOP is a Point without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type PointWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	X float64 `xlsx:"2"`

	Y float64 `xlsx:"3"`
	// insertion for WOP pointer fields
}

var Point_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"X",
	"Y",
}

type BackRepoPointStruct struct {
	// stores PointDB according to their gorm ID
	Map_PointDBID_PointDB map[uint]*PointDB

	// stores PointDB ID according to Point address
	Map_PointPtr_PointDBID map[*models.Point]uint

	// stores Point according to their gorm ID
	Map_PointDBID_PointPtr map[uint]*models.Point

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoPoint *BackRepoPointStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoPoint.stage
	return
}

func (backRepoPoint *BackRepoPointStruct) GetDB() *gorm.DB {
	return backRepoPoint.db
}

// GetPointDBFromPointPtr is a handy function to access the back repo instance from the stage instance
func (backRepoPoint *BackRepoPointStruct) GetPointDBFromPointPtr(point *models.Point) (pointDB *PointDB) {
	id := backRepoPoint.Map_PointPtr_PointDBID[point]
	pointDB = backRepoPoint.Map_PointDBID_PointDB[id]
	return
}

// BackRepoPoint.CommitPhaseOne commits all staged instances of Point to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoPoint *BackRepoPointStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for point := range stage.Points {
		backRepoPoint.CommitPhaseOneInstance(point)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, point := range backRepoPoint.Map_PointDBID_PointPtr {
		if _, ok := stage.Points[point]; !ok {
			backRepoPoint.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoPoint.CommitDeleteInstance commits deletion of Point to the BackRepo
func (backRepoPoint *BackRepoPointStruct) CommitDeleteInstance(id uint) (Error error) {

	point := backRepoPoint.Map_PointDBID_PointPtr[id]

	// point is not staged anymore, remove pointDB
	pointDB := backRepoPoint.Map_PointDBID_PointDB[id]
	query := backRepoPoint.db.Unscoped().Delete(&pointDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete(backRepoPoint.Map_PointPtr_PointDBID, point)
	delete(backRepoPoint.Map_PointDBID_PointPtr, id)
	delete(backRepoPoint.Map_PointDBID_PointDB, id)

	return
}

// BackRepoPoint.CommitPhaseOneInstance commits point staged instances of Point to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoPoint *BackRepoPointStruct) CommitPhaseOneInstance(point *models.Point) (Error error) {

	// check if the point is not commited yet
	if _, ok := backRepoPoint.Map_PointPtr_PointDBID[point]; ok {
		return
	}

	// initiate point
	var pointDB PointDB
	pointDB.CopyBasicFieldsFromPoint(point)

	query := backRepoPoint.db.Create(&pointDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	backRepoPoint.Map_PointPtr_PointDBID[point] = pointDB.ID
	backRepoPoint.Map_PointDBID_PointPtr[pointDB.ID] = point
	backRepoPoint.Map_PointDBID_PointDB[pointDB.ID] = &pointDB

	return
}

// BackRepoPoint.CommitPhaseTwo commits all staged instances of Point to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPoint *BackRepoPointStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, point := range backRepoPoint.Map_PointDBID_PointPtr {
		backRepoPoint.CommitPhaseTwoInstance(backRepo, idx, point)
	}

	return
}

// BackRepoPoint.CommitPhaseTwoInstance commits {{structname }} of models.Point to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPoint *BackRepoPointStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, point *models.Point) (Error error) {

	// fetch matching pointDB
	if pointDB, ok := backRepoPoint.Map_PointDBID_PointDB[idx]; ok {

		pointDB.CopyBasicFieldsFromPoint(point)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoPoint.db.Save(&pointDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Point intance %s", point.Name))
		return err
	}

	return
}

// BackRepoPoint.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoPoint *BackRepoPointStruct) CheckoutPhaseOne() (Error error) {

	pointDBArray := make([]PointDB, 0)
	query := backRepoPoint.db.Find(&pointDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	pointInstancesToBeRemovedFromTheStage := make(map[*models.Point]any)
	for key, value := range backRepoPoint.stage.Points {
		pointInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, pointDB := range pointDBArray {
		backRepoPoint.CheckoutPhaseOneInstance(&pointDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		point, ok := backRepoPoint.Map_PointDBID_PointPtr[pointDB.ID]
		if ok {
			delete(pointInstancesToBeRemovedFromTheStage, point)
		}
	}

	// remove from stage and back repo's 3 maps all points that are not in the checkout
	for point := range pointInstancesToBeRemovedFromTheStage {
		point.Unstage(backRepoPoint.GetStage())

		// remove instance from the back repo 3 maps
		pointID := backRepoPoint.Map_PointPtr_PointDBID[point]
		delete(backRepoPoint.Map_PointPtr_PointDBID, point)
		delete(backRepoPoint.Map_PointDBID_PointDB, pointID)
		delete(backRepoPoint.Map_PointDBID_PointPtr, pointID)
	}

	return
}

// CheckoutPhaseOneInstance takes a pointDB that has been found in the DB, updates the backRepo and stages the
// models version of the pointDB
func (backRepoPoint *BackRepoPointStruct) CheckoutPhaseOneInstance(pointDB *PointDB) (Error error) {

	point, ok := backRepoPoint.Map_PointDBID_PointPtr[pointDB.ID]
	if !ok {
		point = new(models.Point)

		backRepoPoint.Map_PointDBID_PointPtr[pointDB.ID] = point
		backRepoPoint.Map_PointPtr_PointDBID[point] = pointDB.ID

		// append model store with the new element
		point.Name = pointDB.Name_Data.String
		point.Stage(backRepoPoint.GetStage())
	}
	pointDB.CopyBasicFieldsToPoint(point)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	point.Stage(backRepoPoint.GetStage())

	// preserve pointer to pointDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_PointDBID_PointDB)[pointDB hold variable pointers
	pointDB_Data := *pointDB
	preservedPtrToPoint := &pointDB_Data
	backRepoPoint.Map_PointDBID_PointDB[pointDB.ID] = preservedPtrToPoint

	return
}

// BackRepoPoint.CheckoutPhaseTwo Checkouts all staged instances of Point to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPoint *BackRepoPointStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, pointDB := range backRepoPoint.Map_PointDBID_PointDB {
		backRepoPoint.CheckoutPhaseTwoInstance(backRepo, pointDB)
	}
	return
}

// BackRepoPoint.CheckoutPhaseTwoInstance Checkouts staged instances of Point to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPoint *BackRepoPointStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, pointDB *PointDB) (Error error) {

	point := backRepoPoint.Map_PointDBID_PointPtr[pointDB.ID]
	_ = point // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	return
}

// CommitPoint allows commit of a single point (if already staged)
func (backRepo *BackRepoStruct) CommitPoint(point *models.Point) {
	backRepo.BackRepoPoint.CommitPhaseOneInstance(point)
	if id, ok := backRepo.BackRepoPoint.Map_PointPtr_PointDBID[point]; ok {
		backRepo.BackRepoPoint.CommitPhaseTwoInstance(backRepo, id, point)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitPoint allows checkout of a single point (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutPoint(point *models.Point) {
	// check if the point is staged
	if _, ok := backRepo.BackRepoPoint.Map_PointPtr_PointDBID[point]; ok {

		if id, ok := backRepo.BackRepoPoint.Map_PointPtr_PointDBID[point]; ok {
			var pointDB PointDB
			pointDB.ID = id

			if err := backRepo.BackRepoPoint.db.First(&pointDB, id).Error; err != nil {
				log.Panicln("CheckoutPoint : Problem with getting object with id:", id)
			}
			backRepo.BackRepoPoint.CheckoutPhaseOneInstance(&pointDB)
			backRepo.BackRepoPoint.CheckoutPhaseTwoInstance(backRepo, &pointDB)
		}
	}
}

// CopyBasicFieldsFromPoint
func (pointDB *PointDB) CopyBasicFieldsFromPoint(point *models.Point) {
	// insertion point for fields commit

	pointDB.Name_Data.String = point.Name
	pointDB.Name_Data.Valid = true

	pointDB.X_Data.Float64 = point.X
	pointDB.X_Data.Valid = true

	pointDB.Y_Data.Float64 = point.Y
	pointDB.Y_Data.Valid = true
}

// CopyBasicFieldsFromPointWOP
func (pointDB *PointDB) CopyBasicFieldsFromPointWOP(point *PointWOP) {
	// insertion point for fields commit

	pointDB.Name_Data.String = point.Name
	pointDB.Name_Data.Valid = true

	pointDB.X_Data.Float64 = point.X
	pointDB.X_Data.Valid = true

	pointDB.Y_Data.Float64 = point.Y
	pointDB.Y_Data.Valid = true
}

// CopyBasicFieldsToPoint
func (pointDB *PointDB) CopyBasicFieldsToPoint(point *models.Point) {
	// insertion point for checkout of basic fields (back repo to stage)
	point.Name = pointDB.Name_Data.String
	point.X = pointDB.X_Data.Float64
	point.Y = pointDB.Y_Data.Float64
}

// CopyBasicFieldsToPointWOP
func (pointDB *PointDB) CopyBasicFieldsToPointWOP(point *PointWOP) {
	point.ID = int(pointDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	point.Name = pointDB.Name_Data.String
	point.X = pointDB.X_Data.Float64
	point.Y = pointDB.Y_Data.Float64
}

// Backup generates a json file from a slice of all PointDB instances in the backrepo
func (backRepoPoint *BackRepoPointStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "PointDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*PointDB, 0)
	for _, pointDB := range backRepoPoint.Map_PointDBID_PointDB {
		forBackup = append(forBackup, pointDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Point ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Point file", err.Error())
	}
}

// Backup generates a json file from a slice of all PointDB instances in the backrepo
func (backRepoPoint *BackRepoPointStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*PointDB, 0)
	for _, pointDB := range backRepoPoint.Map_PointDBID_PointDB {
		forBackup = append(forBackup, pointDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Point")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Point_Fields, -1)
	for _, pointDB := range forBackup {

		var pointWOP PointWOP
		pointDB.CopyBasicFieldsToPointWOP(&pointWOP)

		row := sh.AddRow()
		row.WriteStruct(&pointWOP, -1)
	}
}

// RestoreXL from the "Point" sheet all PointDB instances
func (backRepoPoint *BackRepoPointStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoPointid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Point"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoPoint.rowVisitorPoint)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoPoint *BackRepoPointStruct) rowVisitorPoint(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var pointWOP PointWOP
		row.ReadStruct(&pointWOP)

		// add the unmarshalled struct to the stage
		pointDB := new(PointDB)
		pointDB.CopyBasicFieldsFromPointWOP(&pointWOP)

		pointDB_ID_atBackupTime := pointDB.ID
		pointDB.ID = 0
		query := backRepoPoint.db.Create(pointDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoPoint.Map_PointDBID_PointDB[pointDB.ID] = pointDB
		BackRepoPointid_atBckpTime_newID[pointDB_ID_atBackupTime] = pointDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "PointDB.json" in dirPath that stores an array
// of PointDB and stores it in the database
// the map BackRepoPointid_atBckpTime_newID is updated accordingly
func (backRepoPoint *BackRepoPointStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoPointid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "PointDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Point file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*PointDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_PointDBID_PointDB
	for _, pointDB := range forRestore {

		pointDB_ID_atBackupTime := pointDB.ID
		pointDB.ID = 0
		query := backRepoPoint.db.Create(pointDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		backRepoPoint.Map_PointDBID_PointDB[pointDB.ID] = pointDB
		BackRepoPointid_atBckpTime_newID[pointDB_ID_atBackupTime] = pointDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Point file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Point>id_atBckpTime_newID
// to compute new index
func (backRepoPoint *BackRepoPointStruct) RestorePhaseTwo() {

	for _, pointDB := range backRepoPoint.Map_PointDBID_PointDB {

		// next line of code is to avert unused variable compilation error
		_ = pointDB

		// insertion point for reindexing pointers encoding
		// This reindex point.ControlPoints
		if pointDB.Link_ControlPointsDBID.Int64 != 0 {
			pointDB.Link_ControlPointsDBID.Int64 =
				int64(BackRepoLinkid_atBckpTime_newID[uint(pointDB.Link_ControlPointsDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoPoint.db.Model(pointDB).Updates(*pointDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoPointid_atBckpTime_newID map[uint]uint
