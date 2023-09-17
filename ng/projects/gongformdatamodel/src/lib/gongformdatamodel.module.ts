import { NgModule } from '@angular/core';

import { DataModelPanelComponent } from './data-model-panel/data-model-panel.component';

import { GongdocModule } from 'gongdoc'
import { GongdocspecificModule } from 'gongdocspecific'

import { GongtreeModule } from 'gongtree'
import { GongtreespecificModule } from 'gongtreespecific'

import { GongtableModule } from 'gongtable'
import { GongtablespecificModule } from 'gongtablespecific'
import { GongtabledatamodelModule } from 'gongtabledatamodel'

import { GongformModule } from 'gongform'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AngularSplitModule } from 'angular-split';

@NgModule({
  declarations: [
    DataModelPanelComponent
  ],
  imports: [
    GongdocModule,
    GongdocspecificModule,

    GongtreeModule,
    GongtreespecificModule,

    GongtableModule,
    GongtablespecificModule,
    GongtabledatamodelModule,

    MatRadioModule,
    FormsModule,
    CommonModule,

    AngularSplitModule,

    GongformModule,
  ],
  exports: [
    DataModelPanelComponent
  ]
})
export class GongformdatamodelModule { }