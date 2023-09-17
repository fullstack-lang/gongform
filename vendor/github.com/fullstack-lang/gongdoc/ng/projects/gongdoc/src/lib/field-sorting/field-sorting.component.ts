// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { FieldDB } from '../field-db'
import { FieldService } from '../field.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-field-sorting',
  templateUrl: './field-sorting.component.html',
  styleUrls: ['./field-sorting.component.css']
})
export class FieldSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of Field instances that are in the association
  associatedFields = new Array<FieldDB>();

  constructor(
    private fieldService: FieldService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of field instances
    public dialogRef: MatDialogRef<FieldSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getFields()
  }

  getFields(): void {
    this.frontRepoService.pull(this.dialogData.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let field of this.frontRepo.Fields_array) {
          let ID = this.dialogData.ID
          let revPointerID = field[this.dialogData.ReversePointer as keyof FieldDB] as unknown as NullInt64
          let revPointerID_Index = field[this.dialogData.ReversePointer + "_Index" as keyof FieldDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedFields.push(field)
          }
        }

        // sort associated field according to order
        this.associatedFields.sort((t1, t2) => {
          let t1_revPointerID_Index = t1[this.dialogData.ReversePointer + "_Index" as keyof typeof t1] as unknown as NullInt64
          let t2_revPointerID_Index = t2[this.dialogData.ReversePointer + "_Index" as keyof typeof t2] as unknown as NullInt64
          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        });
      }
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.associatedFields, event.previousIndex, event.currentIndex);

    // set the order of Field instances
    let index = 0

    for (let field of this.associatedFields) {
      let revPointerID_Index = field[this.dialogData.ReversePointer + "_Index" as keyof FieldDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedFields.forEach(
      field => {
        this.fieldService.updateField(field, this.dialogData.GONG__StackPath)
          .subscribe(field => {
            this.fieldService.FieldServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer + ' done');
  }
}
