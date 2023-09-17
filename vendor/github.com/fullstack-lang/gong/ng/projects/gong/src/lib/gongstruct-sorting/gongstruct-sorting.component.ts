// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { GongStructDB } from '../gongstruct-db'
import { GongStructService } from '../gongstruct.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-gongstruct-sorting',
  templateUrl: './gongstruct-sorting.component.html',
  styleUrls: ['./gongstruct-sorting.component.css']
})
export class GongStructSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of GongStruct instances that are in the association
  associatedGongStructs = new Array<GongStructDB>();

  constructor(
    private gongstructService: GongStructService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongstruct instances
    public dialogRef: MatDialogRef<GongStructSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getGongStructs()
  }

  getGongStructs(): void {
    this.frontRepoService.pull(this.dialogData.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let gongstruct of this.frontRepo.GongStructs_array) {
          let ID = this.dialogData.ID
          let revPointerID = gongstruct[this.dialogData.ReversePointer as keyof GongStructDB] as unknown as NullInt64
          let revPointerID_Index = gongstruct[this.dialogData.ReversePointer + "_Index" as keyof GongStructDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedGongStructs.push(gongstruct)
          }
        }

        // sort associated gongstruct according to order
        this.associatedGongStructs.sort((t1, t2) => {
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
    moveItemInArray(this.associatedGongStructs, event.previousIndex, event.currentIndex);

    // set the order of GongStruct instances
    let index = 0

    for (let gongstruct of this.associatedGongStructs) {
      let revPointerID_Index = gongstruct[this.dialogData.ReversePointer + "_Index" as keyof GongStructDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedGongStructs.forEach(
      gongstruct => {
        this.gongstructService.updateGongStruct(gongstruct, this.dialogData.GONG__StackPath)
          .subscribe(gongstruct => {
            this.gongstructService.GongStructServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer + ' done');
  }
}
