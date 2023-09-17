import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { CheckBoxDB } from './checkbox-db'
import { CheckBoxService } from './checkbox.service'

import { FormDivDB } from './formdiv-db'
import { FormDivService } from './formdiv.service'

import { FormEditAssocButtonDB } from './formeditassocbutton-db'
import { FormEditAssocButtonService } from './formeditassocbutton.service'

import { FormFieldDB } from './formfield-db'
import { FormFieldService } from './formfield.service'

import { FormFieldDateDB } from './formfielddate-db'
import { FormFieldDateService } from './formfielddate.service'

import { FormFieldDateTimeDB } from './formfielddatetime-db'
import { FormFieldDateTimeService } from './formfielddatetime.service'

import { FormFieldFloat64DB } from './formfieldfloat64-db'
import { FormFieldFloat64Service } from './formfieldfloat64.service'

import { FormFieldIntDB } from './formfieldint-db'
import { FormFieldIntService } from './formfieldint.service'

import { FormFieldSelectDB } from './formfieldselect-db'
import { FormFieldSelectService } from './formfieldselect.service'

import { FormFieldStringDB } from './formfieldstring-db'
import { FormFieldStringService } from './formfieldstring.service'

import { FormFieldTimeDB } from './formfieldtime-db'
import { FormFieldTimeService } from './formfieldtime.service'

import { FormGroupDB } from './formgroup-db'
import { FormGroupService } from './formgroup.service'

import { FormSortAssocButtonDB } from './formsortassocbutton-db'
import { FormSortAssocButtonService } from './formsortassocbutton.service'

import { OptionDB } from './option-db'
import { OptionService } from './option.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  CheckBoxs_array = new Array<CheckBoxDB>(); // array of repo instances
  CheckBoxs = new Map<number, CheckBoxDB>(); // map of repo instances
  CheckBoxs_batch = new Map<number, CheckBoxDB>(); // same but only in last GET (for finding repo instances to delete)
  FormDivs_array = new Array<FormDivDB>(); // array of repo instances
  FormDivs = new Map<number, FormDivDB>(); // map of repo instances
  FormDivs_batch = new Map<number, FormDivDB>(); // same but only in last GET (for finding repo instances to delete)
  FormEditAssocButtons_array = new Array<FormEditAssocButtonDB>(); // array of repo instances
  FormEditAssocButtons = new Map<number, FormEditAssocButtonDB>(); // map of repo instances
  FormEditAssocButtons_batch = new Map<number, FormEditAssocButtonDB>(); // same but only in last GET (for finding repo instances to delete)
  FormFields_array = new Array<FormFieldDB>(); // array of repo instances
  FormFields = new Map<number, FormFieldDB>(); // map of repo instances
  FormFields_batch = new Map<number, FormFieldDB>(); // same but only in last GET (for finding repo instances to delete)
  FormFieldDates_array = new Array<FormFieldDateDB>(); // array of repo instances
  FormFieldDates = new Map<number, FormFieldDateDB>(); // map of repo instances
  FormFieldDates_batch = new Map<number, FormFieldDateDB>(); // same but only in last GET (for finding repo instances to delete)
  FormFieldDateTimes_array = new Array<FormFieldDateTimeDB>(); // array of repo instances
  FormFieldDateTimes = new Map<number, FormFieldDateTimeDB>(); // map of repo instances
  FormFieldDateTimes_batch = new Map<number, FormFieldDateTimeDB>(); // same but only in last GET (for finding repo instances to delete)
  FormFieldFloat64s_array = new Array<FormFieldFloat64DB>(); // array of repo instances
  FormFieldFloat64s = new Map<number, FormFieldFloat64DB>(); // map of repo instances
  FormFieldFloat64s_batch = new Map<number, FormFieldFloat64DB>(); // same but only in last GET (for finding repo instances to delete)
  FormFieldInts_array = new Array<FormFieldIntDB>(); // array of repo instances
  FormFieldInts = new Map<number, FormFieldIntDB>(); // map of repo instances
  FormFieldInts_batch = new Map<number, FormFieldIntDB>(); // same but only in last GET (for finding repo instances to delete)
  FormFieldSelects_array = new Array<FormFieldSelectDB>(); // array of repo instances
  FormFieldSelects = new Map<number, FormFieldSelectDB>(); // map of repo instances
  FormFieldSelects_batch = new Map<number, FormFieldSelectDB>(); // same but only in last GET (for finding repo instances to delete)
  FormFieldStrings_array = new Array<FormFieldStringDB>(); // array of repo instances
  FormFieldStrings = new Map<number, FormFieldStringDB>(); // map of repo instances
  FormFieldStrings_batch = new Map<number, FormFieldStringDB>(); // same but only in last GET (for finding repo instances to delete)
  FormFieldTimes_array = new Array<FormFieldTimeDB>(); // array of repo instances
  FormFieldTimes = new Map<number, FormFieldTimeDB>(); // map of repo instances
  FormFieldTimes_batch = new Map<number, FormFieldTimeDB>(); // same but only in last GET (for finding repo instances to delete)
  FormGroups_array = new Array<FormGroupDB>(); // array of repo instances
  FormGroups = new Map<number, FormGroupDB>(); // map of repo instances
  FormGroups_batch = new Map<number, FormGroupDB>(); // same but only in last GET (for finding repo instances to delete)
  FormSortAssocButtons_array = new Array<FormSortAssocButtonDB>(); // array of repo instances
  FormSortAssocButtons = new Map<number, FormSortAssocButtonDB>(); // map of repo instances
  FormSortAssocButtons_batch = new Map<number, FormSortAssocButtonDB>(); // same but only in last GET (for finding repo instances to delete)
  Options_array = new Array<OptionDB>(); // array of repo instances
  Options = new Map<number, OptionDB>(); // map of repo instances
  Options_batch = new Map<number, OptionDB>(); // same but only in last GET (for finding repo instances to delete)
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number = 0 // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean = false // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string = ""  // The "Aclass"
  SourceField: string = "" // the "AnarrayofbUse"
  IntermediateStruct: string = "" // the "AclassBclassUse" 
  IntermediateStructField: string = "" // the "Bclass" as field
  NextAssociationStruct: string = "" // the "Bclass"

  GONG__StackPath: string = ""
}

export enum SelectionMode {
  ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
  MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  GONG__StackPath: string = ""

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  //
  // Store of all instances of the stack
  //
  frontRepo = new (FrontRepo)

  constructor(
    private http: HttpClient, // insertion point sub template 
    private checkboxService: CheckBoxService,
    private formdivService: FormDivService,
    private formeditassocbuttonService: FormEditAssocButtonService,
    private formfieldService: FormFieldService,
    private formfielddateService: FormFieldDateService,
    private formfielddatetimeService: FormFieldDateTimeService,
    private formfieldfloat64Service: FormFieldFloat64Service,
    private formfieldintService: FormFieldIntService,
    private formfieldselectService: FormFieldSelectService,
    private formfieldstringService: FormFieldStringService,
    private formfieldtimeService: FormFieldTimeService,
    private formgroupService: FormGroupService,
    private formsortassocbuttonService: FormSortAssocButtonService,
    private optionService: OptionService,
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
    let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

    servicePostFunction(instanceToBePosted).subscribe(
      instance => {
        let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
    let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

    serviceDeleteFunction(instanceToBeDeleted).subscribe(
      instance => {
        let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<CheckBoxDB[]>,
    Observable<FormDivDB[]>,
    Observable<FormEditAssocButtonDB[]>,
    Observable<FormFieldDB[]>,
    Observable<FormFieldDateDB[]>,
    Observable<FormFieldDateTimeDB[]>,
    Observable<FormFieldFloat64DB[]>,
    Observable<FormFieldIntDB[]>,
    Observable<FormFieldSelectDB[]>,
    Observable<FormFieldStringDB[]>,
    Observable<FormFieldTimeDB[]>,
    Observable<FormGroupDB[]>,
    Observable<FormSortAssocButtonDB[]>,
    Observable<OptionDB[]>,
  ] = [ // insertion point sub template
      this.checkboxService.getCheckBoxs(this.GONG__StackPath),
      this.formdivService.getFormDivs(this.GONG__StackPath),
      this.formeditassocbuttonService.getFormEditAssocButtons(this.GONG__StackPath),
      this.formfieldService.getFormFields(this.GONG__StackPath),
      this.formfielddateService.getFormFieldDates(this.GONG__StackPath),
      this.formfielddatetimeService.getFormFieldDateTimes(this.GONG__StackPath),
      this.formfieldfloat64Service.getFormFieldFloat64s(this.GONG__StackPath),
      this.formfieldintService.getFormFieldInts(this.GONG__StackPath),
      this.formfieldselectService.getFormFieldSelects(this.GONG__StackPath),
      this.formfieldstringService.getFormFieldStrings(this.GONG__StackPath),
      this.formfieldtimeService.getFormFieldTimes(this.GONG__StackPath),
      this.formgroupService.getFormGroups(this.GONG__StackPath),
      this.formsortassocbuttonService.getFormSortAssocButtons(this.GONG__StackPath),
      this.optionService.getOptions(this.GONG__StackPath),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

    this.GONG__StackPath = GONG__StackPath

    this.observableFrontRepo = [ // insertion point sub template
      this.checkboxService.getCheckBoxs(this.GONG__StackPath),
      this.formdivService.getFormDivs(this.GONG__StackPath),
      this.formeditassocbuttonService.getFormEditAssocButtons(this.GONG__StackPath),
      this.formfieldService.getFormFields(this.GONG__StackPath),
      this.formfielddateService.getFormFieldDates(this.GONG__StackPath),
      this.formfielddatetimeService.getFormFieldDateTimes(this.GONG__StackPath),
      this.formfieldfloat64Service.getFormFieldFloat64s(this.GONG__StackPath),
      this.formfieldintService.getFormFieldInts(this.GONG__StackPath),
      this.formfieldselectService.getFormFieldSelects(this.GONG__StackPath),
      this.formfieldstringService.getFormFieldStrings(this.GONG__StackPath),
      this.formfieldtimeService.getFormFieldTimes(this.GONG__StackPath),
      this.formgroupService.getFormGroups(this.GONG__StackPath),
      this.formsortassocbuttonService.getFormSortAssocButtons(this.GONG__StackPath),
      this.optionService.getOptions(this.GONG__StackPath),
    ]

    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            checkboxs_,
            formdivs_,
            formeditassocbuttons_,
            formfields_,
            formfielddates_,
            formfielddatetimes_,
            formfieldfloat64s_,
            formfieldints_,
            formfieldselects_,
            formfieldstrings_,
            formfieldtimes_,
            formgroups_,
            formsortassocbuttons_,
            options_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var checkboxs: CheckBoxDB[]
            checkboxs = checkboxs_ as CheckBoxDB[]
            var formdivs: FormDivDB[]
            formdivs = formdivs_ as FormDivDB[]
            var formeditassocbuttons: FormEditAssocButtonDB[]
            formeditassocbuttons = formeditassocbuttons_ as FormEditAssocButtonDB[]
            var formfields: FormFieldDB[]
            formfields = formfields_ as FormFieldDB[]
            var formfielddates: FormFieldDateDB[]
            formfielddates = formfielddates_ as FormFieldDateDB[]
            var formfielddatetimes: FormFieldDateTimeDB[]
            formfielddatetimes = formfielddatetimes_ as FormFieldDateTimeDB[]
            var formfieldfloat64s: FormFieldFloat64DB[]
            formfieldfloat64s = formfieldfloat64s_ as FormFieldFloat64DB[]
            var formfieldints: FormFieldIntDB[]
            formfieldints = formfieldints_ as FormFieldIntDB[]
            var formfieldselects: FormFieldSelectDB[]
            formfieldselects = formfieldselects_ as FormFieldSelectDB[]
            var formfieldstrings: FormFieldStringDB[]
            formfieldstrings = formfieldstrings_ as FormFieldStringDB[]
            var formfieldtimes: FormFieldTimeDB[]
            formfieldtimes = formfieldtimes_ as FormFieldTimeDB[]
            var formgroups: FormGroupDB[]
            formgroups = formgroups_ as FormGroupDB[]
            var formsortassocbuttons: FormSortAssocButtonDB[]
            formsortassocbuttons = formsortassocbuttons_ as FormSortAssocButtonDB[]
            var options: OptionDB[]
            options = options_ as OptionDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            this.frontRepo.CheckBoxs_array = checkboxs

            // clear the map that counts CheckBox in the GET
            this.frontRepo.CheckBoxs_batch.clear()

            checkboxs.forEach(
              checkbox => {
                this.frontRepo.CheckBoxs.set(checkbox.ID, checkbox)
                this.frontRepo.CheckBoxs_batch.set(checkbox.ID, checkbox)
              }
            )

            // clear checkboxs that are absent from the batch
            this.frontRepo.CheckBoxs.forEach(
              checkbox => {
                if (this.frontRepo.CheckBoxs_batch.get(checkbox.ID) == undefined) {
                  this.frontRepo.CheckBoxs.delete(checkbox.ID)
                }
              }
            )

            // sort CheckBoxs_array array
            this.frontRepo.CheckBoxs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormDivs_array = formdivs

            // clear the map that counts FormDiv in the GET
            this.frontRepo.FormDivs_batch.clear()

            formdivs.forEach(
              formdiv => {
                this.frontRepo.FormDivs.set(formdiv.ID, formdiv)
                this.frontRepo.FormDivs_batch.set(formdiv.ID, formdiv)
              }
            )

            // clear formdivs that are absent from the batch
            this.frontRepo.FormDivs.forEach(
              formdiv => {
                if (this.frontRepo.FormDivs_batch.get(formdiv.ID) == undefined) {
                  this.frontRepo.FormDivs.delete(formdiv.ID)
                }
              }
            )

            // sort FormDivs_array array
            this.frontRepo.FormDivs_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormEditAssocButtons_array = formeditassocbuttons

            // clear the map that counts FormEditAssocButton in the GET
            this.frontRepo.FormEditAssocButtons_batch.clear()

            formeditassocbuttons.forEach(
              formeditassocbutton => {
                this.frontRepo.FormEditAssocButtons.set(formeditassocbutton.ID, formeditassocbutton)
                this.frontRepo.FormEditAssocButtons_batch.set(formeditassocbutton.ID, formeditassocbutton)
              }
            )

            // clear formeditassocbuttons that are absent from the batch
            this.frontRepo.FormEditAssocButtons.forEach(
              formeditassocbutton => {
                if (this.frontRepo.FormEditAssocButtons_batch.get(formeditassocbutton.ID) == undefined) {
                  this.frontRepo.FormEditAssocButtons.delete(formeditassocbutton.ID)
                }
              }
            )

            // sort FormEditAssocButtons_array array
            this.frontRepo.FormEditAssocButtons_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormFields_array = formfields

            // clear the map that counts FormField in the GET
            this.frontRepo.FormFields_batch.clear()

            formfields.forEach(
              formfield => {
                this.frontRepo.FormFields.set(formfield.ID, formfield)
                this.frontRepo.FormFields_batch.set(formfield.ID, formfield)
              }
            )

            // clear formfields that are absent from the batch
            this.frontRepo.FormFields.forEach(
              formfield => {
                if (this.frontRepo.FormFields_batch.get(formfield.ID) == undefined) {
                  this.frontRepo.FormFields.delete(formfield.ID)
                }
              }
            )

            // sort FormFields_array array
            this.frontRepo.FormFields_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormFieldDates_array = formfielddates

            // clear the map that counts FormFieldDate in the GET
            this.frontRepo.FormFieldDates_batch.clear()

            formfielddates.forEach(
              formfielddate => {
                this.frontRepo.FormFieldDates.set(formfielddate.ID, formfielddate)
                this.frontRepo.FormFieldDates_batch.set(formfielddate.ID, formfielddate)
              }
            )

            // clear formfielddates that are absent from the batch
            this.frontRepo.FormFieldDates.forEach(
              formfielddate => {
                if (this.frontRepo.FormFieldDates_batch.get(formfielddate.ID) == undefined) {
                  this.frontRepo.FormFieldDates.delete(formfielddate.ID)
                }
              }
            )

            // sort FormFieldDates_array array
            this.frontRepo.FormFieldDates_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormFieldDateTimes_array = formfielddatetimes

            // clear the map that counts FormFieldDateTime in the GET
            this.frontRepo.FormFieldDateTimes_batch.clear()

            formfielddatetimes.forEach(
              formfielddatetime => {
                this.frontRepo.FormFieldDateTimes.set(formfielddatetime.ID, formfielddatetime)
                this.frontRepo.FormFieldDateTimes_batch.set(formfielddatetime.ID, formfielddatetime)
              }
            )

            // clear formfielddatetimes that are absent from the batch
            this.frontRepo.FormFieldDateTimes.forEach(
              formfielddatetime => {
                if (this.frontRepo.FormFieldDateTimes_batch.get(formfielddatetime.ID) == undefined) {
                  this.frontRepo.FormFieldDateTimes.delete(formfielddatetime.ID)
                }
              }
            )

            // sort FormFieldDateTimes_array array
            this.frontRepo.FormFieldDateTimes_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormFieldFloat64s_array = formfieldfloat64s

            // clear the map that counts FormFieldFloat64 in the GET
            this.frontRepo.FormFieldFloat64s_batch.clear()

            formfieldfloat64s.forEach(
              formfieldfloat64 => {
                this.frontRepo.FormFieldFloat64s.set(formfieldfloat64.ID, formfieldfloat64)
                this.frontRepo.FormFieldFloat64s_batch.set(formfieldfloat64.ID, formfieldfloat64)
              }
            )

            // clear formfieldfloat64s that are absent from the batch
            this.frontRepo.FormFieldFloat64s.forEach(
              formfieldfloat64 => {
                if (this.frontRepo.FormFieldFloat64s_batch.get(formfieldfloat64.ID) == undefined) {
                  this.frontRepo.FormFieldFloat64s.delete(formfieldfloat64.ID)
                }
              }
            )

            // sort FormFieldFloat64s_array array
            this.frontRepo.FormFieldFloat64s_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormFieldInts_array = formfieldints

            // clear the map that counts FormFieldInt in the GET
            this.frontRepo.FormFieldInts_batch.clear()

            formfieldints.forEach(
              formfieldint => {
                this.frontRepo.FormFieldInts.set(formfieldint.ID, formfieldint)
                this.frontRepo.FormFieldInts_batch.set(formfieldint.ID, formfieldint)
              }
            )

            // clear formfieldints that are absent from the batch
            this.frontRepo.FormFieldInts.forEach(
              formfieldint => {
                if (this.frontRepo.FormFieldInts_batch.get(formfieldint.ID) == undefined) {
                  this.frontRepo.FormFieldInts.delete(formfieldint.ID)
                }
              }
            )

            // sort FormFieldInts_array array
            this.frontRepo.FormFieldInts_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormFieldSelects_array = formfieldselects

            // clear the map that counts FormFieldSelect in the GET
            this.frontRepo.FormFieldSelects_batch.clear()

            formfieldselects.forEach(
              formfieldselect => {
                this.frontRepo.FormFieldSelects.set(formfieldselect.ID, formfieldselect)
                this.frontRepo.FormFieldSelects_batch.set(formfieldselect.ID, formfieldselect)
              }
            )

            // clear formfieldselects that are absent from the batch
            this.frontRepo.FormFieldSelects.forEach(
              formfieldselect => {
                if (this.frontRepo.FormFieldSelects_batch.get(formfieldselect.ID) == undefined) {
                  this.frontRepo.FormFieldSelects.delete(formfieldselect.ID)
                }
              }
            )

            // sort FormFieldSelects_array array
            this.frontRepo.FormFieldSelects_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormFieldStrings_array = formfieldstrings

            // clear the map that counts FormFieldString in the GET
            this.frontRepo.FormFieldStrings_batch.clear()

            formfieldstrings.forEach(
              formfieldstring => {
                this.frontRepo.FormFieldStrings.set(formfieldstring.ID, formfieldstring)
                this.frontRepo.FormFieldStrings_batch.set(formfieldstring.ID, formfieldstring)
              }
            )

            // clear formfieldstrings that are absent from the batch
            this.frontRepo.FormFieldStrings.forEach(
              formfieldstring => {
                if (this.frontRepo.FormFieldStrings_batch.get(formfieldstring.ID) == undefined) {
                  this.frontRepo.FormFieldStrings.delete(formfieldstring.ID)
                }
              }
            )

            // sort FormFieldStrings_array array
            this.frontRepo.FormFieldStrings_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormFieldTimes_array = formfieldtimes

            // clear the map that counts FormFieldTime in the GET
            this.frontRepo.FormFieldTimes_batch.clear()

            formfieldtimes.forEach(
              formfieldtime => {
                this.frontRepo.FormFieldTimes.set(formfieldtime.ID, formfieldtime)
                this.frontRepo.FormFieldTimes_batch.set(formfieldtime.ID, formfieldtime)
              }
            )

            // clear formfieldtimes that are absent from the batch
            this.frontRepo.FormFieldTimes.forEach(
              formfieldtime => {
                if (this.frontRepo.FormFieldTimes_batch.get(formfieldtime.ID) == undefined) {
                  this.frontRepo.FormFieldTimes.delete(formfieldtime.ID)
                }
              }
            )

            // sort FormFieldTimes_array array
            this.frontRepo.FormFieldTimes_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormGroups_array = formgroups

            // clear the map that counts FormGroup in the GET
            this.frontRepo.FormGroups_batch.clear()

            formgroups.forEach(
              formgroup => {
                this.frontRepo.FormGroups.set(formgroup.ID, formgroup)
                this.frontRepo.FormGroups_batch.set(formgroup.ID, formgroup)
              }
            )

            // clear formgroups that are absent from the batch
            this.frontRepo.FormGroups.forEach(
              formgroup => {
                if (this.frontRepo.FormGroups_batch.get(formgroup.ID) == undefined) {
                  this.frontRepo.FormGroups.delete(formgroup.ID)
                }
              }
            )

            // sort FormGroups_array array
            this.frontRepo.FormGroups_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.FormSortAssocButtons_array = formsortassocbuttons

            // clear the map that counts FormSortAssocButton in the GET
            this.frontRepo.FormSortAssocButtons_batch.clear()

            formsortassocbuttons.forEach(
              formsortassocbutton => {
                this.frontRepo.FormSortAssocButtons.set(formsortassocbutton.ID, formsortassocbutton)
                this.frontRepo.FormSortAssocButtons_batch.set(formsortassocbutton.ID, formsortassocbutton)
              }
            )

            // clear formsortassocbuttons that are absent from the batch
            this.frontRepo.FormSortAssocButtons.forEach(
              formsortassocbutton => {
                if (this.frontRepo.FormSortAssocButtons_batch.get(formsortassocbutton.ID) == undefined) {
                  this.frontRepo.FormSortAssocButtons.delete(formsortassocbutton.ID)
                }
              }
            )

            // sort FormSortAssocButtons_array array
            this.frontRepo.FormSortAssocButtons_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Options_array = options

            // clear the map that counts Option in the GET
            this.frontRepo.Options_batch.clear()

            options.forEach(
              option => {
                this.frontRepo.Options.set(option.ID, option)
                this.frontRepo.Options_batch.set(option.ID, option)
              }
            )

            // clear options that are absent from the batch
            this.frontRepo.Options.forEach(
              option => {
                if (this.frontRepo.Options_batch.get(option.ID) == undefined) {
                  this.frontRepo.Options.delete(option.ID)
                }
              }
            )

            // sort Options_array array
            this.frontRepo.Options_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            checkboxs.forEach(
              checkbox => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field FormDiv.CheckBoxs redeeming
                {
                  let _formdiv = this.frontRepo.FormDivs.get(checkbox.FormDiv_CheckBoxsDBID.Int64)
                  if (_formdiv) {
                    if (_formdiv.CheckBoxs == undefined) {
                      _formdiv.CheckBoxs = new Array<CheckBoxDB>()
                    }
                    _formdiv.CheckBoxs.push(checkbox)
                    if (checkbox.FormDiv_CheckBoxs_reverse == undefined) {
                      checkbox.FormDiv_CheckBoxs_reverse = _formdiv
                    }
                  }
                }
              }
            )
            formdivs.forEach(
              formdiv => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field FormEditAssocButton redeeming
                {
                  let _formeditassocbutton = this.frontRepo.FormEditAssocButtons.get(formdiv.FormEditAssocButtonID.Int64)
                  if (_formeditassocbutton) {
                    formdiv.FormEditAssocButton = _formeditassocbutton
                  }
                }
                // insertion point for pointer field FormSortAssocButton redeeming
                {
                  let _formsortassocbutton = this.frontRepo.FormSortAssocButtons.get(formdiv.FormSortAssocButtonID.Int64)
                  if (_formsortassocbutton) {
                    formdiv.FormSortAssocButton = _formsortassocbutton
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field FormGroup.FormDivs redeeming
                {
                  let _formgroup = this.frontRepo.FormGroups.get(formdiv.FormGroup_FormDivsDBID.Int64)
                  if (_formgroup) {
                    if (_formgroup.FormDivs == undefined) {
                      _formgroup.FormDivs = new Array<FormDivDB>()
                    }
                    _formgroup.FormDivs.push(formdiv)
                    if (formdiv.FormGroup_FormDivs_reverse == undefined) {
                      formdiv.FormGroup_FormDivs_reverse = _formgroup
                    }
                  }
                }
              }
            )
            formeditassocbuttons.forEach(
              formeditassocbutton => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formfields.forEach(
              formfield => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field FormFieldString redeeming
                {
                  let _formfieldstring = this.frontRepo.FormFieldStrings.get(formfield.FormFieldStringID.Int64)
                  if (_formfieldstring) {
                    formfield.FormFieldString = _formfieldstring
                  }
                }
                // insertion point for pointer field FormFieldFloat64 redeeming
                {
                  let _formfieldfloat64 = this.frontRepo.FormFieldFloat64s.get(formfield.FormFieldFloat64ID.Int64)
                  if (_formfieldfloat64) {
                    formfield.FormFieldFloat64 = _formfieldfloat64
                  }
                }
                // insertion point for pointer field FormFieldInt redeeming
                {
                  let _formfieldint = this.frontRepo.FormFieldInts.get(formfield.FormFieldIntID.Int64)
                  if (_formfieldint) {
                    formfield.FormFieldInt = _formfieldint
                  }
                }
                // insertion point for pointer field FormFieldDate redeeming
                {
                  let _formfielddate = this.frontRepo.FormFieldDates.get(formfield.FormFieldDateID.Int64)
                  if (_formfielddate) {
                    formfield.FormFieldDate = _formfielddate
                  }
                }
                // insertion point for pointer field FormFieldTime redeeming
                {
                  let _formfieldtime = this.frontRepo.FormFieldTimes.get(formfield.FormFieldTimeID.Int64)
                  if (_formfieldtime) {
                    formfield.FormFieldTime = _formfieldtime
                  }
                }
                // insertion point for pointer field FormFieldDateTime redeeming
                {
                  let _formfielddatetime = this.frontRepo.FormFieldDateTimes.get(formfield.FormFieldDateTimeID.Int64)
                  if (_formfielddatetime) {
                    formfield.FormFieldDateTime = _formfielddatetime
                  }
                }
                // insertion point for pointer field FormFieldSelect redeeming
                {
                  let _formfieldselect = this.frontRepo.FormFieldSelects.get(formfield.FormFieldSelectID.Int64)
                  if (_formfieldselect) {
                    formfield.FormFieldSelect = _formfieldselect
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field FormDiv.FormFields redeeming
                {
                  let _formdiv = this.frontRepo.FormDivs.get(formfield.FormDiv_FormFieldsDBID.Int64)
                  if (_formdiv) {
                    if (_formdiv.FormFields == undefined) {
                      _formdiv.FormFields = new Array<FormFieldDB>()
                    }
                    _formdiv.FormFields.push(formfield)
                    if (formfield.FormDiv_FormFields_reverse == undefined) {
                      formfield.FormDiv_FormFields_reverse = _formdiv
                    }
                  }
                }
              }
            )
            formfielddates.forEach(
              formfielddate => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formfielddatetimes.forEach(
              formfielddatetime => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formfieldfloat64s.forEach(
              formfieldfloat64 => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formfieldints.forEach(
              formfieldint => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formfieldselects.forEach(
              formfieldselect => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Value redeeming
                {
                  let _option = this.frontRepo.Options.get(formfieldselect.ValueID.Int64)
                  if (_option) {
                    formfieldselect.Value = _option
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formfieldstrings.forEach(
              formfieldstring => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formfieldtimes.forEach(
              formfieldtime => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formgroups.forEach(
              formgroup => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            formsortassocbuttons.forEach(
              formsortassocbutton => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            options.forEach(
              option => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field FormFieldSelect.Options redeeming
                {
                  let _formfieldselect = this.frontRepo.FormFieldSelects.get(option.FormFieldSelect_OptionsDBID.Int64)
                  if (_formfieldselect) {
                    if (_formfieldselect.Options == undefined) {
                      _formfieldselect.Options = new Array<OptionDB>()
                    }
                    _formfieldselect.Options.push(option)
                    if (option.FormFieldSelect_Options_reverse == undefined) {
                      option.FormFieldSelect_Options_reverse = _formfieldselect
                    }
                  }
                }
              }
            )

            // 
            // Third Step: sort arrays (slices in go) according to their index
            // insertion point sub template for redeem 
            checkboxs.forEach(
              checkbox => {
                // insertion point for sorting
              }
            )
            formdivs.forEach(
              formdiv => {
                // insertion point for sorting
                formdiv.FormFields?.sort((t1, t2) => {
                  if (t1.FormDiv_FormFieldsDBID_Index.Int64 > t2.FormDiv_FormFieldsDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.FormDiv_FormFieldsDBID_Index.Int64 < t2.FormDiv_FormFieldsDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

                formdiv.CheckBoxs?.sort((t1, t2) => {
                  if (t1.FormDiv_CheckBoxsDBID_Index.Int64 > t2.FormDiv_CheckBoxsDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.FormDiv_CheckBoxsDBID_Index.Int64 < t2.FormDiv_CheckBoxsDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

              }
            )
            formeditassocbuttons.forEach(
              formeditassocbutton => {
                // insertion point for sorting
              }
            )
            formfields.forEach(
              formfield => {
                // insertion point for sorting
              }
            )
            formfielddates.forEach(
              formfielddate => {
                // insertion point for sorting
              }
            )
            formfielddatetimes.forEach(
              formfielddatetime => {
                // insertion point for sorting
              }
            )
            formfieldfloat64s.forEach(
              formfieldfloat64 => {
                // insertion point for sorting
              }
            )
            formfieldints.forEach(
              formfieldint => {
                // insertion point for sorting
              }
            )
            formfieldselects.forEach(
              formfieldselect => {
                // insertion point for sorting
                formfieldselect.Options?.sort((t1, t2) => {
                  if (t1.FormFieldSelect_OptionsDBID_Index.Int64 > t2.FormFieldSelect_OptionsDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.FormFieldSelect_OptionsDBID_Index.Int64 < t2.FormFieldSelect_OptionsDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

              }
            )
            formfieldstrings.forEach(
              formfieldstring => {
                // insertion point for sorting
              }
            )
            formfieldtimes.forEach(
              formfieldtime => {
                // insertion point for sorting
              }
            )
            formgroups.forEach(
              formgroup => {
                // insertion point for sorting
                formgroup.FormDivs?.sort((t1, t2) => {
                  if (t1.FormGroup_FormDivsDBID_Index.Int64 > t2.FormGroup_FormDivsDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.FormGroup_FormDivsDBID_Index.Int64 < t2.FormGroup_FormDivsDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

              }
            )
            formsortassocbuttons.forEach(
              formsortassocbutton => {
                // insertion point for sorting
              }
            )
            options.forEach(
              option => {
                // insertion point for sorting
              }
            )

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // CheckBoxPull performs a GET on CheckBox of the stack and redeem association pointers 
  CheckBoxPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.checkboxService.getCheckBoxs(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            checkboxs,
          ]) => {
            // init the array
            this.frontRepo.CheckBoxs_array = checkboxs

            // clear the map that counts CheckBox in the GET
            this.frontRepo.CheckBoxs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            checkboxs.forEach(
              checkbox => {
                this.frontRepo.CheckBoxs.set(checkbox.ID, checkbox)
                this.frontRepo.CheckBoxs_batch.set(checkbox.ID, checkbox)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field FormDiv.CheckBoxs redeeming
                {
                  let _formdiv = this.frontRepo.FormDivs.get(checkbox.FormDiv_CheckBoxsDBID.Int64)
                  if (_formdiv) {
                    if (_formdiv.CheckBoxs == undefined) {
                      _formdiv.CheckBoxs = new Array<CheckBoxDB>()
                    }
                    _formdiv.CheckBoxs.push(checkbox)
                    if (checkbox.FormDiv_CheckBoxs_reverse == undefined) {
                      checkbox.FormDiv_CheckBoxs_reverse = _formdiv
                    }
                  }
                }
              }
            )

            // clear checkboxs that are absent from the GET
            this.frontRepo.CheckBoxs.forEach(
              checkbox => {
                if (this.frontRepo.CheckBoxs_batch.get(checkbox.ID) == undefined) {
                  this.frontRepo.CheckBoxs.delete(checkbox.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // FormDivPull performs a GET on FormDiv of the stack and redeem association pointers 
  FormDivPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formdivService.getFormDivs(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formdivs,
          ]) => {
            // init the array
            this.frontRepo.FormDivs_array = formdivs

            // clear the map that counts FormDiv in the GET
            this.frontRepo.FormDivs_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formdivs.forEach(
              formdiv => {
                this.frontRepo.FormDivs.set(formdiv.ID, formdiv)
                this.frontRepo.FormDivs_batch.set(formdiv.ID, formdiv)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field FormEditAssocButton redeeming
                {
                  let _formeditassocbutton = this.frontRepo.FormEditAssocButtons.get(formdiv.FormEditAssocButtonID.Int64)
                  if (_formeditassocbutton) {
                    formdiv.FormEditAssocButton = _formeditassocbutton
                  }
                }
                // insertion point for pointer field FormSortAssocButton redeeming
                {
                  let _formsortassocbutton = this.frontRepo.FormSortAssocButtons.get(formdiv.FormSortAssocButtonID.Int64)
                  if (_formsortassocbutton) {
                    formdiv.FormSortAssocButton = _formsortassocbutton
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field FormGroup.FormDivs redeeming
                {
                  let _formgroup = this.frontRepo.FormGroups.get(formdiv.FormGroup_FormDivsDBID.Int64)
                  if (_formgroup) {
                    if (_formgroup.FormDivs == undefined) {
                      _formgroup.FormDivs = new Array<FormDivDB>()
                    }
                    _formgroup.FormDivs.push(formdiv)
                    if (formdiv.FormGroup_FormDivs_reverse == undefined) {
                      formdiv.FormGroup_FormDivs_reverse = _formgroup
                    }
                  }
                }
              }
            )

            // clear formdivs that are absent from the GET
            this.frontRepo.FormDivs.forEach(
              formdiv => {
                if (this.frontRepo.FormDivs_batch.get(formdiv.ID) == undefined) {
                  this.frontRepo.FormDivs.delete(formdiv.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // FormEditAssocButtonPull performs a GET on FormEditAssocButton of the stack and redeem association pointers 
  FormEditAssocButtonPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formeditassocbuttonService.getFormEditAssocButtons(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formeditassocbuttons,
          ]) => {
            // init the array
            this.frontRepo.FormEditAssocButtons_array = formeditassocbuttons

            // clear the map that counts FormEditAssocButton in the GET
            this.frontRepo.FormEditAssocButtons_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formeditassocbuttons.forEach(
              formeditassocbutton => {
                this.frontRepo.FormEditAssocButtons.set(formeditassocbutton.ID, formeditassocbutton)
                this.frontRepo.FormEditAssocButtons_batch.set(formeditassocbutton.ID, formeditassocbutton)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formeditassocbuttons that are absent from the GET
            this.frontRepo.FormEditAssocButtons.forEach(
              formeditassocbutton => {
                if (this.frontRepo.FormEditAssocButtons_batch.get(formeditassocbutton.ID) == undefined) {
                  this.frontRepo.FormEditAssocButtons.delete(formeditassocbutton.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // FormFieldPull performs a GET on FormField of the stack and redeem association pointers 
  FormFieldPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formfieldService.getFormFields(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formfields,
          ]) => {
            // init the array
            this.frontRepo.FormFields_array = formfields

            // clear the map that counts FormField in the GET
            this.frontRepo.FormFields_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formfields.forEach(
              formfield => {
                this.frontRepo.FormFields.set(formfield.ID, formfield)
                this.frontRepo.FormFields_batch.set(formfield.ID, formfield)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field FormFieldString redeeming
                {
                  let _formfieldstring = this.frontRepo.FormFieldStrings.get(formfield.FormFieldStringID.Int64)
                  if (_formfieldstring) {
                    formfield.FormFieldString = _formfieldstring
                  }
                }
                // insertion point for pointer field FormFieldFloat64 redeeming
                {
                  let _formfieldfloat64 = this.frontRepo.FormFieldFloat64s.get(formfield.FormFieldFloat64ID.Int64)
                  if (_formfieldfloat64) {
                    formfield.FormFieldFloat64 = _formfieldfloat64
                  }
                }
                // insertion point for pointer field FormFieldInt redeeming
                {
                  let _formfieldint = this.frontRepo.FormFieldInts.get(formfield.FormFieldIntID.Int64)
                  if (_formfieldint) {
                    formfield.FormFieldInt = _formfieldint
                  }
                }
                // insertion point for pointer field FormFieldDate redeeming
                {
                  let _formfielddate = this.frontRepo.FormFieldDates.get(formfield.FormFieldDateID.Int64)
                  if (_formfielddate) {
                    formfield.FormFieldDate = _formfielddate
                  }
                }
                // insertion point for pointer field FormFieldTime redeeming
                {
                  let _formfieldtime = this.frontRepo.FormFieldTimes.get(formfield.FormFieldTimeID.Int64)
                  if (_formfieldtime) {
                    formfield.FormFieldTime = _formfieldtime
                  }
                }
                // insertion point for pointer field FormFieldDateTime redeeming
                {
                  let _formfielddatetime = this.frontRepo.FormFieldDateTimes.get(formfield.FormFieldDateTimeID.Int64)
                  if (_formfielddatetime) {
                    formfield.FormFieldDateTime = _formfielddatetime
                  }
                }
                // insertion point for pointer field FormFieldSelect redeeming
                {
                  let _formfieldselect = this.frontRepo.FormFieldSelects.get(formfield.FormFieldSelectID.Int64)
                  if (_formfieldselect) {
                    formfield.FormFieldSelect = _formfieldselect
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field FormDiv.FormFields redeeming
                {
                  let _formdiv = this.frontRepo.FormDivs.get(formfield.FormDiv_FormFieldsDBID.Int64)
                  if (_formdiv) {
                    if (_formdiv.FormFields == undefined) {
                      _formdiv.FormFields = new Array<FormFieldDB>()
                    }
                    _formdiv.FormFields.push(formfield)
                    if (formfield.FormDiv_FormFields_reverse == undefined) {
                      formfield.FormDiv_FormFields_reverse = _formdiv
                    }
                  }
                }
              }
            )

            // clear formfields that are absent from the GET
            this.frontRepo.FormFields.forEach(
              formfield => {
                if (this.frontRepo.FormFields_batch.get(formfield.ID) == undefined) {
                  this.frontRepo.FormFields.delete(formfield.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // FormFieldDatePull performs a GET on FormFieldDate of the stack and redeem association pointers 
  FormFieldDatePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formfielddateService.getFormFieldDates(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formfielddates,
          ]) => {
            // init the array
            this.frontRepo.FormFieldDates_array = formfielddates

            // clear the map that counts FormFieldDate in the GET
            this.frontRepo.FormFieldDates_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formfielddates.forEach(
              formfielddate => {
                this.frontRepo.FormFieldDates.set(formfielddate.ID, formfielddate)
                this.frontRepo.FormFieldDates_batch.set(formfielddate.ID, formfielddate)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formfielddates that are absent from the GET
            this.frontRepo.FormFieldDates.forEach(
              formfielddate => {
                if (this.frontRepo.FormFieldDates_batch.get(formfielddate.ID) == undefined) {
                  this.frontRepo.FormFieldDates.delete(formfielddate.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // FormFieldDateTimePull performs a GET on FormFieldDateTime of the stack and redeem association pointers 
  FormFieldDateTimePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formfielddatetimeService.getFormFieldDateTimes(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formfielddatetimes,
          ]) => {
            // init the array
            this.frontRepo.FormFieldDateTimes_array = formfielddatetimes

            // clear the map that counts FormFieldDateTime in the GET
            this.frontRepo.FormFieldDateTimes_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formfielddatetimes.forEach(
              formfielddatetime => {
                this.frontRepo.FormFieldDateTimes.set(formfielddatetime.ID, formfielddatetime)
                this.frontRepo.FormFieldDateTimes_batch.set(formfielddatetime.ID, formfielddatetime)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formfielddatetimes that are absent from the GET
            this.frontRepo.FormFieldDateTimes.forEach(
              formfielddatetime => {
                if (this.frontRepo.FormFieldDateTimes_batch.get(formfielddatetime.ID) == undefined) {
                  this.frontRepo.FormFieldDateTimes.delete(formfielddatetime.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // FormFieldFloat64Pull performs a GET on FormFieldFloat64 of the stack and redeem association pointers 
  FormFieldFloat64Pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formfieldfloat64Service.getFormFieldFloat64s(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formfieldfloat64s,
          ]) => {
            // init the array
            this.frontRepo.FormFieldFloat64s_array = formfieldfloat64s

            // clear the map that counts FormFieldFloat64 in the GET
            this.frontRepo.FormFieldFloat64s_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formfieldfloat64s.forEach(
              formfieldfloat64 => {
                this.frontRepo.FormFieldFloat64s.set(formfieldfloat64.ID, formfieldfloat64)
                this.frontRepo.FormFieldFloat64s_batch.set(formfieldfloat64.ID, formfieldfloat64)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formfieldfloat64s that are absent from the GET
            this.frontRepo.FormFieldFloat64s.forEach(
              formfieldfloat64 => {
                if (this.frontRepo.FormFieldFloat64s_batch.get(formfieldfloat64.ID) == undefined) {
                  this.frontRepo.FormFieldFloat64s.delete(formfieldfloat64.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // FormFieldIntPull performs a GET on FormFieldInt of the stack and redeem association pointers 
  FormFieldIntPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formfieldintService.getFormFieldInts(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formfieldints,
          ]) => {
            // init the array
            this.frontRepo.FormFieldInts_array = formfieldints

            // clear the map that counts FormFieldInt in the GET
            this.frontRepo.FormFieldInts_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formfieldints.forEach(
              formfieldint => {
                this.frontRepo.FormFieldInts.set(formfieldint.ID, formfieldint)
                this.frontRepo.FormFieldInts_batch.set(formfieldint.ID, formfieldint)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formfieldints that are absent from the GET
            this.frontRepo.FormFieldInts.forEach(
              formfieldint => {
                if (this.frontRepo.FormFieldInts_batch.get(formfieldint.ID) == undefined) {
                  this.frontRepo.FormFieldInts.delete(formfieldint.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // FormFieldSelectPull performs a GET on FormFieldSelect of the stack and redeem association pointers 
  FormFieldSelectPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formfieldselectService.getFormFieldSelects(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formfieldselects,
          ]) => {
            // init the array
            this.frontRepo.FormFieldSelects_array = formfieldselects

            // clear the map that counts FormFieldSelect in the GET
            this.frontRepo.FormFieldSelects_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formfieldselects.forEach(
              formfieldselect => {
                this.frontRepo.FormFieldSelects.set(formfieldselect.ID, formfieldselect)
                this.frontRepo.FormFieldSelects_batch.set(formfieldselect.ID, formfieldselect)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Value redeeming
                {
                  let _option = this.frontRepo.Options.get(formfieldselect.ValueID.Int64)
                  if (_option) {
                    formfieldselect.Value = _option
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formfieldselects that are absent from the GET
            this.frontRepo.FormFieldSelects.forEach(
              formfieldselect => {
                if (this.frontRepo.FormFieldSelects_batch.get(formfieldselect.ID) == undefined) {
                  this.frontRepo.FormFieldSelects.delete(formfieldselect.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // FormFieldStringPull performs a GET on FormFieldString of the stack and redeem association pointers 
  FormFieldStringPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formfieldstringService.getFormFieldStrings(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formfieldstrings,
          ]) => {
            // init the array
            this.frontRepo.FormFieldStrings_array = formfieldstrings

            // clear the map that counts FormFieldString in the GET
            this.frontRepo.FormFieldStrings_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formfieldstrings.forEach(
              formfieldstring => {
                this.frontRepo.FormFieldStrings.set(formfieldstring.ID, formfieldstring)
                this.frontRepo.FormFieldStrings_batch.set(formfieldstring.ID, formfieldstring)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formfieldstrings that are absent from the GET
            this.frontRepo.FormFieldStrings.forEach(
              formfieldstring => {
                if (this.frontRepo.FormFieldStrings_batch.get(formfieldstring.ID) == undefined) {
                  this.frontRepo.FormFieldStrings.delete(formfieldstring.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // FormFieldTimePull performs a GET on FormFieldTime of the stack and redeem association pointers 
  FormFieldTimePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formfieldtimeService.getFormFieldTimes(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formfieldtimes,
          ]) => {
            // init the array
            this.frontRepo.FormFieldTimes_array = formfieldtimes

            // clear the map that counts FormFieldTime in the GET
            this.frontRepo.FormFieldTimes_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formfieldtimes.forEach(
              formfieldtime => {
                this.frontRepo.FormFieldTimes.set(formfieldtime.ID, formfieldtime)
                this.frontRepo.FormFieldTimes_batch.set(formfieldtime.ID, formfieldtime)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formfieldtimes that are absent from the GET
            this.frontRepo.FormFieldTimes.forEach(
              formfieldtime => {
                if (this.frontRepo.FormFieldTimes_batch.get(formfieldtime.ID) == undefined) {
                  this.frontRepo.FormFieldTimes.delete(formfieldtime.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // FormGroupPull performs a GET on FormGroup of the stack and redeem association pointers 
  FormGroupPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formgroupService.getFormGroups(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formgroups,
          ]) => {
            // init the array
            this.frontRepo.FormGroups_array = formgroups

            // clear the map that counts FormGroup in the GET
            this.frontRepo.FormGroups_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formgroups.forEach(
              formgroup => {
                this.frontRepo.FormGroups.set(formgroup.ID, formgroup)
                this.frontRepo.FormGroups_batch.set(formgroup.ID, formgroup)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formgroups that are absent from the GET
            this.frontRepo.FormGroups.forEach(
              formgroup => {
                if (this.frontRepo.FormGroups_batch.get(formgroup.ID) == undefined) {
                  this.frontRepo.FormGroups.delete(formgroup.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // FormSortAssocButtonPull performs a GET on FormSortAssocButton of the stack and redeem association pointers 
  FormSortAssocButtonPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.formsortassocbuttonService.getFormSortAssocButtons(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            formsortassocbuttons,
          ]) => {
            // init the array
            this.frontRepo.FormSortAssocButtons_array = formsortassocbuttons

            // clear the map that counts FormSortAssocButton in the GET
            this.frontRepo.FormSortAssocButtons_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            formsortassocbuttons.forEach(
              formsortassocbutton => {
                this.frontRepo.FormSortAssocButtons.set(formsortassocbutton.ID, formsortassocbutton)
                this.frontRepo.FormSortAssocButtons_batch.set(formsortassocbutton.ID, formsortassocbutton)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear formsortassocbuttons that are absent from the GET
            this.frontRepo.FormSortAssocButtons.forEach(
              formsortassocbutton => {
                if (this.frontRepo.FormSortAssocButtons_batch.get(formsortassocbutton.ID) == undefined) {
                  this.frontRepo.FormSortAssocButtons.delete(formsortassocbutton.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // OptionPull performs a GET on Option of the stack and redeem association pointers 
  OptionPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.optionService.getOptions(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            options,
          ]) => {
            // init the array
            this.frontRepo.Options_array = options

            // clear the map that counts Option in the GET
            this.frontRepo.Options_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            options.forEach(
              option => {
                this.frontRepo.Options.set(option.ID, option)
                this.frontRepo.Options_batch.set(option.ID, option)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field FormFieldSelect.Options redeeming
                {
                  let _formfieldselect = this.frontRepo.FormFieldSelects.get(option.FormFieldSelect_OptionsDBID.Int64)
                  if (_formfieldselect) {
                    if (_formfieldselect.Options == undefined) {
                      _formfieldselect.Options = new Array<OptionDB>()
                    }
                    _formfieldselect.Options.push(option)
                    if (option.FormFieldSelect_Options_reverse == undefined) {
                      option.FormFieldSelect_Options_reverse = _formfieldselect
                    }
                  }
                }
              }
            )

            // clear options that are absent from the GET
            this.frontRepo.Options.forEach(
              option => {
                if (this.frontRepo.Options_batch.get(option.ID) == undefined) {
                  this.frontRepo.Options.delete(option.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getCheckBoxUniqueID(id: number): number {
  return 31 * id
}
export function getFormDivUniqueID(id: number): number {
  return 37 * id
}
export function getFormEditAssocButtonUniqueID(id: number): number {
  return 41 * id
}
export function getFormFieldUniqueID(id: number): number {
  return 43 * id
}
export function getFormFieldDateUniqueID(id: number): number {
  return 47 * id
}
export function getFormFieldDateTimeUniqueID(id: number): number {
  return 53 * id
}
export function getFormFieldFloat64UniqueID(id: number): number {
  return 59 * id
}
export function getFormFieldIntUniqueID(id: number): number {
  return 61 * id
}
export function getFormFieldSelectUniqueID(id: number): number {
  return 67 * id
}
export function getFormFieldStringUniqueID(id: number): number {
  return 71 * id
}
export function getFormFieldTimeUniqueID(id: number): number {
  return 73 * id
}
export function getFormGroupUniqueID(id: number): number {
  return 79 * id
}
export function getFormSortAssocButtonUniqueID(id: number): number {
  return 83 * id
}
export function getOptionUniqueID(id: number): number {
  return 89 * id
}
