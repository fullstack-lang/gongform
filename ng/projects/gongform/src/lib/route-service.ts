import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { CheckBoxsTableComponent } from './checkboxs-table/checkboxs-table.component'
import { CheckBoxDetailComponent } from './checkbox-detail/checkbox-detail.component'

import { FormDivsTableComponent } from './formdivs-table/formdivs-table.component'
import { FormDivDetailComponent } from './formdiv-detail/formdiv-detail.component'

import { FormEditAssocButtonsTableComponent } from './formeditassocbuttons-table/formeditassocbuttons-table.component'
import { FormEditAssocButtonDetailComponent } from './formeditassocbutton-detail/formeditassocbutton-detail.component'

import { FormFieldsTableComponent } from './formfields-table/formfields-table.component'
import { FormFieldDetailComponent } from './formfield-detail/formfield-detail.component'

import { FormFieldDatesTableComponent } from './formfielddates-table/formfielddates-table.component'
import { FormFieldDateDetailComponent } from './formfielddate-detail/formfielddate-detail.component'

import { FormFieldDateTimesTableComponent } from './formfielddatetimes-table/formfielddatetimes-table.component'
import { FormFieldDateTimeDetailComponent } from './formfielddatetime-detail/formfielddatetime-detail.component'

import { FormFieldFloat64sTableComponent } from './formfieldfloat64s-table/formfieldfloat64s-table.component'
import { FormFieldFloat64DetailComponent } from './formfieldfloat64-detail/formfieldfloat64-detail.component'

import { FormFieldIntsTableComponent } from './formfieldints-table/formfieldints-table.component'
import { FormFieldIntDetailComponent } from './formfieldint-detail/formfieldint-detail.component'

import { FormFieldSelectsTableComponent } from './formfieldselects-table/formfieldselects-table.component'
import { FormFieldSelectDetailComponent } from './formfieldselect-detail/formfieldselect-detail.component'

import { FormFieldStringsTableComponent } from './formfieldstrings-table/formfieldstrings-table.component'
import { FormFieldStringDetailComponent } from './formfieldstring-detail/formfieldstring-detail.component'

import { FormFieldTimesTableComponent } from './formfieldtimes-table/formfieldtimes-table.component'
import { FormFieldTimeDetailComponent } from './formfieldtime-detail/formfieldtime-detail.component'

import { FormGroupsTableComponent } from './formgroups-table/formgroups-table.component'
import { FormGroupDetailComponent } from './formgroup-detail/formgroup-detail.component'

import { FormSortAssocButtonsTableComponent } from './formsortassocbuttons-table/formsortassocbuttons-table.component'
import { FormSortAssocButtonDetailComponent } from './formsortassocbutton-detail/formsortassocbutton-detail.component'

import { OptionsTableComponent } from './options-table/options-table.component'
import { OptionDetailComponent } from './option-detail/option-detail.component'


@Injectable({
    providedIn: 'root'
})
export class RouteService {
    private routes: Routes = []

    constructor(private router: Router) { }

    public addRoutes(newRoutes: Routes): void {
        const existingRoutes = this.router.config
        this.routes = this.router.config

        for (let newRoute of newRoutes) {
            if (!existingRoutes.includes(newRoute)) {
                this.routes.push(newRoute)
            }
        }
        this.router.resetConfig(this.routes)
    }

    getPathRoot(): string {
        return 'github_com_fullstack_lang_gongform_go'
    }
    getTableOutlet(stackPath: string): string {
        return this.getPathRoot() + '_table' + '_' + stackPath
    }
    getEditorOutlet(stackPath: string): string {
        return this.getPathRoot() + '_editor' + '_' + stackPath
    }
    // insertion point for per gongstruct route/path getters
    getCheckBoxTablePath(): string {
        return this.getPathRoot() + '-checkboxs/:GONG__StackPath'
    }
    getCheckBoxTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCheckBoxTablePath(), component: CheckBoxsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getCheckBoxAdderPath(): string {
        return this.getPathRoot() + '-checkbox-adder/:GONG__StackPath'
    }
    getCheckBoxAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCheckBoxAdderPath(), component: CheckBoxDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCheckBoxAdderForUsePath(): string {
        return this.getPathRoot() + '-checkbox-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getCheckBoxAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCheckBoxAdderForUsePath(), component: CheckBoxDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCheckBoxDetailPath(): string {
        return this.getPathRoot() + '-checkbox-detail/:id/:GONG__StackPath'
    }
    getCheckBoxDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCheckBoxDetailPath(), component: CheckBoxDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormDivTablePath(): string {
        return this.getPathRoot() + '-formdivs/:GONG__StackPath'
    }
    getFormDivTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormDivTablePath(), component: FormDivsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormDivAdderPath(): string {
        return this.getPathRoot() + '-formdiv-adder/:GONG__StackPath'
    }
    getFormDivAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormDivAdderPath(), component: FormDivDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormDivAdderForUsePath(): string {
        return this.getPathRoot() + '-formdiv-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormDivAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormDivAdderForUsePath(), component: FormDivDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormDivDetailPath(): string {
        return this.getPathRoot() + '-formdiv-detail/:id/:GONG__StackPath'
    }
    getFormDivDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormDivDetailPath(), component: FormDivDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormEditAssocButtonTablePath(): string {
        return this.getPathRoot() + '-formeditassocbuttons/:GONG__StackPath'
    }
    getFormEditAssocButtonTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormEditAssocButtonTablePath(), component: FormEditAssocButtonsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormEditAssocButtonAdderPath(): string {
        return this.getPathRoot() + '-formeditassocbutton-adder/:GONG__StackPath'
    }
    getFormEditAssocButtonAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormEditAssocButtonAdderPath(), component: FormEditAssocButtonDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormEditAssocButtonAdderForUsePath(): string {
        return this.getPathRoot() + '-formeditassocbutton-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormEditAssocButtonAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormEditAssocButtonAdderForUsePath(), component: FormEditAssocButtonDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormEditAssocButtonDetailPath(): string {
        return this.getPathRoot() + '-formeditassocbutton-detail/:id/:GONG__StackPath'
    }
    getFormEditAssocButtonDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormEditAssocButtonDetailPath(), component: FormEditAssocButtonDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormFieldTablePath(): string {
        return this.getPathRoot() + '-formfields/:GONG__StackPath'
    }
    getFormFieldTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldTablePath(), component: FormFieldsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormFieldAdderPath(): string {
        return this.getPathRoot() + '-formfield-adder/:GONG__StackPath'
    }
    getFormFieldAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldAdderPath(), component: FormFieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormFieldAdderForUsePath(): string {
        return this.getPathRoot() + '-formfield-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormFieldAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldAdderForUsePath(), component: FormFieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormFieldDetailPath(): string {
        return this.getPathRoot() + '-formfield-detail/:id/:GONG__StackPath'
    }
    getFormFieldDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldDetailPath(), component: FormFieldDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormFieldDateTablePath(): string {
        return this.getPathRoot() + '-formfielddates/:GONG__StackPath'
    }
    getFormFieldDateTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldDateTablePath(), component: FormFieldDatesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormFieldDateAdderPath(): string {
        return this.getPathRoot() + '-formfielddate-adder/:GONG__StackPath'
    }
    getFormFieldDateAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldDateAdderPath(), component: FormFieldDateDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormFieldDateAdderForUsePath(): string {
        return this.getPathRoot() + '-formfielddate-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormFieldDateAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldDateAdderForUsePath(), component: FormFieldDateDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormFieldDateDetailPath(): string {
        return this.getPathRoot() + '-formfielddate-detail/:id/:GONG__StackPath'
    }
    getFormFieldDateDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldDateDetailPath(), component: FormFieldDateDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormFieldDateTimeTablePath(): string {
        return this.getPathRoot() + '-formfielddatetimes/:GONG__StackPath'
    }
    getFormFieldDateTimeTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldDateTimeTablePath(), component: FormFieldDateTimesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormFieldDateTimeAdderPath(): string {
        return this.getPathRoot() + '-formfielddatetime-adder/:GONG__StackPath'
    }
    getFormFieldDateTimeAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldDateTimeAdderPath(), component: FormFieldDateTimeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormFieldDateTimeAdderForUsePath(): string {
        return this.getPathRoot() + '-formfielddatetime-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormFieldDateTimeAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldDateTimeAdderForUsePath(), component: FormFieldDateTimeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormFieldDateTimeDetailPath(): string {
        return this.getPathRoot() + '-formfielddatetime-detail/:id/:GONG__StackPath'
    }
    getFormFieldDateTimeDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldDateTimeDetailPath(), component: FormFieldDateTimeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormFieldFloat64TablePath(): string {
        return this.getPathRoot() + '-formfieldfloat64s/:GONG__StackPath'
    }
    getFormFieldFloat64TableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldFloat64TablePath(), component: FormFieldFloat64sTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormFieldFloat64AdderPath(): string {
        return this.getPathRoot() + '-formfieldfloat64-adder/:GONG__StackPath'
    }
    getFormFieldFloat64AdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldFloat64AdderPath(), component: FormFieldFloat64DetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormFieldFloat64AdderForUsePath(): string {
        return this.getPathRoot() + '-formfieldfloat64-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormFieldFloat64AdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldFloat64AdderForUsePath(), component: FormFieldFloat64DetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormFieldFloat64DetailPath(): string {
        return this.getPathRoot() + '-formfieldfloat64-detail/:id/:GONG__StackPath'
    }
    getFormFieldFloat64DetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldFloat64DetailPath(), component: FormFieldFloat64DetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormFieldIntTablePath(): string {
        return this.getPathRoot() + '-formfieldints/:GONG__StackPath'
    }
    getFormFieldIntTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldIntTablePath(), component: FormFieldIntsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormFieldIntAdderPath(): string {
        return this.getPathRoot() + '-formfieldint-adder/:GONG__StackPath'
    }
    getFormFieldIntAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldIntAdderPath(), component: FormFieldIntDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormFieldIntAdderForUsePath(): string {
        return this.getPathRoot() + '-formfieldint-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormFieldIntAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldIntAdderForUsePath(), component: FormFieldIntDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormFieldIntDetailPath(): string {
        return this.getPathRoot() + '-formfieldint-detail/:id/:GONG__StackPath'
    }
    getFormFieldIntDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldIntDetailPath(), component: FormFieldIntDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormFieldSelectTablePath(): string {
        return this.getPathRoot() + '-formfieldselects/:GONG__StackPath'
    }
    getFormFieldSelectTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldSelectTablePath(), component: FormFieldSelectsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormFieldSelectAdderPath(): string {
        return this.getPathRoot() + '-formfieldselect-adder/:GONG__StackPath'
    }
    getFormFieldSelectAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldSelectAdderPath(), component: FormFieldSelectDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormFieldSelectAdderForUsePath(): string {
        return this.getPathRoot() + '-formfieldselect-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormFieldSelectAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldSelectAdderForUsePath(), component: FormFieldSelectDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormFieldSelectDetailPath(): string {
        return this.getPathRoot() + '-formfieldselect-detail/:id/:GONG__StackPath'
    }
    getFormFieldSelectDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldSelectDetailPath(), component: FormFieldSelectDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormFieldStringTablePath(): string {
        return this.getPathRoot() + '-formfieldstrings/:GONG__StackPath'
    }
    getFormFieldStringTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldStringTablePath(), component: FormFieldStringsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormFieldStringAdderPath(): string {
        return this.getPathRoot() + '-formfieldstring-adder/:GONG__StackPath'
    }
    getFormFieldStringAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldStringAdderPath(), component: FormFieldStringDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormFieldStringAdderForUsePath(): string {
        return this.getPathRoot() + '-formfieldstring-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormFieldStringAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldStringAdderForUsePath(), component: FormFieldStringDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormFieldStringDetailPath(): string {
        return this.getPathRoot() + '-formfieldstring-detail/:id/:GONG__StackPath'
    }
    getFormFieldStringDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldStringDetailPath(), component: FormFieldStringDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormFieldTimeTablePath(): string {
        return this.getPathRoot() + '-formfieldtimes/:GONG__StackPath'
    }
    getFormFieldTimeTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldTimeTablePath(), component: FormFieldTimesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormFieldTimeAdderPath(): string {
        return this.getPathRoot() + '-formfieldtime-adder/:GONG__StackPath'
    }
    getFormFieldTimeAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldTimeAdderPath(), component: FormFieldTimeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormFieldTimeAdderForUsePath(): string {
        return this.getPathRoot() + '-formfieldtime-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormFieldTimeAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldTimeAdderForUsePath(), component: FormFieldTimeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormFieldTimeDetailPath(): string {
        return this.getPathRoot() + '-formfieldtime-detail/:id/:GONG__StackPath'
    }
    getFormFieldTimeDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormFieldTimeDetailPath(), component: FormFieldTimeDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormGroupTablePath(): string {
        return this.getPathRoot() + '-formgroups/:GONG__StackPath'
    }
    getFormGroupTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormGroupTablePath(), component: FormGroupsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormGroupAdderPath(): string {
        return this.getPathRoot() + '-formgroup-adder/:GONG__StackPath'
    }
    getFormGroupAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormGroupAdderPath(), component: FormGroupDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormGroupAdderForUsePath(): string {
        return this.getPathRoot() + '-formgroup-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormGroupAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormGroupAdderForUsePath(), component: FormGroupDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormGroupDetailPath(): string {
        return this.getPathRoot() + '-formgroup-detail/:id/:GONG__StackPath'
    }
    getFormGroupDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormGroupDetailPath(), component: FormGroupDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getFormSortAssocButtonTablePath(): string {
        return this.getPathRoot() + '-formsortassocbuttons/:GONG__StackPath'
    }
    getFormSortAssocButtonTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormSortAssocButtonTablePath(), component: FormSortAssocButtonsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getFormSortAssocButtonAdderPath(): string {
        return this.getPathRoot() + '-formsortassocbutton-adder/:GONG__StackPath'
    }
    getFormSortAssocButtonAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormSortAssocButtonAdderPath(), component: FormSortAssocButtonDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormSortAssocButtonAdderForUsePath(): string {
        return this.getPathRoot() + '-formsortassocbutton-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getFormSortAssocButtonAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormSortAssocButtonAdderForUsePath(), component: FormSortAssocButtonDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getFormSortAssocButtonDetailPath(): string {
        return this.getPathRoot() + '-formsortassocbutton-detail/:id/:GONG__StackPath'
    }
    getFormSortAssocButtonDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getFormSortAssocButtonDetailPath(), component: FormSortAssocButtonDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getOptionTablePath(): string {
        return this.getPathRoot() + '-options/:GONG__StackPath'
    }
    getOptionTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getOptionTablePath(), component: OptionsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getOptionAdderPath(): string {
        return this.getPathRoot() + '-option-adder/:GONG__StackPath'
    }
    getOptionAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getOptionAdderPath(), component: OptionDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getOptionAdderForUsePath(): string {
        return this.getPathRoot() + '-option-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getOptionAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getOptionAdderForUsePath(), component: OptionDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getOptionDetailPath(): string {
        return this.getPathRoot() + '-option-detail/:id/:GONG__StackPath'
    }
    getOptionDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getOptionDetailPath(), component: OptionDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }



    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            // insertion point for all routes getter
            this.getCheckBoxTableRoute(stackPath),
            this.getCheckBoxAdderRoute(stackPath),
            this.getCheckBoxAdderForUseRoute(stackPath),
            this.getCheckBoxDetailRoute(stackPath),

            this.getFormDivTableRoute(stackPath),
            this.getFormDivAdderRoute(stackPath),
            this.getFormDivAdderForUseRoute(stackPath),
            this.getFormDivDetailRoute(stackPath),

            this.getFormEditAssocButtonTableRoute(stackPath),
            this.getFormEditAssocButtonAdderRoute(stackPath),
            this.getFormEditAssocButtonAdderForUseRoute(stackPath),
            this.getFormEditAssocButtonDetailRoute(stackPath),

            this.getFormFieldTableRoute(stackPath),
            this.getFormFieldAdderRoute(stackPath),
            this.getFormFieldAdderForUseRoute(stackPath),
            this.getFormFieldDetailRoute(stackPath),

            this.getFormFieldDateTableRoute(stackPath),
            this.getFormFieldDateAdderRoute(stackPath),
            this.getFormFieldDateAdderForUseRoute(stackPath),
            this.getFormFieldDateDetailRoute(stackPath),

            this.getFormFieldDateTimeTableRoute(stackPath),
            this.getFormFieldDateTimeAdderRoute(stackPath),
            this.getFormFieldDateTimeAdderForUseRoute(stackPath),
            this.getFormFieldDateTimeDetailRoute(stackPath),

            this.getFormFieldFloat64TableRoute(stackPath),
            this.getFormFieldFloat64AdderRoute(stackPath),
            this.getFormFieldFloat64AdderForUseRoute(stackPath),
            this.getFormFieldFloat64DetailRoute(stackPath),

            this.getFormFieldIntTableRoute(stackPath),
            this.getFormFieldIntAdderRoute(stackPath),
            this.getFormFieldIntAdderForUseRoute(stackPath),
            this.getFormFieldIntDetailRoute(stackPath),

            this.getFormFieldSelectTableRoute(stackPath),
            this.getFormFieldSelectAdderRoute(stackPath),
            this.getFormFieldSelectAdderForUseRoute(stackPath),
            this.getFormFieldSelectDetailRoute(stackPath),

            this.getFormFieldStringTableRoute(stackPath),
            this.getFormFieldStringAdderRoute(stackPath),
            this.getFormFieldStringAdderForUseRoute(stackPath),
            this.getFormFieldStringDetailRoute(stackPath),

            this.getFormFieldTimeTableRoute(stackPath),
            this.getFormFieldTimeAdderRoute(stackPath),
            this.getFormFieldTimeAdderForUseRoute(stackPath),
            this.getFormFieldTimeDetailRoute(stackPath),

            this.getFormGroupTableRoute(stackPath),
            this.getFormGroupAdderRoute(stackPath),
            this.getFormGroupAdderForUseRoute(stackPath),
            this.getFormGroupDetailRoute(stackPath),

            this.getFormSortAssocButtonTableRoute(stackPath),
            this.getFormSortAssocButtonAdderRoute(stackPath),
            this.getFormSortAssocButtonAdderForUseRoute(stackPath),
            this.getFormSortAssocButtonDetailRoute(stackPath),

            this.getOptionTableRoute(stackPath),
            this.getOptionAdderRoute(stackPath),
            this.getOptionAdderForUseRoute(stackPath),
            this.getOptionDetailRoute(stackPath),

        ])
    }
}
