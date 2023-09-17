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
import { CellFloat64DB } from '../cellfloat64-db'
import { CellFloat64Service } from '../cellfloat64.service'

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
  selector: 'app-cellfloat64stable',
  templateUrl: './cellfloat64s-table.component.html',
  styleUrls: ['./cellfloat64s-table.component.css'],
})
export class CellFloat64sTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of CellFloat64 instances
  selection: SelectionModel<CellFloat64DB> = new (SelectionModel)
  initialSelection = new Array<CellFloat64DB>()

  // the data source for the table
  cellfloat64s: CellFloat64DB[] = []
  matTableDataSource: MatTableDataSource<CellFloat64DB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.cellfloat64s
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
    this.matTableDataSource.sortingDataAccessor = (cellfloat64DB: CellFloat64DB, property: string) => {
      switch (property) {
        case 'ID':
          return cellfloat64DB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return cellfloat64DB.Name;

        case 'Value':
          return cellfloat64DB.Value;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (cellfloat64DB: CellFloat64DB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the cellfloat64DB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += cellfloat64DB.Name.toLowerCase()
      mergedContent += cellfloat64DB.Value.toString()

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
    private cellfloat64Service: CellFloat64Service,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of cellfloat64 instances
    public dialogRef: MatDialogRef<CellFloat64sTableComponent>,
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
    this.cellfloat64Service.CellFloat64ServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getCellFloat64s()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Value",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Value",
      ]
      this.selection = new SelectionModel<CellFloat64DB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getCellFloat64s()

    this.matTableDataSource = new MatTableDataSource(this.cellfloat64s)
  }

  getCellFloat64s(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.cellfloat64s = this.frontRepo.CellFloat64s_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let cellfloat64 of this.cellfloat64s) {
            let ID = this.dialogData.ID
            let revPointer = cellfloat64[this.dialogData.ReversePointer as keyof CellFloat64DB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(cellfloat64)
            }
            this.selection = new SelectionModel<CellFloat64DB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, CellFloat64DB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to CellFloat64DB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as CellFloat64DB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let cellfloat64 = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as CellFloat64DB
              this.initialSelection.push(cellfloat64)
            }
          }

          this.selection = new SelectionModel<CellFloat64DB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.cellfloat64s
      }
    )
  }

  // newCellFloat64 initiate a new cellfloat64
  // create a new CellFloat64 objet
  newCellFloat64() {
  }

  deleteCellFloat64(cellfloat64ID: number, cellfloat64: CellFloat64DB) {
    // list of cellfloat64s is truncated of cellfloat64 before the delete
    this.cellfloat64s = this.cellfloat64s.filter(h => h !== cellfloat64);

    this.cellfloat64Service.deleteCellFloat64(cellfloat64ID, this.GONG__StackPath).subscribe(
      cellfloat64 => {
        this.cellfloat64Service.CellFloat64ServiceChanged.next("delete")
      }
    );
  }

  editCellFloat64(cellfloat64ID: number, cellfloat64: CellFloat64DB) {

  }

  // set editor outlet
  setEditorRouterOutlet(cellfloat64ID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "cellfloat64" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, cellfloat64ID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.cellfloat64s.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.cellfloat64s.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<CellFloat64DB>()

      // reset all initial selection of cellfloat64 that belong to cellfloat64
      for (let cellfloat64 of this.initialSelection) {
        let index = cellfloat64[this.dialogData.ReversePointer as keyof CellFloat64DB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(cellfloat64)

      }

      // from selection, set cellfloat64 that belong to cellfloat64
      for (let cellfloat64 of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = cellfloat64[this.dialogData.ReversePointer as keyof CellFloat64DB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(cellfloat64)
      }


      // update all cellfloat64 (only update selection & initial selection)
      for (let cellfloat64 of toUpdate) {
        this.cellfloat64Service.updateCellFloat64(cellfloat64, this.GONG__StackPath)
          .subscribe(cellfloat64 => {
            this.cellfloat64Service.CellFloat64ServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, CellFloat64DB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedCellFloat64 = new Set<number>()
      for (let cellfloat64 of this.initialSelection) {
        if (this.selection.selected.includes(cellfloat64)) {
          // console.log("cellfloat64 " + cellfloat64.Name + " is still selected")
        } else {
          console.log("cellfloat64 " + cellfloat64.Name + " has been unselected")
          unselectedCellFloat64.add(cellfloat64.ID)
          console.log("is unselected " + unselectedCellFloat64.has(cellfloat64.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let cellfloat64 = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as CellFloat64DB
      if (unselectedCellFloat64.has(cellfloat64.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<CellFloat64DB>) = new Array<CellFloat64DB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          cellfloat64 => {
            if (!this.initialSelection.includes(cellfloat64)) {
              // console.log("cellfloat64 " + cellfloat64.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + cellfloat64.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = cellfloat64.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = cellfloat64.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("cellfloat64 " + cellfloat64.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<CellFloat64DB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
