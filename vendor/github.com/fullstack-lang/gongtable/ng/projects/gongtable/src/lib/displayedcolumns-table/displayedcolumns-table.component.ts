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
import { DisplayedColumnDB } from '../displayedcolumn-db'
import { DisplayedColumnService } from '../displayedcolumn.service'

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
  selector: 'app-displayedcolumnstable',
  templateUrl: './displayedcolumns-table.component.html',
  styleUrls: ['./displayedcolumns-table.component.css'],
})
export class DisplayedColumnsTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of DisplayedColumn instances
  selection: SelectionModel<DisplayedColumnDB> = new (SelectionModel)
  initialSelection = new Array<DisplayedColumnDB>()

  // the data source for the table
  displayedcolumns: DisplayedColumnDB[] = []
  matTableDataSource: MatTableDataSource<DisplayedColumnDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.displayedcolumns
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
    this.matTableDataSource.sortingDataAccessor = (displayedcolumnDB: DisplayedColumnDB, property: string) => {
      switch (property) {
        case 'ID':
          return displayedcolumnDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return displayedcolumnDB.Name;

        case 'Table_DisplayedColumns':
          if (this.frontRepo.Tables.get(displayedcolumnDB.Table_DisplayedColumnsDBID.Int64) != undefined) {
            return this.frontRepo.Tables.get(displayedcolumnDB.Table_DisplayedColumnsDBID.Int64)!.Name
          } else {
            return ""
          }

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (displayedcolumnDB: DisplayedColumnDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the displayedcolumnDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += displayedcolumnDB.Name.toLowerCase()
      if (displayedcolumnDB.Table_DisplayedColumnsDBID.Int64 != 0) {
        mergedContent += this.frontRepo.Tables.get(displayedcolumnDB.Table_DisplayedColumnsDBID.Int64)!.Name.toLowerCase()
      }


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
    private displayedcolumnService: DisplayedColumnService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of displayedcolumn instances
    public dialogRef: MatDialogRef<DisplayedColumnsTableComponent>,
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
    this.displayedcolumnService.DisplayedColumnServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getDisplayedColumns()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Table_DisplayedColumns",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Table_DisplayedColumns",
      ]
      this.selection = new SelectionModel<DisplayedColumnDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getDisplayedColumns()

    this.matTableDataSource = new MatTableDataSource(this.displayedcolumns)
  }

  getDisplayedColumns(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.displayedcolumns = this.frontRepo.DisplayedColumns_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let displayedcolumn of this.displayedcolumns) {
            let ID = this.dialogData.ID
            let revPointer = displayedcolumn[this.dialogData.ReversePointer as keyof DisplayedColumnDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(displayedcolumn)
            }
            this.selection = new SelectionModel<DisplayedColumnDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, DisplayedColumnDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to DisplayedColumnDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as DisplayedColumnDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let displayedcolumn = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as DisplayedColumnDB
              this.initialSelection.push(displayedcolumn)
            }
          }

          this.selection = new SelectionModel<DisplayedColumnDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.displayedcolumns
      }
    )
  }

  // newDisplayedColumn initiate a new displayedcolumn
  // create a new DisplayedColumn objet
  newDisplayedColumn() {
  }

  deleteDisplayedColumn(displayedcolumnID: number, displayedcolumn: DisplayedColumnDB) {
    // list of displayedcolumns is truncated of displayedcolumn before the delete
    this.displayedcolumns = this.displayedcolumns.filter(h => h !== displayedcolumn);

    this.displayedcolumnService.deleteDisplayedColumn(displayedcolumnID, this.GONG__StackPath).subscribe(
      displayedcolumn => {
        this.displayedcolumnService.DisplayedColumnServiceChanged.next("delete")
      }
    );
  }

  editDisplayedColumn(displayedcolumnID: number, displayedcolumn: DisplayedColumnDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(displayedcolumnID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "displayedcolumn" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, displayedcolumnID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.displayedcolumns.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.displayedcolumns.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<DisplayedColumnDB>()

      // reset all initial selection of displayedcolumn that belong to displayedcolumn
      for (let displayedcolumn of this.initialSelection) {
        let index = displayedcolumn[this.dialogData.ReversePointer as keyof DisplayedColumnDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(displayedcolumn)

      }

      // from selection, set displayedcolumn that belong to displayedcolumn
      for (let displayedcolumn of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = displayedcolumn[this.dialogData.ReversePointer as keyof DisplayedColumnDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(displayedcolumn)
      }


      // update all displayedcolumn (only update selection & initial selection)
      for (let displayedcolumn of toUpdate) {
        this.displayedcolumnService.updateDisplayedColumn(displayedcolumn, this.GONG__StackPath)
          .subscribe(displayedcolumn => {
            this.displayedcolumnService.DisplayedColumnServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, DisplayedColumnDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedDisplayedColumn = new Set<number>()
      for (let displayedcolumn of this.initialSelection) {
        if (this.selection.selected.includes(displayedcolumn)) {
          // console.log("displayedcolumn " + displayedcolumn.Name + " is still selected")
        } else {
          console.log("displayedcolumn " + displayedcolumn.Name + " has been unselected")
          unselectedDisplayedColumn.add(displayedcolumn.ID)
          console.log("is unselected " + unselectedDisplayedColumn.has(displayedcolumn.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let displayedcolumn = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as DisplayedColumnDB
      if (unselectedDisplayedColumn.has(displayedcolumn.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<DisplayedColumnDB>) = new Array<DisplayedColumnDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          displayedcolumn => {
            if (!this.initialSelection.includes(displayedcolumn)) {
              // console.log("displayedcolumn " + displayedcolumn.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + displayedcolumn.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = displayedcolumn.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = displayedcolumn.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("displayedcolumn " + displayedcolumn.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<DisplayedColumnDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
