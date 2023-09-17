// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { FormSortAssocButtonDB } from '../formsortassocbutton-db'
import { FormSortAssocButtonService } from '../formsortassocbutton.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-formsortassocbutton-sorting',
  templateUrl: './formsortassocbutton-sorting.component.html',
  styleUrls: ['./formsortassocbutton-sorting.component.css']
})
export class FormSortAssocButtonSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of FormSortAssocButton instances that are in the association
  associatedFormSortAssocButtons = new Array<FormSortAssocButtonDB>();

  constructor(
    private formsortassocbuttonService: FormSortAssocButtonService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of formsortassocbutton instances
    public dialogRef: MatDialogRef<FormSortAssocButtonSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getFormSortAssocButtons()
  }

  getFormSortAssocButtons(): void {
    this.frontRepoService.pull(this.dialogData.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let formsortassocbutton of this.frontRepo.FormSortAssocButtons_array) {
          let ID = this.dialogData.ID
          let revPointerID = formsortassocbutton[this.dialogData.ReversePointer as keyof FormSortAssocButtonDB] as unknown as NullInt64
          let revPointerID_Index = formsortassocbutton[this.dialogData.ReversePointer + "_Index" as keyof FormSortAssocButtonDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedFormSortAssocButtons.push(formsortassocbutton)
          }
        }

        // sort associated formsortassocbutton according to order
        this.associatedFormSortAssocButtons.sort((t1, t2) => {
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
    moveItemInArray(this.associatedFormSortAssocButtons, event.previousIndex, event.currentIndex);

    // set the order of FormSortAssocButton instances
    let index = 0

    for (let formsortassocbutton of this.associatedFormSortAssocButtons) {
      let revPointerID_Index = formsortassocbutton[this.dialogData.ReversePointer + "_Index" as keyof FormSortAssocButtonDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedFormSortAssocButtons.forEach(
      formsortassocbutton => {
        this.formsortassocbuttonService.updateFormSortAssocButton(formsortassocbutton, this.dialogData.GONG__StackPath)
          .subscribe(formsortassocbutton => {
            this.formsortassocbuttonService.FormSortAssocButtonServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer + ' done');
  }
}
