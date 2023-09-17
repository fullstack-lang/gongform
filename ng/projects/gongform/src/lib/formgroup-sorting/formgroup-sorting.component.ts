// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { FormGroupDB } from '../formgroup-db'
import { FormGroupService } from '../formgroup.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-formgroup-sorting',
  templateUrl: './formgroup-sorting.component.html',
  styleUrls: ['./formgroup-sorting.component.css']
})
export class FormGroupSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of FormGroup instances that are in the association
  associatedFormGroups = new Array<FormGroupDB>();

  constructor(
    private formgroupService: FormGroupService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of formgroup instances
    public dialogRef: MatDialogRef<FormGroupSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getFormGroups()
  }

  getFormGroups(): void {
    this.frontRepoService.pull(this.dialogData.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let formgroup of this.frontRepo.FormGroups_array) {
          let ID = this.dialogData.ID
          let revPointerID = formgroup[this.dialogData.ReversePointer as keyof FormGroupDB] as unknown as NullInt64
          let revPointerID_Index = formgroup[this.dialogData.ReversePointer + "_Index" as keyof FormGroupDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedFormGroups.push(formgroup)
          }
        }

        // sort associated formgroup according to order
        this.associatedFormGroups.sort((t1, t2) => {
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
    moveItemInArray(this.associatedFormGroups, event.previousIndex, event.currentIndex);

    // set the order of FormGroup instances
    let index = 0

    for (let formgroup of this.associatedFormGroups) {
      let revPointerID_Index = formgroup[this.dialogData.ReversePointer + "_Index" as keyof FormGroupDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedFormGroups.forEach(
      formgroup => {
        this.formgroupService.updateFormGroup(formgroup, this.dialogData.GONG__StackPath)
          .subscribe(formgroup => {
            this.formgroupService.FormGroupServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer + ' done');
  }
}
