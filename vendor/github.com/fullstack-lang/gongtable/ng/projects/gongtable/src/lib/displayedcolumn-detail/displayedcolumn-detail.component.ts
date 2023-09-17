// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { DisplayedColumnDB } from '../displayedcolumn-db'
import { DisplayedColumnService } from '../displayedcolumn.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { TableDB } from '../table-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// DisplayedColumnDetailComponent is initilizaed from different routes
// DisplayedColumnDetailComponentState detail different cases 
enum DisplayedColumnDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Table_DisplayedColumns_SET,
}

@Component({
	selector: 'app-displayedcolumn-detail',
	templateUrl: './displayedcolumn-detail.component.html',
	styleUrls: ['./displayedcolumn-detail.component.css'],
})
export class DisplayedColumnDetailComponent implements OnInit {

	// insertion point for declarations

	// the DisplayedColumnDB of interest
	displayedcolumn: DisplayedColumnDB = new DisplayedColumnDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: DisplayedColumnDetailComponentState = DisplayedColumnDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private displayedcolumnService: DisplayedColumnService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private activatedRoute: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.GONG__StackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')!;

		this.activatedRoute.params.subscribe(params => {
			this.onChangedActivatedRoute()
		});
	}
	onChangedActivatedRoute(): void {

		// compute state
		this.id = +this.activatedRoute.snapshot.paramMap.get('id')!;
		this.originStruct = this.activatedRoute.snapshot.paramMap.get('originStruct')!;
		this.originStructFieldName = this.activatedRoute.snapshot.paramMap.get('originStructFieldName')!;

		this.GONG__StackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')!;

		const association = this.activatedRoute.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = DisplayedColumnDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = DisplayedColumnDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "DisplayedColumns":
						// console.log("DisplayedColumn" + " is instanciated with back pointer to instance " + this.id + " Table association DisplayedColumns")
						this.state = DisplayedColumnDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Table_DisplayedColumns_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getDisplayedColumn()

		// observable for changes in structs
		this.displayedcolumnService.DisplayedColumnServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getDisplayedColumn()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getDisplayedColumn(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case DisplayedColumnDetailComponentState.CREATE_INSTANCE:
						this.displayedcolumn = new (DisplayedColumnDB)
						break;
					case DisplayedColumnDetailComponentState.UPDATE_INSTANCE:
						let displayedcolumn = frontRepo.DisplayedColumns.get(this.id)
						console.assert(displayedcolumn != undefined, "missing displayedcolumn with id:" + this.id)
						this.displayedcolumn = displayedcolumn!
						break;
					// insertion point for init of association field
					case DisplayedColumnDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Table_DisplayedColumns_SET:
						this.displayedcolumn = new (DisplayedColumnDB)
						this.displayedcolumn.Table_DisplayedColumns_reverse = frontRepo.Tables.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.displayedcolumn.Table_DisplayedColumns_reverse != undefined) {
			if (this.displayedcolumn.Table_DisplayedColumnsDBID == undefined) {
				this.displayedcolumn.Table_DisplayedColumnsDBID = new NullInt64
			}
			this.displayedcolumn.Table_DisplayedColumnsDBID.Int64 = this.displayedcolumn.Table_DisplayedColumns_reverse.ID
			this.displayedcolumn.Table_DisplayedColumnsDBID.Valid = true
			if (this.displayedcolumn.Table_DisplayedColumnsDBID_Index == undefined) {
				this.displayedcolumn.Table_DisplayedColumnsDBID_Index = new NullInt64
			}
			this.displayedcolumn.Table_DisplayedColumnsDBID_Index.Valid = true
			this.displayedcolumn.Table_DisplayedColumns_reverse = new TableDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case DisplayedColumnDetailComponentState.UPDATE_INSTANCE:
				this.displayedcolumnService.updateDisplayedColumn(this.displayedcolumn, this.GONG__StackPath)
					.subscribe(displayedcolumn => {
						this.displayedcolumnService.DisplayedColumnServiceChanged.next("update")
					});
				break;
			default:
				this.displayedcolumnService.postDisplayedColumn(this.displayedcolumn, this.GONG__StackPath).subscribe(displayedcolumn => {
					this.displayedcolumnService.DisplayedColumnServiceChanged.next("post")
					this.displayedcolumn = new (DisplayedColumnDB) // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string, selectionMode: string,
		sourceField: string, intermediateStructField: string, nextAssociatedStruct: string) {

		console.log("mode " + selectionMode)

		const dialogConfig = new MatDialogConfig();

		let dialogData = new DialogData();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		if (selectionMode == SelectionMode.ONE_MANY_ASSOCIATION_MODE) {

			dialogData.ID = this.displayedcolumn.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(AssociatedStruct).get(
					AssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}
		if (selectionMode == SelectionMode.MANY_MANY_ASSOCIATION_MODE) {
			dialogData.ID = this.displayedcolumn.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "DisplayedColumn"
			dialogData.SourceField = sourceField

			// set up the intermediate struct
			dialogData.IntermediateStruct = AssociatedStruct
			dialogData.IntermediateStructField = intermediateStructField

			// set up the end struct
			dialogData.NextAssociationStruct = nextAssociatedStruct

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(nextAssociatedStruct).get(
					nextAssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}

	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.displayedcolumn.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
			GONG__StackPath: this.GONG__StackPath,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfSortingComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'SortingComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	fillUpNameIfEmpty(event: { value: { Name: string; }; }) {
		if (this.displayedcolumn.Name == "") {
			this.displayedcolumn.Name = event.value.Name
		}
	}

	toggleTextArea(fieldName: string) {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			let displayAsTextArea = this.mapFields_displayAsTextArea.get(fieldName)
			this.mapFields_displayAsTextArea.set(fieldName, !displayAsTextArea)
		} else {
			this.mapFields_displayAsTextArea.set(fieldName, true)
		}
	}

	isATextArea(fieldName: string): boolean {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			return this.mapFields_displayAsTextArea.get(fieldName)!
		} else {
			return false
		}
	}

	compareObjects(o1: any, o2: any) {
		if (o1?.ID == o2?.ID) {
			return true;
		}
		else {
			return false
		}
	}
}
