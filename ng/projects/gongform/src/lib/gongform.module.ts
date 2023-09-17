import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatSortModule } from '@angular/material/sort'
import { MatPaginatorModule } from '@angular/material/paginator'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTreeModule } from '@angular/material/tree';
import { DragDropModule } from '@angular/cdk/drag-drop';

import { AngularSplitModule, SplitComponent } from 'angular-split';

import {
	NgxMatDatetimePickerModule,
	NgxMatNativeDateModule,
	NgxMatTimepickerModule
} from '@angular-material-components/datetime-picker';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent } from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';
import { GongstructSelectionService } from './gongstruct-selection.service'

// insertion point for imports 
import { CheckBoxsTableComponent } from './checkboxs-table/checkboxs-table.component'
import { CheckBoxSortingComponent } from './checkbox-sorting/checkbox-sorting.component'
import { CheckBoxDetailComponent } from './checkbox-detail/checkbox-detail.component'

import { FormDivsTableComponent } from './formdivs-table/formdivs-table.component'
import { FormDivSortingComponent } from './formdiv-sorting/formdiv-sorting.component'
import { FormDivDetailComponent } from './formdiv-detail/formdiv-detail.component'

import { FormEditAssocButtonsTableComponent } from './formeditassocbuttons-table/formeditassocbuttons-table.component'
import { FormEditAssocButtonSortingComponent } from './formeditassocbutton-sorting/formeditassocbutton-sorting.component'
import { FormEditAssocButtonDetailComponent } from './formeditassocbutton-detail/formeditassocbutton-detail.component'

import { FormFieldsTableComponent } from './formfields-table/formfields-table.component'
import { FormFieldSortingComponent } from './formfield-sorting/formfield-sorting.component'
import { FormFieldDetailComponent } from './formfield-detail/formfield-detail.component'

import { FormFieldDatesTableComponent } from './formfielddates-table/formfielddates-table.component'
import { FormFieldDateSortingComponent } from './formfielddate-sorting/formfielddate-sorting.component'
import { FormFieldDateDetailComponent } from './formfielddate-detail/formfielddate-detail.component'

import { FormFieldDateTimesTableComponent } from './formfielddatetimes-table/formfielddatetimes-table.component'
import { FormFieldDateTimeSortingComponent } from './formfielddatetime-sorting/formfielddatetime-sorting.component'
import { FormFieldDateTimeDetailComponent } from './formfielddatetime-detail/formfielddatetime-detail.component'

import { FormFieldFloat64sTableComponent } from './formfieldfloat64s-table/formfieldfloat64s-table.component'
import { FormFieldFloat64SortingComponent } from './formfieldfloat64-sorting/formfieldfloat64-sorting.component'
import { FormFieldFloat64DetailComponent } from './formfieldfloat64-detail/formfieldfloat64-detail.component'

import { FormFieldIntsTableComponent } from './formfieldints-table/formfieldints-table.component'
import { FormFieldIntSortingComponent } from './formfieldint-sorting/formfieldint-sorting.component'
import { FormFieldIntDetailComponent } from './formfieldint-detail/formfieldint-detail.component'

import { FormFieldSelectsTableComponent } from './formfieldselects-table/formfieldselects-table.component'
import { FormFieldSelectSortingComponent } from './formfieldselect-sorting/formfieldselect-sorting.component'
import { FormFieldSelectDetailComponent } from './formfieldselect-detail/formfieldselect-detail.component'

import { FormFieldStringsTableComponent } from './formfieldstrings-table/formfieldstrings-table.component'
import { FormFieldStringSortingComponent } from './formfieldstring-sorting/formfieldstring-sorting.component'
import { FormFieldStringDetailComponent } from './formfieldstring-detail/formfieldstring-detail.component'

import { FormFieldTimesTableComponent } from './formfieldtimes-table/formfieldtimes-table.component'
import { FormFieldTimeSortingComponent } from './formfieldtime-sorting/formfieldtime-sorting.component'
import { FormFieldTimeDetailComponent } from './formfieldtime-detail/formfieldtime-detail.component'

import { FormGroupsTableComponent } from './formgroups-table/formgroups-table.component'
import { FormGroupSortingComponent } from './formgroup-sorting/formgroup-sorting.component'
import { FormGroupDetailComponent } from './formgroup-detail/formgroup-detail.component'

import { FormSortAssocButtonsTableComponent } from './formsortassocbuttons-table/formsortassocbuttons-table.component'
import { FormSortAssocButtonSortingComponent } from './formsortassocbutton-sorting/formsortassocbutton-sorting.component'
import { FormSortAssocButtonDetailComponent } from './formsortassocbutton-detail/formsortassocbutton-detail.component'

import { OptionsTableComponent } from './options-table/options-table.component'
import { OptionSortingComponent } from './option-sorting/option-sorting.component'
import { OptionDetailComponent } from './option-detail/option-detail.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		CheckBoxsTableComponent,
		CheckBoxSortingComponent,
		CheckBoxDetailComponent,

		FormDivsTableComponent,
		FormDivSortingComponent,
		FormDivDetailComponent,

		FormEditAssocButtonsTableComponent,
		FormEditAssocButtonSortingComponent,
		FormEditAssocButtonDetailComponent,

		FormFieldsTableComponent,
		FormFieldSortingComponent,
		FormFieldDetailComponent,

		FormFieldDatesTableComponent,
		FormFieldDateSortingComponent,
		FormFieldDateDetailComponent,

		FormFieldDateTimesTableComponent,
		FormFieldDateTimeSortingComponent,
		FormFieldDateTimeDetailComponent,

		FormFieldFloat64sTableComponent,
		FormFieldFloat64SortingComponent,
		FormFieldFloat64DetailComponent,

		FormFieldIntsTableComponent,
		FormFieldIntSortingComponent,
		FormFieldIntDetailComponent,

		FormFieldSelectsTableComponent,
		FormFieldSelectSortingComponent,
		FormFieldSelectDetailComponent,

		FormFieldStringsTableComponent,
		FormFieldStringSortingComponent,
		FormFieldStringDetailComponent,

		FormFieldTimesTableComponent,
		FormFieldTimeSortingComponent,
		FormFieldTimeDetailComponent,

		FormGroupsTableComponent,
		FormGroupSortingComponent,
		FormGroupDetailComponent,

		FormSortAssocButtonsTableComponent,
		FormSortAssocButtonSortingComponent,
		FormSortAssocButtonDetailComponent,

		OptionsTableComponent,
		OptionSortingComponent,
		OptionDetailComponent,


		SplitterComponent,
		SidebarComponent
	],
	imports: [
		FormsModule,
		ReactiveFormsModule,
		CommonModule,
		RouterModule,

		AppRoutingModule,

		MatSliderModule,
		MatSelectModule,
		MatFormFieldModule,
		MatInputModule,
		MatDatepickerModule,
		MatTableModule,
		MatSortModule,
		MatPaginatorModule,
		MatCheckboxModule,
		MatButtonModule,
		MatIconModule,
		MatToolbarModule,
		MatExpansionModule,
		MatListModule,
		MatDialogModule,
		MatGridListModule,
		MatTreeModule,
		DragDropModule,

		NgxMatDatetimePickerModule,
		NgxMatNativeDateModule,
		NgxMatTimepickerModule,

		AngularSplitModule,
	],
	exports: [
		// insertion point for declarations 
		CheckBoxsTableComponent,
		CheckBoxSortingComponent,
		CheckBoxDetailComponent,

		FormDivsTableComponent,
		FormDivSortingComponent,
		FormDivDetailComponent,

		FormEditAssocButtonsTableComponent,
		FormEditAssocButtonSortingComponent,
		FormEditAssocButtonDetailComponent,

		FormFieldsTableComponent,
		FormFieldSortingComponent,
		FormFieldDetailComponent,

		FormFieldDatesTableComponent,
		FormFieldDateSortingComponent,
		FormFieldDateDetailComponent,

		FormFieldDateTimesTableComponent,
		FormFieldDateTimeSortingComponent,
		FormFieldDateTimeDetailComponent,

		FormFieldFloat64sTableComponent,
		FormFieldFloat64SortingComponent,
		FormFieldFloat64DetailComponent,

		FormFieldIntsTableComponent,
		FormFieldIntSortingComponent,
		FormFieldIntDetailComponent,

		FormFieldSelectsTableComponent,
		FormFieldSelectSortingComponent,
		FormFieldSelectDetailComponent,

		FormFieldStringsTableComponent,
		FormFieldStringSortingComponent,
		FormFieldStringDetailComponent,

		FormFieldTimesTableComponent,
		FormFieldTimeSortingComponent,
		FormFieldTimeDetailComponent,

		FormGroupsTableComponent,
		FormGroupSortingComponent,
		FormGroupDetailComponent,

		FormSortAssocButtonsTableComponent,
		FormSortAssocButtonSortingComponent,
		FormSortAssocButtonDetailComponent,

		OptionsTableComponent,
		OptionSortingComponent,
		OptionDetailComponent,


		SplitterComponent,
		SidebarComponent,

	],
	providers: [
		GongstructSelectionService,
		{
			provide: MatDialogRef,
			useValue: {}
		},
	],
})
export class GongformModule { }
