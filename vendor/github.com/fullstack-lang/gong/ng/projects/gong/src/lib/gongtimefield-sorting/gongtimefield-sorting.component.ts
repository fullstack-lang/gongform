// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { GongTimeFieldDB } from '../gongtimefield-db'
import { GongTimeFieldService } from '../gongtimefield.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-gongtimefield-sorting',
  templateUrl: './gongtimefield-sorting.component.html',
  styleUrls: ['./gongtimefield-sorting.component.css']
})
export class GongTimeFieldSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of GongTimeField instances that are in the association
  associatedGongTimeFields = new Array<GongTimeFieldDB>();

  constructor(
    private gongtimefieldService: GongTimeFieldService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongtimefield instances
    public dialogRef: MatDialogRef<GongTimeFieldSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getGongTimeFields()
  }

  getGongTimeFields(): void {
    this.frontRepoService.pull(this.dialogData.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let gongtimefield of this.frontRepo.GongTimeFields_array) {
          let ID = this.dialogData.ID
          let revPointerID = gongtimefield[this.dialogData.ReversePointer as keyof GongTimeFieldDB] as unknown as NullInt64
          let revPointerID_Index = gongtimefield[this.dialogData.ReversePointer + "_Index" as keyof GongTimeFieldDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedGongTimeFields.push(gongtimefield)
          }
        }

        // sort associated gongtimefield according to order
        this.associatedGongTimeFields.sort((t1, t2) => {
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
    moveItemInArray(this.associatedGongTimeFields, event.previousIndex, event.currentIndex);

    // set the order of GongTimeField instances
    let index = 0

    for (let gongtimefield of this.associatedGongTimeFields) {
      let revPointerID_Index = gongtimefield[this.dialogData.ReversePointer + "_Index" as keyof GongTimeFieldDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedGongTimeFields.forEach(
      gongtimefield => {
        this.gongtimefieldService.updateGongTimeField(gongtimefield, this.dialogData.GONG__StackPath)
          .subscribe(gongtimefield => {
            this.gongtimefieldService.GongTimeFieldServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer + ' done');
  }
}