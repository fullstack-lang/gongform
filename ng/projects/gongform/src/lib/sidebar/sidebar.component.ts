import { Component, Input, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { BehaviorSubject, Subscription } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbFromBackService } from '../commitnbfromback.service'
import { GongstructSelectionService } from '../gongstruct-selection.service'

// insertion point for per struct import code
import { CheckBoxService } from '../checkbox.service'
import { getCheckBoxUniqueID } from '../front-repo.service'
import { FormDivService } from '../formdiv.service'
import { getFormDivUniqueID } from '../front-repo.service'
import { FormEditAssocButtonService } from '../formeditassocbutton.service'
import { getFormEditAssocButtonUniqueID } from '../front-repo.service'
import { FormFieldService } from '../formfield.service'
import { getFormFieldUniqueID } from '../front-repo.service'
import { FormFieldDateService } from '../formfielddate.service'
import { getFormFieldDateUniqueID } from '../front-repo.service'
import { FormFieldDateTimeService } from '../formfielddatetime.service'
import { getFormFieldDateTimeUniqueID } from '../front-repo.service'
import { FormFieldFloat64Service } from '../formfieldfloat64.service'
import { getFormFieldFloat64UniqueID } from '../front-repo.service'
import { FormFieldIntService } from '../formfieldint.service'
import { getFormFieldIntUniqueID } from '../front-repo.service'
import { FormFieldSelectService } from '../formfieldselect.service'
import { getFormFieldSelectUniqueID } from '../front-repo.service'
import { FormFieldStringService } from '../formfieldstring.service'
import { getFormFieldStringUniqueID } from '../front-repo.service'
import { FormFieldTimeService } from '../formfieldtime.service'
import { getFormFieldTimeUniqueID } from '../front-repo.service'
import { FormGroupService } from '../formgroup.service'
import { getFormGroupUniqueID } from '../front-repo.service'
import { FormSortAssocButtonService } from '../formsortassocbutton.service'
import { getFormSortAssocButtonUniqueID } from '../front-repo.service'
import { OptionService } from '../option.service'
import { getOptionUniqueID } from '../front-repo.service'

import { RouteService } from '../route-service';

/**
 * Types of a GongNode / GongFlatNode
 */
export enum GongNodeType {
  STRUCT = "STRUCT",
  INSTANCE = "INSTANCE",
  ONE__ZERO_ONE_ASSOCIATION = 'ONE__ZERO_ONE_ASSOCIATION',
  ONE__ZERO_MANY_ASSOCIATION = 'ONE__ZERO_MANY_ASSOCIATION',
}

/**
 * GongNode is the "data" node
 */
interface GongNode {
  name: string; // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children: GongNode[];
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


/** 
 * GongFlatNode is the dynamic visual node with expandable and level information
 * */
interface GongFlatNode {
  expandable: boolean;
  name: string;
  level: number;
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


@Component({
  selector: 'app-gongform-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css'],
})
export class SidebarComponent implements OnInit {

  /**
  * _transformer generated a displayed node from a data node
  *
  * @param node input data noe
  * @param level input level
  *
  * @returns an ExampleFlatNode
  */
  private _transformer = (node: GongNode, level: number) => {
    return {

      /**
      * in javascript, The !! ensures the resulting type is a boolean (true or false).
      *
      * !!node.children will evaluate to true is the variable is defined
      */
      expandable: !!node.children && node.children.length > 0,
      name: node.name,
      level: level,
      type: node.type,
      structName: node.structName,
      associationField: node.associationField,
      associatedStructName: node.associatedStructName,
      id: node.id,
      uniqueIdPerStack: node.uniqueIdPerStack,
    }
  }

  /**
   * treeControl is passed as the paramter treeControl in the "mat-tree" selector
   *
   * Flat tree control. Able to expand/collapse a subtree recursively for flattened tree.
   *
   * Construct with flat tree data node functions getLevel and isExpandable.
  constructor(
    getLevel: (dataNode: T) => number,
    isExpandable: (dataNode: T) => boolean, 
    options?: FlatTreeControlOptions<T, K> | undefined);
   */
  treeControl = new FlatTreeControl<GongFlatNode>(
    node => node.level,
    node => node.expandable
  );

  /**
   * from mat-tree documentation
   *
   * Tree flattener to convert a normal type of node to node with children & level information.
   */
  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children
  );

  /**
   * data is the other paramter to the "mat-tree" selector
   * 
   * strangely, the dataSource declaration has to follow the treeFlattener declaration
   */
  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  /**
   * hasChild is used by the selector for expandable nodes
   * 
   *  <mat-tree-node *matTreeNodeDef="let node;when: hasChild" matTreeNodePadding>
   * 
   * @param _ 
   * @param node 
   */
  hasChild = (_: number, node: GongFlatNode) => node.expandable;

  // front repo
  frontRepo: FrontRepo = new (FrontRepo)
  commitNbFromBack: number = 0

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  // SelectedStructChanged is the behavior subject that will emit
  // the selected gong struct whose table has to be displayed in the table outlet
  SelectedStructChanged: BehaviorSubject<string> = new BehaviorSubject("");

  subscription: Subscription = new Subscription

  @Input() GONG__StackPath: string = ""

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private gongstructSelectionService: GongstructSelectionService,

    // insertion point for per struct service declaration
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

    private routeService: RouteService,
  ) { }

  ngOnDestroy() {
    // prevent memory leak when component destroyed
    this.subscription.unsubscribe();
  }

  ngOnInit(): void {

    console.log("Sidebar init: " + this.GONG__StackPath)

    // add the routes that will used by this side panel component and
    // by the component that are called from this component
    this.routeService.addDataPanelRoutes(this.GONG__StackPath)

    this.subscription = this.gongstructSelectionService.gongtructSelected$.subscribe(
      gongstructName => {
        // console.log("sidebar gongstruct selected " + gongstructName)

        this.setTableRouterOutlet(gongstructName.toLowerCase() + "s")
      });

    this.refresh()

    this.SelectedStructChanged.subscribe(
      selectedStruct => {
        if (selectedStruct != "") {
          this.setTableRouterOutlet(selectedStruct.toLowerCase() + "s")
        }
      }
    )

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.checkboxService.CheckBoxServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formdivService.FormDivServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formeditassocbuttonService.FormEditAssocButtonServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formfieldService.FormFieldServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formfielddateService.FormFieldDateServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formfielddatetimeService.FormFieldDateTimeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formfieldfloat64Service.FormFieldFloat64ServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formfieldintService.FormFieldIntServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formfieldselectService.FormFieldSelectServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formfieldstringService.FormFieldStringServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formfieldtimeService.FormFieldTimeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formgroupService.FormGroupServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.formsortassocbuttonService.FormSortAssocButtonServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.optionService.OptionServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
  }

  refresh(): void {
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(frontRepo => {
      this.frontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      this.treeControl.dataNodes?.forEach(
        node => {
          if (this.treeControl.isExpanded(node)) {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, true)
          } else {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, false)
          }
        }
      )

      // reset the gong node tree
      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction
      /**
      * fill up the CheckBox part of the mat tree
      */
      let checkboxGongNodeStruct: GongNode = {
        name: "CheckBox",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "CheckBox",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(checkboxGongNodeStruct)

      this.frontRepo.CheckBoxs_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.CheckBoxs_array.forEach(
        checkboxDB => {
          let checkboxGongNodeInstance: GongNode = {
            name: checkboxDB.Name,
            type: GongNodeType.INSTANCE,
            id: checkboxDB.ID,
            uniqueIdPerStack: getCheckBoxUniqueID(checkboxDB.ID),
            structName: "CheckBox",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          checkboxGongNodeStruct.children!.push(checkboxGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the FormDiv part of the mat tree
      */
      let formdivGongNodeStruct: GongNode = {
        name: "FormDiv",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormDiv",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formdivGongNodeStruct)

      this.frontRepo.FormDivs_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormDivs_array.forEach(
        formdivDB => {
          let formdivGongNodeInstance: GongNode = {
            name: formdivDB.Name,
            type: GongNodeType.INSTANCE,
            id: formdivDB.ID,
            uniqueIdPerStack: getFormDivUniqueID(formdivDB.ID),
            structName: "FormDiv",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formdivGongNodeStruct.children!.push(formdivGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer FormFields
          */
          let FormFieldsGongNodeAssociation: GongNode = {
            name: "(FormField) FormFields",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: formdivDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "FormDiv",
            associationField: "FormFields",
            associatedStructName: "FormField",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formdivGongNodeInstance.children.push(FormFieldsGongNodeAssociation)

          formdivDB.FormFields?.forEach(formfieldDB => {
            let formfieldNode: GongNode = {
              name: formfieldDB.Name,
              type: GongNodeType.INSTANCE,
              id: formfieldDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getFormDivUniqueID(formdivDB.ID)
                + 11 * getFormFieldUniqueID(formfieldDB.ID),
              structName: "FormField",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormFieldsGongNodeAssociation.children.push(formfieldNode)
          })

          /**
          * let append a node for the slide of pointer CheckBoxs
          */
          let CheckBoxsGongNodeAssociation: GongNode = {
            name: "(CheckBox) CheckBoxs",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: formdivDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "FormDiv",
            associationField: "CheckBoxs",
            associatedStructName: "CheckBox",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formdivGongNodeInstance.children.push(CheckBoxsGongNodeAssociation)

          formdivDB.CheckBoxs?.forEach(checkboxDB => {
            let checkboxNode: GongNode = {
              name: checkboxDB.Name,
              type: GongNodeType.INSTANCE,
              id: checkboxDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getFormDivUniqueID(formdivDB.ID)
                + 11 * getCheckBoxUniqueID(checkboxDB.ID),
              structName: "CheckBox",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            CheckBoxsGongNodeAssociation.children.push(checkboxNode)
          })

          /**
          * let append a node for the association FormEditAssocButton
          */
          let FormEditAssocButtonGongNodeAssociation: GongNode = {
            name: "(FormEditAssocButton) FormEditAssocButton",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: formdivDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "FormDiv",
            associationField: "FormEditAssocButton",
            associatedStructName: "FormEditAssocButton",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formdivGongNodeInstance.children!.push(FormEditAssocButtonGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation FormEditAssocButton
            */
          if (formdivDB.FormEditAssocButton != undefined) {
            let formdivGongNodeInstance_FormEditAssocButton: GongNode = {
              name: formdivDB.FormEditAssocButton.Name,
              type: GongNodeType.INSTANCE,
              id: formdivDB.FormEditAssocButton.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getFormDivUniqueID(formdivDB.ID)
                + 5 * getFormEditAssocButtonUniqueID(formdivDB.FormEditAssocButton.ID),
              structName: "FormEditAssocButton",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormEditAssocButtonGongNodeAssociation.children.push(formdivGongNodeInstance_FormEditAssocButton)
          }

          /**
          * let append a node for the association FormSortAssocButton
          */
          let FormSortAssocButtonGongNodeAssociation: GongNode = {
            name: "(FormSortAssocButton) FormSortAssocButton",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: formdivDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "FormDiv",
            associationField: "FormSortAssocButton",
            associatedStructName: "FormSortAssocButton",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formdivGongNodeInstance.children!.push(FormSortAssocButtonGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation FormSortAssocButton
            */
          if (formdivDB.FormSortAssocButton != undefined) {
            let formdivGongNodeInstance_FormSortAssocButton: GongNode = {
              name: formdivDB.FormSortAssocButton.Name,
              type: GongNodeType.INSTANCE,
              id: formdivDB.FormSortAssocButton.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getFormDivUniqueID(formdivDB.ID)
                + 5 * getFormSortAssocButtonUniqueID(formdivDB.FormSortAssocButton.ID),
              structName: "FormSortAssocButton",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormSortAssocButtonGongNodeAssociation.children.push(formdivGongNodeInstance_FormSortAssocButton)
          }

        }
      )

      /**
      * fill up the FormEditAssocButton part of the mat tree
      */
      let formeditassocbuttonGongNodeStruct: GongNode = {
        name: "FormEditAssocButton",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormEditAssocButton",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formeditassocbuttonGongNodeStruct)

      this.frontRepo.FormEditAssocButtons_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormEditAssocButtons_array.forEach(
        formeditassocbuttonDB => {
          let formeditassocbuttonGongNodeInstance: GongNode = {
            name: formeditassocbuttonDB.Name,
            type: GongNodeType.INSTANCE,
            id: formeditassocbuttonDB.ID,
            uniqueIdPerStack: getFormEditAssocButtonUniqueID(formeditassocbuttonDB.ID),
            structName: "FormEditAssocButton",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formeditassocbuttonGongNodeStruct.children!.push(formeditassocbuttonGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the FormField part of the mat tree
      */
      let formfieldGongNodeStruct: GongNode = {
        name: "FormField",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormField",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formfieldGongNodeStruct)

      this.frontRepo.FormFields_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormFields_array.forEach(
        formfieldDB => {
          let formfieldGongNodeInstance: GongNode = {
            name: formfieldDB.Name,
            type: GongNodeType.INSTANCE,
            id: formfieldDB.ID,
            uniqueIdPerStack: getFormFieldUniqueID(formfieldDB.ID),
            structName: "FormField",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formfieldGongNodeStruct.children!.push(formfieldGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association FormFieldString
          */
          let FormFieldStringGongNodeAssociation: GongNode = {
            name: "(FormFieldString) FormFieldString",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: formfieldDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "FormField",
            associationField: "FormFieldString",
            associatedStructName: "FormFieldString",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formfieldGongNodeInstance.children!.push(FormFieldStringGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation FormFieldString
            */
          if (formfieldDB.FormFieldString != undefined) {
            let formfieldGongNodeInstance_FormFieldString: GongNode = {
              name: formfieldDB.FormFieldString.Name,
              type: GongNodeType.INSTANCE,
              id: formfieldDB.FormFieldString.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getFormFieldUniqueID(formfieldDB.ID)
                + 5 * getFormFieldStringUniqueID(formfieldDB.FormFieldString.ID),
              structName: "FormFieldString",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormFieldStringGongNodeAssociation.children.push(formfieldGongNodeInstance_FormFieldString)
          }

          /**
          * let append a node for the association FormFieldFloat64
          */
          let FormFieldFloat64GongNodeAssociation: GongNode = {
            name: "(FormFieldFloat64) FormFieldFloat64",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: formfieldDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "FormField",
            associationField: "FormFieldFloat64",
            associatedStructName: "FormFieldFloat64",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formfieldGongNodeInstance.children!.push(FormFieldFloat64GongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation FormFieldFloat64
            */
          if (formfieldDB.FormFieldFloat64 != undefined) {
            let formfieldGongNodeInstance_FormFieldFloat64: GongNode = {
              name: formfieldDB.FormFieldFloat64.Name,
              type: GongNodeType.INSTANCE,
              id: formfieldDB.FormFieldFloat64.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getFormFieldUniqueID(formfieldDB.ID)
                + 5 * getFormFieldFloat64UniqueID(formfieldDB.FormFieldFloat64.ID),
              structName: "FormFieldFloat64",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormFieldFloat64GongNodeAssociation.children.push(formfieldGongNodeInstance_FormFieldFloat64)
          }

          /**
          * let append a node for the association FormFieldInt
          */
          let FormFieldIntGongNodeAssociation: GongNode = {
            name: "(FormFieldInt) FormFieldInt",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: formfieldDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "FormField",
            associationField: "FormFieldInt",
            associatedStructName: "FormFieldInt",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formfieldGongNodeInstance.children!.push(FormFieldIntGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation FormFieldInt
            */
          if (formfieldDB.FormFieldInt != undefined) {
            let formfieldGongNodeInstance_FormFieldInt: GongNode = {
              name: formfieldDB.FormFieldInt.Name,
              type: GongNodeType.INSTANCE,
              id: formfieldDB.FormFieldInt.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getFormFieldUniqueID(formfieldDB.ID)
                + 5 * getFormFieldIntUniqueID(formfieldDB.FormFieldInt.ID),
              structName: "FormFieldInt",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormFieldIntGongNodeAssociation.children.push(formfieldGongNodeInstance_FormFieldInt)
          }

          /**
          * let append a node for the association FormFieldDate
          */
          let FormFieldDateGongNodeAssociation: GongNode = {
            name: "(FormFieldDate) FormFieldDate",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: formfieldDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "FormField",
            associationField: "FormFieldDate",
            associatedStructName: "FormFieldDate",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formfieldGongNodeInstance.children!.push(FormFieldDateGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation FormFieldDate
            */
          if (formfieldDB.FormFieldDate != undefined) {
            let formfieldGongNodeInstance_FormFieldDate: GongNode = {
              name: formfieldDB.FormFieldDate.Name,
              type: GongNodeType.INSTANCE,
              id: formfieldDB.FormFieldDate.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getFormFieldUniqueID(formfieldDB.ID)
                + 5 * getFormFieldDateUniqueID(formfieldDB.FormFieldDate.ID),
              structName: "FormFieldDate",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormFieldDateGongNodeAssociation.children.push(formfieldGongNodeInstance_FormFieldDate)
          }

          /**
          * let append a node for the association FormFieldTime
          */
          let FormFieldTimeGongNodeAssociation: GongNode = {
            name: "(FormFieldTime) FormFieldTime",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: formfieldDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "FormField",
            associationField: "FormFieldTime",
            associatedStructName: "FormFieldTime",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formfieldGongNodeInstance.children!.push(FormFieldTimeGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation FormFieldTime
            */
          if (formfieldDB.FormFieldTime != undefined) {
            let formfieldGongNodeInstance_FormFieldTime: GongNode = {
              name: formfieldDB.FormFieldTime.Name,
              type: GongNodeType.INSTANCE,
              id: formfieldDB.FormFieldTime.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getFormFieldUniqueID(formfieldDB.ID)
                + 5 * getFormFieldTimeUniqueID(formfieldDB.FormFieldTime.ID),
              structName: "FormFieldTime",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormFieldTimeGongNodeAssociation.children.push(formfieldGongNodeInstance_FormFieldTime)
          }

          /**
          * let append a node for the association FormFieldDateTime
          */
          let FormFieldDateTimeGongNodeAssociation: GongNode = {
            name: "(FormFieldDateTime) FormFieldDateTime",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: formfieldDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "FormField",
            associationField: "FormFieldDateTime",
            associatedStructName: "FormFieldDateTime",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formfieldGongNodeInstance.children!.push(FormFieldDateTimeGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation FormFieldDateTime
            */
          if (formfieldDB.FormFieldDateTime != undefined) {
            let formfieldGongNodeInstance_FormFieldDateTime: GongNode = {
              name: formfieldDB.FormFieldDateTime.Name,
              type: GongNodeType.INSTANCE,
              id: formfieldDB.FormFieldDateTime.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getFormFieldUniqueID(formfieldDB.ID)
                + 5 * getFormFieldDateTimeUniqueID(formfieldDB.FormFieldDateTime.ID),
              structName: "FormFieldDateTime",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormFieldDateTimeGongNodeAssociation.children.push(formfieldGongNodeInstance_FormFieldDateTime)
          }

          /**
          * let append a node for the association FormFieldSelect
          */
          let FormFieldSelectGongNodeAssociation: GongNode = {
            name: "(FormFieldSelect) FormFieldSelect",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: formfieldDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "FormField",
            associationField: "FormFieldSelect",
            associatedStructName: "FormFieldSelect",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formfieldGongNodeInstance.children!.push(FormFieldSelectGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation FormFieldSelect
            */
          if (formfieldDB.FormFieldSelect != undefined) {
            let formfieldGongNodeInstance_FormFieldSelect: GongNode = {
              name: formfieldDB.FormFieldSelect.Name,
              type: GongNodeType.INSTANCE,
              id: formfieldDB.FormFieldSelect.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getFormFieldUniqueID(formfieldDB.ID)
                + 5 * getFormFieldSelectUniqueID(formfieldDB.FormFieldSelect.ID),
              structName: "FormFieldSelect",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormFieldSelectGongNodeAssociation.children.push(formfieldGongNodeInstance_FormFieldSelect)
          }

        }
      )

      /**
      * fill up the FormFieldDate part of the mat tree
      */
      let formfielddateGongNodeStruct: GongNode = {
        name: "FormFieldDate",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormFieldDate",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formfielddateGongNodeStruct)

      this.frontRepo.FormFieldDates_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormFieldDates_array.forEach(
        formfielddateDB => {
          let formfielddateGongNodeInstance: GongNode = {
            name: formfielddateDB.Name,
            type: GongNodeType.INSTANCE,
            id: formfielddateDB.ID,
            uniqueIdPerStack: getFormFieldDateUniqueID(formfielddateDB.ID),
            structName: "FormFieldDate",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formfielddateGongNodeStruct.children!.push(formfielddateGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the FormFieldDateTime part of the mat tree
      */
      let formfielddatetimeGongNodeStruct: GongNode = {
        name: "FormFieldDateTime",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormFieldDateTime",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formfielddatetimeGongNodeStruct)

      this.frontRepo.FormFieldDateTimes_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormFieldDateTimes_array.forEach(
        formfielddatetimeDB => {
          let formfielddatetimeGongNodeInstance: GongNode = {
            name: formfielddatetimeDB.Name,
            type: GongNodeType.INSTANCE,
            id: formfielddatetimeDB.ID,
            uniqueIdPerStack: getFormFieldDateTimeUniqueID(formfielddatetimeDB.ID),
            structName: "FormFieldDateTime",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formfielddatetimeGongNodeStruct.children!.push(formfielddatetimeGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the FormFieldFloat64 part of the mat tree
      */
      let formfieldfloat64GongNodeStruct: GongNode = {
        name: "FormFieldFloat64",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormFieldFloat64",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formfieldfloat64GongNodeStruct)

      this.frontRepo.FormFieldFloat64s_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormFieldFloat64s_array.forEach(
        formfieldfloat64DB => {
          let formfieldfloat64GongNodeInstance: GongNode = {
            name: formfieldfloat64DB.Name,
            type: GongNodeType.INSTANCE,
            id: formfieldfloat64DB.ID,
            uniqueIdPerStack: getFormFieldFloat64UniqueID(formfieldfloat64DB.ID),
            structName: "FormFieldFloat64",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formfieldfloat64GongNodeStruct.children!.push(formfieldfloat64GongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the FormFieldInt part of the mat tree
      */
      let formfieldintGongNodeStruct: GongNode = {
        name: "FormFieldInt",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormFieldInt",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formfieldintGongNodeStruct)

      this.frontRepo.FormFieldInts_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormFieldInts_array.forEach(
        formfieldintDB => {
          let formfieldintGongNodeInstance: GongNode = {
            name: formfieldintDB.Name,
            type: GongNodeType.INSTANCE,
            id: formfieldintDB.ID,
            uniqueIdPerStack: getFormFieldIntUniqueID(formfieldintDB.ID),
            structName: "FormFieldInt",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formfieldintGongNodeStruct.children!.push(formfieldintGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the FormFieldSelect part of the mat tree
      */
      let formfieldselectGongNodeStruct: GongNode = {
        name: "FormFieldSelect",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormFieldSelect",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formfieldselectGongNodeStruct)

      this.frontRepo.FormFieldSelects_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormFieldSelects_array.forEach(
        formfieldselectDB => {
          let formfieldselectGongNodeInstance: GongNode = {
            name: formfieldselectDB.Name,
            type: GongNodeType.INSTANCE,
            id: formfieldselectDB.ID,
            uniqueIdPerStack: getFormFieldSelectUniqueID(formfieldselectDB.ID),
            structName: "FormFieldSelect",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formfieldselectGongNodeStruct.children!.push(formfieldselectGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Value
          */
          let ValueGongNodeAssociation: GongNode = {
            name: "(Option) Value",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: formfieldselectDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "FormFieldSelect",
            associationField: "Value",
            associatedStructName: "Option",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formfieldselectGongNodeInstance.children!.push(ValueGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Value
            */
          if (formfieldselectDB.Value != undefined) {
            let formfieldselectGongNodeInstance_Value: GongNode = {
              name: formfieldselectDB.Value.Name,
              type: GongNodeType.INSTANCE,
              id: formfieldselectDB.Value.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getFormFieldSelectUniqueID(formfieldselectDB.ID)
                + 5 * getOptionUniqueID(formfieldselectDB.Value.ID),
              structName: "Option",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ValueGongNodeAssociation.children.push(formfieldselectGongNodeInstance_Value)
          }

          /**
          * let append a node for the slide of pointer Options
          */
          let OptionsGongNodeAssociation: GongNode = {
            name: "(Option) Options",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: formfieldselectDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "FormFieldSelect",
            associationField: "Options",
            associatedStructName: "Option",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formfieldselectGongNodeInstance.children.push(OptionsGongNodeAssociation)

          formfieldselectDB.Options?.forEach(optionDB => {
            let optionNode: GongNode = {
              name: optionDB.Name,
              type: GongNodeType.INSTANCE,
              id: optionDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getFormFieldSelectUniqueID(formfieldselectDB.ID)
                + 11 * getOptionUniqueID(optionDB.ID),
              structName: "Option",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            OptionsGongNodeAssociation.children.push(optionNode)
          })

        }
      )

      /**
      * fill up the FormFieldString part of the mat tree
      */
      let formfieldstringGongNodeStruct: GongNode = {
        name: "FormFieldString",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormFieldString",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formfieldstringGongNodeStruct)

      this.frontRepo.FormFieldStrings_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormFieldStrings_array.forEach(
        formfieldstringDB => {
          let formfieldstringGongNodeInstance: GongNode = {
            name: formfieldstringDB.Name,
            type: GongNodeType.INSTANCE,
            id: formfieldstringDB.ID,
            uniqueIdPerStack: getFormFieldStringUniqueID(formfieldstringDB.ID),
            structName: "FormFieldString",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formfieldstringGongNodeStruct.children!.push(formfieldstringGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the FormFieldTime part of the mat tree
      */
      let formfieldtimeGongNodeStruct: GongNode = {
        name: "FormFieldTime",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormFieldTime",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formfieldtimeGongNodeStruct)

      this.frontRepo.FormFieldTimes_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormFieldTimes_array.forEach(
        formfieldtimeDB => {
          let formfieldtimeGongNodeInstance: GongNode = {
            name: formfieldtimeDB.Name,
            type: GongNodeType.INSTANCE,
            id: formfieldtimeDB.ID,
            uniqueIdPerStack: getFormFieldTimeUniqueID(formfieldtimeDB.ID),
            structName: "FormFieldTime",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formfieldtimeGongNodeStruct.children!.push(formfieldtimeGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the FormGroup part of the mat tree
      */
      let formgroupGongNodeStruct: GongNode = {
        name: "FormGroup",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormGroup",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formgroupGongNodeStruct)

      this.frontRepo.FormGroups_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormGroups_array.forEach(
        formgroupDB => {
          let formgroupGongNodeInstance: GongNode = {
            name: formgroupDB.Name,
            type: GongNodeType.INSTANCE,
            id: formgroupDB.ID,
            uniqueIdPerStack: getFormGroupUniqueID(formgroupDB.ID),
            structName: "FormGroup",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formgroupGongNodeStruct.children!.push(formgroupGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer FormDivs
          */
          let FormDivsGongNodeAssociation: GongNode = {
            name: "(FormDiv) FormDivs",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: formgroupDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "FormGroup",
            associationField: "FormDivs",
            associatedStructName: "FormDiv",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          formgroupGongNodeInstance.children.push(FormDivsGongNodeAssociation)

          formgroupDB.FormDivs?.forEach(formdivDB => {
            let formdivNode: GongNode = {
              name: formdivDB.Name,
              type: GongNodeType.INSTANCE,
              id: formdivDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getFormGroupUniqueID(formgroupDB.ID)
                + 11 * getFormDivUniqueID(formdivDB.ID),
              structName: "FormDiv",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FormDivsGongNodeAssociation.children.push(formdivNode)
          })

        }
      )

      /**
      * fill up the FormSortAssocButton part of the mat tree
      */
      let formsortassocbuttonGongNodeStruct: GongNode = {
        name: "FormSortAssocButton",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "FormSortAssocButton",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(formsortassocbuttonGongNodeStruct)

      this.frontRepo.FormSortAssocButtons_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.FormSortAssocButtons_array.forEach(
        formsortassocbuttonDB => {
          let formsortassocbuttonGongNodeInstance: GongNode = {
            name: formsortassocbuttonDB.Name,
            type: GongNodeType.INSTANCE,
            id: formsortassocbuttonDB.ID,
            uniqueIdPerStack: getFormSortAssocButtonUniqueID(formsortassocbuttonDB.ID),
            structName: "FormSortAssocButton",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          formsortassocbuttonGongNodeStruct.children!.push(formsortassocbuttonGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Option part of the mat tree
      */
      let optionGongNodeStruct: GongNode = {
        name: "Option",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Option",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(optionGongNodeStruct)

      this.frontRepo.Options_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Options_array.forEach(
        optionDB => {
          let optionGongNodeInstance: GongNode = {
            name: optionDB.Name,
            type: GongNodeType.INSTANCE,
            id: optionDB.ID,
            uniqueIdPerStack: getOptionUniqueID(optionDB.ID),
            structName: "Option",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          optionGongNodeStruct.children!.push(optionGongNodeInstance)

          // insertion point for per field code
        }
      )


      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      this.treeControl.dataNodes?.forEach(
        node => {
          if (memoryOfExpandedNodes.get(node.uniqueIdPerStack)) {
            this.treeControl.expand(node)
          }
        }
      )
    })
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutlet(path: string) {
    let outletName = this.routeService.getTableOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + path
    this.router.navigate([{
      outlets: {
        outletName: [fullPath]
      }
    }]);
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutletFromTree(path: string, type: GongNodeType, structName: string, id: number) {

    if (type == GongNodeType.STRUCT) {
      let outletName = this.routeService.getTableOutlet(this.GONG__StackPath)
      let fullPath = this.routeService.getPathRoot() + "-" + path.toLowerCase()
      let outletConf: any = {}
      outletConf[outletName] = [fullPath, this.GONG__StackPath]

      this.router.navigate([{ outlets: outletConf }])
    }

    if (type == GongNodeType.INSTANCE) {
      let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
      let fullPath = this.routeService.getPathRoot() + "-" + structName.toLowerCase() + "-detail"

      let outletConf: any = {}
      outletConf[outletName] = [fullPath, id, this.GONG__StackPath]

      this.router.navigate([{ outlets: outletConf }])
    }
  }

  setEditorRouterOutlet(path: string) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + path.toLowerCase()
    let outletConf : any = {}
    outletConf[outletName] = [fullPath, this.GONG__StackPath]
    this.router.navigate([{ outlets: outletConf }]);
  }

  setEditorSpecialRouterOutlet(node: GongFlatNode) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + node.associatedStructName.toLowerCase() + "-adder"
    this.router.navigate([{
      outlets: {
        outletName: [fullPath, node.id, node.structName, node.associationField, this.GONG__StackPath]
      }
    }]);
  }
}
