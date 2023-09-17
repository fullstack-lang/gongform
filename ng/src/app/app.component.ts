import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongform from 'gongform'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'Gongform Data/Model'
  view = this.default

  views: string[] = [this.default];

  DataStack = "gongform"
  ModelStacks = "github.com/fullstack-lang/gongform/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
