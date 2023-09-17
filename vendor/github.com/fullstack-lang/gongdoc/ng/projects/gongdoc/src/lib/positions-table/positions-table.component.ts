// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional, Input } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData, FrontRepoService, FrontRepo, SelectionMode } from '../front-repo.service'
import { NullInt64 } from '../null-int64'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { ActivatedRoute, Router, RouterState } from '@angular/router';
import { PositionDB } from '../position-db'
import { PositionService } from '../position.service'

// insertion point for additional imports

import { RouteService } from '../route-service';

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-positionstable',
  templateUrl: './positions-table.component.html',
  styleUrls: ['./positions-table.component.css'],
})
export class PositionsTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Position instances
  selection: SelectionModel<PositionDB> = new (SelectionModel)
  initialSelection = new Array<PositionDB>()

  // the data source for the table
  positions: PositionDB[] = []
  matTableDataSource: MatTableDataSource<PositionDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.positions
  frontRepo: FrontRepo = new (FrontRepo)

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort)
  sort: MatSort | undefined
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  ngAfterViewInit() {

    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (positionDB: PositionDB, property: string) => {
      switch (property) {
        case 'ID':
          return positionDB.ID

        // insertion point for specific sorting accessor
        case 'X':
          return positionDB.X;

        case 'Y':
          return positionDB.Y;

        case 'Name':
          return positionDB.Name;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (positionDB: PositionDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the positionDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += positionDB.X.toString()
      mergedContent += positionDB.Y.toString()
      mergedContent += positionDB.Name.toLowerCase()

      let isSelected = mergedContent.includes(filter.toLowerCase())
      return isSelected
    };

    this.matTableDataSource.sort = this.sort!
    this.matTableDataSource.paginator = this.paginator!
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.matTableDataSource.filter = filterValue.trim().toLowerCase();
  }

  constructor(
    private positionService: PositionService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of position instances
    public dialogRef: MatDialogRef<PositionsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
    private activatedRoute: ActivatedRoute,

    private routeService: RouteService,
  ) {

    // compute mode
    if (dialogData == undefined) {
      this.mode = TableComponentMode.DISPLAY_MODE
    } else {
      this.GONG__StackPath = dialogData.GONG__StackPath
      switch (dialogData.SelectionMode) {
        case SelectionMode.ONE_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.ONE_MANY_ASSOCIATION_MODE
          break
        case SelectionMode.MANY_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.MANY_MANY_ASSOCIATION_MODE
          break
        default:
      }
    }

    // observable for changes in structs
    this.positionService.PositionServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getPositions()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "X",
        "Y",
        "Name",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "X",
        "Y",
        "Name",
      ]
      this.selection = new SelectionModel<PositionDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getPositions()

    this.matTableDataSource = new MatTableDataSource(this.positions)
  }

  getPositions(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.positions = this.frontRepo.Positions_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let position of this.positions) {
            let ID = this.dialogData.ID
            let revPointer = position[this.dialogData.ReversePointer as keyof PositionDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(position)
            }
            this.selection = new SelectionModel<PositionDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, PositionDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to PositionDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as PositionDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let position = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as PositionDB
              this.initialSelection.push(position)
            }
          }

          this.selection = new SelectionModel<PositionDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.positions
      }
    )
  }

  // newPosition initiate a new position
  // create a new Position objet
  newPosition() {
  }

  deletePosition(positionID: number, position: PositionDB) {
    // list of positions is truncated of position before the delete
    this.positions = this.positions.filter(h => h !== position);

    this.positionService.deletePosition(positionID, this.GONG__StackPath).subscribe(
      position => {
        this.positionService.PositionServiceChanged.next("delete")
      }
    );
  }

  editPosition(positionID: number, position: PositionDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(positionID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "position" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, positionID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.positions.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.positions.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<PositionDB>()

      // reset all initial selection of position that belong to position
      for (let position of this.initialSelection) {
        let index = position[this.dialogData.ReversePointer as keyof PositionDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(position)

      }

      // from selection, set position that belong to position
      for (let position of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = position[this.dialogData.ReversePointer as keyof PositionDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(position)
      }


      // update all position (only update selection & initial selection)
      for (let position of toUpdate) {
        this.positionService.updatePosition(position, this.GONG__StackPath)
          .subscribe(position => {
            this.positionService.PositionServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, PositionDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedPosition = new Set<number>()
      for (let position of this.initialSelection) {
        if (this.selection.selected.includes(position)) {
          // console.log("position " + position.Name + " is still selected")
        } else {
          console.log("position " + position.Name + " has been unselected")
          unselectedPosition.add(position.ID)
          console.log("is unselected " + unselectedPosition.has(position.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let position = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as PositionDB
      if (unselectedPosition.has(position.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<PositionDB>) = new Array<PositionDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          position => {
            if (!this.initialSelection.includes(position)) {
              // console.log("position " + position.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + position.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = position.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = position.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("position " + position.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<PositionDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}