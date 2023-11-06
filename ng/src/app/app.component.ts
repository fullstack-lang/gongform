import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongform from 'gongform'
import * as gongtable from 'gongtable'


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'Gongform Data/Model'
  form_view = "Form"
  view = this.form_view

  views: string[] = [this.form_view, this.default];

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  FormName = "Form 1"
  StackType = gongform.StackType
  StackNames = gongform.StacksNames

  TableExtraPathEnum = gongtable.TableExtraPathEnum


  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
