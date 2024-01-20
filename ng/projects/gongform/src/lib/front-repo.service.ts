import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { CheckBoxDB } from './checkbox-db'
import { CheckBox, CopyCheckBoxDBToCheckBox } from './checkbox'
import { CheckBoxService } from './checkbox.service'

import { FormDivDB } from './formdiv-db'
import { FormDiv, CopyFormDivDBToFormDiv } from './formdiv'
import { FormDivService } from './formdiv.service'

import { FormEditAssocButtonDB } from './formeditassocbutton-db'
import { FormEditAssocButton, CopyFormEditAssocButtonDBToFormEditAssocButton } from './formeditassocbutton'
import { FormEditAssocButtonService } from './formeditassocbutton.service'

import { FormFieldDB } from './formfield-db'
import { FormField, CopyFormFieldDBToFormField } from './formfield'
import { FormFieldService } from './formfield.service'

import { FormFieldDateDB } from './formfielddate-db'
import { FormFieldDate, CopyFormFieldDateDBToFormFieldDate } from './formfielddate'
import { FormFieldDateService } from './formfielddate.service'

import { FormFieldDateTimeDB } from './formfielddatetime-db'
import { FormFieldDateTime, CopyFormFieldDateTimeDBToFormFieldDateTime } from './formfielddatetime'
import { FormFieldDateTimeService } from './formfielddatetime.service'

import { FormFieldFloat64DB } from './formfieldfloat64-db'
import { FormFieldFloat64, CopyFormFieldFloat64DBToFormFieldFloat64 } from './formfieldfloat64'
import { FormFieldFloat64Service } from './formfieldfloat64.service'

import { FormFieldIntDB } from './formfieldint-db'
import { FormFieldInt, CopyFormFieldIntDBToFormFieldInt } from './formfieldint'
import { FormFieldIntService } from './formfieldint.service'

import { FormFieldSelectDB } from './formfieldselect-db'
import { FormFieldSelect, CopyFormFieldSelectDBToFormFieldSelect } from './formfieldselect'
import { FormFieldSelectService } from './formfieldselect.service'

import { FormFieldStringDB } from './formfieldstring-db'
import { FormFieldString, CopyFormFieldStringDBToFormFieldString } from './formfieldstring'
import { FormFieldStringService } from './formfieldstring.service'

import { FormFieldTimeDB } from './formfieldtime-db'
import { FormFieldTime, CopyFormFieldTimeDBToFormFieldTime } from './formfieldtime'
import { FormFieldTimeService } from './formfieldtime.service'

import { FormGroupDB } from './formgroup-db'
import { FormGroup, CopyFormGroupDBToFormGroup } from './formgroup'
import { FormGroupService } from './formgroup.service'

import { FormSortAssocButtonDB } from './formsortassocbutton-db'
import { FormSortAssocButton, CopyFormSortAssocButtonDBToFormSortAssocButton } from './formsortassocbutton'
import { FormSortAssocButtonService } from './formsortassocbutton.service'

import { OptionDB } from './option-db'
import { Option, CopyOptionDBToOption } from './option'
import { OptionService } from './option.service'

export const StackType = "github.com/fullstack-lang/gongform/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
  CheckBoxs_array = new Array<CheckBoxDB>() // array of repo instances
  CheckBoxs = new Map<number, CheckBoxDB>() // map of repo instances
  CheckBoxs_batch = new Map<number, CheckBoxDB>() // same but only in last GET (for finding repo instances to delete)

  array_CheckBoxs = new Array<CheckBox>() // array of front instances
  map_ID_CheckBox = new Map<number, CheckBox>() // map of front instances

  FormDivs_array = new Array<FormDivDB>() // array of repo instances
  FormDivs = new Map<number, FormDivDB>() // map of repo instances
  FormDivs_batch = new Map<number, FormDivDB>() // same but only in last GET (for finding repo instances to delete)

  array_FormDivs = new Array<FormDiv>() // array of front instances
  map_ID_FormDiv = new Map<number, FormDiv>() // map of front instances

  FormEditAssocButtons_array = new Array<FormEditAssocButtonDB>() // array of repo instances
  FormEditAssocButtons = new Map<number, FormEditAssocButtonDB>() // map of repo instances
  FormEditAssocButtons_batch = new Map<number, FormEditAssocButtonDB>() // same but only in last GET (for finding repo instances to delete)

  array_FormEditAssocButtons = new Array<FormEditAssocButton>() // array of front instances
  map_ID_FormEditAssocButton = new Map<number, FormEditAssocButton>() // map of front instances

  FormFields_array = new Array<FormFieldDB>() // array of repo instances
  FormFields = new Map<number, FormFieldDB>() // map of repo instances
  FormFields_batch = new Map<number, FormFieldDB>() // same but only in last GET (for finding repo instances to delete)

  array_FormFields = new Array<FormField>() // array of front instances
  map_ID_FormField = new Map<number, FormField>() // map of front instances

  FormFieldDates_array = new Array<FormFieldDateDB>() // array of repo instances
  FormFieldDates = new Map<number, FormFieldDateDB>() // map of repo instances
  FormFieldDates_batch = new Map<number, FormFieldDateDB>() // same but only in last GET (for finding repo instances to delete)

  array_FormFieldDates = new Array<FormFieldDate>() // array of front instances
  map_ID_FormFieldDate = new Map<number, FormFieldDate>() // map of front instances

  FormFieldDateTimes_array = new Array<FormFieldDateTimeDB>() // array of repo instances
  FormFieldDateTimes = new Map<number, FormFieldDateTimeDB>() // map of repo instances
  FormFieldDateTimes_batch = new Map<number, FormFieldDateTimeDB>() // same but only in last GET (for finding repo instances to delete)

  array_FormFieldDateTimes = new Array<FormFieldDateTime>() // array of front instances
  map_ID_FormFieldDateTime = new Map<number, FormFieldDateTime>() // map of front instances

  FormFieldFloat64s_array = new Array<FormFieldFloat64DB>() // array of repo instances
  FormFieldFloat64s = new Map<number, FormFieldFloat64DB>() // map of repo instances
  FormFieldFloat64s_batch = new Map<number, FormFieldFloat64DB>() // same but only in last GET (for finding repo instances to delete)

  array_FormFieldFloat64s = new Array<FormFieldFloat64>() // array of front instances
  map_ID_FormFieldFloat64 = new Map<number, FormFieldFloat64>() // map of front instances

  FormFieldInts_array = new Array<FormFieldIntDB>() // array of repo instances
  FormFieldInts = new Map<number, FormFieldIntDB>() // map of repo instances
  FormFieldInts_batch = new Map<number, FormFieldIntDB>() // same but only in last GET (for finding repo instances to delete)

  array_FormFieldInts = new Array<FormFieldInt>() // array of front instances
  map_ID_FormFieldInt = new Map<number, FormFieldInt>() // map of front instances

  FormFieldSelects_array = new Array<FormFieldSelectDB>() // array of repo instances
  FormFieldSelects = new Map<number, FormFieldSelectDB>() // map of repo instances
  FormFieldSelects_batch = new Map<number, FormFieldSelectDB>() // same but only in last GET (for finding repo instances to delete)

  array_FormFieldSelects = new Array<FormFieldSelect>() // array of front instances
  map_ID_FormFieldSelect = new Map<number, FormFieldSelect>() // map of front instances

  FormFieldStrings_array = new Array<FormFieldStringDB>() // array of repo instances
  FormFieldStrings = new Map<number, FormFieldStringDB>() // map of repo instances
  FormFieldStrings_batch = new Map<number, FormFieldStringDB>() // same but only in last GET (for finding repo instances to delete)

  array_FormFieldStrings = new Array<FormFieldString>() // array of front instances
  map_ID_FormFieldString = new Map<number, FormFieldString>() // map of front instances

  FormFieldTimes_array = new Array<FormFieldTimeDB>() // array of repo instances
  FormFieldTimes = new Map<number, FormFieldTimeDB>() // map of repo instances
  FormFieldTimes_batch = new Map<number, FormFieldTimeDB>() // same but only in last GET (for finding repo instances to delete)

  array_FormFieldTimes = new Array<FormFieldTime>() // array of front instances
  map_ID_FormFieldTime = new Map<number, FormFieldTime>() // map of front instances

  FormGroups_array = new Array<FormGroupDB>() // array of repo instances
  FormGroups = new Map<number, FormGroupDB>() // map of repo instances
  FormGroups_batch = new Map<number, FormGroupDB>() // same but only in last GET (for finding repo instances to delete)

  array_FormGroups = new Array<FormGroup>() // array of front instances
  map_ID_FormGroup = new Map<number, FormGroup>() // map of front instances

  FormSortAssocButtons_array = new Array<FormSortAssocButtonDB>() // array of repo instances
  FormSortAssocButtons = new Map<number, FormSortAssocButtonDB>() // map of repo instances
  FormSortAssocButtons_batch = new Map<number, FormSortAssocButtonDB>() // same but only in last GET (for finding repo instances to delete)

  array_FormSortAssocButtons = new Array<FormSortAssocButton>() // array of front instances
  map_ID_FormSortAssocButton = new Map<number, FormSortAssocButton>() // map of front instances

  Options_array = new Array<OptionDB>() // array of repo instances
  Options = new Map<number, OptionDB>() // map of repo instances
  Options_batch = new Map<number, OptionDB>() // same but only in last GET (for finding repo instances to delete)

  array_Options = new Array<Option>() // array of front instances
  map_ID_Option = new Map<number, Option>() // map of front instances


  // getArray allows for a get function that is robust to refactoring of the named struct name
  // for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
  // contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
  getArray<Type>(gongStructName: string): Array<Type> {
    switch (gongStructName) { // deprecated
      // insertion point
      case 'CheckBox':
        return this.CheckBoxs_array as unknown as Array<Type>
      case 'FormDiv':
        return this.FormDivs_array as unknown as Array<Type>
      case 'FormEditAssocButton':
        return this.FormEditAssocButtons_array as unknown as Array<Type>
      case 'FormField':
        return this.FormFields_array as unknown as Array<Type>
      case 'FormFieldDate':
        return this.FormFieldDates_array as unknown as Array<Type>
      case 'FormFieldDateTime':
        return this.FormFieldDateTimes_array as unknown as Array<Type>
      case 'FormFieldFloat64':
        return this.FormFieldFloat64s_array as unknown as Array<Type>
      case 'FormFieldInt':
        return this.FormFieldInts_array as unknown as Array<Type>
      case 'FormFieldSelect':
        return this.FormFieldSelects_array as unknown as Array<Type>
      case 'FormFieldString':
        return this.FormFieldStrings_array as unknown as Array<Type>
      case 'FormFieldTime':
        return this.FormFieldTimes_array as unknown as Array<Type>
      case 'FormGroup':
        return this.FormGroups_array as unknown as Array<Type>
      case 'FormSortAssocButton':
        return this.FormSortAssocButtons_array as unknown as Array<Type>
      case 'Option':
        return this.Options_array as unknown as Array<Type>
      default:
        throw new Error("Type not recognized");
    }
  }

  getFrontArray<Type>(gongStructName: string): Array<Type> {
    switch (gongStructName) {
      // insertion point
      case 'CheckBox':
        return this.array_CheckBoxs as unknown as Array<Type>
      case 'FormDiv':
        return this.array_FormDivs as unknown as Array<Type>
      case 'FormEditAssocButton':
        return this.array_FormEditAssocButtons as unknown as Array<Type>
      case 'FormField':
        return this.array_FormFields as unknown as Array<Type>
      case 'FormFieldDate':
        return this.array_FormFieldDates as unknown as Array<Type>
      case 'FormFieldDateTime':
        return this.array_FormFieldDateTimes as unknown as Array<Type>
      case 'FormFieldFloat64':
        return this.array_FormFieldFloat64s as unknown as Array<Type>
      case 'FormFieldInt':
        return this.array_FormFieldInts as unknown as Array<Type>
      case 'FormFieldSelect':
        return this.array_FormFieldSelects as unknown as Array<Type>
      case 'FormFieldString':
        return this.array_FormFieldStrings as unknown as Array<Type>
      case 'FormFieldTime':
        return this.array_FormFieldTimes as unknown as Array<Type>
      case 'FormGroup':
        return this.array_FormGroups as unknown as Array<Type>
      case 'FormSortAssocButton':
        return this.array_FormSortAssocButtons as unknown as Array<Type>
      case 'Option':
        return this.array_Options as unknown as Array<Type>
      default:
        throw new Error("Type not recognized");
    }
  }

  // getMap allows for a get function that is robust to refactoring of the named struct name
  getMap<Type>(gongStructName: string): Map<number, Type> { // deprecated
    switch (gongStructName) {
      // insertion point
      case 'CheckBox':
        return this.CheckBoxs as unknown as Map<number, Type>
      case 'FormDiv':
        return this.FormDivs as unknown as Map<number, Type>
      case 'FormEditAssocButton':
        return this.FormEditAssocButtons as unknown as Map<number, Type>
      case 'FormField':
        return this.FormFields as unknown as Map<number, Type>
      case 'FormFieldDate':
        return this.FormFieldDates as unknown as Map<number, Type>
      case 'FormFieldDateTime':
        return this.FormFieldDateTimes as unknown as Map<number, Type>
      case 'FormFieldFloat64':
        return this.FormFieldFloat64s as unknown as Map<number, Type>
      case 'FormFieldInt':
        return this.FormFieldInts as unknown as Map<number, Type>
      case 'FormFieldSelect':
        return this.FormFieldSelects as unknown as Map<number, Type>
      case 'FormFieldString':
        return this.FormFieldStrings as unknown as Map<number, Type>
      case 'FormFieldTime':
        return this.FormFieldTimes as unknown as Map<number, Type>
      case 'FormGroup':
        return this.FormGroups as unknown as Map<number, Type>
      case 'FormSortAssocButton':
        return this.FormSortAssocButtons as unknown as Map<number, Type>
      case 'Option':
        return this.Options as unknown as Map<number, Type>
      default:
        throw new Error("Type not recognized");
    }
  }
  
  getFrontMap<Type>(gongStructName: string): Map<number, Type> {
    switch (gongStructName) {
      // insertion point
      case 'CheckBox':
        return this.map_ID_CheckBox as unknown as Map<number, Type>
      case 'FormDiv':
        return this.map_ID_FormDiv as unknown as Map<number, Type>
      case 'FormEditAssocButton':
        return this.map_ID_FormEditAssocButton as unknown as Map<number, Type>
      case 'FormField':
        return this.map_ID_FormField as unknown as Map<number, Type>
      case 'FormFieldDate':
        return this.map_ID_FormFieldDate as unknown as Map<number, Type>
      case 'FormFieldDateTime':
        return this.map_ID_FormFieldDateTime as unknown as Map<number, Type>
      case 'FormFieldFloat64':
        return this.map_ID_FormFieldFloat64 as unknown as Map<number, Type>
      case 'FormFieldInt':
        return this.map_ID_FormFieldInt as unknown as Map<number, Type>
      case 'FormFieldSelect':
        return this.map_ID_FormFieldSelect as unknown as Map<number, Type>
      case 'FormFieldString':
        return this.map_ID_FormFieldString as unknown as Map<number, Type>
      case 'FormFieldTime':
        return this.map_ID_FormFieldTime as unknown as Map<number, Type>
      case 'FormGroup':
        return this.map_ID_FormGroup as unknown as Map<number, Type>
      case 'FormSortAssocButton':
        return this.map_ID_FormSortAssocButton as unknown as Map<number, Type>
      case 'Option':
        return this.map_ID_Option as unknown as Map<number, Type>
      default:
        throw new Error("Type not recognized");
    }
  }
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
  observableFrontRepo: [
    Observable<null>, // see below for the of(null) observable
    // insertion point sub template 
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
  ] = [
      // Using "combineLatest" with a placeholder observable.
      //
      // This allows the typescript compiler to pass when no GongStruct is present in the front API
      //
      // The "of(null)" is a "meaningless" observable that emits a single value (null) and completes.
      // This is used as a workaround to satisfy TypeScript requirements and the "combineLatest" 
      // expectation for a non-empty array of observables.
      of(null), // 
      // insertion point sub template
      this.checkboxService.getCheckBoxs(this.GONG__StackPath, this.frontRepo),
      this.formdivService.getFormDivs(this.GONG__StackPath, this.frontRepo),
      this.formeditassocbuttonService.getFormEditAssocButtons(this.GONG__StackPath, this.frontRepo),
      this.formfieldService.getFormFields(this.GONG__StackPath, this.frontRepo),
      this.formfielddateService.getFormFieldDates(this.GONG__StackPath, this.frontRepo),
      this.formfielddatetimeService.getFormFieldDateTimes(this.GONG__StackPath, this.frontRepo),
      this.formfieldfloat64Service.getFormFieldFloat64s(this.GONG__StackPath, this.frontRepo),
      this.formfieldintService.getFormFieldInts(this.GONG__StackPath, this.frontRepo),
      this.formfieldselectService.getFormFieldSelects(this.GONG__StackPath, this.frontRepo),
      this.formfieldstringService.getFormFieldStrings(this.GONG__StackPath, this.frontRepo),
      this.formfieldtimeService.getFormFieldTimes(this.GONG__StackPath, this.frontRepo),
      this.formgroupService.getFormGroups(this.GONG__StackPath, this.frontRepo),
      this.formsortassocbuttonService.getFormSortAssocButtons(this.GONG__StackPath, this.frontRepo),
      this.optionService.getOptions(this.GONG__StackPath, this.frontRepo),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

    this.GONG__StackPath = GONG__StackPath

    this.observableFrontRepo = [
      of(null), // see above for justification
      // insertion point sub template
      this.checkboxService.getCheckBoxs(this.GONG__StackPath, this.frontRepo),
      this.formdivService.getFormDivs(this.GONG__StackPath, this.frontRepo),
      this.formeditassocbuttonService.getFormEditAssocButtons(this.GONG__StackPath, this.frontRepo),
      this.formfieldService.getFormFields(this.GONG__StackPath, this.frontRepo),
      this.formfielddateService.getFormFieldDates(this.GONG__StackPath, this.frontRepo),
      this.formfielddatetimeService.getFormFieldDateTimes(this.GONG__StackPath, this.frontRepo),
      this.formfieldfloat64Service.getFormFieldFloat64s(this.GONG__StackPath, this.frontRepo),
      this.formfieldintService.getFormFieldInts(this.GONG__StackPath, this.frontRepo),
      this.formfieldselectService.getFormFieldSelects(this.GONG__StackPath, this.frontRepo),
      this.formfieldstringService.getFormFieldStrings(this.GONG__StackPath, this.frontRepo),
      this.formfieldtimeService.getFormFieldTimes(this.GONG__StackPath, this.frontRepo),
      this.formgroupService.getFormGroups(this.GONG__StackPath, this.frontRepo),
      this.formsortassocbuttonService.getFormSortAssocButtons(this.GONG__StackPath, this.frontRepo),
      this.optionService.getOptions(this.GONG__StackPath, this.frontRepo),
    ]

    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([
            ___of_null, // see above for the explanation about of
            // insertion point sub template for declarations 
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
              checkboxDB => {
                this.frontRepo.CheckBoxs.set(checkboxDB.ID, checkboxDB)
                this.frontRepo.CheckBoxs_batch.set(checkboxDB.ID, checkboxDB)
              }
            )

            // clear checkboxs that are absent from the batch
            this.frontRepo.CheckBoxs.forEach(
              checkboxDB => {
                if (this.frontRepo.CheckBoxs_batch.get(checkboxDB.ID) == undefined) {
                  this.frontRepo.CheckBoxs.delete(checkboxDB.ID)
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
              formdivDB => {
                this.frontRepo.FormDivs.set(formdivDB.ID, formdivDB)
                this.frontRepo.FormDivs_batch.set(formdivDB.ID, formdivDB)
              }
            )

            // clear formdivs that are absent from the batch
            this.frontRepo.FormDivs.forEach(
              formdivDB => {
                if (this.frontRepo.FormDivs_batch.get(formdivDB.ID) == undefined) {
                  this.frontRepo.FormDivs.delete(formdivDB.ID)
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
              formeditassocbuttonDB => {
                this.frontRepo.FormEditAssocButtons.set(formeditassocbuttonDB.ID, formeditassocbuttonDB)
                this.frontRepo.FormEditAssocButtons_batch.set(formeditassocbuttonDB.ID, formeditassocbuttonDB)
              }
            )

            // clear formeditassocbuttons that are absent from the batch
            this.frontRepo.FormEditAssocButtons.forEach(
              formeditassocbuttonDB => {
                if (this.frontRepo.FormEditAssocButtons_batch.get(formeditassocbuttonDB.ID) == undefined) {
                  this.frontRepo.FormEditAssocButtons.delete(formeditassocbuttonDB.ID)
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
              formfieldDB => {
                this.frontRepo.FormFields.set(formfieldDB.ID, formfieldDB)
                this.frontRepo.FormFields_batch.set(formfieldDB.ID, formfieldDB)
              }
            )

            // clear formfields that are absent from the batch
            this.frontRepo.FormFields.forEach(
              formfieldDB => {
                if (this.frontRepo.FormFields_batch.get(formfieldDB.ID) == undefined) {
                  this.frontRepo.FormFields.delete(formfieldDB.ID)
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
              formfielddateDB => {
                this.frontRepo.FormFieldDates.set(formfielddateDB.ID, formfielddateDB)
                this.frontRepo.FormFieldDates_batch.set(formfielddateDB.ID, formfielddateDB)
              }
            )

            // clear formfielddates that are absent from the batch
            this.frontRepo.FormFieldDates.forEach(
              formfielddateDB => {
                if (this.frontRepo.FormFieldDates_batch.get(formfielddateDB.ID) == undefined) {
                  this.frontRepo.FormFieldDates.delete(formfielddateDB.ID)
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
              formfielddatetimeDB => {
                this.frontRepo.FormFieldDateTimes.set(formfielddatetimeDB.ID, formfielddatetimeDB)
                this.frontRepo.FormFieldDateTimes_batch.set(formfielddatetimeDB.ID, formfielddatetimeDB)
              }
            )

            // clear formfielddatetimes that are absent from the batch
            this.frontRepo.FormFieldDateTimes.forEach(
              formfielddatetimeDB => {
                if (this.frontRepo.FormFieldDateTimes_batch.get(formfielddatetimeDB.ID) == undefined) {
                  this.frontRepo.FormFieldDateTimes.delete(formfielddatetimeDB.ID)
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
              formfieldfloat64DB => {
                this.frontRepo.FormFieldFloat64s.set(formfieldfloat64DB.ID, formfieldfloat64DB)
                this.frontRepo.FormFieldFloat64s_batch.set(formfieldfloat64DB.ID, formfieldfloat64DB)
              }
            )

            // clear formfieldfloat64s that are absent from the batch
            this.frontRepo.FormFieldFloat64s.forEach(
              formfieldfloat64DB => {
                if (this.frontRepo.FormFieldFloat64s_batch.get(formfieldfloat64DB.ID) == undefined) {
                  this.frontRepo.FormFieldFloat64s.delete(formfieldfloat64DB.ID)
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
              formfieldintDB => {
                this.frontRepo.FormFieldInts.set(formfieldintDB.ID, formfieldintDB)
                this.frontRepo.FormFieldInts_batch.set(formfieldintDB.ID, formfieldintDB)
              }
            )

            // clear formfieldints that are absent from the batch
            this.frontRepo.FormFieldInts.forEach(
              formfieldintDB => {
                if (this.frontRepo.FormFieldInts_batch.get(formfieldintDB.ID) == undefined) {
                  this.frontRepo.FormFieldInts.delete(formfieldintDB.ID)
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
              formfieldselectDB => {
                this.frontRepo.FormFieldSelects.set(formfieldselectDB.ID, formfieldselectDB)
                this.frontRepo.FormFieldSelects_batch.set(formfieldselectDB.ID, formfieldselectDB)
              }
            )

            // clear formfieldselects that are absent from the batch
            this.frontRepo.FormFieldSelects.forEach(
              formfieldselectDB => {
                if (this.frontRepo.FormFieldSelects_batch.get(formfieldselectDB.ID) == undefined) {
                  this.frontRepo.FormFieldSelects.delete(formfieldselectDB.ID)
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
              formfieldstringDB => {
                this.frontRepo.FormFieldStrings.set(formfieldstringDB.ID, formfieldstringDB)
                this.frontRepo.FormFieldStrings_batch.set(formfieldstringDB.ID, formfieldstringDB)
              }
            )

            // clear formfieldstrings that are absent from the batch
            this.frontRepo.FormFieldStrings.forEach(
              formfieldstringDB => {
                if (this.frontRepo.FormFieldStrings_batch.get(formfieldstringDB.ID) == undefined) {
                  this.frontRepo.FormFieldStrings.delete(formfieldstringDB.ID)
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
              formfieldtimeDB => {
                this.frontRepo.FormFieldTimes.set(formfieldtimeDB.ID, formfieldtimeDB)
                this.frontRepo.FormFieldTimes_batch.set(formfieldtimeDB.ID, formfieldtimeDB)
              }
            )

            // clear formfieldtimes that are absent from the batch
            this.frontRepo.FormFieldTimes.forEach(
              formfieldtimeDB => {
                if (this.frontRepo.FormFieldTimes_batch.get(formfieldtimeDB.ID) == undefined) {
                  this.frontRepo.FormFieldTimes.delete(formfieldtimeDB.ID)
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
              formgroupDB => {
                this.frontRepo.FormGroups.set(formgroupDB.ID, formgroupDB)
                this.frontRepo.FormGroups_batch.set(formgroupDB.ID, formgroupDB)
              }
            )

            // clear formgroups that are absent from the batch
            this.frontRepo.FormGroups.forEach(
              formgroupDB => {
                if (this.frontRepo.FormGroups_batch.get(formgroupDB.ID) == undefined) {
                  this.frontRepo.FormGroups.delete(formgroupDB.ID)
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
              formsortassocbuttonDB => {
                this.frontRepo.FormSortAssocButtons.set(formsortassocbuttonDB.ID, formsortassocbuttonDB)
                this.frontRepo.FormSortAssocButtons_batch.set(formsortassocbuttonDB.ID, formsortassocbuttonDB)
              }
            )

            // clear formsortassocbuttons that are absent from the batch
            this.frontRepo.FormSortAssocButtons.forEach(
              formsortassocbuttonDB => {
                if (this.frontRepo.FormSortAssocButtons_batch.get(formsortassocbuttonDB.ID) == undefined) {
                  this.frontRepo.FormSortAssocButtons.delete(formsortassocbuttonDB.ID)
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
              optionDB => {
                this.frontRepo.Options.set(optionDB.ID, optionDB)
                this.frontRepo.Options_batch.set(optionDB.ID, optionDB)
              }
            )

            // clear options that are absent from the batch
            this.frontRepo.Options.forEach(
              optionDB => {
                if (this.frontRepo.Options_batch.get(optionDB.ID) == undefined) {
                  this.frontRepo.Options.delete(optionDB.ID)
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
            // Second Step: reddeem slice of pointers fields
            // insertion point sub template for redeem 
            checkboxs.forEach(
              checkbox => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            formdivs.forEach(
              formdiv => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field FormEditAssocButton redeeming
                {
                  let _formeditassocbutton = this.frontRepo.FormEditAssocButtons.get(formdiv.FormDivPointersEncoding.FormEditAssocButtonID.Int64)
                  if (_formeditassocbutton) {
                    formdiv.FormEditAssocButton = _formeditassocbutton
                  }
                }
                // insertion point for pointer field FormSortAssocButton redeeming
                {
                  let _formsortassocbutton = this.frontRepo.FormSortAssocButtons.get(formdiv.FormDivPointersEncoding.FormSortAssocButtonID.Int64)
                  if (_formsortassocbutton) {
                    formdiv.FormSortAssocButton = _formsortassocbutton
                  }
                }
                // insertion point for pointers decoding
                formdiv.FormFields = new Array<FormFieldDB>()
                for (let _id of formdiv.FormDivPointersEncoding.FormFields) {
                  let _formfield = this.frontRepo.FormFields.get(_id)
                  if (_formfield != undefined) {
                    formdiv.FormFields.push(_formfield!)
                  }
                }
                formdiv.CheckBoxs = new Array<CheckBoxDB>()
                for (let _id of formdiv.FormDivPointersEncoding.CheckBoxs) {
                  let _checkbox = this.frontRepo.CheckBoxs.get(_id)
                  if (_checkbox != undefined) {
                    formdiv.CheckBoxs.push(_checkbox!)
                  }
                }
              }
            )
            formeditassocbuttons.forEach(
              formeditassocbutton => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            formfields.forEach(
              formfield => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field FormFieldString redeeming
                {
                  let _formfieldstring = this.frontRepo.FormFieldStrings.get(formfield.FormFieldPointersEncoding.FormFieldStringID.Int64)
                  if (_formfieldstring) {
                    formfield.FormFieldString = _formfieldstring
                  }
                }
                // insertion point for pointer field FormFieldFloat64 redeeming
                {
                  let _formfieldfloat64 = this.frontRepo.FormFieldFloat64s.get(formfield.FormFieldPointersEncoding.FormFieldFloat64ID.Int64)
                  if (_formfieldfloat64) {
                    formfield.FormFieldFloat64 = _formfieldfloat64
                  }
                }
                // insertion point for pointer field FormFieldInt redeeming
                {
                  let _formfieldint = this.frontRepo.FormFieldInts.get(formfield.FormFieldPointersEncoding.FormFieldIntID.Int64)
                  if (_formfieldint) {
                    formfield.FormFieldInt = _formfieldint
                  }
                }
                // insertion point for pointer field FormFieldDate redeeming
                {
                  let _formfielddate = this.frontRepo.FormFieldDates.get(formfield.FormFieldPointersEncoding.FormFieldDateID.Int64)
                  if (_formfielddate) {
                    formfield.FormFieldDate = _formfielddate
                  }
                }
                // insertion point for pointer field FormFieldTime redeeming
                {
                  let _formfieldtime = this.frontRepo.FormFieldTimes.get(formfield.FormFieldPointersEncoding.FormFieldTimeID.Int64)
                  if (_formfieldtime) {
                    formfield.FormFieldTime = _formfieldtime
                  }
                }
                // insertion point for pointer field FormFieldDateTime redeeming
                {
                  let _formfielddatetime = this.frontRepo.FormFieldDateTimes.get(formfield.FormFieldPointersEncoding.FormFieldDateTimeID.Int64)
                  if (_formfielddatetime) {
                    formfield.FormFieldDateTime = _formfielddatetime
                  }
                }
                // insertion point for pointer field FormFieldSelect redeeming
                {
                  let _formfieldselect = this.frontRepo.FormFieldSelects.get(formfield.FormFieldPointersEncoding.FormFieldSelectID.Int64)
                  if (_formfieldselect) {
                    formfield.FormFieldSelect = _formfieldselect
                  }
                }
                // insertion point for pointers decoding
              }
            )
            formfielddates.forEach(
              formfielddate => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            formfielddatetimes.forEach(
              formfielddatetime => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            formfieldfloat64s.forEach(
              formfieldfloat64 => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            formfieldints.forEach(
              formfieldint => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            formfieldselects.forEach(
              formfieldselect => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Value redeeming
                {
                  let _option = this.frontRepo.Options.get(formfieldselect.FormFieldSelectPointersEncoding.ValueID.Int64)
                  if (_option) {
                    formfieldselect.Value = _option
                  }
                }
                // insertion point for pointers decoding
                formfieldselect.Options = new Array<OptionDB>()
                for (let _id of formfieldselect.FormFieldSelectPointersEncoding.Options) {
                  let _option = this.frontRepo.Options.get(_id)
                  if (_option != undefined) {
                    formfieldselect.Options.push(_option!)
                  }
                }
              }
            )
            formfieldstrings.forEach(
              formfieldstring => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            formfieldtimes.forEach(
              formfieldtime => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            formgroups.forEach(
              formgroup => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
                formgroup.FormDivs = new Array<FormDivDB>()
                for (let _id of formgroup.FormGroupPointersEncoding.FormDivs) {
                  let _formdiv = this.frontRepo.FormDivs.get(_id)
                  if (_formdiv != undefined) {
                    formgroup.FormDivs.push(_formdiv!)
                  }
                }
              }
            )
            formsortassocbuttons.forEach(
              formsortassocbutton => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            options.forEach(
              option => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )

            // 
            // Third Step: reddeem front objects
            // insertion point sub template for redeem 
            
            // init front objects
            this.frontRepo.array_CheckBoxs = []
            this.frontRepo.map_ID_CheckBox.clear()
            this.frontRepo.CheckBoxs_array.forEach(
              checkboxDB => {
                let checkbox = new CheckBox
                CopyCheckBoxDBToCheckBox(checkboxDB, checkbox, this.frontRepo)
                this.frontRepo.array_CheckBoxs.push(checkbox)
                this.frontRepo.map_ID_CheckBox.set(checkbox.ID, checkbox)
              }
            )

            
            // init front objects
            this.frontRepo.array_FormDivs = []
            this.frontRepo.map_ID_FormDiv.clear()
            this.frontRepo.FormDivs_array.forEach(
              formdivDB => {
                let formdiv = new FormDiv
                CopyFormDivDBToFormDiv(formdivDB, formdiv, this.frontRepo)
                this.frontRepo.array_FormDivs.push(formdiv)
                this.frontRepo.map_ID_FormDiv.set(formdiv.ID, formdiv)
              }
            )

            
            // init front objects
            this.frontRepo.array_FormEditAssocButtons = []
            this.frontRepo.map_ID_FormEditAssocButton.clear()
            this.frontRepo.FormEditAssocButtons_array.forEach(
              formeditassocbuttonDB => {
                let formeditassocbutton = new FormEditAssocButton
                CopyFormEditAssocButtonDBToFormEditAssocButton(formeditassocbuttonDB, formeditassocbutton, this.frontRepo)
                this.frontRepo.array_FormEditAssocButtons.push(formeditassocbutton)
                this.frontRepo.map_ID_FormEditAssocButton.set(formeditassocbutton.ID, formeditassocbutton)
              }
            )

            
            // init front objects
            this.frontRepo.array_FormFields = []
            this.frontRepo.map_ID_FormField.clear()
            this.frontRepo.FormFields_array.forEach(
              formfieldDB => {
                let formfield = new FormField
                CopyFormFieldDBToFormField(formfieldDB, formfield, this.frontRepo)
                this.frontRepo.array_FormFields.push(formfield)
                this.frontRepo.map_ID_FormField.set(formfield.ID, formfield)
              }
            )

            
            // init front objects
            this.frontRepo.array_FormFieldDates = []
            this.frontRepo.map_ID_FormFieldDate.clear()
            this.frontRepo.FormFieldDates_array.forEach(
              formfielddateDB => {
                let formfielddate = new FormFieldDate
                CopyFormFieldDateDBToFormFieldDate(formfielddateDB, formfielddate, this.frontRepo)
                this.frontRepo.array_FormFieldDates.push(formfielddate)
                this.frontRepo.map_ID_FormFieldDate.set(formfielddate.ID, formfielddate)
              }
            )

            
            // init front objects
            this.frontRepo.array_FormFieldDateTimes = []
            this.frontRepo.map_ID_FormFieldDateTime.clear()
            this.frontRepo.FormFieldDateTimes_array.forEach(
              formfielddatetimeDB => {
                let formfielddatetime = new FormFieldDateTime
                CopyFormFieldDateTimeDBToFormFieldDateTime(formfielddatetimeDB, formfielddatetime, this.frontRepo)
                this.frontRepo.array_FormFieldDateTimes.push(formfielddatetime)
                this.frontRepo.map_ID_FormFieldDateTime.set(formfielddatetime.ID, formfielddatetime)
              }
            )

            
            // init front objects
            this.frontRepo.array_FormFieldFloat64s = []
            this.frontRepo.map_ID_FormFieldFloat64.clear()
            this.frontRepo.FormFieldFloat64s_array.forEach(
              formfieldfloat64DB => {
                let formfieldfloat64 = new FormFieldFloat64
                CopyFormFieldFloat64DBToFormFieldFloat64(formfieldfloat64DB, formfieldfloat64, this.frontRepo)
                this.frontRepo.array_FormFieldFloat64s.push(formfieldfloat64)
                this.frontRepo.map_ID_FormFieldFloat64.set(formfieldfloat64.ID, formfieldfloat64)
              }
            )

            
            // init front objects
            this.frontRepo.array_FormFieldInts = []
            this.frontRepo.map_ID_FormFieldInt.clear()
            this.frontRepo.FormFieldInts_array.forEach(
              formfieldintDB => {
                let formfieldint = new FormFieldInt
                CopyFormFieldIntDBToFormFieldInt(formfieldintDB, formfieldint, this.frontRepo)
                this.frontRepo.array_FormFieldInts.push(formfieldint)
                this.frontRepo.map_ID_FormFieldInt.set(formfieldint.ID, formfieldint)
              }
            )

            
            // init front objects
            this.frontRepo.array_FormFieldSelects = []
            this.frontRepo.map_ID_FormFieldSelect.clear()
            this.frontRepo.FormFieldSelects_array.forEach(
              formfieldselectDB => {
                let formfieldselect = new FormFieldSelect
                CopyFormFieldSelectDBToFormFieldSelect(formfieldselectDB, formfieldselect, this.frontRepo)
                this.frontRepo.array_FormFieldSelects.push(formfieldselect)
                this.frontRepo.map_ID_FormFieldSelect.set(formfieldselect.ID, formfieldselect)
              }
            )

            
            // init front objects
            this.frontRepo.array_FormFieldStrings = []
            this.frontRepo.map_ID_FormFieldString.clear()
            this.frontRepo.FormFieldStrings_array.forEach(
              formfieldstringDB => {
                let formfieldstring = new FormFieldString
                CopyFormFieldStringDBToFormFieldString(formfieldstringDB, formfieldstring, this.frontRepo)
                this.frontRepo.array_FormFieldStrings.push(formfieldstring)
                this.frontRepo.map_ID_FormFieldString.set(formfieldstring.ID, formfieldstring)
              }
            )

            
            // init front objects
            this.frontRepo.array_FormFieldTimes = []
            this.frontRepo.map_ID_FormFieldTime.clear()
            this.frontRepo.FormFieldTimes_array.forEach(
              formfieldtimeDB => {
                let formfieldtime = new FormFieldTime
                CopyFormFieldTimeDBToFormFieldTime(formfieldtimeDB, formfieldtime, this.frontRepo)
                this.frontRepo.array_FormFieldTimes.push(formfieldtime)
                this.frontRepo.map_ID_FormFieldTime.set(formfieldtime.ID, formfieldtime)
              }
            )

            
            // init front objects
            this.frontRepo.array_FormGroups = []
            this.frontRepo.map_ID_FormGroup.clear()
            this.frontRepo.FormGroups_array.forEach(
              formgroupDB => {
                let formgroup = new FormGroup
                CopyFormGroupDBToFormGroup(formgroupDB, formgroup, this.frontRepo)
                this.frontRepo.array_FormGroups.push(formgroup)
                this.frontRepo.map_ID_FormGroup.set(formgroup.ID, formgroup)
              }
            )

            
            // init front objects
            this.frontRepo.array_FormSortAssocButtons = []
            this.frontRepo.map_ID_FormSortAssocButton.clear()
            this.frontRepo.FormSortAssocButtons_array.forEach(
              formsortassocbuttonDB => {
                let formsortassocbutton = new FormSortAssocButton
                CopyFormSortAssocButtonDBToFormSortAssocButton(formsortassocbuttonDB, formsortassocbutton, this.frontRepo)
                this.frontRepo.array_FormSortAssocButtons.push(formsortassocbutton)
                this.frontRepo.map_ID_FormSortAssocButton.set(formsortassocbutton.ID, formsortassocbutton)
              }
            )

            
            // init front objects
            this.frontRepo.array_Options = []
            this.frontRepo.map_ID_Option.clear()
            this.frontRepo.Options_array.forEach(
              optionDB => {
                let option = new Option
                CopyOptionDBToOption(optionDB, option, this.frontRepo)
                this.frontRepo.array_Options.push(option)
                this.frontRepo.map_ID_Option.set(option.ID, option)
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
          this.checkboxService.getCheckBoxs(this.GONG__StackPath, this.frontRepo)
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
          this.formdivService.getFormDivs(this.GONG__StackPath, this.frontRepo)
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
                  let _formeditassocbutton = this.frontRepo.FormEditAssocButtons.get(formdiv.FormDivPointersEncoding.FormEditAssocButtonID.Int64)
                  if (_formeditassocbutton) {
                    formdiv.FormEditAssocButton = _formeditassocbutton
                  }
                }
                // insertion point for pointer field FormSortAssocButton redeeming
                {
                  let _formsortassocbutton = this.frontRepo.FormSortAssocButtons.get(formdiv.FormDivPointersEncoding.FormSortAssocButtonID.Int64)
                  if (_formsortassocbutton) {
                    formdiv.FormSortAssocButton = _formsortassocbutton
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
          this.formeditassocbuttonService.getFormEditAssocButtons(this.GONG__StackPath, this.frontRepo)
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
          this.formfieldService.getFormFields(this.GONG__StackPath, this.frontRepo)
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
                  let _formfieldstring = this.frontRepo.FormFieldStrings.get(formfield.FormFieldPointersEncoding.FormFieldStringID.Int64)
                  if (_formfieldstring) {
                    formfield.FormFieldString = _formfieldstring
                  }
                }
                // insertion point for pointer field FormFieldFloat64 redeeming
                {
                  let _formfieldfloat64 = this.frontRepo.FormFieldFloat64s.get(formfield.FormFieldPointersEncoding.FormFieldFloat64ID.Int64)
                  if (_formfieldfloat64) {
                    formfield.FormFieldFloat64 = _formfieldfloat64
                  }
                }
                // insertion point for pointer field FormFieldInt redeeming
                {
                  let _formfieldint = this.frontRepo.FormFieldInts.get(formfield.FormFieldPointersEncoding.FormFieldIntID.Int64)
                  if (_formfieldint) {
                    formfield.FormFieldInt = _formfieldint
                  }
                }
                // insertion point for pointer field FormFieldDate redeeming
                {
                  let _formfielddate = this.frontRepo.FormFieldDates.get(formfield.FormFieldPointersEncoding.FormFieldDateID.Int64)
                  if (_formfielddate) {
                    formfield.FormFieldDate = _formfielddate
                  }
                }
                // insertion point for pointer field FormFieldTime redeeming
                {
                  let _formfieldtime = this.frontRepo.FormFieldTimes.get(formfield.FormFieldPointersEncoding.FormFieldTimeID.Int64)
                  if (_formfieldtime) {
                    formfield.FormFieldTime = _formfieldtime
                  }
                }
                // insertion point for pointer field FormFieldDateTime redeeming
                {
                  let _formfielddatetime = this.frontRepo.FormFieldDateTimes.get(formfield.FormFieldPointersEncoding.FormFieldDateTimeID.Int64)
                  if (_formfielddatetime) {
                    formfield.FormFieldDateTime = _formfielddatetime
                  }
                }
                // insertion point for pointer field FormFieldSelect redeeming
                {
                  let _formfieldselect = this.frontRepo.FormFieldSelects.get(formfield.FormFieldPointersEncoding.FormFieldSelectID.Int64)
                  if (_formfieldselect) {
                    formfield.FormFieldSelect = _formfieldselect
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
          this.formfielddateService.getFormFieldDates(this.GONG__StackPath, this.frontRepo)
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
          this.formfielddatetimeService.getFormFieldDateTimes(this.GONG__StackPath, this.frontRepo)
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
          this.formfieldfloat64Service.getFormFieldFloat64s(this.GONG__StackPath, this.frontRepo)
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
          this.formfieldintService.getFormFieldInts(this.GONG__StackPath, this.frontRepo)
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
          this.formfieldselectService.getFormFieldSelects(this.GONG__StackPath, this.frontRepo)
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
                  let _option = this.frontRepo.Options.get(formfieldselect.FormFieldSelectPointersEncoding.ValueID.Int64)
                  if (_option) {
                    formfieldselect.Value = _option
                  }
                }
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
          this.formfieldstringService.getFormFieldStrings(this.GONG__StackPath, this.frontRepo)
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
          this.formfieldtimeService.getFormFieldTimes(this.GONG__StackPath, this.frontRepo)
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
          this.formgroupService.getFormGroups(this.GONG__StackPath, this.frontRepo)
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
          this.formsortassocbuttonService.getFormSortAssocButtons(this.GONG__StackPath, this.frontRepo)
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
          this.optionService.getOptions(this.GONG__StackPath, this.frontRepo)
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
