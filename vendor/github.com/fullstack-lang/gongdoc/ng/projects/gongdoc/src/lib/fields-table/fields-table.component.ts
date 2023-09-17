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
import { FieldDB } from '../field-db'
import { FieldService } from '../field.service'

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
  selector: 'app-fieldstable',
  templateUrl: './fields-table.component.html',
  styleUrls: ['./fields-table.component.css'],
})
export class FieldsTableComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Field instances
  selection: SelectionModel<FieldDB> = new (SelectionModel)
  initialSelection = new Array<FieldDB>()

  // the data source for the table
  fields: FieldDB[] = []
  matTableDataSource: MatTableDataSource<FieldDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.fields
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
    this.matTableDataSource.sortingDataAccessor = (fieldDB: FieldDB, property: string) => {
      switch (property) {
        case 'ID':
          return fieldDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return fieldDB.Name;

        case 'Identifier':
          return fieldDB.Identifier;

        case 'FieldTypeAsString':
          return fieldDB.FieldTypeAsString;

        case 'Structname':
          return fieldDB.Structname;

        case 'Fieldtypename':
          return fieldDB.Fieldtypename;

        case 'GongStructShape_Fields':
          if (this.frontRepo.GongStructShapes.get(fieldDB.GongStructShape_FieldsDBID.Int64) != undefined) {
            return this.frontRepo.GongStructShapes.get(fieldDB.GongStructShape_FieldsDBID.Int64)!.Name
          } else {
            return ""
          }

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (fieldDB: FieldDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the fieldDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += fieldDB.Name.toLowerCase()
      mergedContent += fieldDB.Identifier.toLowerCase()
      mergedContent += fieldDB.FieldTypeAsString.toLowerCase()
      mergedContent += fieldDB.Structname.toLowerCase()
      mergedContent += fieldDB.Fieldtypename.toLowerCase()
      if (fieldDB.GongStructShape_FieldsDBID.Int64 != 0) {
        mergedContent += this.frontRepo.GongStructShapes.get(fieldDB.GongStructShape_FieldsDBID.Int64)!.Name.toLowerCase()
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
    private fieldService: FieldService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of field instances
    public dialogRef: MatDialogRef<FieldsTableComponent>,
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
    this.fieldService.FieldServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getFields()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Delete', // insertion point for columns to display
        "Name",
        "Identifier",
        "FieldTypeAsString",
        "Structname",
        "Fieldtypename",
        "GongStructShape_Fields",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Identifier",
        "FieldTypeAsString",
        "Structname",
        "Fieldtypename",
        "GongStructShape_Fields",
      ]
      this.selection = new SelectionModel<FieldDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    let stackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')
    if (stackPath != undefined) {
      this.GONG__StackPath = stackPath
    }

    this.getFields()

    this.matTableDataSource = new MatTableDataSource(this.fields)
  }

  getFields(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.fields = this.frontRepo.Fields_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let field of this.fields) {
            let ID = this.dialogData.ID
            let revPointer = field[this.dialogData.ReversePointer as keyof FieldDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(field)
            }
            this.selection = new SelectionModel<FieldDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, FieldDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          // we associates on sourceInstance of type SourceStruct with a MANY MANY associations to FieldDB
          // the field name is sourceField
          let sourceFieldArray = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as FieldDB[]
          if (sourceFieldArray != null) {
            for (let associationInstance of sourceFieldArray) {
              let field = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as FieldDB
              this.initialSelection.push(field)
            }
          }

          this.selection = new SelectionModel<FieldDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.fields
      }
    )
  }

  // newField initiate a new field
  // create a new Field objet
  newField() {
  }

  deleteField(fieldID: number, field: FieldDB) {
    // list of fields is truncated of field before the delete
    this.fields = this.fields.filter(h => h !== field);

    this.fieldService.deleteField(fieldID, this.GONG__StackPath).subscribe(
      field => {
        this.fieldService.FieldServiceChanged.next("delete")
      }
    );
  }

  editField(fieldID: number, field: FieldDB) {

  }

  // set editor outlet
  setEditorRouterOutlet(fieldID: number) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + "field" + "-detail"

    let outletConf: any = {}
    outletConf[outletName] = [fullPath, fieldID, this.GONG__StackPath]

    this.router.navigate([{ outlets: outletConf }])
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.fields.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.fields.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<FieldDB>()

      // reset all initial selection of field that belong to field
      for (let field of this.initialSelection) {
        let index = field[this.dialogData.ReversePointer as keyof FieldDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(field)

      }

      // from selection, set field that belong to field
      for (let field of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = field[this.dialogData.ReversePointer as keyof FieldDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(field)
      }


      // update all field (only update selection & initial selection)
      for (let field of toUpdate) {
        this.fieldService.updateField(field, this.GONG__StackPath)
          .subscribe(field => {
            this.fieldService.FieldServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, FieldDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedField = new Set<number>()
      for (let field of this.initialSelection) {
        if (this.selection.selected.includes(field)) {
          // console.log("field " + field.Name + " is still selected")
        } else {
          console.log("field " + field.Name + " has been unselected")
          unselectedField.add(field.ID)
          console.log("is unselected " + unselectedField.has(field.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let field = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as FieldDB
      if (unselectedField.has(field.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<FieldDB>) = new Array<FieldDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          field => {
            if (!this.initialSelection.includes(field)) {
              // console.log("field " + field.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + field.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = field.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = field.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("field " + field.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<FieldDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
