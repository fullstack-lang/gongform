// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { OptionDB } from '../option-db'
import { OptionService } from '../option.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-option-sorting',
  templateUrl: './option-sorting.component.html',
  styleUrls: ['./option-sorting.component.css']
})
export class OptionSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of Option instances that are in the association
  associatedOptions = new Array<OptionDB>();

  constructor(
    private optionService: OptionService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of option instances
    public dialogRef: MatDialogRef<OptionSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getOptions()
  }

  getOptions(): void {
    this.frontRepoService.pull(this.dialogData.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let option of this.frontRepo.Options_array) {
          let ID = this.dialogData.ID
          let revPointerID = option[this.dialogData.ReversePointer as keyof OptionDB] as unknown as NullInt64
          let revPointerID_Index = option[this.dialogData.ReversePointer + "_Index" as keyof OptionDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedOptions.push(option)
          }
        }

        // sort associated option according to order
        this.associatedOptions.sort((t1, t2) => {
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
    moveItemInArray(this.associatedOptions, event.previousIndex, event.currentIndex);

    // set the order of Option instances
    let index = 0

    for (let option of this.associatedOptions) {
      let revPointerID_Index = option[this.dialogData.ReversePointer + "_Index" as keyof OptionDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedOptions.forEach(
      option => {
        this.optionService.updateOption(option, this.dialogData.GONG__StackPath)
          .subscribe(option => {
            this.optionService.OptionServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer + ' done');
  }
}
